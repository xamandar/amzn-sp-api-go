// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShippingV1/vendor_direct_fulfillment_shipping_v1_models"
)

// GetShippingLabelsReader is a Reader for the GetShippingLabels structure.
type GetShippingLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetShippingLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetShippingLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetShippingLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetShippingLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetShippingLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetShippingLabelsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetShippingLabelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetShippingLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetShippingLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetShippingLabelsOK creates a GetShippingLabelsOK with default headers values
func NewGetShippingLabelsOK() *GetShippingLabelsOK {
	return &GetShippingLabelsOK{}
}

/* GetShippingLabelsOK describes a response with status code 200, with default header values.

Success.
*/
type GetShippingLabelsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsOK  %+v", 200, o.Payload)
}
func (o *GetShippingLabelsOK) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsBadRequest creates a GetShippingLabelsBadRequest with default headers values
func NewGetShippingLabelsBadRequest() *GetShippingLabelsBadRequest {
	return &GetShippingLabelsBadRequest{}
}

/* GetShippingLabelsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetShippingLabelsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsBadRequest  %+v", 400, o.Payload)
}
func (o *GetShippingLabelsBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsForbidden creates a GetShippingLabelsForbidden with default headers values
func NewGetShippingLabelsForbidden() *GetShippingLabelsForbidden {
	return &GetShippingLabelsForbidden{}
}

/* GetShippingLabelsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetShippingLabelsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsForbidden  %+v", 403, o.Payload)
}
func (o *GetShippingLabelsForbidden) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsNotFound creates a GetShippingLabelsNotFound with default headers values
func NewGetShippingLabelsNotFound() *GetShippingLabelsNotFound {
	return &GetShippingLabelsNotFound{}
}

/* GetShippingLabelsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetShippingLabelsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsNotFound  %+v", 404, o.Payload)
}
func (o *GetShippingLabelsNotFound) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsUnsupportedMediaType creates a GetShippingLabelsUnsupportedMediaType with default headers values
func NewGetShippingLabelsUnsupportedMediaType() *GetShippingLabelsUnsupportedMediaType {
	return &GetShippingLabelsUnsupportedMediaType{}
}

/* GetShippingLabelsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetShippingLabelsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetShippingLabelsUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsTooManyRequests creates a GetShippingLabelsTooManyRequests with default headers values
func NewGetShippingLabelsTooManyRequests() *GetShippingLabelsTooManyRequests {
	return &GetShippingLabelsTooManyRequests{}
}

/* GetShippingLabelsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetShippingLabelsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetShippingLabelsTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsInternalServerError creates a GetShippingLabelsInternalServerError with default headers values
func NewGetShippingLabelsInternalServerError() *GetShippingLabelsInternalServerError {
	return &GetShippingLabelsInternalServerError{}
}

/* GetShippingLabelsInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type GetShippingLabelsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetShippingLabelsInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelsServiceUnavailable creates a GetShippingLabelsServiceUnavailable with default headers values
func NewGetShippingLabelsServiceUnavailable() *GetShippingLabelsServiceUnavailable {
	return &GetShippingLabelsServiceUnavailable{}
}

/* GetShippingLabelsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetShippingLabelsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse
}

func (o *GetShippingLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/shippingLabels][%d] getShippingLabelsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetShippingLabelsServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse {
	return o.Payload
}

func (o *GetShippingLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetShippingLabelListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
