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

// NewProjectGetScriptsParams creates a new ProjectGetScriptsParams object
// with the default values initialized.
func NewProjectGetScriptsParams() *ProjectGetScriptsParams {
	var ()
	return &ProjectGetScriptsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGetScriptsParamsWithTimeout creates a new ProjectGetScriptsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectGetScriptsParamsWithTimeout(timeout time.Duration) *ProjectGetScriptsParams {
	var ()
	return &ProjectGetScriptsParams{

		timeout: timeout,
	}
}

// NewProjectGetScriptsParamsWithContext creates a new ProjectGetScriptsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectGetScriptsParamsWithContext(ctx context.Context) *ProjectGetScriptsParams {
	var ()
	return &ProjectGetScriptsParams{

		Context: ctx,
	}
}

// NewProjectGetScriptsParamsWithHTTPClient creates a new ProjectGetScriptsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectGetScriptsParamsWithHTTPClient(client *http.Client) *ProjectGetScriptsParams {
	var ()
	return &ProjectGetScriptsParams{
		HTTPClient: client,
	}
}

/*ProjectGetScriptsParams contains all the parameters to send to the API endpoint
for the project get scripts operation typically these are written to a http.Request
*/
type ProjectGetScriptsParams struct {

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

// WithTimeout adds the timeout to the project get scripts params
func (o *ProjectGetScriptsParams) WithTimeout(timeout time.Duration) *ProjectGetScriptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project get scripts params
func (o *ProjectGetScriptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project get scripts params
func (o *ProjectGetScriptsParams) WithContext(ctx context.Context) *ProjectGetScriptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project get scripts params
func (o *ProjectGetScriptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project get scripts params
func (o *ProjectGetScriptsParams) WithHTTPClient(client *http.Client) *ProjectGetScriptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project get scripts params
func (o *ProjectGetScriptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project get scripts params
func (o *ProjectGetScriptsParams) WithAuthorization(authorization string) *ProjectGetScriptsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project get scripts params
func (o *ProjectGetScriptsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the project get scripts params
func (o *ProjectGetScriptsParams) WithFilter(filter *string) *ProjectGetScriptsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the project get scripts params
func (o *ProjectGetScriptsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the project get scripts params
func (o *ProjectGetScriptsParams) WithOrderBy(orderBy []string) *ProjectGetScriptsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the project get scripts params
func (o *ProjectGetScriptsParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProjectID adds the projectID to the project get scripts params
func (o *ProjectGetScriptsParams) WithProjectID(projectID string) *ProjectGetScriptsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project get scripts params
func (o *ProjectGetScriptsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGetScriptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
