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

// NewDataStreamUpdateParams creates a new DataStreamUpdateParams object
// with the default values initialized.
func NewDataStreamUpdateParams() *DataStreamUpdateParams {
	var ()
	return &DataStreamUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataStreamUpdateParamsWithTimeout creates a new DataStreamUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataStreamUpdateParamsWithTimeout(timeout time.Duration) *DataStreamUpdateParams {
	var ()
	return &DataStreamUpdateParams{

		timeout: timeout,
	}
}

// NewDataStreamUpdateParamsWithContext creates a new DataStreamUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataStreamUpdateParamsWithContext(ctx context.Context) *DataStreamUpdateParams {
	var ()
	return &DataStreamUpdateParams{

		Context: ctx,
	}
}

// NewDataStreamUpdateParamsWithHTTPClient creates a new DataStreamUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataStreamUpdateParamsWithHTTPClient(client *http.Client) *DataStreamUpdateParams {
	var ()
	return &DataStreamUpdateParams{
		HTTPClient: client,
	}
}

/*DataStreamUpdateParams contains all the parameters to send to the API endpoint
for the data stream update operation typically these are written to a http.Request
*/
type DataStreamUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.DataStream

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data stream update params
func (o *DataStreamUpdateParams) WithTimeout(timeout time.Duration) *DataStreamUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data stream update params
func (o *DataStreamUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data stream update params
func (o *DataStreamUpdateParams) WithContext(ctx context.Context) *DataStreamUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data stream update params
func (o *DataStreamUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data stream update params
func (o *DataStreamUpdateParams) WithHTTPClient(client *http.Client) *DataStreamUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data stream update params
func (o *DataStreamUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data stream update params
func (o *DataStreamUpdateParams) WithAuthorization(authorization string) *DataStreamUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data stream update params
func (o *DataStreamUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data stream update params
func (o *DataStreamUpdateParams) WithBody(body *models.DataStream) *DataStreamUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data stream update params
func (o *DataStreamUpdateParams) SetBody(body *models.DataStream) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DataStreamUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
