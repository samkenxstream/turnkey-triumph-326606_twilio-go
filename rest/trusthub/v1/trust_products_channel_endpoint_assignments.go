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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateTrustProductChannelEndpointAssignment'
type CreateTrustProductChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType *string `json:"ChannelEndpointType,omitempty"`
}

func (params *CreateTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *CreateTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *CreateTrustProductChannelEndpointAssignmentParams) SetChannelEndpointType(ChannelEndpointType string) *CreateTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointType = &ChannelEndpointType
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateTrustProductChannelEndpointAssignment(TrustProductSid string, params *CreateTrustProductChannelEndpointAssignmentParams) (*TrusthubV1TrustProductChannelEndpointAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointType != nil {
		data.Set("ChannelEndpointType", *params.ChannelEndpointType)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteTrustProductChannelEndpointAssignment(TrustProductSid string, Sid string) error {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchTrustProductChannelEndpointAssignment(TrustProductSid string, Sid string) (*TrusthubV1TrustProductChannelEndpointAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductChannelEndpointAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProductChannelEndpointAssignment'
type ListTrustProductChannelEndpointAssignmentParams struct {
	// The SID of an channel endpoint
	ChannelEndpointSid *string `json:"ChannelEndpointSid,omitempty"`
	// comma separated list of channel endpoint sids
	ChannelEndpointSids *string `json:"ChannelEndpointSids,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSid(ChannelEndpointSid string) *ListTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSid = &ChannelEndpointSid
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetChannelEndpointSids(ChannelEndpointSids string) *ListTrustProductChannelEndpointAssignmentParams {
	params.ChannelEndpointSids = &ChannelEndpointSids
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetPageSize(PageSize int) *ListTrustProductChannelEndpointAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrustProductChannelEndpointAssignmentParams) SetLimit(Limit int) *ListTrustProductChannelEndpointAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrustProductChannelEndpointAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams, pageToken, pageNumber string) (*ListTrustProductChannelEndpointAssignmentResponse, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments"

	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelEndpointSid != nil {
		data.Set("ChannelEndpointSid", *params.ChannelEndpointSid)
	}
	if params != nil && params.ChannelEndpointSids != nil {
		data.Set("ChannelEndpointSids", *params.ChannelEndpointSids)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrustProductChannelEndpointAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams) ([]TrusthubV1TrustProductChannelEndpointAssignment, error) {
	response, err := c.StreamTrustProductChannelEndpointAssignment(TrustProductSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]TrusthubV1TrustProductChannelEndpointAssignment, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams TrustProductChannelEndpointAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProductChannelEndpointAssignment(TrustProductSid string, params *ListTrustProductChannelEndpointAssignmentParams) (chan TrusthubV1TrustProductChannelEndpointAssignment, error) {
	if params == nil {
		params = &ListTrustProductChannelEndpointAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTrustProductChannelEndpointAssignment(TrustProductSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1TrustProductChannelEndpointAssignment, 1)

	go func() {
		for response != nil {
			responseRecords := response.Results
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListTrustProductChannelEndpointAssignmentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTrustProductChannelEndpointAssignmentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTrustProductChannelEndpointAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductChannelEndpointAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
