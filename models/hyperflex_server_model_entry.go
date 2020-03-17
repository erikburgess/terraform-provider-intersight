// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HyperflexServerModelEntry Hyperflex:Server Model Entry
//
// A constraint specifying supported server models in regex format.
//
// swagger:model hyperflexServerModelEntry
type HyperflexServerModelEntry struct {
	HyperflexAbstractAppSetting

	// The conditions that must be satisfied before applying the AppSetting.
	Constraint *HyperflexAppSettingConstraint `json:"Constraint,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexServerModelEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexAbstractAppSetting
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexAbstractAppSetting = aO0

	// AO1
	var dataAO1 struct {
		Constraint *HyperflexAppSettingConstraint `json:"Constraint,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Constraint = dataAO1.Constraint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexServerModelEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.HyperflexAbstractAppSetting)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Constraint *HyperflexAppSettingConstraint `json:"Constraint,omitempty"`
	}

	dataAO1.Constraint = m.Constraint

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex server model entry
func (m *HyperflexServerModelEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexAbstractAppSetting
	if err := m.HyperflexAbstractAppSetting.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexServerModelEntry) validateConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraint) { // not required
		return nil
	}

	if m.Constraint != nil {
		if err := m.Constraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constraint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexServerModelEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexServerModelEntry) UnmarshalBinary(b []byte) error {
	var res HyperflexServerModelEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
