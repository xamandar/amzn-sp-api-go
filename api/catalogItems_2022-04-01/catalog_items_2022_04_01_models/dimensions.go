// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2022_04_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Dimensions Dimensions of an Amazon catalog item or item in its packaging.
//
// swagger:model Dimensions
type Dimensions struct {

	// Height of an item or item package.
	Height *Dimension `json:"height,omitempty"`

	// Length of an item or item package.
	Length *Dimension `json:"length,omitempty"`

	// Weight of an item or item package.
	Weight *Dimension `json:"weight,omitempty"`

	// Width of an item or item package.
	Width *Dimension `json:"width,omitempty"`
}

// Validate validates this dimensions
func (m *Dimensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dimensions) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Height) { // not required
		return nil
	}

	if m.Height != nil {
		if err := m.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) validateLength(formats strfmt.Registry) error {
	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if m.Length != nil {
		if err := m.Length.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("length")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("length")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if m.Weight != nil {
		if err := m.Weight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) validateWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Width) { // not required
		return nil
	}

	if m.Width != nil {
		if err := m.Width.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dimensions based on the context it is used
func (m *Dimensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLength(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dimensions) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Height != nil {
		if err := m.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) contextValidateLength(ctx context.Context, formats strfmt.Registry) error {

	if m.Length != nil {
		if err := m.Length.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("length")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("length")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Weight != nil {
		if err := m.Weight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if m.Width != nil {
		if err := m.Width.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Dimensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dimensions) UnmarshalBinary(b []byte) error {
	var res Dimensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
