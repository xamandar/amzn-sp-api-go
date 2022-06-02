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

// SellerFeedbackType Information about the seller's feedback, including the percentage of positive feedback, and the total number of ratings received.
//
// swagger:model SellerFeedbackType
type SellerFeedbackType struct {

	// The number of ratings received about the seller.
	// Required: true
	FeedbackCount *int64 `json:"FeedbackCount"`

	// The percentage of positive feedback for the seller in the past 365 days.
	SellerPositiveFeedbackRating float64 `json:"SellerPositiveFeedbackRating,omitempty"`
}

// Validate validates this seller feedback type
func (m *SellerFeedbackType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeedbackCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SellerFeedbackType) validateFeedbackCount(formats strfmt.Registry) error {

	if err := validate.Required("FeedbackCount", "body", m.FeedbackCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this seller feedback type based on context it is used
func (m *SellerFeedbackType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SellerFeedbackType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SellerFeedbackType) UnmarshalBinary(b []byte) error {
	var res SellerFeedbackType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}