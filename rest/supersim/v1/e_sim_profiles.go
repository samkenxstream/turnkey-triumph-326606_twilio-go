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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateEsimProfile'
type CreateEsimProfileParams struct {
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` when the status of the eSIM Profile changes. At this stage of the eSIM Profile pilot, the a request to the URL will only be called when the ESimProfile resource changes from `reserving` to `available`.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// Identifier of the eUICC that will claim the eSIM Profile.
	Eid *string `json:"Eid,omitempty"`
}

func (params *CreateEsimProfileParams) SetCallbackMethod(CallbackMethod string) *CreateEsimProfileParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateEsimProfileParams) SetCallbackUrl(CallbackUrl string) *CreateEsimProfileParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateEsimProfileParams) SetEid(Eid string) *CreateEsimProfileParams {
	params.Eid = &Eid
	return params
}

// Order an eSIM Profile.
func (c *ApiService) CreateEsimProfile(params *CreateEsimProfileParams) (*SupersimV1EsimProfile, error) {
	path := "/v1/ESimProfiles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Eid != nil {
		data.Set("Eid", *params.Eid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1EsimProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch an eSIM Profile.
func (c *ApiService) FetchEsimProfile(Sid string) (*SupersimV1EsimProfile, error) {
	path := "/v1/ESimProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1EsimProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEsimProfile'
type ListEsimProfileParams struct {
	// List the eSIM Profiles that have been associated with an EId.
	Eid *string `json:"Eid,omitempty"`
	// Find the eSIM Profile resource related to a [Sim](https://www.twilio.com/docs/wireless/api/sim-resource) resource by providing the SIM SID. Will always return an array with either 1 or 0 records.
	SimSid *string `json:"SimSid,omitempty"`
	// List the eSIM Profiles that are in a given status.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEsimProfileParams) SetEid(Eid string) *ListEsimProfileParams {
	params.Eid = &Eid
	return params
}
func (params *ListEsimProfileParams) SetSimSid(SimSid string) *ListEsimProfileParams {
	params.SimSid = &SimSid
	return params
}
func (params *ListEsimProfileParams) SetStatus(Status string) *ListEsimProfileParams {
	params.Status = &Status
	return params
}
func (params *ListEsimProfileParams) SetPageSize(PageSize int) *ListEsimProfileParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEsimProfileParams) SetLimit(Limit int) *ListEsimProfileParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of EsimProfile records from the API. Request is executed immediately.
func (c *ApiService) PageEsimProfile(params *ListEsimProfileParams, pageToken, pageNumber string) (*ListEsimProfileResponse, error) {
	path := "/v1/ESimProfiles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Eid != nil {
		data.Set("Eid", *params.Eid)
	}
	if params != nil && params.SimSid != nil {
		data.Set("SimSid", *params.SimSid)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
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

	ps := &ListEsimProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists EsimProfile records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEsimProfile(params *ListEsimProfileParams) ([]SupersimV1EsimProfile, error) {
	response, err := c.StreamEsimProfile(params)
	if err != nil {
		return nil, err
	}

	records := make([]SupersimV1EsimProfile, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams EsimProfile records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEsimProfile(params *ListEsimProfileParams) (chan SupersimV1EsimProfile, error) {
	if params == nil {
		params = &ListEsimProfileParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEsimProfile(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan SupersimV1EsimProfile, 1)

	go func() {
		for response != nil {
			responseRecords := response.EsimProfiles
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListEsimProfileResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEsimProfileResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEsimProfileResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEsimProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
