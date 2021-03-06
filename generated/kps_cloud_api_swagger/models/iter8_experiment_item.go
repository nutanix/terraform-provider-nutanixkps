// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Iter8ExperimentItem iter8 experiment item
// swagger:model Iter8ExperimentItem
type Iter8ExperimentItem struct {

	// assessment conclusion
	AssessmentConclusion []string `json:"assessmentConclusion"`

	// baseline
	Baseline string `json:"baseline,omitempty"`

	// baseline percentage
	BaselinePercentage int64 `json:"baselinePercentage,omitempty"`

	// candidate
	Candidate string `json:"candidate,omitempty"`

	// candidate percentage
	CandidatePercentage int64 `json:"candidatePercentage,omitempty"`

	// created at
	CreatedAt int64 `json:"createdAt,omitempty"`

	// ended at
	EndedAt int64 `json:"endedAt,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`

	// started at
	StartedAt int64 `json:"startedAt,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// target service
	TargetService string `json:"targetService,omitempty"`

	// target service namespace
	TargetServiceNamespace string `json:"targetServiceNamespace,omitempty"`
}

// Validate validates this iter8 experiment item
func (m *Iter8ExperimentItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Iter8ExperimentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Iter8ExperimentItem) UnmarshalBinary(b []byte) error {
	var res Iter8ExperimentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
