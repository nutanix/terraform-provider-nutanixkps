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

// ProxyResponse ProxyResponse is used as websocket httpProxy response message
// swagger:model ProxyResponse
type ProxyResponse struct {

	// response
	Response []uint8 `json:"response"`

	// status
	// Required: true
	Status *string `json:"status"`

	// status code
	// Required: true
	StatusCode *int64 `json:"statusCode"`
}

// Validate validates this proxy response
func (m *ProxyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ProxyResponse) validateStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyResponse) UnmarshalBinary(b []byte) error {
	var res ProxyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
