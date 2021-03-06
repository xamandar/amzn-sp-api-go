// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AcceptedRate The specific rate purchased for the shipment, or null if unpurchased.
//
// swagger:model AcceptedRate
type AcceptedRate struct {

	// The weight that was used to calculate the totalCharge.
	BilledWeight *Weight `json:"billedWeight,omitempty"`

	// promise
	Promise *ShippingPromiseSet `json:"promise,omitempty"`

	// service type
	ServiceType ServiceType `json:"serviceType,omitempty"`

	// The total charge that will be billed for the rate.
	TotalCharge *Currency `json:"totalCharge,omitempty"`
}

// Validate validates this accepted rate
func (m *AcceptedRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBilledWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCharge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcceptedRate) validateBilledWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.BilledWeight) { // not required
		return nil
	}

	if m.BilledWeight != nil {
		if err := m.BilledWeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billedWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billedWeight")
			}
			return err
		}
	}

	return nil
}

func (m *AcceptedRate) validatePromise(formats strfmt.Registry) error {
	if swag.IsZero(m.Promise) { // not required
		return nil
	}

	if m.Promise != nil {
		if err := m.Promise.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("promise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("promise")
			}
			return err
		}
	}

	return nil
}

func (m *AcceptedRate) validateServiceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceType) { // not required
		return nil
	}

	if err := m.ServiceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceType")
		}
		return err
	}

	return nil
}

func (m *AcceptedRate) validateTotalCharge(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalCharge) { // not required
		return nil
	}

	if m.TotalCharge != nil {
		if err := m.TotalCharge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totalCharge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("totalCharge")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this accepted rate based on the context it is used
func (m *AcceptedRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBilledWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePromise(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalCharge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcceptedRate) contextValidateBilledWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.BilledWeight != nil {
		if err := m.BilledWeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billedWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billedWeight")
			}
			return err
		}
	}

	return nil
}

func (m *AcceptedRate) contextValidatePromise(ctx context.Context, formats strfmt.Registry) error {

	if m.Promise != nil {
		if err := m.Promise.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("promise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("promise")
			}
			return err
		}
	}

	return nil
}

func (m *AcceptedRate) contextValidateServiceType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServiceType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceType")
		}
		return err
	}

	return nil
}

func (m *AcceptedRate) contextValidateTotalCharge(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalCharge != nil {
		if err := m.TotalCharge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totalCharge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("totalCharge")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcceptedRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcceptedRate) UnmarshalBinary(b []byte) error {
	var res AcceptedRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
