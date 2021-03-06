// Code generated by go-swagger; DO NOT EDIT.

package service_binding

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

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// NewServiceBindingCreateParams creates a new ServiceBindingCreateParams object
// with the default values initialized.
func NewServiceBindingCreateParams() *ServiceBindingCreateParams {
	var ()
	return &ServiceBindingCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBindingCreateParamsWithTimeout creates a new ServiceBindingCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBindingCreateParamsWithTimeout(timeout time.Duration) *ServiceBindingCreateParams {
	var ()
	return &ServiceBindingCreateParams{

		timeout: timeout,
	}
}

// NewServiceBindingCreateParamsWithContext creates a new ServiceBindingCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBindingCreateParamsWithContext(ctx context.Context) *ServiceBindingCreateParams {
	var ()
	return &ServiceBindingCreateParams{

		Context: ctx,
	}
}

// NewServiceBindingCreateParamsWithHTTPClient creates a new ServiceBindingCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBindingCreateParamsWithHTTPClient(client *http.Client) *ServiceBindingCreateParams {
	var ()
	return &ServiceBindingCreateParams{
		HTTPClient: client,
	}
}

/*ServiceBindingCreateParams contains all the parameters to send to the API endpoint
for the service binding create operation typically these are written to a http.Request
*/
type ServiceBindingCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.ServiceBindingParam

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service binding create params
func (o *ServiceBindingCreateParams) WithTimeout(timeout time.Duration) *ServiceBindingCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service binding create params
func (o *ServiceBindingCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service binding create params
func (o *ServiceBindingCreateParams) WithContext(ctx context.Context) *ServiceBindingCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service binding create params
func (o *ServiceBindingCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service binding create params
func (o *ServiceBindingCreateParams) WithHTTPClient(client *http.Client) *ServiceBindingCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service binding create params
func (o *ServiceBindingCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the service binding create params
func (o *ServiceBindingCreateParams) WithAuthorization(authorization string) *ServiceBindingCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the service binding create params
func (o *ServiceBindingCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the service binding create params
func (o *ServiceBindingCreateParams) WithBody(body *models.ServiceBindingParam) *ServiceBindingCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service binding create params
func (o *ServiceBindingCreateParams) SetBody(body *models.ServiceBindingParam) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBindingCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
