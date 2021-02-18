// Code generated by go-swagger; DO NOT EDIT.

package software_update

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

// NewSoftwareUpgradeServiceDomainListParams creates a new SoftwareUpgradeServiceDomainListParams object
// with the default values initialized.
func NewSoftwareUpgradeServiceDomainListParams() *SoftwareUpgradeServiceDomainListParams {
	var ()
	return &SoftwareUpgradeServiceDomainListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareUpgradeServiceDomainListParamsWithTimeout creates a new SoftwareUpgradeServiceDomainListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareUpgradeServiceDomainListParamsWithTimeout(timeout time.Duration) *SoftwareUpgradeServiceDomainListParams {
	var ()
	return &SoftwareUpgradeServiceDomainListParams{

		timeout: timeout,
	}
}

// NewSoftwareUpgradeServiceDomainListParamsWithContext creates a new SoftwareUpgradeServiceDomainListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareUpgradeServiceDomainListParamsWithContext(ctx context.Context) *SoftwareUpgradeServiceDomainListParams {
	var ()
	return &SoftwareUpgradeServiceDomainListParams{

		Context: ctx,
	}
}

// NewSoftwareUpgradeServiceDomainListParamsWithHTTPClient creates a new SoftwareUpgradeServiceDomainListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareUpgradeServiceDomainListParamsWithHTTPClient(client *http.Client) *SoftwareUpgradeServiceDomainListParams {
	var ()
	return &SoftwareUpgradeServiceDomainListParams{
		HTTPClient: client,
	}
}

/*SoftwareUpgradeServiceDomainListParams contains all the parameters to send to the API endpoint
for the software upgrade service domain list operation typically these are written to a http.Request
*/
type SoftwareUpgradeServiceDomainListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*BatchID
	  ID for the batch

	*/
	BatchID string
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

// WithTimeout adds the timeout to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithTimeout(timeout time.Duration) *SoftwareUpgradeServiceDomainListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithContext(ctx context.Context) *SoftwareUpgradeServiceDomainListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithHTTPClient(client *http.Client) *SoftwareUpgradeServiceDomainListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithAuthorization(authorization string) *SoftwareUpgradeServiceDomainListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBatchID adds the batchID to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithBatchID(batchID string) *SoftwareUpgradeServiceDomainListParams {
	o.SetBatchID(batchID)
	return o
}

// SetBatchID adds the batchId to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetBatchID(batchID string) {
	o.BatchID = batchID
}

// WithFilter adds the filter to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithFilter(filter *string) *SoftwareUpgradeServiceDomainListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithOrderBy(orderBy []string) *SoftwareUpgradeServiceDomainListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPageIndex adds the pageIndex to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithPageIndex(pageIndex *int64) *SoftwareUpgradeServiceDomainListParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) WithPageSize(pageSize *int64) *SoftwareUpgradeServiceDomainListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the software upgrade service domain list params
func (o *SoftwareUpgradeServiceDomainListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareUpgradeServiceDomainListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param batchId
	if err := r.SetPathParam("batchId", o.BatchID); err != nil {
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
