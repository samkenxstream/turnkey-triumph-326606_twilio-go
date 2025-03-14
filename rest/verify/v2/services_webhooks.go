/*
 * Twilio - Verify
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

// Optional parameters for the method 'CreateWebhook'
type CreateWebhookParams struct {
	// The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied`
	EventTypes *[]string `json:"EventTypes,omitempty"`
	// The string that you assigned to describe the webhook. **This value should not contain PII.**
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The webhook status. Default value is `enabled`. One of: `enabled` or `disabled`
	Status *string `json:"Status,omitempty"`
	// The webhook version. Default value is `v2` which includes all the latest fields. Version `v1` is legacy and may be removed in the future.
	Version *string `json:"Version,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
}

func (params *CreateWebhookParams) SetEventTypes(EventTypes []string) *CreateWebhookParams {
	params.EventTypes = &EventTypes
	return params
}
func (params *CreateWebhookParams) SetFriendlyName(FriendlyName string) *CreateWebhookParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateWebhookParams) SetStatus(Status string) *CreateWebhookParams {
	params.Status = &Status
	return params
}
func (params *CreateWebhookParams) SetVersion(Version string) *CreateWebhookParams {
	params.Version = &Version
	return params
}
func (params *CreateWebhookParams) SetWebhookUrl(WebhookUrl string) *CreateWebhookParams {
	params.WebhookUrl = &WebhookUrl
	return params
}

// Create a new Webhook for the Service
func (c *ApiService) CreateWebhook(ServiceSid string, params *CreateWebhookParams) (*VerifyV2Webhook, error) {
	path := "/v2/Services/{ServiceSid}/Webhooks"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EventTypes != nil {
		for _, item := range *params.EventTypes {
			data.Add("EventTypes", item)
		}
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Version != nil {
		data.Set("Version", *params.Version)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Webhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Webhook.
func (c *ApiService) DeleteWebhook(ServiceSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Webhooks/{Sid}"
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

// Fetch a specific Webhook.
func (c *ApiService) FetchWebhook(ServiceSid string, Sid string) (*VerifyV2Webhook, error) {
	path := "/v2/Services/{ServiceSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Webhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWebhook'
type ListWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWebhookParams) SetPageSize(PageSize int) *ListWebhookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWebhookParams) SetLimit(Limit int) *ListWebhookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Webhook records from the API. Request is executed immediately.
func (c *ApiService) PageWebhook(ServiceSid string, params *ListWebhookParams, pageToken, pageNumber string) (*ListWebhookResponse, error) {
	path := "/v2/Services/{ServiceSid}/Webhooks"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Webhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWebhook(ServiceSid string, params *ListWebhookParams) ([]VerifyV2Webhook, error) {
	response, err := c.StreamWebhook(ServiceSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]VerifyV2Webhook, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Webhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWebhook(ServiceSid string, params *ListWebhookParams) (chan VerifyV2Webhook, error) {
	if params == nil {
		params = &ListWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageWebhook(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2Webhook, 1)

	go func() {
		for response != nil {
			responseRecords := response.Webhooks
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListWebhookResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListWebhookResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListWebhookResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWebhook'
type UpdateWebhookParams struct {
	// The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied`
	EventTypes *[]string `json:"EventTypes,omitempty"`
	// The string that you assigned to describe the webhook. **This value should not contain PII.**
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The webhook status. Default value is `enabled`. One of: `enabled` or `disabled`
	Status *string `json:"Status,omitempty"`
	// The webhook version. Default value is `v2` which includes all the latest fields. Version `v1` is legacy and may be removed in the future.
	Version *string `json:"Version,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
}

func (params *UpdateWebhookParams) SetEventTypes(EventTypes []string) *UpdateWebhookParams {
	params.EventTypes = &EventTypes
	return params
}
func (params *UpdateWebhookParams) SetFriendlyName(FriendlyName string) *UpdateWebhookParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWebhookParams) SetStatus(Status string) *UpdateWebhookParams {
	params.Status = &Status
	return params
}
func (params *UpdateWebhookParams) SetVersion(Version string) *UpdateWebhookParams {
	params.Version = &Version
	return params
}
func (params *UpdateWebhookParams) SetWebhookUrl(WebhookUrl string) *UpdateWebhookParams {
	params.WebhookUrl = &WebhookUrl
	return params
}

func (c *ApiService) UpdateWebhook(ServiceSid string, Sid string, params *UpdateWebhookParams) (*VerifyV2Webhook, error) {
	path := "/v2/Services/{ServiceSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EventTypes != nil {
		for _, item := range *params.EventTypes {
			data.Add("EventTypes", item)
		}
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Version != nil {
		data.Set("Version", *params.Version)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Webhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
