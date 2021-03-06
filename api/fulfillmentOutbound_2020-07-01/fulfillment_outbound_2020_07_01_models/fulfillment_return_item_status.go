// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FulfillmentReturnItemStatus Indicates if the return item has been processed by a fulfillment center.
//
// swagger:model FulfillmentReturnItemStatus
type FulfillmentReturnItemStatus string

func NewFulfillmentReturnItemStatus(value FulfillmentReturnItemStatus) *FulfillmentReturnItemStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FulfillmentReturnItemStatus.
func (m FulfillmentReturnItemStatus) Pointer() *FulfillmentReturnItemStatus {
	return &m
}

const (

	// FulfillmentReturnItemStatusNew captures enum value "New"
	FulfillmentReturnItemStatusNew FulfillmentReturnItemStatus = "New"

	// FulfillmentReturnItemStatusProcessed captures enum value "Processed"
	FulfillmentReturnItemStatusProcessed FulfillmentReturnItemStatus = "Processed"
)

// for schema
var fulfillmentReturnItemStatusEnum []interface{}

func init() {
	var res []FulfillmentReturnItemStatus
	if err := json.Unmarshal([]byte(`["New","Processed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fulfillmentReturnItemStatusEnum = append(fulfillmentReturnItemStatusEnum, v)
	}
}

func (m FulfillmentReturnItemStatus) validateFulfillmentReturnItemStatusEnum(path, location string, value FulfillmentReturnItemStatus) error {
	if err := validate.EnumCase(path, location, value, fulfillmentReturnItemStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this fulfillment return item status
func (m FulfillmentReturnItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFulfillmentReturnItemStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this fulfillment return item status based on context it is used
func (m FulfillmentReturnItemStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
