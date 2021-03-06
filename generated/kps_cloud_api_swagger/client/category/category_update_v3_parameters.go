// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewCategoryUpdateV3Params creates a new CategoryUpdateV3Params object
// with the default values initialized.
func NewCategoryUpdateV3Params() *CategoryUpdateV3Params {
	var ()
	return &CategoryUpdateV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCategoryUpdateV3ParamsWithTimeout creates a new CategoryUpdateV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCategoryUpdateV3ParamsWithTimeout(timeout time.Duration) *CategoryUpdateV3Params {
	var ()
	return &CategoryUpdateV3Params{

		timeout: timeout,
	}
}

// NewCategoryUpdateV3ParamsWithContext creates a new CategoryUpdateV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewCategoryUpdateV3ParamsWithContext(ctx context.Context) *CategoryUpdateV3Params {
	var ()
	return &CategoryUpdateV3Params{

		Context: ctx,
	}
}

// NewCategoryUpdateV3ParamsWithHTTPClient creates a new CategoryUpdateV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCategoryUpdateV3ParamsWithHTTPClient(client *http.Client) *CategoryUpdateV3Params {
	var ()
	return &CategoryUpdateV3Params{
		HTTPClient: client,
	}
}

/*CategoryUpdateV3Params contains all the parameters to send to the API endpoint
for the category update v3 operation typically these are written to a http.Request
*/
type CategoryUpdateV3Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.Category
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the category update v3 params
func (o *CategoryUpdateV3Params) WithTimeout(timeout time.Duration) *CategoryUpdateV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the category update v3 params
func (o *CategoryUpdateV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the category update v3 params
func (o *CategoryUpdateV3Params) WithContext(ctx context.Context) *CategoryUpdateV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the category update v3 params
func (o *CategoryUpdateV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the category update v3 params
func (o *CategoryUpdateV3Params) WithHTTPClient(client *http.Client) *CategoryUpdateV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the category update v3 params
func (o *CategoryUpdateV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the category update v3 params
func (o *CategoryUpdateV3Params) WithAuthorization(authorization string) *CategoryUpdateV3Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the category update v3 params
func (o *CategoryUpdateV3Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the category update v3 params
func (o *CategoryUpdateV3Params) WithBody(body *models.Category) *CategoryUpdateV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the category update v3 params
func (o *CategoryUpdateV3Params) SetBody(body *models.Category) {
	o.Body = body
}

// WithID adds the id to the category update v3 params
func (o *CategoryUpdateV3Params) WithID(id string) *CategoryUpdateV3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the category update v3 params
func (o *CategoryUpdateV3Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CategoryUpdateV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
