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

// CreateUpdateDeploymentResponse struct for CreateUpdateDeploymentResponse
type CreateUpdateDeploymentResponse struct {
	ReleaseName string `json:"releaseName,omitempty"`
	// deployment notes in base64 encoded format
	Notes string `json:"notes,omitempty"`
}
