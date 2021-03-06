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

// StandardTextBlock The A+ Content standard text box block, comprised of a paragraph with a headline.
//
// swagger:model StandardTextBlock
type StandardTextBlock struct {

	// body
	Body *ParagraphComponent `json:"body,omitempty"`

	// headline
	Headline *TextComponent `json:"headline,omitempty"`
}

// Validate validates this standard text block
func (m *StandardTextBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
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

func (m *StandardTextBlock) validateBody(formats strfmt.Registry) error {
	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *StandardTextBlock) validateHeadline(formats strfmt.Registry) error {
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

// ContextValidate validate this standard text block based on the context it is used
func (m *StandardTextBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBody(ctx, formats); err != nil {
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

func (m *StandardTextBlock) contextValidateBody(ctx context.Context, formats strfmt.Registry) error {

	if m.Body != nil {
		if err := m.Body.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *StandardTextBlock) contextValidateHeadline(ctx context.Context, formats strfmt.Registry) error {

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
func (m *StandardTextBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandardTextBlock) UnmarshalBinary(b []byte) error {
	var res StandardTextBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
