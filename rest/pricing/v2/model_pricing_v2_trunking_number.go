/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2TrunkingNumber struct for PricingV2TrunkingNumber
type PricingV2TrunkingNumber struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The destination phone number, in E.164 format
	DestinationNumber *string `json:"destination_number,omitempty"`
	// The ISO country code
	IsoCountry           *string                                      `json:"iso_country,omitempty"`
	OriginatingCallPrice *PricingV2TrunkingNumberOriginatingCallPrice `json:"originating_call_price,omitempty"`
	// The origination phone number, in E.164 format
	OriginationNumber *string `json:"origination_number,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit               *string                                                    `json:"price_unit,omitempty"`
	TerminatingPrefixPrices *[]PricingV2TrunkingCountryInstanceTerminatingPrefixPrices `json:"terminating_prefix_prices,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
