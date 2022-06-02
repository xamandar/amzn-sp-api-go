// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemLevelFields item level fields
//
// swagger:model ItemLevelFields
type ItemLevelFields struct {

	// additional inputs
	// Required: true
	AdditionalInputs AdditionalInputsList `json:"AdditionalInputs"`

	// The Amazon Standard Identification Number (ASIN) of the item.
	// Required: true
	Asin *string `json:"Asin"`
}

// Validate validates this item level fields
func (m *ItemLevelFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemLevelFields) validateAdditionalInputs(formats strfmt.Registry) error {

	if err := validate.Required("AdditionalInputs", "body", m.AdditionalInputs); err != nil {
		return err
	}

	if err := m.AdditionalInputs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AdditionalInputs")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AdditionalInputs")
		}
		return err
	}

	return nil
}

func (m *ItemLevelFields) validateAsin(formats strfmt.Registry) error {

	if err := validate.Required("Asin", "body", m.Asin); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this item level fields based on the context it is used
func (m *ItemLevelFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemLevelFields) contextValidateAdditionalInputs(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AdditionalInputs.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AdditionalInputs")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AdditionalInputs")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemLevelFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemLevelFields) UnmarshalBinary(b []byte) error {
	var res ItemLevelFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}