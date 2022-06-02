// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LabelSpecification The label specification info.
//
// swagger:model LabelSpecification
type LabelSpecification struct {

	// The format of the label. Enum of PNG only for now.
	// Required: true
	// Enum: [PNG]
	LabelFormat *string `json:"labelFormat"`

	// The label stock size specification in length and height. Enum of 4x6 only for now.
	// Required: true
	// Enum: [4x6]
	LabelStockSize *string `json:"labelStockSize"`
}

// Validate validates this label specification
func (m *LabelSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelStockSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var labelSpecificationTypeLabelFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PNG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		labelSpecificationTypeLabelFormatPropEnum = append(labelSpecificationTypeLabelFormatPropEnum, v)
	}
}

const (

	// LabelSpecificationLabelFormatPNG captures enum value "PNG"
	LabelSpecificationLabelFormatPNG string = "PNG"
)

// prop value enum
func (m *LabelSpecification) validateLabelFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, labelSpecificationTypeLabelFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LabelSpecification) validateLabelFormat(formats strfmt.Registry) error {

	if err := validate.Required("labelFormat", "body", m.LabelFormat); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelFormatEnum("labelFormat", "body", *m.LabelFormat); err != nil {
		return err
	}

	return nil
}

var labelSpecificationTypeLabelStockSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["4x6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		labelSpecificationTypeLabelStockSizePropEnum = append(labelSpecificationTypeLabelStockSizePropEnum, v)
	}
}

const (

	// LabelSpecificationLabelStockSizeNr4x6 captures enum value "4x6"
	LabelSpecificationLabelStockSizeNr4x6 string = "4x6"
)

// prop value enum
func (m *LabelSpecification) validateLabelStockSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, labelSpecificationTypeLabelStockSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LabelSpecification) validateLabelStockSize(formats strfmt.Registry) error {

	if err := validate.Required("labelStockSize", "body", m.LabelStockSize); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelStockSizeEnum("labelStockSize", "body", *m.LabelStockSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this label specification based on context it is used
func (m *LabelSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelSpecification) UnmarshalBinary(b []byte) error {
	var res LabelSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}