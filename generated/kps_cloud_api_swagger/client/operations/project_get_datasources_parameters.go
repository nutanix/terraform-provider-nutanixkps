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

// NewProjectGetDatasourcesParams creates a new ProjectGetDatasourcesParams object
// with the default values initialized.
func NewProjectGetDatasourcesParams() *ProjectGetDatasourcesParams {
	var ()
	return &ProjectGetDatasourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGetDatasourcesParamsWithTimeout creates a new ProjectGetDatasourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectGetDatasourcesParamsWithTimeout(timeout time.Duration) *ProjectGetDatasourcesParams {
	var ()
	return &ProjectGetDatasourcesParams{

		timeout: timeout,
	}
}

// NewProjectGetDatasourcesParamsWithContext creates a new ProjectGetDatasourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectGetDatasourcesParamsWithContext(ctx context.Context) *ProjectGetDatasourcesParams {
	var ()
	return &ProjectGetDatasourcesParams{

		Context: ctx,
	}
}

// NewProjectGetDatasourcesParamsWithHTTPClient creates a new ProjectGetDatasourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectGetDatasourcesParamsWithHTTPClient(client *http.Client) *ProjectGetDatasourcesParams {
	var ()
	return &ProjectGetDatasourcesParams{
		HTTPClient: client,
	}
}

/*ProjectGetDatasourcesParams contains all the parameters to send to the API endpoint
for the project get datasources operation typically these are written to a http.Request
*/
type ProjectGetDatasourcesParams struct {

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

// WithTimeout adds the timeout to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithTimeout(timeout time.Duration) *ProjectGetDatasourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithContext(ctx context.Context) *ProjectGetDatasourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithHTTPClient(client *http.Client) *ProjectGetDatasourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithAuthorization(authorization string) *ProjectGetDatasourcesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithFilter(filter *string) *ProjectGetDatasourcesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithOrderBy(orderBy []string) *ProjectGetDatasourcesParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProjectID adds the projectID to the project get datasources params
func (o *ProjectGetDatasourcesParams) WithProjectID(projectID string) *ProjectGetDatasourcesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project get datasources params
func (o *ProjectGetDatasourcesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGetDatasourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
