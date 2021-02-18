// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Edge Edge is the DB object and object model for edge
//
// An Edge is a Nutanix (Kubernetes) cluster for a tenant.
// swagger:model Edge
type Edge struct {

	// Determines if the edge is currently connected to XI IoT management services.
	Connected bool `json:"connected,omitempty"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Edge description
	Description string `json:"description,omitempty"`

	// Number of devices (nodes) in this edge
	// Required: true
	EdgeDevices *float64 `json:"edgeDevices"`

	// Edge Gateway IP address
	// Required: true
	Gateway *string `json:"gateway"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Edge IP Address
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// A list of Category labels for this edge.
	Labels []*CategoryInfo `json:"labels"`

	// Edge name.
	// Maximum length edge name is determined by kubernetes.
	// Name length limited to 60 as node name is the edge name plus a suffix.
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go
	// Required: true
	Name *string `json:"name"`

	// Edge serial number
	// Required: true
	SerialNumber *string `json:"serialNumber"`

	// ShortID is the unique ID for the given edge.
	// This ID must be unique for each edge, for the given tenant.
	ShortID string `json:"shortId,omitempty"`

	// Edge storage capacity in GB
	// Required: true
	StorageCapacity *float64 `json:"storageCapacity"`

	// Edge storage usage in GB
	// Required: true
	StorageUsage *float64 `json:"storageUsage"`

	// Edge subnet mask
	// Required: true
	Subnet *string `json:"subnet"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// Edge type.
	Type string `json:"type,omitempty"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// ntnx:ignore
	// Version of entity, implemented using timestamp in nano seconds
	// This is set to float64 since JSON numbers are floating point
	// May lose precision due to truncation but should have milli-second precision
	Version float64 `json:"version,omitempty"`

	// role
	Role *NodeRole `json:"role,omitempty"`
}

// Validate validates this edge
func (m *Edge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Edge) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateEdgeDevices(formats strfmt.Registry) error {

	if err := validate.Required("edgeDevices", "body", m.EdgeDevices); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Edge) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("serialNumber", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateStorageCapacity(formats strfmt.Registry) error {

	if err := validate.Required("storageCapacity", "body", m.StorageCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateStorageUsage(formats strfmt.Registry) error {

	if err := validate.Required("storageUsage", "body", m.StorageUsage); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Edge) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Edge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Edge) UnmarshalBinary(b []byte) error {
	var res Edge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
