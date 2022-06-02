// Code generated by go-swagger; DO NOT EDIT.

package listings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/listingsItems_2021-08-01/listings_items_2021_08_01_models"
)

// GetListingsItemReader is a Reader for the GetListingsItem structure.
type GetListingsItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListingsItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListingsItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListingsItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListingsItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetListingsItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetListingsItemRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetListingsItemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetListingsItemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetListingsItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetListingsItemServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListingsItemOK creates a GetListingsItemOK with default headers values
func NewGetListingsItemOK() *GetListingsItemOK {
	return &GetListingsItemOK{}
}

/* GetListingsItemOK describes a response with status code 200, with default header values.

Success.
*/
type GetListingsItemOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.Item
}

func (o *GetListingsItemOK) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemOK  %+v", 200, o.Payload)
}
func (o *GetListingsItemOK) GetPayload() *listings_items_2021_08_01_models.Item {
	return o.Payload
}

func (o *GetListingsItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(listings_items_2021_08_01_models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemBadRequest creates a GetListingsItemBadRequest with default headers values
func NewGetListingsItemBadRequest() *GetListingsItemBadRequest {
	return &GetListingsItemBadRequest{}
}

/* GetListingsItemBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetListingsItemBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemBadRequest) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemBadRequest  %+v", 400, o.Payload)
}
func (o *GetListingsItemBadRequest) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemForbidden creates a GetListingsItemForbidden with default headers values
func NewGetListingsItemForbidden() *GetListingsItemForbidden {
	return &GetListingsItemForbidden{}
}

/* GetListingsItemForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetListingsItemForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemForbidden) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemForbidden  %+v", 403, o.Payload)
}
func (o *GetListingsItemForbidden) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemNotFound creates a GetListingsItemNotFound with default headers values
func NewGetListingsItemNotFound() *GetListingsItemNotFound {
	return &GetListingsItemNotFound{}
}

/* GetListingsItemNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetListingsItemNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemNotFound) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemNotFound  %+v", 404, o.Payload)
}
func (o *GetListingsItemNotFound) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemRequestEntityTooLarge creates a GetListingsItemRequestEntityTooLarge with default headers values
func NewGetListingsItemRequestEntityTooLarge() *GetListingsItemRequestEntityTooLarge {
	return &GetListingsItemRequestEntityTooLarge{}
}

/* GetListingsItemRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetListingsItemRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *GetListingsItemRequestEntityTooLarge) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemUnsupportedMediaType creates a GetListingsItemUnsupportedMediaType with default headers values
func NewGetListingsItemUnsupportedMediaType() *GetListingsItemUnsupportedMediaType {
	return &GetListingsItemUnsupportedMediaType{}
}

/* GetListingsItemUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetListingsItemUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetListingsItemUnsupportedMediaType) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemTooManyRequests creates a GetListingsItemTooManyRequests with default headers values
func NewGetListingsItemTooManyRequests() *GetListingsItemTooManyRequests {
	return &GetListingsItemTooManyRequests{}
}

/* GetListingsItemTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetListingsItemTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetListingsItemTooManyRequests) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemInternalServerError creates a GetListingsItemInternalServerError with default headers values
func NewGetListingsItemInternalServerError() *GetListingsItemInternalServerError {
	return &GetListingsItemInternalServerError{}
}

/* GetListingsItemInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetListingsItemInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemInternalServerError) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemInternalServerError  %+v", 500, o.Payload)
}
func (o *GetListingsItemInternalServerError) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsItemServiceUnavailable creates a GetListingsItemServiceUnavailable with default headers values
func NewGetListingsItemServiceUnavailable() *GetListingsItemServiceUnavailable {
	return &GetListingsItemServiceUnavailable{}
}

/* GetListingsItemServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetListingsItemServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

func (o *GetListingsItemServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/items/{sellerId}/{sku}][%d] getListingsItemServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetListingsItemServiceUnavailable) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsItemServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}