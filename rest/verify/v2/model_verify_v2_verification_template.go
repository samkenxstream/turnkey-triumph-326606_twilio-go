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

// VerifyV2VerificationTemplate struct for VerifyV2VerificationTemplate
type VerifyV2VerificationTemplate struct {
	// Account Sid
	AccountSid *string `json:"account_sid,omitempty"`
	// A string to describe the verification template
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A string that uniquely identifies this Template
	Sid *string `json:"sid,omitempty"`
	// Object with the template translations.
	Translations *interface{} `json:"translations,omitempty"`
}
