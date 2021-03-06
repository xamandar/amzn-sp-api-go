// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubmitAcknowledgementRequest The request schema for the submitAcknowledgment operation.
//
// swagger:model SubmitAcknowledgementRequest
type SubmitAcknowledgementRequest struct {

	// acknowledgements
	Acknowledgements []*OrderAcknowledgement `json:"acknowledgements"`
}

// Validate validates this submit acknowledgement request
func (m *SubmitAcknowledgementRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcknowledgements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitAcknowledgementRequest) validateAcknowledgements(formats strfmt.Registry) error {
	if swag.IsZero(m.Acknowledgements) { // not required
		return nil
	}

	for i := 0; i < len(m.Acknowledgements); i++ {
		if swag.IsZero(m.Acknowledgements[i]) { // not required
			continue
		}

		if m.Acknowledgements[i] != nil {
			if err := m.Acknowledgements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acknowledgements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this submit acknowledgement request based on the context it is used
func (m *SubmitAcknowledgementRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcknowledgements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitAcknowledgementRequest) contextValidateAcknowledgements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Acknowledgements); i++ {

		if m.Acknowledgements[i] != nil {
			if err := m.Acknowledgements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acknowledgements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubmitAcknowledgementRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitAcknowledgementRequest) UnmarshalBinary(b []byte) error {
	var res SubmitAcknowledgementRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
