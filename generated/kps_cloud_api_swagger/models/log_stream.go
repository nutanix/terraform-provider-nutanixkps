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

// LogStream LogStream is object model for requesting log streams.
//
// LogStream is used to specify the edge, app or pipeline, and container
// to stream logs from.
// swagger:model LogStream
type LogStream struct {

	// ID of the application
	ApplicationID string `json:"applicationId,omitempty"`

	// Name of the kubernetes container in the pod to
	// stream logs from.
	// Required: true
	Container *string `json:"container"`

	// ID of the data pipeline
	DataPipelineID string `json:"dataPipelineId,omitempty"`

	// Edge ID from which logs will be streamed.
	// Required: true
	EdgeID *string `json:"edgeId"`
}

// Validate validates this log stream
func (m *LogStream) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogStream) validateContainer(formats strfmt.Registry) error {

	if err := validate.Required("container", "body", m.Container); err != nil {
		return err
	}

	return nil
}

func (m *LogStream) validateEdgeID(formats strfmt.Registry) error {

	if err := validate.Required("edgeId", "body", m.EdgeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogStream) UnmarshalBinary(b []byte) error {
	var res LogStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
