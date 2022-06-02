// Code generated by go-swagger; DO NOT EDIT.

package vendor_invoices_models

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

// ItemQuantity Details of quantity.
//
// swagger:model ItemQuantity
type ItemQuantity struct {

	// Quantity of an item. This value should not be zero.
	// Required: true
	Amount *int64 `json:"amount"`

	// Unit of measure for the quantity.
	// Required: true
	// Enum: [Cases Eaches]
	UnitOfMeasure *string `json:"unitOfMeasure"`

	// The case size, if the unit of measure value is Cases.
	UnitSize int64 `json:"unitSize,omitempty"`
}

// Validate validates this item quantity
func (m *ItemQuantity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemQuantity) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

var itemQuantityTypeUnitOfMeasurePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cases","Eaches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemQuantityTypeUnitOfMeasurePropEnum = append(itemQuantityTypeUnitOfMeasurePropEnum, v)
	}
}

const (

	// ItemQuantityUnitOfMeasureCases captures enum value "Cases"
	ItemQuantityUnitOfMeasureCases string = "Cases"

	// ItemQuantityUnitOfMeasureEaches captures enum value "Eaches"
	ItemQuantityUnitOfMeasureEaches string = "Eaches"
)

// prop value enum
func (m *ItemQuantity) validateUnitOfMeasureEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemQuantityTypeUnitOfMeasurePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemQuantity) validateUnitOfMeasure(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasure", "body", m.UnitOfMeasure); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitOfMeasureEnum("unitOfMeasure", "body", *m.UnitOfMeasure); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this item quantity based on context it is used
func (m *ItemQuantity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemQuantity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemQuantity) UnmarshalBinary(b []byte) error {
	var res ItemQuantity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}