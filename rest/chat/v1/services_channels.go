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

// Optional parameters for the method 'CreateChannel'
type CreateChannelParams struct {
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The visibility of the channel. Can be: `public` or `private` and defaults to `public`.
	Type *string `json:"Type,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. This value must be 64 characters or less in length and be unique within the Service.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateChannelParams) SetAttributes(Attributes string) *CreateChannelParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateChannelParams) SetFriendlyName(FriendlyName string) *CreateChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateChannelParams) SetType(Type string) *CreateChannelParams {
	params.Type = &Type
	return params
}
func (params *CreateChannelParams) SetUniqueName(UniqueName string) *CreateChannelParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateChannel(ServiceSid string, params *CreateChannelParams) (*ChatV1Channel, error) {
	path := "/v1/Services/{ServiceSid}/Channels"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteChannel(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchChannel(ServiceSid string, Sid string) (*ChatV1Channel, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannel'
type ListChannelParams struct {
	// The visibility of the Channels to read. Can be: `public` or `private` and defaults to `public`.
	Type *[]string `json:"Type,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListChannelParams) SetType(Type []string) *ListChannelParams {
	params.Type = &Type
	return params
}
func (params *ListChannelParams) SetPageSize(PageSize int) *ListChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListChannelParams) SetLimit(Limit int) *ListChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Channel records from the API. Request is executed immediately.
func (c *ApiService) PageChannel(ServiceSid string, params *ListChannelParams, pageToken, pageNumber string) (*ListChannelResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Type != nil {
		for _, item := range *params.Type {
			data.Add("Type", item)
		}
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

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Channel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChannel(ServiceSid string, params *ListChannelParams) ([]ChatV1Channel, error) {
	response, errors := c.StreamChannel(ServiceSid, params)

	records := make([]ChatV1Channel, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Channel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChannel(ServiceSid string, params *ListChannelParams) (chan ChatV1Channel, chan error) {
	if params == nil {
		params = &ListChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ChatV1Channel, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageChannel(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamChannel(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamChannel(response *ListChannelResponse, params *ListChannelParams, recordChannel chan ChatV1Channel, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Channels
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListChannelResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListChannelResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateChannel'
type UpdateChannelParams struct {
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. This value must be 64 characters or less in length and be unique within the Service.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateChannelParams) SetAttributes(Attributes string) *UpdateChannelParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateChannelParams) SetFriendlyName(FriendlyName string) *UpdateChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateChannelParams) SetUniqueName(UniqueName string) *UpdateChannelParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateChannel(ServiceSid string, Sid string, params *UpdateChannelParams) (*ChatV1Channel, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
