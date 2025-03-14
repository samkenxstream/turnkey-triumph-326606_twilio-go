/*
 * Twilio - Conversations
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

// Optional parameters for the method 'CreateServiceConversationScopedWebhook'
type CreateServiceConversationScopedWebhookParams struct {
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	// The message index for which and it's successors the webhook will be replayed. Not set by default
	ConfigurationReplayAfter *int `json:"Configuration.ReplayAfter,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
	// The target of this webhook: `webhook`, `studio`, `trigger`
	Target *string `json:"Target,omitempty"`
}

func (params *CreateServiceConversationScopedWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *CreateServiceConversationScopedWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *CreateServiceConversationScopedWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *CreateServiceConversationScopedWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *CreateServiceConversationScopedWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *CreateServiceConversationScopedWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *CreateServiceConversationScopedWebhookParams) SetConfigurationReplayAfter(ConfigurationReplayAfter int) *CreateServiceConversationScopedWebhookParams {
	params.ConfigurationReplayAfter = &ConfigurationReplayAfter
	return params
}
func (params *CreateServiceConversationScopedWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *CreateServiceConversationScopedWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *CreateServiceConversationScopedWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *CreateServiceConversationScopedWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}
func (params *CreateServiceConversationScopedWebhookParams) SetTarget(Target string) *CreateServiceConversationScopedWebhookParams {
	params.Target = &Target
	return params
}

// Create a new webhook scoped to the conversation in a specific service
func (c *ApiService) CreateServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, params *CreateServiceConversationScopedWebhookParams) (*ConversationsV1ServiceConversationScopedWebhook, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationReplayAfter != nil {
		data.Set("Configuration.ReplayAfter", fmt.Sprint(*params.ConfigurationReplayAfter))
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an existing webhook scoped to the conversation
func (c *ApiService) DeleteServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, Sid string) error {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
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

// Fetch the configuration of a conversation-scoped webhook
func (c *ApiService) FetchServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, Sid string) (*ConversationsV1ServiceConversationScopedWebhook, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceConversationScopedWebhook'
type ListServiceConversationScopedWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceConversationScopedWebhookParams) SetPageSize(PageSize int) *ListServiceConversationScopedWebhookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceConversationScopedWebhookParams) SetLimit(Limit int) *ListServiceConversationScopedWebhookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceConversationScopedWebhook records from the API. Request is executed immediately.
func (c *ApiService) PageServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, params *ListServiceConversationScopedWebhookParams, pageToken, pageNumber string) (*ListServiceConversationScopedWebhookResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

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

	ps := &ListServiceConversationScopedWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceConversationScopedWebhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, params *ListServiceConversationScopedWebhookParams) ([]ConversationsV1ServiceConversationScopedWebhook, error) {
	response, err := c.StreamServiceConversationScopedWebhook(ChatServiceSid, ConversationSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ConversationsV1ServiceConversationScopedWebhook, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams ServiceConversationScopedWebhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, params *ListServiceConversationScopedWebhookParams) (chan ConversationsV1ServiceConversationScopedWebhook, error) {
	if params == nil {
		params = &ListServiceConversationScopedWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceConversationScopedWebhook(ChatServiceSid, ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ServiceConversationScopedWebhook, 1)

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
			if record, err = client.GetNext(c.baseURL, response, c.getNextListServiceConversationScopedWebhookResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceConversationScopedWebhookResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceConversationScopedWebhookResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceConversationScopedWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceConversationScopedWebhook'
type UpdateServiceConversationScopedWebhookParams struct {
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
}

func (params *UpdateServiceConversationScopedWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *UpdateServiceConversationScopedWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *UpdateServiceConversationScopedWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *UpdateServiceConversationScopedWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *UpdateServiceConversationScopedWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *UpdateServiceConversationScopedWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *UpdateServiceConversationScopedWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *UpdateServiceConversationScopedWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *UpdateServiceConversationScopedWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *UpdateServiceConversationScopedWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}

// Update an existing conversation-scoped webhook
func (c *ApiService) UpdateServiceConversationScopedWebhook(ChatServiceSid string, ConversationSid string, Sid string, params *UpdateServiceConversationScopedWebhookParams) (*ConversationsV1ServiceConversationScopedWebhook, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
