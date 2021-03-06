// Code generated by go-swagger; DO NOT EDIT.

package aplus_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/aplusContent_2020-11-01/aplus_content_2020_11_01_models"
)

// UpdateContentDocumentReader is a Reader for the UpdateContentDocument structure.
type UpdateContentDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateContentDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateContentDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateContentDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateContentDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateContentDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateContentDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewUpdateContentDocumentGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateContentDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateContentDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateContentDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateContentDocumentOK creates a UpdateContentDocumentOK with default headers values
func NewUpdateContentDocumentOK() *UpdateContentDocumentOK {
	return &UpdateContentDocumentOK{}
}

/* UpdateContentDocumentOK describes a response with status code 200, with default header values.

Success.
*/
type UpdateContentDocumentOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.PostContentDocumentResponse
}

func (o *UpdateContentDocumentOK) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentOK  %+v", 200, o.Payload)
}
func (o *UpdateContentDocumentOK) GetPayload() *aplus_content_2020_11_01_models.PostContentDocumentResponse {
	return o.Payload
}

func (o *UpdateContentDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.PostContentDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentBadRequest creates a UpdateContentDocumentBadRequest with default headers values
func NewUpdateContentDocumentBadRequest() *UpdateContentDocumentBadRequest {
	return &UpdateContentDocumentBadRequest{}
}

/* UpdateContentDocumentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateContentDocumentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentBadRequest) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateContentDocumentBadRequest) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentUnauthorized creates a UpdateContentDocumentUnauthorized with default headers values
func NewUpdateContentDocumentUnauthorized() *UpdateContentDocumentUnauthorized {
	return &UpdateContentDocumentUnauthorized{}
}

/* UpdateContentDocumentUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type UpdateContentDocumentUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateContentDocumentUnauthorized) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentForbidden creates a UpdateContentDocumentForbidden with default headers values
func NewUpdateContentDocumentForbidden() *UpdateContentDocumentForbidden {
	return &UpdateContentDocumentForbidden{}
}

/* UpdateContentDocumentForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateContentDocumentForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentForbidden) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentForbidden  %+v", 403, o.Payload)
}
func (o *UpdateContentDocumentForbidden) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentNotFound creates a UpdateContentDocumentNotFound with default headers values
func NewUpdateContentDocumentNotFound() *UpdateContentDocumentNotFound {
	return &UpdateContentDocumentNotFound{}
}

/* UpdateContentDocumentNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type UpdateContentDocumentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentNotFound) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentNotFound  %+v", 404, o.Payload)
}
func (o *UpdateContentDocumentNotFound) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentGone creates a UpdateContentDocumentGone with default headers values
func NewUpdateContentDocumentGone() *UpdateContentDocumentGone {
	return &UpdateContentDocumentGone{}
}

/* UpdateContentDocumentGone describes a response with status code 410, with default header values.

The specified resource no longer exists.
*/
type UpdateContentDocumentGone struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentGone) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentGone  %+v", 410, o.Payload)
}
func (o *UpdateContentDocumentGone) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentTooManyRequests creates a UpdateContentDocumentTooManyRequests with default headers values
func NewUpdateContentDocumentTooManyRequests() *UpdateContentDocumentTooManyRequests {
	return &UpdateContentDocumentTooManyRequests{}
}

/* UpdateContentDocumentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateContentDocumentTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateContentDocumentTooManyRequests) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentInternalServerError creates a UpdateContentDocumentInternalServerError with default headers values
func NewUpdateContentDocumentInternalServerError() *UpdateContentDocumentInternalServerError {
	return &UpdateContentDocumentInternalServerError{}
}

/* UpdateContentDocumentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateContentDocumentInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateContentDocumentInternalServerError) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContentDocumentServiceUnavailable creates a UpdateContentDocumentServiceUnavailable with default headers values
func NewUpdateContentDocumentServiceUnavailable() *UpdateContentDocumentServiceUnavailable {
	return &UpdateContentDocumentServiceUnavailable{}
}

/* UpdateContentDocumentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateContentDocumentServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *UpdateContentDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentDocuments/{contentReferenceKey}][%d] updateContentDocumentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UpdateContentDocumentServiceUnavailable) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *UpdateContentDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
