// Code generated by go-swagger; DO NOT EDIT.

package vendor_invoices_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Money An amount of money, including units in the form of currency.
//
// swagger:model Money
type Money struct {

	// amount
	Amount Decimal `json:"amount,omitempty"`

	// Three-digit currency code in ISO 4217 format.
	CurrencyCode string `json:"currencyCode,omitempty"`
}

// Validate validates this money
func (m *Money) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Money) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("amount")
		}
		return err
	}

	return nil
}

// ContextValidate validate this money based on the context it is used
func (m *Money) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Money) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Amount.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("amount")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Money) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Money) UnmarshalBinary(b []byte) error {
	var res Money
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}