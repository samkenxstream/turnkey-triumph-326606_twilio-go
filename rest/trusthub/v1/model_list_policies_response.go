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

// ListPoliciesResponse struct for ListPoliciesResponse
type ListPoliciesResponse struct {
	Meta    ListCustomerProfileResponseMeta `json:"meta,omitempty"`
	Results []TrusthubV1Policies            `json:"results,omitempty"`
}
