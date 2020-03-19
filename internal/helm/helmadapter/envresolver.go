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

package helmadapter

import (
	"context"

	"emperror.dev/errors"
	"k8s.io/helm/pkg/helm/environment"

	"github.com/banzaicloud/pipeline/internal/helm"
	legacyHelm "github.com/banzaicloud/pipeline/src/helm"
)

//  envGenerator intermediary adapter component for handling legacy helm env generation
type envGenerator struct {
}

func NewEnvGenerator() helm.EnvResolver {
	return envGenerator{}
}

// GenerateHelmRepoEnv wraps the legacy method to gain control over the behaviour
func (envGenerator) GenerateHelmRepoEnv(orgName string) environment.EnvSettings {

	// ignoring the orgName
	return legacyHelm.GenerateHelmRepoEnv(helm.PlatformHelmEnv)
}

func (envGenerator) ResolveHelmEnv(ctx context.Context, organizationID uint) (helm.HelmEnv, error) {
	return helm.HelmEnv{}, errors.New("no-op method, invalid usage")
}

func (envGenerator) ResolvePlatformEnv(ctx context.Context) (helm.HelmEnv, error) {
	return helm.HelmEnv{}, errors.New("no-op method, invalid usage")
}