// Code generated by go-swagger; DO NOT EDIT.

package sellers_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Marketplace Detailed information about an Amazon market where a seller can list items for sale and customers can view and purchase items.
//
// swagger:model Marketplace
type Marketplace struct {

	// The ISO 3166-1 alpha-2 format country code of the marketplace.
	// Required: true
	// Pattern: ^([A-Z]{2})$
	CountryCode *string `json:"countryCode"`

	// The ISO 4217 format currency code of the marketplace.
	// Required: true
	DefaultCurrencyCode *string `json:"defaultCurrencyCode"`

	// The ISO 639-1 format language code of the marketplace.
	// Required: true
	DefaultLanguageCode *string `json:"defaultLanguageCode"`

	// The domain name of the marketplace.
	// Required: true
	DomainName *string `json:"domainName"`

	// The encrypted marketplace value.
	// Required: true
	ID *string `json:"id"`

	// Marketplace name.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this marketplace
func (m *Marketplace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Marketplace) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.Pattern("countryCode", "body", *m.CountryCode, `^([A-Z]{2})$`); err != nil {
		return err
	}

	return nil
}

func (m *Marketplace) validateDefaultCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("defaultCurrencyCode", "body", m.DefaultCurrencyCode); err != nil {
		return err
	}

	return nil
}

func (m *Marketplace) validateDefaultLanguageCode(formats strfmt.Registry) error {

	if err := validate.Required("defaultLanguageCode", "body", m.DefaultLanguageCode); err != nil {
		return err
	}

	return nil
}

func (m *Marketplace) validateDomainName(formats strfmt.Registry) error {

	if err := validate.Required("domainName", "body", m.DomainName); err != nil {
		return err
	}

	return nil
}

func (m *Marketplace) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Marketplace) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this marketplace based on context it is used
func (m *Marketplace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Marketplace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Marketplace) UnmarshalBinary(b []byte) error {
	var res Marketplace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}