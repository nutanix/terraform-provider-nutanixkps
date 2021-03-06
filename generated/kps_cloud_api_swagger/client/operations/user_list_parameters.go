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

// NewUserListParams creates a new UserListParams object
// with the default values initialized.
func NewUserListParams() *UserListParams {
	var ()
	return &UserListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserListParamsWithTimeout creates a new UserListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserListParamsWithTimeout(timeout time.Duration) *UserListParams {
	var ()
	return &UserListParams{

		timeout: timeout,
	}
}

// NewUserListParamsWithContext creates a new UserListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserListParamsWithContext(ctx context.Context) *UserListParams {
	var ()
	return &UserListParams{

		Context: ctx,
	}
}

// NewUserListParamsWithHTTPClient creates a new UserListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserListParamsWithHTTPClient(client *http.Client) *UserListParams {
	var ()
	return &UserListParams{
		HTTPClient: client,
	}
}

/*UserListParams contains all the parameters to send to the API endpoint
for the user list operation typically these are written to a http.Request
*/
type UserListParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user list params
func (o *UserListParams) WithTimeout(timeout time.Duration) *UserListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list params
func (o *UserListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list params
func (o *UserListParams) WithContext(ctx context.Context) *UserListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list params
func (o *UserListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list params
func (o *UserListParams) WithHTTPClient(client *http.Client) *UserListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list params
func (o *UserListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user list params
func (o *UserListParams) WithAuthorization(authorization string) *UserListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user list params
func (o *UserListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithFilter adds the filter to the user list params
func (o *UserListParams) WithFilter(filter *string) *UserListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the user list params
func (o *UserListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOrderBy adds the orderBy to the user list params
func (o *UserListParams) WithOrderBy(orderBy []string) *UserListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the user list params
func (o *UserListParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *UserListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
