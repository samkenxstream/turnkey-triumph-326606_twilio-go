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

// ConversationsV1Credential struct for ConversationsV1Credential
type ConversationsV1Credential struct {
	// The unique ID of the Account responsible for this credential.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The human-readable name of this credential.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// [APN only] Whether to send the credential to sandbox APNs.
	Sandbox *string `json:"sandbox,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The type of push-notification service the credential is for.
	Type *string `json:"type,omitempty"`
	// An absolute URL for this credential.
	Url *string `json:"url,omitempty"`
}
