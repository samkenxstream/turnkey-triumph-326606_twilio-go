/*
 * Twilio - Flex
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

// Optional parameters for the method 'CreateFlexFlow'
type CreateFlexFlowParams struct {
	// The channel type. One of `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`. By default, Studio’s Send to Flex widget passes it on to the Task attributes for Tasks created based on this Flex Flow. The Task attributes will be used by the Flex UI to render the respective Task as appropriate (applying channel-specific design and length limits). If `channelType` is `facebook`, `whatsapp` or `line`, the Send to Flex widget should set the Task Channel to Programmable Chat.
	ChannelType *string `json:"ChannelType,omitempty"`
	// The SID of the chat service.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The channel contact's Identity.
	ContactIdentity *string `json:"ContactIdentity,omitempty"`
	// Whether the new Flex Flow is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the Flex Flow resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The Task Channel SID (TCXXXX) or unique name (e.g., `sms`) to use for the Task that will be created. Applicable and required when `integrationType` is `task`. The default value is `default`.
	IntegrationChannel *string `json:"Integration.Channel,omitempty"`
	// In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
	IntegrationCreationOnMessage *bool `json:"Integration.CreationOnMessage,omitempty"`
	// The SID of the Studio Flow. Required when `integrationType` is `studio`.
	IntegrationFlowSid *string `json:"Integration.FlowSid,omitempty"`
	// The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationPriority *int `json:"Integration.Priority,omitempty"`
	// The number of times to retry the Studio Flow or webhook in case of failure. Takes integer values from 0 to 3 with the default being 3. Optional when `integrationType` is `studio` or `external`, not applicable otherwise.
	IntegrationRetryCount *int `json:"Integration.RetryCount,omitempty"`
	// The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationTimeout *int `json:"Integration.Timeout,omitempty"`
	// The URL of the external webhook. Required when `integrationType` is `external`.
	IntegrationUrl *string `json:"Integration.Url,omitempty"`
	// The Workflow SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkflowSid *string `json:"Integration.WorkflowSid,omitempty"`
	// The Workspace SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkspaceSid *string `json:"Integration.WorkspaceSid,omitempty"`
	// The software that will handle inbound messages. [Integration Type](https://www.twilio.com/docs/flex/developer/messaging/manage-flows#integration-types) can be: `studio`, `external`, or `task`.
	IntegrationType *string `json:"IntegrationType,omitempty"`
	// When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
	JanitorEnabled *bool `json:"JanitorEnabled,omitempty"`
	// When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
	LongLived *bool `json:"LongLived,omitempty"`
}

func (params *CreateFlexFlowParams) SetChannelType(ChannelType string) *CreateFlexFlowParams {
	params.ChannelType = &ChannelType
	return params
}
func (params *CreateFlexFlowParams) SetChatServiceSid(ChatServiceSid string) *CreateFlexFlowParams {
	params.ChatServiceSid = &ChatServiceSid
	return params
}
func (params *CreateFlexFlowParams) SetContactIdentity(ContactIdentity string) *CreateFlexFlowParams {
	params.ContactIdentity = &ContactIdentity
	return params
}
func (params *CreateFlexFlowParams) SetEnabled(Enabled bool) *CreateFlexFlowParams {
	params.Enabled = &Enabled
	return params
}
func (params *CreateFlexFlowParams) SetFriendlyName(FriendlyName string) *CreateFlexFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationChannel(IntegrationChannel string) *CreateFlexFlowParams {
	params.IntegrationChannel = &IntegrationChannel
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationCreationOnMessage(IntegrationCreationOnMessage bool) *CreateFlexFlowParams {
	params.IntegrationCreationOnMessage = &IntegrationCreationOnMessage
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationFlowSid(IntegrationFlowSid string) *CreateFlexFlowParams {
	params.IntegrationFlowSid = &IntegrationFlowSid
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationPriority(IntegrationPriority int) *CreateFlexFlowParams {
	params.IntegrationPriority = &IntegrationPriority
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationRetryCount(IntegrationRetryCount int) *CreateFlexFlowParams {
	params.IntegrationRetryCount = &IntegrationRetryCount
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationTimeout(IntegrationTimeout int) *CreateFlexFlowParams {
	params.IntegrationTimeout = &IntegrationTimeout
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationUrl(IntegrationUrl string) *CreateFlexFlowParams {
	params.IntegrationUrl = &IntegrationUrl
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationWorkflowSid(IntegrationWorkflowSid string) *CreateFlexFlowParams {
	params.IntegrationWorkflowSid = &IntegrationWorkflowSid
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationWorkspaceSid(IntegrationWorkspaceSid string) *CreateFlexFlowParams {
	params.IntegrationWorkspaceSid = &IntegrationWorkspaceSid
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationType(IntegrationType string) *CreateFlexFlowParams {
	params.IntegrationType = &IntegrationType
	return params
}
func (params *CreateFlexFlowParams) SetJanitorEnabled(JanitorEnabled bool) *CreateFlexFlowParams {
	params.JanitorEnabled = &JanitorEnabled
	return params
}
func (params *CreateFlexFlowParams) SetLongLived(LongLived bool) *CreateFlexFlowParams {
	params.LongLived = &LongLived
	return params
}

func (c *ApiService) CreateFlexFlow(params *CreateFlexFlowParams) (*FlexV1FlexFlow, error) {
	path := "/v1/FlexFlows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelType != nil {
		data.Set("ChannelType", *params.ChannelType)
	}
	if params != nil && params.ChatServiceSid != nil {
		data.Set("ChatServiceSid", *params.ChatServiceSid)
	}
	if params != nil && params.ContactIdentity != nil {
		data.Set("ContactIdentity", *params.ContactIdentity)
	}
	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IntegrationChannel != nil {
		data.Set("Integration.Channel", *params.IntegrationChannel)
	}
	if params != nil && params.IntegrationCreationOnMessage != nil {
		data.Set("Integration.CreationOnMessage", fmt.Sprint(*params.IntegrationCreationOnMessage))
	}
	if params != nil && params.IntegrationFlowSid != nil {
		data.Set("Integration.FlowSid", *params.IntegrationFlowSid)
	}
	if params != nil && params.IntegrationPriority != nil {
		data.Set("Integration.Priority", fmt.Sprint(*params.IntegrationPriority))
	}
	if params != nil && params.IntegrationRetryCount != nil {
		data.Set("Integration.RetryCount", fmt.Sprint(*params.IntegrationRetryCount))
	}
	if params != nil && params.IntegrationTimeout != nil {
		data.Set("Integration.Timeout", fmt.Sprint(*params.IntegrationTimeout))
	}
	if params != nil && params.IntegrationUrl != nil {
		data.Set("Integration.Url", *params.IntegrationUrl)
	}
	if params != nil && params.IntegrationWorkflowSid != nil {
		data.Set("Integration.WorkflowSid", *params.IntegrationWorkflowSid)
	}
	if params != nil && params.IntegrationWorkspaceSid != nil {
		data.Set("Integration.WorkspaceSid", *params.IntegrationWorkspaceSid)
	}
	if params != nil && params.IntegrationType != nil {
		data.Set("IntegrationType", *params.IntegrationType)
	}
	if params != nil && params.JanitorEnabled != nil {
		data.Set("JanitorEnabled", fmt.Sprint(*params.JanitorEnabled))
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteFlexFlow(Sid string) error {
	path := "/v1/FlexFlows/{Sid}"
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

func (c *ApiService) FetchFlexFlow(Sid string) (*FlexV1FlexFlow, error) {
	path := "/v1/FlexFlows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlexFlow'
type ListFlexFlowParams struct {
	// The `friendly_name` of the Flex Flow resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFlexFlowParams) SetFriendlyName(FriendlyName string) *ListFlexFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListFlexFlowParams) SetPageSize(PageSize int) *ListFlexFlowParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFlexFlowParams) SetLimit(Limit int) *ListFlexFlowParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FlexFlow records from the API. Request is executed immediately.
func (c *ApiService) PageFlexFlow(params *ListFlexFlowParams, pageToken, pageNumber string) (*ListFlexFlowResponse, error) {
	path := "/v1/FlexFlows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListFlexFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FlexFlow records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlexFlow(params *ListFlexFlowParams) ([]FlexV1FlexFlow, error) {
	response, err := c.StreamFlexFlow(params)
	if err != nil {
		return nil, err
	}

	records := make([]FlexV1FlexFlow, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams FlexFlow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlexFlow(params *ListFlexFlowParams) (chan FlexV1FlexFlow, error) {
	if params == nil {
		params = &ListFlexFlowParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFlexFlow(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan FlexV1FlexFlow, 1)

	go func() {
		for response != nil {
			responseRecords := response.FlexFlows
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListFlexFlowResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFlexFlowResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFlexFlowResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlexFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateFlexFlow'
type UpdateFlexFlowParams struct {
	// The channel type. One of `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`. By default, Studio’s Send to Flex widget passes it on to the Task attributes for Tasks created based on this Flex Flow. The Task attributes will be used by the Flex UI to render the respective Task as appropriate (applying channel-specific design and length limits). If `channelType` is `facebook`, `whatsapp` or `line`, the Send to Flex widget should set the Task Channel to Programmable Chat.
	ChannelType *string `json:"ChannelType,omitempty"`
	// The SID of the chat service.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The channel contact's Identity.
	ContactIdentity *string `json:"ContactIdentity,omitempty"`
	// Whether the new Flex Flow is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the Flex Flow resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The Task Channel SID (TCXXXX) or unique name (e.g., `sms`) to use for the Task that will be created. Applicable and required when `integrationType` is `task`. The default value is `default`.
	IntegrationChannel *string `json:"Integration.Channel,omitempty"`
	// In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
	IntegrationCreationOnMessage *bool `json:"Integration.CreationOnMessage,omitempty"`
	// The SID of the Studio Flow. Required when `integrationType` is `studio`.
	IntegrationFlowSid *string `json:"Integration.FlowSid,omitempty"`
	// The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationPriority *int `json:"Integration.Priority,omitempty"`
	// The number of times to retry the Studio Flow or webhook in case of failure. Takes integer values from 0 to 3 with the default being 3. Optional when `integrationType` is `studio` or `external`, not applicable otherwise.
	IntegrationRetryCount *int `json:"Integration.RetryCount,omitempty"`
	// The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationTimeout *int `json:"Integration.Timeout,omitempty"`
	// The URL of the external webhook. Required when `integrationType` is `external`.
	IntegrationUrl *string `json:"Integration.Url,omitempty"`
	// The Workflow SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkflowSid *string `json:"Integration.WorkflowSid,omitempty"`
	// The Workspace SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkspaceSid *string `json:"Integration.WorkspaceSid,omitempty"`
	// The software that will handle inbound messages. [Integration Type](https://www.twilio.com/docs/flex/developer/messaging/manage-flows#integration-types) can be: `studio`, `external`, or `task`.
	IntegrationType *string `json:"IntegrationType,omitempty"`
	// When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
	JanitorEnabled *bool `json:"JanitorEnabled,omitempty"`
	// When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
	LongLived *bool `json:"LongLived,omitempty"`
}

func (params *UpdateFlexFlowParams) SetChannelType(ChannelType string) *UpdateFlexFlowParams {
	params.ChannelType = &ChannelType
	return params
}
func (params *UpdateFlexFlowParams) SetChatServiceSid(ChatServiceSid string) *UpdateFlexFlowParams {
	params.ChatServiceSid = &ChatServiceSid
	return params
}
func (params *UpdateFlexFlowParams) SetContactIdentity(ContactIdentity string) *UpdateFlexFlowParams {
	params.ContactIdentity = &ContactIdentity
	return params
}
func (params *UpdateFlexFlowParams) SetEnabled(Enabled bool) *UpdateFlexFlowParams {
	params.Enabled = &Enabled
	return params
}
func (params *UpdateFlexFlowParams) SetFriendlyName(FriendlyName string) *UpdateFlexFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationChannel(IntegrationChannel string) *UpdateFlexFlowParams {
	params.IntegrationChannel = &IntegrationChannel
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationCreationOnMessage(IntegrationCreationOnMessage bool) *UpdateFlexFlowParams {
	params.IntegrationCreationOnMessage = &IntegrationCreationOnMessage
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationFlowSid(IntegrationFlowSid string) *UpdateFlexFlowParams {
	params.IntegrationFlowSid = &IntegrationFlowSid
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationPriority(IntegrationPriority int) *UpdateFlexFlowParams {
	params.IntegrationPriority = &IntegrationPriority
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationRetryCount(IntegrationRetryCount int) *UpdateFlexFlowParams {
	params.IntegrationRetryCount = &IntegrationRetryCount
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationTimeout(IntegrationTimeout int) *UpdateFlexFlowParams {
	params.IntegrationTimeout = &IntegrationTimeout
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationUrl(IntegrationUrl string) *UpdateFlexFlowParams {
	params.IntegrationUrl = &IntegrationUrl
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationWorkflowSid(IntegrationWorkflowSid string) *UpdateFlexFlowParams {
	params.IntegrationWorkflowSid = &IntegrationWorkflowSid
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationWorkspaceSid(IntegrationWorkspaceSid string) *UpdateFlexFlowParams {
	params.IntegrationWorkspaceSid = &IntegrationWorkspaceSid
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationType(IntegrationType string) *UpdateFlexFlowParams {
	params.IntegrationType = &IntegrationType
	return params
}
func (params *UpdateFlexFlowParams) SetJanitorEnabled(JanitorEnabled bool) *UpdateFlexFlowParams {
	params.JanitorEnabled = &JanitorEnabled
	return params
}
func (params *UpdateFlexFlowParams) SetLongLived(LongLived bool) *UpdateFlexFlowParams {
	params.LongLived = &LongLived
	return params
}

func (c *ApiService) UpdateFlexFlow(Sid string, params *UpdateFlexFlowParams) (*FlexV1FlexFlow, error) {
	path := "/v1/FlexFlows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelType != nil {
		data.Set("ChannelType", *params.ChannelType)
	}
	if params != nil && params.ChatServiceSid != nil {
		data.Set("ChatServiceSid", *params.ChatServiceSid)
	}
	if params != nil && params.ContactIdentity != nil {
		data.Set("ContactIdentity", *params.ContactIdentity)
	}
	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IntegrationChannel != nil {
		data.Set("Integration.Channel", *params.IntegrationChannel)
	}
	if params != nil && params.IntegrationCreationOnMessage != nil {
		data.Set("Integration.CreationOnMessage", fmt.Sprint(*params.IntegrationCreationOnMessage))
	}
	if params != nil && params.IntegrationFlowSid != nil {
		data.Set("Integration.FlowSid", *params.IntegrationFlowSid)
	}
	if params != nil && params.IntegrationPriority != nil {
		data.Set("Integration.Priority", fmt.Sprint(*params.IntegrationPriority))
	}
	if params != nil && params.IntegrationRetryCount != nil {
		data.Set("Integration.RetryCount", fmt.Sprint(*params.IntegrationRetryCount))
	}
	if params != nil && params.IntegrationTimeout != nil {
		data.Set("Integration.Timeout", fmt.Sprint(*params.IntegrationTimeout))
	}
	if params != nil && params.IntegrationUrl != nil {
		data.Set("Integration.Url", *params.IntegrationUrl)
	}
	if params != nil && params.IntegrationWorkflowSid != nil {
		data.Set("Integration.WorkflowSid", *params.IntegrationWorkflowSid)
	}
	if params != nil && params.IntegrationWorkspaceSid != nil {
		data.Set("Integration.WorkspaceSid", *params.IntegrationWorkspaceSid)
	}
	if params != nil && params.IntegrationType != nil {
		data.Set("IntegrationType", *params.IntegrationType)
	}
	if params != nil && params.JanitorEnabled != nil {
		data.Set("JanitorEnabled", fmt.Sprint(*params.JanitorEnabled))
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
