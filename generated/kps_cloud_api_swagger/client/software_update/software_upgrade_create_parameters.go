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

// NewSoftwareUpgradeCreateParams creates a new SoftwareUpgradeCreateParams object
// with the default values initialized.
func NewSoftwareUpgradeCreateParams() *SoftwareUpgradeCreateParams {
	var ()
	return &SoftwareUpgradeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareUpgradeCreateParamsWithTimeout creates a new SoftwareUpgradeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareUpgradeCreateParamsWithTimeout(timeout time.Duration) *SoftwareUpgradeCreateParams {
	var ()
	return &SoftwareUpgradeCreateParams{

		timeout: timeout,
	}
}

// NewSoftwareUpgradeCreateParamsWithContext creates a new SoftwareUpgradeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareUpgradeCreateParamsWithContext(ctx context.Context) *SoftwareUpgradeCreateParams {
	var ()
	return &SoftwareUpgradeCreateParams{

		Context: ctx,
	}
}

// NewSoftwareUpgradeCreateParamsWithHTTPClient creates a new SoftwareUpgradeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareUpgradeCreateParamsWithHTTPClient(client *http.Client) *SoftwareUpgradeCreateParams {
	var ()
	return &SoftwareUpgradeCreateParams{
		HTTPClient: client,
	}
}

/*SoftwareUpgradeCreateParams contains all the parameters to send to the API endpoint
for the software upgrade create operation typically these are written to a http.Request
*/
type SoftwareUpgradeCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.SoftwareUpgradeCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) WithTimeout(timeout time.Duration) *SoftwareUpgradeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) WithContext(ctx context.Context) *SoftwareUpgradeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) WithHTTPClient(client *http.Client) *SoftwareUpgradeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) WithAuthorization(authorization string) *SoftwareUpgradeCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) WithBody(body *models.SoftwareUpgradeCreate) *SoftwareUpgradeCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the software upgrade create params
func (o *SoftwareUpgradeCreateParams) SetBody(body *models.SoftwareUpgradeCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareUpgradeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
