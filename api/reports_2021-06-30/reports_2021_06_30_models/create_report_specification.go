// Code generated by go-swagger; DO NOT EDIT.

package reports_2021_06_30_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateReportSpecification Information required to create the report.
//
// swagger:model CreateReportSpecification
type CreateReportSpecification struct {

	// The end of a date and time range, in ISO 8601 date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this.
	// Format: date-time
	DataEndTime strfmt.DateTime `json:"dataEndTime,omitempty"`

	// The start of a date and time range, in ISO 8601 date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this.
	// Format: date-time
	DataStartTime strfmt.DateTime `json:"dataStartTime,omitempty"`

	// A list of marketplace identifiers. The report document's contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise.
	// Required: true
	// Max Items: 25
	// Min Items: 1
	MarketplaceIds []string `json:"marketplaceIds"`

	// report options
	ReportOptions ReportOptions `json:"reportOptions,omitempty"`

	// The report type.
	// Required: true
	ReportType *string `json:"reportType"`
}

// Validate validates this create report specification
func (m *CreateReportSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateReportSpecification) validateDataEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DataEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dataEndTime", "body", "date-time", m.DataEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateReportSpecification) validateDataStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DataStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dataStartTime", "body", "date-time", m.DataStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateReportSpecification) validateMarketplaceIds(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceIds", "body", m.MarketplaceIds); err != nil {
		return err
	}

	iMarketplaceIdsSize := int64(len(m.MarketplaceIds))

	if err := validate.MinItems("marketplaceIds", "body", iMarketplaceIdsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("marketplaceIds", "body", iMarketplaceIdsSize, 25); err != nil {
		return err
	}

	return nil
}

func (m *CreateReportSpecification) validateReportOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ReportOptions) { // not required
		return nil
	}

	if m.ReportOptions != nil {
		if err := m.ReportOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reportOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reportOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateReportSpecification) validateReportType(formats strfmt.Registry) error {

	if err := validate.Required("reportType", "body", m.ReportType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create report specification based on the context it is used
func (m *CreateReportSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReportOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateReportSpecification) contextValidateReportOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ReportOptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reportOptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reportOptions")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateReportSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateReportSpecification) UnmarshalBinary(b []byte) error {
	var res CreateReportSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}