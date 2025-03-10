/*
 * Twilio - Media
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

// Deletes a MediaRecording resource identified by a SID.
func (c *ApiService) DeleteMediaRecording(Sid string) error {
	path := "/v1/MediaRecordings/{Sid}"
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

// Returns a single MediaRecording resource identified by a SID.
func (c *ApiService) FetchMediaRecording(Sid string) (*MediaV1MediaRecording, error) {
	path := "/v1/MediaRecordings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1MediaRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMediaRecording'
type ListMediaRecordingParams struct {
	// The sort order of the list by `date_created`. Can be: `asc` (ascending) or `desc` (descending) with `desc` as the default.
	Order *string `json:"Order,omitempty"`
	// Status to filter by, with possible values `processing`, `completed`, `deleted`, or `failed`.
	Status *string `json:"Status,omitempty"`
	// SID of a MediaProcessor to filter by.
	ProcessorSid *string `json:"ProcessorSid,omitempty"`
	// SID of a MediaRecording source to filter by.
	SourceSid *string `json:"SourceSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMediaRecordingParams) SetOrder(Order string) *ListMediaRecordingParams {
	params.Order = &Order
	return params
}
func (params *ListMediaRecordingParams) SetStatus(Status string) *ListMediaRecordingParams {
	params.Status = &Status
	return params
}
func (params *ListMediaRecordingParams) SetProcessorSid(ProcessorSid string) *ListMediaRecordingParams {
	params.ProcessorSid = &ProcessorSid
	return params
}
func (params *ListMediaRecordingParams) SetSourceSid(SourceSid string) *ListMediaRecordingParams {
	params.SourceSid = &SourceSid
	return params
}
func (params *ListMediaRecordingParams) SetPageSize(PageSize int) *ListMediaRecordingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMediaRecordingParams) SetLimit(Limit int) *ListMediaRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of MediaRecording records from the API. Request is executed immediately.
func (c *ApiService) PageMediaRecording(params *ListMediaRecordingParams, pageToken, pageNumber string) (*ListMediaRecordingResponse, error) {
	path := "/v1/MediaRecordings"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.ProcessorSid != nil {
		data.Set("ProcessorSid", *params.ProcessorSid)
	}
	if params != nil && params.SourceSid != nil {
		data.Set("SourceSid", *params.SourceSid)
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

	ps := &ListMediaRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists MediaRecording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMediaRecording(params *ListMediaRecordingParams) ([]MediaV1MediaRecording, error) {
	response, err := c.StreamMediaRecording(params)
	if err != nil {
		return nil, err
	}

	records := make([]MediaV1MediaRecording, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams MediaRecording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMediaRecording(params *ListMediaRecordingParams) (chan MediaV1MediaRecording, error) {
	if params == nil {
		params = &ListMediaRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMediaRecording(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan MediaV1MediaRecording, 1)

	go func() {
		for response != nil {
			responseRecords := response.MediaRecordings
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListMediaRecordingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMediaRecordingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMediaRecordingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMediaRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
