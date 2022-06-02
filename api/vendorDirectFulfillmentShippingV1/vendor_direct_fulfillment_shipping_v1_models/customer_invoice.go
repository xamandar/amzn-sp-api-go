// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomerInvoice customer invoice
//
// swagger:model CustomerInvoice
type CustomerInvoice struct {

	// The Base64encoded customer invoice.
	// Required: true
	Content *string `json:"content"`

	// The purchase order number for this order.
	// Required: true
	// Pattern: ^[a-zA-Z0-9]+$
	PurchaseOrderNumber *string `json:"purchaseOrderNumber"`
}

// Validate validates this customer invoice
func (m *CustomerInvoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerInvoice) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *CustomerInvoice) validatePurchaseOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderNumber", "body", m.PurchaseOrderNumber); err != nil {
		return err
	}

	if err := validate.Pattern("purchaseOrderNumber", "body", *m.PurchaseOrderNumber, `^[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this customer invoice based on context it is used
func (m *CustomerInvoice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerInvoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerInvoice) UnmarshalBinary(b []byte) error {
	var res CustomerInvoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}