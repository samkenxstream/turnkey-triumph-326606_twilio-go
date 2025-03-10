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

// NotifyV1Binding struct for NotifyV1Binding
type NotifyV1Binding struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The channel-specific address
	Address *string `json:"address,omitempty"`
	// The type of the Binding
	BindingType *string `json:"binding_type,omitempty"`
	// The SID of the Credential resource to be used to send notifications to this Binding
	CredentialSid *string `json:"credential_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Deprecated
	Endpoint *string `json:"endpoint,omitempty"`
	// The `identity` value that identifies the new resource's User
	Identity *string `json:"identity,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The protocol version to use to send the notification
	NotificationProtocolVersion *string `json:"notification_protocol_version,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The list of tags associated with this Binding
	Tags *[]string `json:"tags,omitempty"`
	// The absolute URL of the Binding resource
	Url *string `json:"url,omitempty"`
}
