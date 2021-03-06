// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditLog audit log
// swagger:model AuditLog
type AuditLog struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// edge i ds
	EdgeIDs string `json:"edgeIds,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// request header
	RequestHeader string `json:"requestHeader,omitempty"`

	// request ID
	RequestID string `json:"requestId,omitempty"`

	// request method
	RequestMethod string `json:"requestMethod,omitempty"`

	// request payload
	RequestPayload string `json:"requestPayload,omitempty"`

	// request URL
	RequestURL string `json:"requestUrl,omitempty"`

	// response code
	ResponseCode int64 `json:"responseCode,omitempty"`

	// response length
	ResponseLength int64 `json:"responseLength,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`

	// started at
	// Format: date-time
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// tenant ID
	TenantID string `json:"tenantId,omitempty"`

	// time m s
	TimeMS float32 `json:"timeMs,omitempty"`

	// user email
	UserEmail string `json:"userEmail,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startedAt", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
