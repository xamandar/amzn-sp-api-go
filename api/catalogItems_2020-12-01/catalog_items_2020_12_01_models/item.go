// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2020_12_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Item An item in the Amazon catalog.
//
// swagger:model Item
type Item struct {

	// asin
	// Required: true
	Asin *ItemAsin `json:"asin"`

	// attributes
	Attributes ItemAttributes `json:"attributes,omitempty"`

	// identifiers
	Identifiers ItemIdentifiers `json:"identifiers,omitempty"`

	// images
	Images ItemImages `json:"images,omitempty"`

	// product types
	ProductTypes ItemProductTypes `json:"productTypes,omitempty"`

	// sales ranks
	SalesRanks ItemSalesRanks `json:"salesRanks,omitempty"`

	// summaries
	Summaries ItemSummaries `json:"summaries,omitempty"`

	// variations
	Variations ItemVariations `json:"variations,omitempty"`

	// vendor details
	VendorDetails ItemVendorDetails `json:"vendorDetails,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalesRanks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateAsin(formats strfmt.Registry) error {

	if err := validate.Required("asin", "body", m.Asin); err != nil {
		return err
	}

	if err := validate.Required("asin", "body", m.Asin); err != nil {
		return err
	}

	if m.Asin != nil {
		if err := m.Asin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asin")
			}
			return err
		}
	}

	return nil
}

func (m *Item) validateIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	if err := m.Identifiers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identifiers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identifiers")
		}
		return err
	}

	return nil
}

func (m *Item) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if err := m.Images.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("images")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("images")
		}
		return err
	}

	return nil
}

func (m *Item) validateProductTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductTypes) { // not required
		return nil
	}

	if err := m.ProductTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("productTypes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("productTypes")
		}
		return err
	}

	return nil
}

func (m *Item) validateSalesRanks(formats strfmt.Registry) error {
	if swag.IsZero(m.SalesRanks) { // not required
		return nil
	}

	if err := m.SalesRanks.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("salesRanks")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("salesRanks")
		}
		return err
	}

	return nil
}

func (m *Item) validateSummaries(formats strfmt.Registry) error {
	if swag.IsZero(m.Summaries) { // not required
		return nil
	}

	if err := m.Summaries.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("summaries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("summaries")
		}
		return err
	}

	return nil
}

func (m *Item) validateVariations(formats strfmt.Registry) error {
	if swag.IsZero(m.Variations) { // not required
		return nil
	}

	if err := m.Variations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("variations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("variations")
		}
		return err
	}

	return nil
}

func (m *Item) validateVendorDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.VendorDetails) { // not required
		return nil
	}

	if err := m.VendorDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vendorDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vendorDetails")
		}
		return err
	}

	return nil
}

// ContextValidate validate this item based on the context it is used
func (m *Item) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSalesRanks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVendorDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) contextValidateAsin(ctx context.Context, formats strfmt.Registry) error {

	if m.Asin != nil {
		if err := m.Asin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("asin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("asin")
			}
			return err
		}
	}

	return nil
}

func (m *Item) contextValidateIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Identifiers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identifiers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identifiers")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Images.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("images")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("images")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateProductTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ProductTypes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("productTypes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("productTypes")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateSalesRanks(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SalesRanks.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("salesRanks")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("salesRanks")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateSummaries(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Summaries.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("summaries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("summaries")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateVariations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Variations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("variations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("variations")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateVendorDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VendorDetails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vendorDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vendorDetails")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
