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

// GetEligibleShipmentServicesOldReader is a Reader for the GetEligibleShipmentServicesOld structure.
type GetEligibleShipmentServicesOldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEligibleShipmentServicesOldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEligibleShipmentServicesOldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEligibleShipmentServicesOldBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEligibleShipmentServicesOldUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEligibleShipmentServicesOldForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEligibleShipmentServicesOldNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEligibleShipmentServicesOldTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEligibleShipmentServicesOldInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetEligibleShipmentServicesOldServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEligibleShipmentServicesOldOK creates a GetEligibleShipmentServicesOldOK with default headers values
func NewGetEligibleShipmentServicesOldOK() *GetEligibleShipmentServicesOldOK {
	return &GetEligibleShipmentServicesOldOK{}
}

/* GetEligibleShipmentServicesOldOK describes a response with status code 200, with default header values.

Success
*/
type GetEligibleShipmentServicesOldOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldOK) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldOK  %+v", 200, o.Payload)
}
func (o *GetEligibleShipmentServicesOldOK) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldBadRequest creates a GetEligibleShipmentServicesOldBadRequest with default headers values
func NewGetEligibleShipmentServicesOldBadRequest() *GetEligibleShipmentServicesOldBadRequest {
	return &GetEligibleShipmentServicesOldBadRequest{}
}

/* GetEligibleShipmentServicesOldBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetEligibleShipmentServicesOldBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldBadRequest) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldBadRequest  %+v", 400, o.Payload)
}
func (o *GetEligibleShipmentServicesOldBadRequest) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldUnauthorized creates a GetEligibleShipmentServicesOldUnauthorized with default headers values
func NewGetEligibleShipmentServicesOldUnauthorized() *GetEligibleShipmentServicesOldUnauthorized {
	return &GetEligibleShipmentServicesOldUnauthorized{}
}

/* GetEligibleShipmentServicesOldUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetEligibleShipmentServicesOldUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldUnauthorized) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEligibleShipmentServicesOldUnauthorized) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldForbidden creates a GetEligibleShipmentServicesOldForbidden with default headers values
func NewGetEligibleShipmentServicesOldForbidden() *GetEligibleShipmentServicesOldForbidden {
	return &GetEligibleShipmentServicesOldForbidden{}
}

/* GetEligibleShipmentServicesOldForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetEligibleShipmentServicesOldForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldForbidden) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldForbidden  %+v", 403, o.Payload)
}
func (o *GetEligibleShipmentServicesOldForbidden) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldNotFound creates a GetEligibleShipmentServicesOldNotFound with default headers values
func NewGetEligibleShipmentServicesOldNotFound() *GetEligibleShipmentServicesOldNotFound {
	return &GetEligibleShipmentServicesOldNotFound{}
}

/* GetEligibleShipmentServicesOldNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetEligibleShipmentServicesOldNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldNotFound) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldNotFound  %+v", 404, o.Payload)
}
func (o *GetEligibleShipmentServicesOldNotFound) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldTooManyRequests creates a GetEligibleShipmentServicesOldTooManyRequests with default headers values
func NewGetEligibleShipmentServicesOldTooManyRequests() *GetEligibleShipmentServicesOldTooManyRequests {
	return &GetEligibleShipmentServicesOldTooManyRequests{}
}

/* GetEligibleShipmentServicesOldTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetEligibleShipmentServicesOldTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetEligibleShipmentServicesOldTooManyRequests) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldInternalServerError creates a GetEligibleShipmentServicesOldInternalServerError with default headers values
func NewGetEligibleShipmentServicesOldInternalServerError() *GetEligibleShipmentServicesOldInternalServerError {
	return &GetEligibleShipmentServicesOldInternalServerError{}
}

/* GetEligibleShipmentServicesOldInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetEligibleShipmentServicesOldInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldInternalServerError) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEligibleShipmentServicesOldInternalServerError) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEligibleShipmentServicesOldServiceUnavailable creates a GetEligibleShipmentServicesOldServiceUnavailable with default headers values
func NewGetEligibleShipmentServicesOldServiceUnavailable() *GetEligibleShipmentServicesOldServiceUnavailable {
	return &GetEligibleShipmentServicesOldServiceUnavailable{}
}

/* GetEligibleShipmentServicesOldServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetEligibleShipmentServicesOldServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse
}

func (o *GetEligibleShipmentServicesOldServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /mfn/v0/eligibleServices][%d] getEligibleShipmentServicesOldServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetEligibleShipmentServicesOldServiceUnavailable) GetPayload() *merchant_fulfillment_v0_models.GetEligibleShipmentServicesResponse {
	return o.Payload
}

func (o *GetEligibleShipmentServicesOldServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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