// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Item Details of the item being shipped.
//
// swagger:model Item
type Item struct {

	// Amazon Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier string `json:"amazonProductIdentifier,omitempty"`

	// item details
	ItemDetails *ItemDetails `json:"itemDetails,omitempty"`

	// Item sequence number for the item. The first item will be 001, the second 002, and so on. This number is used as a reference to refer to this item from the carton or pallet level.
	// Required: true
	ItemSequenceNumber *string `json:"itemSequenceNumber"`

	// Total item quantity shipped in this shipment.
	// Required: true
	ShippedQuantity *ItemQuantity `json:"shippedQuantity"`

	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSequenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateItemDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemDetails) { // not required
		return nil
	}

	if m.ItemDetails != nil {
		if err := m.ItemDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("itemDetails")
			}
			return err
		}
	}

	return nil
}

func (m *Item) validateItemSequenceNumber(formats strfmt.Registry) error {

	if err := validate.Required("itemSequenceNumber", "body", m.ItemSequenceNumber); err != nil {
		return err
	}

	return nil
}

func (m *Item) validateShippedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("shippedQuantity", "body", m.ShippedQuantity); err != nil {
		return err
	}

	if m.ShippedQuantity != nil {
		if err := m.ShippedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shippedQuantity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this item based on the context it is used
func (m *Item) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItemDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) contextValidateItemDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemDetails != nil {
		if err := m.ItemDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("itemDetails")
			}
			return err
		}
	}

	return nil
}

func (m *Item) contextValidateShippedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippedQuantity != nil {
		if err := m.ShippedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shippedQuantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}