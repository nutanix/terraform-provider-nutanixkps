// Code generated by go-swagger; DO NOT EDIT.

package log

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

// NewApplicationLogEntriesGetV2Params creates a new ApplicationLogEntriesGetV2Params object
// with the default values initialized.
func NewApplicationLogEntriesGetV2Params() *ApplicationLogEntriesGetV2Params {
	var ()
	return &ApplicationLogEntriesGetV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationLogEntriesGetV2ParamsWithTimeout creates a new ApplicationLogEntriesGetV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplicationLogEntriesGetV2ParamsWithTimeout(timeout time.Duration) *ApplicationLogEntriesGetV2Params {
	var ()
	return &ApplicationLogEntriesGetV2Params{

		timeout: timeout,
	}
}

// NewApplicationLogEntriesGetV2ParamsWithContext creates a new ApplicationLogEntriesGetV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewApplicationLogEntriesGetV2ParamsWithContext(ctx context.Context) *ApplicationLogEntriesGetV2Params {
	var ()
	return &ApplicationLogEntriesGetV2Params{

		Context: ctx,
	}
}

// NewApplicationLogEntriesGetV2ParamsWithHTTPClient creates a new ApplicationLogEntriesGetV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApplicationLogEntriesGetV2ParamsWithHTTPClient(client *http.Client) *ApplicationLogEntriesGetV2Params {
	var ()
	return &ApplicationLogEntriesGetV2Params{
		HTTPClient: client,
	}
}

/*ApplicationLogEntriesGetV2Params contains all the parameters to send to the API endpoint
for the application log entries get v2 operation typically these are written to a http.Request
*/
type ApplicationLogEntriesGetV2Params struct {

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
	/*ID
	  ID of the entity

	*/
	ID string
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

// WithTimeout adds the timeout to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithTimeout(timeout time.Duration) *ApplicationLogEntriesGetV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithContext(ctx context.Context) *ApplicationLogEntriesGetV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithHTTPClient(client *http.Client) *ApplicationLogEntriesGetV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithAuthorization(authorization string) *ApplicationLogEntriesGetV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithFilter(filter *string) *ApplicationLogEntriesGetV2Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithID(id string) *ApplicationLogEntriesGetV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetID(id string) {
	o.ID = id
}

// WithOrderBy adds the orderBy to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithOrderBy(orderBy []string) *ApplicationLogEntriesGetV2Params {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPageIndex adds the pageIndex to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithPageIndex(pageIndex *int64) *ApplicationLogEntriesGetV2Params {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) WithPageSize(pageSize *int64) *ApplicationLogEntriesGetV2Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the application log entries get v2 params
func (o *ApplicationLogEntriesGetV2Params) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationLogEntriesGetV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
