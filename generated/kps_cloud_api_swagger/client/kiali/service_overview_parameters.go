// Code generated by go-swagger; DO NOT EDIT.

package kiali

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

// NewServiceOverviewParams creates a new ServiceOverviewParams object
// with the default values initialized.
func NewServiceOverviewParams() *ServiceOverviewParams {
	var ()
	return &ServiceOverviewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceOverviewParamsWithTimeout creates a new ServiceOverviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceOverviewParamsWithTimeout(timeout time.Duration) *ServiceOverviewParams {
	var ()
	return &ServiceOverviewParams{

		timeout: timeout,
	}
}

// NewServiceOverviewParamsWithContext creates a new ServiceOverviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceOverviewParamsWithContext(ctx context.Context) *ServiceOverviewParams {
	var ()
	return &ServiceOverviewParams{

		Context: ctx,
	}
}

// NewServiceOverviewParamsWithHTTPClient creates a new ServiceOverviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceOverviewParamsWithHTTPClient(client *http.Client) *ServiceOverviewParams {
	var ()
	return &ServiceOverviewParams{
		HTTPClient: client,
	}
}

/*ServiceOverviewParams contains all the parameters to send to the API endpoint
for the service overview operation typically these are written to a http.Request
*/
type ServiceOverviewParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*Service
	  The service name.

	*/
	Service string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service overview params
func (o *ServiceOverviewParams) WithTimeout(timeout time.Duration) *ServiceOverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service overview params
func (o *ServiceOverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service overview params
func (o *ServiceOverviewParams) WithContext(ctx context.Context) *ServiceOverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service overview params
func (o *ServiceOverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service overview params
func (o *ServiceOverviewParams) WithHTTPClient(client *http.Client) *ServiceOverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service overview params
func (o *ServiceOverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the service overview params
func (o *ServiceOverviewParams) WithAuthorization(authorization string) *ServiceOverviewParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the service overview params
func (o *ServiceOverviewParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the service overview params
func (o *ServiceOverviewParams) WithNamespace(namespace string) *ServiceOverviewParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the service overview params
func (o *ServiceOverviewParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithService adds the service to the service overview params
func (o *ServiceOverviewParams) WithService(service string) *ServiceOverviewParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the service overview params
func (o *ServiceOverviewParams) SetService(service string) {
	o.Service = service
}

// WithServiceDomain adds the serviceDomain to the service overview params
func (o *ServiceOverviewParams) WithServiceDomain(serviceDomain string) *ServiceOverviewParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the service overview params
func (o *ServiceOverviewParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceOverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param service
	if err := r.SetPathParam("service", o.Service); err != nil {
		return err
	}

	// query param serviceDomain
	qrServiceDomain := o.ServiceDomain
	qServiceDomain := qrServiceDomain
	if qServiceDomain != "" {
		if err := r.SetQueryParam("serviceDomain", qServiceDomain); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
