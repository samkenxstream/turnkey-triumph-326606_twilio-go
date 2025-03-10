/*
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
	AreaCodeGeomatch *bool `json:"AreaCodeGeomatch,omitempty"`
	// The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`.
	FallbackMethod *string `json:"FallbackMethod,omitempty"`
	// Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
	FallbackToLongCode *bool `json:"FallbackToLongCode,omitempty"`
	// The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service.
	FallbackUrl *string `json:"FallbackUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`.
	InboundMethod *string `json:"InboundMethod,omitempty"`
	// The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service.
	InboundRequestUrl *string `json:"InboundRequestUrl,omitempty"`
	// Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
	MmsConverter *bool `json:"MmsConverter,omitempty"`
	// Reserved.
	ScanMessageContent *string `json:"ScanMessageContent,omitempty"`
	// Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
	SmartEncoding *bool `json:"SmartEncoding,omitempty"`
	// The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
	StickySender *bool `json:"StickySender,omitempty"`
	// Reserved.
	SynchronousValidation *bool `json:"SynchronousValidation,omitempty"`
	// A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service.
	UseInboundWebhookOnNumber *bool `json:"UseInboundWebhookOnNumber,omitempty"`
	// A string that describes the scenario in which the Messaging Service will be used. Examples: [notification, marketing, verification, poll ..].
	Usecase *string `json:"Usecase,omitempty"`
	// How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
	ValidityPeriod *int `json:"ValidityPeriod,omitempty"`
}

func (params *CreateServiceParams) SetAreaCodeGeomatch(AreaCodeGeomatch bool) *CreateServiceParams {
	params.AreaCodeGeomatch = &AreaCodeGeomatch
	return params
}
func (params *CreateServiceParams) SetFallbackMethod(FallbackMethod string) *CreateServiceParams {
	params.FallbackMethod = &FallbackMethod
	return params
}
func (params *CreateServiceParams) SetFallbackToLongCode(FallbackToLongCode bool) *CreateServiceParams {
	params.FallbackToLongCode = &FallbackToLongCode
	return params
}
func (params *CreateServiceParams) SetFallbackUrl(FallbackUrl string) *CreateServiceParams {
	params.FallbackUrl = &FallbackUrl
	return params
}
func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceParams) SetInboundMethod(InboundMethod string) *CreateServiceParams {
	params.InboundMethod = &InboundMethod
	return params
}
func (params *CreateServiceParams) SetInboundRequestUrl(InboundRequestUrl string) *CreateServiceParams {
	params.InboundRequestUrl = &InboundRequestUrl
	return params
}
func (params *CreateServiceParams) SetMmsConverter(MmsConverter bool) *CreateServiceParams {
	params.MmsConverter = &MmsConverter
	return params
}
func (params *CreateServiceParams) SetScanMessageContent(ScanMessageContent string) *CreateServiceParams {
	params.ScanMessageContent = &ScanMessageContent
	return params
}
func (params *CreateServiceParams) SetSmartEncoding(SmartEncoding bool) *CreateServiceParams {
	params.SmartEncoding = &SmartEncoding
	return params
}
func (params *CreateServiceParams) SetStatusCallback(StatusCallback string) *CreateServiceParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateServiceParams) SetStickySender(StickySender bool) *CreateServiceParams {
	params.StickySender = &StickySender
	return params
}
func (params *CreateServiceParams) SetSynchronousValidation(SynchronousValidation bool) *CreateServiceParams {
	params.SynchronousValidation = &SynchronousValidation
	return params
}
func (params *CreateServiceParams) SetUseInboundWebhookOnNumber(UseInboundWebhookOnNumber bool) *CreateServiceParams {
	params.UseInboundWebhookOnNumber = &UseInboundWebhookOnNumber
	return params
}
func (params *CreateServiceParams) SetUsecase(Usecase string) *CreateServiceParams {
	params.Usecase = &Usecase
	return params
}
func (params *CreateServiceParams) SetValidityPeriod(ValidityPeriod int) *CreateServiceParams {
	params.ValidityPeriod = &ValidityPeriod
	return params
}

func (c *ApiService) CreateService(params *CreateServiceParams) (*MessagingV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AreaCodeGeomatch != nil {
		data.Set("AreaCodeGeomatch", fmt.Sprint(*params.AreaCodeGeomatch))
	}
	if params != nil && params.FallbackMethod != nil {
		data.Set("FallbackMethod", *params.FallbackMethod)
	}
	if params != nil && params.FallbackToLongCode != nil {
		data.Set("FallbackToLongCode", fmt.Sprint(*params.FallbackToLongCode))
	}
	if params != nil && params.FallbackUrl != nil {
		data.Set("FallbackUrl", *params.FallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.InboundMethod != nil {
		data.Set("InboundMethod", *params.InboundMethod)
	}
	if params != nil && params.InboundRequestUrl != nil {
		data.Set("InboundRequestUrl", *params.InboundRequestUrl)
	}
	if params != nil && params.MmsConverter != nil {
		data.Set("MmsConverter", fmt.Sprint(*params.MmsConverter))
	}
	if params != nil && params.ScanMessageContent != nil {
		data.Set("ScanMessageContent", *params.ScanMessageContent)
	}
	if params != nil && params.SmartEncoding != nil {
		data.Set("SmartEncoding", fmt.Sprint(*params.SmartEncoding))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StickySender != nil {
		data.Set("StickySender", fmt.Sprint(*params.StickySender))
	}
	if params != nil && params.SynchronousValidation != nil {
		data.Set("SynchronousValidation", fmt.Sprint(*params.SynchronousValidation))
	}
	if params != nil && params.UseInboundWebhookOnNumber != nil {
		data.Set("UseInboundWebhookOnNumber", fmt.Sprint(*params.UseInboundWebhookOnNumber))
	}
	if params != nil && params.Usecase != nil {
		data.Set("Usecase", *params.Usecase)
	}
	if params != nil && params.ValidityPeriod != nil {
		data.Set("ValidityPeriod", fmt.Sprint(*params.ValidityPeriod))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
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

func (c *ApiService) FetchService(Sid string) (*MessagingV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParams) SetLimit(Limit int) *ListServiceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken, pageNumber string) (*ListServiceResponse, error) {
	path := "/v1/Services"

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

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]MessagingV1Service, error) {
	response, err := c.StreamService(params)
	if err != nil {
		return nil, err
	}

	records := make([]MessagingV1Service, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan MessagingV1Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1Service, 1)

	go func() {
		for response != nil {
			responseRecords := response.Services
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListServiceResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	// Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
	AreaCodeGeomatch *bool `json:"AreaCodeGeomatch,omitempty"`
	// The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`.
	FallbackMethod *string `json:"FallbackMethod,omitempty"`
	// Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
	FallbackToLongCode *bool `json:"FallbackToLongCode,omitempty"`
	// The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service.
	FallbackUrl *string `json:"FallbackUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`.
	InboundMethod *string `json:"InboundMethod,omitempty"`
	// The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service.
	InboundRequestUrl *string `json:"InboundRequestUrl,omitempty"`
	// Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
	MmsConverter *bool `json:"MmsConverter,omitempty"`
	// Reserved.
	ScanMessageContent *string `json:"ScanMessageContent,omitempty"`
	// Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
	SmartEncoding *bool `json:"SmartEncoding,omitempty"`
	// The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
	StickySender *bool `json:"StickySender,omitempty"`
	// Reserved.
	SynchronousValidation *bool `json:"SynchronousValidation,omitempty"`
	// A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service.
	UseInboundWebhookOnNumber *bool `json:"UseInboundWebhookOnNumber,omitempty"`
	// A string that describes the scenario in which the Messaging Service will be used. Examples: [notification, marketing, verification, poll ..]
	Usecase *string `json:"Usecase,omitempty"`
	// How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
	ValidityPeriod *int `json:"ValidityPeriod,omitempty"`
}

func (params *UpdateServiceParams) SetAreaCodeGeomatch(AreaCodeGeomatch bool) *UpdateServiceParams {
	params.AreaCodeGeomatch = &AreaCodeGeomatch
	return params
}
func (params *UpdateServiceParams) SetFallbackMethod(FallbackMethod string) *UpdateServiceParams {
	params.FallbackMethod = &FallbackMethod
	return params
}
func (params *UpdateServiceParams) SetFallbackToLongCode(FallbackToLongCode bool) *UpdateServiceParams {
	params.FallbackToLongCode = &FallbackToLongCode
	return params
}
func (params *UpdateServiceParams) SetFallbackUrl(FallbackUrl string) *UpdateServiceParams {
	params.FallbackUrl = &FallbackUrl
	return params
}
func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetInboundMethod(InboundMethod string) *UpdateServiceParams {
	params.InboundMethod = &InboundMethod
	return params
}
func (params *UpdateServiceParams) SetInboundRequestUrl(InboundRequestUrl string) *UpdateServiceParams {
	params.InboundRequestUrl = &InboundRequestUrl
	return params
}
func (params *UpdateServiceParams) SetMmsConverter(MmsConverter bool) *UpdateServiceParams {
	params.MmsConverter = &MmsConverter
	return params
}
func (params *UpdateServiceParams) SetScanMessageContent(ScanMessageContent string) *UpdateServiceParams {
	params.ScanMessageContent = &ScanMessageContent
	return params
}
func (params *UpdateServiceParams) SetSmartEncoding(SmartEncoding bool) *UpdateServiceParams {
	params.SmartEncoding = &SmartEncoding
	return params
}
func (params *UpdateServiceParams) SetStatusCallback(StatusCallback string) *UpdateServiceParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateServiceParams) SetStickySender(StickySender bool) *UpdateServiceParams {
	params.StickySender = &StickySender
	return params
}
func (params *UpdateServiceParams) SetSynchronousValidation(SynchronousValidation bool) *UpdateServiceParams {
	params.SynchronousValidation = &SynchronousValidation
	return params
}
func (params *UpdateServiceParams) SetUseInboundWebhookOnNumber(UseInboundWebhookOnNumber bool) *UpdateServiceParams {
	params.UseInboundWebhookOnNumber = &UseInboundWebhookOnNumber
	return params
}
func (params *UpdateServiceParams) SetUsecase(Usecase string) *UpdateServiceParams {
	params.Usecase = &Usecase
	return params
}
func (params *UpdateServiceParams) SetValidityPeriod(ValidityPeriod int) *UpdateServiceParams {
	params.ValidityPeriod = &ValidityPeriod
	return params
}

func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*MessagingV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AreaCodeGeomatch != nil {
		data.Set("AreaCodeGeomatch", fmt.Sprint(*params.AreaCodeGeomatch))
	}
	if params != nil && params.FallbackMethod != nil {
		data.Set("FallbackMethod", *params.FallbackMethod)
	}
	if params != nil && params.FallbackToLongCode != nil {
		data.Set("FallbackToLongCode", fmt.Sprint(*params.FallbackToLongCode))
	}
	if params != nil && params.FallbackUrl != nil {
		data.Set("FallbackUrl", *params.FallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.InboundMethod != nil {
		data.Set("InboundMethod", *params.InboundMethod)
	}
	if params != nil && params.InboundRequestUrl != nil {
		data.Set("InboundRequestUrl", *params.InboundRequestUrl)
	}
	if params != nil && params.MmsConverter != nil {
		data.Set("MmsConverter", fmt.Sprint(*params.MmsConverter))
	}
	if params != nil && params.ScanMessageContent != nil {
		data.Set("ScanMessageContent", *params.ScanMessageContent)
	}
	if params != nil && params.SmartEncoding != nil {
		data.Set("SmartEncoding", fmt.Sprint(*params.SmartEncoding))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StickySender != nil {
		data.Set("StickySender", fmt.Sprint(*params.StickySender))
	}
	if params != nil && params.SynchronousValidation != nil {
		data.Set("SynchronousValidation", fmt.Sprint(*params.SynchronousValidation))
	}
	if params != nil && params.UseInboundWebhookOnNumber != nil {
		data.Set("UseInboundWebhookOnNumber", fmt.Sprint(*params.UseInboundWebhookOnNumber))
	}
	if params != nil && params.Usecase != nil {
		data.Set("Usecase", *params.Usecase)
	}
	if params != nil && params.ValidityPeriod != nil {
		data.Set("ValidityPeriod", fmt.Sprint(*params.ValidityPeriod))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
