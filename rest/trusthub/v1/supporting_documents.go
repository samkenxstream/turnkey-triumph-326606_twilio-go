/*
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateSupportingDocument'
type CreateSupportingDocumentParams struct {
	// The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
	Attributes *interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The type of the Supporting Document.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateSupportingDocumentParams) SetAttributes(Attributes interface{}) *CreateSupportingDocumentParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateSupportingDocumentParams) SetFriendlyName(FriendlyName string) *CreateSupportingDocumentParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateSupportingDocumentParams) SetType(Type string) *CreateSupportingDocumentParams {
	params.Type = &Type
	return params
}

// Create a new Supporting Document.
func (c *ApiService) CreateSupportingDocument(params *CreateSupportingDocumentParams) (*TrusthubV1SupportingDocument, error) {
	path := "/v1/SupportingDocuments"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1SupportingDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Supporting Document.
func (c *ApiService) DeleteSupportingDocument(Sid string) error {
	path := "/v1/SupportingDocuments/{Sid}"
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

// Fetch specific Supporting Document Instance.
func (c *ApiService) FetchSupportingDocument(Sid string) (*TrusthubV1SupportingDocument, error) {
	path := "/v1/SupportingDocuments/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1SupportingDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSupportingDocument'
type ListSupportingDocumentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSupportingDocumentParams) SetPageSize(PageSize int) *ListSupportingDocumentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSupportingDocumentParams) SetLimit(Limit int) *ListSupportingDocumentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SupportingDocument records from the API. Request is executed immediately.
func (c *ApiService) PageSupportingDocument(params *ListSupportingDocumentParams, pageToken, pageNumber string) (*ListSupportingDocumentResponse, error) {
	path := "/v1/SupportingDocuments"

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

	ps := &ListSupportingDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SupportingDocument records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSupportingDocument(params *ListSupportingDocumentParams) ([]TrusthubV1SupportingDocument, error) {
	response, err := c.StreamSupportingDocument(params)
	if err != nil {
		return nil, err
	}

	records := make([]TrusthubV1SupportingDocument, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams SupportingDocument records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSupportingDocument(params *ListSupportingDocumentParams) (chan TrusthubV1SupportingDocument, error) {
	if params == nil {
		params = &ListSupportingDocumentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSupportingDocument(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1SupportingDocument, 1)

	go func() {
		for response != nil {
			responseRecords := response.Results
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListSupportingDocumentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSupportingDocumentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSupportingDocumentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSupportingDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSupportingDocument'
type UpdateSupportingDocumentParams struct {
	// The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.
	Attributes *interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSupportingDocumentParams) SetAttributes(Attributes interface{}) *UpdateSupportingDocumentParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateSupportingDocumentParams) SetFriendlyName(FriendlyName string) *UpdateSupportingDocumentParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update an existing Supporting Document.
func (c *ApiService) UpdateSupportingDocument(Sid string, params *UpdateSupportingDocumentParams) (*TrusthubV1SupportingDocument, error) {
	path := "/v1/SupportingDocuments/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1SupportingDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
