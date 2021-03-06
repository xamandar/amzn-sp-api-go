// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ShippingServiceIdentifier An Amazon-defined shipping service identifier.
//
// swagger:model ShippingServiceIdentifier
type ShippingServiceIdentifier string

// Validate validates this shipping service identifier
func (m ShippingServiceIdentifier) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this shipping service identifier based on context it is used
func (m ShippingServiceIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
