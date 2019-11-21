// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MemoryUnit Memory:Unit
//
// Represents a DIMM on a server.
//
// swagger:model memoryUnit
type MemoryUnit struct {
	EquipmentBase

	// admin state
	// Read Only: true
	AdminState string `json:"AdminState,omitempty"`

	// array Id
	// Read Only: true
	ArrayID int64 `json:"ArrayId,omitempty"`

	// bank
	// Read Only: true
	Bank int64 `json:"Bank,omitempty"`

	// capacity
	// Read Only: true
	Capacity string `json:"Capacity,omitempty"`

	// clock
	// Read Only: true
	Clock string `json:"Clock,omitempty"`

	// form factor
	// Read Only: true
	FormFactor string `json:"FormFactor,omitempty"`

	// latency
	// Read Only: true
	Latency string `json:"Latency,omitempty"`

	// location
	// Read Only: true
	Location string `json:"Location,omitempty"`

	// A collection of references to the [memory.Array](mo://memory.Array) Managed Object.
	//
	// When this managed object is deleted, the referenced [memory.Array](mo://memory.Array) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

	// memory Id
	// Read Only: true
	MemoryID int64 `json:"MemoryId,omitempty"`

	// oper power state
	// Read Only: true
	OperPowerState string `json:"OperPowerState,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// operability
	// Read Only: true
	Operability string `json:"Operability,omitempty"`

	// presence
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// set
	// Read Only: true
	Set int64 `json:"Set,omitempty"`

	// speed
	// Read Only: true
	Speed string `json:"Speed,omitempty"`

	// thermal
	// Read Only: true
	Thermal string `json:"Thermal,omitempty"`

	// type
	// Read Only: true
	Type string `json:"Type,omitempty"`

	// visibility
	// Read Only: true
	Visibility string `json:"Visibility,omitempty"`

	// width
	// Read Only: true
	Width string `json:"Width,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemoryUnit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		AdminState string `json:"AdminState,omitempty"`

		ArrayID int64 `json:"ArrayId,omitempty"`

		Bank int64 `json:"Bank,omitempty"`

		Capacity string `json:"Capacity,omitempty"`

		Clock string `json:"Clock,omitempty"`

		FormFactor string `json:"FormFactor,omitempty"`

		Latency string `json:"Latency,omitempty"`

		Location string `json:"Location,omitempty"`

		MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

		MemoryID int64 `json:"MemoryId,omitempty"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Set int64 `json:"Set,omitempty"`

		Speed string `json:"Speed,omitempty"`

		Thermal string `json:"Thermal,omitempty"`

		Type string `json:"Type,omitempty"`

		Visibility string `json:"Visibility,omitempty"`

		Width string `json:"Width,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdminState = dataAO1.AdminState

	m.ArrayID = dataAO1.ArrayID

	m.Bank = dataAO1.Bank

	m.Capacity = dataAO1.Capacity

	m.Clock = dataAO1.Clock

	m.FormFactor = dataAO1.FormFactor

	m.Latency = dataAO1.Latency

	m.Location = dataAO1.Location

	m.MemoryArray = dataAO1.MemoryArray

	m.MemoryID = dataAO1.MemoryID

	m.OperPowerState = dataAO1.OperPowerState

	m.OperState = dataAO1.OperState

	m.Operability = dataAO1.Operability

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Set = dataAO1.Set

	m.Speed = dataAO1.Speed

	m.Thermal = dataAO1.Thermal

	m.Type = dataAO1.Type

	m.Visibility = dataAO1.Visibility

	m.Width = dataAO1.Width

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemoryUnit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AdminState string `json:"AdminState,omitempty"`

		ArrayID int64 `json:"ArrayId,omitempty"`

		Bank int64 `json:"Bank,omitempty"`

		Capacity string `json:"Capacity,omitempty"`

		Clock string `json:"Clock,omitempty"`

		FormFactor string `json:"FormFactor,omitempty"`

		Latency string `json:"Latency,omitempty"`

		Location string `json:"Location,omitempty"`

		MemoryArray *MemoryArrayRef `json:"MemoryArray,omitempty"`

		MemoryID int64 `json:"MemoryId,omitempty"`

		OperPowerState string `json:"OperPowerState,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Set int64 `json:"Set,omitempty"`

		Speed string `json:"Speed,omitempty"`

		Thermal string `json:"Thermal,omitempty"`

		Type string `json:"Type,omitempty"`

		Visibility string `json:"Visibility,omitempty"`

		Width string `json:"Width,omitempty"`
	}

	dataAO1.AdminState = m.AdminState

	dataAO1.ArrayID = m.ArrayID

	dataAO1.Bank = m.Bank

	dataAO1.Capacity = m.Capacity

	dataAO1.Clock = m.Clock

	dataAO1.FormFactor = m.FormFactor

	dataAO1.Latency = m.Latency

	dataAO1.Location = m.Location

	dataAO1.MemoryArray = m.MemoryArray

	dataAO1.MemoryID = m.MemoryID

	dataAO1.OperPowerState = m.OperPowerState

	dataAO1.OperState = m.OperState

	dataAO1.Operability = m.Operability

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Set = m.Set

	dataAO1.Speed = m.Speed

	dataAO1.Thermal = m.Thermal

	dataAO1.Type = m.Type

	dataAO1.Visibility = m.Visibility

	dataAO1.Width = m.Width

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this memory unit
func (m *MemoryUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryArray(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemoryUnit) validateMemoryArray(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryArray) { // not required
		return nil
	}

	if m.MemoryArray != nil {
		if err := m.MemoryArray.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MemoryArray")
			}
			return err
		}
	}

	return nil
}

func (m *MemoryUnit) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemoryUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryUnit) UnmarshalBinary(b []byte) error {
	var res MemoryUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
