/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRateLimitResponse struct for ListRateLimitResponse
type ListRateLimitResponse struct {
	Meta       ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
	RateLimits []VerifyV2RateLimit                 `json:"rate_limits,omitempty"`
}
