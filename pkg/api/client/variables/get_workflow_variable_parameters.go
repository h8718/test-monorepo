// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewGetWorkflowVariableParams creates a new GetWorkflowVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowVariableParams() *GetWorkflowVariableParams {
	return &GetWorkflowVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowVariableParamsWithTimeout creates a new GetWorkflowVariableParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowVariableParamsWithTimeout(timeout time.Duration) *GetWorkflowVariableParams {
	return &GetWorkflowVariableParams{
		timeout: timeout,
	}
}

// NewGetWorkflowVariableParamsWithContext creates a new GetWorkflowVariableParams object
// with the ability to set a context for a request.
func NewGetWorkflowVariableParamsWithContext(ctx context.Context) *GetWorkflowVariableParams {
	return &GetWorkflowVariableParams{
		Context: ctx,
	}
}

// NewGetWorkflowVariableParamsWithHTTPClient creates a new GetWorkflowVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowVariableParamsWithHTTPClient(client *http.Client) *GetWorkflowVariableParams {
	return &GetWorkflowVariableParams{
		HTTPClient: client,
	}
}

/* GetWorkflowVariableParams contains all the parameters to send to the API endpoint
   for the get workflow variable operation.

   Typically these are written to a http.Request.
*/
type GetWorkflowVariableParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Var.

	   target variable
	*/
	Var string

	/* Workflow.

	   path to target workflow
	*/
	Workflow string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowVariableParams) WithDefaults() *GetWorkflowVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow variable params
func (o *GetWorkflowVariableParams) WithTimeout(timeout time.Duration) *GetWorkflowVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow variable params
func (o *GetWorkflowVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow variable params
func (o *GetWorkflowVariableParams) WithContext(ctx context.Context) *GetWorkflowVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow variable params
func (o *GetWorkflowVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow variable params
func (o *GetWorkflowVariableParams) WithHTTPClient(client *http.Client) *GetWorkflowVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow variable params
func (o *GetWorkflowVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get workflow variable params
func (o *GetWorkflowVariableParams) WithNamespace(namespace string) *GetWorkflowVariableParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get workflow variable params
func (o *GetWorkflowVariableParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVar adds the varVar to the get workflow variable params
func (o *GetWorkflowVariableParams) WithVar(varVar string) *GetWorkflowVariableParams {
	o.SetVar(varVar)
	return o
}

// SetVar adds the var to the get workflow variable params
func (o *GetWorkflowVariableParams) SetVar(varVar string) {
	o.Var = varVar
}

// WithWorkflow adds the workflow to the get workflow variable params
func (o *GetWorkflowVariableParams) WithWorkflow(workflow string) *GetWorkflowVariableParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the get workflow variable params
func (o *GetWorkflowVariableParams) SetWorkflow(workflow string) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param var
	qrVar := o.Var
	qVar := qrVar
	if qVar != "" {

		if err := r.SetQueryParam("var", qVar); err != nil {
			return err
		}
	}

	// path param workflow
	if err := r.SetPathParam("workflow", o.Workflow); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
