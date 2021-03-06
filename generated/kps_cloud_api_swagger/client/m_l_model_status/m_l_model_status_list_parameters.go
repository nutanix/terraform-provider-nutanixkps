// Code generated by go-swagger; DO NOT EDIT.

package m_l_model_status

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

// NewMLModelStatusListParams creates a new MLModelStatusListParams object
// with the default values initialized.
func NewMLModelStatusListParams() *MLModelStatusListParams {
	var ()
	return &MLModelStatusListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMLModelStatusListParamsWithTimeout creates a new MLModelStatusListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMLModelStatusListParamsWithTimeout(timeout time.Duration) *MLModelStatusListParams {
	var ()
	return &MLModelStatusListParams{

		timeout: timeout,
	}
}

// NewMLModelStatusListParamsWithContext creates a new MLModelStatusListParams object
// with the default values initialized, and the ability to set a context for a request
func NewMLModelStatusListParamsWithContext(ctx context.Context) *MLModelStatusListParams {
	var ()
	return &MLModelStatusListParams{

		Context: ctx,
	}
}

// NewMLModelStatusListParamsWithHTTPClient creates a new MLModelStatusListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMLModelStatusListParamsWithHTTPClient(client *http.Client) *MLModelStatusListParams {
	var ()
	return &MLModelStatusListParams{
		HTTPClient: client,
	}
}

/*MLModelStatusListParams contains all the parameters to send to the API endpoint
for the m l model status list operation typically these are written to a http.Request
*/
type MLModelStatusListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
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

// WithTimeout adds the timeout to the m l model status list params
func (o *MLModelStatusListParams) WithTimeout(timeout time.Duration) *MLModelStatusListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the m l model status list params
func (o *MLModelStatusListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the m l model status list params
func (o *MLModelStatusListParams) WithContext(ctx context.Context) *MLModelStatusListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the m l model status list params
func (o *MLModelStatusListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the m l model status list params
func (o *MLModelStatusListParams) WithHTTPClient(client *http.Client) *MLModelStatusListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the m l model status list params
func (o *MLModelStatusListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the m l model status list params
func (o *MLModelStatusListParams) WithAuthorization(authorization string) *MLModelStatusListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the m l model status list params
func (o *MLModelStatusListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithPageIndex adds the pageIndex to the m l model status list params
func (o *MLModelStatusListParams) WithPageIndex(pageIndex *int64) *MLModelStatusListParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the m l model status list params
func (o *MLModelStatusListParams) SetPageIndex(pageIndex *int64) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the m l model status list params
func (o *MLModelStatusListParams) WithPageSize(pageSize *int64) *MLModelStatusListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the m l model status list params
func (o *MLModelStatusListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *MLModelStatusListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
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
