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
)

// NewScriptRuntimeGetParams creates a new ScriptRuntimeGetParams object
// with the default values initialized.
func NewScriptRuntimeGetParams() *ScriptRuntimeGetParams {
	var ()
	return &ScriptRuntimeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScriptRuntimeGetParamsWithTimeout creates a new ScriptRuntimeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScriptRuntimeGetParamsWithTimeout(timeout time.Duration) *ScriptRuntimeGetParams {
	var ()
	return &ScriptRuntimeGetParams{

		timeout: timeout,
	}
}

// NewScriptRuntimeGetParamsWithContext creates a new ScriptRuntimeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewScriptRuntimeGetParamsWithContext(ctx context.Context) *ScriptRuntimeGetParams {
	var ()
	return &ScriptRuntimeGetParams{

		Context: ctx,
	}
}

// NewScriptRuntimeGetParamsWithHTTPClient creates a new ScriptRuntimeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScriptRuntimeGetParamsWithHTTPClient(client *http.Client) *ScriptRuntimeGetParams {
	var ()
	return &ScriptRuntimeGetParams{
		HTTPClient: client,
	}
}

/*ScriptRuntimeGetParams contains all the parameters to send to the API endpoint
for the script runtime get operation typically these are written to a http.Request
*/
type ScriptRuntimeGetParams struct {

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

// WithTimeout adds the timeout to the script runtime get params
func (o *ScriptRuntimeGetParams) WithTimeout(timeout time.Duration) *ScriptRuntimeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the script runtime get params
func (o *ScriptRuntimeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the script runtime get params
func (o *ScriptRuntimeGetParams) WithContext(ctx context.Context) *ScriptRuntimeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the script runtime get params
func (o *ScriptRuntimeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the script runtime get params
func (o *ScriptRuntimeGetParams) WithHTTPClient(client *http.Client) *ScriptRuntimeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the script runtime get params
func (o *ScriptRuntimeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the script runtime get params
func (o *ScriptRuntimeGetParams) WithAuthorization(authorization string) *ScriptRuntimeGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the script runtime get params
func (o *ScriptRuntimeGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the script runtime get params
func (o *ScriptRuntimeGetParams) WithID(id string) *ScriptRuntimeGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the script runtime get params
func (o *ScriptRuntimeGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScriptRuntimeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
