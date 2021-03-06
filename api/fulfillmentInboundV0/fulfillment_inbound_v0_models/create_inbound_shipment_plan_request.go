// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateInboundShipmentPlanRequest The request schema for the createInboundShipmentPlan operation.
//
// swagger:model CreateInboundShipmentPlanRequest
type CreateInboundShipmentPlanRequest struct {

	// inbound shipment plan request items
	// Required: true
	InboundShipmentPlanRequestItems InboundShipmentPlanRequestItemList `json:"InboundShipmentPlanRequestItems"`

	// The seller's preference for label preparation for an inbound shipment.
	// Required: true
	LabelPrepPreference *LabelPrepPreference `json:"LabelPrepPreference"`

	// The address from which the inbound shipment will be sent.
	// Required: true
	ShipFromAddress *Address `json:"ShipFromAddress"`

	// The two-character country code for the country where the inbound shipment is to be sent.
	//
	// Note: Not required. Specifying both ShipToCountryCode and ShipToCountrySubdivisionCode returns an error.
	//
	//  Values:
	//
	//  ShipToCountryCode values for North America:
	//  * CA – Canada
	//  * MX - Mexico
	//  * US - United States
	//
	// ShipToCountryCode values for MCI sellers in Europe:
	//  * DE – Germany
	//  * ES – Spain
	//  * FR – France
	//  * GB – United Kingdom
	//  * IT – Italy
	//
	// Default: The country code for the seller's home marketplace.
	ShipToCountryCode string `json:"ShipToCountryCode,omitempty"`

	// The two-character country code, followed by a dash and then up to three characters that represent the subdivision of the country where the inbound shipment is to be sent. For example, "IN-MH". In full ISO 3166-2 format.
	//
	// Note: Not required. Specifying both ShipToCountryCode and ShipToCountrySubdivisionCode returns an error.
	ShipToCountrySubdivisionCode string `json:"ShipToCountrySubdivisionCode,omitempty"`
}

// Validate validates this create inbound shipment plan request
func (m *CreateInboundShipmentPlanRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInboundShipmentPlanRequestItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelPrepPreference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateInboundShipmentPlanRequest) validateInboundShipmentPlanRequestItems(formats strfmt.Registry) error {

	if err := validate.Required("InboundShipmentPlanRequestItems", "body", m.InboundShipmentPlanRequestItems); err != nil {
		return err
	}

	if err := m.InboundShipmentPlanRequestItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InboundShipmentPlanRequestItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InboundShipmentPlanRequestItems")
		}
		return err
	}

	return nil
}

func (m *CreateInboundShipmentPlanRequest) validateLabelPrepPreference(formats strfmt.Registry) error {

	if err := validate.Required("LabelPrepPreference", "body", m.LabelPrepPreference); err != nil {
		return err
	}

	if err := validate.Required("LabelPrepPreference", "body", m.LabelPrepPreference); err != nil {
		return err
	}

	if m.LabelPrepPreference != nil {
		if err := m.LabelPrepPreference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LabelPrepPreference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LabelPrepPreference")
			}
			return err
		}
	}

	return nil
}

func (m *CreateInboundShipmentPlanRequest) validateShipFromAddress(formats strfmt.Registry) error {

	if err := validate.Required("ShipFromAddress", "body", m.ShipFromAddress); err != nil {
		return err
	}

	if m.ShipFromAddress != nil {
		if err := m.ShipFromAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipFromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipFromAddress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create inbound shipment plan request based on the context it is used
func (m *CreateInboundShipmentPlanRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInboundShipmentPlanRequestItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelPrepPreference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFromAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateInboundShipmentPlanRequest) contextValidateInboundShipmentPlanRequestItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InboundShipmentPlanRequestItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InboundShipmentPlanRequestItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InboundShipmentPlanRequestItems")
		}
		return err
	}

	return nil
}

func (m *CreateInboundShipmentPlanRequest) contextValidateLabelPrepPreference(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelPrepPreference != nil {
		if err := m.LabelPrepPreference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LabelPrepPreference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LabelPrepPreference")
			}
			return err
		}
	}

	return nil
}

func (m *CreateInboundShipmentPlanRequest) contextValidateShipFromAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFromAddress != nil {
		if err := m.ShipFromAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipFromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipFromAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateInboundShipmentPlanRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateInboundShipmentPlanRequest) UnmarshalBinary(b []byte) error {
	var res CreateInboundShipmentPlanRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
