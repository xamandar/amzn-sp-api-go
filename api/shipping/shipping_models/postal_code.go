// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PostalCode The postal code of that address. It contains a series of letters or digits or both, sometimes including spaces or punctuation.
//
// swagger:model PostalCode
type PostalCode string

// Validate validates this postal code
func (m PostalCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 20); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this postal code based on context it is used
func (m PostalCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}