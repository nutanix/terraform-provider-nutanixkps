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
)

// NewFunctionGetParams creates a new FunctionGetParams object
// with the default values initialized.
func NewFunctionGetParams() *FunctionGetParams {
	var ()
	return &FunctionGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionGetParamsWithTimeout creates a new FunctionGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFunctionGetParamsWithTimeout(timeout time.Duration) *FunctionGetParams {
	var ()
	return &FunctionGetParams{

		timeout: timeout,
	}
}

// NewFunctionGetParamsWithContext creates a new FunctionGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFunctionGetParamsWithContext(ctx context.Context) *FunctionGetParams {
	var ()
	return &FunctionGetParams{

		Context: ctx,
	}
}

// NewFunctionGetParamsWithHTTPClient creates a new FunctionGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFunctionGetParamsWithHTTPClient(client *http.Client) *FunctionGetParams {
	var ()
	return &FunctionGetParams{
		HTTPClient: client,
	}
}

/*FunctionGetParams contains all the parameters to send to the API endpoint
for the function get operation typically these are written to a http.Request
*/
type FunctionGetParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the function get params
func (o *FunctionGetParams) WithTimeout(timeout time.Duration) *FunctionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the function get params
func (o *FunctionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the function get params
func (o *FunctionGetParams) WithContext(ctx context.Context) *FunctionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the function get params
func (o *FunctionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the function get params
func (o *FunctionGetParams) WithHTTPClient(client *http.Client) *FunctionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the function get params
func (o *FunctionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the function get params
func (o *FunctionGetParams) WithAuthorization(authorization string) *FunctionGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the function get params
func (o *FunctionGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the function get params
func (o *FunctionGetParams) WithID(id string) *FunctionGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the function get params
func (o *FunctionGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
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
