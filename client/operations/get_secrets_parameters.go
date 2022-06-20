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

// NewGetSecretsParams creates a new GetSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSecretsParams() *GetSecretsParams {
	return &GetSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecretsParamsWithTimeout creates a new GetSecretsParams object
// with the ability to set a timeout on a request.
func NewGetSecretsParamsWithTimeout(timeout time.Duration) *GetSecretsParams {
	return &GetSecretsParams{
		timeout: timeout,
	}
}

// NewGetSecretsParamsWithContext creates a new GetSecretsParams object
// with the ability to set a context for a request.
func NewGetSecretsParamsWithContext(ctx context.Context) *GetSecretsParams {
	return &GetSecretsParams{
		Context: ctx,
	}
}

// NewGetSecretsParamsWithHTTPClient creates a new GetSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSecretsParamsWithHTTPClient(client *http.Client) *GetSecretsParams {
	return &GetSecretsParams{
		HTTPClient: client,
	}
}

/* GetSecretsParams contains all the parameters to send to the API endpoint
   for the get secrets operation.

   Typically these are written to a http.Request.
*/
type GetSecretsParams struct {

	/* AccessToken.

	   Secrets owner access token.
	*/
	AccessToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecretsParams) WithDefaults() *GetSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get secrets params
func (o *GetSecretsParams) WithTimeout(timeout time.Duration) *GetSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get secrets params
func (o *GetSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get secrets params
func (o *GetSecretsParams) WithContext(ctx context.Context) *GetSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get secrets params
func (o *GetSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get secrets params
func (o *GetSecretsParams) WithHTTPClient(client *http.Client) *GetSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get secrets params
func (o *GetSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the get secrets params
func (o *GetSecretsParams) WithAccessToken(accessToken *string) *GetSecretsParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the get secrets params
func (o *GetSecretsParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
