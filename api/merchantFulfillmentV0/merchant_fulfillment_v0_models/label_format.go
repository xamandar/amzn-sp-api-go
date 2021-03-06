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

// LabelFormat The label format.
//
// swagger:model LabelFormat
type LabelFormat string

func NewLabelFormat(value LabelFormat) *LabelFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated LabelFormat.
func (m LabelFormat) Pointer() *LabelFormat {
	return &m
}

const (

	// LabelFormatPDF captures enum value "PDF"
	LabelFormatPDF LabelFormat = "PDF"

	// LabelFormatPNG captures enum value "PNG"
	LabelFormatPNG LabelFormat = "PNG"

	// LabelFormatZPL203 captures enum value "ZPL203"
	LabelFormatZPL203 LabelFormat = "ZPL203"

	// LabelFormatZPL300 captures enum value "ZPL300"
	LabelFormatZPL300 LabelFormat = "ZPL300"

	// LabelFormatShippingServiceDefault captures enum value "ShippingServiceDefault"
	LabelFormatShippingServiceDefault LabelFormat = "ShippingServiceDefault"
)

// for schema
var labelFormatEnum []interface{}

func init() {
	var res []LabelFormat
	if err := json.Unmarshal([]byte(`["PDF","PNG","ZPL203","ZPL300","ShippingServiceDefault"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		labelFormatEnum = append(labelFormatEnum, v)
	}
}

func (m LabelFormat) validateLabelFormatEnum(path, location string, value LabelFormat) error {
	if err := validate.EnumCase(path, location, value, labelFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this label format
func (m LabelFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLabelFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this label format based on context it is used
func (m LabelFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
