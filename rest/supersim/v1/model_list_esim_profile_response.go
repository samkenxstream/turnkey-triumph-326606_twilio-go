/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEsimProfileResponse struct for ListEsimProfileResponse
type ListEsimProfileResponse struct {
	EsimProfiles []SupersimV1EsimProfile     `json:"esim_profiles,omitempty"`
	Meta         ListEsimProfileResponseMeta `json:"meta,omitempty"`
}
