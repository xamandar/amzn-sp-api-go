// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AplusResponse The base response data for all A+ Content operations when a request is successful or partially successful. Individual operations may extend this with additional data.
//
// swagger:model AplusResponse
type AplusResponse struct {

	// warnings
	Warnings MessageSet `json:"warnings,omitempty"`
}

// Validate validates this aplus response
func (m *AplusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AplusResponse) validateWarnings(formats strfmt.Registry) error {
	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	if err := m.Warnings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("warnings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("warnings")
		}
		return err
	}

	return nil
}

// ContextValidate validate this aplus response based on the context it is used
func (m *AplusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWarnings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AplusResponse) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Warnings.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("warnings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("warnings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AplusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AplusResponse) UnmarshalBinary(b []byte) error {
	var res AplusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
