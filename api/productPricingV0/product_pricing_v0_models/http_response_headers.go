// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HTTPResponseHeaders A mapping of additional HTTP headers to send/receive for the individual batch request.
//
// swagger:model HttpResponseHeaders
type HTTPResponseHeaders struct {

	// The timestamp that the API request was received.  For more information, consult [RFC 2616 Section 14](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Date string `json:"Date,omitempty"`

	// Unique request reference ID.
	XAmznRequestID string `json:"x-amzn-RequestId,omitempty"`

	// Http response headers
	HTTPResponseHeaders map[string]string `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HTTPResponseHeaders) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The timestamp that the API request was received.  For more information, consult [RFC 2616 Section 14](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
		Date string `json:"Date,omitempty"`

		// Unique request reference ID.
		XAmznRequestID string `json:"x-amzn-RequestId,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HTTPResponseHeaders

	rcv.Date = stage1.Date
	rcv.XAmznRequestID = stage1.XAmznRequestID
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Date")
	delete(stage2, "x-amzn-RequestId")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]string)
		for k, v := range stage2 {
			var toadd string
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.HTTPResponseHeaders = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HTTPResponseHeaders) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The timestamp that the API request was received.  For more information, consult [RFC 2616 Section 14](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
		Date string `json:"Date,omitempty"`

		// Unique request reference ID.
		XAmznRequestID string `json:"x-amzn-RequestId,omitempty"`
	}

	stage1.Date = m.Date
	stage1.XAmznRequestID = m.XAmznRequestID

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HTTPResponseHeaders) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HTTPResponseHeaders)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this Http response headers
func (m *HTTPResponseHeaders) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Http response headers based on context it is used
func (m *HTTPResponseHeaders) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HTTPResponseHeaders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPResponseHeaders) UnmarshalBinary(b []byte) error {
	var res HTTPResponseHeaders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
