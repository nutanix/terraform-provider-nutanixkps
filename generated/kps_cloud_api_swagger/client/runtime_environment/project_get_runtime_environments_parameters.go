// Code generated by go-swagger; DO NOT EDIT.

package runtime_environment

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

// NewProjectGetRuntimeEnvironmentsParams creates a new ProjectGetRuntimeEnvironmentsParams object
// with the default values initialized.
func NewProjectGetRuntimeEnvironmentsParams() *ProjectGetRuntimeEnvironmentsParams {
	var ()
	return &ProjectGetRuntimeEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGetRuntimeEnvironmentsParamsWithTimeout creates a new ProjectGetRuntimeEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectGetRuntimeEnvironmentsParamsWithTimeout(timeout time.Duration) *ProjectGetRuntimeEnvironmentsParams {
	var ()
	return &ProjectGetRuntimeEnvironmentsParams{

		timeout: timeout,
	}
}

// NewProjectGetRuntimeEnvironmentsParamsWithContext creates a new ProjectGetRuntimeEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectGetRuntimeEnvironmentsParamsWithContext(ctx context.Context) *ProjectGetRuntimeEnvironmentsParams {
	var ()
	return &ProjectGetRuntimeEnvironmentsParams{

		Context: ctx,
	}
}

// NewProjectGetRuntimeEnvironmentsParamsWithHTTPClient creates a new ProjectGetRuntimeEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectGetRuntimeEnvironmentsParamsWithHTTPClient(client *http.Client) *ProjectGetRuntimeEnvironmentsParams {
	var ()
	return &ProjectGetRuntimeEnvironmentsParams{
		HTTPClient: client,
	}
}

/*ProjectGetRuntimeEnvironmentsParams contains all the parameters to send to the API endpoint
for the project get runtime environments operation typically these are written to a http.Request
*/
type ProjectGetRuntimeEnvironmentsParams struct {

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
	/*PageIndex
	  0-based index of the page to fetch results.

	*/
	PageIndex *int64
	/*PageSize
	  Item count of each page.

	*/
	PageSize *int64
	/*ProjectID
	  ID for the project

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithTimeout(timeout time.Duration) *ProjectGetRuntimeEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithContext(ctx context.Context) *ProjectGetRuntimeEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithHTTPClient(client *http.Client) *ProjectGetRuntimeEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithAuthorization(authorization string) *ProjectGetRuntimeEnvironmentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithFilter(filter *string) *ProjectGetRuntimeEnvironmentsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithOrderBy(orderBy []string) *ProjectGetRuntimeEnvironmentsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPageIndex adds the pageIndex to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithPageIndex(pageIndex *int64) *ProjectGetRuntimeEnvironmentsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithPageSize(pageSize *int64) *ProjectGetRuntimeEnvironmentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) WithProjectID(projectID string) *ProjectGetRuntimeEnvironmentsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project get runtime environments params
func (o *ProjectGetRuntimeEnvironmentsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGetRuntimeEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PageIndex != nil {

		// query param pageIndex
		var qrPageIndex int64
		if o.PageIndex != nil {
			qrPageIndex = *o.PageIndex
		}
		qPageIndex := swag.FormatInt64(qrPageIndex)
		if qPageIndex != "" {
			if err := r.SetQueryParam("pageIndex", qPageIndex); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

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
