// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateShipmentRequest The request schema for the createShipment operation.
//
// swagger:model CreateShipmentRequest
type CreateShipmentRequest struct {

	// client reference Id
	// Required: true
	ClientReferenceID *ClientReferenceID `json:"clientReferenceId"`

	// containers
	// Required: true
	Containers ContainerList `json:"containers"`

	// ship from
	// Required: true
	ShipFrom *Address `json:"shipFrom"`

	// ship to
	// Required: true
	ShipTo *Address `json:"shipTo"`
}

// Validate validates this create shipment request
func (m *CreateShipmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipmentRequest) validateClientReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("clientReferenceId", "body", m.ClientReferenceID); err != nil {
		return err
	}

	if err := validate.Required("clientReferenceId", "body", m.ClientReferenceID); err != nil {
		return err
	}

	if m.ClientReferenceID != nil {
		if err := m.ClientReferenceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clientReferenceId")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	if err := m.Containers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("containers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("containers")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentRequest) validateShipFrom(formats strfmt.Registry) error {

	if err := validate.Required("shipFrom", "body", m.ShipFrom); err != nil {
		return err
	}

	if m.ShipFrom != nil {
		if err := m.ShipFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFrom")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) validateShipTo(formats strfmt.Registry) error {

	if err := validate.Required("shipTo", "body", m.ShipTo); err != nil {
		return err
	}

	if m.ShipTo != nil {
		if err := m.ShipTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipTo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create shipment request based on the context it is used
func (m *CreateShipmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientReferenceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipmentRequest) contextValidateClientReferenceID(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientReferenceID != nil {
		if err := m.ClientReferenceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clientReferenceId")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Containers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("containers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("containers")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateShipFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFrom != nil {
		if err := m.ShipFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFrom")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateShipTo(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipTo != nil {
		if err := m.ShipTo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipTo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateShipmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShipmentRequest) UnmarshalBinary(b []byte) error {
	var res CreateShipmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
