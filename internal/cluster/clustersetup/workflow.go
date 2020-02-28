// Copyright © 2019 Banzai Cloud
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

package clustersetup

import (
	"context"
	"time"

	"emperror.dev/errors"
	"go.uber.org/cadence"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"

	processClient "github.com/banzaicloud/pipeline/internal/app/pipeline/process/client"
)

// WorkflowName can be used to reference the cluster setup workflow.
const WorkflowName = "cluster-setup"

// Workflow orchestrates the post-creation cluster setup flow.
type Workflow struct {
	// InstallInit
	InstallInitManifest bool

	ProcessLogger *processClient.Client
}

// WorkflowInput is the input for a cluster setup workflow.
type WorkflowInput struct {
	// Kubernetes cluster config secret ID.
	ConfigSecretID string

	// Cluster information
	Cluster      Cluster
	Organization Organization

	NodePoolLabels map[string]map[string]string
}

// Cluster represents a Kubernetes cluster.
type Cluster struct {
	ID           uint
	UID          string
	Name         string
	Distribution string
}

// Organization contains information about the organization a cluster belongs to.
type Organization struct {
	ID   uint
	Name string
}

// Execute executes the cluster setup workflow.
func (w Workflow) Execute(ctx workflow.Context, input WorkflowInput) error {
	// Default timeouts and retries
	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: 20 * time.Minute,
		StartToCloseTimeout:    30 * time.Minute,
		WaitForCancellation:    true,
		RetryPolicy: &cadence.RetryPolicy{
			InitialInterval:    2 * time.Second,
			BackoffCoefficient: 1.5,
			MaximumInterval:    30 * time.Second,
			MaximumAttempts:    30,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	process := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {

		winfo := workflow.GetInfo(ctx)
		pe := processClient.ProcessEntry{
			ID:           winfo.WorkflowExecution.ID,
			Name:         winfo.WorkflowType.Name,
			OrgID:        input.Organization.ID,
			ResourceType: processClient.Cluster,
			ResourceID:   input.Cluster.UID,
			StartedAt:    workflow.Now(ctx),
			Status:       processClient.Running,
		}

		err := w.ProcessLogger.LogProcess(context.Background(), pe)
		if err != nil {
			workflow.GetLogger(ctx).Warn("failed to write process log", zap.Error(err))
		}

		return pe
	})

	defer func() {
		workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {

			var pe processClient.ProcessEntry
			err := process.Get(&pe)
			if err != nil {
				workflow.GetLogger(ctx).Warn("failed to get start process log", zap.Error(err))
				return nil
			}

			finishedAt := workflow.Now(ctx)
			pe.FinishedAt = &finishedAt
			pe.Status = processClient.Finished // TODO

			err = w.ProcessLogger.LogProcess(context.Background(), pe)
			if err != nil {
				workflow.GetLogger(ctx).Warn("failed to write process log", zap.Error(err))
			}

			return nil
		})
	}()

	// Install the cluster manifest to the cluster (if configured)
	if w.InstallInitManifest {
		activityInput := InitManifestActivityInput{
			ConfigSecretID: input.ConfigSecretID,
			Cluster:        input.Cluster,
			Organization:   input.Organization,
		}

		err := workflow.ExecuteActivity(ctx, InitManifestActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	{
		activityInput := CreatePipelineNamespaceActivityInput{
			ConfigSecretID: input.ConfigSecretID,
		}

		err := workflow.ExecuteActivity(ctx, CreatePipelineNamespaceActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	{
		activityInput := LabelKubeSystemNamespaceActivityInput{
			ConfigSecretID: input.ConfigSecretID,
		}

		err := workflow.ExecuteActivity(ctx, LabelKubeSystemNamespaceActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	{
		activityInput := InstallTillerActivityInput{
			ConfigSecretID: input.ConfigSecretID,
			Distribution:   input.Cluster.Distribution,
		}

		err := workflow.ExecuteActivity(ctx, InstallTillerActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	{
		activityInput := InstallTillerWaitActivityInput{
			ConfigSecretID: input.ConfigSecretID,
		}

		aop := activityOptions
		aop.HeartbeatTimeout = 1 * time.Minute
		ctx := workflow.WithActivityOptions(ctx, aop)

		err := workflow.ExecuteActivity(ctx, InstallTillerWaitActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			if cadence.IsTimeoutError(err) {
				return errors.New(
					"Cluster setup failed because Tiller couldn't start. " +
						"Usually this happens when worker nodes are not able to join the cluster. " +
						"Check your network settings to make sure the worker nodes can communicate with the Kubernetes API server.",
				)
			}

			return err
		}
	}

	{
		activityInput := InstallNodePoolLabelSetOperatorActivityInput{
			ClusterID: input.Cluster.ID,
		}

		err := workflow.ExecuteActivity(ctx, InstallNodePoolLabelSetOperatorActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	{
		activityInput := ConfigureNodePoolLabelsActivityInput{
			ConfigSecretID: input.ConfigSecretID,
			Labels:         input.NodePoolLabels,
		}

		err := workflow.ExecuteActivity(ctx, ConfigureNodePoolLabelsActivityName, activityInput).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
