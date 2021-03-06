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

// CompetitivePricingType Competitive pricing information for the item.
//
// swagger:model CompetitivePricingType
type CompetitivePricingType struct {

	// competitive prices
	// Required: true
	CompetitivePrices CompetitivePriceList `json:"CompetitivePrices"`

	// number of offer listings
	// Required: true
	NumberOfOfferListings NumberOfOfferListingsList `json:"NumberOfOfferListings"`

	// The trade-in value of the item in the trade-in program.
	TradeInValue *MoneyType `json:"TradeInValue,omitempty"`
}

// Validate validates this competitive pricing type
func (m *CompetitivePricingType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompetitivePrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfOfferListings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeInValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompetitivePricingType) validateCompetitivePrices(formats strfmt.Registry) error {

	if err := validate.Required("CompetitivePrices", "body", m.CompetitivePrices); err != nil {
		return err
	}

	if err := m.CompetitivePrices.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CompetitivePrices")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CompetitivePrices")
		}
		return err
	}

	return nil
}

func (m *CompetitivePricingType) validateNumberOfOfferListings(formats strfmt.Registry) error {

	if err := validate.Required("NumberOfOfferListings", "body", m.NumberOfOfferListings); err != nil {
		return err
	}

	if err := m.NumberOfOfferListings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NumberOfOfferListings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NumberOfOfferListings")
		}
		return err
	}

	return nil
}

func (m *CompetitivePricingType) validateTradeInValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TradeInValue) { // not required
		return nil
	}

	if m.TradeInValue != nil {
		if err := m.TradeInValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TradeInValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TradeInValue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this competitive pricing type based on the context it is used
func (m *CompetitivePricingType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompetitivePrices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumberOfOfferListings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTradeInValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompetitivePricingType) contextValidateCompetitivePrices(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CompetitivePrices.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CompetitivePrices")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CompetitivePrices")
		}
		return err
	}

	return nil
}

func (m *CompetitivePricingType) contextValidateNumberOfOfferListings(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NumberOfOfferListings.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NumberOfOfferListings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("NumberOfOfferListings")
		}
		return err
	}

	return nil
}

func (m *CompetitivePricingType) contextValidateTradeInValue(ctx context.Context, formats strfmt.Registry) error {

	if m.TradeInValue != nil {
		if err := m.TradeInValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TradeInValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TradeInValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompetitivePricingType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompetitivePricingType) UnmarshalBinary(b []byte) error {
	var res CompetitivePricingType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
