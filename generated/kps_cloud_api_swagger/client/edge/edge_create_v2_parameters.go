// Code generated by go-swagger; DO NOT EDIT.

package edge

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

// NewEdgeCreateV2Params creates a new EdgeCreateV2Params object
// with the default values initialized.
func NewEdgeCreateV2Params() *EdgeCreateV2Params {
	var ()
	return &EdgeCreateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeCreateV2ParamsWithTimeout creates a new EdgeCreateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeCreateV2ParamsWithTimeout(timeout time.Duration) *EdgeCreateV2Params {
	var ()
	return &EdgeCreateV2Params{

		timeout: timeout,
	}
}

// NewEdgeCreateV2ParamsWithContext creates a new EdgeCreateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeCreateV2ParamsWithContext(ctx context.Context) *EdgeCreateV2Params {
	var ()
	return &EdgeCreateV2Params{

		Context: ctx,
	}
}

// NewEdgeCreateV2ParamsWithHTTPClient creates a new EdgeCreateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeCreateV2ParamsWithHTTPClient(client *http.Client) *EdgeCreateV2Params {
	var ()
	return &EdgeCreateV2Params{
		HTTPClient: client,
	}
}

/*EdgeCreateV2Params contains all the parameters to send to the API endpoint
for the edge create v2 operation typically these are written to a http.Request
*/
type EdgeCreateV2Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  Parameters and values used when creating an edge

	*/
	Body *models.EdgeV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge create v2 params
func (o *EdgeCreateV2Params) WithTimeout(timeout time.Duration) *EdgeCreateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge create v2 params
func (o *EdgeCreateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge create v2 params
func (o *EdgeCreateV2Params) WithContext(ctx context.Context) *EdgeCreateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge create v2 params
func (o *EdgeCreateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge create v2 params
func (o *EdgeCreateV2Params) WithHTTPClient(client *http.Client) *EdgeCreateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge create v2 params
func (o *EdgeCreateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the edge create v2 params
func (o *EdgeCreateV2Params) WithAuthorization(authorization string) *EdgeCreateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the edge create v2 params
func (o *EdgeCreateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the edge create v2 params
func (o *EdgeCreateV2Params) WithBody(body *models.EdgeV2) *EdgeCreateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge create v2 params
func (o *EdgeCreateV2Params) SetBody(body *models.EdgeV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeCreateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
