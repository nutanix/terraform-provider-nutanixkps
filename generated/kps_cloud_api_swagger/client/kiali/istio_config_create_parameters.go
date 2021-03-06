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

// NewIstioConfigCreateParams creates a new IstioConfigCreateParams object
// with the default values initialized.
func NewIstioConfigCreateParams() *IstioConfigCreateParams {
	var ()
	return &IstioConfigCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIstioConfigCreateParamsWithTimeout creates a new IstioConfigCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIstioConfigCreateParamsWithTimeout(timeout time.Duration) *IstioConfigCreateParams {
	var ()
	return &IstioConfigCreateParams{

		timeout: timeout,
	}
}

// NewIstioConfigCreateParamsWithContext creates a new IstioConfigCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIstioConfigCreateParamsWithContext(ctx context.Context) *IstioConfigCreateParams {
	var ()
	return &IstioConfigCreateParams{

		Context: ctx,
	}
}

// NewIstioConfigCreateParamsWithHTTPClient creates a new IstioConfigCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIstioConfigCreateParamsWithHTTPClient(client *http.Client) *IstioConfigCreateParams {
	var ()
	return &IstioConfigCreateParams{
		HTTPClient: client,
	}
}

/*IstioConfigCreateParams contains all the parameters to send to the API endpoint
for the istio config create operation typically these are written to a http.Request
*/
type IstioConfigCreateParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*ObjectType
	  The Istio object type.

	*/
	ObjectType string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the istio config create params
func (o *IstioConfigCreateParams) WithTimeout(timeout time.Duration) *IstioConfigCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the istio config create params
func (o *IstioConfigCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the istio config create params
func (o *IstioConfigCreateParams) WithContext(ctx context.Context) *IstioConfigCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the istio config create params
func (o *IstioConfigCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the istio config create params
func (o *IstioConfigCreateParams) WithHTTPClient(client *http.Client) *IstioConfigCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the istio config create params
func (o *IstioConfigCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the istio config create params
func (o *IstioConfigCreateParams) WithAuthorization(authorization string) *IstioConfigCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the istio config create params
func (o *IstioConfigCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the istio config create params
func (o *IstioConfigCreateParams) WithNamespace(namespace string) *IstioConfigCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the istio config create params
func (o *IstioConfigCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithObjectType adds the objectType to the istio config create params
func (o *IstioConfigCreateParams) WithObjectType(objectType string) *IstioConfigCreateParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the istio config create params
func (o *IstioConfigCreateParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WithServiceDomain adds the serviceDomain to the istio config create params
func (o *IstioConfigCreateParams) WithServiceDomain(serviceDomain string) *IstioConfigCreateParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the istio config create params
func (o *IstioConfigCreateParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *IstioConfigCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param object_type
	if err := r.SetPathParam("object_type", o.ObjectType); err != nil {
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
