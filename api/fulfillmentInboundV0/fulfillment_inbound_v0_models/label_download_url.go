// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LabelDownloadURL label download URL
//
// swagger:model LabelDownloadURL
type LabelDownloadURL struct {

	// URL to download the label for the package. Note: The URL will only be valid for 15 seconds
	DownloadURL string `json:"DownloadURL,omitempty"`
}

// Validate validates this label download URL
func (m *LabelDownloadURL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this label download URL based on context it is used
func (m *LabelDownloadURL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelDownloadURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelDownloadURL) UnmarshalBinary(b []byte) error {
	var res LabelDownloadURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}