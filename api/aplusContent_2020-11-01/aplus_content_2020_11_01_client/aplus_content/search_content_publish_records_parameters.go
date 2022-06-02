// Code generated by go-swagger; DO NOT EDIT.

package aplus_content

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
)

// NewSearchContentPublishRecordsParams creates a new SearchContentPublishRecordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchContentPublishRecordsParams() *SearchContentPublishRecordsParams {
	return &SearchContentPublishRecordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchContentPublishRecordsParamsWithTimeout creates a new SearchContentPublishRecordsParams object
// with the ability to set a timeout on a request.
func NewSearchContentPublishRecordsParamsWithTimeout(timeout time.Duration) *SearchContentPublishRecordsParams {
	return &SearchContentPublishRecordsParams{
		timeout: timeout,
	}
}

// NewSearchContentPublishRecordsParamsWithContext creates a new SearchContentPublishRecordsParams object
// with the ability to set a context for a request.
func NewSearchContentPublishRecordsParamsWithContext(ctx context.Context) *SearchContentPublishRecordsParams {
	return &SearchContentPublishRecordsParams{
		Context: ctx,
	}
}

// NewSearchContentPublishRecordsParamsWithHTTPClient creates a new SearchContentPublishRecordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchContentPublishRecordsParamsWithHTTPClient(client *http.Client) *SearchContentPublishRecordsParams {
	return &SearchContentPublishRecordsParams{
		HTTPClient: client,
	}
}

/* SearchContentPublishRecordsParams contains all the parameters to send to the API endpoint
   for the search content publish records operation.

   Typically these are written to a http.Request.
*/
type SearchContentPublishRecordsParams struct {

	/* Asin.

	   The Amazon Standard Identification Number (ASIN).
	*/
	Asin string

	/* MarketplaceID.

	   The identifier for the marketplace where the A+ Content is published.
	*/
	MarketplaceID string

	/* PageToken.

	   A page token from the nextPageToken response element returned by your previous call to this operation. nextPageToken is returned when the results of a call exceed the page size. To get the next page of results, call the operation and include pageToken as the only parameter. Specifying pageToken with any other parameter will cause the request to fail. When no nextPageToken value is returned there are no more pages to return. A pageToken value is not usable across different operations.
	*/
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search content publish records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchContentPublishRecordsParams) WithDefaults() *SearchContentPublishRecordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search content publish records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchContentPublishRecordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search content publish records params
func (o *SearchContentPublishRecordsParams) WithTimeout(timeout time.Duration) *SearchContentPublishRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search content publish records params
func (o *SearchContentPublishRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search content publish records params
func (o *SearchContentPublishRecordsParams) WithContext(ctx context.Context) *SearchContentPublishRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search content publish records params
func (o *SearchContentPublishRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search content publish records params
func (o *SearchContentPublishRecordsParams) WithHTTPClient(client *http.Client) *SearchContentPublishRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search content publish records params
func (o *SearchContentPublishRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsin adds the asin to the search content publish records params
func (o *SearchContentPublishRecordsParams) WithAsin(asin string) *SearchContentPublishRecordsParams {
	o.SetAsin(asin)
	return o
}

// SetAsin adds the asin to the search content publish records params
func (o *SearchContentPublishRecordsParams) SetAsin(asin string) {
	o.Asin = asin
}

// WithMarketplaceID adds the marketplaceID to the search content publish records params
func (o *SearchContentPublishRecordsParams) WithMarketplaceID(marketplaceID string) *SearchContentPublishRecordsParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the search content publish records params
func (o *SearchContentPublishRecordsParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithPageToken adds the pageToken to the search content publish records params
func (o *SearchContentPublishRecordsParams) WithPageToken(pageToken *string) *SearchContentPublishRecordsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the search content publish records params
func (o *SearchContentPublishRecordsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *SearchContentPublishRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param asin
	qrAsin := o.Asin
	qAsin := qrAsin
	if qAsin != "" {

		if err := r.SetQueryParam("asin", qAsin); err != nil {
			return err
		}
	}

	// query param marketplaceId
	qrMarketplaceID := o.MarketplaceID
	qMarketplaceID := qrMarketplaceID
	if qMarketplaceID != "" {

		if err := r.SetQueryParam("marketplaceId", qMarketplaceID); err != nil {
			return err
		}
	}

	if o.PageToken != nil {

		// query param pageToken
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("pageToken", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}