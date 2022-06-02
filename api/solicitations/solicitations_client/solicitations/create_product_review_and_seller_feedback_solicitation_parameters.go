// Code generated by go-swagger; DO NOT EDIT.

package solicitations

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

// NewCreateProductReviewAndSellerFeedbackSolicitationParams creates a new CreateProductReviewAndSellerFeedbackSolicitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProductReviewAndSellerFeedbackSolicitationParams() *CreateProductReviewAndSellerFeedbackSolicitationParams {
	return &CreateProductReviewAndSellerFeedbackSolicitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProductReviewAndSellerFeedbackSolicitationParamsWithTimeout creates a new CreateProductReviewAndSellerFeedbackSolicitationParams object
// with the ability to set a timeout on a request.
func NewCreateProductReviewAndSellerFeedbackSolicitationParamsWithTimeout(timeout time.Duration) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	return &CreateProductReviewAndSellerFeedbackSolicitationParams{
		timeout: timeout,
	}
}

// NewCreateProductReviewAndSellerFeedbackSolicitationParamsWithContext creates a new CreateProductReviewAndSellerFeedbackSolicitationParams object
// with the ability to set a context for a request.
func NewCreateProductReviewAndSellerFeedbackSolicitationParamsWithContext(ctx context.Context) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	return &CreateProductReviewAndSellerFeedbackSolicitationParams{
		Context: ctx,
	}
}

// NewCreateProductReviewAndSellerFeedbackSolicitationParamsWithHTTPClient creates a new CreateProductReviewAndSellerFeedbackSolicitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProductReviewAndSellerFeedbackSolicitationParamsWithHTTPClient(client *http.Client) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	return &CreateProductReviewAndSellerFeedbackSolicitationParams{
		HTTPClient: client,
	}
}

/* CreateProductReviewAndSellerFeedbackSolicitationParams contains all the parameters to send to the API endpoint
   for the create product review and seller feedback solicitation operation.

   Typically these are written to a http.Request.
*/
type CreateProductReviewAndSellerFeedbackSolicitationParams struct {

	/* AmazonOrderID.

	   An Amazon order identifier. This specifies the order for which a solicitation is sent.
	*/
	AmazonOrderID string

	/* MarketplaceIds.

	   A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	*/
	MarketplaceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create product review and seller feedback solicitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WithDefaults() *CreateProductReviewAndSellerFeedbackSolicitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create product review and seller feedback solicitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WithTimeout(timeout time.Duration) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WithContext(ctx context.Context) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WithHTTPClient(client *http.Client) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmazonOrderID adds the amazonOrderID to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WithAmazonOrderID(amazonOrderID string) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	o.SetAmazonOrderID(amazonOrderID)
	return o
}

// SetAmazonOrderID adds the amazonOrderId to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) SetAmazonOrderID(amazonOrderID string) {
	o.AmazonOrderID = amazonOrderID
}

// WithMarketplaceIds adds the marketplaceIds to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WithMarketplaceIds(marketplaceIds []string) *CreateProductReviewAndSellerFeedbackSolicitationParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the create product review and seller feedback solicitation params
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param amazonOrderId
	if err := r.SetPathParam("amazonOrderId", o.AmazonOrderID); err != nil {
		return err
	}

	if o.MarketplaceIds != nil {

		// binding items for marketplaceIds
		joinedMarketplaceIds := o.bindParamMarketplaceIds(reg)

		// query array param marketplaceIds
		if err := r.SetQueryParam("marketplaceIds", joinedMarketplaceIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCreateProductReviewAndSellerFeedbackSolicitation binds the parameter marketplaceIds
func (o *CreateProductReviewAndSellerFeedbackSolicitationParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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