// Code generated by go-swagger; DO NOT EDIT.

package service_instance

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

// NewServiceInstanceStatusListParams creates a new ServiceInstanceStatusListParams object
// with the default values initialized.
func NewServiceInstanceStatusListParams() *ServiceInstanceStatusListParams {
	var ()
	return &ServiceInstanceStatusListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceInstanceStatusListParamsWithTimeout creates a new ServiceInstanceStatusListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceInstanceStatusListParamsWithTimeout(timeout time.Duration) *ServiceInstanceStatusListParams {
	var ()
	return &ServiceInstanceStatusListParams{

		timeout: timeout,
	}
}

// NewServiceInstanceStatusListParamsWithContext creates a new ServiceInstanceStatusListParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceInstanceStatusListParamsWithContext(ctx context.Context) *ServiceInstanceStatusListParams {
	var ()
	return &ServiceInstanceStatusListParams{

		Context: ctx,
	}
}

// NewServiceInstanceStatusListParamsWithHTTPClient creates a new ServiceInstanceStatusListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceInstanceStatusListParamsWithHTTPClient(client *http.Client) *ServiceInstanceStatusListParams {
	var ()
	return &ServiceInstanceStatusListParams{
		HTTPClient: client,
	}
}

/*ServiceInstanceStatusListParams contains all the parameters to send to the API endpoint
for the service instance status list operation typically these are written to a http.Request
*/
type ServiceInstanceStatusListParams struct {

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
	/*SvcDomainID*/
	SvcDomainID *string
	/*SvcInstanceID
	  Service Instance ID

	*/
	SvcInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithTimeout(timeout time.Duration) *ServiceInstanceStatusListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithContext(ctx context.Context) *ServiceInstanceStatusListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithHTTPClient(client *http.Client) *ServiceInstanceStatusListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithAuthorization(authorization string) *ServiceInstanceStatusListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithFilter(filter *string) *ServiceInstanceStatusListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithOrderBy(orderBy []string) *ServiceInstanceStatusListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPageIndex adds the pageIndex to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithPageIndex(pageIndex *int64) *ServiceInstanceStatusListParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithPageSize(pageSize *int64) *ServiceInstanceStatusListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSvcDomainID adds the svcDomainID to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithSvcDomainID(svcDomainID *string) *ServiceInstanceStatusListParams {
	o.SetSvcDomainID(svcDomainID)
	return o
}

// SetSvcDomainID adds the svcDomainId to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetSvcDomainID(svcDomainID *string) {
	o.SvcDomainID = svcDomainID
}

// WithSvcInstanceID adds the svcInstanceID to the service instance status list params
func (o *ServiceInstanceStatusListParams) WithSvcInstanceID(svcInstanceID string) *ServiceInstanceStatusListParams {
	o.SetSvcInstanceID(svcInstanceID)
	return o
}

// SetSvcInstanceID adds the svcInstanceId to the service instance status list params
func (o *ServiceInstanceStatusListParams) SetSvcInstanceID(svcInstanceID string) {
	o.SvcInstanceID = svcInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceInstanceStatusListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SvcDomainID != nil {

		// query param svcDomainId
		var qrSvcDomainID string
		if o.SvcDomainID != nil {
			qrSvcDomainID = *o.SvcDomainID
		}
		qSvcDomainID := qrSvcDomainID
		if qSvcDomainID != "" {
			if err := r.SetQueryParam("svcDomainId", qSvcDomainID); err != nil {
				return err
			}
		}

	}

	// path param svcInstanceId
	if err := r.SetPathParam("svcInstanceId", o.SvcInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
