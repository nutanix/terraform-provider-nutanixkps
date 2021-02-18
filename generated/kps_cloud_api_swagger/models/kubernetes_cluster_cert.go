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

// KubernetesClusterCert KubernetesClusterCert is DB and object model
// swagger:model KubernetesClusterCert
type KubernetesClusterCert struct {

	// Root CA certificate for the tenant.
	// Required: true
	CACertificate *string `json:"CACertificate"`

	// Certificate for the kubernetes cluster using old/fixed root CA.
	// Required: true
	Certificate *string `json:"certificate"`

	// ntnx:ignore
	// Certificate for mqtt client on the edge
	// Required: true
	ClientCertificate *string `json:"clientCertificate"`

	// ntnx:ignore
	// Encrypted private key corresponding to the client certificate.
	// Required: true
	ClientPrivateKey *string `json:"clientPrivateKey"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Certificate for the kubernetes cluster using per-tenant root CA.
	// Required: true
	KubernetesClusterCertificate *string `json:"kubernetesClusterCertificate"`

	// ID of the kubernetes cluster this entity belongs to
	// Required: true
	KubernetesClusterID *string `json:"kubernetesClusterID"`

	// Encrypted private key using per-tenant root CA.
	// Required: true
	KubernetesClusterPrivateKey *string `json:"kubernetesClusterPrivateKey"`

	// locked
	Locked bool `json:"locked,omitempty"`

	// Encrypted private key using old/fixed root CA.
	// Required: true
	PrivateKey *string `json:"privateKey"`

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

// Validate validates this kubernetes cluster cert
func (m *KubernetesClusterCert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCACertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientPrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesClusterCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesClusterPrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
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

func (m *KubernetesClusterCert) validateCACertificate(formats strfmt.Registry) error {

	if err := validate.Required("CACertificate", "body", m.CACertificate); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateClientCertificate(formats strfmt.Registry) error {

	if err := validate.Required("clientCertificate", "body", m.ClientCertificate); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateClientPrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("clientPrivateKey", "body", m.ClientPrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateKubernetesClusterCertificate(formats strfmt.Registry) error {

	if err := validate.Required("kubernetesClusterCertificate", "body", m.KubernetesClusterCertificate); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateKubernetesClusterID(formats strfmt.Registry) error {

	if err := validate.Required("kubernetesClusterID", "body", m.KubernetesClusterID); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateKubernetesClusterPrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("kubernetesClusterPrivateKey", "body", m.KubernetesClusterPrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterCert) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesClusterCert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesClusterCert) UnmarshalBinary(b []byte) error {
	var res KubernetesClusterCert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
