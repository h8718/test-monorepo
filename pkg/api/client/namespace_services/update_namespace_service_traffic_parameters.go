// Code generated by go-swagger; DO NOT EDIT.

package namespace_services

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

// NewUpdateNamespaceServiceTrafficParams creates a new UpdateNamespaceServiceTrafficParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNamespaceServiceTrafficParams() *UpdateNamespaceServiceTrafficParams {
	return &UpdateNamespaceServiceTrafficParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNamespaceServiceTrafficParamsWithTimeout creates a new UpdateNamespaceServiceTrafficParams object
// with the ability to set a timeout on a request.
func NewUpdateNamespaceServiceTrafficParamsWithTimeout(timeout time.Duration) *UpdateNamespaceServiceTrafficParams {
	return &UpdateNamespaceServiceTrafficParams{
		timeout: timeout,
	}
}

// NewUpdateNamespaceServiceTrafficParamsWithContext creates a new UpdateNamespaceServiceTrafficParams object
// with the ability to set a context for a request.
func NewUpdateNamespaceServiceTrafficParamsWithContext(ctx context.Context) *UpdateNamespaceServiceTrafficParams {
	return &UpdateNamespaceServiceTrafficParams{
		Context: ctx,
	}
}

// NewUpdateNamespaceServiceTrafficParamsWithHTTPClient creates a new UpdateNamespaceServiceTrafficParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNamespaceServiceTrafficParamsWithHTTPClient(client *http.Client) *UpdateNamespaceServiceTrafficParams {
	return &UpdateNamespaceServiceTrafficParams{
		HTTPClient: client,
	}
}

/* UpdateNamespaceServiceTrafficParams contains all the parameters to send to the API endpoint
   for the update namespace service traffic operation.

   Typically these are written to a http.Request.
*/
type UpdateNamespaceServiceTrafficParams struct {

	/* ServiceTraffic.

	   Payload that contains information on service traffic
	*/
	ServiceTraffic UpdateNamespaceServiceTrafficBody

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* ServiceName.

	   target service name
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update namespace service traffic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNamespaceServiceTrafficParams) WithDefaults() *UpdateNamespaceServiceTrafficParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update namespace service traffic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNamespaceServiceTrafficParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) WithTimeout(timeout time.Duration) *UpdateNamespaceServiceTrafficParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) WithContext(ctx context.Context) *UpdateNamespaceServiceTrafficParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) WithHTTPClient(client *http.Client) *UpdateNamespaceServiceTrafficParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceTraffic adds the serviceTraffic to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) WithServiceTraffic(serviceTraffic UpdateNamespaceServiceTrafficBody) *UpdateNamespaceServiceTrafficParams {
	o.SetServiceTraffic(serviceTraffic)
	return o
}

// SetServiceTraffic adds the serviceTraffic to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) SetServiceTraffic(serviceTraffic UpdateNamespaceServiceTrafficBody) {
	o.ServiceTraffic = serviceTraffic
}

// WithNamespace adds the namespace to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) WithNamespace(namespace string) *UpdateNamespaceServiceTrafficParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithServiceName adds the serviceName to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) WithServiceName(serviceName string) *UpdateNamespaceServiceTrafficParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the update namespace service traffic params
func (o *UpdateNamespaceServiceTrafficParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNamespaceServiceTrafficParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.ServiceTraffic); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
