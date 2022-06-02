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

// GetFeaturesReader is a Reader for the GetFeatures structure.
type GetFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFeaturesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFeaturesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFeaturesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFeaturesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFeaturesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFeaturesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFeaturesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFeaturesOK creates a GetFeaturesOK with default headers values
func NewGetFeaturesOK() *GetFeaturesOK {
	return &GetFeaturesOK{}
}

/* GetFeaturesOK describes a response with status code 200, with default header values.

Success.
*/
type GetFeaturesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesOK  %+v", 200, o.Payload)
}
func (o *GetFeaturesOK) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesBadRequest creates a GetFeaturesBadRequest with default headers values
func NewGetFeaturesBadRequest() *GetFeaturesBadRequest {
	return &GetFeaturesBadRequest{}
}

/* GetFeaturesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFeaturesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesBadRequest  %+v", 400, o.Payload)
}
func (o *GetFeaturesBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesUnauthorized creates a GetFeaturesUnauthorized with default headers values
func NewGetFeaturesUnauthorized() *GetFeaturesUnauthorized {
	return &GetFeaturesUnauthorized{}
}

/* GetFeaturesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFeaturesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFeaturesUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesForbidden creates a GetFeaturesForbidden with default headers values
func NewGetFeaturesForbidden() *GetFeaturesForbidden {
	return &GetFeaturesForbidden{}
}

/* GetFeaturesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFeaturesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesForbidden  %+v", 403, o.Payload)
}
func (o *GetFeaturesForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesNotFound creates a GetFeaturesNotFound with default headers values
func NewGetFeaturesNotFound() *GetFeaturesNotFound {
	return &GetFeaturesNotFound{}
}

/* GetFeaturesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetFeaturesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesNotFound  %+v", 404, o.Payload)
}
func (o *GetFeaturesNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesTooManyRequests creates a GetFeaturesTooManyRequests with default headers values
func NewGetFeaturesTooManyRequests() *GetFeaturesTooManyRequests {
	return &GetFeaturesTooManyRequests{}
}

/* GetFeaturesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFeaturesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFeaturesTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesInternalServerError creates a GetFeaturesInternalServerError with default headers values
func NewGetFeaturesInternalServerError() *GetFeaturesInternalServerError {
	return &GetFeaturesInternalServerError{}
}

/* GetFeaturesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFeaturesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFeaturesInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeaturesServiceUnavailable creates a GetFeaturesServiceUnavailable with default headers values
func NewGetFeaturesServiceUnavailable() *GetFeaturesServiceUnavailable {
	return &GetFeaturesServiceUnavailable{}
}

/* GetFeaturesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFeaturesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse
}

func (o *GetFeaturesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/features][%d] getFeaturesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFeaturesServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFeaturesResponse {
	return o.Payload
}

func (o *GetFeaturesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFeaturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}