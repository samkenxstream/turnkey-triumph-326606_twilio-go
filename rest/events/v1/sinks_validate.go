/*
 * Twilio - Events
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
	"strings"
)

// Optional parameters for the method 'CreateSinkValidate'
type CreateSinkValidateParams struct {
	// A 34 character string that uniquely identifies the test event for a Sink being validated.
	TestId *string `json:"TestId,omitempty"`
}

func (params *CreateSinkValidateParams) SetTestId(TestId string) *CreateSinkValidateParams {
	params.TestId = &TestId
	return params
}

// Validate that a test event for a Sink was received.
func (c *ApiService) CreateSinkValidate(Sid string, params *CreateSinkValidateParams) (*EventsV1SinkValidate, error) {
	path := "/v1/Sinks/{Sid}/Validate"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TestId != nil {
		data.Set("TestId", *params.TestId)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SinkValidate{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
