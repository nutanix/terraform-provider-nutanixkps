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
)

// NewWsMessagingGetNotificationTopicsParams creates a new WsMessagingGetNotificationTopicsParams object
// with the default values initialized.
func NewWsMessagingGetNotificationTopicsParams() *WsMessagingGetNotificationTopicsParams {

	return &WsMessagingGetNotificationTopicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingGetNotificationTopicsParamsWithTimeout creates a new WsMessagingGetNotificationTopicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingGetNotificationTopicsParamsWithTimeout(timeout time.Duration) *WsMessagingGetNotificationTopicsParams {

	return &WsMessagingGetNotificationTopicsParams{

		timeout: timeout,
	}
}

// NewWsMessagingGetNotificationTopicsParamsWithContext creates a new WsMessagingGetNotificationTopicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingGetNotificationTopicsParamsWithContext(ctx context.Context) *WsMessagingGetNotificationTopicsParams {

	return &WsMessagingGetNotificationTopicsParams{

		Context: ctx,
	}
}

// NewWsMessagingGetNotificationTopicsParamsWithHTTPClient creates a new WsMessagingGetNotificationTopicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingGetNotificationTopicsParamsWithHTTPClient(client *http.Client) *WsMessagingGetNotificationTopicsParams {

	return &WsMessagingGetNotificationTopicsParams{
		HTTPClient: client,
	}
}

/*WsMessagingGetNotificationTopicsParams contains all the parameters to send to the API endpoint
for the ws messaging get notification topics operation typically these are written to a http.Request
*/
type WsMessagingGetNotificationTopicsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging get notification topics params
func (o *WsMessagingGetNotificationTopicsParams) WithTimeout(timeout time.Duration) *WsMessagingGetNotificationTopicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging get notification topics params
func (o *WsMessagingGetNotificationTopicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging get notification topics params
func (o *WsMessagingGetNotificationTopicsParams) WithContext(ctx context.Context) *WsMessagingGetNotificationTopicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging get notification topics params
func (o *WsMessagingGetNotificationTopicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging get notification topics params
func (o *WsMessagingGetNotificationTopicsParams) WithHTTPClient(client *http.Client) *WsMessagingGetNotificationTopicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging get notification topics params
func (o *WsMessagingGetNotificationTopicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingGetNotificationTopicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
