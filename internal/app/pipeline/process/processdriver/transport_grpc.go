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

package processdriver

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/banzaicloud/pipeline/.gen/pipeline"
	"github.com/banzaicloud/pipeline/internal/app/pipeline/process"
)

type GRPCAdapter struct {
	service process.Service
}

func (g GRPCAdapter) Log(ctx context.Context, pe *pb.ProcessEntry) (*pb.ProcessEntryResponse, error) {

	p := process.Process{
		ID:           pe.GetId(),
		ParentID:     pe.GetParentId(),
		OrgID:        uint(pe.GetOrgId()),
		Name:         pe.GetName(),
		ResourceType: pe.GetResourceType(),
		ResourceID:   pe.GetResourceId(),
		Status:       pe.GetStatus(),
	}

	p.StartedAt, _ = ptypes.Timestamp(pe.GetStartedAt())

	if pe.GetFinishedAt() != nil {
		finishedAt, _ := ptypes.Timestamp(pe.GetFinishedAt())
		p.FinishedAt = &finishedAt
	}

	_, err := g.service.Log(ctx, p)
	if err != nil {
		return nil, err
	}

	return &pb.ProcessEntryResponse{}, nil
}

// RegisterGRPCHandlers mounts all of the service endpoints into a grpc.Server.
func RegisterGRPCHandlers(service process.Service, grpcServer *grpc.Server) {
	pb.RegisterProcessServer(grpcServer, GRPCAdapter{service: service})
}
