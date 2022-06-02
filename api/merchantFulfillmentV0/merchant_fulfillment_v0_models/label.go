// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Label Data for creating a shipping label and dimensions for printing the label.
//
// swagger:model Label
type Label struct {

	// custom text for label
	CustomTextForLabel CustomTextForLabel `json:"CustomTextForLabel,omitempty"`

	// dimensions
	// Required: true
	Dimensions *LabelDimensions `json:"Dimensions"`

	// file contents
	// Required: true
	FileContents *FileContents `json:"FileContents"`

	// label format
	LabelFormat LabelFormat `json:"LabelFormat,omitempty"`

	// standard Id for label
	StandardIDForLabel StandardIDForLabel `json:"StandardIdForLabel,omitempty"`
}

// Validate validates this label
func (m *Label) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomTextForLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardIDForLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Label) validateCustomTextForLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomTextForLabel) { // not required
		return nil
	}

	if err := m.CustomTextForLabel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CustomTextForLabel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CustomTextForLabel")
		}
		return err
	}

	return nil
}

func (m *Label) validateDimensions(formats strfmt.Registry) error {

	if err := validate.Required("Dimensions", "body", m.Dimensions); err != nil {
		return err
	}

	if m.Dimensions != nil {
		if err := m.Dimensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *Label) validateFileContents(formats strfmt.Registry) error {

	if err := validate.Required("FileContents", "body", m.FileContents); err != nil {
		return err
	}

	if m.FileContents != nil {
		if err := m.FileContents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FileContents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FileContents")
			}
			return err
		}
	}

	return nil
}

func (m *Label) validateLabelFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelFormat) { // not required
		return nil
	}

	if err := m.LabelFormat.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LabelFormat")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LabelFormat")
		}
		return err
	}

	return nil
}

func (m *Label) validateStandardIDForLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardIDForLabel) { // not required
		return nil
	}

	if err := m.StandardIDForLabel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandardIdForLabel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StandardIdForLabel")
		}
		return err
	}

	return nil
}

// ContextValidate validate this label based on the context it is used
func (m *Label) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomTextForLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileContents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardIDForLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Label) contextValidateCustomTextForLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CustomTextForLabel.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CustomTextForLabel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CustomTextForLabel")
		}
		return err
	}

	return nil
}

func (m *Label) contextValidateDimensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Dimensions != nil {
		if err := m.Dimensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *Label) contextValidateFileContents(ctx context.Context, formats strfmt.Registry) error {

	if m.FileContents != nil {
		if err := m.FileContents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FileContents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FileContents")
			}
			return err
		}
	}

	return nil
}

func (m *Label) contextValidateLabelFormat(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LabelFormat.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LabelFormat")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LabelFormat")
		}
		return err
	}

	return nil
}

func (m *Label) contextValidateStandardIDForLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StandardIDForLabel.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandardIdForLabel")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StandardIdForLabel")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Label) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Label) UnmarshalBinary(b []byte) error {
	var res Label
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}