// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PrepGuidance Item preparation instructions.
//
// swagger:model PrepGuidance
type PrepGuidance string

func NewPrepGuidance(value PrepGuidance) *PrepGuidance {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PrepGuidance.
func (m PrepGuidance) Pointer() *PrepGuidance {
	return &m
}

const (

	// PrepGuidanceConsultHelpDocuments captures enum value "ConsultHelpDocuments"
	PrepGuidanceConsultHelpDocuments PrepGuidance = "ConsultHelpDocuments"

	// PrepGuidanceNoAdditionalPrepRequired captures enum value "NoAdditionalPrepRequired"
	PrepGuidanceNoAdditionalPrepRequired PrepGuidance = "NoAdditionalPrepRequired"

	// PrepGuidanceSeePrepInstructionsList captures enum value "SeePrepInstructionsList"
	PrepGuidanceSeePrepInstructionsList PrepGuidance = "SeePrepInstructionsList"
)

// for schema
var prepGuidanceEnum []interface{}

func init() {
	var res []PrepGuidance
	if err := json.Unmarshal([]byte(`["ConsultHelpDocuments","NoAdditionalPrepRequired","SeePrepInstructionsList"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prepGuidanceEnum = append(prepGuidanceEnum, v)
	}
}

func (m PrepGuidance) validatePrepGuidanceEnum(path, location string, value PrepGuidance) error {
	if err := validate.EnumCase(path, location, value, prepGuidanceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this prep guidance
func (m PrepGuidance) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePrepGuidanceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this prep guidance based on context it is used
func (m PrepGuidance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
