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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListUserChannel'
type ListUserChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUserChannelParams) SetPageSize(PageSize int) *ListUserChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUserChannelParams) SetLimit(Limit int) *ListUserChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UserChannel records from the API. Request is executed immediately.
func (c *ApiService) PageUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams, pageToken, pageNumber string) (*ListUserChannelResponse, error) {
	path := "/v1/Services/{ServiceSid}/Users/{UserSid}/Channels"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListUserChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UserChannel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams) ([]ChatV1UserChannel, error) {
	response, err := c.StreamUserChannel(ServiceSid, UserSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ChatV1UserChannel, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams UserChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams) (chan ChatV1UserChannel, error) {
	if params == nil {
		params = &ListUserChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUserChannel(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ChatV1UserChannel, 1)

	go func() {
		for response != nil {
			responseRecords := response.Channels
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListUserChannelResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUserChannelResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUserChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
