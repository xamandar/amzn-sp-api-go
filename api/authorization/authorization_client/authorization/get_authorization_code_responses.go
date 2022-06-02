// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/authorization/authorization_models"
)

// GetAuthorizationCodeReader is a Reader for the GetAuthorizationCode structure.
type GetAuthorizationCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationCodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationCodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationCodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationCodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationCodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationCodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationCodeOK creates a GetAuthorizationCodeOK with default headers values
func NewGetAuthorizationCodeOK() *GetAuthorizationCodeOK {
	return &GetAuthorizationCodeOK{}
}

/* GetAuthorizationCodeOK describes a response with status code 200, with default header values.

Success.
*/
type GetAuthorizationCodeOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeOK) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeOK  %+v", 200, o.Payload)
}
func (o *GetAuthorizationCodeOK) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeBadRequest creates a GetAuthorizationCodeBadRequest with default headers values
func NewGetAuthorizationCodeBadRequest() *GetAuthorizationCodeBadRequest {
	return &GetAuthorizationCodeBadRequest{}
}

/* GetAuthorizationCodeBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetAuthorizationCodeBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeBadRequest  %+v", 400, o.Payload)
}
func (o *GetAuthorizationCodeBadRequest) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeForbidden creates a GetAuthorizationCodeForbidden with default headers values
func NewGetAuthorizationCodeForbidden() *GetAuthorizationCodeForbidden {
	return &GetAuthorizationCodeForbidden{}
}

/* GetAuthorizationCodeForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetAuthorizationCodeForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeForbidden) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeForbidden  %+v", 403, o.Payload)
}
func (o *GetAuthorizationCodeForbidden) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeNotFound creates a GetAuthorizationCodeNotFound with default headers values
func NewGetAuthorizationCodeNotFound() *GetAuthorizationCodeNotFound {
	return &GetAuthorizationCodeNotFound{}
}

/* GetAuthorizationCodeNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetAuthorizationCodeNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeNotFound  %+v", 404, o.Payload)
}
func (o *GetAuthorizationCodeNotFound) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeRequestEntityTooLarge creates a GetAuthorizationCodeRequestEntityTooLarge with default headers values
func NewGetAuthorizationCodeRequestEntityTooLarge() *GetAuthorizationCodeRequestEntityTooLarge {
	return &GetAuthorizationCodeRequestEntityTooLarge{}
}

/* GetAuthorizationCodeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetAuthorizationCodeRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *GetAuthorizationCodeRequestEntityTooLarge) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeUnsupportedMediaType creates a GetAuthorizationCodeUnsupportedMediaType with default headers values
func NewGetAuthorizationCodeUnsupportedMediaType() *GetAuthorizationCodeUnsupportedMediaType {
	return &GetAuthorizationCodeUnsupportedMediaType{}
}

/* GetAuthorizationCodeUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetAuthorizationCodeUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetAuthorizationCodeUnsupportedMediaType) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeTooManyRequests creates a GetAuthorizationCodeTooManyRequests with default headers values
func NewGetAuthorizationCodeTooManyRequests() *GetAuthorizationCodeTooManyRequests {
	return &GetAuthorizationCodeTooManyRequests{}
}

/* GetAuthorizationCodeTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetAuthorizationCodeTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetAuthorizationCodeTooManyRequests) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeInternalServerError creates a GetAuthorizationCodeInternalServerError with default headers values
func NewGetAuthorizationCodeInternalServerError() *GetAuthorizationCodeInternalServerError {
	return &GetAuthorizationCodeInternalServerError{}
}

/* GetAuthorizationCodeInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetAuthorizationCodeInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAuthorizationCodeInternalServerError) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationCodeServiceUnavailable creates a GetAuthorizationCodeServiceUnavailable with default headers values
func NewGetAuthorizationCodeServiceUnavailable() *GetAuthorizationCodeServiceUnavailable {
	return &GetAuthorizationCodeServiceUnavailable{}
}

/* GetAuthorizationCodeServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetAuthorizationCodeServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *authorization_models.GetAuthorizationCodeResponse
}

func (o *GetAuthorizationCodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /authorization/v1/authorizationCode][%d] getAuthorizationCodeServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAuthorizationCodeServiceUnavailable) GetPayload() *authorization_models.GetAuthorizationCodeResponse {
	return o.Payload
}

func (o *GetAuthorizationCodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(authorization_models.GetAuthorizationCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}