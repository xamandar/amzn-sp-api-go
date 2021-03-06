// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

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

// OfferDetail offer detail
//
// swagger:model OfferDetail
type OfferDetail struct {

	// Information about the condition of the item.
	ConditionNotes string `json:"ConditionNotes,omitempty"`

	// When true, the offer is currently in the Buy Box. There can be up to two Buy Box winners at any time per ASIN, one that is eligible for Prime and one that is not eligible for Prime.
	IsBuyBoxWinner bool `json:"IsBuyBoxWinner,omitempty"`

	// When true, the seller of the item is eligible to win the Buy Box.
	IsFeaturedMerchant bool `json:"IsFeaturedMerchant,omitempty"`

	// When true, the offer is fulfilled by Amazon.
	// Required: true
	IsFulfilledByAmazon *bool `json:"IsFulfilledByAmazon"`

	// The price of the item.
	// Required: true
	ListingPrice *MoneyType `json:"ListingPrice"`

	// When true, this is the seller's offer.
	MyOffer bool `json:"MyOffer,omitempty"`

	// The number of Amazon Points offered with the purchase of an item.
	Points *Points `json:"Points,omitempty"`

	// Amazon Prime information.
	PrimeInformation *PrimeInformationType `json:"PrimeInformation,omitempty"`

	// Information about the seller's feedback, including the percentage of positive feedback, and the total number of ratings received.
	SellerFeedbackRating *SellerFeedbackType `json:"SellerFeedbackRating,omitempty"`

	// The seller identifier for the offer.
	SellerID string `json:"SellerId,omitempty"`

	// The shipping cost.
	// Required: true
	Shipping *MoneyType `json:"Shipping"`

	// The maximum time within which the item will likely be shipped once an order has been placed.
	// Required: true
	ShippingTime *DetailedShippingTimeType `json:"ShippingTime"`

	// The state and country from where the item is shipped.
	ShipsFrom *ShipsFromType `json:"ShipsFrom,omitempty"`

	// The subcondition of the item. Subcondition values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other.
	// Required: true
	SubCondition *string `json:"SubCondition"`

	// Indicates the type of customer that the offer is valid for.
	OfferType OfferCustomerType `json:"offerType,omitempty"`

	// quantity discount prices
	QuantityDiscountPrices []*QuantityDiscountPriceType `json:"quantityDiscountPrices"`
}

// Validate validates this offer detail
func (m *OfferDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsFulfilledByAmazon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListingPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimeInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerFeedbackRating(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipsFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantityDiscountPrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OfferDetail) validateIsFulfilledByAmazon(formats strfmt.Registry) error {

	if err := validate.Required("IsFulfilledByAmazon", "body", m.IsFulfilledByAmazon); err != nil {
		return err
	}

	return nil
}

func (m *OfferDetail) validateListingPrice(formats strfmt.Registry) error {

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

func (m *OfferDetail) validatePoints(formats strfmt.Registry) error {
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

func (m *OfferDetail) validatePrimeInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimeInformation) { // not required
		return nil
	}

	if m.PrimeInformation != nil {
		if err := m.PrimeInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrimeInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrimeInformation")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) validateSellerFeedbackRating(formats strfmt.Registry) error {
	if swag.IsZero(m.SellerFeedbackRating) { // not required
		return nil
	}

	if m.SellerFeedbackRating != nil {
		if err := m.SellerFeedbackRating.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SellerFeedbackRating")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SellerFeedbackRating")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) validateShipping(formats strfmt.Registry) error {

	if err := validate.Required("Shipping", "body", m.Shipping); err != nil {
		return err
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

func (m *OfferDetail) validateShippingTime(formats strfmt.Registry) error {

	if err := validate.Required("ShippingTime", "body", m.ShippingTime); err != nil {
		return err
	}

	if m.ShippingTime != nil {
		if err := m.ShippingTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingTime")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) validateShipsFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipsFrom) { // not required
		return nil
	}

	if m.ShipsFrom != nil {
		if err := m.ShipsFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipsFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipsFrom")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) validateSubCondition(formats strfmt.Registry) error {

	if err := validate.Required("SubCondition", "body", m.SubCondition); err != nil {
		return err
	}

	return nil
}

func (m *OfferDetail) validateOfferType(formats strfmt.Registry) error {
	if swag.IsZero(m.OfferType) { // not required
		return nil
	}

	if err := m.OfferType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("offerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("offerType")
		}
		return err
	}

	return nil
}

func (m *OfferDetail) validateQuantityDiscountPrices(formats strfmt.Registry) error {
	if swag.IsZero(m.QuantityDiscountPrices) { // not required
		return nil
	}

	for i := 0; i < len(m.QuantityDiscountPrices); i++ {
		if swag.IsZero(m.QuantityDiscountPrices[i]) { // not required
			continue
		}

		if m.QuantityDiscountPrices[i] != nil {
			if err := m.QuantityDiscountPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quantityDiscountPrices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quantityDiscountPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this offer detail based on the context it is used
func (m *OfferDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListingPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimeInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellerFeedbackRating(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipsFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOfferType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuantityDiscountPrices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OfferDetail) contextValidateListingPrice(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OfferDetail) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OfferDetail) contextValidatePrimeInformation(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimeInformation != nil {
		if err := m.PrimeInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrimeInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrimeInformation")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) contextValidateSellerFeedbackRating(ctx context.Context, formats strfmt.Registry) error {

	if m.SellerFeedbackRating != nil {
		if err := m.SellerFeedbackRating.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SellerFeedbackRating")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SellerFeedbackRating")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) contextValidateShipping(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OfferDetail) contextValidateShippingTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingTime != nil {
		if err := m.ShippingTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingTime")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) contextValidateShipsFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipsFrom != nil {
		if err := m.ShipsFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipsFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipsFrom")
			}
			return err
		}
	}

	return nil
}

func (m *OfferDetail) contextValidateOfferType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OfferType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("offerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("offerType")
		}
		return err
	}

	return nil
}

func (m *OfferDetail) contextValidateQuantityDiscountPrices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QuantityDiscountPrices); i++ {

		if m.QuantityDiscountPrices[i] != nil {
			if err := m.QuantityDiscountPrices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quantityDiscountPrices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quantityDiscountPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OfferDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OfferDetail) UnmarshalBinary(b []byte) error {
	var res OfferDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
