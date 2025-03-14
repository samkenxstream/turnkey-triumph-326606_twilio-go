/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSchemaVersionResponse struct for ListSchemaVersionResponse
type ListSchemaVersionResponse struct {
	Meta           ListSchemaVersionResponseMeta `json:"meta,omitempty"`
	SchemaVersions []EventsV1SchemaVersion       `json:"schema_versions,omitempty"`
}
