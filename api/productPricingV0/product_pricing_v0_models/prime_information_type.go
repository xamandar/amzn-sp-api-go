// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrimeInformationType Amazon Prime information.
//
// swagger:model PrimeInformationType
type PrimeInformationType struct {

	// Indicates whether the offer is an Amazon Prime offer throughout the entire marketplace where it is listed.
	// Required: true
	IsNationalPrime *bool `json:"IsNationalPrime"`

	// Indicates whether the offer is an Amazon Prime offer.
	// Required: true
	IsPrime *bool `json:"IsPrime"`
}

// Validate validates this prime information type
func (m *PrimeInformationType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsNationalPrime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPrime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrimeInformationType) validateIsNationalPrime(formats strfmt.Registry) error {

	if err := validate.Required("IsNationalPrime", "body", m.IsNationalPrime); err != nil {
		return err
	}

	return nil
}

func (m *PrimeInformationType) validateIsPrime(formats strfmt.Registry) error {

	if err := validate.Required("IsPrime", "body", m.IsPrime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this prime information type based on context it is used
func (m *PrimeInformationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrimeInformationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrimeInformationType) UnmarshalBinary(b []byte) error {
	var res PrimeInformationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}