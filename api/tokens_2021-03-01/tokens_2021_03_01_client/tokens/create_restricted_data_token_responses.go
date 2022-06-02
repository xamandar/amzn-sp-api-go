// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/tokens_2021-03-01/tokens_2021_03_01_models"
)

// CreateRestrictedDataTokenReader is a Reader for the CreateRestrictedDataToken structure.
type CreateRestrictedDataTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRestrictedDataTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRestrictedDataTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRestrictedDataTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRestrictedDataTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRestrictedDataTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRestrictedDataTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateRestrictedDataTokenUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateRestrictedDataTokenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRestrictedDataTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateRestrictedDataTokenServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRestrictedDataTokenOK creates a CreateRestrictedDataTokenOK with default headers values
func NewCreateRestrictedDataTokenOK() *CreateRestrictedDataTokenOK {
	return &CreateRestrictedDataTokenOK{}
}

/* CreateRestrictedDataTokenOK describes a response with status code 200, with default header values.

Success.
*/
type CreateRestrictedDataTokenOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.CreateRestrictedDataTokenResponse
}

func (o *CreateRestrictedDataTokenOK) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenOK  %+v", 200, o.Payload)
}
func (o *CreateRestrictedDataTokenOK) GetPayload() *tokens_2021_03_01_models.CreateRestrictedDataTokenResponse {
	return o.Payload
}

func (o *CreateRestrictedDataTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.CreateRestrictedDataTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenBadRequest creates a CreateRestrictedDataTokenBadRequest with default headers values
func NewCreateRestrictedDataTokenBadRequest() *CreateRestrictedDataTokenBadRequest {
	return &CreateRestrictedDataTokenBadRequest{}
}

/* CreateRestrictedDataTokenBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateRestrictedDataTokenBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRestrictedDataTokenBadRequest) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenUnauthorized creates a CreateRestrictedDataTokenUnauthorized with default headers values
func NewCreateRestrictedDataTokenUnauthorized() *CreateRestrictedDataTokenUnauthorized {
	return &CreateRestrictedDataTokenUnauthorized{}
}

/* CreateRestrictedDataTokenUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CreateRestrictedDataTokenUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateRestrictedDataTokenUnauthorized) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenForbidden creates a CreateRestrictedDataTokenForbidden with default headers values
func NewCreateRestrictedDataTokenForbidden() *CreateRestrictedDataTokenForbidden {
	return &CreateRestrictedDataTokenForbidden{}
}

/* CreateRestrictedDataTokenForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateRestrictedDataTokenForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenForbidden  %+v", 403, o.Payload)
}
func (o *CreateRestrictedDataTokenForbidden) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenNotFound creates a CreateRestrictedDataTokenNotFound with default headers values
func NewCreateRestrictedDataTokenNotFound() *CreateRestrictedDataTokenNotFound {
	return &CreateRestrictedDataTokenNotFound{}
}

/* CreateRestrictedDataTokenNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CreateRestrictedDataTokenNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenNotFound  %+v", 404, o.Payload)
}
func (o *CreateRestrictedDataTokenNotFound) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenUnsupportedMediaType creates a CreateRestrictedDataTokenUnsupportedMediaType with default headers values
func NewCreateRestrictedDataTokenUnsupportedMediaType() *CreateRestrictedDataTokenUnsupportedMediaType {
	return &CreateRestrictedDataTokenUnsupportedMediaType{}
}

/* CreateRestrictedDataTokenUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateRestrictedDataTokenUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateRestrictedDataTokenUnsupportedMediaType) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenTooManyRequests creates a CreateRestrictedDataTokenTooManyRequests with default headers values
func NewCreateRestrictedDataTokenTooManyRequests() *CreateRestrictedDataTokenTooManyRequests {
	return &CreateRestrictedDataTokenTooManyRequests{}
}

/* CreateRestrictedDataTokenTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateRestrictedDataTokenTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateRestrictedDataTokenTooManyRequests) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenInternalServerError creates a CreateRestrictedDataTokenInternalServerError with default headers values
func NewCreateRestrictedDataTokenInternalServerError() *CreateRestrictedDataTokenInternalServerError {
	return &CreateRestrictedDataTokenInternalServerError{}
}

/* CreateRestrictedDataTokenInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateRestrictedDataTokenInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRestrictedDataTokenInternalServerError) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRestrictedDataTokenServiceUnavailable creates a CreateRestrictedDataTokenServiceUnavailable with default headers values
func NewCreateRestrictedDataTokenServiceUnavailable() *CreateRestrictedDataTokenServiceUnavailable {
	return &CreateRestrictedDataTokenServiceUnavailable{}
}

/* CreateRestrictedDataTokenServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateRestrictedDataTokenServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *tokens_2021_03_01_models.ErrorList
}

func (o *CreateRestrictedDataTokenServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /tokens/2021-03-01/restrictedDataToken][%d] createRestrictedDataTokenServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateRestrictedDataTokenServiceUnavailable) GetPayload() *tokens_2021_03_01_models.ErrorList {
	return o.Payload
}

func (o *CreateRestrictedDataTokenServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(tokens_2021_03_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}