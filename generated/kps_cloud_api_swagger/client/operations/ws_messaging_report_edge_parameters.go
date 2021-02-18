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

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// NewWsMessagingReportEdgeParams creates a new WsMessagingReportEdgeParams object
// with the default values initialized.
func NewWsMessagingReportEdgeParams() *WsMessagingReportEdgeParams {
	var ()
	return &WsMessagingReportEdgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingReportEdgeParamsWithTimeout creates a new WsMessagingReportEdgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingReportEdgeParamsWithTimeout(timeout time.Duration) *WsMessagingReportEdgeParams {
	var ()
	return &WsMessagingReportEdgeParams{

		timeout: timeout,
	}
}

// NewWsMessagingReportEdgeParamsWithContext creates a new WsMessagingReportEdgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingReportEdgeParamsWithContext(ctx context.Context) *WsMessagingReportEdgeParams {
	var ()
	return &WsMessagingReportEdgeParams{

		Context: ctx,
	}
}

// NewWsMessagingReportEdgeParamsWithHTTPClient creates a new WsMessagingReportEdgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingReportEdgeParamsWithHTTPClient(client *http.Client) *WsMessagingReportEdgeParams {
	var ()
	return &WsMessagingReportEdgeParams{
		HTTPClient: client,
	}
}

/*WsMessagingReportEdgeParams contains all the parameters to send to the API endpoint
for the ws messaging report edge operation typically these are written to a http.Request
*/
type WsMessagingReportEdgeParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseEdge

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) WithTimeout(timeout time.Duration) *WsMessagingReportEdgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) WithContext(ctx context.Context) *WsMessagingReportEdgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) WithHTTPClient(client *http.Client) *WsMessagingReportEdgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) WithRequest(request *models.ObjectRequestBaseEdge) *WsMessagingReportEdgeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging report edge params
func (o *WsMessagingReportEdgeParams) SetRequest(request *models.ObjectRequestBaseEdge) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingReportEdgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
