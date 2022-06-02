// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Fee Fee type and cost.
//
// swagger:model Fee
type Fee struct {

	// The amount of the fee.
	// Required: true
	Amount *Money `json:"amount"`

	// The type of fee.
	// Required: true
	// Enum: [FBAPerUnitFulfillmentFee FBAPerOrderFulfillmentFee FBATransportationFee FBAFulfillmentCODFee]
	Name *string `json:"name"`
}

// Validate validates this fee
func (m *Fee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fee) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

var feeTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FBAPerUnitFulfillmentFee","FBAPerOrderFulfillmentFee","FBATransportationFee","FBAFulfillmentCODFee"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feeTypeNamePropEnum = append(feeTypeNamePropEnum, v)
	}
}

const (

	// FeeNameFBAPerUnitFulfillmentFee captures enum value "FBAPerUnitFulfillmentFee"
	FeeNameFBAPerUnitFulfillmentFee string = "FBAPerUnitFulfillmentFee"

	// FeeNameFBAPerOrderFulfillmentFee captures enum value "FBAPerOrderFulfillmentFee"
	FeeNameFBAPerOrderFulfillmentFee string = "FBAPerOrderFulfillmentFee"

	// FeeNameFBATransportationFee captures enum value "FBATransportationFee"
	FeeNameFBATransportationFee string = "FBATransportationFee"

	// FeeNameFBAFulfillmentCODFee captures enum value "FBAFulfillmentCODFee"
	FeeNameFBAFulfillmentCODFee string = "FBAFulfillmentCODFee"
)

// prop value enum
func (m *Fee) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, feeTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Fee) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fee based on the context it is used
func (m *Fee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {
		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Fee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Fee) UnmarshalBinary(b []byte) error {
	var res Fee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}