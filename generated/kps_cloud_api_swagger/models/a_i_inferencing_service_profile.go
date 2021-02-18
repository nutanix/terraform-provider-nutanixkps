// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AIInferencingServiceProfile AIInferencingServiceProfile has configuration setting for different framework types.
// swagger:model AIInferencingServiceProfile
type AIInferencingServiceProfile struct {

	// enable
	Enable bool `json:"Enable,omitempty"`

	// runtime
	Runtime []*AIInferencingRuntime `json:"Runtime"`
}

// Validate validates this a i inferencing service profile
func (m *AIInferencingServiceProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuntime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AIInferencingServiceProfile) validateRuntime(formats strfmt.Registry) error {

	if swag.IsZero(m.Runtime) { // not required
		return nil
	}

	for i := 0; i < len(m.Runtime); i++ {
		if swag.IsZero(m.Runtime[i]) { // not required
			continue
		}

		if m.Runtime[i] != nil {
			if err := m.Runtime[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Runtime" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AIInferencingServiceProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AIInferencingServiceProfile) UnmarshalBinary(b []byte) error {
	var res AIInferencingServiceProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
