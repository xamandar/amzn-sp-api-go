// Code generated by go-swagger; DO NOT EDIT.

package easy_ship_2022_03_23_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrackingDetails Representation of tracking metadata.
//
// swagger:model TrackingDetails
type TrackingDetails struct {

	// The tracking identifier for the scheduled package.
	TrackingID String `json:"trackingId,omitempty"`
}

// Validate validates this tracking details
func (m *TrackingDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrackingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackingDetails) validateTrackingID(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackingID) { // not required
		return nil
	}

	if err := m.TrackingID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trackingId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("trackingId")
		}
		return err
	}

	return nil
}

// ContextValidate validate this tracking details based on the context it is used
func (m *TrackingDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrackingID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackingDetails) contextValidateTrackingID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TrackingID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trackingId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("trackingId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrackingDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrackingDetails) UnmarshalBinary(b []byte) error {
	var res TrackingDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}