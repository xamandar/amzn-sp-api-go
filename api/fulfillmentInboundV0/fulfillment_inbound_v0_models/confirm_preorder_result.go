// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfirmPreorderResult confirm preorder result
//
// swagger:model ConfirmPreorderResult
type ConfirmPreorderResult struct {

	// Date that determines which pre-order items in the shipment are eligible for pre-order. The pre-order Buy Box will appear for any pre-order item in the shipment with a release date on or after this date. In YYYY-MM-DD format.
	// Format: date
	ConfirmedFulfillableDate DateStringType `json:"ConfirmedFulfillableDate,omitempty"`

	// Date passed in with the NeedByDate parameter. The confirmed shipment must arrive at the Amazon fulfillment center by this date to avoid delivery promise breaks for pre-ordered items. In YYYY-MM-DD format.
	// Format: date
	ConfirmedNeedByDate DateStringType `json:"ConfirmedNeedByDate,omitempty"`
}

// Validate validates this confirm preorder result
func (m *ConfirmPreorderResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfirmedFulfillableDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfirmedNeedByDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfirmPreorderResult) validateConfirmedFulfillableDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfirmedFulfillableDate) { // not required
		return nil
	}

	if err := m.ConfirmedFulfillableDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ConfirmedFulfillableDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ConfirmedFulfillableDate")
		}
		return err
	}

	return nil
}

func (m *ConfirmPreorderResult) validateConfirmedNeedByDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfirmedNeedByDate) { // not required
		return nil
	}

	if err := m.ConfirmedNeedByDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ConfirmedNeedByDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ConfirmedNeedByDate")
		}
		return err
	}

	return nil
}

// ContextValidate validate this confirm preorder result based on the context it is used
func (m *ConfirmPreorderResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfirmedFulfillableDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfirmedNeedByDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfirmPreorderResult) contextValidateConfirmedFulfillableDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConfirmedFulfillableDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ConfirmedFulfillableDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ConfirmedFulfillableDate")
		}
		return err
	}

	return nil
}

func (m *ConfirmPreorderResult) contextValidateConfirmedNeedByDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConfirmedNeedByDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ConfirmedNeedByDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ConfirmedNeedByDate")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmPreorderResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmPreorderResult) UnmarshalBinary(b []byte) error {
	var res ConfirmPreorderResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}