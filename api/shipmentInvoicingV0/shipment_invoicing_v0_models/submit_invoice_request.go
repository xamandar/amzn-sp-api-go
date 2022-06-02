// Code generated by go-swagger; DO NOT EDIT.

package shipment_invoicing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubmitInvoiceRequest The request schema for the submitInvoice operation.
//
// swagger:model SubmitInvoiceRequest
type SubmitInvoiceRequest struct {

	// MD5 sum for validating the invoice data. For more information about calculating this value, see [Working with Content-MD5 Checksums](https://docs.developer.amazonservices.com/en_US/dev_guide/DG_MD5.html).
	// Required: true
	ContentMD5Value *string `json:"ContentMD5Value"`

	// invoice content
	// Required: true
	// Format: byte
	InvoiceContent *Blob `json:"InvoiceContent"`

	// An Amazon marketplace identifier.
	MarketplaceID string `json:"MarketplaceId,omitempty"`
}

// Validate validates this submit invoice request
func (m *SubmitInvoiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentMD5Value(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitInvoiceRequest) validateContentMD5Value(formats strfmt.Registry) error {

	if err := validate.Required("ContentMD5Value", "body", m.ContentMD5Value); err != nil {
		return err
	}

	return nil
}

func (m *SubmitInvoiceRequest) validateInvoiceContent(formats strfmt.Registry) error {

	if err := validate.Required("InvoiceContent", "body", m.InvoiceContent); err != nil {
		return err
	}

	if err := validate.Required("InvoiceContent", "body", m.InvoiceContent); err != nil {
		return err
	}

	if m.InvoiceContent != nil {
		if err := m.InvoiceContent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InvoiceContent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InvoiceContent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this submit invoice request based on the context it is used
func (m *SubmitInvoiceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvoiceContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitInvoiceRequest) contextValidateInvoiceContent(ctx context.Context, formats strfmt.Registry) error {

	if m.InvoiceContent != nil {
		if err := m.InvoiceContent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InvoiceContent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InvoiceContent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubmitInvoiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitInvoiceRequest) UnmarshalBinary(b []byte) error {
	var res SubmitInvoiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}