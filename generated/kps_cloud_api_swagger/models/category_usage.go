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

// CategoryUsage CategoryUsage captures usage for a category or category value
// swagger:model CategoryUsage
type CategoryUsage struct {

	// IDs of applications using this category
	// Required: true
	ApplicationIDs []string `json:"applicationIds"`

	// IDs of data pipelines using this category
	// Required: true
	DataPipelineIDs []string `json:"dataPipelineIds"`

	// IDs of data sources using this category
	// Required: true
	DataSourceIDs []string `json:"dataSourceIds"`

	// IDs of edges using this category
	// Required: true
	EdgeIDs []string `json:"edgeIds"`

	// IDs of projects using this category
	// Required: true
	ProjectIDs []string `json:"projectIds"`
}

// Validate validates this category usage
func (m *CategoryUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPipelineIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectIDs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CategoryUsage) validateApplicationIDs(formats strfmt.Registry) error {

	if err := validate.Required("applicationIds", "body", m.ApplicationIDs); err != nil {
		return err
	}

	return nil
}

func (m *CategoryUsage) validateDataPipelineIDs(formats strfmt.Registry) error {

	if err := validate.Required("dataPipelineIds", "body", m.DataPipelineIDs); err != nil {
		return err
	}

	return nil
}

func (m *CategoryUsage) validateDataSourceIDs(formats strfmt.Registry) error {

	if err := validate.Required("dataSourceIds", "body", m.DataSourceIDs); err != nil {
		return err
	}

	return nil
}

func (m *CategoryUsage) validateEdgeIDs(formats strfmt.Registry) error {

	if err := validate.Required("edgeIds", "body", m.EdgeIDs); err != nil {
		return err
	}

	return nil
}

func (m *CategoryUsage) validateProjectIDs(formats strfmt.Registry) error {

	if err := validate.Required("projectIds", "body", m.ProjectIDs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CategoryUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategoryUsage) UnmarshalBinary(b []byte) error {
	var res CategoryUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
