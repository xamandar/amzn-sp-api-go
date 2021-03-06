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

// TransportResult The workflow status for a shipment with an Amazon-partnered carrier.
//
// swagger:model TransportResult
type TransportResult struct {

	// An error code that identifies the type of error that occured.
	ErrorCode string `json:"ErrorCode,omitempty"`

	// A message that describes the error condition.
	ErrorDescription string `json:"ErrorDescription,omitempty"`

	// transport status
	// Required: true
	TransportStatus *TransportStatus `json:"TransportStatus"`
}

// Validate validates this transport result
func (m *TransportResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransportStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportResult) validateTransportStatus(formats strfmt.Registry) error {

	if err := validate.Required("TransportStatus", "body", m.TransportStatus); err != nil {
		return err
	}

	if err := validate.Required("TransportStatus", "body", m.TransportStatus); err != nil {
		return err
	}

	if m.TransportStatus != nil {
		if err := m.TransportStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transport result based on the context it is used
func (m *TransportResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransportStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportResult) contextValidateTransportStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.TransportStatus != nil {
		if err := m.TransportStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransportResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportResult) UnmarshalBinary(b []byte) error {
	var res TransportResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
