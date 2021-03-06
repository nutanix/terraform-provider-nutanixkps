// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataPipelineContainers DataPipelineContainers encapsulates the container names
// for a specific data pipeline on a specific edge.
// swagger:model DataPipelineContainers
type DataPipelineContainers struct {

	// container names
	ContainerNames []string `json:"containerNames"`

	// data pipeline ID
	DataPipelineID string `json:"dataPipelineId,omitempty"`

	// edge ID
	EdgeID string `json:"edgeId,omitempty"`
}

// Validate validates this data pipeline containers
func (m *DataPipelineContainers) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataPipelineContainers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPipelineContainers) UnmarshalBinary(b []byte) error {
	var res DataPipelineContainers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
