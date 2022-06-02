// Code generated by go-swagger; DO NOT EDIT.

package vendor_invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentPaymentsV1/vendor_direct_fulfillment_payments_v1_models"
)

// SubmitInvoiceReader is a Reader for the SubmitInvoice structure.
type SubmitInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSubmitInvoiceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitInvoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitInvoiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitInvoiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSubmitInvoiceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSubmitInvoiceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSubmitInvoiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitInvoiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSubmitInvoiceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitInvoiceAccepted creates a SubmitInvoiceAccepted with default headers values
func NewSubmitInvoiceAccepted() *SubmitInvoiceAccepted {
	return &SubmitInvoiceAccepted{}
}

/* SubmitInvoiceAccepted describes a response with status code 202, with default header values.

Success.
*/
type SubmitInvoiceAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceAccepted  %+v", 202, o.Payload)
}
func (o *SubmitInvoiceAccepted) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceBadRequest creates a SubmitInvoiceBadRequest with default headers values
func NewSubmitInvoiceBadRequest() *SubmitInvoiceBadRequest {
	return &SubmitInvoiceBadRequest{}
}

/* SubmitInvoiceBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SubmitInvoiceBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceBadRequest  %+v", 400, o.Payload)
}
func (o *SubmitInvoiceBadRequest) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceForbidden creates a SubmitInvoiceForbidden with default headers values
func NewSubmitInvoiceForbidden() *SubmitInvoiceForbidden {
	return &SubmitInvoiceForbidden{}
}

/* SubmitInvoiceForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SubmitInvoiceForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceForbidden  %+v", 403, o.Payload)
}
func (o *SubmitInvoiceForbidden) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceNotFound creates a SubmitInvoiceNotFound with default headers values
func NewSubmitInvoiceNotFound() *SubmitInvoiceNotFound {
	return &SubmitInvoiceNotFound{}
}

/* SubmitInvoiceNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SubmitInvoiceNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceNotFound  %+v", 404, o.Payload)
}
func (o *SubmitInvoiceNotFound) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceRequestEntityTooLarge creates a SubmitInvoiceRequestEntityTooLarge with default headers values
func NewSubmitInvoiceRequestEntityTooLarge() *SubmitInvoiceRequestEntityTooLarge {
	return &SubmitInvoiceRequestEntityTooLarge{}
}

/* SubmitInvoiceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SubmitInvoiceRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *SubmitInvoiceRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceUnsupportedMediaType creates a SubmitInvoiceUnsupportedMediaType with default headers values
func NewSubmitInvoiceUnsupportedMediaType() *SubmitInvoiceUnsupportedMediaType {
	return &SubmitInvoiceUnsupportedMediaType{}
}

/* SubmitInvoiceUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type SubmitInvoiceUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *SubmitInvoiceUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceTooManyRequests creates a SubmitInvoiceTooManyRequests with default headers values
func NewSubmitInvoiceTooManyRequests() *SubmitInvoiceTooManyRequests {
	return &SubmitInvoiceTooManyRequests{}
}

/* SubmitInvoiceTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SubmitInvoiceTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceTooManyRequests  %+v", 429, o.Payload)
}
func (o *SubmitInvoiceTooManyRequests) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceInternalServerError creates a SubmitInvoiceInternalServerError with default headers values
func NewSubmitInvoiceInternalServerError() *SubmitInvoiceInternalServerError {
	return &SubmitInvoiceInternalServerError{}
}

/* SubmitInvoiceInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type SubmitInvoiceInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceInternalServerError  %+v", 500, o.Payload)
}
func (o *SubmitInvoiceInternalServerError) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInvoiceServiceUnavailable creates a SubmitInvoiceServiceUnavailable with default headers values
func NewSubmitInvoiceServiceUnavailable() *SubmitInvoiceServiceUnavailable {
	return &SubmitInvoiceServiceUnavailable{}
}

/* SubmitInvoiceServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SubmitInvoiceServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse
}

func (o *SubmitInvoiceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/payments/v1/invoices][%d] submitInvoiceServiceUnavailable  %+v", 503, o.Payload)
}
func (o *SubmitInvoiceServiceUnavailable) GetPayload() *vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse {
	return o.Payload
}

func (o *SubmitInvoiceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_payments_v1_models.SubmitInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}