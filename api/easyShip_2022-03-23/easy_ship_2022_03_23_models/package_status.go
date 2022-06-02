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

// PackageStatus The status of the package.
//
// swagger:model PackageStatus
type PackageStatus string

func NewPackageStatus(value PackageStatus) *PackageStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PackageStatus.
func (m PackageStatus) Pointer() *PackageStatus {
	return &m
}

const (

	// PackageStatusReadyForPickup captures enum value "ReadyForPickup"
	PackageStatusReadyForPickup PackageStatus = "ReadyForPickup"

	// PackageStatusPickedUp captures enum value "PickedUp"
	PackageStatusPickedUp PackageStatus = "PickedUp"

	// PackageStatusAtOriginFC captures enum value "AtOriginFC"
	PackageStatusAtOriginFC PackageStatus = "AtOriginFC"

	// PackageStatusAtDestinationFC captures enum value "AtDestinationFC"
	PackageStatusAtDestinationFC PackageStatus = "AtDestinationFC"

	// PackageStatusDelivered captures enum value "Delivered"
	PackageStatusDelivered PackageStatus = "Delivered"

	// PackageStatusRejected captures enum value "Rejected"
	PackageStatusRejected PackageStatus = "Rejected"

	// PackageStatusUndeliverable captures enum value "Undeliverable"
	PackageStatusUndeliverable PackageStatus = "Undeliverable"

	// PackageStatusReturnedToSeller captures enum value "ReturnedToSeller"
	PackageStatusReturnedToSeller PackageStatus = "ReturnedToSeller"

	// PackageStatusLostInTransit captures enum value "LostInTransit"
	PackageStatusLostInTransit PackageStatus = "LostInTransit"

	// PackageStatusLabelCanceled captures enum value "LabelCanceled"
	PackageStatusLabelCanceled PackageStatus = "LabelCanceled"

	// PackageStatusDamagedInTransit captures enum value "DamagedInTransit"
	PackageStatusDamagedInTransit PackageStatus = "DamagedInTransit"

	// PackageStatusOutForDelivery captures enum value "OutForDelivery"
	PackageStatusOutForDelivery PackageStatus = "OutForDelivery"
)

// for schema
var packageStatusEnum []interface{}

func init() {
	var res []PackageStatus
	if err := json.Unmarshal([]byte(`["ReadyForPickup","PickedUp","AtOriginFC","AtDestinationFC","Delivered","Rejected","Undeliverable","ReturnedToSeller","LostInTransit","LabelCanceled","DamagedInTransit","OutForDelivery"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packageStatusEnum = append(packageStatusEnum, v)
	}
}

func (m PackageStatus) validatePackageStatusEnum(path, location string, value PackageStatus) error {
	if err := validate.EnumCase(path, location, value, packageStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this package status
func (m PackageStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePackageStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this package status based on context it is used
func (m PackageStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}