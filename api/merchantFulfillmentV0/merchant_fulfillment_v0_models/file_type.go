// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FileType The file type for a label.
//
// swagger:model FileType
type FileType string

func NewFileType(value FileType) *FileType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FileType.
func (m FileType) Pointer() *FileType {
	return &m
}

const (

	// FileTypeApplicationPdf captures enum value "application/pdf"
	FileTypeApplicationPdf FileType = "application/pdf"

	// FileTypeApplicationZpl captures enum value "application/zpl"
	FileTypeApplicationZpl FileType = "application/zpl"

	// FileTypeImagePng captures enum value "image/png"
	FileTypeImagePng FileType = "image/png"
)

// for schema
var fileTypeEnum []interface{}

func init() {
	var res []FileType
	if err := json.Unmarshal([]byte(`["application/pdf","application/zpl","image/png"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileTypeEnum = append(fileTypeEnum, v)
	}
}

func (m FileType) validateFileTypeEnum(path, location string, value FileType) error {
	if err := validate.EnumCase(path, location, value, fileTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this file type
func (m FileType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFileTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this file type based on context it is used
func (m FileType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}