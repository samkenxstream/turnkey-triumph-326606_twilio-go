/*
 * Twilio - Media
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListPlayerStreamerResponse struct for ListPlayerStreamerResponse
type ListPlayerStreamerResponse struct {
	Meta            ListMediaProcessorResponseMeta `json:"meta,omitempty"`
	PlayerStreamers []MediaV1PlayerStreamer        `json:"player_streamers,omitempty"`
}
