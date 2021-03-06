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

// ShipmentItem An item of a shipment, refund, guarantee claim, or chargeback.
//
// swagger:model ShipmentItem
type ShipmentItem struct {

	// The cost of Amazon Points granted for a shipment item.
	CostOfPointsGranted *Currency `json:"CostOfPointsGranted,omitempty"`

	// The cost of Amazon Points returned for a shipment item. This value is only returned for refunds, guarantee claims, and chargeback events.
	CostOfPointsReturned *Currency `json:"CostOfPointsReturned,omitempty"`

	// A list of charge adjustments associated with the shipment item. This value is only returned for refunds, guarantee claims, and chargeback events.
	ItemChargeAdjustmentList ChargeComponentList `json:"ItemChargeAdjustmentList,omitempty"`

	// A list of charges associated with the shipment item.
	ItemChargeList ChargeComponentList `json:"ItemChargeList,omitempty"`

	// A list of fee adjustments associated with the shipment item. This value is only returned for refunds, guarantee claims, and chargeback events.
	ItemFeeAdjustmentList FeeComponentList `json:"ItemFeeAdjustmentList,omitempty"`

	// A list of fees associated with the shipment item.
	ItemFeeList FeeComponentList `json:"ItemFeeList,omitempty"`

	// A list of taxes withheld information for a shipment item.
	ItemTaxWithheldList TaxWithheldComponentList `json:"ItemTaxWithheldList,omitempty"`

	// An Amazon-defined order adjustment identifier defined for refunds, guarantee claims, and chargeback events.
	OrderAdjustmentItemID string `json:"OrderAdjustmentItemId,omitempty"`

	// An Amazon-defined order item identifier.
	OrderItemID string `json:"OrderItemId,omitempty"`

	// A list of promotion adjustments associated with the shipment item. This value is only returned for refunds, guarantee claims, and chargeback events.
	PromotionAdjustmentList PromotionList `json:"PromotionAdjustmentList,omitempty"`

	// promotion list
	PromotionList PromotionList `json:"PromotionList,omitempty"`

	// The number of items shipped.
	QuantityShipped int32 `json:"QuantityShipped,omitempty"`

	// The seller SKU of the item. The seller SKU is qualified by the seller's seller ID, which is included with every call to the Selling Partner API.
	SellerSKU string `json:"SellerSKU,omitempty"`
}

// Validate validates this shipment item
func (m *ShipmentItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostOfPointsGranted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostOfPointsReturned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemChargeAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemChargeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemFeeAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemFeeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemTaxWithheldList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromotionAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromotionList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentItem) validateCostOfPointsGranted(formats strfmt.Registry) error {
	if swag.IsZero(m.CostOfPointsGranted) { // not required
		return nil
	}

	if m.CostOfPointsGranted != nil {
		if err := m.CostOfPointsGranted.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CostOfPointsGranted")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CostOfPointsGranted")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentItem) validateCostOfPointsReturned(formats strfmt.Registry) error {
	if swag.IsZero(m.CostOfPointsReturned) { // not required
		return nil
	}

	if m.CostOfPointsReturned != nil {
		if err := m.CostOfPointsReturned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CostOfPointsReturned")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CostOfPointsReturned")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentItem) validateItemChargeAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemChargeAdjustmentList) { // not required
		return nil
	}

	if err := m.ItemChargeAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemChargeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemChargeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) validateItemChargeList(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemChargeList) { // not required
		return nil
	}

	if err := m.ItemChargeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemChargeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemChargeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) validateItemFeeAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemFeeAdjustmentList) { // not required
		return nil
	}

	if err := m.ItemFeeAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemFeeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemFeeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) validateItemFeeList(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemFeeList) { // not required
		return nil
	}

	if err := m.ItemFeeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemFeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemFeeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) validateItemTaxWithheldList(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemTaxWithheldList) { // not required
		return nil
	}

	if err := m.ItemTaxWithheldList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemTaxWithheldList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemTaxWithheldList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) validatePromotionAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.PromotionAdjustmentList) { // not required
		return nil
	}

	if err := m.PromotionAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PromotionAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PromotionAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) validatePromotionList(formats strfmt.Registry) error {
	if swag.IsZero(m.PromotionList) { // not required
		return nil
	}

	if err := m.PromotionList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PromotionList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PromotionList")
		}
		return err
	}

	return nil
}

// ContextValidate validate this shipment item based on the context it is used
func (m *ShipmentItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCostOfPointsGranted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCostOfPointsReturned(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemChargeAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemChargeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemFeeAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemFeeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemTaxWithheldList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePromotionAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePromotionList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentItem) contextValidateCostOfPointsGranted(ctx context.Context, formats strfmt.Registry) error {

	if m.CostOfPointsGranted != nil {
		if err := m.CostOfPointsGranted.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CostOfPointsGranted")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CostOfPointsGranted")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentItem) contextValidateCostOfPointsReturned(ctx context.Context, formats strfmt.Registry) error {

	if m.CostOfPointsReturned != nil {
		if err := m.CostOfPointsReturned.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CostOfPointsReturned")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CostOfPointsReturned")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentItem) contextValidateItemChargeAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemChargeAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemChargeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemChargeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) contextValidateItemChargeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemChargeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemChargeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemChargeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) contextValidateItemFeeAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemFeeAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemFeeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemFeeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) contextValidateItemFeeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemFeeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemFeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemFeeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) contextValidateItemTaxWithheldList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemTaxWithheldList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemTaxWithheldList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemTaxWithheldList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) contextValidatePromotionAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PromotionAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PromotionAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PromotionAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentItem) contextValidatePromotionList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PromotionList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PromotionList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PromotionList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentItem) UnmarshalBinary(b []byte) error {
	var res ShipmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
