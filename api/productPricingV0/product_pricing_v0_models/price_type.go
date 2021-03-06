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

// PriceType price type
//
// swagger:model PriceType
type PriceType struct {

	// The value calculated by adding ListingPrice + Shipping - Points. Note that if the landed price is not returned, the listing price represents the product with the lowest landed price.
	LandedPrice *MoneyType `json:"LandedPrice,omitempty"`

	// The listing price of the item including any promotions that apply.
	// Required: true
	ListingPrice *MoneyType `json:"ListingPrice"`

	// The number of Amazon Points offered with the purchase of an item, and their monetary value.
	Points *Points `json:"Points,omitempty"`

	// The shipping cost of the product. Note that the shipping cost is not always available.
	Shipping *MoneyType `json:"Shipping,omitempty"`
}

// Validate validates this price type
func (m *PriceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLandedPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListingPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipping(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceType) validateLandedPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.LandedPrice) { // not required
		return nil
	}

	if m.LandedPrice != nil {
		if err := m.LandedPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LandedPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LandedPrice")
			}
			return err
		}
	}

	return nil
}

func (m *PriceType) validateListingPrice(formats strfmt.Registry) error {

	if err := validate.Required("ListingPrice", "body", m.ListingPrice); err != nil {
		return err
	}

	if m.ListingPrice != nil {
		if err := m.ListingPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ListingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ListingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *PriceType) validatePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if m.Points != nil {
		if err := m.Points.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Points")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Points")
			}
			return err
		}
	}

	return nil
}

func (m *PriceType) validateShipping(formats strfmt.Registry) error {
	if swag.IsZero(m.Shipping) { // not required
		return nil
	}

	if m.Shipping != nil {
		if err := m.Shipping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Shipping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Shipping")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this price type based on the context it is used
func (m *PriceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLandedPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateListingPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceType) contextValidateLandedPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.LandedPrice != nil {
		if err := m.LandedPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LandedPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LandedPrice")
			}
			return err
		}
	}

	return nil
}

func (m *PriceType) contextValidateListingPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.ListingPrice != nil {
		if err := m.ListingPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ListingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ListingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *PriceType) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

	if m.Points != nil {
		if err := m.Points.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Points")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Points")
			}
			return err
		}
	}

	return nil
}

func (m *PriceType) contextValidateShipping(ctx context.Context, formats strfmt.Registry) error {

	if m.Shipping != nil {
		if err := m.Shipping.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Shipping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Shipping")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PriceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceType) UnmarshalBinary(b []byte) error {
	var res PriceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
