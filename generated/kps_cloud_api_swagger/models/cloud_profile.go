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

// CloudProfile CloudProfile is the object model for cloud credentials.
// swagger:model CloudProfile
type CloudProfile struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Describes the cloud service profile.
	// Required: true
	Description *string `json:"description"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// ntnx:ignore
	//
	// Internal Flag - encrypted - for internal migration use
	IFlagEncrypted bool `json:"iflagEncrypted,omitempty"`

	// Name for the cloud profile.
	// Required: true
	Name *string `json:"name"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// Cloud type for this cloud profile.
	// Set value to one of the following: AWS, GCP, Azure
	// Required: true
	// Enum: [AWS GCP Azure]
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

	// aws credential
	AwsCredential *AWSCredential `json:"awsCredential,omitempty"`

	// az credential
	AzCredential *AZCredential `json:"azCredential,omitempty"`

	// gcp credential
	GcpCredential *GCPCredential `json:"gcpCredential,omitempty"`
}

// Validate validates this cloud profile
func (m *CloudProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
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

	if err := m.validateAwsCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudProfile) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudProfile) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CloudProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudProfile) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

var cloudProfileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP","Azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudProfileTypeTypePropEnum = append(cloudProfileTypeTypePropEnum, v)
	}
}

const (

	// CloudProfileTypeAWS captures enum value "AWS"
	CloudProfileTypeAWS string = "AWS"

	// CloudProfileTypeGCP captures enum value "GCP"
	CloudProfileTypeGCP string = "GCP"

	// CloudProfileTypeAzure captures enum value "Azure"
	CloudProfileTypeAzure string = "Azure"
)

// prop value enum
func (m *CloudProfile) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudProfileTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudProfile) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CloudProfile) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudProfile) validateAwsCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsCredential) { // not required
		return nil
	}

	if m.AwsCredential != nil {
		if err := m.AwsCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsCredential")
			}
			return err
		}
	}

	return nil
}

func (m *CloudProfile) validateAzCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.AzCredential) { // not required
		return nil
	}

	if m.AzCredential != nil {
		if err := m.AzCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azCredential")
			}
			return err
		}
	}

	return nil
}

func (m *CloudProfile) validateGcpCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.GcpCredential) { // not required
		return nil
	}

	if m.GcpCredential != nil {
		if err := m.GcpCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpCredential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProfile) UnmarshalBinary(b []byte) error {
	var res CloudProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
