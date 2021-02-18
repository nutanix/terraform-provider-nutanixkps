// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EdgeInventoryDeltaResponse EdgeInventoryDeltaResponse is the response for get edge inventory delta. It
// consists of Deleted, Created, and Updated fields of type
// EdgeInventoryDeleted, EdgeInventoryDetails and EdgeInventoryDetails
// described above.
// swagger:model EdgeInventoryDeltaResponse
type EdgeInventoryDeltaResponse struct {

	// created
	Created *EdgeInventoryDetails `json:"Created,omitempty"`

	// deleted
	Deleted *EdgeInventoryDeleted `json:"Deleted,omitempty"`

	// updated
	Updated *EdgeInventoryDetails `json:"Updated,omitempty"`
}

// Validate validates this edge inventory delta response
func (m *EdgeInventoryDeltaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeInventoryDeltaResponse) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if m.Created != nil {
		if err := m.Created.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Created")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeInventoryDeltaResponse) validateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.Deleted) { // not required
		return nil
	}

	if m.Deleted != nil {
		if err := m.Deleted.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Deleted")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeInventoryDeltaResponse) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if m.Updated != nil {
		if err := m.Updated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Updated")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeInventoryDeltaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeInventoryDeltaResponse) UnmarshalBinary(b []byte) error {
	var res EdgeInventoryDeltaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
