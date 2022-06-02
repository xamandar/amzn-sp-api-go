// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ShipmentConfirmation shipment confirmation
//
// swagger:model ShipmentConfirmation
type ShipmentConfirmation struct {

	// Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package.
	Containers []*Container `json:"containers"`

	// Provide the details of the items in this shipment. If any of the item details field is common at a package or a pallet level, then provide them at the corresponding package.
	// Required: true
	Items []*Item `json:"items"`

	// Purchase order number corresponding to the shipment.
	// Required: true
	// Pattern: ^[a-zA-Z0-9]+$
	PurchaseOrderNumber *string `json:"purchaseOrderNumber"`

	// ID of the selling party or vendor.
	// Required: true
	SellingParty *PartyIdentification `json:"sellingParty"`

	// Warehouse code of vendor.
	// Required: true
	ShipFromParty *PartyIdentification `json:"shipFromParty"`

	// Shipment information.
	// Required: true
	ShipmentDetails *ShipmentDetails `json:"shipmentDetails"`
}

// Validate validates this shipment confirmation
func (m *ShipmentConfirmation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellingParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFromParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentConfirmation) validateContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	for i := 0; i < len(m.Containers); i++ {
		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {
			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ShipmentConfirmation) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ShipmentConfirmation) validatePurchaseOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderNumber", "body", m.PurchaseOrderNumber); err != nil {
		return err
	}

	if err := validate.Pattern("purchaseOrderNumber", "body", *m.PurchaseOrderNumber, `^[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ShipmentConfirmation) validateSellingParty(formats strfmt.Registry) error {

	if err := validate.Required("sellingParty", "body", m.SellingParty); err != nil {
		return err
	}

	if m.SellingParty != nil {
		if err := m.SellingParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentConfirmation) validateShipFromParty(formats strfmt.Registry) error {

	if err := validate.Required("shipFromParty", "body", m.ShipFromParty); err != nil {
		return err
	}

	if m.ShipFromParty != nil {
		if err := m.ShipFromParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFromParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFromParty")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentConfirmation) validateShipmentDetails(formats strfmt.Registry) error {

	if err := validate.Required("shipmentDetails", "body", m.ShipmentDetails); err != nil {
		return err
	}

	if m.ShipmentDetails != nil {
		if err := m.ShipmentDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this shipment confirmation based on the context it is used
func (m *ShipmentConfirmation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFromParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentConfirmation) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Containers); i++ {

		if m.Containers[i] != nil {
			if err := m.Containers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ShipmentConfirmation) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ShipmentConfirmation) contextValidateSellingParty(ctx context.Context, formats strfmt.Registry) error {

	if m.SellingParty != nil {
		if err := m.SellingParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentConfirmation) contextValidateShipFromParty(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFromParty != nil {
		if err := m.ShipFromParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFromParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFromParty")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentConfirmation) contextValidateShipmentDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentDetails != nil {
		if err := m.ShipmentDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentConfirmation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentConfirmation) UnmarshalBinary(b []byte) error {
	var res ShipmentConfirmation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}