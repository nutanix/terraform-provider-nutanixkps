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

// NewWsMessagingOnUpdateLogCollectorParams creates a new WsMessagingOnUpdateLogCollectorParams object
// with the default values initialized.
func NewWsMessagingOnUpdateLogCollectorParams() *WsMessagingOnUpdateLogCollectorParams {
	var ()
	return &WsMessagingOnUpdateLogCollectorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnUpdateLogCollectorParamsWithTimeout creates a new WsMessagingOnUpdateLogCollectorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnUpdateLogCollectorParamsWithTimeout(timeout time.Duration) *WsMessagingOnUpdateLogCollectorParams {
	var ()
	return &WsMessagingOnUpdateLogCollectorParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnUpdateLogCollectorParamsWithContext creates a new WsMessagingOnUpdateLogCollectorParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnUpdateLogCollectorParamsWithContext(ctx context.Context) *WsMessagingOnUpdateLogCollectorParams {
	var ()
	return &WsMessagingOnUpdateLogCollectorParams{

		Context: ctx,
	}
}

// NewWsMessagingOnUpdateLogCollectorParamsWithHTTPClient creates a new WsMessagingOnUpdateLogCollectorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnUpdateLogCollectorParamsWithHTTPClient(client *http.Client) *WsMessagingOnUpdateLogCollectorParams {
	var ()
	return &WsMessagingOnUpdateLogCollectorParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnUpdateLogCollectorParams contains all the parameters to send to the API endpoint
for the ws messaging on update log collector operation typically these are written to a http.Request
*/
type WsMessagingOnUpdateLogCollectorParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseLogCollector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) WithTimeout(timeout time.Duration) *WsMessagingOnUpdateLogCollectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) WithContext(ctx context.Context) *WsMessagingOnUpdateLogCollectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) WithHTTPClient(client *http.Client) *WsMessagingOnUpdateLogCollectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) WithRequest(request *models.ObjectRequestBaseLogCollector) *WsMessagingOnUpdateLogCollectorParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on update log collector params
func (o *WsMessagingOnUpdateLogCollectorParams) SetRequest(request *models.ObjectRequestBaseLogCollector) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnUpdateLogCollectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
