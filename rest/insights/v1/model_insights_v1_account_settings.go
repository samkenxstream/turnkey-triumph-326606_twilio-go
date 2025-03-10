/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// InsightsV1AccountSettings struct for InsightsV1AccountSettings
type InsightsV1AccountSettings struct {
	AccountSid       *string `json:"account_sid,omitempty"`
	AdvancedFeatures *bool   `json:"advanced_features,omitempty"`
	Url              *string `json:"url,omitempty"`
	VoiceTrace       *bool   `json:"voice_trace,omitempty"`
}
