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

// ListConversationMessageResponse struct for ListConversationMessageResponse
type ListConversationMessageResponse struct {
	Messages []ConversationsV1ConversationMessage `json:"messages,omitempty"`
	Meta     ListConfigurationAddressResponseMeta `json:"meta,omitempty"`
}
