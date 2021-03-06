// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ItemQuantity The number of items.
//
// swagger:model ItemQuantity
type ItemQuantity int32

// Validate validates this item quantity
func (m ItemQuantity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this item quantity based on context it is used
func (m ItemQuantity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
