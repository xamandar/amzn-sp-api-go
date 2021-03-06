// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ContentBadge A flag that provides additional information about an A+ Content document.
//
// swagger:model ContentBadge
type ContentBadge string

func NewContentBadge(value ContentBadge) *ContentBadge {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ContentBadge.
func (m ContentBadge) Pointer() *ContentBadge {
	return &m
}

const (

	// ContentBadgeBULK captures enum value "BULK"
	ContentBadgeBULK ContentBadge = "BULK"

	// ContentBadgeGENERATED captures enum value "GENERATED"
	ContentBadgeGENERATED ContentBadge = "GENERATED"

	// ContentBadgeLAUNCHPAD captures enum value "LAUNCHPAD"
	ContentBadgeLAUNCHPAD ContentBadge = "LAUNCHPAD"

	// ContentBadgePREMIUM captures enum value "PREMIUM"
	ContentBadgePREMIUM ContentBadge = "PREMIUM"

	// ContentBadgeSTANDARD captures enum value "STANDARD"
	ContentBadgeSTANDARD ContentBadge = "STANDARD"
)

// for schema
var contentBadgeEnum []interface{}

func init() {
	var res []ContentBadge
	if err := json.Unmarshal([]byte(`["BULK","GENERATED","LAUNCHPAD","PREMIUM","STANDARD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentBadgeEnum = append(contentBadgeEnum, v)
	}
}

func (m ContentBadge) validateContentBadgeEnum(path, location string, value ContentBadge) error {
	if err := validate.EnumCase(path, location, value, contentBadgeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this content badge
func (m ContentBadge) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContentBadgeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this content badge based on context it is used
func (m ContentBadge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
