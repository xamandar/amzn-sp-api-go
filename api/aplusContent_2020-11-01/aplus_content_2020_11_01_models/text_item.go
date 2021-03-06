// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TextItem Rich positional text, usually presented as a collection of bullet points.
//
// swagger:model TextItem
type TextItem struct {

	// The rank or index of this text item within the collection. Different items cannot occupy the same position within a single collection.
	// Required: true
	// Maximum: 100
	// Minimum: 1
	Position *int64 `json:"position"`

	// text
	// Required: true
	Text *TextComponent `json:"text"`
}

// Validate validates this text item
func (m *TextItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextItem) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	if err := validate.MinimumInt("position", "body", *m.Position, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", *m.Position, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *TextItem) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	if m.Text != nil {
		if err := m.Text.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("text")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("text")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this text item based on the context it is used
func (m *TextItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextItem) contextValidateText(ctx context.Context, formats strfmt.Registry) error {

	if m.Text != nil {
		if err := m.Text.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("text")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("text")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextItem) UnmarshalBinary(b []byte) error {
	var res TextItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
