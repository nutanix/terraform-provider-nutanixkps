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

// MLModelMetadata MLModelMetadata is base object model for the machine learning model.
//
// It serves as the payload when creating a machine learning model.
// swagger:model MLModelMetadata
type MLModelMetadata struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Describe the machine learning model.  Maximum length is 200 characters.
	// Required: true
	Description *string `json:"description"`

	// Machine learning model framework type.
	// Required: true
	// Enum: [TensorFlow 1.13.1 OpenVINO 2019_R2 TensorFlow 2.1.0]
	FrameworkType *string `json:"frameworkType"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Name for the machine learning model. Maximum length is 200 characters.
	// Required: true
	Name *string `json:"name"`

	// Parent project ID associated with this machine learning model.
	// Required: true
	ProjectID *string `json:"projectId"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// ntnx:ignore
	// Version of entity, implemented using timestamp in nano seconds
	// This is set to float64 since JSON numbers are floating point
	// May lose precision due to truncation but should have milli-second precision
	Version float64 `json:"version,omitempty"`
}

// Validate validates this m l model metadata
func (m *MLModelMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrameworkType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MLModelMetadata) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MLModelMetadata) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var mLModelMetadataTypeFrameworkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TensorFlow 1.13.1","OpenVINO 2019_R2","TensorFlow 2.1.0"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mLModelMetadataTypeFrameworkTypePropEnum = append(mLModelMetadataTypeFrameworkTypePropEnum, v)
	}
}

const (

	// MLModelMetadataFrameworkTypeTensorFlow1131 captures enum value "TensorFlow 1.13.1"
	MLModelMetadataFrameworkTypeTensorFlow1131 string = "TensorFlow 1.13.1"

	// MLModelMetadataFrameworkTypeOpenVINO2019R2 captures enum value "OpenVINO 2019_R2"
	MLModelMetadataFrameworkTypeOpenVINO2019R2 string = "OpenVINO 2019_R2"

	// MLModelMetadataFrameworkTypeTensorFlow210 captures enum value "TensorFlow 2.1.0"
	MLModelMetadataFrameworkTypeTensorFlow210 string = "TensorFlow 2.1.0"
)

// prop value enum
func (m *MLModelMetadata) validateFrameworkTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mLModelMetadataTypeFrameworkTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MLModelMetadata) validateFrameworkType(formats strfmt.Registry) error {

	if err := validate.Required("frameworkType", "body", m.FrameworkType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFrameworkTypeEnum("frameworkType", "body", *m.FrameworkType); err != nil {
		return err
	}

	return nil
}

func (m *MLModelMetadata) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MLModelMetadata) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *MLModelMetadata) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *MLModelMetadata) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MLModelMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MLModelMetadata) UnmarshalBinary(b []byte) error {
	var res MLModelMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
