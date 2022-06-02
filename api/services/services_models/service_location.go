// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceLocation Information about the location of the service job.
//
// swagger:model ServiceLocation
type ServiceLocation struct {

	// The shipping address for the service job.
	Address *Address `json:"address,omitempty"`

	// The location of the service job.
	// Enum: [IN_HOME IN_STORE ONLINE]
	ServiceLocationType string `json:"serviceLocationType,omitempty"`
}

// Validate validates this service location
func (m *ServiceLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceLocationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceLocation) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

var serviceLocationTypeServiceLocationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_HOME","IN_STORE","ONLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceLocationTypeServiceLocationTypePropEnum = append(serviceLocationTypeServiceLocationTypePropEnum, v)
	}
}

const (

	// ServiceLocationServiceLocationTypeINHOME captures enum value "IN_HOME"
	ServiceLocationServiceLocationTypeINHOME string = "IN_HOME"

	// ServiceLocationServiceLocationTypeINSTORE captures enum value "IN_STORE"
	ServiceLocationServiceLocationTypeINSTORE string = "IN_STORE"

	// ServiceLocationServiceLocationTypeONLINE captures enum value "ONLINE"
	ServiceLocationServiceLocationTypeONLINE string = "ONLINE"
)

// prop value enum
func (m *ServiceLocation) validateServiceLocationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceLocationTypeServiceLocationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceLocation) validateServiceLocationType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceLocationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceLocationTypeEnum("serviceLocationType", "body", m.ServiceLocationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service location based on the context it is used
func (m *ServiceLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceLocation) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceLocation) UnmarshalBinary(b []byte) error {
	var res ServiceLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}