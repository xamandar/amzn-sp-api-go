// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2020_12_01_models

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

// ItemVariationsByMarketplace Variation details for the Amazon catalog item for the indicated Amazon marketplace.
//
// swagger:model ItemVariationsByMarketplace
type ItemVariationsByMarketplace struct {

	// Identifiers (ASINs) of the related items.
	// Required: true
	Asins []string `json:"asins"`

	// Amazon marketplace identifier.
	// Required: true
	MarketplaceID *string `json:"marketplaceId"`

	// Type of variation relationship of the Amazon catalog item in the request to the related item(s): "PARENT" or "CHILD".
	// Example: PARENT
	// Required: true
	// Enum: [PARENT CHILD]
	VariationType *string `json:"variationType"`
}

// Validate validates this item variations by marketplace
func (m *ItemVariationsByMarketplace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemVariationsByMarketplace) validateAsins(formats strfmt.Registry) error {

	if err := validate.Required("asins", "body", m.Asins); err != nil {
		return err
	}

	return nil
}

func (m *ItemVariationsByMarketplace) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	return nil
}

var itemVariationsByMarketplaceTypeVariationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARENT","CHILD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemVariationsByMarketplaceTypeVariationTypePropEnum = append(itemVariationsByMarketplaceTypeVariationTypePropEnum, v)
	}
}

const (

	// ItemVariationsByMarketplaceVariationTypePARENT captures enum value "PARENT"
	ItemVariationsByMarketplaceVariationTypePARENT string = "PARENT"

	// ItemVariationsByMarketplaceVariationTypeCHILD captures enum value "CHILD"
	ItemVariationsByMarketplaceVariationTypeCHILD string = "CHILD"
)

// prop value enum
func (m *ItemVariationsByMarketplace) validateVariationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemVariationsByMarketplaceTypeVariationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemVariationsByMarketplace) validateVariationType(formats strfmt.Registry) error {

	if err := validate.Required("variationType", "body", m.VariationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateVariationTypeEnum("variationType", "body", *m.VariationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this item variations by marketplace based on context it is used
func (m *ItemVariationsByMarketplace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemVariationsByMarketplace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemVariationsByMarketplace) UnmarshalBinary(b []byte) error {
	var res ItemVariationsByMarketplace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
