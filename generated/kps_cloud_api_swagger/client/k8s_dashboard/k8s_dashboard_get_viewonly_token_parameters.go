// Code generated by go-swagger; DO NOT EDIT.

package k8s_dashboard

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

// NewK8sDashboardGetViewonlyTokenParams creates a new K8sDashboardGetViewonlyTokenParams object
// with the default values initialized.
func NewK8sDashboardGetViewonlyTokenParams() *K8sDashboardGetViewonlyTokenParams {
	var ()
	return &K8sDashboardGetViewonlyTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewK8sDashboardGetViewonlyTokenParamsWithTimeout creates a new K8sDashboardGetViewonlyTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewK8sDashboardGetViewonlyTokenParamsWithTimeout(timeout time.Duration) *K8sDashboardGetViewonlyTokenParams {
	var ()
	return &K8sDashboardGetViewonlyTokenParams{

		timeout: timeout,
	}
}

// NewK8sDashboardGetViewonlyTokenParamsWithContext creates a new K8sDashboardGetViewonlyTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewK8sDashboardGetViewonlyTokenParamsWithContext(ctx context.Context) *K8sDashboardGetViewonlyTokenParams {
	var ()
	return &K8sDashboardGetViewonlyTokenParams{

		Context: ctx,
	}
}

// NewK8sDashboardGetViewonlyTokenParamsWithHTTPClient creates a new K8sDashboardGetViewonlyTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewK8sDashboardGetViewonlyTokenParamsWithHTTPClient(client *http.Client) *K8sDashboardGetViewonlyTokenParams {
	var ()
	return &K8sDashboardGetViewonlyTokenParams{
		HTTPClient: client,
	}
}

/*K8sDashboardGetViewonlyTokenParams contains all the parameters to send to the API endpoint
for the k8s dashboard get viewonly token operation typically these are written to a http.Request
*/
type K8sDashboardGetViewonlyTokenParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*SvcDomainID
	  ID for the service domain

	*/
	SvcDomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) WithTimeout(timeout time.Duration) *K8sDashboardGetViewonlyTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) WithContext(ctx context.Context) *K8sDashboardGetViewonlyTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) WithHTTPClient(client *http.Client) *K8sDashboardGetViewonlyTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) WithAuthorization(authorization string) *K8sDashboardGetViewonlyTokenParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSvcDomainID adds the svcDomainID to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) WithSvcDomainID(svcDomainID string) *K8sDashboardGetViewonlyTokenParams {
	o.SetSvcDomainID(svcDomainID)
	return o
}

// SetSvcDomainID adds the svcDomainId to the k8s dashboard get viewonly token params
func (o *K8sDashboardGetViewonlyTokenParams) SetSvcDomainID(svcDomainID string) {
	o.SvcDomainID = svcDomainID
}

// WriteToRequest writes these params to a swagger request
func (o *K8sDashboardGetViewonlyTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param svcDomainId
	if err := r.SetPathParam("svcDomainId", o.SvcDomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
