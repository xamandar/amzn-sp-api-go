// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewListFinancialEventsByGroupIDParams creates a new ListFinancialEventsByGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListFinancialEventsByGroupIDParams() *ListFinancialEventsByGroupIDParams {
	return &ListFinancialEventsByGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListFinancialEventsByGroupIDParamsWithTimeout creates a new ListFinancialEventsByGroupIDParams object
// with the ability to set a timeout on a request.
func NewListFinancialEventsByGroupIDParamsWithTimeout(timeout time.Duration) *ListFinancialEventsByGroupIDParams {
	return &ListFinancialEventsByGroupIDParams{
		timeout: timeout,
	}
}

// NewListFinancialEventsByGroupIDParamsWithContext creates a new ListFinancialEventsByGroupIDParams object
// with the ability to set a context for a request.
func NewListFinancialEventsByGroupIDParamsWithContext(ctx context.Context) *ListFinancialEventsByGroupIDParams {
	return &ListFinancialEventsByGroupIDParams{
		Context: ctx,
	}
}

// NewListFinancialEventsByGroupIDParamsWithHTTPClient creates a new ListFinancialEventsByGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewListFinancialEventsByGroupIDParamsWithHTTPClient(client *http.Client) *ListFinancialEventsByGroupIDParams {
	return &ListFinancialEventsByGroupIDParams{
		HTTPClient: client,
	}
}

/* ListFinancialEventsByGroupIDParams contains all the parameters to send to the API endpoint
   for the list financial events by group Id operation.

   Typically these are written to a http.Request.
*/
type ListFinancialEventsByGroupIDParams struct {

	/* MaxResultsPerPage.

	   The maximum number of results to return per page.

	   Format: int32
	   Default: 100
	*/
	MaxResultsPerPage *int32

	/* NextToken.

	   A string token returned in the response of your previous request.
	*/
	NextToken *string

	/* EventGroupID.

	   The identifier of the financial event group to which the events belong.
	*/
	EventGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list financial events by group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFinancialEventsByGroupIDParams) WithDefaults() *ListFinancialEventsByGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list financial events by group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFinancialEventsByGroupIDParams) SetDefaults() {
	var (
		maxResultsPerPageDefault = int32(100)
	)

	val := ListFinancialEventsByGroupIDParams{
		MaxResultsPerPage: &maxResultsPerPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) WithTimeout(timeout time.Duration) *ListFinancialEventsByGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) WithContext(ctx context.Context) *ListFinancialEventsByGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) WithHTTPClient(client *http.Client) *ListFinancialEventsByGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaxResultsPerPage adds the maxResultsPerPage to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) WithMaxResultsPerPage(maxResultsPerPage *int32) *ListFinancialEventsByGroupIDParams {
	o.SetMaxResultsPerPage(maxResultsPerPage)
	return o
}

// SetMaxResultsPerPage adds the maxResultsPerPage to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) SetMaxResultsPerPage(maxResultsPerPage *int32) {
	o.MaxResultsPerPage = maxResultsPerPage
}

// WithNextToken adds the nextToken to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) WithNextToken(nextToken *string) *ListFinancialEventsByGroupIDParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithEventGroupID adds the eventGroupID to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) WithEventGroupID(eventGroupID string) *ListFinancialEventsByGroupIDParams {
	o.SetEventGroupID(eventGroupID)
	return o
}

// SetEventGroupID adds the eventGroupId to the list financial events by group Id params
func (o *ListFinancialEventsByGroupIDParams) SetEventGroupID(eventGroupID string) {
	o.EventGroupID = eventGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *ListFinancialEventsByGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MaxResultsPerPage != nil {

		// query param MaxResultsPerPage
		var qrMaxResultsPerPage int32

		if o.MaxResultsPerPage != nil {
			qrMaxResultsPerPage = *o.MaxResultsPerPage
		}
		qMaxResultsPerPage := swag.FormatInt32(qrMaxResultsPerPage)
		if qMaxResultsPerPage != "" {

			if err := r.SetQueryParam("MaxResultsPerPage", qMaxResultsPerPage); err != nil {
				return err
			}
		}
	}

	if o.NextToken != nil {

		// query param NextToken
		var qrNextToken string

		if o.NextToken != nil {
			qrNextToken = *o.NextToken
		}
		qNextToken := qrNextToken
		if qNextToken != "" {

			if err := r.SetQueryParam("NextToken", qNextToken); err != nil {
				return err
			}
		}
	}

	// path param eventGroupId
	if err := r.SetPathParam("eventGroupId", o.EventGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}