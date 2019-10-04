/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments.
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

// DeploymentScaleStatus struct for DeploymentScaleStatus
type DeploymentScaleStatus struct {
	CurrentReplicas int32  `json:"currentReplicas,omitempty"`
	DesiredReplicas int32  `json:"desiredReplicas,omitempty"`
	Message         string `json:"message,omitempty"`
}
