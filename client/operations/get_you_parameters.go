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

// NewGetYouParams creates a new GetYouParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetYouParams() *GetYouParams {
	return &GetYouParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetYouParamsWithTimeout creates a new GetYouParams object
// with the ability to set a timeout on a request.
func NewGetYouParamsWithTimeout(timeout time.Duration) *GetYouParams {
	return &GetYouParams{
		timeout: timeout,
	}
}

// NewGetYouParamsWithContext creates a new GetYouParams object
// with the ability to set a context for a request.
func NewGetYouParamsWithContext(ctx context.Context) *GetYouParams {
	return &GetYouParams{
		Context: ctx,
	}
}

// NewGetYouParamsWithHTTPClient creates a new GetYouParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetYouParamsWithHTTPClient(client *http.Client) *GetYouParams {
	return &GetYouParams{
		HTTPClient: client,
	}
}

/*
GetYouParams contains all the parameters to send to the API endpoint

	for the get you operation.

	Typically these are written to a http.Request.
*/
type GetYouParams struct {

	/* AccessToken.

	   API access_token of user.
	*/
	AccessToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get you params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetYouParams) WithDefaults() *GetYouParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get you params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetYouParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get you params
func (o *GetYouParams) WithTimeout(timeout time.Duration) *GetYouParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get you params
func (o *GetYouParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get you params
func (o *GetYouParams) WithContext(ctx context.Context) *GetYouParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get you params
func (o *GetYouParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get you params
func (o *GetYouParams) WithHTTPClient(client *http.Client) *GetYouParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get you params
func (o *GetYouParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the get you params
func (o *GetYouParams) WithAccessToken(accessToken *string) *GetYouParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the get you params
func (o *GetYouParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetYouParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
