// Code generated by go-swagger; DO NOT EDIT.

package software_update

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

// NewSoftwareDownloadCreateParams creates a new SoftwareDownloadCreateParams object
// with the default values initialized.
func NewSoftwareDownloadCreateParams() *SoftwareDownloadCreateParams {
	var ()
	return &SoftwareDownloadCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareDownloadCreateParamsWithTimeout creates a new SoftwareDownloadCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareDownloadCreateParamsWithTimeout(timeout time.Duration) *SoftwareDownloadCreateParams {
	var ()
	return &SoftwareDownloadCreateParams{

		timeout: timeout,
	}
}

// NewSoftwareDownloadCreateParamsWithContext creates a new SoftwareDownloadCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareDownloadCreateParamsWithContext(ctx context.Context) *SoftwareDownloadCreateParams {
	var ()
	return &SoftwareDownloadCreateParams{

		Context: ctx,
	}
}

// NewSoftwareDownloadCreateParamsWithHTTPClient creates a new SoftwareDownloadCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareDownloadCreateParamsWithHTTPClient(client *http.Client) *SoftwareDownloadCreateParams {
	var ()
	return &SoftwareDownloadCreateParams{
		HTTPClient: client,
	}
}

/*SoftwareDownloadCreateParams contains all the parameters to send to the API endpoint
for the software download create operation typically these are written to a http.Request
*/
type SoftwareDownloadCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.SoftwareDownloadCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the software download create params
func (o *SoftwareDownloadCreateParams) WithTimeout(timeout time.Duration) *SoftwareDownloadCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software download create params
func (o *SoftwareDownloadCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software download create params
func (o *SoftwareDownloadCreateParams) WithContext(ctx context.Context) *SoftwareDownloadCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software download create params
func (o *SoftwareDownloadCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software download create params
func (o *SoftwareDownloadCreateParams) WithHTTPClient(client *http.Client) *SoftwareDownloadCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software download create params
func (o *SoftwareDownloadCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software download create params
func (o *SoftwareDownloadCreateParams) WithAuthorization(authorization string) *SoftwareDownloadCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software download create params
func (o *SoftwareDownloadCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the software download create params
func (o *SoftwareDownloadCreateParams) WithBody(body *models.SoftwareDownloadCreate) *SoftwareDownloadCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the software download create params
func (o *SoftwareDownloadCreateParams) SetBody(body *models.SoftwareDownloadCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareDownloadCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
