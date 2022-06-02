// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PromisedDeliveryDate The promised delivery date and time of a shipment.
//
// swagger:model PromisedDeliveryDate
type PromisedDeliveryDate strfmt.DateTime

// UnmarshalJSON sets a PromisedDeliveryDate value from JSON input
func (m *PromisedDeliveryDate) UnmarshalJSON(b []byte) error {
	return ((*strfmt.DateTime)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a PromisedDeliveryDate value as JSON output
func (m PromisedDeliveryDate) MarshalJSON() ([]byte, error) {
	return (strfmt.DateTime(m)).MarshalJSON()
}

// Validate validates this promised delivery date
func (m PromisedDeliveryDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.FormatOf("", "body", "date-time", strfmt.DateTime(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this promised delivery date based on context it is used
func (m PromisedDeliveryDate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PromisedDeliveryDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromisedDeliveryDate) UnmarshalBinary(b []byte) error {
	var res PromisedDeliveryDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}