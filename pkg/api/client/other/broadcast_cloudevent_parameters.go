// Code generated by go-swagger; DO NOT EDIT.

package other

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

// NewBroadcastCloudeventParams creates a new BroadcastCloudeventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBroadcastCloudeventParams() *BroadcastCloudeventParams {
	return &BroadcastCloudeventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBroadcastCloudeventParamsWithTimeout creates a new BroadcastCloudeventParams object
// with the ability to set a timeout on a request.
func NewBroadcastCloudeventParamsWithTimeout(timeout time.Duration) *BroadcastCloudeventParams {
	return &BroadcastCloudeventParams{
		timeout: timeout,
	}
}

// NewBroadcastCloudeventParamsWithContext creates a new BroadcastCloudeventParams object
// with the ability to set a context for a request.
func NewBroadcastCloudeventParamsWithContext(ctx context.Context) *BroadcastCloudeventParams {
	return &BroadcastCloudeventParams{
		Context: ctx,
	}
}

// NewBroadcastCloudeventParamsWithHTTPClient creates a new BroadcastCloudeventParams object
// with the ability to set a custom HTTPClient for a request.
func NewBroadcastCloudeventParamsWithHTTPClient(client *http.Client) *BroadcastCloudeventParams {
	return &BroadcastCloudeventParams{
		HTTPClient: client,
	}
}

/* BroadcastCloudeventParams contains all the parameters to send to the API endpoint
   for the broadcast cloudevent operation.

   Typically these are written to a http.Request.
*/
type BroadcastCloudeventParams struct {

	/* Cloudevent.

	   Cloud Event request to be sent.
	*/
	Cloudevent interface{}

	/* Namespace.

	   target namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the broadcast cloudevent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BroadcastCloudeventParams) WithDefaults() *BroadcastCloudeventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the broadcast cloudevent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BroadcastCloudeventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) WithTimeout(timeout time.Duration) *BroadcastCloudeventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) WithContext(ctx context.Context) *BroadcastCloudeventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) WithHTTPClient(client *http.Client) *BroadcastCloudeventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudevent adds the cloudevent to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) WithCloudevent(cloudevent interface{}) *BroadcastCloudeventParams {
	o.SetCloudevent(cloudevent)
	return o
}

// SetCloudevent adds the cloudevent to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) SetCloudevent(cloudevent interface{}) {
	o.Cloudevent = cloudevent
}

// WithNamespace adds the namespace to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) WithNamespace(namespace string) *BroadcastCloudeventParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the broadcast cloudevent params
func (o *BroadcastCloudeventParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BroadcastCloudeventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Cloudevent != nil {
		if err := r.SetBodyParam(o.Cloudevent); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
