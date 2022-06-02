// Code generated by go-swagger; DO NOT EDIT.

package reports_2021_06_30_models

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

// ReportDocument Information required for the report document.
//
// swagger:model ReportDocument
type ReportDocument struct {

	// If present, the report document contents have been compressed with the provided algorithm.
	// Enum: [GZIP]
	CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`

	// The identifier for the report document. This identifier is unique only in combination with a seller ID.
	// Required: true
	ReportDocumentID *string `json:"reportDocumentId"`

	// A presigned URL for the report document. This URL expires after 5 minutes.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this report document
func (m *ReportDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompressionAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportDocumentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportDocumentTypeCompressionAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GZIP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportDocumentTypeCompressionAlgorithmPropEnum = append(reportDocumentTypeCompressionAlgorithmPropEnum, v)
	}
}

const (

	// ReportDocumentCompressionAlgorithmGZIP captures enum value "GZIP"
	ReportDocumentCompressionAlgorithmGZIP string = "GZIP"
)

// prop value enum
func (m *ReportDocument) validateCompressionAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportDocumentTypeCompressionAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportDocument) validateCompressionAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(m.CompressionAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompressionAlgorithmEnum("compressionAlgorithm", "body", m.CompressionAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *ReportDocument) validateReportDocumentID(formats strfmt.Registry) error {

	if err := validate.Required("reportDocumentId", "body", m.ReportDocumentID); err != nil {
		return err
	}

	return nil
}

func (m *ReportDocument) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report document based on context it is used
func (m *ReportDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportDocument) UnmarshalBinary(b []byte) error {
	var res ReportDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}