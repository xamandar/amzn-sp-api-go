// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/services/services_models"
)

// CancelServiceJobByServiceJobIDReader is a Reader for the CancelServiceJobByServiceJobID structure.
type CancelServiceJobByServiceJobIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelServiceJobByServiceJobIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelServiceJobByServiceJobIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelServiceJobByServiceJobIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelServiceJobByServiceJobIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelServiceJobByServiceJobIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCancelServiceJobByServiceJobIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCancelServiceJobByServiceJobIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCancelServiceJobByServiceJobIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCancelServiceJobByServiceJobIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelServiceJobByServiceJobIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelServiceJobByServiceJobIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelServiceJobByServiceJobIDOK creates a CancelServiceJobByServiceJobIDOK with default headers values
func NewCancelServiceJobByServiceJobIDOK() *CancelServiceJobByServiceJobIDOK {
	return &CancelServiceJobByServiceJobIDOK{}
}

/* CancelServiceJobByServiceJobIDOK describes a response with status code 200, with default header values.

Success response
*/
type CancelServiceJobByServiceJobIDOK struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDOK) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdOK  %+v", 200, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDOK) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDBadRequest creates a CancelServiceJobByServiceJobIDBadRequest with default headers values
func NewCancelServiceJobByServiceJobIDBadRequest() *CancelServiceJobByServiceJobIDBadRequest {
	return &CancelServiceJobByServiceJobIDBadRequest{}
}

/* CancelServiceJobByServiceJobIDBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CancelServiceJobByServiceJobIDBadRequest struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdBadRequest  %+v", 400, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDBadRequest) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDForbidden creates a CancelServiceJobByServiceJobIDForbidden with default headers values
func NewCancelServiceJobByServiceJobIDForbidden() *CancelServiceJobByServiceJobIDForbidden {
	return &CancelServiceJobByServiceJobIDForbidden{}
}

/* CancelServiceJobByServiceJobIDForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CancelServiceJobByServiceJobIDForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdForbidden  %+v", 403, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDForbidden) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDNotFound creates a CancelServiceJobByServiceJobIDNotFound with default headers values
func NewCancelServiceJobByServiceJobIDNotFound() *CancelServiceJobByServiceJobIDNotFound {
	return &CancelServiceJobByServiceJobIDNotFound{}
}

/* CancelServiceJobByServiceJobIDNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CancelServiceJobByServiceJobIDNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdNotFound  %+v", 404, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDNotFound) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDRequestEntityTooLarge creates a CancelServiceJobByServiceJobIDRequestEntityTooLarge with default headers values
func NewCancelServiceJobByServiceJobIDRequestEntityTooLarge() *CancelServiceJobByServiceJobIDRequestEntityTooLarge {
	return &CancelServiceJobByServiceJobIDRequestEntityTooLarge{}
}

/* CancelServiceJobByServiceJobIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CancelServiceJobByServiceJobIDRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDRequestEntityTooLarge) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDUnsupportedMediaType creates a CancelServiceJobByServiceJobIDUnsupportedMediaType with default headers values
func NewCancelServiceJobByServiceJobIDUnsupportedMediaType() *CancelServiceJobByServiceJobIDUnsupportedMediaType {
	return &CancelServiceJobByServiceJobIDUnsupportedMediaType{}
}

/* CancelServiceJobByServiceJobIDUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CancelServiceJobByServiceJobIDUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDUnsupportedMediaType) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDUnprocessableEntity creates a CancelServiceJobByServiceJobIDUnprocessableEntity with default headers values
func NewCancelServiceJobByServiceJobIDUnprocessableEntity() *CancelServiceJobByServiceJobIDUnprocessableEntity {
	return &CancelServiceJobByServiceJobIDUnprocessableEntity{}
}

/* CancelServiceJobByServiceJobIDUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity. Unable to process the contained instructions
*/
type CancelServiceJobByServiceJobIDUnprocessableEntity struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDUnprocessableEntity) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDTooManyRequests creates a CancelServiceJobByServiceJobIDTooManyRequests with default headers values
func NewCancelServiceJobByServiceJobIDTooManyRequests() *CancelServiceJobByServiceJobIDTooManyRequests {
	return &CancelServiceJobByServiceJobIDTooManyRequests{}
}

/* CancelServiceJobByServiceJobIDTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CancelServiceJobByServiceJobIDTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDTooManyRequests) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDInternalServerError creates a CancelServiceJobByServiceJobIDInternalServerError with default headers values
func NewCancelServiceJobByServiceJobIDInternalServerError() *CancelServiceJobByServiceJobIDInternalServerError {
	return &CancelServiceJobByServiceJobIDInternalServerError{}
}

/* CancelServiceJobByServiceJobIDInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CancelServiceJobByServiceJobIDInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdInternalServerError  %+v", 500, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDInternalServerError) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelServiceJobByServiceJobIDServiceUnavailable creates a CancelServiceJobByServiceJobIDServiceUnavailable with default headers values
func NewCancelServiceJobByServiceJobIDServiceUnavailable() *CancelServiceJobByServiceJobIDServiceUnavailable {
	return &CancelServiceJobByServiceJobIDServiceUnavailable{}
}

/* CancelServiceJobByServiceJobIDServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CancelServiceJobByServiceJobIDServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CancelServiceJobByServiceJobIDResponse
}

func (o *CancelServiceJobByServiceJobIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/cancellations][%d] cancelServiceJobByServiceJobIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CancelServiceJobByServiceJobIDServiceUnavailable) GetPayload() *services_models.CancelServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CancelServiceJobByServiceJobIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CancelServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
