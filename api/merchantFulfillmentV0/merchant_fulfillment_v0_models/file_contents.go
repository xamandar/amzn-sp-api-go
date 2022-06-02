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

// FileContents The document data and checksum.
//
// swagger:model FileContents
type FileContents struct {

	// An MD5 hash to validate the PDF document data, in the form of a Base64-encoded string.
	// Required: true
	Checksum *string `json:"Checksum"`

	// Data for printing labels, in the form of a Base64-encoded, GZip-compressed string.
	// Required: true
	Contents *string `json:"Contents"`

	// file type
	// Required: true
	FileType *FileType `json:"FileType"`
}

// Validate validates this file contents
func (m *FileContents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecksum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileContents) validateChecksum(formats strfmt.Registry) error {

	if err := validate.Required("Checksum", "body", m.Checksum); err != nil {
		return err
	}

	return nil
}

func (m *FileContents) validateContents(formats strfmt.Registry) error {

	if err := validate.Required("Contents", "body", m.Contents); err != nil {
		return err
	}

	return nil
}

func (m *FileContents) validateFileType(formats strfmt.Registry) error {

	if err := validate.Required("FileType", "body", m.FileType); err != nil {
		return err
	}

	if err := validate.Required("FileType", "body", m.FileType); err != nil {
		return err
	}

	if m.FileType != nil {
		if err := m.FileType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FileType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file contents based on the context it is used
func (m *FileContents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileContents) contextValidateFileType(ctx context.Context, formats strfmt.Registry) error {

	if m.FileType != nil {
		if err := m.FileType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FileType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileContents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileContents) UnmarshalBinary(b []byte) error {
	var res FileContents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}