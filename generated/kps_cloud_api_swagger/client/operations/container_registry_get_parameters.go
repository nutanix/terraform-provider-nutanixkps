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

// NewContainerRegistryGetParams creates a new ContainerRegistryGetParams object
// with the default values initialized.
func NewContainerRegistryGetParams() *ContainerRegistryGetParams {
	var ()
	return &ContainerRegistryGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRegistryGetParamsWithTimeout creates a new ContainerRegistryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerRegistryGetParamsWithTimeout(timeout time.Duration) *ContainerRegistryGetParams {
	var ()
	return &ContainerRegistryGetParams{

		timeout: timeout,
	}
}

// NewContainerRegistryGetParamsWithContext creates a new ContainerRegistryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerRegistryGetParamsWithContext(ctx context.Context) *ContainerRegistryGetParams {
	var ()
	return &ContainerRegistryGetParams{

		Context: ctx,
	}
}

// NewContainerRegistryGetParamsWithHTTPClient creates a new ContainerRegistryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerRegistryGetParamsWithHTTPClient(client *http.Client) *ContainerRegistryGetParams {
	var ()
	return &ContainerRegistryGetParams{
		HTTPClient: client,
	}
}

/*ContainerRegistryGetParams contains all the parameters to send to the API endpoint
for the container registry get operation typically these are written to a http.Request
*/
type ContainerRegistryGetParams struct {

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

// WithTimeout adds the timeout to the container registry get params
func (o *ContainerRegistryGetParams) WithTimeout(timeout time.Duration) *ContainerRegistryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container registry get params
func (o *ContainerRegistryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container registry get params
func (o *ContainerRegistryGetParams) WithContext(ctx context.Context) *ContainerRegistryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container registry get params
func (o *ContainerRegistryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container registry get params
func (o *ContainerRegistryGetParams) WithHTTPClient(client *http.Client) *ContainerRegistryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container registry get params
func (o *ContainerRegistryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the container registry get params
func (o *ContainerRegistryGetParams) WithAuthorization(authorization string) *ContainerRegistryGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the container registry get params
func (o *ContainerRegistryGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the container registry get params
func (o *ContainerRegistryGetParams) WithID(id string) *ContainerRegistryGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container registry get params
func (o *ContainerRegistryGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRegistryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
