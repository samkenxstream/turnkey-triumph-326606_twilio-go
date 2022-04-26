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

// Optional parameters for the method 'CreateMediaProcessor'
type CreateMediaProcessorParams struct {
	// The [Media Extension](/docs/live/api/media-extensions-overview) name or URL. Ex: `video-composer-v1`
	Extension *string `json:"Extension,omitempty"`
	// The context of the Media Extension, represented as a JSON dictionary. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about the context to send.
	ExtensionContext *string `json:"ExtensionContext,omitempty"`
	// User-defined environment variables for the Media Extension, represented as a JSON dictionary of key/value strings. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about whether you need to provide this.
	ExtensionEnvironment *interface{} `json:"ExtensionEnvironment,omitempty"`
	// The maximum time, in seconds, that the MediaProcessor can run before automatically ends. The default value is 300 seconds, and the maximum value is 90000 seconds. Once this maximum duration is reached, Twilio will end the MediaProcessor, regardless of whether media is still streaming.
	MaxDuration *int `json:"MaxDuration,omitempty"`
	// The URL to which Twilio will send asynchronous webhook requests for every MediaProcessor event. See [Status Callbacks](/docs/live/status-callbacks) for details.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method Twilio should use to call the `status_callback` URL. Can be `POST` or `GET` and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
}

func (params *CreateMediaProcessorParams) SetExtension(Extension string) *CreateMediaProcessorParams {
	params.Extension = &Extension
	return params
}
func (params *CreateMediaProcessorParams) SetExtensionContext(ExtensionContext string) *CreateMediaProcessorParams {
	params.ExtensionContext = &ExtensionContext
	return params
}
func (params *CreateMediaProcessorParams) SetExtensionEnvironment(ExtensionEnvironment interface{}) *CreateMediaProcessorParams {
	params.ExtensionEnvironment = &ExtensionEnvironment
	return params
}
func (params *CreateMediaProcessorParams) SetMaxDuration(MaxDuration int) *CreateMediaProcessorParams {
	params.MaxDuration = &MaxDuration
	return params
}
func (params *CreateMediaProcessorParams) SetStatusCallback(StatusCallback string) *CreateMediaProcessorParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateMediaProcessorParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateMediaProcessorParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}

func (c *ApiService) CreateMediaProcessor(params *CreateMediaProcessorParams) (*MediaV1MediaProcessor, error) {
	path := "/v1/MediaProcessors"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Extension != nil {
		data.Set("Extension", *params.Extension)
	}
	if params != nil && params.ExtensionContext != nil {
		data.Set("ExtensionContext", *params.ExtensionContext)
	}
	if params != nil && params.ExtensionEnvironment != nil {
		v, err := json.Marshal(params.ExtensionEnvironment)

		if err != nil {
			return nil, err
		}

		data.Set("ExtensionEnvironment", string(v))
	}
	if params != nil && params.MaxDuration != nil {
		data.Set("MaxDuration", fmt.Sprint(*params.MaxDuration))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1MediaProcessor{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Returns a single MediaProcessor resource identified by a SID.
func (c *ApiService) FetchMediaProcessor(Sid string) (*MediaV1MediaProcessor, error) {
	path := "/v1/MediaProcessors/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1MediaProcessor{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMediaProcessor'
type ListMediaProcessorParams struct {
	// The sort order of the list by `date_created`. Can be: `asc` (ascending) or `desc` (descending) with `desc` as the default.
	Order *string `json:"Order,omitempty"`
	// Status to filter by, with possible values `started`, `ended` or `failed`.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMediaProcessorParams) SetOrder(Order string) *ListMediaProcessorParams {
	params.Order = &Order
	return params
}
func (params *ListMediaProcessorParams) SetStatus(Status string) *ListMediaProcessorParams {
	params.Status = &Status
	return params
}
func (params *ListMediaProcessorParams) SetPageSize(PageSize int) *ListMediaProcessorParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMediaProcessorParams) SetLimit(Limit int) *ListMediaProcessorParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of MediaProcessor records from the API. Request is executed immediately.
func (c *ApiService) PageMediaProcessor(params *ListMediaProcessorParams, pageToken, pageNumber string) (*ListMediaProcessorResponse, error) {
	path := "/v1/MediaProcessors"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
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

	ps := &ListMediaProcessorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists MediaProcessor records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMediaProcessor(params *ListMediaProcessorParams) ([]MediaV1MediaProcessor, error) {
	response, errors := c.StreamMediaProcessor(params)

	records := make([]MediaV1MediaProcessor, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams MediaProcessor records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMediaProcessor(params *ListMediaProcessorParams) (chan MediaV1MediaProcessor, chan error) {
	if params == nil {
		params = &ListMediaProcessorParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MediaV1MediaProcessor, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMediaProcessor(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMediaProcessor(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMediaProcessor(response *ListMediaProcessorResponse, params *ListMediaProcessorParams, recordChannel chan MediaV1MediaProcessor, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.MediaProcessors
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMediaProcessorResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMediaProcessorResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMediaProcessorResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMediaProcessorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMediaProcessor'
type UpdateMediaProcessorParams struct {
	// The status of the MediaProcessor. Can be `ended`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateMediaProcessorParams) SetStatus(Status string) *UpdateMediaProcessorParams {
	params.Status = &Status
	return params
}

// Updates a MediaProcessor resource identified by a SID.
func (c *ApiService) UpdateMediaProcessor(Sid string, params *UpdateMediaProcessorParams) (*MediaV1MediaProcessor, error) {
	path := "/v1/MediaProcessors/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1MediaProcessor{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
