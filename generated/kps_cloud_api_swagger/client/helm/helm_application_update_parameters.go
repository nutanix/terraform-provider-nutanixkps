// Code generated by go-swagger; DO NOT EDIT.

package helm

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

// NewHelmApplicationUpdateParams creates a new HelmApplicationUpdateParams object
// with the default values initialized.
func NewHelmApplicationUpdateParams() *HelmApplicationUpdateParams {
	var ()
	return &HelmApplicationUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHelmApplicationUpdateParamsWithTimeout creates a new HelmApplicationUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHelmApplicationUpdateParamsWithTimeout(timeout time.Duration) *HelmApplicationUpdateParams {
	var ()
	return &HelmApplicationUpdateParams{

		timeout: timeout,
	}
}

// NewHelmApplicationUpdateParamsWithContext creates a new HelmApplicationUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewHelmApplicationUpdateParamsWithContext(ctx context.Context) *HelmApplicationUpdateParams {
	var ()
	return &HelmApplicationUpdateParams{

		Context: ctx,
	}
}

// NewHelmApplicationUpdateParamsWithHTTPClient creates a new HelmApplicationUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHelmApplicationUpdateParamsWithHTTPClient(client *http.Client) *HelmApplicationUpdateParams {
	var ()
	return &HelmApplicationUpdateParams{
		HTTPClient: client,
	}
}

/*HelmApplicationUpdateParams contains all the parameters to send to the API endpoint
for the helm application update operation typically these are written to a http.Request
*/
type HelmApplicationUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Application*/
	Application string
	/*Chart*/
	Chart runtime.NamedReadCloser
	/*ID
	  ID of the entity

	*/
	ID string
	/*Values*/
	Values runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the helm application update params
func (o *HelmApplicationUpdateParams) WithTimeout(timeout time.Duration) *HelmApplicationUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the helm application update params
func (o *HelmApplicationUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the helm application update params
func (o *HelmApplicationUpdateParams) WithContext(ctx context.Context) *HelmApplicationUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the helm application update params
func (o *HelmApplicationUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the helm application update params
func (o *HelmApplicationUpdateParams) WithHTTPClient(client *http.Client) *HelmApplicationUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the helm application update params
func (o *HelmApplicationUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the helm application update params
func (o *HelmApplicationUpdateParams) WithAuthorization(authorization string) *HelmApplicationUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the helm application update params
func (o *HelmApplicationUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithApplication adds the application to the helm application update params
func (o *HelmApplicationUpdateParams) WithApplication(application string) *HelmApplicationUpdateParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the helm application update params
func (o *HelmApplicationUpdateParams) SetApplication(application string) {
	o.Application = application
}

// WithChart adds the chart to the helm application update params
func (o *HelmApplicationUpdateParams) WithChart(chart runtime.NamedReadCloser) *HelmApplicationUpdateParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the helm application update params
func (o *HelmApplicationUpdateParams) SetChart(chart runtime.NamedReadCloser) {
	o.Chart = chart
}

// WithID adds the id to the helm application update params
func (o *HelmApplicationUpdateParams) WithID(id string) *HelmApplicationUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the helm application update params
func (o *HelmApplicationUpdateParams) SetID(id string) {
	o.ID = id
}

// WithValues adds the values to the helm application update params
func (o *HelmApplicationUpdateParams) WithValues(values runtime.NamedReadCloser) *HelmApplicationUpdateParams {
	o.SetValues(values)
	return o
}

// SetValues adds the values to the helm application update params
func (o *HelmApplicationUpdateParams) SetValues(values runtime.NamedReadCloser) {
	o.Values = values
}

// WriteToRequest writes these params to a swagger request
func (o *HelmApplicationUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// form param application
	frApplication := o.Application
	fApplication := frApplication
	if fApplication != "" {
		if err := r.SetFormParam("application", fApplication); err != nil {
			return err
		}
	}

	if o.Chart != nil {

		if o.Chart != nil {

			// form file param chart
			if err := r.SetFileParam("chart", o.Chart); err != nil {
				return err
			}

		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Values != nil {

		if o.Values != nil {

			// form file param values
			if err := r.SetFileParam("values", o.Values); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
