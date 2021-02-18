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

// NewGetThreeScaleHandlersParams creates a new GetThreeScaleHandlersParams object
// with the default values initialized.
func NewGetThreeScaleHandlersParams() *GetThreeScaleHandlersParams {
	var ()
	return &GetThreeScaleHandlersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetThreeScaleHandlersParamsWithTimeout creates a new GetThreeScaleHandlersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetThreeScaleHandlersParamsWithTimeout(timeout time.Duration) *GetThreeScaleHandlersParams {
	var ()
	return &GetThreeScaleHandlersParams{

		timeout: timeout,
	}
}

// NewGetThreeScaleHandlersParamsWithContext creates a new GetThreeScaleHandlersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetThreeScaleHandlersParamsWithContext(ctx context.Context) *GetThreeScaleHandlersParams {
	var ()
	return &GetThreeScaleHandlersParams{

		Context: ctx,
	}
}

// NewGetThreeScaleHandlersParamsWithHTTPClient creates a new GetThreeScaleHandlersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetThreeScaleHandlersParamsWithHTTPClient(client *http.Client) *GetThreeScaleHandlersParams {
	var ()
	return &GetThreeScaleHandlersParams{
		HTTPClient: client,
	}
}

/*GetThreeScaleHandlersParams contains all the parameters to send to the API endpoint
for the get three scale handlers operation typically these are written to a http.Request
*/
type GetThreeScaleHandlersParams struct {

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

// WithTimeout adds the timeout to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) WithTimeout(timeout time.Duration) *GetThreeScaleHandlersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) WithContext(ctx context.Context) *GetThreeScaleHandlersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) WithHTTPClient(client *http.Client) *GetThreeScaleHandlersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) WithAuthorization(authorization string) *GetThreeScaleHandlersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithServiceDomain adds the serviceDomain to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) WithServiceDomain(serviceDomain string) *GetThreeScaleHandlersParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the get three scale handlers params
func (o *GetThreeScaleHandlersParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *GetThreeScaleHandlersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
