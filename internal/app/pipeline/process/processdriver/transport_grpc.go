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

	"github.com/banzaicloud/pipeline/.gen/pipeline"
	"github.com/banzaicloud/pipeline/internal/app/pipeline/process"
	"google.golang.org/grpc"
)

type GRPCAdapter struct {
	service process.Service
}

func (g GRPCAdapter) Log(context.Context, *pipeline.ProcessEntry) (*pipeline.ProcessEntryResponse, error) {
	return &pipeline.ProcessEntryResponse{}, nil
}

// RegisterGRPCHandlers mounts all of the service endpoints into an grpc.Server.
func RegisterGRPCHandlers(service process.Service, grpcServer *grpc.Server) {
	pipeline.RegisterProcessServer(grpcServer, GRPCAdapter{service: service})
}
