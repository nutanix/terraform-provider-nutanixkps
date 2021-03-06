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

// NewWsMessagingTeardownSSHTunnelingParams creates a new WsMessagingTeardownSSHTunnelingParams object
// with the default values initialized.
func NewWsMessagingTeardownSSHTunnelingParams() *WsMessagingTeardownSSHTunnelingParams {
	var ()
	return &WsMessagingTeardownSSHTunnelingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingTeardownSSHTunnelingParamsWithTimeout creates a new WsMessagingTeardownSSHTunnelingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingTeardownSSHTunnelingParamsWithTimeout(timeout time.Duration) *WsMessagingTeardownSSHTunnelingParams {
	var ()
	return &WsMessagingTeardownSSHTunnelingParams{

		timeout: timeout,
	}
}

// NewWsMessagingTeardownSSHTunnelingParamsWithContext creates a new WsMessagingTeardownSSHTunnelingParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingTeardownSSHTunnelingParamsWithContext(ctx context.Context) *WsMessagingTeardownSSHTunnelingParams {
	var ()
	return &WsMessagingTeardownSSHTunnelingParams{

		Context: ctx,
	}
}

// NewWsMessagingTeardownSSHTunnelingParamsWithHTTPClient creates a new WsMessagingTeardownSSHTunnelingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingTeardownSSHTunnelingParamsWithHTTPClient(client *http.Client) *WsMessagingTeardownSSHTunnelingParams {
	var ()
	return &WsMessagingTeardownSSHTunnelingParams{
		HTTPClient: client,
	}
}

/*WsMessagingTeardownSSHTunnelingParams contains all the parameters to send to the API endpoint
for the ws messaging teardown SSH tunneling operation typically these are written to a http.Request
*/
type WsMessagingTeardownSSHTunnelingParams struct {

	/*Request*/
	Request *models.ObjectRequestTeardownSSHTunneling

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) WithTimeout(timeout time.Duration) *WsMessagingTeardownSSHTunnelingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) WithContext(ctx context.Context) *WsMessagingTeardownSSHTunnelingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) WithHTTPClient(client *http.Client) *WsMessagingTeardownSSHTunnelingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) WithRequest(request *models.ObjectRequestTeardownSSHTunneling) *WsMessagingTeardownSSHTunnelingParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging teardown SSH tunneling params
func (o *WsMessagingTeardownSSHTunnelingParams) SetRequest(request *models.ObjectRequestTeardownSSHTunneling) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingTeardownSSHTunnelingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
