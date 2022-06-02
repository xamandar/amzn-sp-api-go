// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionReference transaction reference
//
// swagger:model TransactionReference
type TransactionReference struct {

	// GUID to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this transaction reference
func (m *TransactionReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this transaction reference based on context it is used
func (m *TransactionReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionReference) UnmarshalBinary(b []byte) error {
	var res TransactionReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}