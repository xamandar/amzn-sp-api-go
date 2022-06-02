// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StateOrProvinceCode The state or province code. **Note.** Required in the Canada, US, and UK marketplaces. Also required for shipments originating from China.
//
// swagger:model StateOrProvinceCode
type StateOrProvinceCode string

// Validate validates this state or province code
func (m StateOrProvinceCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MaxLength("", "body", string(m), 30); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this state or province code based on context it is used
func (m StateOrProvinceCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}