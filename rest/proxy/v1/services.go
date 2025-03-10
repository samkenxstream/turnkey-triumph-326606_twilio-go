/*
 * Twilio - Proxy
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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// The URL we should call when the interaction status changes.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
	ChatInstanceSid *string `json:"ChatInstanceSid,omitempty"`
	// The default `ttl` value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value.
	DefaultTtl *int `json:"DefaultTtl,omitempty"`
	// Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America.
	GeoMatchLevel *string `json:"GeoMatchLevel,omitempty"`
	// The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
	InterceptCallbackUrl *string `json:"InterceptCallbackUrl,omitempty"`
	// The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
	NumberSelectionBehavior *string `json:"NumberSelectionBehavior,omitempty"`
	// The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
	OutOfSessionCallbackUrl *string `json:"OutOfSessionCallbackUrl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateServiceParams) SetCallbackUrl(CallbackUrl string) *CreateServiceParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateServiceParams) SetChatInstanceSid(ChatInstanceSid string) *CreateServiceParams {
	params.ChatInstanceSid = &ChatInstanceSid
	return params
}
func (params *CreateServiceParams) SetDefaultTtl(DefaultTtl int) *CreateServiceParams {
	params.DefaultTtl = &DefaultTtl
	return params
}
func (params *CreateServiceParams) SetGeoMatchLevel(GeoMatchLevel string) *CreateServiceParams {
	params.GeoMatchLevel = &GeoMatchLevel
	return params
}
func (params *CreateServiceParams) SetInterceptCallbackUrl(InterceptCallbackUrl string) *CreateServiceParams {
	params.InterceptCallbackUrl = &InterceptCallbackUrl
	return params
}
func (params *CreateServiceParams) SetNumberSelectionBehavior(NumberSelectionBehavior string) *CreateServiceParams {
	params.NumberSelectionBehavior = &NumberSelectionBehavior
	return params
}
func (params *CreateServiceParams) SetOutOfSessionCallbackUrl(OutOfSessionCallbackUrl string) *CreateServiceParams {
	params.OutOfSessionCallbackUrl = &OutOfSessionCallbackUrl
	return params
}
func (params *CreateServiceParams) SetUniqueName(UniqueName string) *CreateServiceParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new Service for Twilio Proxy
func (c *ApiService) CreateService(params *CreateServiceParams) (*ProxyV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.ChatInstanceSid != nil {
		data.Set("ChatInstanceSid", *params.ChatInstanceSid)
	}
	if params != nil && params.DefaultTtl != nil {
		data.Set("DefaultTtl", fmt.Sprint(*params.DefaultTtl))
	}
	if params != nil && params.GeoMatchLevel != nil {
		data.Set("GeoMatchLevel", *params.GeoMatchLevel)
	}
	if params != nil && params.InterceptCallbackUrl != nil {
		data.Set("InterceptCallbackUrl", *params.InterceptCallbackUrl)
	}
	if params != nil && params.NumberSelectionBehavior != nil {
		data.Set("NumberSelectionBehavior", *params.NumberSelectionBehavior)
	}
	if params != nil && params.OutOfSessionCallbackUrl != nil {
		data.Set("OutOfSessionCallbackUrl", *params.OutOfSessionCallbackUrl)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Service.
func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
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

// Fetch a specific Service.
func (c *ApiService) FetchService(Sid string) (*ProxyV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParams) SetLimit(Limit int) *ListServiceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken, pageNumber string) (*ListServiceResponse, error) {
	path := "/v1/Services"

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

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]ProxyV1Service, error) {
	response, err := c.StreamService(params)
	if err != nil {
		return nil, err
	}

	records := make([]ProxyV1Service, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan ProxyV1Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ProxyV1Service, 1)

	go func() {
		for response != nil {
			responseRecords := response.Services
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListServiceResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	// The URL we should call when the interaction status changes.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
	ChatInstanceSid *string `json:"ChatInstanceSid,omitempty"`
	// The default `ttl` value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value.
	DefaultTtl *int `json:"DefaultTtl,omitempty"`
	// Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America.
	GeoMatchLevel *string `json:"GeoMatchLevel,omitempty"`
	// The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
	InterceptCallbackUrl *string `json:"InterceptCallbackUrl,omitempty"`
	// The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
	NumberSelectionBehavior *string `json:"NumberSelectionBehavior,omitempty"`
	// The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
	OutOfSessionCallbackUrl *string `json:"OutOfSessionCallbackUrl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateServiceParams) SetCallbackUrl(CallbackUrl string) *UpdateServiceParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *UpdateServiceParams) SetChatInstanceSid(ChatInstanceSid string) *UpdateServiceParams {
	params.ChatInstanceSid = &ChatInstanceSid
	return params
}
func (params *UpdateServiceParams) SetDefaultTtl(DefaultTtl int) *UpdateServiceParams {
	params.DefaultTtl = &DefaultTtl
	return params
}
func (params *UpdateServiceParams) SetGeoMatchLevel(GeoMatchLevel string) *UpdateServiceParams {
	params.GeoMatchLevel = &GeoMatchLevel
	return params
}
func (params *UpdateServiceParams) SetInterceptCallbackUrl(InterceptCallbackUrl string) *UpdateServiceParams {
	params.InterceptCallbackUrl = &InterceptCallbackUrl
	return params
}
func (params *UpdateServiceParams) SetNumberSelectionBehavior(NumberSelectionBehavior string) *UpdateServiceParams {
	params.NumberSelectionBehavior = &NumberSelectionBehavior
	return params
}
func (params *UpdateServiceParams) SetOutOfSessionCallbackUrl(OutOfSessionCallbackUrl string) *UpdateServiceParams {
	params.OutOfSessionCallbackUrl = &OutOfSessionCallbackUrl
	return params
}
func (params *UpdateServiceParams) SetUniqueName(UniqueName string) *UpdateServiceParams {
	params.UniqueName = &UniqueName
	return params
}

// Update a specific Service.
func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*ProxyV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.ChatInstanceSid != nil {
		data.Set("ChatInstanceSid", *params.ChatInstanceSid)
	}
	if params != nil && params.DefaultTtl != nil {
		data.Set("DefaultTtl", fmt.Sprint(*params.DefaultTtl))
	}
	if params != nil && params.GeoMatchLevel != nil {
		data.Set("GeoMatchLevel", *params.GeoMatchLevel)
	}
	if params != nil && params.InterceptCallbackUrl != nil {
		data.Set("InterceptCallbackUrl", *params.InterceptCallbackUrl)
	}
	if params != nil && params.NumberSelectionBehavior != nil {
		data.Set("NumberSelectionBehavior", *params.NumberSelectionBehavior)
	}
	if params != nil && params.OutOfSessionCallbackUrl != nil {
		data.Set("OutOfSessionCallbackUrl", *params.OutOfSessionCallbackUrl)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
