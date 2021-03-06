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

// RuntimeEnvironmentListPayload payload for RuntimeEnvironmentListResponse
// swagger:model RuntimeEnvironmentListPayload
type RuntimeEnvironmentListPayload struct {

	// Specify result order. Zero or more entries with format: &ltkey> [desc]
	// where orderByKeys lists allowed keys in each response.
	OrderBy string `json:"orderBy,omitempty"`

	// Keys that can be used in orderBy.
	OrderByKeys []string `json:"orderByKeys"`

	// 0-based index of the page to fetch results.
	// Required: true
	PageIndex *int64 `json:"pageIndex"`

	// Item count of each page.
	// Required: true
	PageSize *int64 `json:"pageSize"`

	// list of runtime environments
	// Required: true
	RuntimeEnvironmentList []*RuntimeEnvironment `json:"result"`

	// Count of all items matching the query.
	// Required: true
	TotalCount *int64 `json:"totalCount"`
}

// Validate validates this runtime environment list payload
func (m *RuntimeEnvironmentListPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeEnvironmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeEnvironmentListPayload) validatePageIndex(formats strfmt.Registry) error {

	if err := validate.Required("pageIndex", "body", m.PageIndex); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironmentListPayload) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeEnvironmentListPayload) validateRuntimeEnvironmentList(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.RuntimeEnvironmentList); err != nil {
		return err
	}

	for i := 0; i < len(m.RuntimeEnvironmentList); i++ {
		if swag.IsZero(m.RuntimeEnvironmentList[i]) { // not required
			continue
		}

		if m.RuntimeEnvironmentList[i] != nil {
			if err := m.RuntimeEnvironmentList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuntimeEnvironmentListPayload) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeEnvironmentListPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeEnvironmentListPayload) UnmarshalBinary(b []byte) error {
	var res RuntimeEnvironmentListPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
