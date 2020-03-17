// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OsPlaceHolder Os:Place Holder
//
// Definition for place holders in templates/post install scripts.
//
// swagger:model osPlaceHolder
type OsPlaceHolder struct {
	MoBaseComplexType

	OsPlaceHolderAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OsPlaceHolder) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 OsPlaceHolderAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.OsPlaceHolderAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OsPlaceHolder) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.OsPlaceHolderAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this os place holder
func (m *OsPlaceHolder) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OsPlaceHolderAO1P1
	if err := m.OsPlaceHolderAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OsPlaceHolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsPlaceHolder) UnmarshalBinary(b []byte) error {
	var res OsPlaceHolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OsPlaceHolderAO1P1 os place holder a o1 p1
//
// swagger:model OsPlaceHolderAO1P1
type OsPlaceHolderAO1P1 struct {

	// Flag to indicate if value is set. Value will be used to check if any edit.
	IsValueSet *bool `json:"IsValueSet,omitempty"`

	// Definition of place holder.
	Type *WorkflowPrimitiveDataType `json:"Type,omitempty"`

	// Value for placeholder provided by user.
	Value interface{} `json:"Value,omitempty"`

	// os place holder a o1 p1
	OsPlaceHolderAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *OsPlaceHolderAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Flag to indicate if value is set. Value will be used to check if any edit.
		IsValueSet *bool `json:"IsValueSet,omitempty"`

		// Definition of place holder.
		Type *WorkflowPrimitiveDataType `json:"Type,omitempty"`

		// Value for placeholder provided by user.
		Value interface{} `json:"Value,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv OsPlaceHolderAO1P1

	rcv.IsValueSet = stage1.IsValueSet
	rcv.Type = stage1.Type
	rcv.Value = stage1.Value
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "IsValueSet")
	delete(stage2, "Type")
	delete(stage2, "Value")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.OsPlaceHolderAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m OsPlaceHolderAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Flag to indicate if value is set. Value will be used to check if any edit.
		IsValueSet *bool `json:"IsValueSet,omitempty"`

		// Definition of place holder.
		Type *WorkflowPrimitiveDataType `json:"Type,omitempty"`

		// Value for placeholder provided by user.
		Value interface{} `json:"Value,omitempty"`
	}

	stage1.IsValueSet = m.IsValueSet
	stage1.Type = m.Type
	stage1.Value = m.Value

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.OsPlaceHolderAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.OsPlaceHolderAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this os place holder a o1 p1
func (m *OsPlaceHolderAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsPlaceHolderAO1P1) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OsPlaceHolderAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsPlaceHolderAO1P1) UnmarshalBinary(b []byte) error {
	var res OsPlaceHolderAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
