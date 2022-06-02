// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostContentDocumentSuspendSubmissionResponse post content document suspend submission response
//
// swagger:model PostContentDocumentSuspendSubmissionResponse
type PostContentDocumentSuspendSubmissionResponse struct {
	AplusResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostContentDocumentSuspendSubmissionResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AplusResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AplusResponse = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostContentDocumentSuspendSubmissionResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AplusResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post content document suspend submission response
func (m *PostContentDocumentSuspendSubmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AplusResponse
	if err := m.AplusResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post content document suspend submission response based on the context it is used
func (m *PostContentDocumentSuspendSubmissionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AplusResponse
	if err := m.AplusResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostContentDocumentSuspendSubmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostContentDocumentSuspendSubmissionResponse) UnmarshalBinary(b []byte) error {
	var res PostContentDocumentSuspendSubmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}