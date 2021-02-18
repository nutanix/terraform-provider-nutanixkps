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

// ObjectRequestBaseDataSource ObjectRequestBaseDataSource is used as websocket DataSource message
// swagger:model ObjectRequestBaseDataSource
type ObjectRequestBaseDataSource struct {

	// tenant ID
	// Required: true
	TenantID *string `json:"tenantId"`

	// doc
	// Required: true
	Doc *DataSource `json:"doc"`
}

// Validate validates this object request base data source
func (m *ObjectRequestBaseDataSource) Validate(formats strfmt.Registry) error {
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

func (m *ObjectRequestBaseDataSource) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRequestBaseDataSource) validateDoc(formats strfmt.Registry) error {

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
func (m *ObjectRequestBaseDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectRequestBaseDataSource) UnmarshalBinary(b []byte) error {
	var res ObjectRequestBaseDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
