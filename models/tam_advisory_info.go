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

// TamAdvisoryInfo Tam:Advisory Info
//
// State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.
//
// swagger:model tamAdvisoryInfo
type TamAdvisoryInfo struct {
	MoBaseMo

	// Reference to the source Intersight advisory.
	//
	Advisory *TamAdvisoryRef `json:"Advisory,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory.
	//
	// Enum: [active acknowledged]
	State *string `json:"State,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamAdvisoryInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Advisory *TamAdvisoryRef `json:"Advisory,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		State *string `json:"State,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Advisory = dataAO1.Advisory

	m.Organization = dataAO1.Organization

	m.State = dataAO1.State

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamAdvisoryInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Advisory *TamAdvisoryRef `json:"Advisory,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		State *string `json:"State,omitempty"`
	}

	dataAO1.Advisory = m.Advisory

	dataAO1.Organization = m.Organization

	dataAO1.State = m.State

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam advisory info
func (m *TamAdvisoryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvisory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TamAdvisoryInfo) validateAdvisory(formats strfmt.Registry) error {

	if swag.IsZero(m.Advisory) { // not required
		return nil
	}

	if m.Advisory != nil {
		if err := m.Advisory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Advisory")
			}
			return err
		}
	}

	return nil
}

func (m *TamAdvisoryInfo) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

var tamAdvisoryInfoTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","acknowledged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tamAdvisoryInfoTypeStatePropEnum = append(tamAdvisoryInfoTypeStatePropEnum, v)
	}
}

// property enum
func (m *TamAdvisoryInfo) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tamAdvisoryInfoTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TamAdvisoryInfo) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("State", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TamAdvisoryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamAdvisoryInfo) UnmarshalBinary(b []byte) error {
	var res TamAdvisoryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}