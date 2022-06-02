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

// SellerInputDefinition Specifies characteristics that apply to a seller input.
//
// swagger:model SellerInputDefinition
type SellerInputDefinition struct {

	// constraints
	// Required: true
	Constraints Constraints `json:"Constraints"`

	// The data type of the additional input field.
	// Required: true
	DataType *string `json:"DataType"`

	// The display text for the additional input field.
	// Required: true
	InputDisplayText *string `json:"InputDisplayText"`

	// Whether the seller input applies to the item or the shipment.
	InputTarget InputTargetType `json:"InputTarget,omitempty"`

	// When true, the additional input field is required.
	// Required: true
	IsRequired *bool `json:"IsRequired"`

	// restricted set values
	RestrictedSetValues RestrictedSetValues `json:"RestrictedSetValues,omitempty"`

	// stored value
	// Required: true
	StoredValue *AdditionalSellerInput `json:"StoredValue"`
}

// Validate validates this seller input definition
func (m *SellerInputDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputDisplayText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictedSetValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoredValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SellerInputDefinition) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("Constraints", "body", m.Constraints); err != nil {
		return err
	}

	if err := m.Constraints.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Constraints")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Constraints")
		}
		return err
	}

	return nil
}

func (m *SellerInputDefinition) validateDataType(formats strfmt.Registry) error {

	if err := validate.Required("DataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *SellerInputDefinition) validateInputDisplayText(formats strfmt.Registry) error {

	if err := validate.Required("InputDisplayText", "body", m.InputDisplayText); err != nil {
		return err
	}

	return nil
}

func (m *SellerInputDefinition) validateInputTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.InputTarget) { // not required
		return nil
	}

	if err := m.InputTarget.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InputTarget")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InputTarget")
		}
		return err
	}

	return nil
}

func (m *SellerInputDefinition) validateIsRequired(formats strfmt.Registry) error {

	if err := validate.Required("IsRequired", "body", m.IsRequired); err != nil {
		return err
	}

	return nil
}

func (m *SellerInputDefinition) validateRestrictedSetValues(formats strfmt.Registry) error {
	if swag.IsZero(m.RestrictedSetValues) { // not required
		return nil
	}

	if err := m.RestrictedSetValues.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RestrictedSetValues")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("RestrictedSetValues")
		}
		return err
	}

	return nil
}

func (m *SellerInputDefinition) validateStoredValue(formats strfmt.Registry) error {

	if err := validate.Required("StoredValue", "body", m.StoredValue); err != nil {
		return err
	}

	if m.StoredValue != nil {
		if err := m.StoredValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StoredValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StoredValue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this seller input definition based on the context it is used
func (m *SellerInputDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestrictedSetValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoredValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SellerInputDefinition) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Constraints.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Constraints")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Constraints")
		}
		return err
	}

	return nil
}

func (m *SellerInputDefinition) contextValidateInputTarget(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InputTarget.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InputTarget")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InputTarget")
		}
		return err
	}

	return nil
}

func (m *SellerInputDefinition) contextValidateRestrictedSetValues(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RestrictedSetValues.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RestrictedSetValues")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("RestrictedSetValues")
		}
		return err
	}

	return nil
}

func (m *SellerInputDefinition) contextValidateStoredValue(ctx context.Context, formats strfmt.Registry) error {

	if m.StoredValue != nil {
		if err := m.StoredValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StoredValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StoredValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SellerInputDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SellerInputDefinition) UnmarshalBinary(b []byte) error {
	var res SellerInputDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}