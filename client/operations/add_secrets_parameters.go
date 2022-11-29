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

// NewAddSecretsParams creates a new AddSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSecretsParams() *AddSecretsParams {
	return &AddSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSecretsParamsWithTimeout creates a new AddSecretsParams object
// with the ability to set a timeout on a request.
func NewAddSecretsParamsWithTimeout(timeout time.Duration) *AddSecretsParams {
	return &AddSecretsParams{
		timeout: timeout,
	}
}

// NewAddSecretsParamsWithContext creates a new AddSecretsParams object
// with the ability to set a context for a request.
func NewAddSecretsParamsWithContext(ctx context.Context) *AddSecretsParams {
	return &AddSecretsParams{
		Context: ctx,
	}
}

// NewAddSecretsParamsWithHTTPClient creates a new AddSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddSecretsParamsWithHTTPClient(client *http.Client) *AddSecretsParams {
	return &AddSecretsParams{
		HTTPClient: client,
	}
}

/*
AddSecretsParams contains all the parameters to send to the API endpoint

	for the add secrets operation.

	Typically these are written to a http.Request.
*/
type AddSecretsParams struct {

	/* AccessToken.

	   Secrets owner access token.
	*/
	AccessToken *string

	/* Overwrite.

	   Whether existing secret keys should be overwritten.
	*/
	Overwrite *bool

	/* Secrets.

	   Optional. List of secrets to be added.
	*/
	Secrets map[string]AddSecretsParamsBodyAnon

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSecretsParams) WithDefaults() *AddSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add secrets params
func (o *AddSecretsParams) WithTimeout(timeout time.Duration) *AddSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add secrets params
func (o *AddSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add secrets params
func (o *AddSecretsParams) WithContext(ctx context.Context) *AddSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add secrets params
func (o *AddSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add secrets params
func (o *AddSecretsParams) WithHTTPClient(client *http.Client) *AddSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add secrets params
func (o *AddSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the add secrets params
func (o *AddSecretsParams) WithAccessToken(accessToken *string) *AddSecretsParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the add secrets params
func (o *AddSecretsParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WithOverwrite adds the overwrite to the add secrets params
func (o *AddSecretsParams) WithOverwrite(overwrite *bool) *AddSecretsParams {
	o.SetOverwrite(overwrite)
	return o
}

// SetOverwrite adds the overwrite to the add secrets params
func (o *AddSecretsParams) SetOverwrite(overwrite *bool) {
	o.Overwrite = overwrite
}

// WithSecrets adds the secrets to the add secrets params
func (o *AddSecretsParams) WithSecrets(secrets map[string]AddSecretsParamsBodyAnon) *AddSecretsParams {
	o.SetSecrets(secrets)
	return o
}

// SetSecrets adds the secrets to the add secrets params
func (o *AddSecretsParams) SetSecrets(secrets map[string]AddSecretsParamsBodyAnon) {
	o.Secrets = secrets
}

// WriteToRequest writes these params to a swagger request
func (o *AddSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Overwrite != nil {

		// query param overwrite
		var qrOverwrite bool

		if o.Overwrite != nil {
			qrOverwrite = *o.Overwrite
		}
		qOverwrite := swag.FormatBool(qrOverwrite)
		if qOverwrite != "" {

			if err := r.SetQueryParam("overwrite", qOverwrite); err != nil {
				return err
			}
		}
	}
	if o.Secrets != nil {
		if err := r.SetBodyParam(o.Secrets); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
