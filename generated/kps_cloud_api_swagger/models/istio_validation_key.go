// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// IstioValidationKey IstioValidationKey is the key value composed of an Istio ObjectType and Name.
// swagger:model IstioValidationKey
type IstioValidationKey struct {

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// object type
	ObjectType string `json:"objectType,omitempty"`
}

// Validate validates this istio validation key
func (m *IstioValidationKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IstioValidationKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IstioValidationKey) UnmarshalBinary(b []byte) error {
	var res IstioValidationKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
