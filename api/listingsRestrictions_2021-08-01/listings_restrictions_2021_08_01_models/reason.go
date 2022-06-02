// Code generated by go-swagger; DO NOT EDIT.

package listings_restrictions_2021_08_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Reason A reason for the restriction, including path forward links that may allow Selling Partners to remove the restriction, if available.
//
// swagger:model Reason
type Reason struct {

	// A list of path forward links that may allow Selling Partners to remove the restriction.
	Links []*Link `json:"links"`

	// A message describing the reason for the restriction.
	// Required: true
	Message *string `json:"message"`

	// A code indicating why the listing is restricted.
	// Enum: [APPROVAL_REQUIRED ASIN_NOT_FOUND NOT_ELIGIBLE]
	ReasonCode string `json:"reasonCode,omitempty"`
}

// Validate validates this reason
func (m *Reason) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasonCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Reason) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Reason) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

var reasonTypeReasonCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPROVAL_REQUIRED","ASIN_NOT_FOUND","NOT_ELIGIBLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reasonTypeReasonCodePropEnum = append(reasonTypeReasonCodePropEnum, v)
	}
}

const (

	// ReasonReasonCodeAPPROVALREQUIRED captures enum value "APPROVAL_REQUIRED"
	ReasonReasonCodeAPPROVALREQUIRED string = "APPROVAL_REQUIRED"

	// ReasonReasonCodeASINNOTFOUND captures enum value "ASIN_NOT_FOUND"
	ReasonReasonCodeASINNOTFOUND string = "ASIN_NOT_FOUND"

	// ReasonReasonCodeNOTELIGIBLE captures enum value "NOT_ELIGIBLE"
	ReasonReasonCodeNOTELIGIBLE string = "NOT_ELIGIBLE"
)

// prop value enum
func (m *Reason) validateReasonCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reasonTypeReasonCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Reason) validateReasonCode(formats strfmt.Registry) error {
	if swag.IsZero(m.ReasonCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonCodeEnum("reasonCode", "body", m.ReasonCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this reason based on the context it is used
func (m *Reason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Reason) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Reason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reason) UnmarshalBinary(b []byte) error {
	var res Reason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}