/*
 * Twilio - Notify
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

// NotifyV1Notification struct for NotifyV1Notification
type NotifyV1Notification struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The actions to display for the notification
	Action *string `json:"action,omitempty"`
	// Deprecated
	Alexa *interface{} `json:"alexa,omitempty"`
	// The APNS-specific payload that overrides corresponding attributes in a generic payload for APNS Bindings
	Apn *interface{} `json:"apn,omitempty"`
	// The notification body text
	Body *string `json:"body,omitempty"`
	// The custom key-value pairs of the notification's payload
	Data *interface{} `json:"data,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// Deprecated
	FacebookMessenger *interface{} `json:"facebook_messenger,omitempty"`
	// The FCM-specific payload that overrides corresponding attributes in generic payload for FCM Bindings
	Fcm *interface{} `json:"fcm,omitempty"`
	// The GCM-specific payload that overrides corresponding attributes in generic payload for GCM Bindings
	Gcm *interface{} `json:"gcm,omitempty"`
	// The list of identity values of the Users to notify
	Identities *[]string `json:"identities,omitempty"`
	// The priority of the notification
	Priority *string `json:"priority,omitempty"`
	// The list of Segments to notify
	Segments *[]string `json:"segments,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SMS-specific payload that overrides corresponding attributes in generic payload for SMS Bindings
	Sms *interface{} `json:"sms,omitempty"`
	// The name of the sound to be played for the notification
	Sound *string `json:"sound,omitempty"`
	// The tags that select the Bindings to notify
	Tags *[]string `json:"tags,omitempty"`
	// The notification title
	Title *string `json:"title,omitempty"`
	// How long, in seconds, the notification is valid
	Ttl *int `json:"ttl,omitempty"`
}
