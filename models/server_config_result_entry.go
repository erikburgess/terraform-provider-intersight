// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerConfigResultEntry Server:Config Result Entry
//
// The profile configuration (deploy, validation) results details information.
//
// swagger:model serverConfigResultEntry
type ServerConfigResultEntry struct {
	PolicyAbstractConfigResultEntry

	// A collection of references to the [server.ConfigResult](mo://server.ConfigResult) Managed Object.
	//
	// When this managed object is deleted, the referenced [server.ConfigResult](mo://server.ConfigResult) MO unsets its reference to this deleted MO.
	//
	ConfigResult *ServerConfigResultRef `json:"ConfigResult,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServerConfigResultEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractConfigResultEntry
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractConfigResultEntry = aO0

	// AO1
	var dataAO1 struct {
		ConfigResult *ServerConfigResultRef `json:"ConfigResult,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConfigResult = dataAO1.ConfigResult

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServerConfigResultEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractConfigResultEntry)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ConfigResult *ServerConfigResultRef `json:"ConfigResult,omitempty"`
	}

	dataAO1.ConfigResult = m.ConfigResult

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this server config result entry
func (m *ServerConfigResultEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractConfigResultEntry
	if err := m.PolicyAbstractConfigResultEntry.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerConfigResultEntry) validateConfigResult(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigResult) { // not required
		return nil
	}

	if m.ConfigResult != nil {
		if err := m.ConfigResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigResult")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerConfigResultEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerConfigResultEntry) UnmarshalBinary(b []byte) error {
	var res ServerConfigResultEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
