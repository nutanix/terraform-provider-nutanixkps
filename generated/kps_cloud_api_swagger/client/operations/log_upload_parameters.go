// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewLogUploadParams creates a new LogUploadParams object
// with the default values initialized.
func NewLogUploadParams() *LogUploadParams {
	var ()
	return &LogUploadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogUploadParamsWithTimeout creates a new LogUploadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogUploadParamsWithTimeout(timeout time.Duration) *LogUploadParams {
	var ()
	return &LogUploadParams{

		timeout: timeout,
	}
}

// NewLogUploadParamsWithContext creates a new LogUploadParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogUploadParamsWithContext(ctx context.Context) *LogUploadParams {
	var ()
	return &LogUploadParams{

		Context: ctx,
	}
}

// NewLogUploadParamsWithHTTPClient creates a new LogUploadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogUploadParamsWithHTTPClient(client *http.Client) *LogUploadParams {
	var ()
	return &LogUploadParams{
		HTTPClient: client,
	}
}

/*LogUploadParams contains all the parameters to send to the API endpoint
for the log upload operation typically these are written to a http.Request
*/
type LogUploadParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Payload*/
	Payload *models.ObjectRequestLogUpload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the log upload params
func (o *LogUploadParams) WithTimeout(timeout time.Duration) *LogUploadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log upload params
func (o *LogUploadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log upload params
func (o *LogUploadParams) WithContext(ctx context.Context) *LogUploadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log upload params
func (o *LogUploadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log upload params
func (o *LogUploadParams) WithHTTPClient(client *http.Client) *LogUploadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log upload params
func (o *LogUploadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the log upload params
func (o *LogUploadParams) WithAuthorization(authorization string) *LogUploadParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the log upload params
func (o *LogUploadParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithPayload adds the payload to the log upload params
func (o *LogUploadParams) WithPayload(payload *models.ObjectRequestLogUpload) *LogUploadParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the log upload params
func (o *LogUploadParams) SetPayload(payload *models.ObjectRequestLogUpload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *LogUploadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
