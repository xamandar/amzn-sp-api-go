// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShipsFromType The state and country from where the item is shipped.
//
// swagger:model ShipsFromType
type ShipsFromType struct {

	// The country from where the item is shipped.
	Country string `json:"Country,omitempty"`

	// The state from where the item is shipped.
	State string `json:"State,omitempty"`
}

// Validate validates this ships from type
func (m *ShipsFromType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ships from type based on context it is used
func (m *ShipsFromType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShipsFromType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipsFromType) UnmarshalBinary(b []byte) error {
	var res ShipsFromType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}