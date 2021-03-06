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

// TrackingID The tracking id generated to each shipment. It contains a series of letters or digits or both.
//
// swagger:model TrackingId
type TrackingID string

// Validate validates this tracking Id
func (m TrackingID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 60); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tracking Id based on context it is used
func (m TrackingID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
