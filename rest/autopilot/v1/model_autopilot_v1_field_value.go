/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// AutopilotV1FieldValue struct for AutopilotV1FieldValue
type AutopilotV1FieldValue struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the FieldType associated with the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Field Type associated with the Field Value
	FieldTypeSid *string `json:"field_type_sid,omitempty"`
	// The ISO language-country tag that identifies the language of the value
	Language *string `json:"language,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The word for which the field value is a synonym of
	SynonymOf *string `json:"synonym_of,omitempty"`
	// The absolute URL of the FieldValue resource
	Url *string `json:"url,omitempty"`
	// The Field Value data
	Value *string `json:"value,omitempty"`
}
