// Code generated by go-swagger; DO NOT EDIT.

package container_registry

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

// NewProjectGetContainerRegistriesV2Params creates a new ProjectGetContainerRegistriesV2Params object
// with the default values initialized.
func NewProjectGetContainerRegistriesV2Params() *ProjectGetContainerRegistriesV2Params {
	var ()
	return &ProjectGetContainerRegistriesV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGetContainerRegistriesV2ParamsWithTimeout creates a new ProjectGetContainerRegistriesV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectGetContainerRegistriesV2ParamsWithTimeout(timeout time.Duration) *ProjectGetContainerRegistriesV2Params {
	var ()
	return &ProjectGetContainerRegistriesV2Params{

		timeout: timeout,
	}
}

// NewProjectGetContainerRegistriesV2ParamsWithContext creates a new ProjectGetContainerRegistriesV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewProjectGetContainerRegistriesV2ParamsWithContext(ctx context.Context) *ProjectGetContainerRegistriesV2Params {
	var ()
	return &ProjectGetContainerRegistriesV2Params{

		Context: ctx,
	}
}

// NewProjectGetContainerRegistriesV2ParamsWithHTTPClient creates a new ProjectGetContainerRegistriesV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectGetContainerRegistriesV2ParamsWithHTTPClient(client *http.Client) *ProjectGetContainerRegistriesV2Params {
	var ()
	return &ProjectGetContainerRegistriesV2Params{
		HTTPClient: client,
	}
}

/*ProjectGetContainerRegistriesV2Params contains all the parameters to send to the API endpoint
for the project get container registries v2 operation typically these are written to a http.Request
*/
type ProjectGetContainerRegistriesV2Params struct {

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

// WithTimeout adds the timeout to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithTimeout(timeout time.Duration) *ProjectGetContainerRegistriesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithContext(ctx context.Context) *ProjectGetContainerRegistriesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithHTTPClient(client *http.Client) *ProjectGetContainerRegistriesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithAuthorization(authorization string) *ProjectGetContainerRegistriesV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithFilter(filter *string) *ProjectGetContainerRegistriesV2Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithOrderBy(orderBy []string) *ProjectGetContainerRegistriesV2Params {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPageIndex adds the pageIndex to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithPageIndex(pageIndex *int64) *ProjectGetContainerRegistriesV2Params {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithPageSize(pageSize *int64) *ProjectGetContainerRegistriesV2Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) WithProjectID(projectID string) *ProjectGetContainerRegistriesV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project get container registries v2 params
func (o *ProjectGetContainerRegistriesV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGetContainerRegistriesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
