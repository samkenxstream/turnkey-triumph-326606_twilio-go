/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListQueryResponse struct for ListQueryResponse
type ListQueryResponse struct {
	Meta    ListAssistantResponseMeta `json:"meta,omitempty"`
	Queries []AutopilotV1Query        `json:"queries,omitempty"`
}
