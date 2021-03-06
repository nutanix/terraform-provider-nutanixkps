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

// NewSpansListParams creates a new SpansListParams object
// with the default values initialized.
func NewSpansListParams() *SpansListParams {
	var ()
	return &SpansListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpansListParamsWithTimeout creates a new SpansListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpansListParamsWithTimeout(timeout time.Duration) *SpansListParams {
	var ()
	return &SpansListParams{

		timeout: timeout,
	}
}

// NewSpansListParamsWithContext creates a new SpansListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSpansListParamsWithContext(ctx context.Context) *SpansListParams {
	var ()
	return &SpansListParams{

		Context: ctx,
	}
}

// NewSpansListParamsWithHTTPClient creates a new SpansListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSpansListParamsWithHTTPClient(client *http.Client) *SpansListParams {
	var ()
	return &SpansListParams{
		HTTPClient: client,
	}
}

/*SpansListParams contains all the parameters to send to the API endpoint
for the spans list operation typically these are written to a http.Request
*/
type SpansListParams struct {

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

// WithTimeout adds the timeout to the spans list params
func (o *SpansListParams) WithTimeout(timeout time.Duration) *SpansListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the spans list params
func (o *SpansListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the spans list params
func (o *SpansListParams) WithContext(ctx context.Context) *SpansListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the spans list params
func (o *SpansListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the spans list params
func (o *SpansListParams) WithHTTPClient(client *http.Client) *SpansListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the spans list params
func (o *SpansListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the spans list params
func (o *SpansListParams) WithAuthorization(authorization string) *SpansListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the spans list params
func (o *SpansListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the spans list params
func (o *SpansListParams) WithNamespace(namespace string) *SpansListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the spans list params
func (o *SpansListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithService adds the service to the spans list params
func (o *SpansListParams) WithService(service string) *SpansListParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the spans list params
func (o *SpansListParams) SetService(service string) {
	o.Service = service
}

// WithServiceDomain adds the serviceDomain to the spans list params
func (o *SpansListParams) WithServiceDomain(serviceDomain string) *SpansListParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the spans list params
func (o *SpansListParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *SpansListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
