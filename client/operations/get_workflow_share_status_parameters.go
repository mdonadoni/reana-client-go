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

// NewGetWorkflowShareStatusParams creates a new GetWorkflowShareStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowShareStatusParams() *GetWorkflowShareStatusParams {
	return &GetWorkflowShareStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowShareStatusParamsWithTimeout creates a new GetWorkflowShareStatusParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowShareStatusParamsWithTimeout(timeout time.Duration) *GetWorkflowShareStatusParams {
	return &GetWorkflowShareStatusParams{
		timeout: timeout,
	}
}

// NewGetWorkflowShareStatusParamsWithContext creates a new GetWorkflowShareStatusParams object
// with the ability to set a context for a request.
func NewGetWorkflowShareStatusParamsWithContext(ctx context.Context) *GetWorkflowShareStatusParams {
	return &GetWorkflowShareStatusParams{
		Context: ctx,
	}
}

// NewGetWorkflowShareStatusParamsWithHTTPClient creates a new GetWorkflowShareStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowShareStatusParamsWithHTTPClient(client *http.Client) *GetWorkflowShareStatusParams {
	return &GetWorkflowShareStatusParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowShareStatusParams contains all the parameters to send to the API endpoint

	for the get workflow share status operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowShareStatusParams struct {

	/* AccessToken.

	   The API access_token of workflow owner.
	*/
	AccessToken *string

	/* WorkflowIDOrName.

	   Required. Workflow UUID or name.
	*/
	WorkflowIDOrName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow share status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowShareStatusParams) WithDefaults() *GetWorkflowShareStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow share status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowShareStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow share status params
func (o *GetWorkflowShareStatusParams) WithTimeout(timeout time.Duration) *GetWorkflowShareStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow share status params
func (o *GetWorkflowShareStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow share status params
func (o *GetWorkflowShareStatusParams) WithContext(ctx context.Context) *GetWorkflowShareStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow share status params
func (o *GetWorkflowShareStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow share status params
func (o *GetWorkflowShareStatusParams) WithHTTPClient(client *http.Client) *GetWorkflowShareStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow share status params
func (o *GetWorkflowShareStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the get workflow share status params
func (o *GetWorkflowShareStatusParams) WithAccessToken(accessToken *string) *GetWorkflowShareStatusParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the get workflow share status params
func (o *GetWorkflowShareStatusParams) SetAccessToken(accessToken *string) {
	o.AccessToken = accessToken
}

// WithWorkflowIDOrName adds the workflowIDOrName to the get workflow share status params
func (o *GetWorkflowShareStatusParams) WithWorkflowIDOrName(workflowIDOrName string) *GetWorkflowShareStatusParams {
	o.SetWorkflowIDOrName(workflowIDOrName)
	return o
}

// SetWorkflowIDOrName adds the workflowIdOrName to the get workflow share status params
func (o *GetWorkflowShareStatusParams) SetWorkflowIDOrName(workflowIDOrName string) {
	o.WorkflowIDOrName = workflowIDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowShareStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workflow_id_or_name
	if err := r.SetPathParam("workflow_id_or_name", o.WorkflowIDOrName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}