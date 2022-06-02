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

// AvailableDeliveryExperienceOption The available delivery confirmation options, and the fee charged, if any.
//
// swagger:model AvailableDeliveryExperienceOption
type AvailableDeliveryExperienceOption struct {

	// charge
	// Required: true
	Charge *CurrencyAmount `json:"Charge"`

	// delivery experience option
	// Required: true
	DeliveryExperienceOption *DeliveryExperienceOption `json:"DeliveryExperienceOption"`
}

// Validate validates this available delivery experience option
func (m *AvailableDeliveryExperienceOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryExperienceOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableDeliveryExperienceOption) validateCharge(formats strfmt.Registry) error {

	if err := validate.Required("Charge", "body", m.Charge); err != nil {
		return err
	}

	if m.Charge != nil {
		if err := m.Charge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Charge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Charge")
			}
			return err
		}
	}

	return nil
}

func (m *AvailableDeliveryExperienceOption) validateDeliveryExperienceOption(formats strfmt.Registry) error {

	if err := validate.Required("DeliveryExperienceOption", "body", m.DeliveryExperienceOption); err != nil {
		return err
	}

	if err := validate.Required("DeliveryExperienceOption", "body", m.DeliveryExperienceOption); err != nil {
		return err
	}

	if m.DeliveryExperienceOption != nil {
		if err := m.DeliveryExperienceOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryExperienceOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryExperienceOption")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this available delivery experience option based on the context it is used
func (m *AvailableDeliveryExperienceOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCharge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryExperienceOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableDeliveryExperienceOption) contextValidateCharge(ctx context.Context, formats strfmt.Registry) error {

	if m.Charge != nil {
		if err := m.Charge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Charge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Charge")
			}
			return err
		}
	}

	return nil
}

func (m *AvailableDeliveryExperienceOption) contextValidateDeliveryExperienceOption(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryExperienceOption != nil {
		if err := m.DeliveryExperienceOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryExperienceOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryExperienceOption")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailableDeliveryExperienceOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableDeliveryExperienceOption) UnmarshalBinary(b []byte) error {
	var res AvailableDeliveryExperienceOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}