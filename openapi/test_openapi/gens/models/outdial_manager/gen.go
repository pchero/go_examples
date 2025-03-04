// Package outdial_manager provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package outdial_manager

import (
	"time"

	externalRef0 "testoapi/gens/models/common"
)

// Defines values for OutdialtargetStatus.
const (
	StatusDone        OutdialtargetStatus = "done"
	StatusIdle        OutdialtargetStatus = "idle"
	StatusProgressing OutdialtargetStatus = "progressing"
)

// Outdial defines model for Outdial.
type Outdial struct {
	// CampaignId The unique identifier for the campaign associated with the outdial.
	CampaignId *string `json:"campaign_id,omitempty"`

	// CustomerId The unique identifier for the customer associated with the outdial.
	CustomerId *string `json:"customer_id,omitempty"`

	// Data The data associated with the outdial.
	Data *string `json:"data,omitempty"`

	// Detail The detailed description of the outdial.
	Detail *string `json:"detail,omitempty"`

	// Id The unique identifier for the outdial.
	Id *string `json:"id,omitempty"`

	// Name The name of the outdial.
	Name *string `json:"name,omitempty"`

	// TmCreate Timestamp when the outdial was created.
	TmCreate *string `json:"tm_create,omitempty"`

	// TmDelete Timestamp when the outdial was deleted.
	TmDelete *string `json:"tm_delete,omitempty"`

	// TmUpdate Timestamp when the outdial was last updated.
	TmUpdate *string `json:"tm_update,omitempty"`
}

// Outdialtarget defines model for Outdialtarget.
type Outdialtarget struct {
	// Data The data associated with the outdial.
	Data *string `json:"data,omitempty"`

	// Destination0 Contains source or destination detail info.
	Destination0 *externalRef0.Address `json:"destination_0,omitempty"`

	// Destination1 Contains source or destination detail info.
	Destination1 *externalRef0.Address `json:"destination_1,omitempty"`

	// Destination2 Contains source or destination detail info.
	Destination2 *externalRef0.Address `json:"destination_2,omitempty"`

	// Destination3 Contains source or destination detail info.
	Destination3 *externalRef0.Address `json:"destination_3,omitempty"`

	// Destination4 Contains source or destination detail info.
	Destination4 *externalRef0.Address `json:"destination_4,omitempty"`

	// Detail Additional details about the outdial.
	Detail *string `json:"detail,omitempty"`

	// Id The unique identifier for the outdial.
	Id *string `json:"id,omitempty"`

	// Name The name of the outdial.
	Name *string `json:"name,omitempty"`

	// OutdialId The outdial reference ID.
	OutdialId *string `json:"outdial_id,omitempty"`

	// Status The status of the outdial.
	Status *OutdialtargetStatus `json:"status,omitempty"`

	// TmCreate The creation timestamp.
	TmCreate *time.Time `json:"tm_create,omitempty"`

	// TmDelete The deletion timestamp.
	TmDelete *time.Time `json:"tm_delete,omitempty"`

	// TmUpdate The update timestamp.
	TmUpdate *time.Time `json:"tm_update,omitempty"`

	// TryCount0 The try count for destination 0.
	TryCount0 *int `json:"try_count_0,omitempty"`

	// TryCount1 The try count for destination 1.
	TryCount1 *int `json:"try_count_1,omitempty"`

	// TryCount2 The try count for destination 2.
	TryCount2 *int `json:"try_count_2,omitempty"`

	// TryCount3 The try count for destination 3.
	TryCount3 *int `json:"try_count_3,omitempty"`

	// TryCount4 The try count for destination 4.
	TryCount4 *int `json:"try_count_4,omitempty"`
}

// OutdialtargetStatus The status of the outdial.
type OutdialtargetStatus string
