/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'CreateDialingPermissionsCountryBulkUpdate'
type CreateDialingPermissionsCountryBulkUpdateParams struct {
	// URL encoded JSON array of update objects. example : `[ { \\\"iso_code\\\": \\\"GB\\\", \\\"low_risk_numbers_enabled\\\": \\\"true\\\", \\\"high_risk_special_numbers_enabled\\\":\\\"true\\\", \\\"high_risk_tollfraud_numbers_enabled\\\": \\\"false\\\" } ]`
	UpdateRequest *string `json:"UpdateRequest,omitempty"`
}

func (params *CreateDialingPermissionsCountryBulkUpdateParams) SetUpdateRequest(UpdateRequest string) *CreateDialingPermissionsCountryBulkUpdateParams {
	params.UpdateRequest = &UpdateRequest
	return params
}

// Create a bulk update request to change voice dialing country permissions of one or more countries identified by the corresponding [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
func (c *ApiService) CreateDialingPermissionsCountryBulkUpdate(params *CreateDialingPermissionsCountryBulkUpdateParams) (*VoiceV1DialingPermissionsCountryBulkUpdate, error) {
	path := "/v1/DialingPermissions/BulkCountryUpdates"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UpdateRequest != nil {
		data.Set("UpdateRequest", *params.UpdateRequest)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1DialingPermissionsCountryBulkUpdate{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
