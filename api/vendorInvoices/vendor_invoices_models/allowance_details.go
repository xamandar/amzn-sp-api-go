// Code generated by go-swagger; DO NOT EDIT.

package vendor_invoices_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AllowanceDetails Monetary and tax details of the allowance.
//
// swagger:model AllowanceDetails
type AllowanceDetails struct {

	// Total monetary amount related to this allowance.
	// Required: true
	AllowanceAmount *Money `json:"allowanceAmount"`

	// Description of the allowance.
	Description string `json:"description,omitempty"`

	// Tax amount details applied on this allowance.
	TaxDetails []*TaxDetails `json:"taxDetails"`

	// Type of the allowance applied.
	// Required: true
	// Enum: [Discount DiscountIncentive Defective Promotional UnsaleableMerchandise Special]
	Type *string `json:"type"`
}

// Validate validates this allowance details
func (m *AllowanceDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowanceAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllowanceDetails) validateAllowanceAmount(formats strfmt.Registry) error {

	if err := validate.Required("allowanceAmount", "body", m.AllowanceAmount); err != nil {
		return err
	}

	if m.AllowanceAmount != nil {
		if err := m.AllowanceAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowanceAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allowanceAmount")
			}
			return err
		}
	}

	return nil
}

func (m *AllowanceDetails) validateTaxDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.TaxDetails); i++ {
		if swag.IsZero(m.TaxDetails[i]) { // not required
			continue
		}

		if m.TaxDetails[i] != nil {
			if err := m.TaxDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var allowanceDetailsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Discount","DiscountIncentive","Defective","Promotional","UnsaleableMerchandise","Special"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		allowanceDetailsTypeTypePropEnum = append(allowanceDetailsTypeTypePropEnum, v)
	}
}

const (

	// AllowanceDetailsTypeDiscount captures enum value "Discount"
	AllowanceDetailsTypeDiscount string = "Discount"

	// AllowanceDetailsTypeDiscountIncentive captures enum value "DiscountIncentive"
	AllowanceDetailsTypeDiscountIncentive string = "DiscountIncentive"

	// AllowanceDetailsTypeDefective captures enum value "Defective"
	AllowanceDetailsTypeDefective string = "Defective"

	// AllowanceDetailsTypePromotional captures enum value "Promotional"
	AllowanceDetailsTypePromotional string = "Promotional"

	// AllowanceDetailsTypeUnsaleableMerchandise captures enum value "UnsaleableMerchandise"
	AllowanceDetailsTypeUnsaleableMerchandise string = "UnsaleableMerchandise"

	// AllowanceDetailsTypeSpecial captures enum value "Special"
	AllowanceDetailsTypeSpecial string = "Special"
)

// prop value enum
func (m *AllowanceDetails) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, allowanceDetailsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AllowanceDetails) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this allowance details based on the context it is used
func (m *AllowanceDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowanceAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllowanceDetails) contextValidateAllowanceAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.AllowanceAmount != nil {
		if err := m.AllowanceAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowanceAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allowanceAmount")
			}
			return err
		}
	}

	return nil
}

func (m *AllowanceDetails) contextValidateTaxDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaxDetails); i++ {

		if m.TaxDetails[i] != nil {
			if err := m.TaxDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllowanceDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowanceDetails) UnmarshalBinary(b []byte) error {
	var res AllowanceDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}