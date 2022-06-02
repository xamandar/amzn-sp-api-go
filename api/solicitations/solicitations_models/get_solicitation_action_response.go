// Code generated by go-swagger; DO NOT EDIT.

package solicitations_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetSolicitationActionResponse Describes a solicitation action that can be taken for an order. Provides a JSON Hypertext Application Language (HAL) link to the JSON schema document that describes the expected input.
//
// swagger:model GetSolicitationActionResponse
type GetSolicitationActionResponse struct {

	// embedded
	Embedded *GetSolicitationActionResponseEmbedded `json:"_embedded,omitempty"`

	// links
	Links *GetSolicitationActionResponseLinks `json:"_links,omitempty"`

	// errors
	Errors ErrorList `json:"errors,omitempty"`

	// payload
	Payload *SolicitationsAction `json:"payload,omitempty"`
}

// Validate validates this get solicitation action response
func (m *GetSolicitationActionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSolicitationActionResponse) validateEmbedded(formats strfmt.Registry) error {
	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *GetSolicitationActionResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *GetSolicitationActionResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if err := m.Errors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *GetSolicitationActionResponse) validatePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get solicitation action response based on the context it is used
func (m *GetSolicitationActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmbedded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSolicitationActionResponse) contextValidateEmbedded(ctx context.Context, formats strfmt.Registry) error {

	if m.Embedded != nil {
		if err := m.Embedded.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *GetSolicitationActionResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *GetSolicitationActionResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Errors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *GetSolicitationActionResponse) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if m.Payload != nil {
		if err := m.Payload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSolicitationActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSolicitationActionResponse) UnmarshalBinary(b []byte) error {
	var res GetSolicitationActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetSolicitationActionResponseEmbedded get solicitation action response embedded
//
// swagger:model GetSolicitationActionResponseEmbedded
type GetSolicitationActionResponseEmbedded struct {

	// schema
	Schema *GetSchemaResponse `json:"schema,omitempty"`
}

// Validate validates this get solicitation action response embedded
func (m *GetSolicitationActionResponseEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSolicitationActionResponseEmbedded) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get solicitation action response embedded based on the context it is used
func (m *GetSolicitationActionResponseEmbedded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSolicitationActionResponseEmbedded) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSolicitationActionResponseEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSolicitationActionResponseEmbedded) UnmarshalBinary(b []byte) error {
	var res GetSolicitationActionResponseEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetSolicitationActionResponseLinks get solicitation action response links
//
// swagger:model GetSolicitationActionResponseLinks
type GetSolicitationActionResponseLinks struct {

	// schema
	// Required: true
	Schema *LinkObject `json:"schema"`

	// self
	// Required: true
	Self *LinkObject `json:"self"`
}

// Validate validates this get solicitation action response links
func (m *GetSolicitationActionResponseLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSolicitationActionResponseLinks) validateSchema(formats strfmt.Registry) error {

	if err := validate.Required("_links"+"."+"schema", "body", m.Schema); err != nil {
		return err
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

func (m *GetSolicitationActionResponseLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("_links"+"."+"self", "body", m.Self); err != nil {
		return err
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get solicitation action response links based on the context it is used
func (m *GetSolicitationActionResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSolicitationActionResponseLinks) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

func (m *GetSolicitationActionResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSolicitationActionResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSolicitationActionResponseLinks) UnmarshalBinary(b []byte) error {
	var res GetSolicitationActionResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}