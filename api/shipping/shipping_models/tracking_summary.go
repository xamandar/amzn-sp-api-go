// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrackingSummary The tracking summary.
//
// swagger:model TrackingSummary
type TrackingSummary struct {

	// The derived status based on the events in the eventHistory.
	// Max Length: 60
	// Min Length: 1
	Status string `json:"status,omitempty"`
}

// Validate validates this tracking summary
func (m *TrackingSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackingSummary) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MinLength("status", "body", m.Status, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("status", "body", m.Status, 60); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tracking summary based on context it is used
func (m *TrackingSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrackingSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrackingSummary) UnmarshalBinary(b []byte) error {
	var res TrackingSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
