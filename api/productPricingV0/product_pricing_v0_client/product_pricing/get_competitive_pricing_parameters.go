// Code generated by go-swagger; DO NOT EDIT.

package product_pricing

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

// NewGetCompetitivePricingParams creates a new GetCompetitivePricingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompetitivePricingParams() *GetCompetitivePricingParams {
	return &GetCompetitivePricingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompetitivePricingParamsWithTimeout creates a new GetCompetitivePricingParams object
// with the ability to set a timeout on a request.
func NewGetCompetitivePricingParamsWithTimeout(timeout time.Duration) *GetCompetitivePricingParams {
	return &GetCompetitivePricingParams{
		timeout: timeout,
	}
}

// NewGetCompetitivePricingParamsWithContext creates a new GetCompetitivePricingParams object
// with the ability to set a context for a request.
func NewGetCompetitivePricingParamsWithContext(ctx context.Context) *GetCompetitivePricingParams {
	return &GetCompetitivePricingParams{
		Context: ctx,
	}
}

// NewGetCompetitivePricingParamsWithHTTPClient creates a new GetCompetitivePricingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompetitivePricingParamsWithHTTPClient(client *http.Client) *GetCompetitivePricingParams {
	return &GetCompetitivePricingParams{
		HTTPClient: client,
	}
}

/* GetCompetitivePricingParams contains all the parameters to send to the API endpoint
   for the get competitive pricing operation.

   Typically these are written to a http.Request.
*/
type GetCompetitivePricingParams struct {

	/* Asins.

	   A list of up to twenty Amazon Standard Identification Number (ASIN) values used to identify items in the given marketplace.
	*/
	Asins []string

	/* CustomerType.

	   Indicates whether to request pricing information from the point of view of Consumer or Business buyers. Default is Consumer.
	*/
	CustomerType *string

	/* ItemType.

	   Indicates whether ASIN values or seller SKU values are used to identify items. If you specify Asin, the information in the response will be dependent on the list of Asins you provide in the Asins parameter. If you specify Sku, the information in the response will be dependent on the list of Skus you provide in the Skus parameter. Possible values: Asin, Sku.
	*/
	ItemType string

	/* MarketplaceID.

	   A marketplace identifier. Specifies the marketplace for which prices are returned.
	*/
	MarketplaceID string

	/* Skus.

	   A list of up to twenty seller SKU values used to identify items in the given marketplace.
	*/
	Skus []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get competitive pricing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompetitivePricingParams) WithDefaults() *GetCompetitivePricingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get competitive pricing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompetitivePricingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithTimeout(timeout time.Duration) *GetCompetitivePricingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithContext(ctx context.Context) *GetCompetitivePricingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithHTTPClient(client *http.Client) *GetCompetitivePricingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsins adds the asins to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithAsins(asins []string) *GetCompetitivePricingParams {
	o.SetAsins(asins)
	return o
}

// SetAsins adds the asins to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetAsins(asins []string) {
	o.Asins = asins
}

// WithCustomerType adds the customerType to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithCustomerType(customerType *string) *GetCompetitivePricingParams {
	o.SetCustomerType(customerType)
	return o
}

// SetCustomerType adds the customerType to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetCustomerType(customerType *string) {
	o.CustomerType = customerType
}

// WithItemType adds the itemType to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithItemType(itemType string) *GetCompetitivePricingParams {
	o.SetItemType(itemType)
	return o
}

// SetItemType adds the itemType to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetItemType(itemType string) {
	o.ItemType = itemType
}

// WithMarketplaceID adds the marketplaceID to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithMarketplaceID(marketplaceID string) *GetCompetitivePricingParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithSkus adds the skus to the get competitive pricing params
func (o *GetCompetitivePricingParams) WithSkus(skus []string) *GetCompetitivePricingParams {
	o.SetSkus(skus)
	return o
}

// SetSkus adds the skus to the get competitive pricing params
func (o *GetCompetitivePricingParams) SetSkus(skus []string) {
	o.Skus = skus
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompetitivePricingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Asins != nil {

		// binding items for Asins
		joinedAsins := o.bindParamAsins(reg)

		// query array param Asins
		if err := r.SetQueryParam("Asins", joinedAsins...); err != nil {
			return err
		}
	}

	if o.CustomerType != nil {

		// query param CustomerType
		var qrCustomerType string

		if o.CustomerType != nil {
			qrCustomerType = *o.CustomerType
		}
		qCustomerType := qrCustomerType
		if qCustomerType != "" {

			if err := r.SetQueryParam("CustomerType", qCustomerType); err != nil {
				return err
			}
		}
	}

	// query param ItemType
	qrItemType := o.ItemType
	qItemType := qrItemType
	if qItemType != "" {

		if err := r.SetQueryParam("ItemType", qItemType); err != nil {
			return err
		}
	}

	// query param MarketplaceId
	qrMarketplaceID := o.MarketplaceID
	qMarketplaceID := qrMarketplaceID
	if qMarketplaceID != "" {

		if err := r.SetQueryParam("MarketplaceId", qMarketplaceID); err != nil {
			return err
		}
	}

	if o.Skus != nil {

		// binding items for Skus
		joinedSkus := o.bindParamSkus(reg)

		// query array param Skus
		if err := r.SetQueryParam("Skus", joinedSkus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCompetitivePricing binds the parameter Asins
func (o *GetCompetitivePricingParams) bindParamAsins(formats strfmt.Registry) []string {
	asinsIR := o.Asins

	var asinsIC []string
	for _, asinsIIR := range asinsIR { // explode []string

		asinsIIV := asinsIIR // string as string
		asinsIC = append(asinsIC, asinsIIV)
	}

	// items.CollectionFormat: ""
	asinsIS := swag.JoinByFormat(asinsIC, "")

	return asinsIS
}

// bindParamGetCompetitivePricing binds the parameter Skus
func (o *GetCompetitivePricingParams) bindParamSkus(formats strfmt.Registry) []string {
	skusIR := o.Skus

	var skusIC []string
	for _, skusIIR := range skusIR { // explode []string

		skusIIV := skusIIR // string as string
		skusIC = append(skusIC, skusIIV)
	}

	// items.CollectionFormat: ""
	skusIS := swag.JoinByFormat(skusIC, "")

	return skusIS
}