/*
 * Twilio - Conversations
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

// ConversationsV1ServiceUserConversation struct for ConversationsV1ServiceUserConversation
type ConversationsV1ServiceUserConversation struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	Attributes *string `json:"attributes,omitempty"`
	// The unique ID of the Conversation Service this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the Conversation for this User Conversation.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The current state of this User Conversation
	ConversationState *string `json:"conversation_state,omitempty"`
	// Creator of this conversation.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date that this conversation was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this conversation was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The human-readable name of this conversation.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The index of the last read Message .
	LastReadMessageIndex *int `json:"last_read_message_index,omitempty"`
	// Absolute URLs to access the participant and conversation of this user conversation.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The Notification Level of this User Conversation.
	NotificationLevel *string `json:"notification_level,omitempty"`
	// Participant Sid.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// Timer date values for this conversation.
	Timers *interface{} `json:"timers,omitempty"`
	// An application-defined string that uniquely identifies the Conversation resource.
	UniqueName *string `json:"unique_name,omitempty"`
	// The number of unread Messages in the Conversation.
	UnreadMessagesCount *int    `json:"unread_messages_count,omitempty"`
	Url                 *string `json:"url,omitempty"`
	// The unique ID for the User.
	UserSid *string `json:"user_sid,omitempty"`
}
