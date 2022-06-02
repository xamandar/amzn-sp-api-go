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

// ShippingServiceOptions Extra services provided by a carrier.
//
// swagger:model ShippingServiceOptions
type ShippingServiceOptions struct {

	// When true, the carrier will pick up the package.
	//
	// Note: Scheduled carrier pickup is available only using Dynamex (US), DPD (UK), and Royal Mail (UK).
	// Required: true
	CarrierWillPickUp *bool `json:"CarrierWillPickUp"`

	// carrier will pick up option
	CarrierWillPickUpOption CarrierWillPickUpOption `json:"CarrierWillPickUpOption,omitempty"`

	// The declared value of the shipment. The carrier uses this value to determine the amount to use to insure the shipment. If DeclaredValue is greater than the carrier's minimum insurance amount, the seller is charged for the additional insurance as determined by the carrier. For information about optional insurance coverage, see the Seller Central Help [UK](https://sellercentral.amazon.co.uk/gp/help/200204080) [US](https://sellercentral.amazon.com/gp/help/200204080).
	DeclaredValue *CurrencyAmount `json:"DeclaredValue,omitempty"`

	// The delivery confirmation level.
	// Required: true
	DeliveryExperience *DeliveryExperienceType `json:"DeliveryExperience"`

	// The seller's preferred label format.
	LabelFormat LabelFormat `json:"LabelFormat,omitempty"`
}

// Validate validates this shipping service options
func (m *ShippingServiceOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierWillPickUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCarrierWillPickUpOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeclaredValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryExperience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingServiceOptions) validateCarrierWillPickUp(formats strfmt.Registry) error {

	if err := validate.Required("CarrierWillPickUp", "body", m.CarrierWillPickUp); err != nil {
		return err
	}

	return nil
}

func (m *ShippingServiceOptions) validateCarrierWillPickUpOption(formats strfmt.Registry) error {
	if swag.IsZero(m.CarrierWillPickUpOption) { // not required
		return nil
	}

	if err := m.CarrierWillPickUpOption.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CarrierWillPickUpOption")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CarrierWillPickUpOption")
		}
		return err
	}

	return nil
}

func (m *ShippingServiceOptions) validateDeclaredValue(formats strfmt.Registry) error {
	if swag.IsZero(m.DeclaredValue) { // not required
		return nil
	}

	if m.DeclaredValue != nil {
		if err := m.DeclaredValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeclaredValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeclaredValue")
			}
			return err
		}
	}

	return nil
}

func (m *ShippingServiceOptions) validateDeliveryExperience(formats strfmt.Registry) error {

	if err := validate.Required("DeliveryExperience", "body", m.DeliveryExperience); err != nil {
		return err
	}

	if err := validate.Required("DeliveryExperience", "body", m.DeliveryExperience); err != nil {
		return err
	}

	if m.DeliveryExperience != nil {
		if err := m.DeliveryExperience.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryExperience")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryExperience")
			}
			return err
		}
	}

	return nil
}

func (m *ShippingServiceOptions) validateLabelFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelFormat) { // not required
		return nil
	}

	if err := m.LabelFormat.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LabelFormat")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LabelFormat")
		}
		return err
	}

	return nil
}

// ContextValidate validate this shipping service options based on the context it is used
func (m *ShippingServiceOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCarrierWillPickUpOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeclaredValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryExperience(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingServiceOptions) contextValidateCarrierWillPickUpOption(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CarrierWillPickUpOption.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CarrierWillPickUpOption")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CarrierWillPickUpOption")
		}
		return err
	}

	return nil
}

func (m *ShippingServiceOptions) contextValidateDeclaredValue(ctx context.Context, formats strfmt.Registry) error {

	if m.DeclaredValue != nil {
		if err := m.DeclaredValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeclaredValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeclaredValue")
			}
			return err
		}
	}

	return nil
}

func (m *ShippingServiceOptions) contextValidateDeliveryExperience(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryExperience != nil {
		if err := m.DeliveryExperience.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryExperience")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryExperience")
			}
			return err
		}
	}

	return nil
}

func (m *ShippingServiceOptions) contextValidateLabelFormat(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LabelFormat.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LabelFormat")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LabelFormat")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShippingServiceOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShippingServiceOptions) UnmarshalBinary(b []byte) error {
	var res ShippingServiceOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}