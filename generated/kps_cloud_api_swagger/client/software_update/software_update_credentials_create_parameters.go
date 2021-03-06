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

// NewSoftwareUpdateCredentialsCreateParams creates a new SoftwareUpdateCredentialsCreateParams object
// with the default values initialized.
func NewSoftwareUpdateCredentialsCreateParams() *SoftwareUpdateCredentialsCreateParams {
	var ()
	return &SoftwareUpdateCredentialsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareUpdateCredentialsCreateParamsWithTimeout creates a new SoftwareUpdateCredentialsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareUpdateCredentialsCreateParamsWithTimeout(timeout time.Duration) *SoftwareUpdateCredentialsCreateParams {
	var ()
	return &SoftwareUpdateCredentialsCreateParams{

		timeout: timeout,
	}
}

// NewSoftwareUpdateCredentialsCreateParamsWithContext creates a new SoftwareUpdateCredentialsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareUpdateCredentialsCreateParamsWithContext(ctx context.Context) *SoftwareUpdateCredentialsCreateParams {
	var ()
	return &SoftwareUpdateCredentialsCreateParams{

		Context: ctx,
	}
}

// NewSoftwareUpdateCredentialsCreateParamsWithHTTPClient creates a new SoftwareUpdateCredentialsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareUpdateCredentialsCreateParamsWithHTTPClient(client *http.Client) *SoftwareUpdateCredentialsCreateParams {
	var ()
	return &SoftwareUpdateCredentialsCreateParams{
		HTTPClient: client,
	}
}

/*SoftwareUpdateCredentialsCreateParams contains all the parameters to send to the API endpoint
for the software update credentials create operation typically these are written to a http.Request
*/
type SoftwareUpdateCredentialsCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.SoftwareUpdateCredentials

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) WithTimeout(timeout time.Duration) *SoftwareUpdateCredentialsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) WithContext(ctx context.Context) *SoftwareUpdateCredentialsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) WithHTTPClient(client *http.Client) *SoftwareUpdateCredentialsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) WithAuthorization(authorization string) *SoftwareUpdateCredentialsCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) WithBody(body *models.SoftwareUpdateCredentials) *SoftwareUpdateCredentialsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the software update credentials create params
func (o *SoftwareUpdateCredentialsCreateParams) SetBody(body *models.SoftwareUpdateCredentials) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareUpdateCredentialsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
