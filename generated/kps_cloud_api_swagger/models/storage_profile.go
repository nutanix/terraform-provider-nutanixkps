// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageProfile StorageProfile is the object model for storage profile.
// swagger:model StorageProfile
type StorageProfile struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// ntnx:ignore
	//
	// Internal Flag - encrypted - for internal migration use
	IFlagEncrypted bool `json:"iflagEncrypted,omitempty"`

	// Flag to specify if it is default storage profile
	IsDefault bool `json:"isDefault,omitempty"`

	// Name for the storage profile.
	// Required: true
	Name *string `json:"name"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// Storage type for this Storage profile.
	// Required: true
	// Enum: [NutanixVolumes EBS vSphere]
	Type *string `json:"type"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// ntnx:ignore
	// Version of entity, implemented using timestamp in nano seconds
	// This is set to float64 since JSON numbers are floating point
	// May lose precision due to truncation but should have milli-second precision
	Version float64 `json:"version,omitempty"`

	// ebs storage config
	EbsStorageConfig *EBSStorageProfileConfig `json:"ebsStorageConfig,omitempty"`

	// nutanix volumes config
	NutanixVolumesConfig *NutanixVolumesStorageProfileConfig `json:"nutanixVolumesConfig,omitempty"`

	// v sphere storage config
	VSphereStorageConfig VSphereStorageProfileConfig `json:"vSphereStorageConfig,omitempty"`
}

// Validate validates this storage profile
func (m *StorageProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEbsStorageConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNutanixVolumesConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageProfile) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfile) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

var storageProfileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NutanixVolumes","EBS","vSphere"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageProfileTypeTypePropEnum = append(storageProfileTypeTypePropEnum, v)
	}
}

const (

	// StorageProfileTypeNutanixVolumes captures enum value "NutanixVolumes"
	StorageProfileTypeNutanixVolumes string = "NutanixVolumes"

	// StorageProfileTypeEBS captures enum value "EBS"
	StorageProfileTypeEBS string = "EBS"

	// StorageProfileTypeVSphere captures enum value "vSphere"
	StorageProfileTypeVSphere string = "vSphere"
)

// prop value enum
func (m *StorageProfile) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageProfileTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageProfile) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfile) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfile) validateEbsStorageConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.EbsStorageConfig) { // not required
		return nil
	}

	if m.EbsStorageConfig != nil {
		if err := m.EbsStorageConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ebsStorageConfig")
			}
			return err
		}
	}

	return nil
}

func (m *StorageProfile) validateNutanixVolumesConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NutanixVolumesConfig) { // not required
		return nil
	}

	if m.NutanixVolumesConfig != nil {
		if err := m.NutanixVolumesConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nutanixVolumesConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageProfile) UnmarshalBinary(b []byte) error {
	var res StorageProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
