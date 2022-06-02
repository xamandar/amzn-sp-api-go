// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ItemDescription The description of the item.
//
// swagger:model ItemDescription
type ItemDescription string

// Validate validates this item description
func (m ItemDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this item description based on context it is used
func (m ItemDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}