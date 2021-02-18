// Code generated by go-swagger; DO NOT EDIT.

package m_l_model

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

// NewMLModelVersionUpdateParams creates a new MLModelVersionUpdateParams object
// with the default values initialized.
func NewMLModelVersionUpdateParams() *MLModelVersionUpdateParams {
	var ()
	return &MLModelVersionUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMLModelVersionUpdateParamsWithTimeout creates a new MLModelVersionUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMLModelVersionUpdateParamsWithTimeout(timeout time.Duration) *MLModelVersionUpdateParams {
	var ()
	return &MLModelVersionUpdateParams{

		timeout: timeout,
	}
}

// NewMLModelVersionUpdateParamsWithContext creates a new MLModelVersionUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMLModelVersionUpdateParamsWithContext(ctx context.Context) *MLModelVersionUpdateParams {
	var ()
	return &MLModelVersionUpdateParams{

		Context: ctx,
	}
}

// NewMLModelVersionUpdateParamsWithHTTPClient creates a new MLModelVersionUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMLModelVersionUpdateParamsWithHTTPClient(client *http.Client) *MLModelVersionUpdateParams {
	var ()
	return &MLModelVersionUpdateParams{
		HTTPClient: client,
	}
}

/*MLModelVersionUpdateParams contains all the parameters to send to the API endpoint
for the m l model version update operation typically these are written to a http.Request
*/
type MLModelVersionUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from the login API response.

	*/
	Authorization string
	/*Description
	  Model version description.

	*/
	Description *string
	/*ID
	  ID of the entity

	*/
	ID string
	/*ModelVersion
	  Model version, a positive integer.

	*/
	ModelVersion int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the m l model version update params
func (o *MLModelVersionUpdateParams) WithTimeout(timeout time.Duration) *MLModelVersionUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the m l model version update params
func (o *MLModelVersionUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the m l model version update params
func (o *MLModelVersionUpdateParams) WithContext(ctx context.Context) *MLModelVersionUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the m l model version update params
func (o *MLModelVersionUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the m l model version update params
func (o *MLModelVersionUpdateParams) WithHTTPClient(client *http.Client) *MLModelVersionUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the m l model version update params
func (o *MLModelVersionUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the m l model version update params
func (o *MLModelVersionUpdateParams) WithAuthorization(authorization string) *MLModelVersionUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the m l model version update params
func (o *MLModelVersionUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDescription adds the description to the m l model version update params
func (o *MLModelVersionUpdateParams) WithDescription(description *string) *MLModelVersionUpdateParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the m l model version update params
func (o *MLModelVersionUpdateParams) SetDescription(description *string) {
	o.Description = description
}

// WithID adds the id to the m l model version update params
func (o *MLModelVersionUpdateParams) WithID(id string) *MLModelVersionUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the m l model version update params
func (o *MLModelVersionUpdateParams) SetID(id string) {
	o.ID = id
}

// WithModelVersion adds the modelVersion to the m l model version update params
func (o *MLModelVersionUpdateParams) WithModelVersion(modelVersion int64) *MLModelVersionUpdateParams {
	o.SetModelVersion(modelVersion)
	return o
}

// SetModelVersion adds the modelVersion to the m l model version update params
func (o *MLModelVersionUpdateParams) SetModelVersion(modelVersion int64) {
	o.ModelVersion = modelVersion
}

// WriteToRequest writes these params to a swagger request
func (o *MLModelVersionUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param model_version
	if err := r.SetPathParam("model_version", swag.FormatInt64(o.ModelVersion)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
