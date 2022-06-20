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
)

// NewStatusParams creates a new StatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusParams() *StatusParams {
	return &StatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusParamsWithTimeout creates a new StatusParams object
// with the ability to set a timeout on a request.
func NewStatusParamsWithTimeout(timeout time.Duration) *StatusParams {
	return &StatusParams{
		timeout: timeout,
	}
}

// NewStatusParamsWithContext creates a new StatusParams object
// with the ability to set a context for a request.
func NewStatusParamsWithContext(ctx context.Context) *StatusParams {
	return &StatusParams{
		Context: ctx,
	}
}

// NewStatusParamsWithHTTPClient creates a new StatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusParamsWithHTTPClient(client *http.Client) *StatusParams {
	return &StatusParams{
		HTTPClient: client,
	}
}

/* StatusParams contains all the parameters to send to the API endpoint
   for the status operation.

   Typically these are written to a http.Request.
*/
type StatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusParams) WithDefaults() *StatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status params
func (o *StatusParams) WithTimeout(timeout time.Duration) *StatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status params
func (o *StatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status params
func (o *StatusParams) WithContext(ctx context.Context) *StatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status params
func (o *StatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status params
func (o *StatusParams) WithHTTPClient(client *http.Client) *StatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status params
func (o *StatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
