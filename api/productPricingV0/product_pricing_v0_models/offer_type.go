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

// OfferType offer type
//
// swagger:model OfferType
type OfferType struct {

	// Contains pricing information that includes promotions and contains the shipping cost.
	// Required: true
	BuyingPrice *PriceType `json:"BuyingPrice"`

	// The fulfillment channel for the offer listing. Possible values:
	//
	// * Amazon - Fulfilled by Amazon.
	// * Merchant - Fulfilled by the seller.
	// Required: true
	FulfillmentChannel *string `json:"FulfillmentChannel"`

	// The item condition for the offer listing. Possible values: New, Used, Collectible, Refurbished, or Club.
	// Required: true
	ItemCondition *string `json:"ItemCondition"`

	// The item subcondition for the offer listing. Possible values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, or Other.
	// Required: true
	ItemSubCondition *string `json:"ItemSubCondition"`

	// The current price excluding any promotions that apply to the product. Excludes the shipping cost.
	// Required: true
	RegularPrice *MoneyType `json:"RegularPrice"`

	// The seller stock keeping unit (SKU) of the item.
	// Required: true
	SellerSKU *string `json:"SellerSKU"`

	// The current listing price for Business buyers.
	BusinessPrice *MoneyType `json:"businessPrice,omitempty"`

	// Indicates the type of customer that the offer is valid for.
	OfferType OfferCustomerType `json:"offerType,omitempty"`

	// quantity discount prices
	QuantityDiscountPrices []*QuantityDiscountPriceType `json:"quantityDiscountPrices"`
}

// Validate validates this offer type
func (m *OfferType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuyingPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSubCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegularPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerSKU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusinessPrice(formats); err != nil {
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

func (m *OfferType) validateBuyingPrice(formats strfmt.Registry) error {

	if err := validate.Required("BuyingPrice", "body", m.BuyingPrice); err != nil {
		return err
	}

	if m.BuyingPrice != nil {
		if err := m.BuyingPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OfferType) validateFulfillmentChannel(formats strfmt.Registry) error {

	if err := validate.Required("FulfillmentChannel", "body", m.FulfillmentChannel); err != nil {
		return err
	}

	return nil
}

func (m *OfferType) validateItemCondition(formats strfmt.Registry) error {

	if err := validate.Required("ItemCondition", "body", m.ItemCondition); err != nil {
		return err
	}

	return nil
}

func (m *OfferType) validateItemSubCondition(formats strfmt.Registry) error {

	if err := validate.Required("ItemSubCondition", "body", m.ItemSubCondition); err != nil {
		return err
	}

	return nil
}

func (m *OfferType) validateRegularPrice(formats strfmt.Registry) error {

	if err := validate.Required("RegularPrice", "body", m.RegularPrice); err != nil {
		return err
	}

	if m.RegularPrice != nil {
		if err := m.RegularPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegularPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RegularPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OfferType) validateSellerSKU(formats strfmt.Registry) error {

	if err := validate.Required("SellerSKU", "body", m.SellerSKU); err != nil {
		return err
	}

	return nil
}

func (m *OfferType) validateBusinessPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessPrice) { // not required
		return nil
	}

	if m.BusinessPrice != nil {
		if err := m.BusinessPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OfferType) validateOfferType(formats strfmt.Registry) error {
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

func (m *OfferType) validateQuantityDiscountPrices(formats strfmt.Registry) error {
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

// ContextValidate validate this offer type based on the context it is used
func (m *OfferType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuyingPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegularPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBusinessPrice(ctx, formats); err != nil {
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

func (m *OfferType) contextValidateBuyingPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyingPrice != nil {
		if err := m.BuyingPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OfferType) contextValidateRegularPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.RegularPrice != nil {
		if err := m.RegularPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegularPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RegularPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OfferType) contextValidateBusinessPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.BusinessPrice != nil {
		if err := m.BusinessPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OfferType) contextValidateOfferType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OfferType) contextValidateQuantityDiscountPrices(ctx context.Context, formats strfmt.Registry) error {

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
func (m *OfferType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OfferType) UnmarshalBinary(b []byte) error {
	var res OfferType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
