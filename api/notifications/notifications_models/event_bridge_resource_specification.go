// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventBridgeResourceSpecification The information required to create an Amazon EventBridge destination.
//
// swagger:model EventBridgeResourceSpecification
type EventBridgeResourceSpecification struct {

	// The identifier for the AWS account that is responsible for charges related to receiving notifications.
	// Required: true
	AccountID *string `json:"accountId"`

	// The AWS region in which you will be receiving the notifications.
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this event bridge resource specification
func (m *EventBridgeResourceSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventBridgeResourceSpecification) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *EventBridgeResourceSpecification) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event bridge resource specification based on context it is used
func (m *EventBridgeResourceSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventBridgeResourceSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventBridgeResourceSpecification) UnmarshalBinary(b []byte) error {
	var res EventBridgeResourceSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
