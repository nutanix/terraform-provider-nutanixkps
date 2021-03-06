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

// NewK8sDashboardGetAdminTokenParams creates a new K8sDashboardGetAdminTokenParams object
// with the default values initialized.
func NewK8sDashboardGetAdminTokenParams() *K8sDashboardGetAdminTokenParams {
	var ()
	return &K8sDashboardGetAdminTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewK8sDashboardGetAdminTokenParamsWithTimeout creates a new K8sDashboardGetAdminTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewK8sDashboardGetAdminTokenParamsWithTimeout(timeout time.Duration) *K8sDashboardGetAdminTokenParams {
	var ()
	return &K8sDashboardGetAdminTokenParams{

		timeout: timeout,
	}
}

// NewK8sDashboardGetAdminTokenParamsWithContext creates a new K8sDashboardGetAdminTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewK8sDashboardGetAdminTokenParamsWithContext(ctx context.Context) *K8sDashboardGetAdminTokenParams {
	var ()
	return &K8sDashboardGetAdminTokenParams{

		Context: ctx,
	}
}

// NewK8sDashboardGetAdminTokenParamsWithHTTPClient creates a new K8sDashboardGetAdminTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewK8sDashboardGetAdminTokenParamsWithHTTPClient(client *http.Client) *K8sDashboardGetAdminTokenParams {
	var ()
	return &K8sDashboardGetAdminTokenParams{
		HTTPClient: client,
	}
}

/*K8sDashboardGetAdminTokenParams contains all the parameters to send to the API endpoint
for the k8s dashboard get admin token operation typically these are written to a http.Request
*/
type K8sDashboardGetAdminTokenParams struct {

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

// WithTimeout adds the timeout to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) WithTimeout(timeout time.Duration) *K8sDashboardGetAdminTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) WithContext(ctx context.Context) *K8sDashboardGetAdminTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) WithHTTPClient(client *http.Client) *K8sDashboardGetAdminTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) WithAuthorization(authorization string) *K8sDashboardGetAdminTokenParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSvcDomainID adds the svcDomainID to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) WithSvcDomainID(svcDomainID string) *K8sDashboardGetAdminTokenParams {
	o.SetSvcDomainID(svcDomainID)
	return o
}

// SetSvcDomainID adds the svcDomainId to the k8s dashboard get admin token params
func (o *K8sDashboardGetAdminTokenParams) SetSvcDomainID(svcDomainID string) {
	o.SvcDomainID = svcDomainID
}

// WriteToRequest writes these params to a swagger request
func (o *K8sDashboardGetAdminTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
