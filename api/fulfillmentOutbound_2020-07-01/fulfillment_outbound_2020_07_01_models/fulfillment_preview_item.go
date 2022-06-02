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

// FulfillmentPreviewItem Item information for a shipment in a fulfillment order preview.
//
// swagger:model FulfillmentPreviewItem
type FulfillmentPreviewItem struct {

	// The estimated shipping weight of the item quantity for a single item, as identified by sellerSku, in a shipment.
	EstimatedShippingWeight *Weight `json:"estimatedShippingWeight,omitempty"`

	// The item quantity.
	// Required: true
	Quantity *Quantity `json:"quantity"`

	// A fulfillment order item identifier that the seller created with a call to the createFulfillmentOrder operation.
	// Required: true
	SellerFulfillmentOrderItemID *string `json:"sellerFulfillmentOrderItemId"`

	// The seller SKU of the item.
	// Required: true
	SellerSku *string `json:"sellerSku"`

	// The method used to calculate the estimated shipping weight.
	// Enum: [Package Dimensional]
	ShippingWeightCalculationMethod string `json:"shippingWeightCalculationMethod,omitempty"`
}

// Validate validates this fulfillment preview item
func (m *FulfillmentPreviewItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEstimatedShippingWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerFulfillmentOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingWeightCalculationMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentPreviewItem) validateEstimatedShippingWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.EstimatedShippingWeight) { // not required
		return nil
	}

	if m.EstimatedShippingWeight != nil {
		if err := m.EstimatedShippingWeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("estimatedShippingWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("estimatedShippingWeight")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentPreviewItem) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if m.Quantity != nil {
		if err := m.Quantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quantity")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentPreviewItem) validateSellerFulfillmentOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("sellerFulfillmentOrderItemId", "body", m.SellerFulfillmentOrderItemID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentPreviewItem) validateSellerSku(formats strfmt.Registry) error {

	if err := validate.Required("sellerSku", "body", m.SellerSku); err != nil {
		return err
	}

	return nil
}

var fulfillmentPreviewItemTypeShippingWeightCalculationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Package","Dimensional"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fulfillmentPreviewItemTypeShippingWeightCalculationMethodPropEnum = append(fulfillmentPreviewItemTypeShippingWeightCalculationMethodPropEnum, v)
	}
}

const (

	// FulfillmentPreviewItemShippingWeightCalculationMethodPackage captures enum value "Package"
	FulfillmentPreviewItemShippingWeightCalculationMethodPackage string = "Package"

	// FulfillmentPreviewItemShippingWeightCalculationMethodDimensional captures enum value "Dimensional"
	FulfillmentPreviewItemShippingWeightCalculationMethodDimensional string = "Dimensional"
)

// prop value enum
func (m *FulfillmentPreviewItem) validateShippingWeightCalculationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fulfillmentPreviewItemTypeShippingWeightCalculationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FulfillmentPreviewItem) validateShippingWeightCalculationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingWeightCalculationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateShippingWeightCalculationMethodEnum("shippingWeightCalculationMethod", "body", m.ShippingWeightCalculationMethod); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fulfillment preview item based on the context it is used
func (m *FulfillmentPreviewItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEstimatedShippingWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentPreviewItem) contextValidateEstimatedShippingWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.EstimatedShippingWeight != nil {
		if err := m.EstimatedShippingWeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("estimatedShippingWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("estimatedShippingWeight")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentPreviewItem) contextValidateQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.Quantity != nil {
		if err := m.Quantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FulfillmentPreviewItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FulfillmentPreviewItem) UnmarshalBinary(b []byte) error {
	var res FulfillmentPreviewItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}