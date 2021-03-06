// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorOrders/vendor_orders_models"
)

// GetPurchaseOrderReader is a Reader for the GetPurchaseOrder structure.
type GetPurchaseOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurchaseOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurchaseOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPurchaseOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPurchaseOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPurchaseOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPurchaseOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetPurchaseOrderUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPurchaseOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurchaseOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPurchaseOrderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPurchaseOrderOK creates a GetPurchaseOrderOK with default headers values
func NewGetPurchaseOrderOK() *GetPurchaseOrderOK {
	return &GetPurchaseOrderOK{}
}

/* GetPurchaseOrderOK describes a response with status code 200, with default header values.

Success.
*/
type GetPurchaseOrderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderOK) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderOK  %+v", 200, o.Payload)
}
func (o *GetPurchaseOrderOK) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderBadRequest creates a GetPurchaseOrderBadRequest with default headers values
func NewGetPurchaseOrderBadRequest() *GetPurchaseOrderBadRequest {
	return &GetPurchaseOrderBadRequest{}
}

/* GetPurchaseOrderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPurchaseOrderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderBadRequest  %+v", 400, o.Payload)
}
func (o *GetPurchaseOrderBadRequest) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderUnauthorized creates a GetPurchaseOrderUnauthorized with default headers values
func NewGetPurchaseOrderUnauthorized() *GetPurchaseOrderUnauthorized {
	return &GetPurchaseOrderUnauthorized{}
}

/* GetPurchaseOrderUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetPurchaseOrderUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPurchaseOrderUnauthorized) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderForbidden creates a GetPurchaseOrderForbidden with default headers values
func NewGetPurchaseOrderForbidden() *GetPurchaseOrderForbidden {
	return &GetPurchaseOrderForbidden{}
}

/* GetPurchaseOrderForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetPurchaseOrderForbidden struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderForbidden  %+v", 403, o.Payload)
}
func (o *GetPurchaseOrderForbidden) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderNotFound creates a GetPurchaseOrderNotFound with default headers values
func NewGetPurchaseOrderNotFound() *GetPurchaseOrderNotFound {
	return &GetPurchaseOrderNotFound{}
}

/* GetPurchaseOrderNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetPurchaseOrderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderNotFound  %+v", 404, o.Payload)
}
func (o *GetPurchaseOrderNotFound) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderUnsupportedMediaType creates a GetPurchaseOrderUnsupportedMediaType with default headers values
func NewGetPurchaseOrderUnsupportedMediaType() *GetPurchaseOrderUnsupportedMediaType {
	return &GetPurchaseOrderUnsupportedMediaType{}
}

/* GetPurchaseOrderUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetPurchaseOrderUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetPurchaseOrderUnsupportedMediaType) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderTooManyRequests creates a GetPurchaseOrderTooManyRequests with default headers values
func NewGetPurchaseOrderTooManyRequests() *GetPurchaseOrderTooManyRequests {
	return &GetPurchaseOrderTooManyRequests{}
}

/* GetPurchaseOrderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPurchaseOrderTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetPurchaseOrderTooManyRequests) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderInternalServerError creates a GetPurchaseOrderInternalServerError with default headers values
func NewGetPurchaseOrderInternalServerError() *GetPurchaseOrderInternalServerError {
	return &GetPurchaseOrderInternalServerError{}
}

/* GetPurchaseOrderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetPurchaseOrderInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPurchaseOrderInternalServerError) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrderServiceUnavailable creates a GetPurchaseOrderServiceUnavailable with default headers values
func NewGetPurchaseOrderServiceUnavailable() *GetPurchaseOrderServiceUnavailable {
	return &GetPurchaseOrderServiceUnavailable{}
}

/* GetPurchaseOrderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPurchaseOrderServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrderResponse
}

func (o *GetPurchaseOrderServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}][%d] getPurchaseOrderServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetPurchaseOrderServiceUnavailable) GetPayload() *vendor_orders_models.GetPurchaseOrderResponse {
	return o.Payload
}

func (o *GetPurchaseOrderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
