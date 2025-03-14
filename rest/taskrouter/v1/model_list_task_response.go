/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTaskResponse struct for ListTaskResponse
type ListTaskResponse struct {
	Meta  ListWorkspaceResponseMeta `json:"meta,omitempty"`
	Tasks []TaskrouterV1Task        `json:"tasks,omitempty"`
}
