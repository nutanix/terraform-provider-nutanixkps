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

// Service Service is object model for service
// swagger:model Service
type Service struct {

	// service type
	// Required: true
	// Enum: [IoT PaaS]
	ServiceType *string `json:"serviceType"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IoT","PaaS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceTypeServiceTypePropEnum = append(serviceTypeServiceTypePropEnum, v)
	}
}

const (

	// ServiceServiceTypeIoT captures enum value "IoT"
	ServiceServiceTypeIoT string = "IoT"

	// ServiceServiceTypePaaS captures enum value "PaaS"
	ServiceServiceTypePaaS string = "PaaS"
)

// prop value enum
func (m *Service) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceTypeServiceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Service) validateServiceType(formats strfmt.Registry) error {

	if err := validate.Required("serviceType", "body", m.ServiceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateServiceTypeEnum("serviceType", "body", *m.ServiceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Service) UnmarshalBinary(b []byte) error {
	var res Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
