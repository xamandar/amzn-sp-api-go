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

// PutTransportDetailsReader is a Reader for the PutTransportDetails structure.
type PutTransportDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTransportDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTransportDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTransportDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTransportDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTransportDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTransportDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTransportDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTransportDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTransportDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTransportDetailsOK creates a PutTransportDetailsOK with default headers values
func NewPutTransportDetailsOK() *PutTransportDetailsOK {
	return &PutTransportDetailsOK{}
}

/* PutTransportDetailsOK describes a response with status code 200, with default header values.

Success.
*/
type PutTransportDetailsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsOK) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsOK  %+v", 200, o.Payload)
}
func (o *PutTransportDetailsOK) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsBadRequest creates a PutTransportDetailsBadRequest with default headers values
func NewPutTransportDetailsBadRequest() *PutTransportDetailsBadRequest {
	return &PutTransportDetailsBadRequest{}
}

/* PutTransportDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type PutTransportDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *PutTransportDetailsBadRequest) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsUnauthorized creates a PutTransportDetailsUnauthorized with default headers values
func NewPutTransportDetailsUnauthorized() *PutTransportDetailsUnauthorized {
	return &PutTransportDetailsUnauthorized{}
}

/* PutTransportDetailsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type PutTransportDetailsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *PutTransportDetailsUnauthorized) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsForbidden creates a PutTransportDetailsForbidden with default headers values
func NewPutTransportDetailsForbidden() *PutTransportDetailsForbidden {
	return &PutTransportDetailsForbidden{}
}

/* PutTransportDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type PutTransportDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsForbidden) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsForbidden  %+v", 403, o.Payload)
}
func (o *PutTransportDetailsForbidden) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsNotFound creates a PutTransportDetailsNotFound with default headers values
func NewPutTransportDetailsNotFound() *PutTransportDetailsNotFound {
	return &PutTransportDetailsNotFound{}
}

/* PutTransportDetailsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type PutTransportDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsNotFound) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsNotFound  %+v", 404, o.Payload)
}
func (o *PutTransportDetailsNotFound) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsTooManyRequests creates a PutTransportDetailsTooManyRequests with default headers values
func NewPutTransportDetailsTooManyRequests() *PutTransportDetailsTooManyRequests {
	return &PutTransportDetailsTooManyRequests{}
}

/* PutTransportDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type PutTransportDetailsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsTooManyRequests  %+v", 429, o.Payload)
}
func (o *PutTransportDetailsTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsInternalServerError creates a PutTransportDetailsInternalServerError with default headers values
func NewPutTransportDetailsInternalServerError() *PutTransportDetailsInternalServerError {
	return &PutTransportDetailsInternalServerError{}
}

/* PutTransportDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type PutTransportDetailsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsInternalServerError  %+v", 500, o.Payload)
}
func (o *PutTransportDetailsInternalServerError) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTransportDetailsServiceUnavailable creates a PutTransportDetailsServiceUnavailable with default headers values
func NewPutTransportDetailsServiceUnavailable() *PutTransportDetailsServiceUnavailable {
	return &PutTransportDetailsServiceUnavailable{}
}

/* PutTransportDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type PutTransportDetailsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.PutTransportDetailsResponse
}

func (o *PutTransportDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}/transport][%d] putTransportDetailsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PutTransportDetailsServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.PutTransportDetailsResponse {
	return o.Payload
}

func (o *PutTransportDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.PutTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}