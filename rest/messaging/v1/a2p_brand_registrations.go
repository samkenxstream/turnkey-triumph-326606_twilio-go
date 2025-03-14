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

// Optional parameters for the method 'CreateBrandRegistrations'
type CreateBrandRegistrationsParams struct {
	// A2P Messaging Profile Bundle Sid.
	A2PProfileBundleSid *string `json:"A2PProfileBundleSid,omitempty"`
	// Type of brand being created. One of: \\\"STANDARD\\\", \\\"STARTER\\\". STARTER is for low volume, starter use cases. STANDARD is for all other use cases.
	BrandType *string `json:"BrandType,omitempty"`
	// Customer Profile Bundle Sid.
	CustomerProfileBundleSid *string `json:"CustomerProfileBundleSid,omitempty"`
	// A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided.
	Mock *bool `json:"Mock,omitempty"`
	// A flag to disable automatic secondary vetting for brands which it would otherwise be done.
	SkipAutomaticSecVet *bool `json:"SkipAutomaticSecVet,omitempty"`
}

func (params *CreateBrandRegistrationsParams) SetA2PProfileBundleSid(A2PProfileBundleSid string) *CreateBrandRegistrationsParams {
	params.A2PProfileBundleSid = &A2PProfileBundleSid
	return params
}
func (params *CreateBrandRegistrationsParams) SetBrandType(BrandType string) *CreateBrandRegistrationsParams {
	params.BrandType = &BrandType
	return params
}
func (params *CreateBrandRegistrationsParams) SetCustomerProfileBundleSid(CustomerProfileBundleSid string) *CreateBrandRegistrationsParams {
	params.CustomerProfileBundleSid = &CustomerProfileBundleSid
	return params
}
func (params *CreateBrandRegistrationsParams) SetMock(Mock bool) *CreateBrandRegistrationsParams {
	params.Mock = &Mock
	return params
}
func (params *CreateBrandRegistrationsParams) SetSkipAutomaticSecVet(SkipAutomaticSecVet bool) *CreateBrandRegistrationsParams {
	params.SkipAutomaticSecVet = &SkipAutomaticSecVet
	return params
}

func (c *ApiService) CreateBrandRegistrations(params *CreateBrandRegistrationsParams) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.A2PProfileBundleSid != nil {
		data.Set("A2PProfileBundleSid", *params.A2PProfileBundleSid)
	}
	if params != nil && params.BrandType != nil {
		data.Set("BrandType", *params.BrandType)
	}
	if params != nil && params.CustomerProfileBundleSid != nil {
		data.Set("CustomerProfileBundleSid", *params.CustomerProfileBundleSid)
	}
	if params != nil && params.Mock != nil {
		data.Set("Mock", fmt.Sprint(*params.Mock))
	}
	if params != nil && params.SkipAutomaticSecVet != nil {
		data.Set("SkipAutomaticSecVet", fmt.Sprint(*params.SkipAutomaticSecVet))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchBrandRegistrations(Sid string) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBrandRegistrations'
type ListBrandRegistrationsParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBrandRegistrationsParams) SetPageSize(PageSize int) *ListBrandRegistrationsParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBrandRegistrationsParams) SetLimit(Limit int) *ListBrandRegistrationsParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of BrandRegistrations records from the API. Request is executed immediately.
func (c *ApiService) PageBrandRegistrations(params *ListBrandRegistrationsParams, pageToken, pageNumber string) (*ListBrandRegistrationsResponse, error) {
	path := "/v1/a2p/BrandRegistrations"

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

	ps := &ListBrandRegistrationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BrandRegistrations records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBrandRegistrations(params *ListBrandRegistrationsParams) ([]MessagingV1BrandRegistrations, error) {
	response, err := c.StreamBrandRegistrations(params)
	if err != nil {
		return nil, err
	}

	records := make([]MessagingV1BrandRegistrations, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams BrandRegistrations records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBrandRegistrations(params *ListBrandRegistrationsParams) (chan MessagingV1BrandRegistrations, error) {
	if params == nil {
		params = &ListBrandRegistrationsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBrandRegistrations(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1BrandRegistrations, 1)

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
			if record, err = client.GetNext(c.baseURL, response, c.getNextListBrandRegistrationsResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBrandRegistrationsResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBrandRegistrationsResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBrandRegistrationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *ApiService) UpdateBrandRegistrations(Sid string) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
