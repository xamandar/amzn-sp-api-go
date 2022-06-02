// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuyerInfo Buyer information
//
// swagger:model BuyerInfo
type BuyerInfo struct {

	// The county of the buyer.
	BuyerCounty string `json:"BuyerCounty,omitempty"`

	// The anonymized email address of the buyer.
	BuyerEmail string `json:"BuyerEmail,omitempty"`

	// The name of the buyer.
	BuyerName string `json:"BuyerName,omitempty"`

	// Tax information about the buyer.
	BuyerTaxInfo *BuyerTaxInfo `json:"BuyerTaxInfo,omitempty"`

	// The purchase order (PO) number entered by the buyer at checkout. Returned only for orders where the buyer entered a PO number at checkout.
	PurchaseOrderNumber string `json:"PurchaseOrderNumber,omitempty"`
}

// Validate validates this buyer info
func (m *BuyerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuyerTaxInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuyerInfo) validateBuyerTaxInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyerTaxInfo) { // not required
		return nil
	}

	if m.BuyerTaxInfo != nil {
		if err := m.BuyerTaxInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerTaxInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerTaxInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this buyer info based on the context it is used
func (m *BuyerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuyerTaxInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuyerInfo) contextValidateBuyerTaxInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyerTaxInfo != nil {
		if err := m.BuyerTaxInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerTaxInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerTaxInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuyerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuyerInfo) UnmarshalBinary(b []byte) error {
	var res BuyerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}