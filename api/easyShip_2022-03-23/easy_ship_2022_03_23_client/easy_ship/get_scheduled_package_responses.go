// Code generated by go-swagger; DO NOT EDIT.

package easy_ship

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/easyShip_2022-03-23/easy_ship_2022_03_23_models"
)

// GetScheduledPackageReader is a Reader for the GetScheduledPackage structure.
type GetScheduledPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduledPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduledPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScheduledPackageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScheduledPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScheduledPackageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScheduledPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScheduledPackageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScheduledPackageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScheduledPackageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScheduledPackageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScheduledPackageOK creates a GetScheduledPackageOK with default headers values
func NewGetScheduledPackageOK() *GetScheduledPackageOK {
	return &GetScheduledPackageOK{}
}

/* GetScheduledPackageOK describes a response with status code 200, with default header values.

Success.
*/
type GetScheduledPackageOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.Package
}

func (o *GetScheduledPackageOK) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageOK  %+v", 200, o.Payload)
}
func (o *GetScheduledPackageOK) GetPayload() *easy_ship_2022_03_23_models.Package {
	return o.Payload
}

func (o *GetScheduledPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageBadRequest creates a GetScheduledPackageBadRequest with default headers values
func NewGetScheduledPackageBadRequest() *GetScheduledPackageBadRequest {
	return &GetScheduledPackageBadRequest{}
}

/* GetScheduledPackageBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetScheduledPackageBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageBadRequest) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageBadRequest  %+v", 400, o.Payload)
}
func (o *GetScheduledPackageBadRequest) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageUnauthorized creates a GetScheduledPackageUnauthorized with default headers values
func NewGetScheduledPackageUnauthorized() *GetScheduledPackageUnauthorized {
	return &GetScheduledPackageUnauthorized{}
}

/* GetScheduledPackageUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetScheduledPackageUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageUnauthorized  %+v", 401, o.Payload)
}
func (o *GetScheduledPackageUnauthorized) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageForbidden creates a GetScheduledPackageForbidden with default headers values
func NewGetScheduledPackageForbidden() *GetScheduledPackageForbidden {
	return &GetScheduledPackageForbidden{}
}

/* GetScheduledPackageForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetScheduledPackageForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageForbidden) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageForbidden  %+v", 403, o.Payload)
}
func (o *GetScheduledPackageForbidden) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageNotFound creates a GetScheduledPackageNotFound with default headers values
func NewGetScheduledPackageNotFound() *GetScheduledPackageNotFound {
	return &GetScheduledPackageNotFound{}
}

/* GetScheduledPackageNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetScheduledPackageNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageNotFound) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageNotFound  %+v", 404, o.Payload)
}
func (o *GetScheduledPackageNotFound) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageUnsupportedMediaType creates a GetScheduledPackageUnsupportedMediaType with default headers values
func NewGetScheduledPackageUnsupportedMediaType() *GetScheduledPackageUnsupportedMediaType {
	return &GetScheduledPackageUnsupportedMediaType{}
}

/* GetScheduledPackageUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetScheduledPackageUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetScheduledPackageUnsupportedMediaType) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageTooManyRequests creates a GetScheduledPackageTooManyRequests with default headers values
func NewGetScheduledPackageTooManyRequests() *GetScheduledPackageTooManyRequests {
	return &GetScheduledPackageTooManyRequests{}
}

/* GetScheduledPackageTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetScheduledPackageTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetScheduledPackageTooManyRequests) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageInternalServerError creates a GetScheduledPackageInternalServerError with default headers values
func NewGetScheduledPackageInternalServerError() *GetScheduledPackageInternalServerError {
	return &GetScheduledPackageInternalServerError{}
}

/* GetScheduledPackageInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetScheduledPackageInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageInternalServerError  %+v", 500, o.Payload)
}
func (o *GetScheduledPackageInternalServerError) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScheduledPackageServiceUnavailable creates a GetScheduledPackageServiceUnavailable with default headers values
func NewGetScheduledPackageServiceUnavailable() *GetScheduledPackageServiceUnavailable {
	return &GetScheduledPackageServiceUnavailable{}
}

/* GetScheduledPackageServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetScheduledPackageServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *GetScheduledPackageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /easyShip/2022-03-23/package][%d] getScheduledPackageServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetScheduledPackageServiceUnavailable) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *GetScheduledPackageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
