// Code generated by go-swagger; DO NOT EDIT.

package service_instance

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

// NewServiceInstanceGetParams creates a new ServiceInstanceGetParams object
// with the default values initialized.
func NewServiceInstanceGetParams() *ServiceInstanceGetParams {
	var ()
	return &ServiceInstanceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceInstanceGetParamsWithTimeout creates a new ServiceInstanceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceInstanceGetParamsWithTimeout(timeout time.Duration) *ServiceInstanceGetParams {
	var ()
	return &ServiceInstanceGetParams{

		timeout: timeout,
	}
}

// NewServiceInstanceGetParamsWithContext creates a new ServiceInstanceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceInstanceGetParamsWithContext(ctx context.Context) *ServiceInstanceGetParams {
	var ()
	return &ServiceInstanceGetParams{

		Context: ctx,
	}
}

// NewServiceInstanceGetParamsWithHTTPClient creates a new ServiceInstanceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceInstanceGetParamsWithHTTPClient(client *http.Client) *ServiceInstanceGetParams {
	var ()
	return &ServiceInstanceGetParams{
		HTTPClient: client,
	}
}

/*ServiceInstanceGetParams contains all the parameters to send to the API endpoint
for the service instance get operation typically these are written to a http.Request
*/
type ServiceInstanceGetParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*SvcInstanceID
	  Service Instance ID

	*/
	SvcInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service instance get params
func (o *ServiceInstanceGetParams) WithTimeout(timeout time.Duration) *ServiceInstanceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service instance get params
func (o *ServiceInstanceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service instance get params
func (o *ServiceInstanceGetParams) WithContext(ctx context.Context) *ServiceInstanceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service instance get params
func (o *ServiceInstanceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service instance get params
func (o *ServiceInstanceGetParams) WithHTTPClient(client *http.Client) *ServiceInstanceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service instance get params
func (o *ServiceInstanceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the service instance get params
func (o *ServiceInstanceGetParams) WithAuthorization(authorization string) *ServiceInstanceGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the service instance get params
func (o *ServiceInstanceGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSvcInstanceID adds the svcInstanceID to the service instance get params
func (o *ServiceInstanceGetParams) WithSvcInstanceID(svcInstanceID string) *ServiceInstanceGetParams {
	o.SetSvcInstanceID(svcInstanceID)
	return o
}

// SetSvcInstanceID adds the svcInstanceId to the service instance get params
func (o *ServiceInstanceGetParams) SetSvcInstanceID(svcInstanceID string) {
	o.SvcInstanceID = svcInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceInstanceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param svcInstanceId
	if err := r.SetPathParam("svcInstanceId", o.SvcInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
