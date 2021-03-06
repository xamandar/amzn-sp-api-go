// Code generated by go-swagger; DO NOT EDIT.

package fba_smalland_light_models

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

// FeeLineItem Fee details for a specific fee.
//
// swagger:model FeeLineItem
type FeeLineItem struct {

	// Amount charged to the seller for the specific fee type.
	// Required: true
	FeeCharge *MoneyType `json:"feeCharge"`

	// The type of fee charged to the seller.
	// Required: true
	// Enum: [FBAWeightBasedFee FBAPerOrderFulfillmentFee FBAPerUnitFulfillmentFee Commission]
	FeeType *string `json:"feeType"`
}

// Validate validates this fee line item
func (m *FeeLineItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeeCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeeLineItem) validateFeeCharge(formats strfmt.Registry) error {

	if err := validate.Required("feeCharge", "body", m.FeeCharge); err != nil {
		return err
	}

	if m.FeeCharge != nil {
		if err := m.FeeCharge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feeCharge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("feeCharge")
			}
			return err
		}
	}

	return nil
}

var feeLineItemTypeFeeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FBAWeightBasedFee","FBAPerOrderFulfillmentFee","FBAPerUnitFulfillmentFee","Commission"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feeLineItemTypeFeeTypePropEnum = append(feeLineItemTypeFeeTypePropEnum, v)
	}
}

const (

	// FeeLineItemFeeTypeFBAWeightBasedFee captures enum value "FBAWeightBasedFee"
	FeeLineItemFeeTypeFBAWeightBasedFee string = "FBAWeightBasedFee"

	// FeeLineItemFeeTypeFBAPerOrderFulfillmentFee captures enum value "FBAPerOrderFulfillmentFee"
	FeeLineItemFeeTypeFBAPerOrderFulfillmentFee string = "FBAPerOrderFulfillmentFee"

	// FeeLineItemFeeTypeFBAPerUnitFulfillmentFee captures enum value "FBAPerUnitFulfillmentFee"
	FeeLineItemFeeTypeFBAPerUnitFulfillmentFee string = "FBAPerUnitFulfillmentFee"

	// FeeLineItemFeeTypeCommission captures enum value "Commission"
	FeeLineItemFeeTypeCommission string = "Commission"
)

// prop value enum
func (m *FeeLineItem) validateFeeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, feeLineItemTypeFeeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeeLineItem) validateFeeType(formats strfmt.Registry) error {

	if err := validate.Required("feeType", "body", m.FeeType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFeeTypeEnum("feeType", "body", *m.FeeType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fee line item based on the context it is used
func (m *FeeLineItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeeCharge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeeLineItem) contextValidateFeeCharge(ctx context.Context, formats strfmt.Registry) error {

	if m.FeeCharge != nil {
		if err := m.FeeCharge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feeCharge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("feeCharge")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeeLineItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeeLineItem) UnmarshalBinary(b []byte) error {
	var res FeeLineItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
