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

// PkeOnAzureClusterNetwork struct for PkeOnAzureClusterNetwork
type PkeOnAzureClusterNetwork struct {
	Name string `json:"name,omitempty"`
	// When referencing an already existing virtual network this field does not need to be specified.
	Cidr string `json:"cidr,omitempty"`
}
