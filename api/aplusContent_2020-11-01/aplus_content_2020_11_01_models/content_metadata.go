// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentMetadata The metadata of an A+ Content document.
//
// swagger:model ContentMetadata
type ContentMetadata struct {

	// badge set
	// Required: true
	BadgeSet ContentBadgeSet `json:"badgeSet"`

	// marketplace Id
	// Required: true
	MarketplaceID *MarketplaceID `json:"marketplaceId"`

	// The A+ Content document name.
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// status
	// Required: true
	Status *ContentStatus `json:"status"`

	// The approximate age of the A+ Content document and metadata.
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"updateTime"`
}

// Validate validates this content metadata
func (m *ContentMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBadgeSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentMetadata) validateBadgeSet(formats strfmt.Registry) error {

	if err := validate.Required("badgeSet", "body", m.BadgeSet); err != nil {
		return err
	}

	if err := m.BadgeSet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("badgeSet")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("badgeSet")
		}
		return err
	}

	return nil
}

func (m *ContentMetadata) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketplaceId")
			}
			return err
		}
	}

	return nil
}

func (m *ContentMetadata) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *ContentMetadata) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ContentMetadata) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("updateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("updateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this content metadata based on the context it is used
func (m *ContentMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBadgeSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarketplaceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentMetadata) contextValidateBadgeSet(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BadgeSet.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("badgeSet")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("badgeSet")
		}
		return err
	}

	return nil
}

func (m *ContentMetadata) contextValidateMarketplaceID(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketplaceId")
			}
			return err
		}
	}

	return nil
}

func (m *ContentMetadata) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentMetadata) UnmarshalBinary(b []byte) error {
	var res ContentMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}