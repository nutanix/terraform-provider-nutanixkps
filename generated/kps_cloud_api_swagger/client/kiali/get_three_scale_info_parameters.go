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

// NewGetThreeScaleInfoParams creates a new GetThreeScaleInfoParams object
// with the default values initialized.
func NewGetThreeScaleInfoParams() *GetThreeScaleInfoParams {
	var ()
	return &GetThreeScaleInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetThreeScaleInfoParamsWithTimeout creates a new GetThreeScaleInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetThreeScaleInfoParamsWithTimeout(timeout time.Duration) *GetThreeScaleInfoParams {
	var ()
	return &GetThreeScaleInfoParams{

		timeout: timeout,
	}
}

// NewGetThreeScaleInfoParamsWithContext creates a new GetThreeScaleInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetThreeScaleInfoParamsWithContext(ctx context.Context) *GetThreeScaleInfoParams {
	var ()
	return &GetThreeScaleInfoParams{

		Context: ctx,
	}
}

// NewGetThreeScaleInfoParamsWithHTTPClient creates a new GetThreeScaleInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetThreeScaleInfoParamsWithHTTPClient(client *http.Client) *GetThreeScaleInfoParams {
	var ()
	return &GetThreeScaleInfoParams{
		HTTPClient: client,
	}
}

/*GetThreeScaleInfoParams contains all the parameters to send to the API endpoint
for the get three scale info operation typically these are written to a http.Request
*/
type GetThreeScaleInfoParams struct {

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

// WithTimeout adds the timeout to the get three scale info params
func (o *GetThreeScaleInfoParams) WithTimeout(timeout time.Duration) *GetThreeScaleInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get three scale info params
func (o *GetThreeScaleInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get three scale info params
func (o *GetThreeScaleInfoParams) WithContext(ctx context.Context) *GetThreeScaleInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get three scale info params
func (o *GetThreeScaleInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get three scale info params
func (o *GetThreeScaleInfoParams) WithHTTPClient(client *http.Client) *GetThreeScaleInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get three scale info params
func (o *GetThreeScaleInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get three scale info params
func (o *GetThreeScaleInfoParams) WithAuthorization(authorization string) *GetThreeScaleInfoParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get three scale info params
func (o *GetThreeScaleInfoParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithServiceDomain adds the serviceDomain to the get three scale info params
func (o *GetThreeScaleInfoParams) WithServiceDomain(serviceDomain string) *GetThreeScaleInfoParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the get three scale info params
func (o *GetThreeScaleInfoParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *GetThreeScaleInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
