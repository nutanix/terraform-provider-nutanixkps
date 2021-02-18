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

// NewWsMessagingOnUpdateProjectServiceParams creates a new WsMessagingOnUpdateProjectServiceParams object
// with the default values initialized.
func NewWsMessagingOnUpdateProjectServiceParams() *WsMessagingOnUpdateProjectServiceParams {
	var ()
	return &WsMessagingOnUpdateProjectServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnUpdateProjectServiceParamsWithTimeout creates a new WsMessagingOnUpdateProjectServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnUpdateProjectServiceParamsWithTimeout(timeout time.Duration) *WsMessagingOnUpdateProjectServiceParams {
	var ()
	return &WsMessagingOnUpdateProjectServiceParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnUpdateProjectServiceParamsWithContext creates a new WsMessagingOnUpdateProjectServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnUpdateProjectServiceParamsWithContext(ctx context.Context) *WsMessagingOnUpdateProjectServiceParams {
	var ()
	return &WsMessagingOnUpdateProjectServiceParams{

		Context: ctx,
	}
}

// NewWsMessagingOnUpdateProjectServiceParamsWithHTTPClient creates a new WsMessagingOnUpdateProjectServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnUpdateProjectServiceParamsWithHTTPClient(client *http.Client) *WsMessagingOnUpdateProjectServiceParams {
	var ()
	return &WsMessagingOnUpdateProjectServiceParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnUpdateProjectServiceParams contains all the parameters to send to the API endpoint
for the ws messaging on update project service operation typically these are written to a http.Request
*/
type WsMessagingOnUpdateProjectServiceParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseProjectService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) WithTimeout(timeout time.Duration) *WsMessagingOnUpdateProjectServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) WithContext(ctx context.Context) *WsMessagingOnUpdateProjectServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) WithHTTPClient(client *http.Client) *WsMessagingOnUpdateProjectServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) WithRequest(request *models.ObjectRequestBaseProjectService) *WsMessagingOnUpdateProjectServiceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on update project service params
func (o *WsMessagingOnUpdateProjectServiceParams) SetRequest(request *models.ObjectRequestBaseProjectService) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnUpdateProjectServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
