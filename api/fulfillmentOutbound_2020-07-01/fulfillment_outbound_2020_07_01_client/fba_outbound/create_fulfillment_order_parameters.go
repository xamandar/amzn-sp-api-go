// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// NewCreateFulfillmentOrderParams creates a new CreateFulfillmentOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFulfillmentOrderParams() *CreateFulfillmentOrderParams {
	return &CreateFulfillmentOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFulfillmentOrderParamsWithTimeout creates a new CreateFulfillmentOrderParams object
// with the ability to set a timeout on a request.
func NewCreateFulfillmentOrderParamsWithTimeout(timeout time.Duration) *CreateFulfillmentOrderParams {
	return &CreateFulfillmentOrderParams{
		timeout: timeout,
	}
}

// NewCreateFulfillmentOrderParamsWithContext creates a new CreateFulfillmentOrderParams object
// with the ability to set a context for a request.
func NewCreateFulfillmentOrderParamsWithContext(ctx context.Context) *CreateFulfillmentOrderParams {
	return &CreateFulfillmentOrderParams{
		Context: ctx,
	}
}

// NewCreateFulfillmentOrderParamsWithHTTPClient creates a new CreateFulfillmentOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFulfillmentOrderParamsWithHTTPClient(client *http.Client) *CreateFulfillmentOrderParams {
	return &CreateFulfillmentOrderParams{
		HTTPClient: client,
	}
}

/* CreateFulfillmentOrderParams contains all the parameters to send to the API endpoint
   for the create fulfillment order operation.

   Typically these are written to a http.Request.
*/
type CreateFulfillmentOrderParams struct {

	// Body.
	Body *fulfillment_outbound_2020_07_01_models.CreateFulfillmentOrderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create fulfillment order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFulfillmentOrderParams) WithDefaults() *CreateFulfillmentOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create fulfillment order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFulfillmentOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) WithTimeout(timeout time.Duration) *CreateFulfillmentOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) WithContext(ctx context.Context) *CreateFulfillmentOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) WithHTTPClient(client *http.Client) *CreateFulfillmentOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) WithBody(body *fulfillment_outbound_2020_07_01_models.CreateFulfillmentOrderRequest) *CreateFulfillmentOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create fulfillment order params
func (o *CreateFulfillmentOrderParams) SetBody(body *fulfillment_outbound_2020_07_01_models.CreateFulfillmentOrderRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFulfillmentOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}