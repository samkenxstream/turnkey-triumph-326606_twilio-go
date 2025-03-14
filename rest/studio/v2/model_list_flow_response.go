/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFlowResponse struct for ListFlowResponse
type ListFlowResponse struct {
	Flows []StudioV2Flow       `json:"flows,omitempty"`
	Meta  ListFlowResponseMeta `json:"meta,omitempty"`
}
