// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResourcePermissions ResourcePermissions holds permission flags for an object type
// True means allowed.
// swagger:model ResourcePermissions
type ResourcePermissions struct {

	// create
	Create bool `json:"create,omitempty"`

	// delete
	Delete bool `json:"delete,omitempty"`

	// update
	Update bool `json:"update,omitempty"`
}

// Validate validates this resource permissions
func (m *ResourcePermissions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourcePermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePermissions) UnmarshalBinary(b []byte) error {
	var res ResourcePermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
