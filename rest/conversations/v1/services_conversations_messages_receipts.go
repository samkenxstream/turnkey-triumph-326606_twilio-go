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

// Fetch the delivery and read receipts of the conversation message
func (c *ApiService) FetchServiceConversationMessageReceipt(ChatServiceSid string, ConversationSid string, MessageSid string, Sid string) (*ConversationsV1ServiceConversationMessageReceipt, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationMessageReceipt{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceConversationMessageReceipt'
type ListServiceConversationMessageReceiptParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceConversationMessageReceiptParams) SetPageSize(PageSize int) *ListServiceConversationMessageReceiptParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceConversationMessageReceiptParams) SetLimit(Limit int) *ListServiceConversationMessageReceiptParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceConversationMessageReceipt records from the API. Request is executed immediately.
func (c *ApiService) PageServiceConversationMessageReceipt(ChatServiceSid string, ConversationSid string, MessageSid string, params *ListServiceConversationMessageReceiptParams, pageToken, pageNumber string) (*ListServiceConversationMessageReceiptResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

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

	ps := &ListServiceConversationMessageReceiptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceConversationMessageReceipt records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceConversationMessageReceipt(ChatServiceSid string, ConversationSid string, MessageSid string, params *ListServiceConversationMessageReceiptParams) ([]ConversationsV1ServiceConversationMessageReceipt, error) {
	response, err := c.StreamServiceConversationMessageReceipt(ChatServiceSid, ConversationSid, MessageSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ConversationsV1ServiceConversationMessageReceipt, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams ServiceConversationMessageReceipt records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceConversationMessageReceipt(ChatServiceSid string, ConversationSid string, MessageSid string, params *ListServiceConversationMessageReceiptParams) (chan ConversationsV1ServiceConversationMessageReceipt, error) {
	if params == nil {
		params = &ListServiceConversationMessageReceiptParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceConversationMessageReceipt(ChatServiceSid, ConversationSid, MessageSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ServiceConversationMessageReceipt, 1)

	go func() {
		for response != nil {
			responseRecords := response.DeliveryReceipts
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListServiceConversationMessageReceiptResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceConversationMessageReceiptResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceConversationMessageReceiptResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceConversationMessageReceiptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
