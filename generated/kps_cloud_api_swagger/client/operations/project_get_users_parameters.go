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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProjectGetUsersParams creates a new ProjectGetUsersParams object
// with the default values initialized.
func NewProjectGetUsersParams() *ProjectGetUsersParams {
	var ()
	return &ProjectGetUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGetUsersParamsWithTimeout creates a new ProjectGetUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectGetUsersParamsWithTimeout(timeout time.Duration) *ProjectGetUsersParams {
	var ()
	return &ProjectGetUsersParams{

		timeout: timeout,
	}
}

// NewProjectGetUsersParamsWithContext creates a new ProjectGetUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectGetUsersParamsWithContext(ctx context.Context) *ProjectGetUsersParams {
	var ()
	return &ProjectGetUsersParams{

		Context: ctx,
	}
}

// NewProjectGetUsersParamsWithHTTPClient creates a new ProjectGetUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectGetUsersParamsWithHTTPClient(client *http.Client) *ProjectGetUsersParams {
	var ()
	return &ProjectGetUsersParams{
		HTTPClient: client,
	}
}

/*ProjectGetUsersParams contains all the parameters to send to the API endpoint
for the project get users operation typically these are written to a http.Request
*/
type ProjectGetUsersParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Filter
	  Specify result filter. Format is similar to a SQL WHERE clause. For example,
	to filter object by name with prefix foo, use: name LIKE 'foo%'.
	Supported filter keys are the same as order by keys.

	*/
	Filter *string
	/*OrderBy
	  Specify result order. Zero or more entries with format: &ltkey> [desc]
	where orderByKeys lists allowed keys in each response.

	*/
	OrderBy []string
	/*ProjectID
	  ID for the project

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project get users params
func (o *ProjectGetUsersParams) WithTimeout(timeout time.Duration) *ProjectGetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project get users params
func (o *ProjectGetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project get users params
func (o *ProjectGetUsersParams) WithContext(ctx context.Context) *ProjectGetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project get users params
func (o *ProjectGetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project get users params
func (o *ProjectGetUsersParams) WithHTTPClient(client *http.Client) *ProjectGetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project get users params
func (o *ProjectGetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project get users params
func (o *ProjectGetUsersParams) WithAuthorization(authorization string) *ProjectGetUsersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project get users params
func (o *ProjectGetUsersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the project get users params
func (o *ProjectGetUsersParams) WithFilter(filter *string) *ProjectGetUsersParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the project get users params
func (o *ProjectGetUsersParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the project get users params
func (o *ProjectGetUsersParams) WithOrderBy(orderBy []string) *ProjectGetUsersParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the project get users params
func (o *ProjectGetUsersParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProjectID adds the projectID to the project get users params
func (o *ProjectGetUsersParams) WithProjectID(projectID string) *ProjectGetUsersParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project get users params
func (o *ProjectGetUsersParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	valuesOrderBy := o.OrderBy

	joinedOrderBy := swag.JoinByFormat(valuesOrderBy, "")
	// query array param orderBy
	if err := r.SetQueryParam("orderBy", joinedOrderBy...); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
