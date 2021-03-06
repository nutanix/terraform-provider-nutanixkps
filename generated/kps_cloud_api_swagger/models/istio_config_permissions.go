// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// IstioConfigPermissions IstioConfigPermissions holds a map of ResourcesPermissions per namespace
// swagger:model IstioConfigPermissions
type IstioConfigPermissions map[string]ResourcesPermissions

// Validate validates this istio config permissions
func (m IstioConfigPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
