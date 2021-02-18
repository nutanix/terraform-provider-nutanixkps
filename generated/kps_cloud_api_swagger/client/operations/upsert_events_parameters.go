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

// NewUpsertEventsParams creates a new UpsertEventsParams object
// with the default values initialized.
func NewUpsertEventsParams() *UpsertEventsParams {
	var ()
	return &UpsertEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertEventsParamsWithTimeout creates a new UpsertEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpsertEventsParamsWithTimeout(timeout time.Duration) *UpsertEventsParams {
	var ()
	return &UpsertEventsParams{

		timeout: timeout,
	}
}

// NewUpsertEventsParamsWithContext creates a new UpsertEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpsertEventsParamsWithContext(ctx context.Context) *UpsertEventsParams {
	var ()
	return &UpsertEventsParams{

		Context: ctx,
	}
}

// NewUpsertEventsParamsWithHTTPClient creates a new UpsertEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpsertEventsParamsWithHTTPClient(client *http.Client) *UpsertEventsParams {
	var ()
	return &UpsertEventsParams{
		HTTPClient: client,
	}
}

/*UpsertEventsParams contains all the parameters to send to the API endpoint
for the upsert events operation typically these are written to a http.Request
*/
type UpsertEventsParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  This is events upsert request description

	*/
	Body *models.EventUpsertRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upsert events params
func (o *UpsertEventsParams) WithTimeout(timeout time.Duration) *UpsertEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert events params
func (o *UpsertEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert events params
func (o *UpsertEventsParams) WithContext(ctx context.Context) *UpsertEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert events params
func (o *UpsertEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert events params
func (o *UpsertEventsParams) WithHTTPClient(client *http.Client) *UpsertEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert events params
func (o *UpsertEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the upsert events params
func (o *UpsertEventsParams) WithAuthorization(authorization string) *UpsertEventsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the upsert events params
func (o *UpsertEventsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the upsert events params
func (o *UpsertEventsParams) WithBody(body *models.EventUpsertRequest) *UpsertEventsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upsert events params
func (o *UpsertEventsParams) SetBody(body *models.EventUpsertRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
