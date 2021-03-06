// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KialiMetric KialiMetric holds the Prometheus Matrix model, which contains one or more time series (depending on grouping)
// swagger:model KialiMetric
type KialiMetric struct {

	// matrix
	Matrix Matrix `json:"matrix,omitempty"`
}

// Validate validates this kiali metric
func (m *KialiMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KialiMetric) validateMatrix(formats strfmt.Registry) error {

	if swag.IsZero(m.Matrix) { // not required
		return nil
	}

	if err := m.Matrix.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("matrix")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KialiMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KialiMetric) UnmarshalBinary(b []byte) error {
	var res KialiMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
