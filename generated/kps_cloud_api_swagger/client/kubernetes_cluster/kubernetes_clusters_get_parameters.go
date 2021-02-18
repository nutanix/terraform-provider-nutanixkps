// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_cluster

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

// NewKubernetesClustersGetParams creates a new KubernetesClustersGetParams object
// with the default values initialized.
func NewKubernetesClustersGetParams() *KubernetesClustersGetParams {
	var ()
	return &KubernetesClustersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesClustersGetParamsWithTimeout creates a new KubernetesClustersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKubernetesClustersGetParamsWithTimeout(timeout time.Duration) *KubernetesClustersGetParams {
	var ()
	return &KubernetesClustersGetParams{

		timeout: timeout,
	}
}

// NewKubernetesClustersGetParamsWithContext creates a new KubernetesClustersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewKubernetesClustersGetParamsWithContext(ctx context.Context) *KubernetesClustersGetParams {
	var ()
	return &KubernetesClustersGetParams{

		Context: ctx,
	}
}

// NewKubernetesClustersGetParamsWithHTTPClient creates a new KubernetesClustersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKubernetesClustersGetParamsWithHTTPClient(client *http.Client) *KubernetesClustersGetParams {
	var ()
	return &KubernetesClustersGetParams{
		HTTPClient: client,
	}
}

/*KubernetesClustersGetParams contains all the parameters to send to the API endpoint
for the kubernetes clusters get operation typically these are written to a http.Request
*/
type KubernetesClustersGetParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) WithTimeout(timeout time.Duration) *KubernetesClustersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) WithContext(ctx context.Context) *KubernetesClustersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) WithHTTPClient(client *http.Client) *KubernetesClustersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) WithAuthorization(authorization string) *KubernetesClustersGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) WithID(id string) *KubernetesClustersGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the kubernetes clusters get params
func (o *KubernetesClustersGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesClustersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
