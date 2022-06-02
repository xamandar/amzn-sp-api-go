// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/fulfillmentInboundV0/fulfillment_inbound_v0_models"
)

// ConfirmPreorderReader is a Reader for the ConfirmPreorder structure.
type ConfirmPreorderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfirmPreorderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfirmPreorderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfirmPreorderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewConfirmPreorderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConfirmPreorderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConfirmPreorderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewConfirmPreorderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfirmPreorderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewConfirmPreorderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfirmPreorderOK creates a ConfirmPreorderOK with default headers values
func NewConfirmPreorderOK() *ConfirmPreorderOK {
	return &ConfirmPreorderOK{}
}

/* ConfirmPreorderOK describes a response with status code 200, with default header values.

Success.
*/
type ConfirmPreorderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderOK) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderOK  %+v", 200, o.Payload)
}
func (o *ConfirmPreorderOK) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderBadRequest creates a ConfirmPreorderBadRequest with default headers values
func NewConfirmPreorderBadRequest() *ConfirmPreorderBadRequest {
	return &ConfirmPreorderBadRequest{}
}

/* ConfirmPreorderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ConfirmPreorderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderBadRequest  %+v", 400, o.Payload)
}
func (o *ConfirmPreorderBadRequest) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderUnauthorized creates a ConfirmPreorderUnauthorized with default headers values
func NewConfirmPreorderUnauthorized() *ConfirmPreorderUnauthorized {
	return &ConfirmPreorderUnauthorized{}
}

/* ConfirmPreorderUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type ConfirmPreorderUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderUnauthorized  %+v", 401, o.Payload)
}
func (o *ConfirmPreorderUnauthorized) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderForbidden creates a ConfirmPreorderForbidden with default headers values
func NewConfirmPreorderForbidden() *ConfirmPreorderForbidden {
	return &ConfirmPreorderForbidden{}
}

/* ConfirmPreorderForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ConfirmPreorderForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderForbidden) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderForbidden  %+v", 403, o.Payload)
}
func (o *ConfirmPreorderForbidden) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderNotFound creates a ConfirmPreorderNotFound with default headers values
func NewConfirmPreorderNotFound() *ConfirmPreorderNotFound {
	return &ConfirmPreorderNotFound{}
}

/* ConfirmPreorderNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type ConfirmPreorderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderNotFound) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderNotFound  %+v", 404, o.Payload)
}
func (o *ConfirmPreorderNotFound) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderTooManyRequests creates a ConfirmPreorderTooManyRequests with default headers values
func NewConfirmPreorderTooManyRequests() *ConfirmPreorderTooManyRequests {
	return &ConfirmPreorderTooManyRequests{}
}

/* ConfirmPreorderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ConfirmPreorderTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderTooManyRequests  %+v", 429, o.Payload)
}
func (o *ConfirmPreorderTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderInternalServerError creates a ConfirmPreorderInternalServerError with default headers values
func NewConfirmPreorderInternalServerError() *ConfirmPreorderInternalServerError {
	return &ConfirmPreorderInternalServerError{}
}

/* ConfirmPreorderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ConfirmPreorderInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderInternalServerError  %+v", 500, o.Payload)
}
func (o *ConfirmPreorderInternalServerError) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmPreorderServiceUnavailable creates a ConfirmPreorderServiceUnavailable with default headers values
func NewConfirmPreorderServiceUnavailable() *ConfirmPreorderServiceUnavailable {
	return &ConfirmPreorderServiceUnavailable{}
}

/* ConfirmPreorderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ConfirmPreorderServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.ConfirmPreorderResponse
}

func (o *ConfirmPreorderServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/preorder/confirm][%d] confirmPreorderServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ConfirmPreorderServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.ConfirmPreorderResponse {
	return o.Payload
}

func (o *ConfirmPreorderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.ConfirmPreorderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}