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

// SupersimV1UsageRecord struct for SupersimV1UsageRecord
type SupersimV1UsageRecord struct {
	// The SID of the Account that incurred the usage.
	AccountSid *string `json:"account_sid,omitempty"`
	// Total data downloaded in bytes, aggregated by the query parameters.
	DataDownload *int `json:"data_download,omitempty"`
	// Total of data_upload and data_download.
	DataTotal *int `json:"data_total,omitempty"`
	// Total data uploaded in bytes, aggregated by the query parameters.
	DataUpload *int `json:"data_upload,omitempty"`
	// SID of the Fleet resource on which the usage occurred.
	FleetSid *string `json:"fleet_sid,omitempty"`
	// Alpha-2 ISO Country Code of the country the usage occurred in.
	IsoCountry *string `json:"iso_country,omitempty"`
	// SID of the Network resource on which the usage occurred.
	NetworkSid *string `json:"network_sid,omitempty"`
	// The time period for which the usage is reported.
	Period *interface{} `json:"period,omitempty"`
	// SID of a Sim resource to which the UsageRecord belongs.
	SimSid *string `json:"sim_sid,omitempty"`
}
