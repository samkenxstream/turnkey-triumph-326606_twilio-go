/*
 * Twilio - Serverless
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

// ServerlessV1Environment struct for ServerlessV1Environment
type ServerlessV1Environment struct {
	// The SID of the Account that created the Environment resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the build deployed in the environment
	BuildSid *string `json:"build_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Environment resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Environment resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The domain name for all Functions and Assets deployed in the Environment
	DomainName *string `json:"domain_name,omitempty"`
	// A URL-friendly name that represents the environment
	DomainSuffix *string `json:"domain_suffix,omitempty"`
	// The URLs of the Environment resource's nested resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the Service that the Environment resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Environment resource
	Sid *string `json:"sid,omitempty"`
	// A user-defined string that uniquely identifies the Environment resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Environment resource
	Url *string `json:"url,omitempty"`
}
