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
	"time"
)

// MediaV1MediaProcessor struct for MediaV1MediaProcessor
type MediaV1MediaProcessor struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The reason why a MediaProcessor ended
	EndedReason *string `json:"ended_reason,omitempty"`
	// The Media Extension name or URL
	Extension *string `json:"extension,omitempty"`
	// The Media Extension context
	ExtensionContext *string `json:"extension_context,omitempty"`
	// Maximum MediaProcessor duration in seconds
	MaxDuration *int `json:"max_duration,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the MediaProcessor
	Status *string `json:"status,omitempty"`
	// The URL to which Twilio will send MediaProcessor event updates
	StatusCallback *string `json:"status_callback,omitempty"`
	// The HTTP method Twilio should use to call the `status_callback` URL
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
