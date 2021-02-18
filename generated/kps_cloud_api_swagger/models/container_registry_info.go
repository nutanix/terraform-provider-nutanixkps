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

// ContainerRegistryInfo container registry info
// swagger:model ContainerRegistryInfo
type ContainerRegistryInfo struct {

	// Email address for the container registry profile user.
	// Required: true
	Email *string `json:"email"`

	// Password for the container registry profile.
	// Required: true
	Pwd *string `json:"pwd"`

	// User name for the container registry profile.
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this container registry info
func (m *ContainerRegistryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePwd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerRegistryInfo) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *ContainerRegistryInfo) validatePwd(formats strfmt.Registry) error {

	if err := validate.Required("pwd", "body", m.Pwd); err != nil {
		return err
	}

	return nil
}

func (m *ContainerRegistryInfo) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerRegistryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerRegistryInfo) UnmarshalBinary(b []byte) error {
	var res ContainerRegistryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
