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

// RenderApplicationPayload RenderApplicationPayload describes edge services on edge for app template
// engine to render application YAML.
// swagger:model RenderApplicationPayload
type RenderApplicationPayload struct {

	// edge services
	EdgeServices map[string]EdgeService `json:"edgeServices,omitempty"`
}

// Validate validates this render application payload
func (m *RenderApplicationPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RenderApplicationPayload) validateEdgeServices(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeServices) { // not required
		return nil
	}

	for k := range m.EdgeServices {

		if err := validate.Required("edgeServices"+"."+k, "body", m.EdgeServices[k]); err != nil {
			return err
		}
		if val, ok := m.EdgeServices[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RenderApplicationPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RenderApplicationPayload) UnmarshalBinary(b []byte) error {
	var res RenderApplicationPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
