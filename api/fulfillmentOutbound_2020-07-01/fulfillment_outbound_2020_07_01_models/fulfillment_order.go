// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FulfillmentOrder General information about a fulfillment order, including its status.
//
// swagger:model FulfillmentOrder
type FulfillmentOrder struct {

	// cod settings
	CodSettings *CODSettings `json:"codSettings,omitempty"`

	// delivery window
	DeliveryWindow *DeliveryWindow `json:"deliveryWindow,omitempty"`

	// The destination address submitted with the createFulfillmentOrder operation.
	// Required: true
	DestinationAddress *Address `json:"destinationAddress"`

	// A text block submitted with the createFulfillmentOrder operation. Displays in recipient-facing materials such as the packing slip.
	// Required: true
	DisplayableOrderComment *string `json:"displayableOrderComment"`

	// A date and time submitted with the createFulfillmentOrder operation. Displays as the order date in recipient-facing materials such as the packing slip.
	// Required: true
	// Format: date-time
	DisplayableOrderDate *Timestamp `json:"displayableOrderDate"`

	// A fulfillment order identifier submitted with the createFulfillmentOrder operation. Displays as the order identifier in recipient-facing materials such as the packing slip.
	// Required: true
	DisplayableOrderID *string `json:"displayableOrderId"`

	// A list of features and their fulfillment policies to apply to the order.
	FeatureConstraints []*FeatureSettings `json:"featureConstraints"`

	// fulfillment action
	FulfillmentAction FulfillmentAction `json:"fulfillmentAction,omitempty"`

	// fulfillment order status
	// Required: true
	FulfillmentOrderStatus *FulfillmentOrderStatus `json:"fulfillmentOrderStatus"`

	// fulfillment policy
	FulfillmentPolicy FulfillmentPolicy `json:"fulfillmentPolicy,omitempty"`

	// The identifier for the marketplace the fulfillment order is placed against.
	// Required: true
	MarketplaceID *string `json:"marketplaceId"`

	// notification emails
	NotificationEmails NotificationEmailList `json:"notificationEmails,omitempty"`

	// The date and time that the fulfillment order was received by an Amazon fulfillment center.
	// Required: true
	// Format: date-time
	ReceivedDate *Timestamp `json:"receivedDate"`

	// The fulfillment order identifier submitted with the createFulfillmentOrder operation.
	// Required: true
	SellerFulfillmentOrderID *string `json:"sellerFulfillmentOrderId"`

	// shipping speed category
	// Required: true
	ShippingSpeedCategory *ShippingSpeedCategory `json:"shippingSpeedCategory"`

	// The date and time that the status of the fulfillment order last changed, in ISO 8601 date time format.
	// Required: true
	// Format: date-time
	StatusUpdatedDate *Timestamp `json:"statusUpdatedDate"`
}

// Validate validates this fulfillment order
func (m *FulfillmentOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayableOrderComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayableOrderDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayableOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentOrderStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerFulfillmentOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingSpeedCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentOrder) validateCodSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.CodSettings) { // not required
		return nil
	}

	if m.CodSettings != nil {
		if err := m.CodSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codSettings")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateDeliveryWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryWindow) { // not required
		return nil
	}

	if m.DeliveryWindow != nil {
		if err := m.DeliveryWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliveryWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deliveryWindow")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateDestinationAddress(formats strfmt.Registry) error {

	if err := validate.Required("destinationAddress", "body", m.DestinationAddress); err != nil {
		return err
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateDisplayableOrderComment(formats strfmt.Registry) error {

	if err := validate.Required("displayableOrderComment", "body", m.DisplayableOrderComment); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateDisplayableOrderDate(formats strfmt.Registry) error {

	if err := validate.Required("displayableOrderDate", "body", m.DisplayableOrderDate); err != nil {
		return err
	}

	if err := validate.Required("displayableOrderDate", "body", m.DisplayableOrderDate); err != nil {
		return err
	}

	if m.DisplayableOrderDate != nil {
		if err := m.DisplayableOrderDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayableOrderDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("displayableOrderDate")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateDisplayableOrderID(formats strfmt.Registry) error {

	if err := validate.Required("displayableOrderId", "body", m.DisplayableOrderID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateFeatureConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.FeatureConstraints) { // not required
		return nil
	}

	for i := 0; i < len(m.FeatureConstraints); i++ {
		if swag.IsZero(m.FeatureConstraints[i]) { // not required
			continue
		}

		if m.FeatureConstraints[i] != nil {
			if err := m.FeatureConstraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FulfillmentOrder) validateFulfillmentAction(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentAction) { // not required
		return nil
	}

	if err := m.FulfillmentAction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentAction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentAction")
		}
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateFulfillmentOrderStatus(formats strfmt.Registry) error {

	if err := validate.Required("fulfillmentOrderStatus", "body", m.FulfillmentOrderStatus); err != nil {
		return err
	}

	if err := validate.Required("fulfillmentOrderStatus", "body", m.FulfillmentOrderStatus); err != nil {
		return err
	}

	if m.FulfillmentOrderStatus != nil {
		if err := m.FulfillmentOrderStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fulfillmentOrderStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fulfillmentOrderStatus")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateFulfillmentPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentPolicy) { // not required
		return nil
	}

	if err := m.FulfillmentPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentPolicy")
		}
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateNotificationEmails(formats strfmt.Registry) error {
	if swag.IsZero(m.NotificationEmails) { // not required
		return nil
	}

	if err := m.NotificationEmails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notificationEmails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("notificationEmails")
		}
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateReceivedDate(formats strfmt.Registry) error {

	if err := validate.Required("receivedDate", "body", m.ReceivedDate); err != nil {
		return err
	}

	if err := validate.Required("receivedDate", "body", m.ReceivedDate); err != nil {
		return err
	}

	if m.ReceivedDate != nil {
		if err := m.ReceivedDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receivedDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("receivedDate")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateSellerFulfillmentOrderID(formats strfmt.Registry) error {

	if err := validate.Required("sellerFulfillmentOrderId", "body", m.SellerFulfillmentOrderID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentOrder) validateShippingSpeedCategory(formats strfmt.Registry) error {

	if err := validate.Required("shippingSpeedCategory", "body", m.ShippingSpeedCategory); err != nil {
		return err
	}

	if err := validate.Required("shippingSpeedCategory", "body", m.ShippingSpeedCategory); err != nil {
		return err
	}

	if m.ShippingSpeedCategory != nil {
		if err := m.ShippingSpeedCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippingSpeedCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shippingSpeedCategory")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) validateStatusUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("statusUpdatedDate", "body", m.StatusUpdatedDate); err != nil {
		return err
	}

	if err := validate.Required("statusUpdatedDate", "body", m.StatusUpdatedDate); err != nil {
		return err
	}

	if m.StatusUpdatedDate != nil {
		if err := m.StatusUpdatedDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusUpdatedDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusUpdatedDate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fulfillment order based on the context it is used
func (m *FulfillmentOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCodSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestinationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayableOrderDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatureConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentOrderStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotificationEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReceivedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingSpeedCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusUpdatedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentOrder) contextValidateCodSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.CodSettings != nil {
		if err := m.CodSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codSettings")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateDeliveryWindow(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryWindow != nil {
		if err := m.DeliveryWindow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliveryWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deliveryWindow")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateDestinationAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateDisplayableOrderDate(ctx context.Context, formats strfmt.Registry) error {

	if m.DisplayableOrderDate != nil {
		if err := m.DisplayableOrderDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayableOrderDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("displayableOrderDate")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateFeatureConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FeatureConstraints); i++ {

		if m.FeatureConstraints[i] != nil {
			if err := m.FeatureConstraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FulfillmentOrder) contextValidateFulfillmentAction(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentAction.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentAction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentAction")
		}
		return err
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateFulfillmentOrderStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.FulfillmentOrderStatus != nil {
		if err := m.FulfillmentOrderStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fulfillmentOrderStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fulfillmentOrderStatus")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateFulfillmentPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentPolicy")
		}
		return err
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateNotificationEmails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NotificationEmails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notificationEmails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("notificationEmails")
		}
		return err
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateReceivedDate(ctx context.Context, formats strfmt.Registry) error {

	if m.ReceivedDate != nil {
		if err := m.ReceivedDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receivedDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("receivedDate")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateShippingSpeedCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingSpeedCategory != nil {
		if err := m.ShippingSpeedCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippingSpeedCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shippingSpeedCategory")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentOrder) contextValidateStatusUpdatedDate(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusUpdatedDate != nil {
		if err := m.StatusUpdatedDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusUpdatedDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusUpdatedDate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FulfillmentOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FulfillmentOrder) UnmarshalBinary(b []byte) error {
	var res FulfillmentOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}