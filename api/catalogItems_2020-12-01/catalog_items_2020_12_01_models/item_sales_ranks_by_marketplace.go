// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2020_12_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemSalesRanksByMarketplace Sales ranks of an Amazon catalog item for the indicated Amazon marketplace.
//
// swagger:model ItemSalesRanksByMarketplace
type ItemSalesRanksByMarketplace struct {

	// Amazon marketplace identifier.
	// Required: true
	MarketplaceID *string `json:"marketplaceId"`

	// Sales ranks of an Amazon catalog item for an Amazon marketplace.
	// Required: true
	Ranks []*ItemSalesRank `json:"ranks"`
}

// Validate validates this item sales ranks by marketplace
func (m *ItemSalesRanksByMarketplace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRanks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSalesRanksByMarketplace) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	return nil
}

func (m *ItemSalesRanksByMarketplace) validateRanks(formats strfmt.Registry) error {

	if err := validate.Required("ranks", "body", m.Ranks); err != nil {
		return err
	}

	for i := 0; i < len(m.Ranks); i++ {
		if swag.IsZero(m.Ranks[i]) { // not required
			continue
		}

		if m.Ranks[i] != nil {
			if err := m.Ranks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ranks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this item sales ranks by marketplace based on the context it is used
func (m *ItemSalesRanksByMarketplace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRanks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSalesRanksByMarketplace) contextValidateRanks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ranks); i++ {

		if m.Ranks[i] != nil {
			if err := m.Ranks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ranks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemSalesRanksByMarketplace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSalesRanksByMarketplace) UnmarshalBinary(b []byte) error {
	var res ItemSalesRanksByMarketplace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}