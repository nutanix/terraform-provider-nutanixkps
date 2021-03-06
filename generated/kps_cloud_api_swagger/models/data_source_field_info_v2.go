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

// DataSourceFieldInfoV2 DataSourceFieldInfoV2 - data source field info struct
//
// A field represents specific information within a data source topic payload.
// A topic payload may contain multiple fields.
// The user defines fields extractable from a topic
// by specifying DataSourceFieldInfo for each field of a data source in the user interface.
// The fieldType for a field is used to extract the field value from the topic payload.
// swagger:model DataSourceFieldInfoV2
type DataSourceFieldInfoV2 struct {

	// Name of the field.
	// A unique name within the the data source.
	// Required: true
	Name *string `json:"name"`

	// topic for the field
	// The topic specified depends on the protocol in the data source. Specify the mqqtTopic for the MQTT protocol.
	// For the RTSP protocol, the topic is the server endpoint or named protocol stream in the RSTP URL.
	// Required: true
	Topic *string `json:"topic"`
}

// Validate validates this data source field info v2
func (m *DataSourceFieldInfoV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSourceFieldInfoV2) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DataSourceFieldInfoV2) validateTopic(formats strfmt.Registry) error {

	if err := validate.Required("topic", "body", m.Topic); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSourceFieldInfoV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSourceFieldInfoV2) UnmarshalBinary(b []byte) error {
	var res DataSourceFieldInfoV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
