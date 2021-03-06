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

// NewWsMessagingHTTPProxyParams creates a new WsMessagingHTTPProxyParams object
// with the default values initialized.
func NewWsMessagingHTTPProxyParams() *WsMessagingHTTPProxyParams {
	var ()
	return &WsMessagingHTTPProxyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingHTTPProxyParamsWithTimeout creates a new WsMessagingHTTPProxyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingHTTPProxyParamsWithTimeout(timeout time.Duration) *WsMessagingHTTPProxyParams {
	var ()
	return &WsMessagingHTTPProxyParams{

		timeout: timeout,
	}
}

// NewWsMessagingHTTPProxyParamsWithContext creates a new WsMessagingHTTPProxyParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingHTTPProxyParamsWithContext(ctx context.Context) *WsMessagingHTTPProxyParams {
	var ()
	return &WsMessagingHTTPProxyParams{

		Context: ctx,
	}
}

// NewWsMessagingHTTPProxyParamsWithHTTPClient creates a new WsMessagingHTTPProxyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingHTTPProxyParamsWithHTTPClient(client *http.Client) *WsMessagingHTTPProxyParams {
	var ()
	return &WsMessagingHTTPProxyParams{
		HTTPClient: client,
	}
}

/*WsMessagingHTTPProxyParams contains all the parameters to send to the API endpoint
for the ws messaging HTTP proxy operation typically these are written to a http.Request
*/
type WsMessagingHTTPProxyParams struct {

	/*Request*/
	Request *models.ProxyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) WithTimeout(timeout time.Duration) *WsMessagingHTTPProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) WithContext(ctx context.Context) *WsMessagingHTTPProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) WithHTTPClient(client *http.Client) *WsMessagingHTTPProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) WithRequest(request *models.ProxyRequest) *WsMessagingHTTPProxyParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging HTTP proxy params
func (o *WsMessagingHTTPProxyParams) SetRequest(request *models.ProxyRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingHTTPProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
