// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MLModelVersionURLGetResponsePayload m l model version URL get response payload
// swagger:model MLModelVersionURLGetResponsePayload
type MLModelVersionURLGetResponsePayload struct {

	// URL
	URL string `json:"url,omitempty"`
}

// Validate validates this m l model version URL get response payload
func (m *MLModelVersionURLGetResponsePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MLModelVersionURLGetResponsePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MLModelVersionURLGetResponsePayload) UnmarshalBinary(b []byte) error {
	var res MLModelVersionURLGetResponsePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
