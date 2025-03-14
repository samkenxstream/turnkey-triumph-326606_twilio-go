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

// ListTaskChannelResponse struct for ListTaskChannelResponse
type ListTaskChannelResponse struct {
	Channels []TaskrouterV1TaskChannel `json:"channels,omitempty"`
	Meta     ListWorkspaceResponseMeta `json:"meta,omitempty"`
}
