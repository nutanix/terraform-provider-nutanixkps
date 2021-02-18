// Code generated by go-swagger; DO NOT EDIT.

package data_driver_instance

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

// NewDataDriverInstanceCreateParams creates a new DataDriverInstanceCreateParams object
// with the default values initialized.
func NewDataDriverInstanceCreateParams() *DataDriverInstanceCreateParams {
	var ()
	return &DataDriverInstanceCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataDriverInstanceCreateParamsWithTimeout creates a new DataDriverInstanceCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataDriverInstanceCreateParamsWithTimeout(timeout time.Duration) *DataDriverInstanceCreateParams {
	var ()
	return &DataDriverInstanceCreateParams{

		timeout: timeout,
	}
}

// NewDataDriverInstanceCreateParamsWithContext creates a new DataDriverInstanceCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataDriverInstanceCreateParamsWithContext(ctx context.Context) *DataDriverInstanceCreateParams {
	var ()
	return &DataDriverInstanceCreateParams{

		Context: ctx,
	}
}

// NewDataDriverInstanceCreateParamsWithHTTPClient creates a new DataDriverInstanceCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataDriverInstanceCreateParamsWithHTTPClient(client *http.Client) *DataDriverInstanceCreateParams {
	var ()
	return &DataDriverInstanceCreateParams{
		HTTPClient: client,
	}
}

/*DataDriverInstanceCreateParams contains all the parameters to send to the API endpoint
for the data driver instance create operation typically these are written to a http.Request
*/
type DataDriverInstanceCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  Parameters and values used when creating or updating a data driver instance

	*/
	Body *models.DataDriverInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data driver instance create params
func (o *DataDriverInstanceCreateParams) WithTimeout(timeout time.Duration) *DataDriverInstanceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data driver instance create params
func (o *DataDriverInstanceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data driver instance create params
func (o *DataDriverInstanceCreateParams) WithContext(ctx context.Context) *DataDriverInstanceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data driver instance create params
func (o *DataDriverInstanceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data driver instance create params
func (o *DataDriverInstanceCreateParams) WithHTTPClient(client *http.Client) *DataDriverInstanceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data driver instance create params
func (o *DataDriverInstanceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data driver instance create params
func (o *DataDriverInstanceCreateParams) WithAuthorization(authorization string) *DataDriverInstanceCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data driver instance create params
func (o *DataDriverInstanceCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data driver instance create params
func (o *DataDriverInstanceCreateParams) WithBody(body *models.DataDriverInstance) *DataDriverInstanceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data driver instance create params
func (o *DataDriverInstanceCreateParams) SetBody(body *models.DataDriverInstance) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DataDriverInstanceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
