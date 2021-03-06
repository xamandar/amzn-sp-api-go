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

// GetFulfillmentPreviewReader is a Reader for the GetFulfillmentPreview structure.
type GetFulfillmentPreviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFulfillmentPreviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFulfillmentPreviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFulfillmentPreviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFulfillmentPreviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFulfillmentPreviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFulfillmentPreviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFulfillmentPreviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFulfillmentPreviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFulfillmentPreviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFulfillmentPreviewOK creates a GetFulfillmentPreviewOK with default headers values
func NewGetFulfillmentPreviewOK() *GetFulfillmentPreviewOK {
	return &GetFulfillmentPreviewOK{}
}

/* GetFulfillmentPreviewOK describes a response with status code 200, with default header values.

Success.
*/
type GetFulfillmentPreviewOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewOK) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewOK  %+v", 200, o.Payload)
}
func (o *GetFulfillmentPreviewOK) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewBadRequest creates a GetFulfillmentPreviewBadRequest with default headers values
func NewGetFulfillmentPreviewBadRequest() *GetFulfillmentPreviewBadRequest {
	return &GetFulfillmentPreviewBadRequest{}
}

/* GetFulfillmentPreviewBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFulfillmentPreviewBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewBadRequest) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewBadRequest  %+v", 400, o.Payload)
}
func (o *GetFulfillmentPreviewBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewUnauthorized creates a GetFulfillmentPreviewUnauthorized with default headers values
func NewGetFulfillmentPreviewUnauthorized() *GetFulfillmentPreviewUnauthorized {
	return &GetFulfillmentPreviewUnauthorized{}
}

/* GetFulfillmentPreviewUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFulfillmentPreviewUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFulfillmentPreviewUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewForbidden creates a GetFulfillmentPreviewForbidden with default headers values
func NewGetFulfillmentPreviewForbidden() *GetFulfillmentPreviewForbidden {
	return &GetFulfillmentPreviewForbidden{}
}

/* GetFulfillmentPreviewForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFulfillmentPreviewForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewForbidden) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewForbidden  %+v", 403, o.Payload)
}
func (o *GetFulfillmentPreviewForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewNotFound creates a GetFulfillmentPreviewNotFound with default headers values
func NewGetFulfillmentPreviewNotFound() *GetFulfillmentPreviewNotFound {
	return &GetFulfillmentPreviewNotFound{}
}

/* GetFulfillmentPreviewNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetFulfillmentPreviewNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewNotFound) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewNotFound  %+v", 404, o.Payload)
}
func (o *GetFulfillmentPreviewNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewTooManyRequests creates a GetFulfillmentPreviewTooManyRequests with default headers values
func NewGetFulfillmentPreviewTooManyRequests() *GetFulfillmentPreviewTooManyRequests {
	return &GetFulfillmentPreviewTooManyRequests{}
}

/* GetFulfillmentPreviewTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFulfillmentPreviewTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFulfillmentPreviewTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewInternalServerError creates a GetFulfillmentPreviewInternalServerError with default headers values
func NewGetFulfillmentPreviewInternalServerError() *GetFulfillmentPreviewInternalServerError {
	return &GetFulfillmentPreviewInternalServerError{}
}

/* GetFulfillmentPreviewInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFulfillmentPreviewInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFulfillmentPreviewInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentPreviewServiceUnavailable creates a GetFulfillmentPreviewServiceUnavailable with default headers values
func NewGetFulfillmentPreviewServiceUnavailable() *GetFulfillmentPreviewServiceUnavailable {
	return &GetFulfillmentPreviewServiceUnavailable{}
}

/* GetFulfillmentPreviewServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFulfillmentPreviewServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse
}

func (o *GetFulfillmentPreviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fba/outbound/2020-07-01/fulfillmentOrders/preview][%d] getFulfillmentPreviewServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFulfillmentPreviewServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse {
	return o.Payload
}

func (o *GetFulfillmentPreviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentPreviewResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
