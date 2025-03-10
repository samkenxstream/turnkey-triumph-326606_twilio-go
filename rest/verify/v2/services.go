/*
 * Twilio - Verify
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
	// The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
	CodeLength *int `json:"CodeLength,omitempty"`
	// Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
	CustomCodeEnabled *bool `json:"CustomCodeEnabled,omitempty"`
	// The default message [template](https://www.twilio.com/docs/verify/api/templates). Will be used for all SMS verifications unless explicitly overriden. SMS channel only.
	DefaultTemplateSid *string `json:"DefaultTemplateSid,omitempty"`
	// Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: `Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code`
	DoNotShareWarningEnabled *bool `json:"DoNotShareWarningEnabled,omitempty"`
	// Whether to ask the user to press a number before delivering the verify code in a phone call.
	DtmfInputRequired *bool `json:"DtmfInputRequired,omitempty"`
	// A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.**
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to perform a lookup with each verification started and return info about the phone number.
	LookupEnabled *bool `json:"LookupEnabled,omitempty"`
	// Whether to pass PSD2 transaction parameters when starting a verification.
	Psd2Enabled *bool `json:"Psd2Enabled,omitempty"`
	// Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
	PushApnCredentialSid *string `json:"Push.ApnCredentialSid,omitempty"`
	// Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
	PushFcmCredentialSid *string `json:"Push.FcmCredentialSid,omitempty"`
	// Optional configuration for the Push factors. If true, include the date in the Challenge's reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true
	PushIncludeDate *bool `json:"Push.IncludeDate,omitempty"`
	// Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
	SkipSmsToLandlines *bool `json:"SkipSmsToLandlines,omitempty"`
	// Optional configuration for the TOTP factors. Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6
	TotpCodeLength *int `json:"Totp.CodeLength,omitempty"`
	// Optional configuration for the TOTP factors. Set TOTP Issuer for this service. This will allow to configure the issuer of the TOTP URI. Defaults to the service friendly name if not provided.
	TotpIssuer *string `json:"Totp.Issuer,omitempty"`
	// Optional configuration for the TOTP factors. The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1
	TotpSkew *int `json:"Totp.Skew,omitempty"`
	// Optional configuration for the TOTP factors. Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds
	TotpTimeStep *int `json:"Totp.TimeStep,omitempty"`
	// The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
	TtsName *string `json:"TtsName,omitempty"`
}

func (params *CreateServiceParams) SetCodeLength(CodeLength int) *CreateServiceParams {
	params.CodeLength = &CodeLength
	return params
}
func (params *CreateServiceParams) SetCustomCodeEnabled(CustomCodeEnabled bool) *CreateServiceParams {
	params.CustomCodeEnabled = &CustomCodeEnabled
	return params
}
func (params *CreateServiceParams) SetDefaultTemplateSid(DefaultTemplateSid string) *CreateServiceParams {
	params.DefaultTemplateSid = &DefaultTemplateSid
	return params
}
func (params *CreateServiceParams) SetDoNotShareWarningEnabled(DoNotShareWarningEnabled bool) *CreateServiceParams {
	params.DoNotShareWarningEnabled = &DoNotShareWarningEnabled
	return params
}
func (params *CreateServiceParams) SetDtmfInputRequired(DtmfInputRequired bool) *CreateServiceParams {
	params.DtmfInputRequired = &DtmfInputRequired
	return params
}
func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceParams) SetLookupEnabled(LookupEnabled bool) *CreateServiceParams {
	params.LookupEnabled = &LookupEnabled
	return params
}
func (params *CreateServiceParams) SetPsd2Enabled(Psd2Enabled bool) *CreateServiceParams {
	params.Psd2Enabled = &Psd2Enabled
	return params
}
func (params *CreateServiceParams) SetPushApnCredentialSid(PushApnCredentialSid string) *CreateServiceParams {
	params.PushApnCredentialSid = &PushApnCredentialSid
	return params
}
func (params *CreateServiceParams) SetPushFcmCredentialSid(PushFcmCredentialSid string) *CreateServiceParams {
	params.PushFcmCredentialSid = &PushFcmCredentialSid
	return params
}
func (params *CreateServiceParams) SetPushIncludeDate(PushIncludeDate bool) *CreateServiceParams {
	params.PushIncludeDate = &PushIncludeDate
	return params
}
func (params *CreateServiceParams) SetSkipSmsToLandlines(SkipSmsToLandlines bool) *CreateServiceParams {
	params.SkipSmsToLandlines = &SkipSmsToLandlines
	return params
}
func (params *CreateServiceParams) SetTotpCodeLength(TotpCodeLength int) *CreateServiceParams {
	params.TotpCodeLength = &TotpCodeLength
	return params
}
func (params *CreateServiceParams) SetTotpIssuer(TotpIssuer string) *CreateServiceParams {
	params.TotpIssuer = &TotpIssuer
	return params
}
func (params *CreateServiceParams) SetTotpSkew(TotpSkew int) *CreateServiceParams {
	params.TotpSkew = &TotpSkew
	return params
}
func (params *CreateServiceParams) SetTotpTimeStep(TotpTimeStep int) *CreateServiceParams {
	params.TotpTimeStep = &TotpTimeStep
	return params
}
func (params *CreateServiceParams) SetTtsName(TtsName string) *CreateServiceParams {
	params.TtsName = &TtsName
	return params
}

// Create a new Verification Service.
func (c *ApiService) CreateService(params *CreateServiceParams) (*VerifyV2Service, error) {
	path := "/v2/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CodeLength != nil {
		data.Set("CodeLength", fmt.Sprint(*params.CodeLength))
	}
	if params != nil && params.CustomCodeEnabled != nil {
		data.Set("CustomCodeEnabled", fmt.Sprint(*params.CustomCodeEnabled))
	}
	if params != nil && params.DefaultTemplateSid != nil {
		data.Set("DefaultTemplateSid", *params.DefaultTemplateSid)
	}
	if params != nil && params.DoNotShareWarningEnabled != nil {
		data.Set("DoNotShareWarningEnabled", fmt.Sprint(*params.DoNotShareWarningEnabled))
	}
	if params != nil && params.DtmfInputRequired != nil {
		data.Set("DtmfInputRequired", fmt.Sprint(*params.DtmfInputRequired))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.LookupEnabled != nil {
		data.Set("LookupEnabled", fmt.Sprint(*params.LookupEnabled))
	}
	if params != nil && params.Psd2Enabled != nil {
		data.Set("Psd2Enabled", fmt.Sprint(*params.Psd2Enabled))
	}
	if params != nil && params.PushApnCredentialSid != nil {
		data.Set("Push.ApnCredentialSid", *params.PushApnCredentialSid)
	}
	if params != nil && params.PushFcmCredentialSid != nil {
		data.Set("Push.FcmCredentialSid", *params.PushFcmCredentialSid)
	}
	if params != nil && params.PushIncludeDate != nil {
		data.Set("Push.IncludeDate", fmt.Sprint(*params.PushIncludeDate))
	}
	if params != nil && params.SkipSmsToLandlines != nil {
		data.Set("SkipSmsToLandlines", fmt.Sprint(*params.SkipSmsToLandlines))
	}
	if params != nil && params.TotpCodeLength != nil {
		data.Set("Totp.CodeLength", fmt.Sprint(*params.TotpCodeLength))
	}
	if params != nil && params.TotpIssuer != nil {
		data.Set("Totp.Issuer", *params.TotpIssuer)
	}
	if params != nil && params.TotpSkew != nil {
		data.Set("Totp.Skew", fmt.Sprint(*params.TotpSkew))
	}
	if params != nil && params.TotpTimeStep != nil {
		data.Set("Totp.TimeStep", fmt.Sprint(*params.TotpTimeStep))
	}
	if params != nil && params.TtsName != nil {
		data.Set("TtsName", *params.TtsName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Verification Service Instance.
func (c *ApiService) DeleteService(Sid string) error {
	path := "/v2/Services/{Sid}"
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

// Fetch specific Verification Service Instance.
func (c *ApiService) FetchService(Sid string) (*VerifyV2Service, error) {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Service{}
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
	path := "/v2/Services"

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
func (c *ApiService) ListService(params *ListServiceParams) ([]VerifyV2Service, error) {
	response, err := c.StreamService(params)
	if err != nil {
		return nil, err
	}

	records := make([]VerifyV2Service, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan VerifyV2Service, error) {
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
	channel := make(chan VerifyV2Service, 1)

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
	// The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
	CodeLength *int `json:"CodeLength,omitempty"`
	// Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
	CustomCodeEnabled *bool `json:"CustomCodeEnabled,omitempty"`
	// The default message [template](https://www.twilio.com/docs/verify/api/templates). Will be used for all SMS verifications unless explicitly overriden. SMS channel only.
	DefaultTemplateSid *string `json:"DefaultTemplateSid,omitempty"`
	// Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.**
	DoNotShareWarningEnabled *bool `json:"DoNotShareWarningEnabled,omitempty"`
	// Whether to ask the user to press a number before delivering the verify code in a phone call.
	DtmfInputRequired *bool `json:"DtmfInputRequired,omitempty"`
	// A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.**
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to perform a lookup with each verification started and return info about the phone number.
	LookupEnabled *bool `json:"LookupEnabled,omitempty"`
	// Whether to pass PSD2 transaction parameters when starting a verification.
	Psd2Enabled *bool `json:"Psd2Enabled,omitempty"`
	// Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
	PushApnCredentialSid *string `json:"Push.ApnCredentialSid,omitempty"`
	// Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
	PushFcmCredentialSid *string `json:"Push.FcmCredentialSid,omitempty"`
	// Optional configuration for the Push factors. If true, include the date in the Challenge's reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true
	PushIncludeDate *bool `json:"Push.IncludeDate,omitempty"`
	// Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
	SkipSmsToLandlines *bool `json:"SkipSmsToLandlines,omitempty"`
	// Optional configuration for the TOTP factors. Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6
	TotpCodeLength *int `json:"Totp.CodeLength,omitempty"`
	// Optional configuration for the TOTP factors. Set TOTP Issuer for this service. This will allow to configure the issuer of the TOTP URI.
	TotpIssuer *string `json:"Totp.Issuer,omitempty"`
	// Optional configuration for the TOTP factors. The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1
	TotpSkew *int `json:"Totp.Skew,omitempty"`
	// Optional configuration for the TOTP factors. Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds
	TotpTimeStep *int `json:"Totp.TimeStep,omitempty"`
	// The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
	TtsName *string `json:"TtsName,omitempty"`
}

func (params *UpdateServiceParams) SetCodeLength(CodeLength int) *UpdateServiceParams {
	params.CodeLength = &CodeLength
	return params
}
func (params *UpdateServiceParams) SetCustomCodeEnabled(CustomCodeEnabled bool) *UpdateServiceParams {
	params.CustomCodeEnabled = &CustomCodeEnabled
	return params
}
func (params *UpdateServiceParams) SetDefaultTemplateSid(DefaultTemplateSid string) *UpdateServiceParams {
	params.DefaultTemplateSid = &DefaultTemplateSid
	return params
}
func (params *UpdateServiceParams) SetDoNotShareWarningEnabled(DoNotShareWarningEnabled bool) *UpdateServiceParams {
	params.DoNotShareWarningEnabled = &DoNotShareWarningEnabled
	return params
}
func (params *UpdateServiceParams) SetDtmfInputRequired(DtmfInputRequired bool) *UpdateServiceParams {
	params.DtmfInputRequired = &DtmfInputRequired
	return params
}
func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetLookupEnabled(LookupEnabled bool) *UpdateServiceParams {
	params.LookupEnabled = &LookupEnabled
	return params
}
func (params *UpdateServiceParams) SetPsd2Enabled(Psd2Enabled bool) *UpdateServiceParams {
	params.Psd2Enabled = &Psd2Enabled
	return params
}
func (params *UpdateServiceParams) SetPushApnCredentialSid(PushApnCredentialSid string) *UpdateServiceParams {
	params.PushApnCredentialSid = &PushApnCredentialSid
	return params
}
func (params *UpdateServiceParams) SetPushFcmCredentialSid(PushFcmCredentialSid string) *UpdateServiceParams {
	params.PushFcmCredentialSid = &PushFcmCredentialSid
	return params
}
func (params *UpdateServiceParams) SetPushIncludeDate(PushIncludeDate bool) *UpdateServiceParams {
	params.PushIncludeDate = &PushIncludeDate
	return params
}
func (params *UpdateServiceParams) SetSkipSmsToLandlines(SkipSmsToLandlines bool) *UpdateServiceParams {
	params.SkipSmsToLandlines = &SkipSmsToLandlines
	return params
}
func (params *UpdateServiceParams) SetTotpCodeLength(TotpCodeLength int) *UpdateServiceParams {
	params.TotpCodeLength = &TotpCodeLength
	return params
}
func (params *UpdateServiceParams) SetTotpIssuer(TotpIssuer string) *UpdateServiceParams {
	params.TotpIssuer = &TotpIssuer
	return params
}
func (params *UpdateServiceParams) SetTotpSkew(TotpSkew int) *UpdateServiceParams {
	params.TotpSkew = &TotpSkew
	return params
}
func (params *UpdateServiceParams) SetTotpTimeStep(TotpTimeStep int) *UpdateServiceParams {
	params.TotpTimeStep = &TotpTimeStep
	return params
}
func (params *UpdateServiceParams) SetTtsName(TtsName string) *UpdateServiceParams {
	params.TtsName = &TtsName
	return params
}

// Update a specific Verification Service.
func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*VerifyV2Service, error) {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CodeLength != nil {
		data.Set("CodeLength", fmt.Sprint(*params.CodeLength))
	}
	if params != nil && params.CustomCodeEnabled != nil {
		data.Set("CustomCodeEnabled", fmt.Sprint(*params.CustomCodeEnabled))
	}
	if params != nil && params.DefaultTemplateSid != nil {
		data.Set("DefaultTemplateSid", *params.DefaultTemplateSid)
	}
	if params != nil && params.DoNotShareWarningEnabled != nil {
		data.Set("DoNotShareWarningEnabled", fmt.Sprint(*params.DoNotShareWarningEnabled))
	}
	if params != nil && params.DtmfInputRequired != nil {
		data.Set("DtmfInputRequired", fmt.Sprint(*params.DtmfInputRequired))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.LookupEnabled != nil {
		data.Set("LookupEnabled", fmt.Sprint(*params.LookupEnabled))
	}
	if params != nil && params.Psd2Enabled != nil {
		data.Set("Psd2Enabled", fmt.Sprint(*params.Psd2Enabled))
	}
	if params != nil && params.PushApnCredentialSid != nil {
		data.Set("Push.ApnCredentialSid", *params.PushApnCredentialSid)
	}
	if params != nil && params.PushFcmCredentialSid != nil {
		data.Set("Push.FcmCredentialSid", *params.PushFcmCredentialSid)
	}
	if params != nil && params.PushIncludeDate != nil {
		data.Set("Push.IncludeDate", fmt.Sprint(*params.PushIncludeDate))
	}
	if params != nil && params.SkipSmsToLandlines != nil {
		data.Set("SkipSmsToLandlines", fmt.Sprint(*params.SkipSmsToLandlines))
	}
	if params != nil && params.TotpCodeLength != nil {
		data.Set("Totp.CodeLength", fmt.Sprint(*params.TotpCodeLength))
	}
	if params != nil && params.TotpIssuer != nil {
		data.Set("Totp.Issuer", *params.TotpIssuer)
	}
	if params != nil && params.TotpSkew != nil {
		data.Set("Totp.Skew", fmt.Sprint(*params.TotpSkew))
	}
	if params != nil && params.TotpTimeStep != nil {
		data.Set("Totp.TimeStep", fmt.Sprint(*params.TotpTimeStep))
	}
	if params != nil && params.TtsName != nil {
		data.Set("TtsName", *params.TtsName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
