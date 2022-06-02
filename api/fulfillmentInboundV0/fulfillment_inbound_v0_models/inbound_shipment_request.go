// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InboundShipmentRequest The request schema for an inbound shipment.
//
// swagger:model InboundShipmentRequest
type InboundShipmentRequest struct {

	// inbound shipment header
	// Required: true
	InboundShipmentHeader *InboundShipmentHeader `json:"InboundShipmentHeader"`

	// inbound shipment items
	// Required: true
	InboundShipmentItems InboundShipmentItemList `json:"InboundShipmentItems"`

	// A marketplace identifier. Specifies the marketplace where the product would be stored.
	// Required: true
	MarketplaceID *string `json:"MarketplaceId"`
}

// Validate validates this inbound shipment request
func (m *InboundShipmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInboundShipmentHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInboundShipmentItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundShipmentRequest) validateInboundShipmentHeader(formats strfmt.Registry) error {

	if err := validate.Required("InboundShipmentHeader", "body", m.InboundShipmentHeader); err != nil {
		return err
	}

	if m.InboundShipmentHeader != nil {
		if err := m.InboundShipmentHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InboundShipmentHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InboundShipmentHeader")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipmentRequest) validateInboundShipmentItems(formats strfmt.Registry) error {

	if err := validate.Required("InboundShipmentItems", "body", m.InboundShipmentItems); err != nil {
		return err
	}

	if err := m.InboundShipmentItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InboundShipmentItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InboundShipmentItems")
		}
		return err
	}

	return nil
}

func (m *InboundShipmentRequest) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("MarketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inbound shipment request based on the context it is used
func (m *InboundShipmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInboundShipmentHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInboundShipmentItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundShipmentRequest) contextValidateInboundShipmentHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.InboundShipmentHeader != nil {
		if err := m.InboundShipmentHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InboundShipmentHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InboundShipmentHeader")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipmentRequest) contextValidateInboundShipmentItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InboundShipmentItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InboundShipmentItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InboundShipmentItems")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InboundShipmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundShipmentRequest) UnmarshalBinary(b []byte) error {
	var res InboundShipmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}