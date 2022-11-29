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

// NewGitlabProjectsParams creates a new GitlabProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGitlabProjectsParams() *GitlabProjectsParams {
	return &GitlabProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGitlabProjectsParamsWithTimeout creates a new GitlabProjectsParams object
// with the ability to set a timeout on a request.
func NewGitlabProjectsParamsWithTimeout(timeout time.Duration) *GitlabProjectsParams {
	return &GitlabProjectsParams{
		timeout: timeout,
	}
}

// NewGitlabProjectsParamsWithContext creates a new GitlabProjectsParams object
// with the ability to set a context for a request.
func NewGitlabProjectsParamsWithContext(ctx context.Context) *GitlabProjectsParams {
	return &GitlabProjectsParams{
		Context: ctx,
	}
}

// NewGitlabProjectsParamsWithHTTPClient creates a new GitlabProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGitlabProjectsParamsWithHTTPClient(client *http.Client) *GitlabProjectsParams {
	return &GitlabProjectsParams{
		HTTPClient: client,
	}
}

/*
GitlabProjectsParams contains all the parameters to send to the API endpoint

	for the gitlab projects operation.

	Typically these are written to a http.Request.
*/
type GitlabProjectsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gitlab projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GitlabProjectsParams) WithDefaults() *GitlabProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gitlab projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GitlabProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gitlab projects params
func (o *GitlabProjectsParams) WithTimeout(timeout time.Duration) *GitlabProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gitlab projects params
func (o *GitlabProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gitlab projects params
func (o *GitlabProjectsParams) WithContext(ctx context.Context) *GitlabProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gitlab projects params
func (o *GitlabProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gitlab projects params
func (o *GitlabProjectsParams) WithHTTPClient(client *http.Client) *GitlabProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gitlab projects params
func (o *GitlabProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GitlabProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
