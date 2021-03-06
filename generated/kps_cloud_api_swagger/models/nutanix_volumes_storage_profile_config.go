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

// NutanixVolumesStorageProfileConfig NutanixVolumesStorageProfileConfig - struct for Nutanix volume Storage Profile config.
// swagger:model NutanixVolumesStorageProfileConfig
type NutanixVolumesStorageProfileConfig struct {

	// data services IP
	// Required: true
	DataServicesIP *string `json:"dataServicesIP"`

	// data services port
	// Required: true
	DataServicesPort *int64 `json:"dataServicesPort"`

	// flash mode
	FlashMode bool `json:"flashMode,omitempty"`

	// prism element cluster port
	// Required: true
	PrismElementClusterPort *int64 `json:"prismElementClusterPort"`

	// prism element cluster v IP
	// Required: true
	PrismElementClusterVIP *string `json:"prismElementClusterVIP"`

	// prism element password
	// Required: true
	PrismElementPassword *string `json:"prismElementPassword"`

	// prism element user name
	// Required: true
	PrismElementUserName *string `json:"prismElementUserName"`

	// storage container name
	// Required: true
	StorageContainerName *string `json:"storageContainerName"`
}

// Validate validates this nutanix volumes storage profile config
func (m *NutanixVolumesStorageProfileConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataServicesIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataServicesPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrismElementClusterPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrismElementClusterVIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrismElementPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrismElementUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageContainerName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validateDataServicesIP(formats strfmt.Registry) error {

	if err := validate.Required("dataServicesIP", "body", m.DataServicesIP); err != nil {
		return err
	}

	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validateDataServicesPort(formats strfmt.Registry) error {

	if err := validate.Required("dataServicesPort", "body", m.DataServicesPort); err != nil {
		return err
	}

	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validatePrismElementClusterPort(formats strfmt.Registry) error {

	if err := validate.Required("prismElementClusterPort", "body", m.PrismElementClusterPort); err != nil {
		return err
	}

	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validatePrismElementClusterVIP(formats strfmt.Registry) error {

	if err := validate.Required("prismElementClusterVIP", "body", m.PrismElementClusterVIP); err != nil {
		return err
	}

	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validatePrismElementPassword(formats strfmt.Registry) error {

	if err := validate.Required("prismElementPassword", "body", m.PrismElementPassword); err != nil {
		return err
	}

	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validatePrismElementUserName(formats strfmt.Registry) error {

	if err := validate.Required("prismElementUserName", "body", m.PrismElementUserName); err != nil {
		return err
	}

	return nil
}

func (m *NutanixVolumesStorageProfileConfig) validateStorageContainerName(formats strfmt.Registry) error {

	if err := validate.Required("storageContainerName", "body", m.StorageContainerName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NutanixVolumesStorageProfileConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NutanixVolumesStorageProfileConfig) UnmarshalBinary(b []byte) error {
	var res NutanixVolumesStorageProfileConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
