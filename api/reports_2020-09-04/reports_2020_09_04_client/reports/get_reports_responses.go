// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/reports_2020-09-04/reports_2020_09_04_models"
)

// GetReportsReader is a Reader for the GetReports structure.
type GetReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetReportsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReportsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReportsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportsOK creates a GetReportsOK with default headers values
func NewGetReportsOK() *GetReportsOK {
	return &GetReportsOK{}
}

/* GetReportsOK describes a response with status code 200, with default header values.

Success.
*/
type GetReportsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsOK) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsOK  %+v", 200, o.Payload)
}
func (o *GetReportsOK) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsBadRequest creates a GetReportsBadRequest with default headers values
func NewGetReportsBadRequest() *GetReportsBadRequest {
	return &GetReportsBadRequest{}
}

/* GetReportsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetReportsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsBadRequest  %+v", 400, o.Payload)
}
func (o *GetReportsBadRequest) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsUnauthorized creates a GetReportsUnauthorized with default headers values
func NewGetReportsUnauthorized() *GetReportsUnauthorized {
	return &GetReportsUnauthorized{}
}

/* GetReportsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetReportsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReportsUnauthorized) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsForbidden creates a GetReportsForbidden with default headers values
func NewGetReportsForbidden() *GetReportsForbidden {
	return &GetReportsForbidden{}
}

/* GetReportsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetReportsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsForbidden  %+v", 403, o.Payload)
}
func (o *GetReportsForbidden) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsNotFound creates a GetReportsNotFound with default headers values
func NewGetReportsNotFound() *GetReportsNotFound {
	return &GetReportsNotFound{}
}

/* GetReportsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetReportsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsNotFound) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsNotFound  %+v", 404, o.Payload)
}
func (o *GetReportsNotFound) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsUnsupportedMediaType creates a GetReportsUnsupportedMediaType with default headers values
func NewGetReportsUnsupportedMediaType() *GetReportsUnsupportedMediaType {
	return &GetReportsUnsupportedMediaType{}
}

/* GetReportsUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetReportsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetReportsUnsupportedMediaType) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsTooManyRequests creates a GetReportsTooManyRequests with default headers values
func NewGetReportsTooManyRequests() *GetReportsTooManyRequests {
	return &GetReportsTooManyRequests{}
}

/* GetReportsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetReportsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetReportsTooManyRequests) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsInternalServerError creates a GetReportsInternalServerError with default headers values
func NewGetReportsInternalServerError() *GetReportsInternalServerError {
	return &GetReportsInternalServerError{}
}

/* GetReportsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetReportsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReportsInternalServerError) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsServiceUnavailable creates a GetReportsServiceUnavailable with default headers values
func NewGetReportsServiceUnavailable() *GetReportsServiceUnavailable {
	return &GetReportsServiceUnavailable{}
}

/* GetReportsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetReportsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportsResponse
}

func (o *GetReportsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/reports][%d] getReportsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetReportsServiceUnavailable) GetPayload() *reports_2020_09_04_models.GetReportsResponse {
	return o.Payload
}

func (o *GetReportsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
