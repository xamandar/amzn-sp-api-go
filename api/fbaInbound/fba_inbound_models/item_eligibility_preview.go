// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound_models

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

// ItemEligibilityPreview The response object which contains the ASIN, marketplaceId if required, eligibility program, the eligibility status (boolean), and a list of ineligibility reason codes.
//
// swagger:model ItemEligibilityPreview
type ItemEligibilityPreview struct {

	// The ASIN for which eligibility was determined.
	// Required: true
	Asin *string `json:"asin"`

	// Potential Ineligibility Reason Codes.
	IneligibilityReasonList []string `json:"ineligibilityReasonList"`

	// Indicates if the item is eligible for the program.
	// Required: true
	IsEligibleForProgram *bool `json:"isEligibleForProgram"`

	// The marketplace for which eligibility was determined.
	MarketplaceID string `json:"marketplaceId,omitempty"`

	// The program for which eligibility was determined.
	// Required: true
	// Enum: [INBOUND COMMINGLING]
	Program *string `json:"program"`
}

// Validate validates this item eligibility preview
func (m *ItemEligibilityPreview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIneligibilityReasonList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEligibleForProgram(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgram(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemEligibilityPreview) validateAsin(formats strfmt.Registry) error {

	if err := validate.Required("asin", "body", m.Asin); err != nil {
		return err
	}

	return nil
}

var itemEligibilityPreviewIneligibilityReasonListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FBA_INB_0004","FBA_INB_0006","FBA_INB_0007","FBA_INB_0008","FBA_INB_0009","FBA_INB_0010","FBA_INB_0011","FBA_INB_0012","FBA_INB_0013","FBA_INB_0014","FBA_INB_0015","FBA_INB_0016","FBA_INB_0017","FBA_INB_0018","FBA_INB_0019","FBA_INB_0034","FBA_INB_0035","FBA_INB_0036","FBA_INB_0037","FBA_INB_0038","FBA_INB_0050","FBA_INB_0051","FBA_INB_0053","FBA_INB_0055","FBA_INB_0056","FBA_INB_0059","FBA_INB_0065","FBA_INB_0066","FBA_INB_0067","FBA_INB_0068","FBA_INB_0095","FBA_INB_0097","FBA_INB_0098","FBA_INB_0099","FBA_INB_0100","FBA_INB_0103","FBA_INB_0104","FBA_INB_0197","UNKNOWN_INB_ERROR_CODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemEligibilityPreviewIneligibilityReasonListItemsEnum = append(itemEligibilityPreviewIneligibilityReasonListItemsEnum, v)
	}
}

func (m *ItemEligibilityPreview) validateIneligibilityReasonListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemEligibilityPreviewIneligibilityReasonListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemEligibilityPreview) validateIneligibilityReasonList(formats strfmt.Registry) error {
	if swag.IsZero(m.IneligibilityReasonList) { // not required
		return nil
	}

	for i := 0; i < len(m.IneligibilityReasonList); i++ {

		// value enum
		if err := m.validateIneligibilityReasonListItemsEnum("ineligibilityReasonList"+"."+strconv.Itoa(i), "body", m.IneligibilityReasonList[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ItemEligibilityPreview) validateIsEligibleForProgram(formats strfmt.Registry) error {

	if err := validate.Required("isEligibleForProgram", "body", m.IsEligibleForProgram); err != nil {
		return err
	}

	return nil
}

var itemEligibilityPreviewTypeProgramPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INBOUND","COMMINGLING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemEligibilityPreviewTypeProgramPropEnum = append(itemEligibilityPreviewTypeProgramPropEnum, v)
	}
}

const (

	// ItemEligibilityPreviewProgramINBOUND captures enum value "INBOUND"
	ItemEligibilityPreviewProgramINBOUND string = "INBOUND"

	// ItemEligibilityPreviewProgramCOMMINGLING captures enum value "COMMINGLING"
	ItemEligibilityPreviewProgramCOMMINGLING string = "COMMINGLING"
)

// prop value enum
func (m *ItemEligibilityPreview) validateProgramEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemEligibilityPreviewTypeProgramPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemEligibilityPreview) validateProgram(formats strfmt.Registry) error {

	if err := validate.Required("program", "body", m.Program); err != nil {
		return err
	}

	// value enum
	if err := m.validateProgramEnum("program", "body", *m.Program); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this item eligibility preview based on context it is used
func (m *ItemEligibilityPreview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemEligibilityPreview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemEligibilityPreview) UnmarshalBinary(b []byte) error {
	var res ItemEligibilityPreview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
