// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OfferCountType The total number of offers for the specified condition and fulfillment channel.
//
// swagger:model OfferCountType
type OfferCountType struct {

	// The number of offers in a fulfillment channel that meet a specific condition.
	OfferCount int32 `json:"OfferCount,omitempty"`

	// Indicates the condition of the item. For example: New, Used, Collectible, Refurbished, or Club.
	Condition string `json:"condition,omitempty"`

	// Indicates whether the item is fulfilled by Amazon or by the seller.
	FulfillmentChannel FulfillmentChannelType `json:"fulfillmentChannel,omitempty"`
}

// Validate validates this offer count type
func (m *OfferCountType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFulfillmentChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OfferCountType) validateFulfillmentChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentChannel) { // not required
		return nil
	}

	if err := m.FulfillmentChannel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentChannel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentChannel")
		}
		return err
	}

	return nil
}

// ContextValidate validate this offer count type based on the context it is used
func (m *OfferCountType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFulfillmentChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OfferCountType) contextValidateFulfillmentChannel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentChannel.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentChannel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentChannel")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OfferCountType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OfferCountType) UnmarshalBinary(b []byte) error {
	var res OfferCountType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}