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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateConversationMessage'
type CreateConversationMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The channel specific identifier of the message's author. Defaults to `system`.
	Author *string `json:"Author,omitempty"`
	// The content of the message, can be up to 1,600 characters long.
	Body *string `json:"Body,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated. `null` if the message has not been edited.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The Media SID to be attached to the new Message.
	MediaSid *string `json:"MediaSid,omitempty"`
}

func (params *CreateConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateConversationMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateConversationMessageParams) SetAttributes(Attributes string) *CreateConversationMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateConversationMessageParams) SetAuthor(Author string) *CreateConversationMessageParams {
	params.Author = &Author
	return params
}
func (params *CreateConversationMessageParams) SetBody(Body string) *CreateConversationMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateConversationMessageParams) SetDateCreated(DateCreated time.Time) *CreateConversationMessageParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateConversationMessageParams) SetDateUpdated(DateUpdated time.Time) *CreateConversationMessageParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateConversationMessageParams) SetMediaSid(MediaSid string) *CreateConversationMessageParams {
	params.MediaSid = &MediaSid
	return params
}

// Add a new message to the conversation
func (c *ApiService) CreateConversationMessage(ConversationSid string, params *CreateConversationMessageParams) (*ConversationsV1ConversationMessage, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Author != nil {
		data.Set("Author", *params.Author)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.MediaSid != nil {
		data.Set("MediaSid", *params.MediaSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteConversationMessage'
type DeleteConversationMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteConversationMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a message from the conversation
func (c *ApiService) DeleteConversationMessage(ConversationSid string, Sid string, params *DeleteConversationMessageParams) error {
	path := "/v1/Conversations/{ConversationSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a message from the conversation
func (c *ApiService) FetchConversationMessage(ConversationSid string, Sid string) (*ConversationsV1ConversationMessage, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversationMessage'
type ListConversationMessageParams struct {
	// The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending), with `asc` as the default.
	Order *string `json:"Order,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConversationMessageParams) SetOrder(Order string) *ListConversationMessageParams {
	params.Order = &Order
	return params
}
func (params *ListConversationMessageParams) SetPageSize(PageSize int) *ListConversationMessageParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConversationMessageParams) SetLimit(Limit int) *ListConversationMessageParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConversationMessage records from the API. Request is executed immediately.
func (c *ApiService) PageConversationMessage(ConversationSid string, params *ListConversationMessageParams, pageToken, pageNumber string) (*ListConversationMessageResponse, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages"

	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
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

	ps := &ListConversationMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConversationMessage records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConversationMessage(ConversationSid string, params *ListConversationMessageParams) ([]ConversationsV1ConversationMessage, error) {
	response, err := c.StreamConversationMessage(ConversationSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ConversationsV1ConversationMessage, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams ConversationMessage records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConversationMessage(ConversationSid string, params *ListConversationMessageParams) (chan ConversationsV1ConversationMessage, error) {
	if params == nil {
		params = &ListConversationMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationMessage(ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ConversationMessage, 1)

	go func() {
		for response != nil {
			responseRecords := response.Messages
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListConversationMessageResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConversationMessageResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConversationMessageResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConversationMessage'
type UpdateConversationMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The channel specific identifier of the message's author. Defaults to `system`.
	Author *string `json:"Author,omitempty"`
	// The content of the message, can be up to 1,600 characters long.
	Body *string `json:"Body,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated. `null` if the message has not been edited.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
}

func (params *UpdateConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateConversationMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateConversationMessageParams) SetAttributes(Attributes string) *UpdateConversationMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateConversationMessageParams) SetAuthor(Author string) *UpdateConversationMessageParams {
	params.Author = &Author
	return params
}
func (params *UpdateConversationMessageParams) SetBody(Body string) *UpdateConversationMessageParams {
	params.Body = &Body
	return params
}
func (params *UpdateConversationMessageParams) SetDateCreated(DateCreated time.Time) *UpdateConversationMessageParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateConversationMessageParams) SetDateUpdated(DateUpdated time.Time) *UpdateConversationMessageParams {
	params.DateUpdated = &DateUpdated
	return params
}

// Update an existing message in the conversation
func (c *ApiService) UpdateConversationMessage(ConversationSid string, Sid string, params *UpdateConversationMessageParams) (*ConversationsV1ConversationMessage, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Author != nil {
		data.Set("Author", *params.Author)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
