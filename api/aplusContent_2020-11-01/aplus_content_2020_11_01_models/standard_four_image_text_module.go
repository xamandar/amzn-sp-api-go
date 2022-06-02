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

// StandardFourImageTextModule Four standard images with text, presented across a single row.
//
// swagger:model StandardFourImageTextModule
type StandardFourImageTextModule struct {

	// block1
	Block1 *StandardImageTextBlock `json:"block1,omitempty"`

	// block2
	Block2 *StandardImageTextBlock `json:"block2,omitempty"`

	// block3
	Block3 *StandardImageTextBlock `json:"block3,omitempty"`

	// block4
	Block4 *StandardImageTextBlock `json:"block4,omitempty"`

	// headline
	Headline *TextComponent `json:"headline,omitempty"`
}

// Validate validates this standard four image text module
func (m *StandardFourImageTextModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlock1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlock2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlock3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlock4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeadline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardFourImageTextModule) validateBlock1(formats strfmt.Registry) error {
	if swag.IsZero(m.Block1) { // not required
		return nil
	}

	if m.Block1 != nil {
		if err := m.Block1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block1")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) validateBlock2(formats strfmt.Registry) error {
	if swag.IsZero(m.Block2) { // not required
		return nil
	}

	if m.Block2 != nil {
		if err := m.Block2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block2")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) validateBlock3(formats strfmt.Registry) error {
	if swag.IsZero(m.Block3) { // not required
		return nil
	}

	if m.Block3 != nil {
		if err := m.Block3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block3")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) validateBlock4(formats strfmt.Registry) error {
	if swag.IsZero(m.Block4) { // not required
		return nil
	}

	if m.Block4 != nil {
		if err := m.Block4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block4")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) validateHeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.Headline) { // not required
		return nil
	}

	if m.Headline != nil {
		if err := m.Headline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("headline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("headline")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this standard four image text module based on the context it is used
func (m *StandardFourImageTextModule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlock1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlock2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlock3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlock4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeadline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardFourImageTextModule) contextValidateBlock1(ctx context.Context, formats strfmt.Registry) error {

	if m.Block1 != nil {
		if err := m.Block1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block1")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) contextValidateBlock2(ctx context.Context, formats strfmt.Registry) error {

	if m.Block2 != nil {
		if err := m.Block2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block2")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) contextValidateBlock3(ctx context.Context, formats strfmt.Registry) error {

	if m.Block3 != nil {
		if err := m.Block3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block3")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) contextValidateBlock4(ctx context.Context, formats strfmt.Registry) error {

	if m.Block4 != nil {
		if err := m.Block4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block4")
			}
			return err
		}
	}

	return nil
}

func (m *StandardFourImageTextModule) contextValidateHeadline(ctx context.Context, formats strfmt.Registry) error {

	if m.Headline != nil {
		if err := m.Headline.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("headline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("headline")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandardFourImageTextModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandardFourImageTextModule) UnmarshalBinary(b []byte) error {
	var res StandardFourImageTextModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}