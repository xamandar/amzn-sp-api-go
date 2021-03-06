// Code generated by go-swagger; DO NOT EDIT.

package fba_smalland_light_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SmallAndLightEnrollmentStatus The Small and Light enrollment status of the item.
//
// swagger:model SmallAndLightEnrollmentStatus
type SmallAndLightEnrollmentStatus string

func NewSmallAndLightEnrollmentStatus(value SmallAndLightEnrollmentStatus) *SmallAndLightEnrollmentStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SmallAndLightEnrollmentStatus.
func (m SmallAndLightEnrollmentStatus) Pointer() *SmallAndLightEnrollmentStatus {
	return &m
}

const (

	// SmallAndLightEnrollmentStatusENROLLED captures enum value "ENROLLED"
	SmallAndLightEnrollmentStatusENROLLED SmallAndLightEnrollmentStatus = "ENROLLED"

	// SmallAndLightEnrollmentStatusNOTENROLLED captures enum value "NOT_ENROLLED"
	SmallAndLightEnrollmentStatusNOTENROLLED SmallAndLightEnrollmentStatus = "NOT_ENROLLED"
)

// for schema
var smallAndLightEnrollmentStatusEnum []interface{}

func init() {
	var res []SmallAndLightEnrollmentStatus
	if err := json.Unmarshal([]byte(`["ENROLLED","NOT_ENROLLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smallAndLightEnrollmentStatusEnum = append(smallAndLightEnrollmentStatusEnum, v)
	}
}

func (m SmallAndLightEnrollmentStatus) validateSmallAndLightEnrollmentStatusEnum(path, location string, value SmallAndLightEnrollmentStatus) error {
	if err := validate.EnumCase(path, location, value, smallAndLightEnrollmentStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this small and light enrollment status
func (m SmallAndLightEnrollmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSmallAndLightEnrollmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this small and light enrollment status based on context it is used
func (m SmallAndLightEnrollmentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
