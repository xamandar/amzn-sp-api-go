// Code generated by go-swagger; DO NOT EDIT.

package fees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/productFeesV0/product_fees_v0_models"
)

// GetMyFeesEstimatesReader is a Reader for the GetMyFeesEstimates structure.
type GetMyFeesEstimatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMyFeesEstimatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMyFeesEstimatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMyFeesEstimatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMyFeesEstimatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMyFeesEstimatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMyFeesEstimatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMyFeesEstimatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMyFeesEstimatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMyFeesEstimatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMyFeesEstimatesOK creates a GetMyFeesEstimatesOK with default headers values
func NewGetMyFeesEstimatesOK() *GetMyFeesEstimatesOK {
	return &GetMyFeesEstimatesOK{}
}

/* GetMyFeesEstimatesOK describes a response with status code 200, with default header values.

Success.
*/
type GetMyFeesEstimatesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload product_fees_v0_models.GetMyFeesEstimatesResponse
}

func (o *GetMyFeesEstimatesOK) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesOK  %+v", 200, o.Payload)
}
func (o *GetMyFeesEstimatesOK) GetPayload() product_fees_v0_models.GetMyFeesEstimatesResponse {
	return o.Payload
}

func (o *GetMyFeesEstimatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesBadRequest creates a GetMyFeesEstimatesBadRequest with default headers values
func NewGetMyFeesEstimatesBadRequest() *GetMyFeesEstimatesBadRequest {
	return &GetMyFeesEstimatesBadRequest{}
}

/* GetMyFeesEstimatesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetMyFeesEstimatesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesBadRequest  %+v", 400, o.Payload)
}
func (o *GetMyFeesEstimatesBadRequest) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesUnauthorized creates a GetMyFeesEstimatesUnauthorized with default headers values
func NewGetMyFeesEstimatesUnauthorized() *GetMyFeesEstimatesUnauthorized {
	return &GetMyFeesEstimatesUnauthorized{}
}

/* GetMyFeesEstimatesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetMyFeesEstimatesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetMyFeesEstimatesUnauthorized) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesForbidden creates a GetMyFeesEstimatesForbidden with default headers values
func NewGetMyFeesEstimatesForbidden() *GetMyFeesEstimatesForbidden {
	return &GetMyFeesEstimatesForbidden{}
}

/* GetMyFeesEstimatesForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include **Access Denied**, **Unauthorized**, **Expired Token**, or **Invalid Signature**.
*/
type GetMyFeesEstimatesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesForbidden) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesForbidden  %+v", 403, o.Payload)
}
func (o *GetMyFeesEstimatesForbidden) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesNotFound creates a GetMyFeesEstimatesNotFound with default headers values
func NewGetMyFeesEstimatesNotFound() *GetMyFeesEstimatesNotFound {
	return &GetMyFeesEstimatesNotFound{}
}

/* GetMyFeesEstimatesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetMyFeesEstimatesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesNotFound) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesNotFound  %+v", 404, o.Payload)
}
func (o *GetMyFeesEstimatesNotFound) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesTooManyRequests creates a GetMyFeesEstimatesTooManyRequests with default headers values
func NewGetMyFeesEstimatesTooManyRequests() *GetMyFeesEstimatesTooManyRequests {
	return &GetMyFeesEstimatesTooManyRequests{}
}

/* GetMyFeesEstimatesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetMyFeesEstimatesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetMyFeesEstimatesTooManyRequests) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesInternalServerError creates a GetMyFeesEstimatesInternalServerError with default headers values
func NewGetMyFeesEstimatesInternalServerError() *GetMyFeesEstimatesInternalServerError {
	return &GetMyFeesEstimatesInternalServerError{}
}

/* GetMyFeesEstimatesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetMyFeesEstimatesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetMyFeesEstimatesInternalServerError) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimatesServiceUnavailable creates a GetMyFeesEstimatesServiceUnavailable with default headers values
func NewGetMyFeesEstimatesServiceUnavailable() *GetMyFeesEstimatesServiceUnavailable {
	return &GetMyFeesEstimatesServiceUnavailable{}
}

/* GetMyFeesEstimatesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetMyFeesEstimatesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimatesErrorList
}

func (o *GetMyFeesEstimatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/feesEstimate][%d] getMyFeesEstimatesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetMyFeesEstimatesServiceUnavailable) GetPayload() *product_fees_v0_models.GetMyFeesEstimatesErrorList {
	return o.Payload
}

func (o *GetMyFeesEstimatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimatesErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}