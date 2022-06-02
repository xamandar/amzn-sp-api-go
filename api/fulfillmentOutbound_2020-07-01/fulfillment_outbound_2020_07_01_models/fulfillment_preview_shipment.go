// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FulfillmentPreviewShipment Delivery and item information for a shipment in a fulfillment order preview.
//
// swagger:model FulfillmentPreviewShipment
type FulfillmentPreviewShipment struct {

	// The earliest date that the shipment is expected to arrive at its destination.
	// Format: date-time
	EarliestArrivalDate Timestamp `json:"earliestArrivalDate,omitempty"`

	// The earliest date that the shipment is expected to be sent from the fulfillment center, in ISO 8601 date time format.
	// Format: date-time
	EarliestShipDate Timestamp `json:"earliestShipDate,omitempty"`

	// Information about the items in the shipment.
	// Required: true
	FulfillmentPreviewItems FulfillmentPreviewItemList `json:"fulfillmentPreviewItems"`

	// The latest date that the shipment is expected to arrive at its destination, in ISO 8601 date time format.
	// Format: date-time
	LatestArrivalDate Timestamp `json:"latestArrivalDate,omitempty"`

	// The latest date that the shipment is expected to be sent from the fulfillment center, in ISO 8601 date time format.
	// Format: date-time
	LatestShipDate Timestamp `json:"latestShipDate,omitempty"`

	// Provides additional insight into the shipment timeline when exact delivery dates are not able to be precomputed.
	ShippingNotes []string `json:"shippingNotes"`
}

// Validate validates this fulfillment preview shipment
func (m *FulfillmentPreviewShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEarliestArrivalDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEarliestShipDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentPreviewItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestArrivalDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestShipDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentPreviewShipment) validateEarliestArrivalDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EarliestArrivalDate) { // not required
		return nil
	}

	if err := m.EarliestArrivalDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("earliestArrivalDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("earliestArrivalDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) validateEarliestShipDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EarliestShipDate) { // not required
		return nil
	}

	if err := m.EarliestShipDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("earliestShipDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("earliestShipDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) validateFulfillmentPreviewItems(formats strfmt.Registry) error {

	if err := validate.Required("fulfillmentPreviewItems", "body", m.FulfillmentPreviewItems); err != nil {
		return err
	}

	if err := m.FulfillmentPreviewItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentPreviewItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentPreviewItems")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) validateLatestArrivalDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestArrivalDate) { // not required
		return nil
	}

	if err := m.LatestArrivalDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("latestArrivalDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("latestArrivalDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) validateLatestShipDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestShipDate) { // not required
		return nil
	}

	if err := m.LatestShipDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("latestShipDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("latestShipDate")
		}
		return err
	}

	return nil
}

// ContextValidate validate this fulfillment preview shipment based on the context it is used
func (m *FulfillmentPreviewShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEarliestArrivalDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEarliestShipDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentPreviewItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestArrivalDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestShipDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentPreviewShipment) contextValidateEarliestArrivalDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EarliestArrivalDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("earliestArrivalDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("earliestArrivalDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) contextValidateEarliestShipDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EarliestShipDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("earliestShipDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("earliestShipDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) contextValidateFulfillmentPreviewItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentPreviewItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentPreviewItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentPreviewItems")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) contextValidateLatestArrivalDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LatestArrivalDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("latestArrivalDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("latestArrivalDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentPreviewShipment) contextValidateLatestShipDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LatestShipDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("latestShipDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("latestShipDate")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FulfillmentPreviewShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FulfillmentPreviewShipment) UnmarshalBinary(b []byte) error {
	var res FulfillmentPreviewShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}