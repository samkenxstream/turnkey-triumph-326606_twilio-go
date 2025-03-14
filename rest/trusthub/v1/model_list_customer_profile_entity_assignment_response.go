/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCustomerProfileEntityAssignmentResponse struct for ListCustomerProfileEntityAssignmentResponse
type ListCustomerProfileEntityAssignmentResponse struct {
	Meta    ListCustomerProfileResponseMeta             `json:"meta,omitempty"`
	Results []TrusthubV1CustomerProfileEntityAssignment `json:"results,omitempty"`
}
