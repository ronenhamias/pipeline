// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"emperror.dev/errors"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/banzaicloud/pipeline/.gen/pipeline"
)

type Client struct {
	grpcClient pb.ProcessClient
}

type Config struct {
	Address    string
	CACertFile string
	CertFile   string
	KeyFile    string
}

func NewClient(c Config) (*Client, error) {
	rootCAs := x509.NewCertPool()

	caCertPEM, err := ioutil.ReadFile(c.CACertFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read GRPC CA certificate")
	}

	if !rootCAs.AppendCertsFromPEM(caCertPEM) {
		return nil, errors.New("failed to append GRPC CA certificate")
	}

	clientCert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load GRPC client certificate")
	}

	tlsConfig := tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{clientCert},
	}

	creds := credentials.NewTLS(&tlsConfig)

	conn, err := grpc.Dial(c.Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewProcessClient(conn)

	return &Client{grpcClient: grpcClient}, nil
}

type ProcessEntry struct {
	ID           string
	ParentID     string
	OrgID        uint
	Name         string
	ResourceType ResourceType
	ResourceID   string
	Status       Status
	StartedAt    time.Time
	FinishedAt   *time.Time
}

type ProcessEvent struct {
	ProcessID string
	Name      string
	Log       string
	Timestamp time.Time
}

type ResourceType string
type Status string

const (
	Cluster ResourceType = "cluster"

	Running  Status = "running"
	Failed   Status = "failed"
	Finished Status = "finished"
)

func (c *Client) LogProcess(ctx context.Context, e ProcessEntry) error {

	pe := pb.ProcessEntry{
		Id:           e.ID,
		ParentId:     e.ParentID,
		OrgId:        uint32(e.OrgID),
		Name:         e.Name,
		ResourceType: string(e.ResourceType),
		ResourceId:   e.ResourceID,
		Status:       string(e.Status),
	}

	pe.StartedAt, _ = ptypes.TimestampProto(e.StartedAt)

	if e.FinishedAt != nil {
		finishedAt, _ := ptypes.TimestampProto(*e.FinishedAt)
		pe.FinishedAt = finishedAt
	}

	_, err := c.grpcClient.LogProcess(ctx, &pe)

	return err
}

func (c *Client) LogEvent(ctx context.Context, e ProcessEvent) error {

	pe := pb.ProcessEvent{
		ProcessId: e.ProcessID,
		Name:      e.Name,
		Log:       e.Log,
	}

	pe.Timestamp, _ = ptypes.TimestampProto(e.Timestamp)

	_, err := c.grpcClient.LogEvent(ctx, &pe)

	return err
}
