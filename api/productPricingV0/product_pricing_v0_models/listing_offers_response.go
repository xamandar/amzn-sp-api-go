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

// ListingOffersResponse listing offers response
//
// swagger:model ListingOffersResponse
type ListingOffersResponse struct {
	BatchOffersResponse

	// request
	Request *ListingOffersRequestParams `json:"request,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ListingOffersResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BatchOffersResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BatchOffersResponse = aO0

	// AO1
	var dataAO1 struct {
		Request *ListingOffersRequestParams `json:"request,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Request = dataAO1.Request

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ListingOffersResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BatchOffersResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Request *ListingOffersRequestParams `json:"request,omitempty"`
	}

	dataAO1.Request = m.Request

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this listing offers response
func (m *ListingOffersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BatchOffersResponse
	if err := m.BatchOffersResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListingOffersResponse) validateRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this listing offers response based on the context it is used
func (m *ListingOffersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BatchOffersResponse
	if err := m.BatchOffersResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListingOffersResponse) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {
		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListingOffersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListingOffersResponse) UnmarshalBinary(b []byte) error {
	var res ListingOffersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
