/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Address struct for ApiV2010Address
type ApiV2010Address struct {
	// The SID of the Account that is responsible for the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The city in which the address is located
	City *string `json:"city,omitempty"`
	// The name associated with the address
	CustomerName *string `json:"customer_name,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// Whether emergency calling has been enabled on this number
	EmergencyEnabled *bool `json:"emergency_enabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The ISO country code of the address
	IsoCountry *string `json:"iso_country,omitempty"`
	// The postal code of the address
	PostalCode *string `json:"postal_code,omitempty"`
	// The state or region of the address
	Region *string `json:"region,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The number and street address of the address
	Street *string `json:"street,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// Whether the address has been validated to comply with local regulation
	Validated *bool `json:"validated,omitempty"`
	// Whether the address has been verified to comply with regulation
	Verified *bool `json:"verified,omitempty"`
}
