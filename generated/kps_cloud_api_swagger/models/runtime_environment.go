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

// RuntimeEnvironment RuntimeEnvironment is the DB object and object model for function runtime.
//
// A RuntimeEnvironment is a Docker container runtime for functions.
// Karbon Platform Services includes several RuntimeEnvironments for built-in (and user defined) functions.
// User can also create custom RuntimeEnvironments which may be
// derived from Karbon Platform Services built-in RuntimeEnvironments.
// swagger:model RuntimeEnvironment
type RuntimeEnvironment struct {

	// ntnx:ignore
	//
	// Whether this is a built-in script runtime. Always set this to false for user created script runtime.
	// Required: true
	Builtin *bool `json:"builtin"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Runtime description.
	Description string `json:"description,omitempty"`

	// DockerProfile ID (Container registry profile) used by this script runtime.
	DockerProfileID string `json:"dockerProfileID,omitempty"`

	// Docker repository URI of the script runtime.
	DockerRepoURI string `json:"dockerRepoURI,omitempty"`

	// Dockerfile for the script runtime. Serves as documentation for the script runtime.
	Dockerfile string `json:"dockerfile,omitempty"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Runtime enviroment script language.
	// Required: true
	Language *string `json:"language"`

	// Name of the runtime environment.
	// Required: true
	Name *string `json:"name"`

	// ID of parent project, required for custom (non-built-in) script runtimes.
	ProjectID string `json:"projectId,omitempty"`

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

// Validate validates this runtime environment
func (m *RuntimeEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuiltin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *RuntimeEnvironment) validateBuiltin(formats strfmt.Registry) error {

	if err := validate.Required("builtin", "body", m.Builtin); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironment) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironment) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironment) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeEnvironment) UnmarshalBinary(b []byte) error {
	var res RuntimeEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
