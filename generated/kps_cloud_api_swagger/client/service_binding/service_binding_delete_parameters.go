// Code generated by go-swagger; DO NOT EDIT.

package service_binding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServiceBindingDeleteParams creates a new ServiceBindingDeleteParams object
// with the default values initialized.
func NewServiceBindingDeleteParams() *ServiceBindingDeleteParams {
	var ()
	return &ServiceBindingDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBindingDeleteParamsWithTimeout creates a new ServiceBindingDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBindingDeleteParamsWithTimeout(timeout time.Duration) *ServiceBindingDeleteParams {
	var ()
	return &ServiceBindingDeleteParams{

		timeout: timeout,
	}
}

// NewServiceBindingDeleteParamsWithContext creates a new ServiceBindingDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBindingDeleteParamsWithContext(ctx context.Context) *ServiceBindingDeleteParams {
	var ()
	return &ServiceBindingDeleteParams{

		Context: ctx,
	}
}

// NewServiceBindingDeleteParamsWithHTTPClient creates a new ServiceBindingDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBindingDeleteParamsWithHTTPClient(client *http.Client) *ServiceBindingDeleteParams {
	var ()
	return &ServiceBindingDeleteParams{
		HTTPClient: client,
	}
}

/*ServiceBindingDeleteParams contains all the parameters to send to the API endpoint
for the service binding delete operation typically these are written to a http.Request
*/
type ServiceBindingDeleteParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*SvcBindingID
	  Service Binding ID

	*/
	SvcBindingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service binding delete params
func (o *ServiceBindingDeleteParams) WithTimeout(timeout time.Duration) *ServiceBindingDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service binding delete params
func (o *ServiceBindingDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service binding delete params
func (o *ServiceBindingDeleteParams) WithContext(ctx context.Context) *ServiceBindingDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service binding delete params
func (o *ServiceBindingDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service binding delete params
func (o *ServiceBindingDeleteParams) WithHTTPClient(client *http.Client) *ServiceBindingDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service binding delete params
func (o *ServiceBindingDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the service binding delete params
func (o *ServiceBindingDeleteParams) WithAuthorization(authorization string) *ServiceBindingDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the service binding delete params
func (o *ServiceBindingDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSvcBindingID adds the svcBindingID to the service binding delete params
func (o *ServiceBindingDeleteParams) WithSvcBindingID(svcBindingID string) *ServiceBindingDeleteParams {
	o.SetSvcBindingID(svcBindingID)
	return o
}

// SetSvcBindingID adds the svcBindingId to the service binding delete params
func (o *ServiceBindingDeleteParams) SetSvcBindingID(svcBindingID string) {
	o.SvcBindingID = svcBindingID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBindingDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param svcBindingId
	if err := r.SetPathParam("svcBindingId", o.SvcBindingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
