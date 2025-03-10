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

// AutopilotV1Sample struct for AutopilotV1Sample
type AutopilotV1Sample struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the Task associated with the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An ISO language-country string that specifies the language used for the sample
	Language *string `json:"language,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The communication channel from which the sample was captured
	SourceChannel *string `json:"source_channel,omitempty"`
	// The text example of how end users might express the task
	TaggedText *string `json:"tagged_text,omitempty"`
	// The SID of the Task associated with the resource
	TaskSid *string `json:"task_sid,omitempty"`
	// The absolute URL of the Sample resource
	Url *string `json:"url,omitempty"`
}
