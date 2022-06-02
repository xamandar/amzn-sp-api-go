// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/merchantFulfillmentV0/merchant_fulfillment_v0_models"
)

// GetEligibleShipmentServicesReader is a Reader for the GetEligibleShipmentServices structure.
type GetEligibleShipmentServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEligibleShipmentServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEligibleShipmentServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEligibleShipmentServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEligibleShipmentServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEligibleShipmentServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEligibleShipmentServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEligibleShipmentServicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEligibleShipmentServicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetEligibleShipmentServicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEligibleShipmentServicesOK creates a GetEligibleShipmentServicesOK with default headers values
func NewGetEligibleShipmentServicesOK() *GetEligibleShipmentServicesOK {
	return &GetEligibleShipmentServicesOK{}
}

/* GetEligibleShipmentServicesOK describes a response with status code 200, with default header values.

Success
*/
type GetEligibleShipmentServicesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOK) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesOK  %+v", 200, o.Payload)
}
func (o *GetEligibleShipmentServicesOK) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesBadRequest creates a GetEligibleShipmentServicesBadRequest with default headers values
func NewGetEligibleShipmentServicesBadRequest() *GetEligibleShipmentServicesBadRequest {
	return &GetEligibleShipmentServicesBadRequest{}
}

/* GetEligibleShipmentServicesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetEligibleShipmentServicesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetEligibleShipmentServicesBadRequest) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesUnauthorized creates a GetEligibleShipmentServicesUnauthorized with default headers values
func NewGetEligibleShipmentServicesUnauthorized() *GetEligibleShipmentServicesUnauthorized {
	return &GetEligibleShipmentServicesUnauthorized{}
}

/* GetEligibleShipmentServicesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetEligibleShipmentServicesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEligibleShipmentServicesUnauthorized) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesForbidden creates a GetEligibleShipmentServicesForbidden with default headers values
func NewGetEligibleShipmentServicesForbidden() *GetEligibleShipmentServicesForbidden {
	return &GetEligibleShipmentServicesForbidden{}
}

/* GetEligibleShipmentServicesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetEligibleShipmentServicesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesForbidden) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesForbidden  %+v", 403, o.Payload)
}
func (o *GetEligibleShipmentServicesForbidden) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesNotFound creates a GetEligibleShipmentServicesNotFound with default headers values
func NewGetEligibleShipmentServicesNotFound() *GetEligibleShipmentServicesNotFound {
	return &GetEligibleShipmentServicesNotFound{}
}

/* GetEligibleShipmentServicesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetEligibleShipmentServicesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesNotFound) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesNotFound  %+v", 404, o.Payload)
}
func (o *GetEligibleShipmentServicesNotFound) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesTooManyRequests creates a GetEligibleShipmentServicesTooManyRequests with default headers values
func NewGetEligibleShipmentServicesTooManyRequests() *GetEligibleShipmentServicesTooManyRequests {
	return &GetEligibleShipmentServicesTooManyRequests{}
}

/* GetEligibleShipmentServicesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetEligibleShipmentServicesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetEligibleShipmentServicesTooManyRequests) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesInternalServerError creates a GetEligibleShipmentServicesInternalServerError with default headers values
func NewGetEligibleShipmentServicesInternalServerError() *GetEligibleShipmentServicesInternalServerError {
	return &GetEligibleShipmentServicesInternalServerError{}
}

/* GetEligibleShipmentServicesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetEligibleShipmentServicesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEligibleShipmentServicesInternalServerError) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEligibleShipmentServicesServiceUnavailable creates a GetEligibleShipmentServicesServiceUnavailable with default headers values
func NewGetEligibleShipmentServicesServiceUnavailable() *GetEligibleShipmentServicesServiceUnavailable {
	return &GetEligibleShipmentServicesServiceUnavailable{}
}

/* GetEligibleShipmentServicesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetEligibleShipmentServicesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleShippingServices][%d] getEligibleShipmentServicesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetEligibleShipmentServicesServiceUnavailable) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}