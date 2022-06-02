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

// ItemSummaryByMarketplace Summary details of a listings item for an Amazon marketplace.
//
// swagger:model ItemSummaryByMarketplace
type ItemSummaryByMarketplace struct {

	// Amazon Standard Identification Number (ASIN) of the listings item.
	// Required: true
	Asin *string `json:"asin"`

	// Identifies the condition of the listings item.
	// Enum: [new_new new_open_box new_oem refurbished_refurbished used_like_new used_very_good used_good used_acceptable collectible_like_new collectible_very_good collectible_good collectible_acceptable club_club]
	ConditionType string `json:"conditionType,omitempty"`

	// Date the listings item was created, in ISO 8601 format.
	// Required: true
	// Format: date-time
	CreatedDate *strfmt.DateTime `json:"createdDate"`

	// Fulfillment network stock keeping unit is an identifier used by Amazon fulfillment centers to identify each unique item.
	FnSku string `json:"fnSku,omitempty"`

	// Name, or title, associated with an Amazon catalog item.
	// Required: true
	ItemName *string `json:"itemName"`

	// Date the listings item was last updated, in ISO 8601 format.
	// Required: true
	// Format: date-time
	LastUpdatedDate *strfmt.DateTime `json:"lastUpdatedDate"`

	// Main image for the listings item.
	MainImage *ItemImage `json:"mainImage,omitempty"`

	// A marketplace identifier. Identifies the Amazon marketplace for the listings item.
	// Required: true
	MarketplaceID *string `json:"marketplaceId"`

	// The Amazon product type of the listings item.
	// Required: true
	ProductType *string `json:"productType"`

	// Statuses that apply to the listings item.
	// Required: true
	Status []string `json:"status"`
}

// Validate validates this item summary by marketplace
func (m *ItemSummaryByMarketplace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMainImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSummaryByMarketplace) validateAsin(formats strfmt.Registry) error {

	if err := validate.Required("asin", "body", m.Asin); err != nil {
		return err
	}

	return nil
}

var itemSummaryByMarketplaceTypeConditionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["new_new","new_open_box","new_oem","refurbished_refurbished","used_like_new","used_very_good","used_good","used_acceptable","collectible_like_new","collectible_very_good","collectible_good","collectible_acceptable","club_club"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSummaryByMarketplaceTypeConditionTypePropEnum = append(itemSummaryByMarketplaceTypeConditionTypePropEnum, v)
	}
}

const (

	// ItemSummaryByMarketplaceConditionTypeNewNew captures enum value "new_new"
	ItemSummaryByMarketplaceConditionTypeNewNew string = "new_new"

	// ItemSummaryByMarketplaceConditionTypeNewOpenBox captures enum value "new_open_box"
	ItemSummaryByMarketplaceConditionTypeNewOpenBox string = "new_open_box"

	// ItemSummaryByMarketplaceConditionTypeNewOem captures enum value "new_oem"
	ItemSummaryByMarketplaceConditionTypeNewOem string = "new_oem"

	// ItemSummaryByMarketplaceConditionTypeRefurbishedRefurbished captures enum value "refurbished_refurbished"
	ItemSummaryByMarketplaceConditionTypeRefurbishedRefurbished string = "refurbished_refurbished"

	// ItemSummaryByMarketplaceConditionTypeUsedLikeNew captures enum value "used_like_new"
	ItemSummaryByMarketplaceConditionTypeUsedLikeNew string = "used_like_new"

	// ItemSummaryByMarketplaceConditionTypeUsedVeryGood captures enum value "used_very_good"
	ItemSummaryByMarketplaceConditionTypeUsedVeryGood string = "used_very_good"

	// ItemSummaryByMarketplaceConditionTypeUsedGood captures enum value "used_good"
	ItemSummaryByMarketplaceConditionTypeUsedGood string = "used_good"

	// ItemSummaryByMarketplaceConditionTypeUsedAcceptable captures enum value "used_acceptable"
	ItemSummaryByMarketplaceConditionTypeUsedAcceptable string = "used_acceptable"

	// ItemSummaryByMarketplaceConditionTypeCollectibleLikeNew captures enum value "collectible_like_new"
	ItemSummaryByMarketplaceConditionTypeCollectibleLikeNew string = "collectible_like_new"

	// ItemSummaryByMarketplaceConditionTypeCollectibleVeryGood captures enum value "collectible_very_good"
	ItemSummaryByMarketplaceConditionTypeCollectibleVeryGood string = "collectible_very_good"

	// ItemSummaryByMarketplaceConditionTypeCollectibleGood captures enum value "collectible_good"
	ItemSummaryByMarketplaceConditionTypeCollectibleGood string = "collectible_good"

	// ItemSummaryByMarketplaceConditionTypeCollectibleAcceptable captures enum value "collectible_acceptable"
	ItemSummaryByMarketplaceConditionTypeCollectibleAcceptable string = "collectible_acceptable"

	// ItemSummaryByMarketplaceConditionTypeClubClub captures enum value "club_club"
	ItemSummaryByMarketplaceConditionTypeClubClub string = "club_club"
)

// prop value enum
func (m *ItemSummaryByMarketplace) validateConditionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSummaryByMarketplaceTypeConditionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSummaryByMarketplace) validateConditionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConditionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConditionTypeEnum("conditionType", "body", m.ConditionType); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("createdDate", "body", m.CreatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateItemName(formats strfmt.Registry) error {

	if err := validate.Required("itemName", "body", m.ItemName); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateLastUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdatedDate", "body", m.LastUpdatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateMainImage(formats strfmt.Registry) error {
	if swag.IsZero(m.MainImage) { // not required
		return nil
	}

	if m.MainImage != nil {
		if err := m.MainImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mainImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mainImage")
			}
			return err
		}
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateProductType(formats strfmt.Registry) error {

	if err := validate.Required("productType", "body", m.ProductType); err != nil {
		return err
	}

	return nil
}

var itemSummaryByMarketplaceStatusItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BUYABLE","DISCOVERABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSummaryByMarketplaceStatusItemsEnum = append(itemSummaryByMarketplaceStatusItemsEnum, v)
	}
}

func (m *ItemSummaryByMarketplace) validateStatusItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSummaryByMarketplaceStatusItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSummaryByMarketplace) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	for i := 0; i < len(m.Status); i++ {

		// value enum
		if err := m.validateStatusItemsEnum("status"+"."+strconv.Itoa(i), "body", m.Status[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this item summary by marketplace based on the context it is used
func (m *ItemSummaryByMarketplace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMainImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSummaryByMarketplace) contextValidateMainImage(ctx context.Context, formats strfmt.Registry) error {

	if m.MainImage != nil {
		if err := m.MainImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mainImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mainImage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemSummaryByMarketplace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSummaryByMarketplace) UnmarshalBinary(b []byte) error {
	var res ItemSummaryByMarketplace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}