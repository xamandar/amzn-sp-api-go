// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaxWithholdingPeriod Period which taxwithholding on seller's account is calculated.
//
// swagger:model TaxWithholdingPeriod
type TaxWithholdingPeriod struct {

	// End of the time range.
	// Format: date-time
	EndDate Date `json:"EndDate,omitempty"`

	// Start of the time range.
	// Format: date-time
	StartDate Date `json:"StartDate,omitempty"`
}

// Validate validates this tax withholding period
func (m *TaxWithholdingPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxWithholdingPeriod) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := m.EndDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EndDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EndDate")
		}
		return err
	}

	return nil
}

func (m *TaxWithholdingPeriod) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := m.StartDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StartDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StartDate")
		}
		return err
	}

	return nil
}

// ContextValidate validate this tax withholding period based on the context it is used
func (m *TaxWithholdingPeriod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxWithholdingPeriod) contextValidateEndDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EndDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EndDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EndDate")
		}
		return err
	}

	return nil
}

func (m *TaxWithholdingPeriod) contextValidateStartDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StartDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StartDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StartDate")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxWithholdingPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxWithholdingPeriod) UnmarshalBinary(b []byte) error {
	var res TaxWithholdingPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}