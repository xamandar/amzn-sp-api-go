// Code generated by go-swagger; DO NOT EDIT.

package solicitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/solicitations/solicitations_models"
)

// CreateProductReviewAndSellerFeedbackSolicitationReader is a Reader for the CreateProductReviewAndSellerFeedbackSolicitation structure.
type CreateProductReviewAndSellerFeedbackSolicitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProductReviewAndSellerFeedbackSolicitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProductReviewAndSellerFeedbackSolicitationCreated creates a CreateProductReviewAndSellerFeedbackSolicitationCreated with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationCreated() *CreateProductReviewAndSellerFeedbackSolicitationCreated {
	return &CreateProductReviewAndSellerFeedbackSolicitationCreated{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type CreateProductReviewAndSellerFeedbackSolicitationCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationCreated) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationCreated  %+v", 201, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationCreated) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationBadRequest creates a CreateProductReviewAndSellerFeedbackSolicitationBadRequest with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationBadRequest() *CreateProductReviewAndSellerFeedbackSolicitationBadRequest {
	return &CreateProductReviewAndSellerFeedbackSolicitationBadRequest{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateProductReviewAndSellerFeedbackSolicitationBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationBadRequest) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationBadRequest  %+v", 400, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationBadRequest) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationForbidden creates a CreateProductReviewAndSellerFeedbackSolicitationForbidden with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationForbidden() *CreateProductReviewAndSellerFeedbackSolicitationForbidden {
	return &CreateProductReviewAndSellerFeedbackSolicitationForbidden{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateProductReviewAndSellerFeedbackSolicitationForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationForbidden) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationForbidden  %+v", 403, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationForbidden) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationNotFound creates a CreateProductReviewAndSellerFeedbackSolicitationNotFound with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationNotFound() *CreateProductReviewAndSellerFeedbackSolicitationNotFound {
	return &CreateProductReviewAndSellerFeedbackSolicitationNotFound{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateProductReviewAndSellerFeedbackSolicitationNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationNotFound) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationNotFound  %+v", 404, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationNotFound) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge creates a CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge() *CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge {
	return &CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType creates a CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType() *CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType {
	return &CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationTooManyRequests creates a CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationTooManyRequests() *CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests {
	return &CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationInternalServerError creates a CreateProductReviewAndSellerFeedbackSolicitationInternalServerError with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationInternalServerError() *CreateProductReviewAndSellerFeedbackSolicitationInternalServerError {
	return &CreateProductReviewAndSellerFeedbackSolicitationInternalServerError{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateProductReviewAndSellerFeedbackSolicitationInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationInternalServerError) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable creates a CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable with default headers values
func NewCreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable() *CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable {
	return &CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable{}
}

/* CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback][%d] createProductReviewAndSellerFeedbackSolicitationServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable) GetPayload() *solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse {
	return o.Payload
}

func (o *CreateProductReviewAndSellerFeedbackSolicitationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(solicitations_models.CreateProductReviewAndSellerFeedbackSolicitationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}