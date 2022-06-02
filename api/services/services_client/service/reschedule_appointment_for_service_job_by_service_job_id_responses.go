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

// RescheduleAppointmentForServiceJobByServiceJobIDReader is a Reader for the RescheduleAppointmentForServiceJobByServiceJobID structure.
type RescheduleAppointmentForServiceJobByServiceJobIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RescheduleAppointmentForServiceJobByServiceJobIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewRescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDOK creates a RescheduleAppointmentForServiceJobByServiceJobIDOK with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDOK() *RescheduleAppointmentForServiceJobByServiceJobIDOK {
	return &RescheduleAppointmentForServiceJobByServiceJobIDOK{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDOK describes a response with status code 200, with default header values.

Success response
*/
type RescheduleAppointmentForServiceJobByServiceJobIDOK struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDOK) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdOK  %+v", 200, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDOK) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDBadRequest creates a RescheduleAppointmentForServiceJobByServiceJobIDBadRequest with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDBadRequest() *RescheduleAppointmentForServiceJobByServiceJobIDBadRequest {
	return &RescheduleAppointmentForServiceJobByServiceJobIDBadRequest{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDBadRequest struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdBadRequest  %+v", 400, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDBadRequest) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDForbidden creates a RescheduleAppointmentForServiceJobByServiceJobIDForbidden with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDForbidden() *RescheduleAppointmentForServiceJobByServiceJobIDForbidden {
	return &RescheduleAppointmentForServiceJobByServiceJobIDForbidden{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDForbidden) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdForbidden  %+v", 403, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDForbidden) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDNotFound creates a RescheduleAppointmentForServiceJobByServiceJobIDNotFound with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDNotFound() *RescheduleAppointmentForServiceJobByServiceJobIDNotFound {
	return &RescheduleAppointmentForServiceJobByServiceJobIDNotFound{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDNotFound) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdNotFound  %+v", 404, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDNotFound) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge creates a RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge() *RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge {
	return &RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType creates a RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType() *RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType {
	return &RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity creates a RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity() *RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity {
	return &RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity. Unable to process the contained instructions
*/
type RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests creates a RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests() *RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests {
	return &RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDInternalServerError creates a RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDInternalServerError() *RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError {
	return &RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdInternalServerError  %+v", 500, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable creates a RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable with default headers values
func NewRescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable() *RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable {
	return &RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable{}
}

/* RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.SetAppointmentResponse
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceJobs/{serviceJobId}/appointments/{appointmentId}][%d] rescheduleAppointmentForServiceJobByServiceJobIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable) GetPayload() *services_models.SetAppointmentResponse {
	return o.Payload
}

func (o *RescheduleAppointmentForServiceJobByServiceJobIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.SetAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}