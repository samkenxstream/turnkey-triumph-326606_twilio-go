/*
 * Twilio - Monitor
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

func (c *ApiService) FetchEvent(Sid string) (*MonitorV1Event, error) {
	path := "/v1/Events/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MonitorV1Event{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEvent'
type ListEventParams struct {
	// Only include events initiated by this Actor. Useful for auditing actions taken by specific users or API credentials.
	ActorSid *string `json:"ActorSid,omitempty"`
	// Only include events of this [Event Type](https://www.twilio.com/docs/usage/monitor-events#event-types).
	EventType *string `json:"EventType,omitempty"`
	// Only include events that refer to this resource. Useful for discovering the history of a specific resource.
	ResourceSid *string `json:"ResourceSid,omitempty"`
	// Only include events that originated from this IP address. Useful for tracking suspicious activity originating from the API or the Twilio Console.
	SourceIpAddress *string `json:"SourceIpAddress,omitempty"`
	// Only include events that occurred on or after this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only include events that occurred on or before this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEventParams) SetActorSid(ActorSid string) *ListEventParams {
	params.ActorSid = &ActorSid
	return params
}
func (params *ListEventParams) SetEventType(EventType string) *ListEventParams {
	params.EventType = &EventType
	return params
}
func (params *ListEventParams) SetResourceSid(ResourceSid string) *ListEventParams {
	params.ResourceSid = &ResourceSid
	return params
}
func (params *ListEventParams) SetSourceIpAddress(SourceIpAddress string) *ListEventParams {
	params.SourceIpAddress = &SourceIpAddress
	return params
}
func (params *ListEventParams) SetStartDate(StartDate time.Time) *ListEventParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListEventParams) SetEndDate(EndDate time.Time) *ListEventParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListEventParams) SetPageSize(PageSize int) *ListEventParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEventParams) SetLimit(Limit int) *ListEventParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Event records from the API. Request is executed immediately.
func (c *ApiService) PageEvent(params *ListEventParams, pageToken, pageNumber string) (*ListEventResponse, error) {
	path := "/v1/Events"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActorSid != nil {
		data.Set("ActorSid", *params.ActorSid)
	}
	if params != nil && params.EventType != nil {
		data.Set("EventType", *params.EventType)
	}
	if params != nil && params.ResourceSid != nil {
		data.Set("ResourceSid", *params.ResourceSid)
	}
	if params != nil && params.SourceIpAddress != nil {
		data.Set("SourceIpAddress", *params.SourceIpAddress)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
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

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Event records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEvent(params *ListEventParams) ([]MonitorV1Event, error) {
	response, err := c.StreamEvent(params)
	if err != nil {
		return nil, err
	}

	records := make([]MonitorV1Event, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Event records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEvent(params *ListEventParams) (chan MonitorV1Event, error) {
	if params == nil {
		params = &ListEventParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEvent(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan MonitorV1Event, 1)

	go func() {
		for response != nil {
			responseRecords := response.Events
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListEventResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEventResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEventResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
