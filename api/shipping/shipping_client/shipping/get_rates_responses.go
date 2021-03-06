// Code generated by go-swagger; DO NOT EDIT.

package shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/shipping/shipping_models"
)

// GetRatesReader is a Reader for the GetRates structure.
type GetRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRatesOK creates a GetRatesOK with default headers values
func NewGetRatesOK() *GetRatesOK {
	return &GetRatesOK{}
}

/* GetRatesOK describes a response with status code 200, with default header values.

Success.
*/
type GetRatesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesOK) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesOK  %+v", 200, o.Payload)
}
func (o *GetRatesOK) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesBadRequest creates a GetRatesBadRequest with default headers values
func NewGetRatesBadRequest() *GetRatesBadRequest {
	return &GetRatesBadRequest{}
}

/* GetRatesBadRequest describes a response with status code 400, with default header values.

Request is missing or has invalid parameters and cannot be parsed.
*/
type GetRatesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesBadRequest  %+v", 400, o.Payload)
}
func (o *GetRatesBadRequest) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesUnauthorized creates a GetRatesUnauthorized with default headers values
func NewGetRatesUnauthorized() *GetRatesUnauthorized {
	return &GetRatesUnauthorized{}
}

/* GetRatesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetRatesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetRatesUnauthorized) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesForbidden creates a GetRatesForbidden with default headers values
func NewGetRatesForbidden() *GetRatesForbidden {
	return &GetRatesForbidden{}
}

/* GetRatesForbidden describes a response with status code 403, with default header values.

403 can be caused for reasons like Access Denied, Unauthorized, Expired Token, Invalid Signature or Resource Not Found.
*/
type GetRatesForbidden struct {

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesForbidden) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesForbidden  %+v", 403, o.Payload)
}
func (o *GetRatesForbidden) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesNotFound creates a GetRatesNotFound with default headers values
func NewGetRatesNotFound() *GetRatesNotFound {
	return &GetRatesNotFound{}
}

/* GetRatesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetRatesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesNotFound) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesNotFound  %+v", 404, o.Payload)
}
func (o *GetRatesNotFound) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesTooManyRequests creates a GetRatesTooManyRequests with default headers values
func NewGetRatesTooManyRequests() *GetRatesTooManyRequests {
	return &GetRatesTooManyRequests{}
}

/* GetRatesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetRatesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetRatesTooManyRequests) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesInternalServerError creates a GetRatesInternalServerError with default headers values
func NewGetRatesInternalServerError() *GetRatesInternalServerError {
	return &GetRatesInternalServerError{}
}

/* GetRatesInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type GetRatesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRatesInternalServerError) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesServiceUnavailable creates a GetRatesServiceUnavailable with default headers values
func NewGetRatesServiceUnavailable() *GetRatesServiceUnavailable {
	return &GetRatesServiceUnavailable{}
}

/* GetRatesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetRatesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *shipping_models.GetRatesResponse
}

func (o *GetRatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /shipping/v1/rates][%d] getRatesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetRatesServiceUnavailable) GetPayload() *shipping_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
