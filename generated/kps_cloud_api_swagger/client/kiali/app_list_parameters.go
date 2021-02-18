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

// NewAppListParams creates a new AppListParams object
// with the default values initialized.
func NewAppListParams() *AppListParams {
	var ()
	return &AppListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppListParamsWithTimeout creates a new AppListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppListParamsWithTimeout(timeout time.Duration) *AppListParams {
	var ()
	return &AppListParams{

		timeout: timeout,
	}
}

// NewAppListParamsWithContext creates a new AppListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppListParamsWithContext(ctx context.Context) *AppListParams {
	var ()
	return &AppListParams{

		Context: ctx,
	}
}

// NewAppListParamsWithHTTPClient creates a new AppListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppListParamsWithHTTPClient(client *http.Client) *AppListParams {
	var ()
	return &AppListParams{
		HTTPClient: client,
	}
}

/*AppListParams contains all the parameters to send to the API endpoint
for the app list operation typically these are written to a http.Request
*/
type AppListParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the app list params
func (o *AppListParams) WithTimeout(timeout time.Duration) *AppListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app list params
func (o *AppListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app list params
func (o *AppListParams) WithContext(ctx context.Context) *AppListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app list params
func (o *AppListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app list params
func (o *AppListParams) WithHTTPClient(client *http.Client) *AppListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app list params
func (o *AppListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the app list params
func (o *AppListParams) WithAuthorization(authorization string) *AppListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the app list params
func (o *AppListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the app list params
func (o *AppListParams) WithNamespace(namespace string) *AppListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the app list params
func (o *AppListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithServiceDomain adds the serviceDomain to the app list params
func (o *AppListParams) WithServiceDomain(serviceDomain string) *AppListParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the app list params
func (o *AppListParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *AppListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
