/*
 * Twilio - Conversations
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

// ConversationsV1ServiceBinding struct for ConversationsV1ServiceBinding
type ConversationsV1ServiceBinding struct {
	// The unique ID of the Account responsible for this binding.
	AccountSid *string `json:"account_sid,omitempty"`
	// The push technology to use for the binding.
	BindingType *string `json:"binding_type,omitempty"`
	// The SID of the Conversation Service that the resource is associated with.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The SID of the Credential for the binding.
	CredentialSid *string `json:"credential_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique endpoint identifier for the Binding.
	Endpoint *string `json:"endpoint,omitempty"`
	// The identity of Conversation User associated with this binding.
	Identity *string `json:"identity,omitempty"`
	// The Conversation message types the binding is subscribed to.
	MessageTypes *[]string `json:"message_types,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// An absolute URL for this binding.
	Url *string `json:"url,omitempty"`
}
