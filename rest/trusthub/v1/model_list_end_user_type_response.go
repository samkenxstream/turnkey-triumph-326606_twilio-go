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

// ListEndUserTypeResponse struct for ListEndUserTypeResponse
type ListEndUserTypeResponse struct {
	EndUserTypes []TrusthubV1EndUserType         `json:"end_user_types,omitempty"`
	Meta         ListCustomerProfileResponseMeta `json:"meta,omitempty"`
}
