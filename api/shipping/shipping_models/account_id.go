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

// AccountID This is the Amazon Shipping account id generated during the Amazon Shipping onboarding process.
//
// swagger:model AccountId
type AccountID string

// Validate validates this account Id
func (m AccountID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MaxLength("", "body", string(m), 10); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this account Id based on context it is used
func (m AccountID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
