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

// ServiceInstanceParam ServiceInstanceParam holds the common parameters for creating or updating Service Instance
// swagger:model ServiceInstanceParam
type ServiceInstanceParam struct {

	// description
	Description string `json:"description,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// scope ID
	// Required: true
	ScopeID *string `json:"scopeId"`

	// svc class ID
	// Required: true
	SvcClassID *string `json:"svcClassId"`
}

// Validate validates this service instance param
func (m *ServiceInstanceParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScopeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcClassID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceInstanceParam) validateScopeID(formats strfmt.Registry) error {

	if err := validate.Required("scopeId", "body", m.ScopeID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceInstanceParam) validateSvcClassID(formats strfmt.Registry) error {

	if err := validate.Required("svcClassId", "body", m.SvcClassID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstanceParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstanceParam) UnmarshalBinary(b []byte) error {
	var res ServiceInstanceParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
