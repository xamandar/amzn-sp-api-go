// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetEligibleShipmentServicesRequest Request schema.
//
// swagger:model GetEligibleShipmentServicesRequest
type GetEligibleShipmentServicesRequest struct {

	// Shipment information required for requesting shipping service offers.
	// Required: true
	ShipmentRequestDetails *ShipmentRequestDetails `json:"ShipmentRequestDetails"`

	// shipping offering filter
	ShippingOfferingFilter *ShippingOfferingFilter `json:"ShippingOfferingFilter,omitempty"`
}

// Validate validates this get eligible shipment services request
func (m *GetEligibleShipmentServicesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipmentRequestDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingOfferingFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetEligibleShipmentServicesRequest) validateShipmentRequestDetails(formats strfmt.Registry) error {

	if err := validate.Required("ShipmentRequestDetails", "body", m.ShipmentRequestDetails); err != nil {
		return err
	}

	if m.ShipmentRequestDetails != nil {
		if err := m.ShipmentRequestDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipmentRequestDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipmentRequestDetails")
			}
			return err
		}
	}

	return nil
}

func (m *GetEligibleShipmentServicesRequest) validateShippingOfferingFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingOfferingFilter) { // not required
		return nil
	}

	if m.ShippingOfferingFilter != nil {
		if err := m.ShippingOfferingFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingOfferingFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingOfferingFilter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get eligible shipment services request based on the context it is used
func (m *GetEligibleShipmentServicesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShipmentRequestDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingOfferingFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetEligibleShipmentServicesRequest) contextValidateShipmentRequestDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentRequestDetails != nil {
		if err := m.ShipmentRequestDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipmentRequestDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipmentRequestDetails")
			}
			return err
		}
	}

	return nil
}

func (m *GetEligibleShipmentServicesRequest) contextValidateShippingOfferingFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingOfferingFilter != nil {
		if err := m.ShippingOfferingFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingOfferingFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingOfferingFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetEligibleShipmentServicesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEligibleShipmentServicesRequest) UnmarshalBinary(b []byte) error {
	var res GetEligibleShipmentServicesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}