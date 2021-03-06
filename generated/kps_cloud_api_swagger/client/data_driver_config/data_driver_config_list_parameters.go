// Code generated by go-swagger; DO NOT EDIT.

package data_driver_config

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

// NewDataDriverConfigListParams creates a new DataDriverConfigListParams object
// with the default values initialized.
func NewDataDriverConfigListParams() *DataDriverConfigListParams {
	var ()
	return &DataDriverConfigListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataDriverConfigListParamsWithTimeout creates a new DataDriverConfigListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataDriverConfigListParamsWithTimeout(timeout time.Duration) *DataDriverConfigListParams {
	var ()
	return &DataDriverConfigListParams{

		timeout: timeout,
	}
}

// NewDataDriverConfigListParamsWithContext creates a new DataDriverConfigListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataDriverConfigListParamsWithContext(ctx context.Context) *DataDriverConfigListParams {
	var ()
	return &DataDriverConfigListParams{

		Context: ctx,
	}
}

// NewDataDriverConfigListParamsWithHTTPClient creates a new DataDriverConfigListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataDriverConfigListParamsWithHTTPClient(client *http.Client) *DataDriverConfigListParams {
	var ()
	return &DataDriverConfigListParams{
		HTTPClient: client,
	}
}

/*DataDriverConfigListParams contains all the parameters to send to the API endpoint
for the data driver config list operation typically these are written to a http.Request
*/
type DataDriverConfigListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data driver config list params
func (o *DataDriverConfigListParams) WithTimeout(timeout time.Duration) *DataDriverConfigListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data driver config list params
func (o *DataDriverConfigListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data driver config list params
func (o *DataDriverConfigListParams) WithContext(ctx context.Context) *DataDriverConfigListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data driver config list params
func (o *DataDriverConfigListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data driver config list params
func (o *DataDriverConfigListParams) WithHTTPClient(client *http.Client) *DataDriverConfigListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data driver config list params
func (o *DataDriverConfigListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data driver config list params
func (o *DataDriverConfigListParams) WithAuthorization(authorization string) *DataDriverConfigListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data driver config list params
func (o *DataDriverConfigListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *DataDriverConfigListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
