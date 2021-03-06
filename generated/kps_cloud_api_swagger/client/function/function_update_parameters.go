// Code generated by go-swagger; DO NOT EDIT.

package function

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

// NewFunctionUpdateParams creates a new FunctionUpdateParams object
// with the default values initialized.
func NewFunctionUpdateParams() *FunctionUpdateParams {
	var ()
	return &FunctionUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionUpdateParamsWithTimeout creates a new FunctionUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFunctionUpdateParamsWithTimeout(timeout time.Duration) *FunctionUpdateParams {
	var ()
	return &FunctionUpdateParams{

		timeout: timeout,
	}
}

// NewFunctionUpdateParamsWithContext creates a new FunctionUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewFunctionUpdateParamsWithContext(ctx context.Context) *FunctionUpdateParams {
	var ()
	return &FunctionUpdateParams{

		Context: ctx,
	}
}

// NewFunctionUpdateParamsWithHTTPClient creates a new FunctionUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFunctionUpdateParamsWithHTTPClient(client *http.Client) *FunctionUpdateParams {
	var ()
	return &FunctionUpdateParams{
		HTTPClient: client,
	}
}

/*FunctionUpdateParams contains all the parameters to send to the API endpoint
for the function update operation typically these are written to a http.Request
*/
type FunctionUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.Function
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the function update params
func (o *FunctionUpdateParams) WithTimeout(timeout time.Duration) *FunctionUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the function update params
func (o *FunctionUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the function update params
func (o *FunctionUpdateParams) WithContext(ctx context.Context) *FunctionUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the function update params
func (o *FunctionUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the function update params
func (o *FunctionUpdateParams) WithHTTPClient(client *http.Client) *FunctionUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the function update params
func (o *FunctionUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the function update params
func (o *FunctionUpdateParams) WithAuthorization(authorization string) *FunctionUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the function update params
func (o *FunctionUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the function update params
func (o *FunctionUpdateParams) WithBody(body *models.Function) *FunctionUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the function update params
func (o *FunctionUpdateParams) SetBody(body *models.Function) {
	o.Body = body
}

// WithID adds the id to the function update params
func (o *FunctionUpdateParams) WithID(id string) *FunctionUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the function update params
func (o *FunctionUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
