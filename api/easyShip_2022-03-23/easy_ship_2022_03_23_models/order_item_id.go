// Code generated by go-swagger; DO NOT EDIT.

package easy_ship_2022_03_23_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OrderItemID The Amazon-defined order item identifier.
//
// swagger:model OrderItemId
type OrderItemID string

// Validate validates this order item Id
func (m OrderItemID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MaxLength("", "body", string(m), 255); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this order item Id based on context it is used
func (m OrderItemID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
