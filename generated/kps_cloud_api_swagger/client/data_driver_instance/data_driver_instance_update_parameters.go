// Code generated by go-swagger; DO NOT EDIT.

package data_driver_instance

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

// NewDataDriverInstanceUpdateParams creates a new DataDriverInstanceUpdateParams object
// with the default values initialized.
func NewDataDriverInstanceUpdateParams() *DataDriverInstanceUpdateParams {
	var ()
	return &DataDriverInstanceUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataDriverInstanceUpdateParamsWithTimeout creates a new DataDriverInstanceUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataDriverInstanceUpdateParamsWithTimeout(timeout time.Duration) *DataDriverInstanceUpdateParams {
	var ()
	return &DataDriverInstanceUpdateParams{

		timeout: timeout,
	}
}

// NewDataDriverInstanceUpdateParamsWithContext creates a new DataDriverInstanceUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataDriverInstanceUpdateParamsWithContext(ctx context.Context) *DataDriverInstanceUpdateParams {
	var ()
	return &DataDriverInstanceUpdateParams{

		Context: ctx,
	}
}

// NewDataDriverInstanceUpdateParamsWithHTTPClient creates a new DataDriverInstanceUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataDriverInstanceUpdateParamsWithHTTPClient(client *http.Client) *DataDriverInstanceUpdateParams {
	var ()
	return &DataDriverInstanceUpdateParams{
		HTTPClient: client,
	}
}

/*DataDriverInstanceUpdateParams contains all the parameters to send to the API endpoint
for the data driver instance update operation typically these are written to a http.Request
*/
type DataDriverInstanceUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  Parameters and values used when creating or updating a data driver instance

	*/
	Body *models.DataDriverInstance
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) WithTimeout(timeout time.Duration) *DataDriverInstanceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) WithContext(ctx context.Context) *DataDriverInstanceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) WithHTTPClient(client *http.Client) *DataDriverInstanceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) WithAuthorization(authorization string) *DataDriverInstanceUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) WithBody(body *models.DataDriverInstance) *DataDriverInstanceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) SetBody(body *models.DataDriverInstance) {
	o.Body = body
}

// WithID adds the id to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) WithID(id string) *DataDriverInstanceUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the data driver instance update params
func (o *DataDriverInstanceUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DataDriverInstanceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
