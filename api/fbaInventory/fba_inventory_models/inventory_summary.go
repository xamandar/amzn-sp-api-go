// Code generated by go-swagger; DO NOT EDIT.

package fba_inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InventorySummary Inventory summary for a specific item.
//
// swagger:model InventorySummary
type InventorySummary struct {

	// The Amazon Standard Identification Number (ASIN) of an item.
	Asin string `json:"asin,omitempty"`

	// The condition of the item as described by the seller (for example, New Item).
	Condition string `json:"condition,omitempty"`

	// Amazon's fulfillment network SKU identifier.
	FnSku string `json:"fnSku,omitempty"`

	// inventory details
	InventoryDetails *InventoryDetails `json:"inventoryDetails,omitempty"`

	// The date and time that any quantity was last updated.
	// Format: date-time
	LastUpdatedTime strfmt.DateTime `json:"lastUpdatedTime,omitempty"`

	// The localized language product title of the item within the specific marketplace.
	ProductName string `json:"productName,omitempty"`

	// The seller SKU of the item.
	SellerSku string `json:"sellerSku,omitempty"`

	// The total number of units in an inbound shipment or in Amazon fulfillment centers.
	TotalQuantity int64 `json:"totalQuantity,omitempty"`
}

// Validate validates this inventory summary
func (m *InventorySummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInventoryDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventorySummary) validateInventoryDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.InventoryDetails) { // not required
		return nil
	}

	if m.InventoryDetails != nil {
		if err := m.InventoryDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventoryDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventoryDetails")
			}
			return err
		}
	}

	return nil
}

func (m *InventorySummary) validateLastUpdatedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedTime", "body", "date-time", m.LastUpdatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inventory summary based on the context it is used
func (m *InventorySummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInventoryDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventorySummary) contextValidateInventoryDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.InventoryDetails != nil {
		if err := m.InventoryDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventoryDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventoryDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventorySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventorySummary) UnmarshalBinary(b []byte) error {
	var res InventorySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}