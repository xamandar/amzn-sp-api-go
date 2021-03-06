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

// SearchContentPublishRecordsResponse search content publish records response
//
// swagger:model SearchContentPublishRecordsResponse
type SearchContentPublishRecordsResponse struct {
	AplusPaginatedResponse

	// publish record list
	// Required: true
	PublishRecordList PublishRecordList `json:"publishRecordList"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SearchContentPublishRecordsResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AplusPaginatedResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AplusPaginatedResponse = aO0

	// AO1
	var dataAO1 struct {
		PublishRecordList PublishRecordList `json:"publishRecordList"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.PublishRecordList = dataAO1.PublishRecordList

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SearchContentPublishRecordsResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AplusPaginatedResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		PublishRecordList PublishRecordList `json:"publishRecordList"`
	}

	dataAO1.PublishRecordList = m.PublishRecordList

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this search content publish records response
func (m *SearchContentPublishRecordsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AplusPaginatedResponse
	if err := m.AplusPaginatedResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishRecordList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchContentPublishRecordsResponse) validatePublishRecordList(formats strfmt.Registry) error {

	if err := validate.Required("publishRecordList", "body", m.PublishRecordList); err != nil {
		return err
	}

	if err := m.PublishRecordList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("publishRecordList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("publishRecordList")
		}
		return err
	}

	return nil
}

// ContextValidate validate this search content publish records response based on the context it is used
func (m *SearchContentPublishRecordsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AplusPaginatedResponse
	if err := m.AplusPaginatedResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublishRecordList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchContentPublishRecordsResponse) contextValidatePublishRecordList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PublishRecordList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("publishRecordList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("publishRecordList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchContentPublishRecordsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchContentPublishRecordsResponse) UnmarshalBinary(b []byte) error {
	var res SearchContentPublishRecordsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
