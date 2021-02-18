// Code generated by go-swagger; DO NOT EDIT.

package log

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

// NewLogRequestDownloadV2Params creates a new LogRequestDownloadV2Params object
// with the default values initialized.
func NewLogRequestDownloadV2Params() *LogRequestDownloadV2Params {
	var ()
	return &LogRequestDownloadV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogRequestDownloadV2ParamsWithTimeout creates a new LogRequestDownloadV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogRequestDownloadV2ParamsWithTimeout(timeout time.Duration) *LogRequestDownloadV2Params {
	var ()
	return &LogRequestDownloadV2Params{

		timeout: timeout,
	}
}

// NewLogRequestDownloadV2ParamsWithContext creates a new LogRequestDownloadV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewLogRequestDownloadV2ParamsWithContext(ctx context.Context) *LogRequestDownloadV2Params {
	var ()
	return &LogRequestDownloadV2Params{

		Context: ctx,
	}
}

// NewLogRequestDownloadV2ParamsWithHTTPClient creates a new LogRequestDownloadV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogRequestDownloadV2ParamsWithHTTPClient(client *http.Client) *LogRequestDownloadV2Params {
	var ()
	return &LogRequestDownloadV2Params{
		HTTPClient: client,
	}
}

/*LogRequestDownloadV2Params contains all the parameters to send to the API endpoint
for the log request download v2 operation typically these are written to a http.Request
*/
type LogRequestDownloadV2Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Request*/
	Request *models.RequestLogDownloadPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the log request download v2 params
func (o *LogRequestDownloadV2Params) WithTimeout(timeout time.Duration) *LogRequestDownloadV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log request download v2 params
func (o *LogRequestDownloadV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log request download v2 params
func (o *LogRequestDownloadV2Params) WithContext(ctx context.Context) *LogRequestDownloadV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log request download v2 params
func (o *LogRequestDownloadV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log request download v2 params
func (o *LogRequestDownloadV2Params) WithHTTPClient(client *http.Client) *LogRequestDownloadV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log request download v2 params
func (o *LogRequestDownloadV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the log request download v2 params
func (o *LogRequestDownloadV2Params) WithAuthorization(authorization string) *LogRequestDownloadV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the log request download v2 params
func (o *LogRequestDownloadV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the log request download v2 params
func (o *LogRequestDownloadV2Params) WithRequest(request *models.RequestLogDownloadPayload) *LogRequestDownloadV2Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the log request download v2 params
func (o *LogRequestDownloadV2Params) SetRequest(request *models.RequestLogDownloadPayload) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *LogRequestDownloadV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
