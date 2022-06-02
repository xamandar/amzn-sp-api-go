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
	"github.com/go-openapi/swag"
)

// NewListContentDocumentAsinRelationsParams creates a new ListContentDocumentAsinRelationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListContentDocumentAsinRelationsParams() *ListContentDocumentAsinRelationsParams {
	return &ListContentDocumentAsinRelationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListContentDocumentAsinRelationsParamsWithTimeout creates a new ListContentDocumentAsinRelationsParams object
// with the ability to set a timeout on a request.
func NewListContentDocumentAsinRelationsParamsWithTimeout(timeout time.Duration) *ListContentDocumentAsinRelationsParams {
	return &ListContentDocumentAsinRelationsParams{
		timeout: timeout,
	}
}

// NewListContentDocumentAsinRelationsParamsWithContext creates a new ListContentDocumentAsinRelationsParams object
// with the ability to set a context for a request.
func NewListContentDocumentAsinRelationsParamsWithContext(ctx context.Context) *ListContentDocumentAsinRelationsParams {
	return &ListContentDocumentAsinRelationsParams{
		Context: ctx,
	}
}

// NewListContentDocumentAsinRelationsParamsWithHTTPClient creates a new ListContentDocumentAsinRelationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListContentDocumentAsinRelationsParamsWithHTTPClient(client *http.Client) *ListContentDocumentAsinRelationsParams {
	return &ListContentDocumentAsinRelationsParams{
		HTTPClient: client,
	}
}

/* ListContentDocumentAsinRelationsParams contains all the parameters to send to the API endpoint
   for the list content document asin relations operation.

   Typically these are written to a http.Request.
*/
type ListContentDocumentAsinRelationsParams struct {

	/* AsinSet.

	   The set of ASINs.
	*/
	AsinSet []string

	/* ContentReferenceKey.

	   The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ Content identifier.
	*/
	ContentReferenceKey string

	/* IncludedDataSet.

	   The set of A+ Content data types to include in the response. If you do not include this parameter, the operation returns the related ASINs without metadata.
	*/
	IncludedDataSet []string

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

// WithDefaults hydrates default values in the list content document asin relations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListContentDocumentAsinRelationsParams) WithDefaults() *ListContentDocumentAsinRelationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list content document asin relations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListContentDocumentAsinRelationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithTimeout(timeout time.Duration) *ListContentDocumentAsinRelationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithContext(ctx context.Context) *ListContentDocumentAsinRelationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithHTTPClient(client *http.Client) *ListContentDocumentAsinRelationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsinSet adds the asinSet to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithAsinSet(asinSet []string) *ListContentDocumentAsinRelationsParams {
	o.SetAsinSet(asinSet)
	return o
}

// SetAsinSet adds the asinSet to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetAsinSet(asinSet []string) {
	o.AsinSet = asinSet
}

// WithContentReferenceKey adds the contentReferenceKey to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithContentReferenceKey(contentReferenceKey string) *ListContentDocumentAsinRelationsParams {
	o.SetContentReferenceKey(contentReferenceKey)
	return o
}

// SetContentReferenceKey adds the contentReferenceKey to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetContentReferenceKey(contentReferenceKey string) {
	o.ContentReferenceKey = contentReferenceKey
}

// WithIncludedDataSet adds the includedDataSet to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithIncludedDataSet(includedDataSet []string) *ListContentDocumentAsinRelationsParams {
	o.SetIncludedDataSet(includedDataSet)
	return o
}

// SetIncludedDataSet adds the includedDataSet to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetIncludedDataSet(includedDataSet []string) {
	o.IncludedDataSet = includedDataSet
}

// WithMarketplaceID adds the marketplaceID to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithMarketplaceID(marketplaceID string) *ListContentDocumentAsinRelationsParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithPageToken adds the pageToken to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) WithPageToken(pageToken *string) *ListContentDocumentAsinRelationsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list content document asin relations params
func (o *ListContentDocumentAsinRelationsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListContentDocumentAsinRelationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AsinSet != nil {

		// binding items for asinSet
		joinedAsinSet := o.bindParamAsinSet(reg)

		// query array param asinSet
		if err := r.SetQueryParam("asinSet", joinedAsinSet...); err != nil {
			return err
		}
	}

	// path param contentReferenceKey
	if err := r.SetPathParam("contentReferenceKey", o.ContentReferenceKey); err != nil {
		return err
	}

	if o.IncludedDataSet != nil {

		// binding items for includedDataSet
		joinedIncludedDataSet := o.bindParamIncludedDataSet(reg)

		// query array param includedDataSet
		if err := r.SetQueryParam("includedDataSet", joinedIncludedDataSet...); err != nil {
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

// bindParamListContentDocumentAsinRelations binds the parameter asinSet
func (o *ListContentDocumentAsinRelationsParams) bindParamAsinSet(formats strfmt.Registry) []string {
	asinSetIR := o.AsinSet

	var asinSetIC []string
	for _, asinSetIIR := range asinSetIR { // explode []string

		asinSetIIV := asinSetIIR // string as string
		asinSetIC = append(asinSetIC, asinSetIIV)
	}

	// items.CollectionFormat: "csv"
	asinSetIS := swag.JoinByFormat(asinSetIC, "csv")

	return asinSetIS
}

// bindParamListContentDocumentAsinRelations binds the parameter includedDataSet
func (o *ListContentDocumentAsinRelationsParams) bindParamIncludedDataSet(formats strfmt.Registry) []string {
	includedDataSetIR := o.IncludedDataSet

	var includedDataSetIC []string
	for _, includedDataSetIIR := range includedDataSetIR { // explode []string

		includedDataSetIIV := includedDataSetIIR // string as string
		includedDataSetIC = append(includedDataSetIC, includedDataSetIIV)
	}

	// items.CollectionFormat: "csv"
	includedDataSetIS := swag.JoinByFormat(includedDataSetIC, "csv")

	return includedDataSetIS
}