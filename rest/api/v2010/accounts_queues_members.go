/*
 * Twilio - Api
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

// Optional parameters for the method 'FetchMember'
type FetchMemberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchMemberParams) SetPathAccountSid(PathAccountSid string) *FetchMemberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific member from the queue
func (c *ApiService) FetchMember(QueueSid string, CallSid string, params *FetchMemberParams) (*ApiV2010Member, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"QueueSid"+"}", QueueSid, -1)
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Member{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMember'
type ListMemberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMemberParams) SetPathAccountSid(PathAccountSid string) *ListMemberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListMemberParams) SetPageSize(PageSize int) *ListMemberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMemberParams) SetLimit(Limit int) *ListMemberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Member records from the API. Request is executed immediately.
func (c *ApiService) PageMember(QueueSid string, params *ListMemberParams, pageToken, pageNumber string) (*ListMemberResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"QueueSid"+"}", QueueSid, -1)

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

	ps := &ListMemberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Member records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMember(QueueSid string, params *ListMemberParams) ([]ApiV2010Member, error) {
	response, err := c.StreamMember(QueueSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ApiV2010Member, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Member records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMember(QueueSid string, params *ListMemberParams) (chan ApiV2010Member, error) {
	if params == nil {
		params = &ListMemberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMember(QueueSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010Member, 1)

	go func() {
		for response != nil {
			responseRecords := response.QueueMembers
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListMemberResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMemberResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMemberResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMemberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMember'
type UpdateMemberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How to pass the update request data. Can be `GET` or `POST` and the default is `POST`. `POST` sends the data as encoded form data and `GET` sends the data as query parameters.
	Method *string `json:"Method,omitempty"`
	// The absolute URL of the Queue resource.
	Url *string `json:"Url,omitempty"`
}

func (params *UpdateMemberParams) SetPathAccountSid(PathAccountSid string) *UpdateMemberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateMemberParams) SetMethod(Method string) *UpdateMemberParams {
	params.Method = &Method
	return params
}
func (params *UpdateMemberParams) SetUrl(Url string) *UpdateMemberParams {
	params.Url = &Url
	return params
}

// Dequeue a member from a queue and have the member&#39;s call begin executing the TwiML document at that URL
func (c *ApiService) UpdateMember(QueueSid string, CallSid string, params *UpdateMemberParams) (*ApiV2010Member, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"QueueSid"+"}", QueueSid, -1)
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Method != nil {
		data.Set("Method", *params.Method)
	}
	if params != nil && params.Url != nil {
		data.Set("Url", *params.Url)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Member{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
