// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AcknowledgementStatusDetails Details of item quantity ordered
//
// swagger:model AcknowledgementStatusDetails
type AcknowledgementStatusDetails struct {

	// Item quantity accepted by vendor to be shipped.
	AcceptedQuantity *ItemQuantity `json:"acceptedQuantity,omitempty"`

	// The date when the line item was confirmed by vendor. Must be in ISO-8601 date/time format.
	// Format: date-time
	AcknowledgementDate strfmt.DateTime `json:"acknowledgementDate,omitempty"`

	// Item quantity rejected by vendor.
	RejectedQuantity *ItemQuantity `json:"rejectedQuantity,omitempty"`
}

// Validate validates this acknowledgement status details
func (m *AcknowledgementStatusDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcknowledgementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcknowledgementStatusDetails) validateAcceptedQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptedQuantity) { // not required
		return nil
	}

	if m.AcceptedQuantity != nil {
		if err := m.AcceptedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptedQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *AcknowledgementStatusDetails) validateAcknowledgementDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AcknowledgementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("acknowledgementDate", "body", "date-time", m.AcknowledgementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AcknowledgementStatusDetails) validateRejectedQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectedQuantity) { // not required
		return nil
	}

	if m.RejectedQuantity != nil {
		if err := m.RejectedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rejectedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rejectedQuantity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this acknowledgement status details based on the context it is used
func (m *AcknowledgementStatusDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRejectedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcknowledgementStatusDetails) contextValidateAcceptedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.AcceptedQuantity != nil {
		if err := m.AcceptedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptedQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *AcknowledgementStatusDetails) contextValidateRejectedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.RejectedQuantity != nil {
		if err := m.RejectedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rejectedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rejectedQuantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcknowledgementStatusDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcknowledgementStatusDetails) UnmarshalBinary(b []byte) error {
	var res AcknowledgementStatusDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
