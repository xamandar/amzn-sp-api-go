// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentOrders_2021-12-28/vendor_direct_fulfillment_orders_2021_12_28_models"
)

// GetOrdersReader is a Reader for the GetOrders structure.
type GetOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrdersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrdersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrdersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrdersOK creates a GetOrdersOK with default headers values
func NewGetOrdersOK() *GetOrdersOK {
	return &GetOrdersOK{}
}

/* GetOrdersOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrdersOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.OrderList
}

func (o *GetOrdersOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersOK  %+v", 200, o.Payload)
}
func (o *GetOrdersOK) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.OrderList {
	return o.Payload
}

func (o *GetOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.OrderList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersBadRequest creates a GetOrdersBadRequest with default headers values
func NewGetOrdersBadRequest() *GetOrdersBadRequest {
	return &GetOrdersBadRequest{}
}

/* GetOrdersBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrdersBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersBadRequest  %+v", 400, o.Payload)
}
func (o *GetOrdersBadRequest) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersForbidden creates a GetOrdersForbidden with default headers values
func NewGetOrdersForbidden() *GetOrdersForbidden {
	return &GetOrdersForbidden{}
}

/* GetOrdersForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrdersForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersForbidden  %+v", 403, o.Payload)
}
func (o *GetOrdersForbidden) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersNotFound creates a GetOrdersNotFound with default headers values
func NewGetOrdersNotFound() *GetOrdersNotFound {
	return &GetOrdersNotFound{}
}

/* GetOrdersNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrdersNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersNotFound  %+v", 404, o.Payload)
}
func (o *GetOrdersNotFound) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersUnsupportedMediaType creates a GetOrdersUnsupportedMediaType with default headers values
func NewGetOrdersUnsupportedMediaType() *GetOrdersUnsupportedMediaType {
	return &GetOrdersUnsupportedMediaType{}
}

/* GetOrdersUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetOrdersUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetOrdersUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersTooManyRequests creates a GetOrdersTooManyRequests with default headers values
func NewGetOrdersTooManyRequests() *GetOrdersTooManyRequests {
	return &GetOrdersTooManyRequests{}
}

/* GetOrdersTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrdersTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetOrdersTooManyRequests) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersInternalServerError creates a GetOrdersInternalServerError with default headers values
func NewGetOrdersInternalServerError() *GetOrdersInternalServerError {
	return &GetOrdersInternalServerError{}
}

/* GetOrdersInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrdersInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOrdersInternalServerError) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersServiceUnavailable creates a GetOrdersServiceUnavailable with default headers values
func NewGetOrdersServiceUnavailable() *GetOrdersServiceUnavailable {
	return &GetOrdersServiceUnavailable{}
}

/* GetOrdersServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrdersServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList
}

func (o *GetOrdersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/orders/2021-12-28/purchaseOrders][%d] getOrdersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetOrdersServiceUnavailable) GetPayload() *vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetOrdersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_orders_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
