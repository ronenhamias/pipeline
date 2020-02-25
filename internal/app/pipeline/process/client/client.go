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
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/banzaicloud/pipeline/.gen/pipeline"
)

type Client struct {
	grpcClient pb.ProcessClient
}

type Config struct {
	Address string
}

func NewClient(c Config) (*Client, error) {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewProcessClient(conn)

	return &Client{grpcClient: grpcClient}, nil
}

func (c *Client) Log(ctx context.Context, orgID uint, name string, resourceType string, resourceID string, status string, startedAt time.Time, finishedAt *time.Time) error {

	// activityInfo := activity.GetInfo(ctx)

	// pe := pb.ProcessEntry{
	// 	Id:       activityInfo.ActivityID,
	// 	ParentId: activityInfo.WorkflowExecution.ID,
	// 	OrgId:    orgID,
	// 	Name:     activityInfo.ActivityType.Name,
	// }

	pe := pb.ProcessEntry{
		Id:           "1234-5678",
		ParentId:     "9876-5432",
		OrgId:        uint32(orgID),
		Name:         name,
		ResourceType: resourceType,
		ResourceId:   resourceID,
		Status:       status,
	}

	pe.StartedAt, _ = ptypes.TimestampProto(startedAt)

	if finishedAt != nil {
		finishedAt, _ := ptypes.TimestampProto(*finishedAt)
		pe.FinishedAt = finishedAt
	}

	_, err := c.grpcClient.Log(ctx, &pe)

	return err
}
