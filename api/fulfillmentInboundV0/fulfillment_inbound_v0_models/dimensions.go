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

// Dimensions The dimension values and unit of measurement.
//
// swagger:model Dimensions
type Dimensions struct {

	// The height dimension.
	// Required: true
	Height *BigDecimalType `json:"Height"`

	// The length dimension.
	// Required: true
	Length *BigDecimalType `json:"Length"`

	// The unit of measurement for the dimensions.
	// Required: true
	Unit *UnitOfMeasurement `json:"Unit"`

	// The width dimension.
	// Required: true
	Width *BigDecimalType `json:"Width"`
}

// Validate validates this dimensions
func (m *Dimensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dimensions) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("Height", "body", m.Height); err != nil {
		return err
	}

	if err := validate.Required("Height", "body", m.Height); err != nil {
		return err
	}

	if m.Height != nil {
		if err := m.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Height")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) validateLength(formats strfmt.Registry) error {

	if err := validate.Required("Length", "body", m.Length); err != nil {
		return err
	}

	if err := validate.Required("Length", "body", m.Length); err != nil {
		return err
	}

	if m.Length != nil {
		if err := m.Length.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Length")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Length")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("Unit", "body", m.Unit); err != nil {
		return err
	}

	if err := validate.Required("Unit", "body", m.Unit); err != nil {
		return err
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Unit")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) validateWidth(formats strfmt.Registry) error {

	if err := validate.Required("Width", "body", m.Width); err != nil {
		return err
	}

	if err := validate.Required("Width", "body", m.Width); err != nil {
		return err
	}

	if m.Width != nil {
		if err := m.Width.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Width")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dimensions based on the context it is used
func (m *Dimensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLength(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dimensions) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Height != nil {
		if err := m.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Height")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) contextValidateLength(ctx context.Context, formats strfmt.Registry) error {

	if m.Length != nil {
		if err := m.Length.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Length")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Length")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Unit")
			}
			return err
		}
	}

	return nil
}

func (m *Dimensions) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if m.Width != nil {
		if err := m.Width.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Width")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Dimensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dimensions) UnmarshalBinary(b []byte) error {
	var res Dimensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
