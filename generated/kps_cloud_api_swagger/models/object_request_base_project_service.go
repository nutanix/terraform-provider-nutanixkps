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

// ObjectRequestBaseProjectService ObjectRequestBaseProjectService is used as a websocket ProjectService message
// swagger:model ObjectRequestBaseProjectService
type ObjectRequestBaseProjectService struct {

	// tenant ID
	// Required: true
	TenantID *string `json:"tenantId"`

	// doc
	// Required: true
	Doc *ProjectService `json:"doc"`
}

// Validate validates this object request base project service
func (m *ObjectRequestBaseProjectService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectRequestBaseProjectService) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRequestBaseProjectService) validateDoc(formats strfmt.Registry) error {

	if err := validate.Required("doc", "body", m.Doc); err != nil {
		return err
	}

	if m.Doc != nil {
		if err := m.Doc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("doc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectRequestBaseProjectService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectRequestBaseProjectService) UnmarshalBinary(b []byte) error {
	var res ObjectRequestBaseProjectService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
