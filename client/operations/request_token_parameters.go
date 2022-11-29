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

// NewRequestTokenParams creates a new RequestTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRequestTokenParams() *RequestTokenParams {
	return &RequestTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRequestTokenParamsWithTimeout creates a new RequestTokenParams object
// with the ability to set a timeout on a request.
func NewRequestTokenParamsWithTimeout(timeout time.Duration) *RequestTokenParams {
	return &RequestTokenParams{
		timeout: timeout,
	}
}

// NewRequestTokenParamsWithContext creates a new RequestTokenParams object
// with the ability to set a context for a request.
func NewRequestTokenParamsWithContext(ctx context.Context) *RequestTokenParams {
	return &RequestTokenParams{
		Context: ctx,
	}
}

// NewRequestTokenParamsWithHTTPClient creates a new RequestTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewRequestTokenParamsWithHTTPClient(client *http.Client) *RequestTokenParams {
	return &RequestTokenParams{
		HTTPClient: client,
	}
}

/*
RequestTokenParams contains all the parameters to send to the API endpoint

	for the request token operation.

	Typically these are written to a http.Request.
*/
type RequestTokenParams struct {

	/* AccessToken.

	   API access_token of user.
	*/
	AccessToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the request token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestTokenParams) WithDefaults() *RequestTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the request token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the request token params
func (o *RequestTokenParams) WithTimeout(timeout time.Duration) *RequestTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request token params
func (o *RequestTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request token params
func (o *RequestTokenParams) WithContext(ctx context.Context) *RequestTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request token params
func (o *RequestTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request token params
func (o *RequestTokenParams) WithHTTPClient(client *http.Client) *RequestTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request token params
func (o *RequestTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the request token params
func (o *RequestTokenParams) WithAccessToken(accessToken *string) *RequestTokenParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the request token params
func (o *RequestTokenParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WriteToRequest writes these params to a swagger request
func (o *RequestTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessToken != nil {

		// query param access_token
		var qrAccessToken string

		if o.AccessToken != nil {
			qrAccessToken = *o.AccessToken
		}
		qAccessToken := qrAccessToken
		if qAccessToken != "" {

			if err := r.SetQueryParam("access_token", qAccessToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
