// Code generated by go-swagger; DO NOT EDIT.

package shipment_invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/shipmentInvoicingV0/shipment_invoicing_v0_models"
)

// GetShipmentDetailsReader is a Reader for the GetShipmentDetails structure.
type GetShipmentDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetShipmentDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetShipmentDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetShipmentDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetShipmentDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetShipmentDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetShipmentDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetShipmentDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetShipmentDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetShipmentDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetShipmentDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetShipmentDetailsOK creates a GetShipmentDetailsOK with default headers values
func NewGetShipmentDetailsOK() *GetShipmentDetailsOK {
	return &GetShipmentDetailsOK{}
}

/* GetShipmentDetailsOK describes a response with status code 200, with default header values.

Success.
*/
type GetShipmentDetailsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsOK  %+v", 200, o.Payload)
}
func (o *GetShipmentDetailsOK) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsBadRequest creates a GetShipmentDetailsBadRequest with default headers values
func NewGetShipmentDetailsBadRequest() *GetShipmentDetailsBadRequest {
	return &GetShipmentDetailsBadRequest{}
}

/* GetShipmentDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetShipmentDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *GetShipmentDetailsBadRequest) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsUnauthorized creates a GetShipmentDetailsUnauthorized with default headers values
func NewGetShipmentDetailsUnauthorized() *GetShipmentDetailsUnauthorized {
	return &GetShipmentDetailsUnauthorized{}
}

/* GetShipmentDetailsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetShipmentDetailsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetShipmentDetailsUnauthorized) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsForbidden creates a GetShipmentDetailsForbidden with default headers values
func NewGetShipmentDetailsForbidden() *GetShipmentDetailsForbidden {
	return &GetShipmentDetailsForbidden{}
}

/* GetShipmentDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetShipmentDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsForbidden  %+v", 403, o.Payload)
}
func (o *GetShipmentDetailsForbidden) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsNotFound creates a GetShipmentDetailsNotFound with default headers values
func NewGetShipmentDetailsNotFound() *GetShipmentDetailsNotFound {
	return &GetShipmentDetailsNotFound{}
}

/* GetShipmentDetailsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetShipmentDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsNotFound  %+v", 404, o.Payload)
}
func (o *GetShipmentDetailsNotFound) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsUnsupportedMediaType creates a GetShipmentDetailsUnsupportedMediaType with default headers values
func NewGetShipmentDetailsUnsupportedMediaType() *GetShipmentDetailsUnsupportedMediaType {
	return &GetShipmentDetailsUnsupportedMediaType{}
}

/* GetShipmentDetailsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetShipmentDetailsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetShipmentDetailsUnsupportedMediaType) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsTooManyRequests creates a GetShipmentDetailsTooManyRequests with default headers values
func NewGetShipmentDetailsTooManyRequests() *GetShipmentDetailsTooManyRequests {
	return &GetShipmentDetailsTooManyRequests{}
}

/* GetShipmentDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetShipmentDetailsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetShipmentDetailsTooManyRequests) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsInternalServerError creates a GetShipmentDetailsInternalServerError with default headers values
func NewGetShipmentDetailsInternalServerError() *GetShipmentDetailsInternalServerError {
	return &GetShipmentDetailsInternalServerError{}
}

/* GetShipmentDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetShipmentDetailsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetShipmentDetailsInternalServerError) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentDetailsServiceUnavailable creates a GetShipmentDetailsServiceUnavailable with default headers values
func NewGetShipmentDetailsServiceUnavailable() *GetShipmentDetailsServiceUnavailable {
	return &GetShipmentDetailsServiceUnavailable{}
}

/* GetShipmentDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetShipmentDetailsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetShipmentDetailsResponse
}

func (o *GetShipmentDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}][%d] getShipmentDetailsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetShipmentDetailsServiceUnavailable) GetPayload() *shipment_invoicing_v0_models.GetShipmentDetailsResponse {
	return o.Payload
}

func (o *GetShipmentDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetShipmentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
