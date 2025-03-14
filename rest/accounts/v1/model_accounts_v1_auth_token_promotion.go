/*
 * Twilio - Accounts
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

// AccountsV1AuthTokenPromotion struct for AccountsV1AuthTokenPromotion
type AccountsV1AuthTokenPromotion struct {
	// The SID of the Account that the secondary Auth Token was created for
	AccountSid *string `json:"account_sid,omitempty"`
	// The promoted Auth Token
	AuthToken *string `json:"auth_token,omitempty"`
	// The ISO 8601 formatted date and time in UTC when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 formatted date and time in UTC when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URI for this resource, relative to `https://accounts.twilio.com`
	Url *string `json:"url,omitempty"`
}
