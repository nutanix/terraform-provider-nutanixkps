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

// NewScriptDeleteParams creates a new ScriptDeleteParams object
// with the default values initialized.
func NewScriptDeleteParams() *ScriptDeleteParams {
	var ()
	return &ScriptDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScriptDeleteParamsWithTimeout creates a new ScriptDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScriptDeleteParamsWithTimeout(timeout time.Duration) *ScriptDeleteParams {
	var ()
	return &ScriptDeleteParams{

		timeout: timeout,
	}
}

// NewScriptDeleteParamsWithContext creates a new ScriptDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewScriptDeleteParamsWithContext(ctx context.Context) *ScriptDeleteParams {
	var ()
	return &ScriptDeleteParams{

		Context: ctx,
	}
}

// NewScriptDeleteParamsWithHTTPClient creates a new ScriptDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScriptDeleteParamsWithHTTPClient(client *http.Client) *ScriptDeleteParams {
	var ()
	return &ScriptDeleteParams{
		HTTPClient: client,
	}
}

/*ScriptDeleteParams contains all the parameters to send to the API endpoint
for the script delete operation typically these are written to a http.Request
*/
type ScriptDeleteParams struct {

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

// WithTimeout adds the timeout to the script delete params
func (o *ScriptDeleteParams) WithTimeout(timeout time.Duration) *ScriptDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the script delete params
func (o *ScriptDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the script delete params
func (o *ScriptDeleteParams) WithContext(ctx context.Context) *ScriptDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the script delete params
func (o *ScriptDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the script delete params
func (o *ScriptDeleteParams) WithHTTPClient(client *http.Client) *ScriptDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the script delete params
func (o *ScriptDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the script delete params
func (o *ScriptDeleteParams) WithAuthorization(authorization string) *ScriptDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the script delete params
func (o *ScriptDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the script delete params
func (o *ScriptDeleteParams) WithID(id string) *ScriptDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the script delete params
func (o *ScriptDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScriptDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
