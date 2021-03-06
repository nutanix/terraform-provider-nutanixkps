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

// NewMLModelVersionCreateParams creates a new MLModelVersionCreateParams object
// with the default values initialized.
func NewMLModelVersionCreateParams() *MLModelVersionCreateParams {
	var ()
	return &MLModelVersionCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMLModelVersionCreateParamsWithTimeout creates a new MLModelVersionCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMLModelVersionCreateParamsWithTimeout(timeout time.Duration) *MLModelVersionCreateParams {
	var ()
	return &MLModelVersionCreateParams{

		timeout: timeout,
	}
}

// NewMLModelVersionCreateParamsWithContext creates a new MLModelVersionCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMLModelVersionCreateParamsWithContext(ctx context.Context) *MLModelVersionCreateParams {
	var ()
	return &MLModelVersionCreateParams{

		Context: ctx,
	}
}

// NewMLModelVersionCreateParamsWithHTTPClient creates a new MLModelVersionCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMLModelVersionCreateParamsWithHTTPClient(client *http.Client) *MLModelVersionCreateParams {
	var ()
	return &MLModelVersionCreateParams{
		HTTPClient: client,
	}
}

/*MLModelVersionCreateParams contains all the parameters to send to the API endpoint
for the m l model version create operation typically these are written to a http.Request
*/
type MLModelVersionCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from the login API response.

	*/
	Authorization string
	/*Payload*/
	Payload runtime.NamedReadCloser
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

// WithTimeout adds the timeout to the m l model version create params
func (o *MLModelVersionCreateParams) WithTimeout(timeout time.Duration) *MLModelVersionCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the m l model version create params
func (o *MLModelVersionCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the m l model version create params
func (o *MLModelVersionCreateParams) WithContext(ctx context.Context) *MLModelVersionCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the m l model version create params
func (o *MLModelVersionCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the m l model version create params
func (o *MLModelVersionCreateParams) WithHTTPClient(client *http.Client) *MLModelVersionCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the m l model version create params
func (o *MLModelVersionCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the m l model version create params
func (o *MLModelVersionCreateParams) WithAuthorization(authorization string) *MLModelVersionCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the m l model version create params
func (o *MLModelVersionCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithPayload adds the payload to the m l model version create params
func (o *MLModelVersionCreateParams) WithPayload(payload runtime.NamedReadCloser) *MLModelVersionCreateParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the m l model version create params
func (o *MLModelVersionCreateParams) SetPayload(payload runtime.NamedReadCloser) {
	o.Payload = payload
}

// WithDescription adds the description to the m l model version create params
func (o *MLModelVersionCreateParams) WithDescription(description *string) *MLModelVersionCreateParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the m l model version create params
func (o *MLModelVersionCreateParams) SetDescription(description *string) {
	o.Description = description
}

// WithID adds the id to the m l model version create params
func (o *MLModelVersionCreateParams) WithID(id string) *MLModelVersionCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the m l model version create params
func (o *MLModelVersionCreateParams) SetID(id string) {
	o.ID = id
}

// WithModelVersion adds the modelVersion to the m l model version create params
func (o *MLModelVersionCreateParams) WithModelVersion(modelVersion int64) *MLModelVersionCreateParams {
	o.SetModelVersion(modelVersion)
	return o
}

// SetModelVersion adds the modelVersion to the m l model version create params
func (o *MLModelVersionCreateParams) SetModelVersion(modelVersion int64) {
	o.ModelVersion = modelVersion
}

// WriteToRequest writes these params to a swagger request
func (o *MLModelVersionCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// form file param Payload
	if err := r.SetFileParam("Payload", o.Payload); err != nil {
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

	// query param model_version
	qrModelVersion := o.ModelVersion
	qModelVersion := swag.FormatInt64(qrModelVersion)
	if qModelVersion != "" {
		if err := r.SetQueryParam("model_version", qModelVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
