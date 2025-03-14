/*
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateBrandVetting'
type CreateBrandVettingParams struct {
	// The unique ID of the vetting
	VettingId *string `json:"VettingId,omitempty"`
	// The third-party provider of the vettings to create .
	VettingProvider *string `json:"VettingProvider,omitempty"`
}

func (params *CreateBrandVettingParams) SetVettingId(VettingId string) *CreateBrandVettingParams {
	params.VettingId = &VettingId
	return params
}
func (params *CreateBrandVettingParams) SetVettingProvider(VettingProvider string) *CreateBrandVettingParams {
	params.VettingProvider = &VettingProvider
	return params
}

func (c *ApiService) CreateBrandVetting(BrandSid string, params *CreateBrandVettingParams) (*MessagingV1BrandVetting, error) {
	path := "/v1/a2p/BrandRegistrations/{BrandSid}/Vettings"
	path = strings.Replace(path, "{"+"BrandSid"+"}", BrandSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.VettingId != nil {
		data.Set("VettingId", *params.VettingId)
	}
	if params != nil && params.VettingProvider != nil {
		data.Set("VettingProvider", *params.VettingProvider)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandVetting{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchBrandVetting(BrandSid string, BrandVettingSid string) (*MessagingV1BrandVetting, error) {
	path := "/v1/a2p/BrandRegistrations/{BrandSid}/Vettings/{BrandVettingSid}"
	path = strings.Replace(path, "{"+"BrandSid"+"}", BrandSid, -1)
	path = strings.Replace(path, "{"+"BrandVettingSid"+"}", BrandVettingSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandVetting{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBrandVetting'
type ListBrandVettingParams struct {
	// The third-party provider of the vettings to read
	VettingProvider *string `json:"VettingProvider,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBrandVettingParams) SetVettingProvider(VettingProvider string) *ListBrandVettingParams {
	params.VettingProvider = &VettingProvider
	return params
}
func (params *ListBrandVettingParams) SetPageSize(PageSize int) *ListBrandVettingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBrandVettingParams) SetLimit(Limit int) *ListBrandVettingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of BrandVetting records from the API. Request is executed immediately.
func (c *ApiService) PageBrandVetting(BrandSid string, params *ListBrandVettingParams, pageToken, pageNumber string) (*ListBrandVettingResponse, error) {
	path := "/v1/a2p/BrandRegistrations/{BrandSid}/Vettings"

	path = strings.Replace(path, "{"+"BrandSid"+"}", BrandSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.VettingProvider != nil {
		data.Set("VettingProvider", *params.VettingProvider)
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

	ps := &ListBrandVettingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BrandVetting records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBrandVetting(BrandSid string, params *ListBrandVettingParams) ([]MessagingV1BrandVetting, error) {
	response, err := c.StreamBrandVetting(BrandSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]MessagingV1BrandVetting, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams BrandVetting records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBrandVetting(BrandSid string, params *ListBrandVettingParams) (chan MessagingV1BrandVetting, error) {
	if params == nil {
		params = &ListBrandVettingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBrandVetting(BrandSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1BrandVetting, 1)

	go func() {
		for response != nil {
			responseRecords := response.Data
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListBrandVettingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBrandVettingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBrandVettingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBrandVettingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
