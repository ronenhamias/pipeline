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

// PostHooks struct for PostHooks
type PostHooks struct {
	InstallLogging       LoggingPostHookInstallLogging         `json:"InstallLogging,omitempty"`
	PostHookFunctionName map[string]interface{}                `json:"PostHookFunctionName,omitempty"`
	InstallServiceMesh   ServiceMeshPostHookInstallServiceMesh `json:"InstallServiceMesh,omitempty"`
}
