// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// GetFulfillmentOrderReader is a Reader for the GetFulfillmentOrder structure.
type GetFulfillmentOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFulfillmentOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFulfillmentOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFulfillmentOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFulfillmentOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFulfillmentOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFulfillmentOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFulfillmentOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFulfillmentOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFulfillmentOrderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFulfillmentOrderOK creates a GetFulfillmentOrderOK with default headers values
func NewGetFulfillmentOrderOK() *GetFulfillmentOrderOK {
	return &GetFulfillmentOrderOK{}
}

/* GetFulfillmentOrderOK describes a response with status code 200, with default header values.

Success.
*/
type GetFulfillmentOrderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderOK  %+v", 200, o.Payload)
}
func (o *GetFulfillmentOrderOK) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderBadRequest creates a GetFulfillmentOrderBadRequest with default headers values
func NewGetFulfillmentOrderBadRequest() *GetFulfillmentOrderBadRequest {
	return &GetFulfillmentOrderBadRequest{}
}

/* GetFulfillmentOrderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFulfillmentOrderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}
func (o *GetFulfillmentOrderBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderUnauthorized creates a GetFulfillmentOrderUnauthorized with default headers values
func NewGetFulfillmentOrderUnauthorized() *GetFulfillmentOrderUnauthorized {
	return &GetFulfillmentOrderUnauthorized{}
}

/* GetFulfillmentOrderUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFulfillmentOrderUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFulfillmentOrderUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderForbidden creates a GetFulfillmentOrderForbidden with default headers values
func NewGetFulfillmentOrderForbidden() *GetFulfillmentOrderForbidden {
	return &GetFulfillmentOrderForbidden{}
}

/* GetFulfillmentOrderForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFulfillmentOrderForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderForbidden  %+v", 403, o.Payload)
}
func (o *GetFulfillmentOrderForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderNotFound creates a GetFulfillmentOrderNotFound with default headers values
func NewGetFulfillmentOrderNotFound() *GetFulfillmentOrderNotFound {
	return &GetFulfillmentOrderNotFound{}
}

/* GetFulfillmentOrderNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetFulfillmentOrderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderNotFound  %+v", 404, o.Payload)
}
func (o *GetFulfillmentOrderNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderTooManyRequests creates a GetFulfillmentOrderTooManyRequests with default headers values
func NewGetFulfillmentOrderTooManyRequests() *GetFulfillmentOrderTooManyRequests {
	return &GetFulfillmentOrderTooManyRequests{}
}

/* GetFulfillmentOrderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFulfillmentOrderTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFulfillmentOrderTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderInternalServerError creates a GetFulfillmentOrderInternalServerError with default headers values
func NewGetFulfillmentOrderInternalServerError() *GetFulfillmentOrderInternalServerError {
	return &GetFulfillmentOrderInternalServerError{}
}

/* GetFulfillmentOrderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFulfillmentOrderInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFulfillmentOrderInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderServiceUnavailable creates a GetFulfillmentOrderServiceUnavailable with default headers values
func NewGetFulfillmentOrderServiceUnavailable() *GetFulfillmentOrderServiceUnavailable {
	return &GetFulfillmentOrderServiceUnavailable{}
}

/* GetFulfillmentOrderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFulfillmentOrderServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

func (o *GetFulfillmentOrderServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFulfillmentOrderServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}