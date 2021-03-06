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

// NewK8sDashboardGetUserTokenParams creates a new K8sDashboardGetUserTokenParams object
// with the default values initialized.
func NewK8sDashboardGetUserTokenParams() *K8sDashboardGetUserTokenParams {
	var ()
	return &K8sDashboardGetUserTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewK8sDashboardGetUserTokenParamsWithTimeout creates a new K8sDashboardGetUserTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewK8sDashboardGetUserTokenParamsWithTimeout(timeout time.Duration) *K8sDashboardGetUserTokenParams {
	var ()
	return &K8sDashboardGetUserTokenParams{

		timeout: timeout,
	}
}

// NewK8sDashboardGetUserTokenParamsWithContext creates a new K8sDashboardGetUserTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewK8sDashboardGetUserTokenParamsWithContext(ctx context.Context) *K8sDashboardGetUserTokenParams {
	var ()
	return &K8sDashboardGetUserTokenParams{

		Context: ctx,
	}
}

// NewK8sDashboardGetUserTokenParamsWithHTTPClient creates a new K8sDashboardGetUserTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewK8sDashboardGetUserTokenParamsWithHTTPClient(client *http.Client) *K8sDashboardGetUserTokenParams {
	var ()
	return &K8sDashboardGetUserTokenParams{
		HTTPClient: client,
	}
}

/*K8sDashboardGetUserTokenParams contains all the parameters to send to the API endpoint
for the k8s dashboard get user token operation typically these are written to a http.Request
*/
type K8sDashboardGetUserTokenParams struct {

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

// WithTimeout adds the timeout to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) WithTimeout(timeout time.Duration) *K8sDashboardGetUserTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) WithContext(ctx context.Context) *K8sDashboardGetUserTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) WithHTTPClient(client *http.Client) *K8sDashboardGetUserTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) WithAuthorization(authorization string) *K8sDashboardGetUserTokenParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSvcDomainID adds the svcDomainID to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) WithSvcDomainID(svcDomainID string) *K8sDashboardGetUserTokenParams {
	o.SetSvcDomainID(svcDomainID)
	return o
}

// SetSvcDomainID adds the svcDomainId to the k8s dashboard get user token params
func (o *K8sDashboardGetUserTokenParams) SetSvcDomainID(svcDomainID string) {
	o.SvcDomainID = svcDomainID
}

// WriteToRequest writes these params to a swagger request
func (o *K8sDashboardGetUserTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
