// Code generated by go-swagger; DO NOT EDIT.

package product_pricing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/productPricingV0/product_pricing_v0_models"
)

// GetCompetitivePricingReader is a Reader for the GetCompetitivePricing structure.
type GetCompetitivePricingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompetitivePricingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCompetitivePricingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCompetitivePricingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCompetitivePricingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCompetitivePricingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCompetitivePricingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCompetitivePricingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCompetitivePricingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCompetitivePricingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCompetitivePricingOK creates a GetCompetitivePricingOK with default headers values
func NewGetCompetitivePricingOK() *GetCompetitivePricingOK {
	return &GetCompetitivePricingOK{}
}

/* GetCompetitivePricingOK describes a response with status code 200, with default header values.

Success.
*/
type GetCompetitivePricingOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingOK) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingOK  %+v", 200, o.Payload)
}
func (o *GetCompetitivePricingOK) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingBadRequest creates a GetCompetitivePricingBadRequest with default headers values
func NewGetCompetitivePricingBadRequest() *GetCompetitivePricingBadRequest {
	return &GetCompetitivePricingBadRequest{}
}

/* GetCompetitivePricingBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetCompetitivePricingBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingBadRequest  %+v", 400, o.Payload)
}
func (o *GetCompetitivePricingBadRequest) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingUnauthorized creates a GetCompetitivePricingUnauthorized with default headers values
func NewGetCompetitivePricingUnauthorized() *GetCompetitivePricingUnauthorized {
	return &GetCompetitivePricingUnauthorized{}
}

/* GetCompetitivePricingUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetCompetitivePricingUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCompetitivePricingUnauthorized) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingForbidden creates a GetCompetitivePricingForbidden with default headers values
func NewGetCompetitivePricingForbidden() *GetCompetitivePricingForbidden {
	return &GetCompetitivePricingForbidden{}
}

/* GetCompetitivePricingForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetCompetitivePricingForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingForbidden) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingForbidden  %+v", 403, o.Payload)
}
func (o *GetCompetitivePricingForbidden) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingNotFound creates a GetCompetitivePricingNotFound with default headers values
func NewGetCompetitivePricingNotFound() *GetCompetitivePricingNotFound {
	return &GetCompetitivePricingNotFound{}
}

/* GetCompetitivePricingNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetCompetitivePricingNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingNotFound) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingNotFound  %+v", 404, o.Payload)
}
func (o *GetCompetitivePricingNotFound) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingTooManyRequests creates a GetCompetitivePricingTooManyRequests with default headers values
func NewGetCompetitivePricingTooManyRequests() *GetCompetitivePricingTooManyRequests {
	return &GetCompetitivePricingTooManyRequests{}
}

/* GetCompetitivePricingTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetCompetitivePricingTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCompetitivePricingTooManyRequests) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingInternalServerError creates a GetCompetitivePricingInternalServerError with default headers values
func NewGetCompetitivePricingInternalServerError() *GetCompetitivePricingInternalServerError {
	return &GetCompetitivePricingInternalServerError{}
}

/* GetCompetitivePricingInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetCompetitivePricingInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCompetitivePricingInternalServerError) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompetitivePricingServiceUnavailable creates a GetCompetitivePricingServiceUnavailable with default headers values
func NewGetCompetitivePricingServiceUnavailable() *GetCompetitivePricingServiceUnavailable {
	return &GetCompetitivePricingServiceUnavailable{}
}

/* GetCompetitivePricingServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetCompetitivePricingServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetCompetitivePricingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/competitivePrice][%d] getCompetitivePricingServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCompetitivePricingServiceUnavailable) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetCompetitivePricingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}