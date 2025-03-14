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

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// PricingV2TrunkingCountryInstanceTerminatingPrefixPrices struct for PricingV2TrunkingCountryInstanceTerminatingPrefixPrices
type PricingV2TrunkingCountryInstanceTerminatingPrefixPrices struct {
	BasePrice           float32  `json:"base_price,omitempty"`
	CurrentPrice        float32  `json:"current_price,omitempty"`
	DestinationPrefixes []string `json:"destination_prefixes,omitempty"`
	FriendlyName        string   `json:"friendly_name,omitempty"`
	OriginationPrefixes []string `json:"origination_prefixes,omitempty"`
}

func (response *PricingV2TrunkingCountryInstanceTerminatingPrefixPrices) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		BasePrice           interface{} `json:"base_price"`
		CurrentPrice        interface{} `json:"current_price"`
		DestinationPrefixes []string    `json:"destination_prefixes"`
		FriendlyName        string      `json:"friendly_name"`
		OriginationPrefixes []string    `json:"origination_prefixes"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = PricingV2TrunkingCountryInstanceTerminatingPrefixPrices{
		DestinationPrefixes: raw.DestinationPrefixes,
		FriendlyName:        raw.FriendlyName,
		OriginationPrefixes: raw.OriginationPrefixes,
	}

	responseBasePrice, err := client.UnmarshalFloat32(&raw.BasePrice)
	if err != nil {
		return err
	}
	response.BasePrice = *responseBasePrice

	responseCurrentPrice, err := client.UnmarshalFloat32(&raw.CurrentPrice)
	if err != nil {
		return err
	}
	response.CurrentPrice = *responseCurrentPrice

	return
}
