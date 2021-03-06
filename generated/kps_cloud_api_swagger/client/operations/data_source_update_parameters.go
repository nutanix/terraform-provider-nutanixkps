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

// NewDataSourceUpdateParams creates a new DataSourceUpdateParams object
// with the default values initialized.
func NewDataSourceUpdateParams() *DataSourceUpdateParams {
	var ()
	return &DataSourceUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataSourceUpdateParamsWithTimeout creates a new DataSourceUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataSourceUpdateParamsWithTimeout(timeout time.Duration) *DataSourceUpdateParams {
	var ()
	return &DataSourceUpdateParams{

		timeout: timeout,
	}
}

// NewDataSourceUpdateParamsWithContext creates a new DataSourceUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataSourceUpdateParamsWithContext(ctx context.Context) *DataSourceUpdateParams {
	var ()
	return &DataSourceUpdateParams{

		Context: ctx,
	}
}

// NewDataSourceUpdateParamsWithHTTPClient creates a new DataSourceUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataSourceUpdateParamsWithHTTPClient(client *http.Client) *DataSourceUpdateParams {
	var ()
	return &DataSourceUpdateParams{
		HTTPClient: client,
	}
}

/*DataSourceUpdateParams contains all the parameters to send to the API endpoint
for the data source update operation typically these are written to a http.Request
*/
type DataSourceUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.DataSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data source update params
func (o *DataSourceUpdateParams) WithTimeout(timeout time.Duration) *DataSourceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data source update params
func (o *DataSourceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data source update params
func (o *DataSourceUpdateParams) WithContext(ctx context.Context) *DataSourceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data source update params
func (o *DataSourceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data source update params
func (o *DataSourceUpdateParams) WithHTTPClient(client *http.Client) *DataSourceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data source update params
func (o *DataSourceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data source update params
func (o *DataSourceUpdateParams) WithAuthorization(authorization string) *DataSourceUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data source update params
func (o *DataSourceUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data source update params
func (o *DataSourceUpdateParams) WithBody(body *models.DataSource) *DataSourceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data source update params
func (o *DataSourceUpdateParams) SetBody(body *models.DataSource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DataSourceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
