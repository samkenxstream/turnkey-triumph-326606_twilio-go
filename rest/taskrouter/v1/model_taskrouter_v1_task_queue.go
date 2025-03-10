/*
 * Twilio - Taskrouter
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

// TaskrouterV1TaskQueue struct for TaskrouterV1TaskQueue
type TaskrouterV1TaskQueue struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The name of the Activity to assign Workers when a task is assigned for them
	AssignmentActivityName *string `json:"assignment_activity_name,omitempty"`
	// The SID of the Activity to assign Workers when a task is assigned for them
	AssignmentActivitySid *string `json:"assignment_activity_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The maximum number of Workers to reserve
	MaxReservedWorkers *int `json:"max_reserved_workers,omitempty"`
	// The name of the Activity to assign Workers once a task is reserved for them
	ReservationActivityName *string `json:"reservation_activity_name,omitempty"`
	// The SID of the Activity to assign Workers once a task is reserved for them
	ReservationActivitySid *string `json:"reservation_activity_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue
	TargetWorkers *string `json:"target_workers,omitempty"`
	// How Tasks will be assigned to Workers
	TaskOrder *string `json:"task_order,omitempty"`
	// The absolute URL of the TaskQueue resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the TaskQueue
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
