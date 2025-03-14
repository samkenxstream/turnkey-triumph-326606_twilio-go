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

import (
	"time"
)

// InsightsV1Conference struct for InsightsV1Conference
type InsightsV1Conference struct {
	// Account SID.
	AccountSid *string `json:"account_sid,omitempty"`
	// Conference SID.
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// Duration of the conference in seconds.
	ConnectDurationSeconds *int `json:"connect_duration_seconds,omitempty"`
	// Conference creation date/time.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Potential issues detected during the conference.
	DetectedIssues *interface{} `json:"detected_issues,omitempty"`
	// Conference duration in seconds.
	DurationSeconds *int `json:"duration_seconds,omitempty"`
	// Conference end reason.
	EndReason *string `json:"end_reason,omitempty"`
	// Conference end date/time.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Call SID that ended the conference.
	EndedBy *string `json:"ended_by,omitempty"`
	// Custom label for the conference.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// Actual maximum concurrent participants.
	MaxConcurrentParticipants *int `json:"max_concurrent_participants,omitempty"`
	// Max participants specified in config.
	MaxParticipants *int `json:"max_participants,omitempty"`
	// Region where the conference was mixed.
	MixerRegion *string `json:"mixer_region,omitempty"`
	// Configuration-requested conference mixer region.
	MixerRegionRequested *string `json:"mixer_region_requested,omitempty"`
	// Processing state for the Conference Summary resource.
	ProcessingState *string `json:"processing_state,omitempty"`
	// Boolean. Indicates whether recording was enabled.
	RecordingEnabled *bool `json:"recording_enabled,omitempty"`
	// Timestamp in ISO 8601 format when the conference started.
	StartTime *time.Time `json:"start_time,omitempty"`
	// Status of conference
	Status *string `json:"status,omitempty"`
	// Object. Contains details about conference tags.
	TagInfo *interface{} `json:"tag_info,omitempty"`
	// Tags for detected conference conditions and participant behaviors.
	Tags *[]string `json:"tags,omitempty"`
	// Unique conference participants.
	UniqueParticipants *int `json:"unique_participants,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
