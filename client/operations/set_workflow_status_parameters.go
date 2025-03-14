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

// NewSetWorkflowStatusParams creates a new SetWorkflowStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetWorkflowStatusParams() *SetWorkflowStatusParams {
	return &SetWorkflowStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetWorkflowStatusParamsWithTimeout creates a new SetWorkflowStatusParams object
// with the ability to set a timeout on a request.
func NewSetWorkflowStatusParamsWithTimeout(timeout time.Duration) *SetWorkflowStatusParams {
	return &SetWorkflowStatusParams{
		timeout: timeout,
	}
}

// NewSetWorkflowStatusParamsWithContext creates a new SetWorkflowStatusParams object
// with the ability to set a context for a request.
func NewSetWorkflowStatusParamsWithContext(ctx context.Context) *SetWorkflowStatusParams {
	return &SetWorkflowStatusParams{
		Context: ctx,
	}
}

// NewSetWorkflowStatusParamsWithHTTPClient creates a new SetWorkflowStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetWorkflowStatusParamsWithHTTPClient(client *http.Client) *SetWorkflowStatusParams {
	return &SetWorkflowStatusParams{
		HTTPClient: client,
	}
}

/*
SetWorkflowStatusParams contains all the parameters to send to the API endpoint

	for the set workflow status operation.

	Typically these are written to a http.Request.
*/
type SetWorkflowStatusParams struct {

	/* AccessToken.

	   The API access_token of workflow owner.
	*/
	AccessToken *string

	/* Parameters.

	   Optional. Additional parameters to customise the workflow status change.
	*/
	Parameters SetWorkflowStatusBody

	/* Status.

	   Required. New workflow status.
	*/
	Status string

	/* WorkflowIDOrName.

	   Required. Analysis UUID or name.
	*/
	WorkflowIDOrName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set workflow status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetWorkflowStatusParams) WithDefaults() *SetWorkflowStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set workflow status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetWorkflowStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set workflow status params
func (o *SetWorkflowStatusParams) WithTimeout(timeout time.Duration) *SetWorkflowStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set workflow status params
func (o *SetWorkflowStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set workflow status params
func (o *SetWorkflowStatusParams) WithContext(ctx context.Context) *SetWorkflowStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set workflow status params
func (o *SetWorkflowStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set workflow status params
func (o *SetWorkflowStatusParams) WithHTTPClient(client *http.Client) *SetWorkflowStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set workflow status params
func (o *SetWorkflowStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the set workflow status params
func (o *SetWorkflowStatusParams) WithAccessToken(accessToken *string) *SetWorkflowStatusParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the set workflow status params
func (o *SetWorkflowStatusParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WithParameters adds the parameters to the set workflow status params
func (o *SetWorkflowStatusParams) WithParameters(parameters SetWorkflowStatusBody) *SetWorkflowStatusParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the set workflow status params
func (o *SetWorkflowStatusParams) SetParameters(parameters SetWorkflowStatusBody) {
	o.Parameters = parameters
}

// WithStatus adds the status to the set workflow status params
func (o *SetWorkflowStatusParams) WithStatus(status string) *SetWorkflowStatusParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the set workflow status params
func (o *SetWorkflowStatusParams) SetStatus(status string) {
	o.Status = status
}

// WithWorkflowIDOrName adds the workflowIDOrName to the set workflow status params
func (o *SetWorkflowStatusParams) WithWorkflowIDOrName(workflowIDOrName string) *SetWorkflowStatusParams {
	o.SetWorkflowIDOrName(workflowIDOrName)
	return o
}

// SetWorkflowIDOrName adds the workflowIdOrName to the set workflow status params
func (o *SetWorkflowStatusParams) SetWorkflowIDOrName(workflowIDOrName string) {
	o.WorkflowIDOrName = workflowIDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *SetWorkflowStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param status
	qrStatus := o.Status
	qStatus := qrStatus
	if qStatus != "" {

		if err := r.SetQueryParam("status", qStatus); err != nil {
			return err
		}
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
