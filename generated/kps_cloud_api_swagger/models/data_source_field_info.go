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

// DataSourceFieldInfo DataSourceFieldInfo - data source field info struct
//
// A field represents specific information within a data source topic payload.
// A topic payload may contain multiple fields.
// The user defines fields extractable from a topic
// by specifying DataSourceFieldInfo for each field of a data source in the user interface.
// The fieldType for a field is used to extract the field value from the topic payload.
// swagger:model DataSourceFieldInfo
type DataSourceFieldInfo struct {

	// Data type for the field.
	// For example, Temperature, Pressure, Custom, and so on.
	// Specify Custom for the entire topic payload. No special extraction is performed.
	// When you specify Custom, Karbon Platform Services might not perform intelligent operations automatically
	// when you specify other fields like Temperature.
	// In the future custom extraction functions
	// for each field might be allowed.
	// DataSource dataType is derived from fieldType of all fields in the data source.
	// Required: true
	FieldType *string `json:"fieldType"`

	// Topic for the field.
	// The topic specified depends on the protocol in the data source. Specify the mqqtTopic for the MQTT protocol.
	// For the RTSP protocol, the topic is the server endpoint or named protocol stream in the RSTP URL.
	// Required: true
	MQTTTopic *string `json:"mqttTopic"`

	// A unique name within the the data source.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this data source field info
func (m *DataSourceFieldInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMQTTTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSourceFieldInfo) validateFieldType(formats strfmt.Registry) error {

	if err := validate.Required("fieldType", "body", m.FieldType); err != nil {
		return err
	}

	return nil
}

func (m *DataSourceFieldInfo) validateMQTTTopic(formats strfmt.Registry) error {

	if err := validate.Required("mqttTopic", "body", m.MQTTTopic); err != nil {
		return err
	}

	return nil
}

func (m *DataSourceFieldInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSourceFieldInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSourceFieldInfo) UnmarshalBinary(b []byte) error {
	var res DataSourceFieldInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
