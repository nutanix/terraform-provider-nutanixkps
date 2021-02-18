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

// ProxyRequest ProxyRequest is used as websocket httpProxy request message
// swagger:model ProxyRequest
type ProxyRequest struct {

	// request
	// Required: true
	Request []uint8 `json:"request"`

	// URL
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this proxy request
func (m *ProxyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyRequest) validateRequest(formats strfmt.Registry) error {

	if err := validate.Required("request", "body", m.Request); err != nil {
		return err
	}

	return nil
}

func (m *ProxyRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyRequest) UnmarshalBinary(b []byte) error {
	var res ProxyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
