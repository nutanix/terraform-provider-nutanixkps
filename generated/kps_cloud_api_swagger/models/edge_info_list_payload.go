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

// EdgeInfoListPayload payload for EdgeInfoListResponseV2
// swagger:model EdgeInfoListPayload
type EdgeInfoListPayload struct {

	// list of edge usage info
	// Required: true
	EdgeUsageInfoList []*EdgeUsageInfo `json:"result"`

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

	// Count of all items matching the query.
	// Required: true
	TotalCount *int64 `json:"totalCount"`
}

// Validate validates this edge info list payload
func (m *EdgeInfoListPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeUsageInfoList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
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

func (m *EdgeInfoListPayload) validateEdgeUsageInfoList(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.EdgeUsageInfoList); err != nil {
		return err
	}

	for i := 0; i < len(m.EdgeUsageInfoList); i++ {
		if swag.IsZero(m.EdgeUsageInfoList[i]) { // not required
			continue
		}

		if m.EdgeUsageInfoList[i] != nil {
			if err := m.EdgeUsageInfoList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeInfoListPayload) validatePageIndex(formats strfmt.Registry) error {

	if err := validate.Required("pageIndex", "body", m.PageIndex); err != nil {
		return err
	}

	return nil
}

func (m *EdgeInfoListPayload) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *EdgeInfoListPayload) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeInfoListPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeInfoListPayload) UnmarshalBinary(b []byte) error {
	var res EdgeInfoListPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
