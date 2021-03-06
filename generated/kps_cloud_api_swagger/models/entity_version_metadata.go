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

// EntityVersionMetadata EntityVersionMetadata contains the database ID and UpdatedAt of an entity
// associated with an edge.
// EntityVersionMetadata is used with the get edge inventory delta payload
// to indicate the entity version, by entity ID, for each entity
// associated with an the edge.
// We use UpdatedAt as opposed to Version since
// Version can get truncated during JSON conversion.
// swagger:model EntityVersionMetadata
type EntityVersionMetadata struct {

	// ID
	ID string `json:"id,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this entity version metadata
func (m *EntityVersionMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityVersionMetadata) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityVersionMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityVersionMetadata) UnmarshalBinary(b []byte) error {
	var res EntityVersionMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
