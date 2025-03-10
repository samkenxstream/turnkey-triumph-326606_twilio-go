/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListBuildResponse struct for ListBuildResponse
type ListBuildResponse struct {
	Builds []ServerlessV1Build     `json:"builds,omitempty"`
	Meta   ListServiceResponseMeta `json:"meta,omitempty"`
}
