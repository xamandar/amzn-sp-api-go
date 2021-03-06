// Code generated by go-swagger; DO NOT EDIT.

package definitions_product_types_2020_09_01_models

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

// SchemaLink schema link
//
// swagger:model SchemaLink
type SchemaLink struct {

	// Checksum hash of the schema (Base64 MD5). Can be used to verify schema contents, identify changes between schema versions, and for caching.
	// Required: true
	Checksum *string `json:"checksum"`

	// link
	// Required: true
	Link *SchemaLinkLink `json:"link"`
}

// Validate validates this schema link
func (m *SchemaLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecksum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemaLink) validateChecksum(formats strfmt.Registry) error {

	if err := validate.Required("checksum", "body", m.Checksum); err != nil {
		return err
	}

	return nil
}

func (m *SchemaLink) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this schema link based on the context it is used
func (m *SchemaLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemaLink) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {
		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemaLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemaLink) UnmarshalBinary(b []byte) error {
	var res SchemaLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SchemaLinkLink Link to retrieve the schema.
//
// swagger:model SchemaLinkLink
type SchemaLinkLink struct {

	// URI resource for the link.
	// Required: true
	Resource *string `json:"resource"`

	// HTTP method for the link operation.
	// Required: true
	// Enum: [GET]
	Verb *string `json:"verb"`
}

// Validate validates this schema link link
func (m *SchemaLinkLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemaLinkLink) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("link"+"."+"resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

var schemaLinkLinkTypeVerbPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemaLinkLinkTypeVerbPropEnum = append(schemaLinkLinkTypeVerbPropEnum, v)
	}
}

const (

	// SchemaLinkLinkVerbGET captures enum value "GET"
	SchemaLinkLinkVerbGET string = "GET"
)

// prop value enum
func (m *SchemaLinkLink) validateVerbEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, schemaLinkLinkTypeVerbPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SchemaLinkLink) validateVerb(formats strfmt.Registry) error {

	if err := validate.Required("link"+"."+"verb", "body", m.Verb); err != nil {
		return err
	}

	// value enum
	if err := m.validateVerbEnum("link"+"."+"verb", "body", *m.Verb); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schema link link based on context it is used
func (m *SchemaLinkLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchemaLinkLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemaLinkLink) UnmarshalBinary(b []byte) error {
	var res SchemaLinkLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
