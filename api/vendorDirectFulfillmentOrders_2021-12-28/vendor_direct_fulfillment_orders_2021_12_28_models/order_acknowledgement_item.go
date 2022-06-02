// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_2021_12_28_models

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

// OrderAcknowledgementItem Details of an individual order being acknowledged.
//
// swagger:model OrderAcknowledgementItem
type OrderAcknowledgementItem struct {

	// The date and time when the order is acknowledged, in ISO-8601 date/time format. For example: 2018-07-16T23:00:00Z / 2018-07-16T23:00:00-05:00 / 2018-07-16T23:00:00-08:00.
	// Required: true
	// Format: date-time
	AcknowledgementDate *strfmt.DateTime `json:"acknowledgementDate"`

	// Status of acknowledgement.
	// Required: true
	AcknowledgementStatus *AcknowledgementStatus `json:"acknowledgementStatus"`

	// Item details including acknowledged quantity.
	// Required: true
	ItemAcknowledgements []*OrderItemAcknowledgement `json:"itemAcknowledgements"`

	// The purchase order number for this order. Formatting Notes: alpha-numeric code.
	// Required: true
	PurchaseOrderNumber *string `json:"purchaseOrderNumber"`

	// PartyID as vendor code.
	// Required: true
	SellingParty *PartyIdentification `json:"sellingParty"`

	// PartyID as the vendor's warehouseId.
	// Required: true
	ShipFromParty *PartyIdentification `json:"shipFromParty"`

	// The vendor's order number for this order.
	// Required: true
	VendorOrderNumber *string `json:"vendorOrderNumber"`
}

// Validate validates this order acknowledgement item
func (m *OrderAcknowledgementItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcknowledgementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcknowledgementStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemAcknowledgements(formats); err != nil {
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

	if err := m.validateVendorOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAcknowledgementItem) validateAcknowledgementDate(formats strfmt.Registry) error {

	if err := validate.Required("acknowledgementDate", "body", m.AcknowledgementDate); err != nil {
		return err
	}

	if err := validate.FormatOf("acknowledgementDate", "body", "date-time", m.AcknowledgementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrderAcknowledgementItem) validateAcknowledgementStatus(formats strfmt.Registry) error {

	if err := validate.Required("acknowledgementStatus", "body", m.AcknowledgementStatus); err != nil {
		return err
	}

	if m.AcknowledgementStatus != nil {
		if err := m.AcknowledgementStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acknowledgementStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acknowledgementStatus")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAcknowledgementItem) validateItemAcknowledgements(formats strfmt.Registry) error {

	if err := validate.Required("itemAcknowledgements", "body", m.ItemAcknowledgements); err != nil {
		return err
	}

	for i := 0; i < len(m.ItemAcknowledgements); i++ {
		if swag.IsZero(m.ItemAcknowledgements[i]) { // not required
			continue
		}

		if m.ItemAcknowledgements[i] != nil {
			if err := m.ItemAcknowledgements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderAcknowledgementItem) validatePurchaseOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderNumber", "body", m.PurchaseOrderNumber); err != nil {
		return err
	}

	return nil
}

func (m *OrderAcknowledgementItem) validateSellingParty(formats strfmt.Registry) error {

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

func (m *OrderAcknowledgementItem) validateShipFromParty(formats strfmt.Registry) error {

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

func (m *OrderAcknowledgementItem) validateVendorOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("vendorOrderNumber", "body", m.VendorOrderNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order acknowledgement item based on the context it is used
func (m *OrderAcknowledgementItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcknowledgementStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemAcknowledgements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFromParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAcknowledgementItem) contextValidateAcknowledgementStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.AcknowledgementStatus != nil {
		if err := m.AcknowledgementStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acknowledgementStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acknowledgementStatus")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAcknowledgementItem) contextValidateItemAcknowledgements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ItemAcknowledgements); i++ {

		if m.ItemAcknowledgements[i] != nil {
			if err := m.ItemAcknowledgements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderAcknowledgementItem) contextValidateSellingParty(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OrderAcknowledgementItem) contextValidateShipFromParty(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *OrderAcknowledgementItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderAcknowledgementItem) UnmarshalBinary(b []byte) error {
	var res OrderAcknowledgementItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}