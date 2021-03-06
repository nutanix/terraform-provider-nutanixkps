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

// NewWsMessagingOnCreateLogCollectorParams creates a new WsMessagingOnCreateLogCollectorParams object
// with the default values initialized.
func NewWsMessagingOnCreateLogCollectorParams() *WsMessagingOnCreateLogCollectorParams {
	var ()
	return &WsMessagingOnCreateLogCollectorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnCreateLogCollectorParamsWithTimeout creates a new WsMessagingOnCreateLogCollectorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnCreateLogCollectorParamsWithTimeout(timeout time.Duration) *WsMessagingOnCreateLogCollectorParams {
	var ()
	return &WsMessagingOnCreateLogCollectorParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnCreateLogCollectorParamsWithContext creates a new WsMessagingOnCreateLogCollectorParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnCreateLogCollectorParamsWithContext(ctx context.Context) *WsMessagingOnCreateLogCollectorParams {
	var ()
	return &WsMessagingOnCreateLogCollectorParams{

		Context: ctx,
	}
}

// NewWsMessagingOnCreateLogCollectorParamsWithHTTPClient creates a new WsMessagingOnCreateLogCollectorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnCreateLogCollectorParamsWithHTTPClient(client *http.Client) *WsMessagingOnCreateLogCollectorParams {
	var ()
	return &WsMessagingOnCreateLogCollectorParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnCreateLogCollectorParams contains all the parameters to send to the API endpoint
for the ws messaging on create log collector operation typically these are written to a http.Request
*/
type WsMessagingOnCreateLogCollectorParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseLogCollector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) WithTimeout(timeout time.Duration) *WsMessagingOnCreateLogCollectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) WithContext(ctx context.Context) *WsMessagingOnCreateLogCollectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) WithHTTPClient(client *http.Client) *WsMessagingOnCreateLogCollectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) WithRequest(request *models.ObjectRequestBaseLogCollector) *WsMessagingOnCreateLogCollectorParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on create log collector params
func (o *WsMessagingOnCreateLogCollectorParams) SetRequest(request *models.ObjectRequestBaseLogCollector) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnCreateLogCollectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
