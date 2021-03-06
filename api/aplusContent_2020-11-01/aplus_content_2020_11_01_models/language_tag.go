// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// LanguageTag The IETF language tag. This only supports the primary language subtag with one secondary language subtag. The secondary language subtag is almost always a regional designation. This does not support additional subtags beyond the primary and secondary subtags.
// **Pattern:** ^[a-z]{2,}-[A-Z0-9]{2,}$
//
// swagger:model LanguageTag
type LanguageTag string

// Validate validates this language tag
func (m LanguageTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 5); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this language tag based on context it is used
func (m LanguageTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
