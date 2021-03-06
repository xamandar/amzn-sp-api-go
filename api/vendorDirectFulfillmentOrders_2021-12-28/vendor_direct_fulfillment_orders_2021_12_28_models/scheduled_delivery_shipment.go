// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScheduledDeliveryShipment Dates for the scheduled delivery shipments.
//
// swagger:model ScheduledDeliveryShipment
type ScheduledDeliveryShipment struct {

	// Earliest nominated delivery date for the scheduled delivery.
	// Format: date-time
	EarliestNominatedDeliveryDate strfmt.DateTime `json:"earliestNominatedDeliveryDate,omitempty"`

	// Latest nominated delivery date for the scheduled delivery.
	// Format: date-time
	LatestNominatedDeliveryDate strfmt.DateTime `json:"latestNominatedDeliveryDate,omitempty"`

	// Scheduled delivery service type.
	ScheduledDeliveryServiceType string `json:"scheduledDeliveryServiceType,omitempty"`
}

// Validate validates this scheduled delivery shipment
func (m *ScheduledDeliveryShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEarliestNominatedDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestNominatedDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduledDeliveryShipment) validateEarliestNominatedDeliveryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EarliestNominatedDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("earliestNominatedDeliveryDate", "body", "date-time", m.EarliestNominatedDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScheduledDeliveryShipment) validateLatestNominatedDeliveryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestNominatedDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("latestNominatedDeliveryDate", "body", "date-time", m.LatestNominatedDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scheduled delivery shipment based on context it is used
func (m *ScheduledDeliveryShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledDeliveryShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledDeliveryShipment) UnmarshalBinary(b []byte) error {
	var res ScheduledDeliveryShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
