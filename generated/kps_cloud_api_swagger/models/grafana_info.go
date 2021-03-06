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

// GrafanaInfo GrafanaInfo provides information to access Grafana dashboards
// swagger:model GrafanaInfo
type GrafanaInfo struct {

	// external links
	ExternalLinks []*ExternalLink `json:"externalLinks"`
}

// Validate validates this grafana info
func (m *GrafanaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrafanaInfo) validateExternalLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalLinks) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalLinks); i++ {
		if swag.IsZero(m.ExternalLinks[i]) { // not required
			continue
		}

		if m.ExternalLinks[i] != nil {
			if err := m.ExternalLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalLinks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GrafanaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrafanaInfo) UnmarshalBinary(b []byte) error {
	var res GrafanaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
