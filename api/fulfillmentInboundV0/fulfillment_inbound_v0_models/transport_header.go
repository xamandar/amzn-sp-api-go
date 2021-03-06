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

// TransportHeader The shipping identifier, information about whether the shipment is by an Amazon-partnered carrier, and information about whether the shipment is Small Parcel or Less Than Truckload/Full Truckload (LTL/FTL).
//
// swagger:model TransportHeader
type TransportHeader struct {

	// Indicates whether a putTransportDetails request is for a partnered carrier.
	//
	// Possible values:
	//
	// * true – Request is for an Amazon-partnered carrier.
	//
	// * false – Request is for a non-Amazon-partnered carrier.
	// Required: true
	IsPartnered *bool `json:"IsPartnered"`

	// The Amazon seller identifier.
	// Required: true
	SellerID *string `json:"SellerId"`

	// A shipment identifier originally returned by the createInboundShipmentPlan operation.
	// Required: true
	ShipmentID *string `json:"ShipmentId"`

	// shipment type
	// Required: true
	ShipmentType *ShipmentType `json:"ShipmentType"`
}

// Validate validates this transport header
func (m *TransportHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsPartnered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportHeader) validateIsPartnered(formats strfmt.Registry) error {

	if err := validate.Required("IsPartnered", "body", m.IsPartnered); err != nil {
		return err
	}

	return nil
}

func (m *TransportHeader) validateSellerID(formats strfmt.Registry) error {

	if err := validate.Required("SellerId", "body", m.SellerID); err != nil {
		return err
	}

	return nil
}

func (m *TransportHeader) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("ShipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	return nil
}

func (m *TransportHeader) validateShipmentType(formats strfmt.Registry) error {

	if err := validate.Required("ShipmentType", "body", m.ShipmentType); err != nil {
		return err
	}

	if err := validate.Required("ShipmentType", "body", m.ShipmentType); err != nil {
		return err
	}

	if m.ShipmentType != nil {
		if err := m.ShipmentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipmentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipmentType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transport header based on the context it is used
func (m *TransportHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShipmentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportHeader) contextValidateShipmentType(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentType != nil {
		if err := m.ShipmentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipmentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipmentType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransportHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportHeader) UnmarshalBinary(b []byte) error {
	var res TransportHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
