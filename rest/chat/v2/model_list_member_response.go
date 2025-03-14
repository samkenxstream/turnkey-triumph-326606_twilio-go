/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListMemberResponse struct for ListMemberResponse
type ListMemberResponse struct {
	Members []ChatV2Member             `json:"members,omitempty"`
	Meta    ListCredentialResponseMeta `json:"meta,omitempty"`
}
