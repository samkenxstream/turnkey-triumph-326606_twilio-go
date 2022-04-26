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

// Optional parameters for the method 'ListBillingPeriod'
type ListBillingPeriodParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBillingPeriodParams) SetPageSize(PageSize int) *ListBillingPeriodParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBillingPeriodParams) SetLimit(Limit int) *ListBillingPeriodParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of BillingPeriod records from the API. Request is executed immediately.
func (c *ApiService) PageBillingPeriod(SimSid string, params *ListBillingPeriodParams, pageToken, pageNumber string) (*ListBillingPeriodResponse, error) {
	path := "/v1/Sims/{SimSid}/BillingPeriods"

	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BillingPeriod records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBillingPeriod(SimSid string, params *ListBillingPeriodParams) ([]SupersimV1BillingPeriod, error) {
	response, errors := c.StreamBillingPeriod(SimSid, params)

	records := make([]SupersimV1BillingPeriod, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams BillingPeriod records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBillingPeriod(SimSid string, params *ListBillingPeriodParams) (chan SupersimV1BillingPeriod, chan error) {
	if params == nil {
		params = &ListBillingPeriodParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1BillingPeriod, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageBillingPeriod(SimSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamBillingPeriod(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamBillingPeriod(response *ListBillingPeriodResponse, params *ListBillingPeriodParams, recordChannel chan SupersimV1BillingPeriod, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.BillingPeriods
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListBillingPeriodResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListBillingPeriodResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListBillingPeriodResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
