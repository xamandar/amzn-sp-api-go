// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrepDetails Preparation instructions and who is responsible for the preparation.
//
// swagger:model PrepDetails
type PrepDetails struct {

	// prep instruction
	// Required: true
	PrepInstruction *PrepInstruction `json:"PrepInstruction"`

	// prep owner
	// Required: true
	PrepOwner *PrepOwner `json:"PrepOwner"`
}

// Validate validates this prep details
func (m *PrepDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrepInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrepDetails) validatePrepInstruction(formats strfmt.Registry) error {

	if err := validate.Required("PrepInstruction", "body", m.PrepInstruction); err != nil {
		return err
	}

	if err := validate.Required("PrepInstruction", "body", m.PrepInstruction); err != nil {
		return err
	}

	if m.PrepInstruction != nil {
		if err := m.PrepInstruction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrepInstruction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrepInstruction")
			}
			return err
		}
	}

	return nil
}

func (m *PrepDetails) validatePrepOwner(formats strfmt.Registry) error {

	if err := validate.Required("PrepOwner", "body", m.PrepOwner); err != nil {
		return err
	}

	if err := validate.Required("PrepOwner", "body", m.PrepOwner); err != nil {
		return err
	}

	if m.PrepOwner != nil {
		if err := m.PrepOwner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrepOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrepOwner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this prep details based on the context it is used
func (m *PrepDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrepInstruction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrepOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrepDetails) contextValidatePrepInstruction(ctx context.Context, formats strfmt.Registry) error {

	if m.PrepInstruction != nil {
		if err := m.PrepInstruction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrepInstruction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrepInstruction")
			}
			return err
		}
	}

	return nil
}

func (m *PrepDetails) contextValidatePrepOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.PrepOwner != nil {
		if err := m.PrepOwner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrepOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrepOwner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrepDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrepDetails) UnmarshalBinary(b []byte) error {
	var res PrepDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}