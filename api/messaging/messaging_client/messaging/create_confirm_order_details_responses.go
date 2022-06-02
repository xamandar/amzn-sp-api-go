// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/messaging/messaging_models"
)

// CreateConfirmOrderDetailsReader is a Reader for the CreateConfirmOrderDetails structure.
type CreateConfirmOrderDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConfirmOrderDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConfirmOrderDetailsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConfirmOrderDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateConfirmOrderDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateConfirmOrderDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateConfirmOrderDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateConfirmOrderDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateConfirmOrderDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateConfirmOrderDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateConfirmOrderDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConfirmOrderDetailsCreated creates a CreateConfirmOrderDetailsCreated with default headers values
func NewCreateConfirmOrderDetailsCreated() *CreateConfirmOrderDetailsCreated {
	return &CreateConfirmOrderDetailsCreated{}
}

/* CreateConfirmOrderDetailsCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type CreateConfirmOrderDetailsCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsCreated  %+v", 201, o.Payload)
}
func (o *CreateConfirmOrderDetailsCreated) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsBadRequest creates a CreateConfirmOrderDetailsBadRequest with default headers values
func NewCreateConfirmOrderDetailsBadRequest() *CreateConfirmOrderDetailsBadRequest {
	return &CreateConfirmOrderDetailsBadRequest{}
}

/* CreateConfirmOrderDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateConfirmOrderDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateConfirmOrderDetailsBadRequest) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsForbidden creates a CreateConfirmOrderDetailsForbidden with default headers values
func NewCreateConfirmOrderDetailsForbidden() *CreateConfirmOrderDetailsForbidden {
	return &CreateConfirmOrderDetailsForbidden{}
}

/* CreateConfirmOrderDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateConfirmOrderDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsForbidden  %+v", 403, o.Payload)
}
func (o *CreateConfirmOrderDetailsForbidden) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsNotFound creates a CreateConfirmOrderDetailsNotFound with default headers values
func NewCreateConfirmOrderDetailsNotFound() *CreateConfirmOrderDetailsNotFound {
	return &CreateConfirmOrderDetailsNotFound{}
}

/* CreateConfirmOrderDetailsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateConfirmOrderDetailsNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsNotFound  %+v", 404, o.Payload)
}
func (o *CreateConfirmOrderDetailsNotFound) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsRequestEntityTooLarge creates a CreateConfirmOrderDetailsRequestEntityTooLarge with default headers values
func NewCreateConfirmOrderDetailsRequestEntityTooLarge() *CreateConfirmOrderDetailsRequestEntityTooLarge {
	return &CreateConfirmOrderDetailsRequestEntityTooLarge{}
}

/* CreateConfirmOrderDetailsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateConfirmOrderDetailsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *CreateConfirmOrderDetailsRequestEntityTooLarge) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsUnsupportedMediaType creates a CreateConfirmOrderDetailsUnsupportedMediaType with default headers values
func NewCreateConfirmOrderDetailsUnsupportedMediaType() *CreateConfirmOrderDetailsUnsupportedMediaType {
	return &CreateConfirmOrderDetailsUnsupportedMediaType{}
}

/* CreateConfirmOrderDetailsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateConfirmOrderDetailsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateConfirmOrderDetailsUnsupportedMediaType) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsTooManyRequests creates a CreateConfirmOrderDetailsTooManyRequests with default headers values
func NewCreateConfirmOrderDetailsTooManyRequests() *CreateConfirmOrderDetailsTooManyRequests {
	return &CreateConfirmOrderDetailsTooManyRequests{}
}

/* CreateConfirmOrderDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateConfirmOrderDetailsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateConfirmOrderDetailsTooManyRequests) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsInternalServerError creates a CreateConfirmOrderDetailsInternalServerError with default headers values
func NewCreateConfirmOrderDetailsInternalServerError() *CreateConfirmOrderDetailsInternalServerError {
	return &CreateConfirmOrderDetailsInternalServerError{}
}

/* CreateConfirmOrderDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateConfirmOrderDetailsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateConfirmOrderDetailsInternalServerError) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmOrderDetailsServiceUnavailable creates a CreateConfirmOrderDetailsServiceUnavailable with default headers values
func NewCreateConfirmOrderDetailsServiceUnavailable() *CreateConfirmOrderDetailsServiceUnavailable {
	return &CreateConfirmOrderDetailsServiceUnavailable{}
}

/* CreateConfirmOrderDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateConfirmOrderDetailsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateConfirmOrderDetailsResponse
}

func (o *CreateConfirmOrderDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails][%d] createConfirmOrderDetailsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateConfirmOrderDetailsServiceUnavailable) GetPayload() *messaging_models.CreateConfirmOrderDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmOrderDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateConfirmOrderDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}