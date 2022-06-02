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

// CartonReferenceDetails carton reference details
//
// swagger:model CartonReferenceDetails
type CartonReferenceDetails struct {

	// Pallet level carton count is mandatory for single item pallet and optional for mixed item pallet.
	CartonCount int64 `json:"cartonCount,omitempty"`

	// Array of reference numbers for the carton that are part of this pallet/shipment. Please provide the cartonSequenceNumber from the 'cartons' segment to refer to that carton's details here.
	// Required: true
	CartonReferenceNumbers []string `json:"cartonReferenceNumbers"`
}

// Validate validates this carton reference details
func (m *CartonReferenceDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCartonReferenceNumbers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartonReferenceDetails) validateCartonReferenceNumbers(formats strfmt.Registry) error {

	if err := validate.Required("cartonReferenceNumbers", "body", m.CartonReferenceNumbers); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this carton reference details based on context it is used
func (m *CartonReferenceDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartonReferenceDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartonReferenceDetails) UnmarshalBinary(b []byte) error {
	var res CartonReferenceDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}