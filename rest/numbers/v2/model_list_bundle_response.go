/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListBundleResponse struct for ListBundleResponse
type ListBundleResponse struct {
	Meta    ListBundleResponseMeta `json:"meta,omitempty"`
	Results []NumbersV2Bundle      `json:"results,omitempty"`
}
