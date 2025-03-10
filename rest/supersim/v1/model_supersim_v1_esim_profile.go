/*
 * Twilio - Supersim
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

// SupersimV1EsimProfile struct for SupersimV1EsimProfile
type SupersimV1EsimProfile struct {
	// The SID of the Account to which the eSIM Profile resource belongs
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Identifier of the eUICC that can claim the eSIM Profile
	Eid *string `json:"eid,omitempty"`
	// Code indicating the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state
	ErrorCode *string `json:"error_code,omitempty"`
	// Error message describing the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state
	ErrorMessage *string `json:"error_message,omitempty"`
	// The ICCID associated with the Sim resource
	Iccid *string `json:"iccid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Sim resource that this eSIM Profile controls
	SimSid *string `json:"sim_sid,omitempty"`
	// Address of the SM-DP+ server from which the Profile will be downloaded
	SmdpPlusAddress *string `json:"smdp_plus_address,omitempty"`
	// The status of the eSIM Profile
	Status *string `json:"status,omitempty"`
	// The absolute URL of the eSIM Profile resource
	Url *string `json:"url,omitempty"`
}
