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

// SupersimV1BillingPeriod struct for SupersimV1BillingPeriod
type SupersimV1BillingPeriod struct {
	// The SID of the Account the Super SIM belongs to
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The end time of the Billing Period
	EndTime *time.Time `json:"end_time,omitempty"`
	// The type of the Billing Period
	PeriodType *string `json:"period_type,omitempty"`
	// The SID of the Billing Period
	Sid *string `json:"sid,omitempty"`
	// The SID of the Super SIM the Billing Period belongs to
	SimSid *string `json:"sim_sid,omitempty"`
	// The start time of the Billing Period
	StartTime *time.Time `json:"start_time,omitempty"`
}
