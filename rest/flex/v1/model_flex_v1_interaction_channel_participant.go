/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// FlexV1InteractionChannelParticipant struct for FlexV1InteractionChannelParticipant
type FlexV1InteractionChannelParticipant struct {
	// The Channel Sid for this Participant.
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The Interaction Sid for this channel.
	InteractionSid *string `json:"interaction_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Participant type.
	Type *string `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
}
