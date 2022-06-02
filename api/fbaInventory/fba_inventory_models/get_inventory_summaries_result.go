// Code generated by go-swagger; DO NOT EDIT.

package fba_inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetInventorySummariesResult The payload schema for the getInventorySummaries operation.
//
// swagger:model GetInventorySummariesResult
type GetInventorySummariesResult struct {

	// granularity
	// Required: true
	Granularity *Granularity `json:"granularity"`

	// inventory summaries
	// Required: true
	InventorySummaries InventorySummaries `json:"inventorySummaries"`
}

// Validate validates this get inventory summaries result
func (m *GetInventorySummariesResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGranularity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInventorySummaries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInventorySummariesResult) validateGranularity(formats strfmt.Registry) error {

	if err := validate.Required("granularity", "body", m.Granularity); err != nil {
		return err
	}

	if m.Granularity != nil {
		if err := m.Granularity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("granularity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("granularity")
			}
			return err
		}
	}

	return nil
}

func (m *GetInventorySummariesResult) validateInventorySummaries(formats strfmt.Registry) error {

	if err := validate.Required("inventorySummaries", "body", m.InventorySummaries); err != nil {
		return err
	}

	if err := m.InventorySummaries.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("inventorySummaries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("inventorySummaries")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get inventory summaries result based on the context it is used
func (m *GetInventorySummariesResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGranularity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInventorySummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInventorySummariesResult) contextValidateGranularity(ctx context.Context, formats strfmt.Registry) error {

	if m.Granularity != nil {
		if err := m.Granularity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("granularity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("granularity")
			}
			return err
		}
	}

	return nil
}

func (m *GetInventorySummariesResult) contextValidateInventorySummaries(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InventorySummaries.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("inventorySummaries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("inventorySummaries")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetInventorySummariesResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInventorySummariesResult) UnmarshalBinary(b []byte) error {
	var res GetInventorySummariesResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}