// Code generated by go-swagger; DO NOT EDIT.

package shipment_invoicing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Address The shipping address details of the shipment.
//
// swagger:model Address
type Address struct {

	// The street address.
	AddressLine1 string `json:"AddressLine1,omitempty"`

	// Additional street address information, if required.
	AddressLine2 string `json:"AddressLine2,omitempty"`

	// Additional street address information, if required.
	AddressLine3 string `json:"AddressLine3,omitempty"`

	// address type
	AddressType AddressTypeEnum `json:"AddressType,omitempty"`

	// The city.
	City string `json:"City,omitempty"`

	// The country code.
	CountryCode string `json:"CountryCode,omitempty"`

	// The county.
	County string `json:"County,omitempty"`

	// The district.
	District string `json:"District,omitempty"`

	// The name.
	Name string `json:"Name,omitempty"`

	// The phone number.
	Phone string `json:"Phone,omitempty"`

	// The postal code.
	PostalCode string `json:"PostalCode,omitempty"`

	// The state or region.
	StateOrRegion string `json:"StateOrRegion,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) validateAddressType(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressType) { // not required
		return nil
	}

	if err := m.AddressType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AddressType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this address based on the context it is used
func (m *Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddressType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) contextValidateAddressType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AddressType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AddressType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
