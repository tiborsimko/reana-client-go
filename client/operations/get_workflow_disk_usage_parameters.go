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

// NewGetWorkflowDiskUsageParams creates a new GetWorkflowDiskUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowDiskUsageParams() *GetWorkflowDiskUsageParams {
	return &GetWorkflowDiskUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowDiskUsageParamsWithTimeout creates a new GetWorkflowDiskUsageParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowDiskUsageParamsWithTimeout(timeout time.Duration) *GetWorkflowDiskUsageParams {
	return &GetWorkflowDiskUsageParams{
		timeout: timeout,
	}
}

// NewGetWorkflowDiskUsageParamsWithContext creates a new GetWorkflowDiskUsageParams object
// with the ability to set a context for a request.
func NewGetWorkflowDiskUsageParamsWithContext(ctx context.Context) *GetWorkflowDiskUsageParams {
	return &GetWorkflowDiskUsageParams{
		Context: ctx,
	}
}

// NewGetWorkflowDiskUsageParamsWithHTTPClient creates a new GetWorkflowDiskUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowDiskUsageParamsWithHTTPClient(client *http.Client) *GetWorkflowDiskUsageParams {
	return &GetWorkflowDiskUsageParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowDiskUsageParams contains all the parameters to send to the API endpoint

	for the get workflow disk usage operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowDiskUsageParams struct {

	/* AccessToken.

	   The API access_token of workflow owner.
	*/
	AccessToken *string

	/* Parameters.

	   Optional. Additional input parameters and operational options.
	*/
	Parameters GetWorkflowDiskUsageBody

	/* WorkflowIDOrName.

	   Required. Analysis UUID or name.
	*/
	WorkflowIDOrName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow disk usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowDiskUsageParams) WithDefaults() *GetWorkflowDiskUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow disk usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowDiskUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) WithTimeout(timeout time.Duration) *GetWorkflowDiskUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) WithContext(ctx context.Context) *GetWorkflowDiskUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) WithHTTPClient(client *http.Client) *GetWorkflowDiskUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) WithAccessToken(accessToken *string) *GetWorkflowDiskUsageParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WithParameters adds the parameters to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) WithParameters(parameters GetWorkflowDiskUsageBody) *GetWorkflowDiskUsageParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) SetParameters(parameters GetWorkflowDiskUsageBody) {
	o.Parameters = parameters
}

// WithWorkflowIDOrName adds the workflowIDOrName to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) WithWorkflowIDOrName(workflowIDOrName string) *GetWorkflowDiskUsageParams {
	o.SetWorkflowIDOrName(workflowIDOrName)
	return o
}

// SetWorkflowIDOrName adds the workflowIdOrName to the get workflow disk usage params
func (o *GetWorkflowDiskUsageParams) SetWorkflowIDOrName(workflowIDOrName string) {
	o.WorkflowIDOrName = workflowIDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowDiskUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Parameters); err != nil {
		return err
	}

	// path param workflow_id_or_name
	if err := r.SetPathParam("workflow_id_or_name", o.WorkflowIDOrName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
