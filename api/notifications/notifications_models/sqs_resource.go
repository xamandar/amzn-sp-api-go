// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SqsResource The information required to create an Amazon Simple Queue Service (Amazon SQS) queue destination.
//
// swagger:model SqsResource
type SqsResource struct {

	// The Amazon Resource Name (ARN) associated with the SQS queue.
	// Required: true
	// Max Length: 1000
	// Pattern: ^arn:aws:sqs:\S+:\S+:\S+
	Arn *string `json:"arn"`
}

// Validate validates this sqs resource
func (m *SqsResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SqsResource) validateArn(formats strfmt.Registry) error {

	if err := validate.Required("arn", "body", m.Arn); err != nil {
		return err
	}

	if err := validate.MaxLength("arn", "body", *m.Arn, 1000); err != nil {
		return err
	}

	if err := validate.Pattern("arn", "body", *m.Arn, `^arn:aws:sqs:\S+:\S+:\S+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sqs resource based on context it is used
func (m *SqsResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SqsResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SqsResource) UnmarshalBinary(b []byte) error {
	var res SqsResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}