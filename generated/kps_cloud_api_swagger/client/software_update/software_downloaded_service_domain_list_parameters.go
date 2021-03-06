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
)

// NewSoftwareDownloadedServiceDomainListParams creates a new SoftwareDownloadedServiceDomainListParams object
// with the default values initialized.
func NewSoftwareDownloadedServiceDomainListParams() *SoftwareDownloadedServiceDomainListParams {
	var ()
	return &SoftwareDownloadedServiceDomainListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareDownloadedServiceDomainListParamsWithTimeout creates a new SoftwareDownloadedServiceDomainListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareDownloadedServiceDomainListParamsWithTimeout(timeout time.Duration) *SoftwareDownloadedServiceDomainListParams {
	var ()
	return &SoftwareDownloadedServiceDomainListParams{

		timeout: timeout,
	}
}

// NewSoftwareDownloadedServiceDomainListParamsWithContext creates a new SoftwareDownloadedServiceDomainListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareDownloadedServiceDomainListParamsWithContext(ctx context.Context) *SoftwareDownloadedServiceDomainListParams {
	var ()
	return &SoftwareDownloadedServiceDomainListParams{

		Context: ctx,
	}
}

// NewSoftwareDownloadedServiceDomainListParamsWithHTTPClient creates a new SoftwareDownloadedServiceDomainListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareDownloadedServiceDomainListParamsWithHTTPClient(client *http.Client) *SoftwareDownloadedServiceDomainListParams {
	var ()
	return &SoftwareDownloadedServiceDomainListParams{
		HTTPClient: client,
	}
}

/*SoftwareDownloadedServiceDomainListParams contains all the parameters to send to the API endpoint
for the software downloaded service domain list operation typically these are written to a http.Request
*/
type SoftwareDownloadedServiceDomainListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Release
	  release

	*/
	Release string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) WithTimeout(timeout time.Duration) *SoftwareDownloadedServiceDomainListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) WithContext(ctx context.Context) *SoftwareDownloadedServiceDomainListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) WithHTTPClient(client *http.Client) *SoftwareDownloadedServiceDomainListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) WithAuthorization(authorization string) *SoftwareDownloadedServiceDomainListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRelease adds the release to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) WithRelease(release string) *SoftwareDownloadedServiceDomainListParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the software downloaded service domain list params
func (o *SoftwareDownloadedServiceDomainListParams) SetRelease(release string) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareDownloadedServiceDomainListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param release
	if err := r.SetPathParam("release", o.Release); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
