// Code generated by go-swagger; DO NOT EDIT.

package cloud_profile

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

// NewCloudProfileListParams creates a new CloudProfileListParams object
// with the default values initialized.
func NewCloudProfileListParams() *CloudProfileListParams {
	var ()
	return &CloudProfileListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudProfileListParamsWithTimeout creates a new CloudProfileListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudProfileListParamsWithTimeout(timeout time.Duration) *CloudProfileListParams {
	var ()
	return &CloudProfileListParams{

		timeout: timeout,
	}
}

// NewCloudProfileListParamsWithContext creates a new CloudProfileListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudProfileListParamsWithContext(ctx context.Context) *CloudProfileListParams {
	var ()
	return &CloudProfileListParams{

		Context: ctx,
	}
}

// NewCloudProfileListParamsWithHTTPClient creates a new CloudProfileListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudProfileListParamsWithHTTPClient(client *http.Client) *CloudProfileListParams {
	var ()
	return &CloudProfileListParams{
		HTTPClient: client,
	}
}

/*CloudProfileListParams contains all the parameters to send to the API endpoint
for the cloud profile list operation typically these are written to a http.Request
*/
type CloudProfileListParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cloud profile list params
func (o *CloudProfileListParams) WithTimeout(timeout time.Duration) *CloudProfileListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud profile list params
func (o *CloudProfileListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud profile list params
func (o *CloudProfileListParams) WithContext(ctx context.Context) *CloudProfileListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud profile list params
func (o *CloudProfileListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud profile list params
func (o *CloudProfileListParams) WithHTTPClient(client *http.Client) *CloudProfileListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud profile list params
func (o *CloudProfileListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud profile list params
func (o *CloudProfileListParams) WithAuthorization(authorization string) *CloudProfileListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud profile list params
func (o *CloudProfileListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the cloud profile list params
func (o *CloudProfileListParams) WithFilter(filter *string) *CloudProfileListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the cloud profile list params
func (o *CloudProfileListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the cloud profile list params
func (o *CloudProfileListParams) WithOrderBy(orderBy []string) *CloudProfileListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the cloud profile list params
func (o *CloudProfileListParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPageIndex adds the pageIndex to the cloud profile list params
func (o *CloudProfileListParams) WithPageIndex(pageIndex *int64) *CloudProfileListParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the cloud profile list params
func (o *CloudProfileListParams) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the cloud profile list params
func (o *CloudProfileListParams) WithPageSize(pageSize *int64) *CloudProfileListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the cloud profile list params
func (o *CloudProfileListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *CloudProfileListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
