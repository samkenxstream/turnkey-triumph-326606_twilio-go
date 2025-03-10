/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSmsCommandResponse struct for ListSmsCommandResponse
type ListSmsCommandResponse struct {
	Meta        ListEsimProfileResponseMeta `json:"meta,omitempty"`
	SmsCommands []SupersimV1SmsCommand      `json:"sms_commands,omitempty"`
}
