// Code generated by go-swagger; DO NOT EDIT.

package user_api_token

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

// NewUserAPITokenUpdateParams creates a new UserAPITokenUpdateParams object
// with the default values initialized.
func NewUserAPITokenUpdateParams() *UserAPITokenUpdateParams {
	var ()
	return &UserAPITokenUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAPITokenUpdateParamsWithTimeout creates a new UserAPITokenUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAPITokenUpdateParamsWithTimeout(timeout time.Duration) *UserAPITokenUpdateParams {
	var ()
	return &UserAPITokenUpdateParams{

		timeout: timeout,
	}
}

// NewUserAPITokenUpdateParamsWithContext creates a new UserAPITokenUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAPITokenUpdateParamsWithContext(ctx context.Context) *UserAPITokenUpdateParams {
	var ()
	return &UserAPITokenUpdateParams{

		Context: ctx,
	}
}

// NewUserAPITokenUpdateParamsWithHTTPClient creates a new UserAPITokenUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAPITokenUpdateParamsWithHTTPClient(client *http.Client) *UserAPITokenUpdateParams {
	var ()
	return &UserAPITokenUpdateParams{
		HTTPClient: client,
	}
}

/*UserAPITokenUpdateParams contains all the parameters to send to the API endpoint
for the user Api token update operation typically these are written to a http.Request
*/
type UserAPITokenUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.UserAPIToken
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user Api token update params
func (o *UserAPITokenUpdateParams) WithTimeout(timeout time.Duration) *UserAPITokenUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user Api token update params
func (o *UserAPITokenUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user Api token update params
func (o *UserAPITokenUpdateParams) WithContext(ctx context.Context) *UserAPITokenUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user Api token update params
func (o *UserAPITokenUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user Api token update params
func (o *UserAPITokenUpdateParams) WithHTTPClient(client *http.Client) *UserAPITokenUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user Api token update params
func (o *UserAPITokenUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user Api token update params
func (o *UserAPITokenUpdateParams) WithAuthorization(authorization string) *UserAPITokenUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user Api token update params
func (o *UserAPITokenUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the user Api token update params
func (o *UserAPITokenUpdateParams) WithBody(body *models.UserAPIToken) *UserAPITokenUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user Api token update params
func (o *UserAPITokenUpdateParams) SetBody(body *models.UserAPIToken) {
	o.Body = body
}

// WithID adds the id to the user Api token update params
func (o *UserAPITokenUpdateParams) WithID(id string) *UserAPITokenUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user Api token update params
func (o *UserAPITokenUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserAPITokenUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
