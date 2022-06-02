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

// CreateDigitalAccessKeyReader is a Reader for the CreateDigitalAccessKey structure.
type CreateDigitalAccessKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDigitalAccessKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDigitalAccessKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDigitalAccessKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDigitalAccessKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDigitalAccessKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateDigitalAccessKeyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDigitalAccessKeyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDigitalAccessKeyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDigitalAccessKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateDigitalAccessKeyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDigitalAccessKeyCreated creates a CreateDigitalAccessKeyCreated with default headers values
func NewCreateDigitalAccessKeyCreated() *CreateDigitalAccessKeyCreated {
	return &CreateDigitalAccessKeyCreated{}
}

/* CreateDigitalAccessKeyCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type CreateDigitalAccessKeyCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyCreated  %+v", 201, o.Payload)
}
func (o *CreateDigitalAccessKeyCreated) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyBadRequest creates a CreateDigitalAccessKeyBadRequest with default headers values
func NewCreateDigitalAccessKeyBadRequest() *CreateDigitalAccessKeyBadRequest {
	return &CreateDigitalAccessKeyBadRequest{}
}

/* CreateDigitalAccessKeyBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateDigitalAccessKeyBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDigitalAccessKeyBadRequest) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyForbidden creates a CreateDigitalAccessKeyForbidden with default headers values
func NewCreateDigitalAccessKeyForbidden() *CreateDigitalAccessKeyForbidden {
	return &CreateDigitalAccessKeyForbidden{}
}

/* CreateDigitalAccessKeyForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateDigitalAccessKeyForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyForbidden  %+v", 403, o.Payload)
}
func (o *CreateDigitalAccessKeyForbidden) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyNotFound creates a CreateDigitalAccessKeyNotFound with default headers values
func NewCreateDigitalAccessKeyNotFound() *CreateDigitalAccessKeyNotFound {
	return &CreateDigitalAccessKeyNotFound{}
}

/* CreateDigitalAccessKeyNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateDigitalAccessKeyNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyNotFound  %+v", 404, o.Payload)
}
func (o *CreateDigitalAccessKeyNotFound) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-requestid
	hdrXAmznRequestid := response.GetHeader("x-amzn-requestid")

	if hdrXAmznRequestid != "" {
		o.XAmznRequestid = hdrXAmznRequestid
	}

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyRequestEntityTooLarge creates a CreateDigitalAccessKeyRequestEntityTooLarge with default headers values
func NewCreateDigitalAccessKeyRequestEntityTooLarge() *CreateDigitalAccessKeyRequestEntityTooLarge {
	return &CreateDigitalAccessKeyRequestEntityTooLarge{}
}

/* CreateDigitalAccessKeyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateDigitalAccessKeyRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *CreateDigitalAccessKeyRequestEntityTooLarge) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyUnsupportedMediaType creates a CreateDigitalAccessKeyUnsupportedMediaType with default headers values
func NewCreateDigitalAccessKeyUnsupportedMediaType() *CreateDigitalAccessKeyUnsupportedMediaType {
	return &CreateDigitalAccessKeyUnsupportedMediaType{}
}

/* CreateDigitalAccessKeyUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateDigitalAccessKeyUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateDigitalAccessKeyUnsupportedMediaType) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyTooManyRequests creates a CreateDigitalAccessKeyTooManyRequests with default headers values
func NewCreateDigitalAccessKeyTooManyRequests() *CreateDigitalAccessKeyTooManyRequests {
	return &CreateDigitalAccessKeyTooManyRequests{}
}

/* CreateDigitalAccessKeyTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateDigitalAccessKeyTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateDigitalAccessKeyTooManyRequests) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyInternalServerError creates a CreateDigitalAccessKeyInternalServerError with default headers values
func NewCreateDigitalAccessKeyInternalServerError() *CreateDigitalAccessKeyInternalServerError {
	return &CreateDigitalAccessKeyInternalServerError{}
}

/* CreateDigitalAccessKeyInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateDigitalAccessKeyInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDigitalAccessKeyInternalServerError) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigitalAccessKeyServiceUnavailable creates a CreateDigitalAccessKeyServiceUnavailable with default headers values
func NewCreateDigitalAccessKeyServiceUnavailable() *CreateDigitalAccessKeyServiceUnavailable {
	return &CreateDigitalAccessKeyServiceUnavailable{}
}

/* CreateDigitalAccessKeyServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateDigitalAccessKeyServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestid string

	Payload *messaging_models.CreateDigitalAccessKeyResponse
}

func (o *CreateDigitalAccessKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey][%d] createDigitalAccessKeyServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateDigitalAccessKeyServiceUnavailable) GetPayload() *messaging_models.CreateDigitalAccessKeyResponse {
	return o.Payload
}

func (o *CreateDigitalAccessKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateDigitalAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}