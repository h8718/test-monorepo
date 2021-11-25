// Code generated by go-swagger; DO NOT EDIT.

package registries

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

// NewGetGlobalPrivateRegistriesParams creates a new GetGlobalPrivateRegistriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalPrivateRegistriesParams() *GetGlobalPrivateRegistriesParams {
	return &GetGlobalPrivateRegistriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalPrivateRegistriesParamsWithTimeout creates a new GetGlobalPrivateRegistriesParams object
// with the ability to set a timeout on a request.
func NewGetGlobalPrivateRegistriesParamsWithTimeout(timeout time.Duration) *GetGlobalPrivateRegistriesParams {
	return &GetGlobalPrivateRegistriesParams{
		timeout: timeout,
	}
}

// NewGetGlobalPrivateRegistriesParamsWithContext creates a new GetGlobalPrivateRegistriesParams object
// with the ability to set a context for a request.
func NewGetGlobalPrivateRegistriesParamsWithContext(ctx context.Context) *GetGlobalPrivateRegistriesParams {
	return &GetGlobalPrivateRegistriesParams{
		Context: ctx,
	}
}

// NewGetGlobalPrivateRegistriesParamsWithHTTPClient creates a new GetGlobalPrivateRegistriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalPrivateRegistriesParamsWithHTTPClient(client *http.Client) *GetGlobalPrivateRegistriesParams {
	return &GetGlobalPrivateRegistriesParams{
		HTTPClient: client,
	}
}

/* GetGlobalPrivateRegistriesParams contains all the parameters to send to the API endpoint
   for the get global private registries operation.

   Typically these are written to a http.Request.
*/
type GetGlobalPrivateRegistriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global private registries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalPrivateRegistriesParams) WithDefaults() *GetGlobalPrivateRegistriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global private registries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalPrivateRegistriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global private registries params
func (o *GetGlobalPrivateRegistriesParams) WithTimeout(timeout time.Duration) *GetGlobalPrivateRegistriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global private registries params
func (o *GetGlobalPrivateRegistriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global private registries params
func (o *GetGlobalPrivateRegistriesParams) WithContext(ctx context.Context) *GetGlobalPrivateRegistriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global private registries params
func (o *GetGlobalPrivateRegistriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global private registries params
func (o *GetGlobalPrivateRegistriesParams) WithHTTPClient(client *http.Client) *GetGlobalPrivateRegistriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global private registries params
func (o *GetGlobalPrivateRegistriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalPrivateRegistriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
