// Code generated by go-swagger; DO NOT EDIT.

package listings_items_2021_08_01_models

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

// ListingsItemSubmissionResponse Response containing the results of a submission to the Selling Partner API for Listings Items.
//
// swagger:model ListingsItemSubmissionResponse
type ListingsItemSubmissionResponse struct {

	// Listings item issues related to the listings item submission.
	Issues []*Issue `json:"issues"`

	// A selling partner provided identifier for an Amazon listing.
	// Required: true
	Sku *string `json:"sku"`

	// The status of the listings item submission.
	// Required: true
	// Enum: [ACCEPTED INVALID]
	Status *string `json:"status"`

	// The unique identifier of the listings item submission.
	// Required: true
	SubmissionID *string `json:"submissionId"`
}

// Validate validates this listings item submission response
func (m *ListingsItemSubmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListingsItemSubmissionResponse) validateIssues(formats strfmt.Registry) error {
	if swag.IsZero(m.Issues) { // not required
		return nil
	}

	for i := 0; i < len(m.Issues); i++ {
		if swag.IsZero(m.Issues[i]) { // not required
			continue
		}

		if m.Issues[i] != nil {
			if err := m.Issues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("issues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("issues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListingsItemSubmissionResponse) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}

var listingsItemSubmissionResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCEPTED","INVALID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listingsItemSubmissionResponseTypeStatusPropEnum = append(listingsItemSubmissionResponseTypeStatusPropEnum, v)
	}
}

const (

	// ListingsItemSubmissionResponseStatusACCEPTED captures enum value "ACCEPTED"
	ListingsItemSubmissionResponseStatusACCEPTED string = "ACCEPTED"

	// ListingsItemSubmissionResponseStatusINVALID captures enum value "INVALID"
	ListingsItemSubmissionResponseStatusINVALID string = "INVALID"
)

// prop value enum
func (m *ListingsItemSubmissionResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listingsItemSubmissionResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ListingsItemSubmissionResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ListingsItemSubmissionResponse) validateSubmissionID(formats strfmt.Registry) error {

	if err := validate.Required("submissionId", "body", m.SubmissionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this listings item submission response based on the context it is used
func (m *ListingsItemSubmissionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIssues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListingsItemSubmissionResponse) contextValidateIssues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Issues); i++ {

		if m.Issues[i] != nil {
			if err := m.Issues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("issues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("issues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListingsItemSubmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListingsItemSubmissionResponse) UnmarshalBinary(b []byte) error {
	var res ListingsItemSubmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
