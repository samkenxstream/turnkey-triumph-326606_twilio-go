/*
 * Twilio - Voice
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

// Optional parameters for the method 'CreateSourceIpMapping'
type CreateSourceIpMappingParams struct {
	// The Twilio-provided string that uniquely identifies the IP Record resource to map from.
	IpRecordSid *string `json:"IpRecordSid,omitempty"`
	// The SID of the SIP Domain that the IP Record should be mapped to.
	SipDomainSid *string `json:"SipDomainSid,omitempty"`
}

func (params *CreateSourceIpMappingParams) SetIpRecordSid(IpRecordSid string) *CreateSourceIpMappingParams {
	params.IpRecordSid = &IpRecordSid
	return params
}
func (params *CreateSourceIpMappingParams) SetSipDomainSid(SipDomainSid string) *CreateSourceIpMappingParams {
	params.SipDomainSid = &SipDomainSid
	return params
}

func (c *ApiService) CreateSourceIpMapping(params *CreateSourceIpMappingParams) (*VoiceV1SourceIpMapping, error) {
	path := "/v1/SourceIpMappings"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IpRecordSid != nil {
		data.Set("IpRecordSid", *params.IpRecordSid)
	}
	if params != nil && params.SipDomainSid != nil {
		data.Set("SipDomainSid", *params.SipDomainSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1SourceIpMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteSourceIpMapping(Sid string) error {
	path := "/v1/SourceIpMappings/{Sid}"
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

func (c *ApiService) FetchSourceIpMapping(Sid string) (*VoiceV1SourceIpMapping, error) {
	path := "/v1/SourceIpMappings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1SourceIpMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSourceIpMapping'
type ListSourceIpMappingParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSourceIpMappingParams) SetPageSize(PageSize int) *ListSourceIpMappingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSourceIpMappingParams) SetLimit(Limit int) *ListSourceIpMappingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SourceIpMapping records from the API. Request is executed immediately.
func (c *ApiService) PageSourceIpMapping(params *ListSourceIpMappingParams, pageToken, pageNumber string) (*ListSourceIpMappingResponse, error) {
	path := "/v1/SourceIpMappings"

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

	ps := &ListSourceIpMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SourceIpMapping records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSourceIpMapping(params *ListSourceIpMappingParams) ([]VoiceV1SourceIpMapping, error) {
	response, err := c.StreamSourceIpMapping(params)
	if err != nil {
		return nil, err
	}

	records := make([]VoiceV1SourceIpMapping, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams SourceIpMapping records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSourceIpMapping(params *ListSourceIpMappingParams) (chan VoiceV1SourceIpMapping, error) {
	if params == nil {
		params = &ListSourceIpMappingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSourceIpMapping(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan VoiceV1SourceIpMapping, 1)

	go func() {
		for response != nil {
			responseRecords := response.SourceIpMappings
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListSourceIpMappingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSourceIpMappingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSourceIpMappingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSourceIpMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSourceIpMapping'
type UpdateSourceIpMappingParams struct {
	// The SID of the SIP Domain that the IP Record should be mapped to.
	SipDomainSid *string `json:"SipDomainSid,omitempty"`
}

func (params *UpdateSourceIpMappingParams) SetSipDomainSid(SipDomainSid string) *UpdateSourceIpMappingParams {
	params.SipDomainSid = &SipDomainSid
	return params
}

func (c *ApiService) UpdateSourceIpMapping(Sid string, params *UpdateSourceIpMappingParams) (*VoiceV1SourceIpMapping, error) {
	path := "/v1/SourceIpMappings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SipDomainSid != nil {
		data.Set("SipDomainSid", *params.SipDomainSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1SourceIpMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
