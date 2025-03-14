/*
 * Twilio - Trusthub
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

// TrusthubV1CustomerProfileEntityAssignment struct for TrusthubV1CustomerProfileEntityAssignment
type TrusthubV1CustomerProfileEntityAssignment struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string that identifies the CustomerProfile resource.
	CustomerProfileSid *string `json:"customer_profile_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The sid of an object bag
	ObjectSid *string `json:"object_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Identity resource
	Url *string `json:"url,omitempty"`
}
