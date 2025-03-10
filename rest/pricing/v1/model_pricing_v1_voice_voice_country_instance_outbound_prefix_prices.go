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

// PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices struct for PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices
type PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices struct {
	BasePrice    float32  `json:"base_price,omitempty"`
	CurrentPrice float32  `json:"current_price,omitempty"`
	FriendlyName string   `json:"friendly_name,omitempty"`
	Prefixes     []string `json:"prefixes,omitempty"`
}

func (response *PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		BasePrice    interface{} `json:"base_price"`
		CurrentPrice interface{} `json:"current_price"`
		FriendlyName string      `json:"friendly_name"`
		Prefixes     []string    `json:"prefixes"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices{
		FriendlyName: raw.FriendlyName,
		Prefixes:     raw.Prefixes,
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
