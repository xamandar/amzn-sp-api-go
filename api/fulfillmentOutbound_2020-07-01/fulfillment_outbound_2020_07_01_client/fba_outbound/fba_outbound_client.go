// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fba outbound API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fba outbound API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelFulfillmentOrder(params *CancelFulfillmentOrderParams, opts ...ClientOption) (*CancelFulfillmentOrderOK, error)

	CreateFulfillmentOrder(params *CreateFulfillmentOrderParams, opts ...ClientOption) (*CreateFulfillmentOrderOK, error)

	CreateFulfillmentReturn(params *CreateFulfillmentReturnParams, opts ...ClientOption) (*CreateFulfillmentReturnOK, error)

	GetFeatureInventory(params *GetFeatureInventoryParams, opts ...ClientOption) (*GetFeatureInventoryOK, error)

	GetFeatureSKU(params *GetFeatureSKUParams, opts ...ClientOption) (*GetFeatureSKUOK, error)

	GetFeatures(params *GetFeaturesParams, opts ...ClientOption) (*GetFeaturesOK, error)

	GetFulfillmentOrder(params *GetFulfillmentOrderParams, opts ...ClientOption) (*GetFulfillmentOrderOK, error)

	GetFulfillmentPreview(params *GetFulfillmentPreviewParams, opts ...ClientOption) (*GetFulfillmentPreviewOK, error)

	GetPackageTrackingDetails(params *GetPackageTrackingDetailsParams, opts ...ClientOption) (*GetPackageTrackingDetailsOK, error)

	ListAllFulfillmentOrders(params *ListAllFulfillmentOrdersParams, opts ...ClientOption) (*ListAllFulfillmentOrdersOK, error)

	ListReturnReasonCodes(params *ListReturnReasonCodesParams, opts ...ClientOption) (*ListReturnReasonCodesOK, error)

	UpdateFulfillmentOrder(params *UpdateFulfillmentOrderParams, opts ...ClientOption) (*UpdateFulfillmentOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelFulfillmentOrder Requests that Amazon stop attempting to fulfill the fulfillment order indicated by the specified order identifier.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CancelFulfillmentOrder(params *CancelFulfillmentOrderParams, opts ...ClientOption) (*CancelFulfillmentOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelFulfillmentOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelFulfillmentOrder",
		Method:             "PUT",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelFulfillmentOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CancelFulfillmentOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancelFulfillmentOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateFulfillmentOrder Requests that Amazon ship items from the seller's inventory in Amazon's fulfillment network to a destination address.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CreateFulfillmentOrder(params *CreateFulfillmentOrderParams, opts ...ClientOption) (*CreateFulfillmentOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFulfillmentOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFulfillmentOrder",
		Method:             "POST",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFulfillmentOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateFulfillmentOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFulfillmentOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateFulfillmentReturn Creates a fulfillment return.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CreateFulfillmentReturn(params *CreateFulfillmentReturnParams, opts ...ClientOption) (*CreateFulfillmentReturnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFulfillmentReturnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFulfillmentReturn",
		Method:             "PUT",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFulfillmentReturnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateFulfillmentReturnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFulfillmentReturn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeatureInventory Returns a list of inventory items that are eligible for the fulfillment feature you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFeatureInventory(params *GetFeatureInventoryParams, opts ...ClientOption) (*GetFeatureInventoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeatureInventoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFeatureInventory",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/features/inventory/{featureName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeatureInventoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFeatureInventoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFeatureInventory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeatureSKU Returns the number of items with the sellerSKU you specify that can have orders fulfilled using the specified feature. Note that if the sellerSKU isn't eligible, the response will contain an empty skuInfo object.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFeatureSKU(params *GetFeatureSKUParams, opts ...ClientOption) (*GetFeatureSKUOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeatureSKUParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFeatureSKU",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/features/inventory/{featureName}/{sellerSku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeatureSKUReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFeatureSKUOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFeatureSKU: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeatures Returns a list of features available for Multi-Channel Fulfillment orders in the marketplace you specify, and whether the seller for which you made the call is enrolled for each feature.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFeatures(params *GetFeaturesParams, opts ...ClientOption) (*GetFeaturesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeaturesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFeatures",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeaturesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFeaturesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFeatures: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFulfillmentOrder Returns the fulfillment order indicated by the specified order identifier.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFulfillmentOrder(params *GetFulfillmentOrderParams, opts ...ClientOption) (*GetFulfillmentOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFulfillmentOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFulfillmentOrder",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFulfillmentOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFulfillmentOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFulfillmentOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFulfillmentPreview Returns a list of fulfillment order previews based on shipping criteria that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFulfillmentPreview(params *GetFulfillmentPreviewParams, opts ...ClientOption) (*GetFulfillmentPreviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFulfillmentPreviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFulfillmentPreview",
		Method:             "POST",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders/preview",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFulfillmentPreviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFulfillmentPreviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFulfillmentPreview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPackageTrackingDetails Returns delivery tracking information for a package in an outbound shipment for a Multi-Channel Fulfillment order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetPackageTrackingDetails(params *GetPackageTrackingDetailsParams, opts ...ClientOption) (*GetPackageTrackingDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPackageTrackingDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPackageTrackingDetails",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/tracking",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPackageTrackingDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPackageTrackingDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPackageTrackingDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllFulfillmentOrders Returns a list of fulfillment orders fulfilled after (or at) a specified date-time, or indicated by the next token parameter.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) ListAllFulfillmentOrders(params *ListAllFulfillmentOrdersParams, opts ...ClientOption) (*ListAllFulfillmentOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllFulfillmentOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAllFulfillmentOrders",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAllFulfillmentOrdersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllFulfillmentOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllFulfillmentOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListReturnReasonCodes Returns a list of return reason codes for a seller SKU in a given marketplace.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) ListReturnReasonCodes(params *ListReturnReasonCodesParams, opts ...ClientOption) (*ListReturnReasonCodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReturnReasonCodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listReturnReasonCodes",
		Method:             "GET",
		PathPattern:        "/fba/outbound/2020-07-01/returnReasonCodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReturnReasonCodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListReturnReasonCodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listReturnReasonCodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateFulfillmentOrder Updates and/or requests shipment for a fulfillment order with an order hold on it.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 30 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) UpdateFulfillmentOrder(params *UpdateFulfillmentOrderParams, opts ...ClientOption) (*UpdateFulfillmentOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFulfillmentOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFulfillmentOrder",
		Method:             "PUT",
		PathPattern:        "/fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFulfillmentOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateFulfillmentOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFulfillmentOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}