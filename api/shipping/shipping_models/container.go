// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

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

// Container Container in the shipment.
//
// swagger:model Container
type Container struct {

	// container reference Id
	// Required: true
	ContainerReferenceID *ContainerReferenceID `json:"containerReferenceId"`

	// The type of physical container being used. (always 'PACKAGE')
	// Enum: [PACKAGE]
	ContainerType string `json:"containerType,omitempty"`

	// The length, width, height, and weight of the container.
	// Required: true
	Dimensions *Dimensions `json:"dimensions"`

	// A list of the items in the container.
	// Required: true
	Items []*ContainerItem `json:"items"`

	// The total value of all items in the container.
	// Required: true
	Value *Currency `json:"value"`

	// The weight of the container.
	// Required: true
	Weight *Weight `json:"weight"`
}

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) validateContainerReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("containerReferenceId", "body", m.ContainerReferenceID); err != nil {
		return err
	}

	if err := validate.Required("containerReferenceId", "body", m.ContainerReferenceID); err != nil {
		return err
	}

	if m.ContainerReferenceID != nil {
		if err := m.ContainerReferenceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerReferenceId")
			}
			return err
		}
	}

	return nil
}

var containerTypeContainerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PACKAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerTypeContainerTypePropEnum = append(containerTypeContainerTypePropEnum, v)
	}
}

const (

	// ContainerContainerTypePACKAGE captures enum value "PACKAGE"
	ContainerContainerTypePACKAGE string = "PACKAGE"
)

// prop value enum
func (m *Container) validateContainerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, containerTypeContainerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Container) validateContainerType(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContainerTypeEnum("containerType", "body", m.ContainerType); err != nil {
		return err
	}

	return nil
}

func (m *Container) validateDimensions(formats strfmt.Registry) error {

	if err := validate.Required("dimensions", "body", m.Dimensions); err != nil {
		return err
	}

	if m.Dimensions != nil {
		if err := m.Dimensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *Container) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Container) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

func (m *Container) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	if m.Weight != nil {
		if err := m.Weight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container based on the context it is used
func (m *Container) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerReferenceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) contextValidateContainerReferenceID(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerReferenceID != nil {
		if err := m.ContainerReferenceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerReferenceId")
			}
			return err
		}
	}

	return nil
}

func (m *Container) contextValidateDimensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Dimensions != nil {
		if err := m.Dimensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *Container) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Container) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

func (m *Container) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Weight != nil {
		if err := m.Weight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}