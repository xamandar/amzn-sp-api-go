// Code generated by go-swagger; DO NOT EDIT.

package small_and_light

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
	"github.com/go-openapi/swag"
)

// NewGetSmallAndLightEnrollmentBySellerSKUParams creates a new GetSmallAndLightEnrollmentBySellerSKUParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSmallAndLightEnrollmentBySellerSKUParams() *GetSmallAndLightEnrollmentBySellerSKUParams {
	return &GetSmallAndLightEnrollmentBySellerSKUParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmallAndLightEnrollmentBySellerSKUParamsWithTimeout creates a new GetSmallAndLightEnrollmentBySellerSKUParams object
// with the ability to set a timeout on a request.
func NewGetSmallAndLightEnrollmentBySellerSKUParamsWithTimeout(timeout time.Duration) *GetSmallAndLightEnrollmentBySellerSKUParams {
	return &GetSmallAndLightEnrollmentBySellerSKUParams{
		timeout: timeout,
	}
}

// NewGetSmallAndLightEnrollmentBySellerSKUParamsWithContext creates a new GetSmallAndLightEnrollmentBySellerSKUParams object
// with the ability to set a context for a request.
func NewGetSmallAndLightEnrollmentBySellerSKUParamsWithContext(ctx context.Context) *GetSmallAndLightEnrollmentBySellerSKUParams {
	return &GetSmallAndLightEnrollmentBySellerSKUParams{
		Context: ctx,
	}
}

// NewGetSmallAndLightEnrollmentBySellerSKUParamsWithHTTPClient creates a new GetSmallAndLightEnrollmentBySellerSKUParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSmallAndLightEnrollmentBySellerSKUParamsWithHTTPClient(client *http.Client) *GetSmallAndLightEnrollmentBySellerSKUParams {
	return &GetSmallAndLightEnrollmentBySellerSKUParams{
		HTTPClient: client,
	}
}

/* GetSmallAndLightEnrollmentBySellerSKUParams contains all the parameters to send to the API endpoint
   for the get small and light enrollment by seller s k u operation.

   Typically these are written to a http.Request.
*/
type GetSmallAndLightEnrollmentBySellerSKUParams struct {

	/* MarketplaceIds.

	   The marketplace for which the enrollment status is retrieved. Note: Accepts a single marketplace only.
	*/
	MarketplaceIds []string

	/* SellerSKU.

	   The seller SKU that identifies the item.
	*/
	SellerSKU string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get small and light enrollment by seller s k u params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WithDefaults() *GetSmallAndLightEnrollmentBySellerSKUParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get small and light enrollment by seller s k u params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WithTimeout(timeout time.Duration) *GetSmallAndLightEnrollmentBySellerSKUParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WithContext(ctx context.Context) *GetSmallAndLightEnrollmentBySellerSKUParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WithHTTPClient(client *http.Client) *GetSmallAndLightEnrollmentBySellerSKUParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMarketplaceIds adds the marketplaceIds to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WithMarketplaceIds(marketplaceIds []string) *GetSmallAndLightEnrollmentBySellerSKUParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithSellerSKU adds the sellerSKU to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WithSellerSKU(sellerSKU string) *GetSmallAndLightEnrollmentBySellerSKUParams {
	o.SetSellerSKU(sellerSKU)
	return o
}

// SetSellerSKU adds the sellerSKU to the get small and light enrollment by seller s k u params
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) SetSellerSKU(sellerSKU string) {
	o.SellerSKU = sellerSKU
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MarketplaceIds != nil {

		// binding items for marketplaceIds
		joinedMarketplaceIds := o.bindParamMarketplaceIds(reg)

		// query array param marketplaceIds
		if err := r.SetQueryParam("marketplaceIds", joinedMarketplaceIds...); err != nil {
			return err
		}
	}

	// path param sellerSKU
	if err := r.SetPathParam("sellerSKU", o.SellerSKU); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetSmallAndLightEnrollmentBySellerSKU binds the parameter marketplaceIds
func (o *GetSmallAndLightEnrollmentBySellerSKUParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
	marketplaceIdsIR := o.MarketplaceIds

	var marketplaceIdsIC []string
	for _, marketplaceIdsIIR := range marketplaceIdsIR { // explode []string

		marketplaceIdsIIV := marketplaceIdsIIR // string as string
		marketplaceIdsIC = append(marketplaceIdsIC, marketplaceIdsIIV)
	}

	// items.CollectionFormat: ""
	marketplaceIdsIS := swag.JoinByFormat(marketplaceIdsIC, "")

	return marketplaceIdsIS
}