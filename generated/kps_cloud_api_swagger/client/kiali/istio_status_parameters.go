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

// NewIstioStatusParams creates a new IstioStatusParams object
// with the default values initialized.
func NewIstioStatusParams() *IstioStatusParams {
	var ()
	return &IstioStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIstioStatusParamsWithTimeout creates a new IstioStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIstioStatusParamsWithTimeout(timeout time.Duration) *IstioStatusParams {
	var ()
	return &IstioStatusParams{

		timeout: timeout,
	}
}

// NewIstioStatusParamsWithContext creates a new IstioStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewIstioStatusParamsWithContext(ctx context.Context) *IstioStatusParams {
	var ()
	return &IstioStatusParams{

		Context: ctx,
	}
}

// NewIstioStatusParamsWithHTTPClient creates a new IstioStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIstioStatusParamsWithHTTPClient(client *http.Client) *IstioStatusParams {
	var ()
	return &IstioStatusParams{
		HTTPClient: client,
	}
}

/*IstioStatusParams contains all the parameters to send to the API endpoint
for the istio status operation typically these are written to a http.Request
*/
type IstioStatusParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the istio status params
func (o *IstioStatusParams) WithTimeout(timeout time.Duration) *IstioStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the istio status params
func (o *IstioStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the istio status params
func (o *IstioStatusParams) WithContext(ctx context.Context) *IstioStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the istio status params
func (o *IstioStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the istio status params
func (o *IstioStatusParams) WithHTTPClient(client *http.Client) *IstioStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the istio status params
func (o *IstioStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the istio status params
func (o *IstioStatusParams) WithAuthorization(authorization string) *IstioStatusParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the istio status params
func (o *IstioStatusParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithServiceDomain adds the serviceDomain to the istio status params
func (o *IstioStatusParams) WithServiceDomain(serviceDomain string) *IstioStatusParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the istio status params
func (o *IstioStatusParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *IstioStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
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
