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

// NewWsMessagingOnDeleteApplicationParams creates a new WsMessagingOnDeleteApplicationParams object
// with the default values initialized.
func NewWsMessagingOnDeleteApplicationParams() *WsMessagingOnDeleteApplicationParams {
	var ()
	return &WsMessagingOnDeleteApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnDeleteApplicationParamsWithTimeout creates a new WsMessagingOnDeleteApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnDeleteApplicationParamsWithTimeout(timeout time.Duration) *WsMessagingOnDeleteApplicationParams {
	var ()
	return &WsMessagingOnDeleteApplicationParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnDeleteApplicationParamsWithContext creates a new WsMessagingOnDeleteApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnDeleteApplicationParamsWithContext(ctx context.Context) *WsMessagingOnDeleteApplicationParams {
	var ()
	return &WsMessagingOnDeleteApplicationParams{

		Context: ctx,
	}
}

// NewWsMessagingOnDeleteApplicationParamsWithHTTPClient creates a new WsMessagingOnDeleteApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnDeleteApplicationParamsWithHTTPClient(client *http.Client) *WsMessagingOnDeleteApplicationParams {
	var ()
	return &WsMessagingOnDeleteApplicationParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnDeleteApplicationParams contains all the parameters to send to the API endpoint
for the ws messaging on delete application operation typically these are written to a http.Request
*/
type WsMessagingOnDeleteApplicationParams struct {

	/*Request*/
	Request *models.DeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) WithTimeout(timeout time.Duration) *WsMessagingOnDeleteApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) WithContext(ctx context.Context) *WsMessagingOnDeleteApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) WithHTTPClient(client *http.Client) *WsMessagingOnDeleteApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) WithRequest(request *models.DeleteRequest) *WsMessagingOnDeleteApplicationParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on delete application params
func (o *WsMessagingOnDeleteApplicationParams) SetRequest(request *models.DeleteRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnDeleteApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
