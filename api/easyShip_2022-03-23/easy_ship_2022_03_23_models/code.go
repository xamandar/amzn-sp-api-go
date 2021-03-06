// Code generated by go-swagger; DO NOT EDIT.

package easy_ship_2022_03_23_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Code An error code that identifies the type of error that occurred. The error codes listed below are specific to the Easy Ship section.
//
// swagger:model Code
type Code string

func NewCode(value Code) *Code {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Code.
func (m Code) Pointer() *Code {
	return &m
}

const (

	// CodeInvalidInput captures enum value "InvalidInput"
	CodeInvalidInput Code = "InvalidInput"

	// CodeInvalidTimeSlotID captures enum value "InvalidTimeSlotId"
	CodeInvalidTimeSlotID Code = "InvalidTimeSlotId"

	// CodeScheduledPackageAlreadyExists captures enum value "ScheduledPackageAlreadyExists"
	CodeScheduledPackageAlreadyExists Code = "ScheduledPackageAlreadyExists"

	// CodeScheduleWindowExpired captures enum value "ScheduleWindowExpired"
	CodeScheduleWindowExpired Code = "ScheduleWindowExpired"

	// CodeRetryableAfterGettingNewSlots captures enum value "RetryableAfterGettingNewSlots"
	CodeRetryableAfterGettingNewSlots Code = "RetryableAfterGettingNewSlots"

	// CodeTimeSlotNotAvailable captures enum value "TimeSlotNotAvailable"
	CodeTimeSlotNotAvailable Code = "TimeSlotNotAvailable"

	// CodeResourceNotFound captures enum value "ResourceNotFound"
	CodeResourceNotFound Code = "ResourceNotFound"

	// CodeInvalidOrderState captures enum value "InvalidOrderState"
	CodeInvalidOrderState Code = "InvalidOrderState"

	// CodeRegionNotSupported captures enum value "RegionNotSupported"
	CodeRegionNotSupported Code = "RegionNotSupported"

	// CodeOrderNotEligibleForRescheduling captures enum value "OrderNotEligibleForRescheduling"
	CodeOrderNotEligibleForRescheduling Code = "OrderNotEligibleForRescheduling"

	// CodeInternalServerError captures enum value "InternalServerError"
	CodeInternalServerError Code = "InternalServerError"
)

// for schema
var codeEnum []interface{}

func init() {
	var res []Code
	if err := json.Unmarshal([]byte(`["InvalidInput","InvalidTimeSlotId","ScheduledPackageAlreadyExists","ScheduleWindowExpired","RetryableAfterGettingNewSlots","TimeSlotNotAvailable","ResourceNotFound","InvalidOrderState","RegionNotSupported","OrderNotEligibleForRescheduling","InternalServerError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		codeEnum = append(codeEnum, v)
	}
}

func (m Code) validateCodeEnum(path, location string, value Code) error {
	if err := validate.EnumCase(path, location, value, codeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this code
func (m Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this code based on context it is used
func (m Code) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
