// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentMetadataRecord The metadata for an A+ Content document, with additional information for content management.
//
// swagger:model ContentMetadataRecord
type ContentMetadataRecord struct {

	// content metadata
	// Required: true
	ContentMetadata *ContentMetadata `json:"contentMetadata"`

	// content reference key
	// Required: true
	ContentReferenceKey *ContentReferenceKey `json:"contentReferenceKey"`
}

// Validate validates this content metadata record
func (m *ContentMetadataRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentReferenceKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentMetadataRecord) validateContentMetadata(formats strfmt.Registry) error {

	if err := validate.Required("contentMetadata", "body", m.ContentMetadata); err != nil {
		return err
	}

	if m.ContentMetadata != nil {
		if err := m.ContentMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *ContentMetadataRecord) validateContentReferenceKey(formats strfmt.Registry) error {

	if err := validate.Required("contentReferenceKey", "body", m.ContentReferenceKey); err != nil {
		return err
	}

	if err := validate.Required("contentReferenceKey", "body", m.ContentReferenceKey); err != nil {
		return err
	}

	if m.ContentReferenceKey != nil {
		if err := m.ContentReferenceKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentReferenceKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentReferenceKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this content metadata record based on the context it is used
func (m *ContentMetadataRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentReferenceKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentMetadataRecord) contextValidateContentMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentMetadata != nil {
		if err := m.ContentMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *ContentMetadataRecord) contextValidateContentReferenceKey(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentReferenceKey != nil {
		if err := m.ContentReferenceKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentReferenceKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentReferenceKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentMetadataRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentMetadataRecord) UnmarshalBinary(b []byte) error {
	var res ContentMetadataRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
