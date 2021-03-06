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

// NewPodLogsParams creates a new PodLogsParams object
// with the default values initialized.
func NewPodLogsParams() *PodLogsParams {
	var ()
	return &PodLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPodLogsParamsWithTimeout creates a new PodLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPodLogsParamsWithTimeout(timeout time.Duration) *PodLogsParams {
	var ()
	return &PodLogsParams{

		timeout: timeout,
	}
}

// NewPodLogsParamsWithContext creates a new PodLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPodLogsParamsWithContext(ctx context.Context) *PodLogsParams {
	var ()
	return &PodLogsParams{

		Context: ctx,
	}
}

// NewPodLogsParamsWithHTTPClient creates a new PodLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPodLogsParamsWithHTTPClient(client *http.Client) *PodLogsParams {
	var ()
	return &PodLogsParams{
		HTTPClient: client,
	}
}

/*PodLogsParams contains all the parameters to send to the API endpoint
for the pod logs operation typically these are written to a http.Request
*/
type PodLogsParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Container
	  The pod container name. Optional for single-container pod. Otherwise required.

	*/
	Container *string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*Pod
	  The pod name.

	*/
	Pod string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string
	/*SinceTime
	  The start time for fetching logs. UNIX time in seconds. Default is all logs.

	*/
	SinceTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pod logs params
func (o *PodLogsParams) WithTimeout(timeout time.Duration) *PodLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod logs params
func (o *PodLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod logs params
func (o *PodLogsParams) WithContext(ctx context.Context) *PodLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod logs params
func (o *PodLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod logs params
func (o *PodLogsParams) WithHTTPClient(client *http.Client) *PodLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod logs params
func (o *PodLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the pod logs params
func (o *PodLogsParams) WithAuthorization(authorization string) *PodLogsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the pod logs params
func (o *PodLogsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContainer adds the container to the pod logs params
func (o *PodLogsParams) WithContainer(container *string) *PodLogsParams {
	o.SetContainer(container)
	return o
}

// SetContainer adds the container to the pod logs params
func (o *PodLogsParams) SetContainer(container *string) {
	o.Container = container
}

// WithNamespace adds the namespace to the pod logs params
func (o *PodLogsParams) WithNamespace(namespace string) *PodLogsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the pod logs params
func (o *PodLogsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPod adds the pod to the pod logs params
func (o *PodLogsParams) WithPod(pod string) *PodLogsParams {
	o.SetPod(pod)
	return o
}

// SetPod adds the pod to the pod logs params
func (o *PodLogsParams) SetPod(pod string) {
	o.Pod = pod
}

// WithServiceDomain adds the serviceDomain to the pod logs params
func (o *PodLogsParams) WithServiceDomain(serviceDomain string) *PodLogsParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the pod logs params
func (o *PodLogsParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WithSinceTime adds the sinceTime to the pod logs params
func (o *PodLogsParams) WithSinceTime(sinceTime *string) *PodLogsParams {
	o.SetSinceTime(sinceTime)
	return o
}

// SetSinceTime adds the sinceTime to the pod logs params
func (o *PodLogsParams) SetSinceTime(sinceTime *string) {
	o.SinceTime = sinceTime
}

// WriteToRequest writes these params to a swagger request
func (o *PodLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Container != nil {

		// query param container
		var qrContainer string
		if o.Container != nil {
			qrContainer = *o.Container
		}
		qContainer := qrContainer
		if qContainer != "" {
			if err := r.SetQueryParam("container", qContainer); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param pod
	if err := r.SetPathParam("pod", o.Pod); err != nil {
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

	if o.SinceTime != nil {

		// query param sinceTime
		var qrSinceTime string
		if o.SinceTime != nil {
			qrSinceTime = *o.SinceTime
		}
		qSinceTime := qrSinceTime
		if qSinceTime != "" {
			if err := r.SetQueryParam("sinceTime", qSinceTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
