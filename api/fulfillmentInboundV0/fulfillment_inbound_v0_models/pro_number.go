// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ProNumber The PRO number ("progressive number" or "progressive ID") assigned to the shipment by the carrier.
//
// swagger:model ProNumber
type ProNumber string

// Validate validates this pro number
func (m ProNumber) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pro number based on context it is used
func (m ProNumber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
