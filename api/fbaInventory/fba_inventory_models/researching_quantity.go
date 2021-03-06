// Code generated by go-swagger; DO NOT EDIT.

package fba_inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResearchingQuantity The number of misplaced or warehouse damaged units that are actively being confirmed at our fulfillment centers.
//
// swagger:model ResearchingQuantity
type ResearchingQuantity struct {

	// A list of quantity details for items currently being researched.
	ResearchingQuantityBreakdown []*ResearchingQuantityEntry `json:"researchingQuantityBreakdown"`

	// The total number of units currently being researched in Amazon's fulfillment network.
	TotalResearchingQuantity int64 `json:"totalResearchingQuantity,omitempty"`
}

// Validate validates this researching quantity
func (m *ResearchingQuantity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResearchingQuantityBreakdown(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResearchingQuantity) validateResearchingQuantityBreakdown(formats strfmt.Registry) error {
	if swag.IsZero(m.ResearchingQuantityBreakdown) { // not required
		return nil
	}

	for i := 0; i < len(m.ResearchingQuantityBreakdown); i++ {
		if swag.IsZero(m.ResearchingQuantityBreakdown[i]) { // not required
			continue
		}

		if m.ResearchingQuantityBreakdown[i] != nil {
			if err := m.ResearchingQuantityBreakdown[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("researchingQuantityBreakdown" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("researchingQuantityBreakdown" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this researching quantity based on the context it is used
func (m *ResearchingQuantity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResearchingQuantityBreakdown(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResearchingQuantity) contextValidateResearchingQuantityBreakdown(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResearchingQuantityBreakdown); i++ {

		if m.ResearchingQuantityBreakdown[i] != nil {
			if err := m.ResearchingQuantityBreakdown[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("researchingQuantityBreakdown" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("researchingQuantityBreakdown" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResearchingQuantity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResearchingQuantity) UnmarshalBinary(b []byte) error {
	var res ResearchingQuantity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
