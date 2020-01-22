// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MemoryPersistentMemoryNamespace Memory:Persistent Memory Namespace
//
// This represents a Persistent Memory Namespace configured within the persistent memory region on a server.
//
// swagger:model memoryPersistentMemoryNamespace
type MemoryPersistentMemoryNamespace struct {
	InventoryBase

	// This represents the capacity in GB of a Persistent Memory Namespace.
	//
	// Read Only: true
	Capacity string `json:"Capacity,omitempty"`

	// This represents the health state of a Persistent Memory Namespace.
	//
	// Read Only: true
	HealthState string `json:"HealthState,omitempty"`

	// This represents the label version of a Persistent Memory Namespace.
	//
	// Read Only: true
	LabelVersion string `json:"LabelVersion,omitempty"`

	// A collection of references to the [memory.PersistentMemoryRegion](mo://memory.PersistentMemoryRegion) Managed Object.
	//
	// When this managed object is deleted, the referenced [memory.PersistentMemoryRegion](mo://memory.PersistentMemoryRegion) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	MemoryPersistentMemoryRegion *MemoryPersistentMemoryRegionRef `json:"MemoryPersistentMemoryRegion,omitempty"`

	// This represents the mode of a Persistent Memory Namespace.
	//
	// Read Only: true
	Mode string `json:"Mode,omitempty"`

	// This represents the name of a Persistent Memory Namespace.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// This represents the uuid of a Persistent Memory Namespace.
	//
	// Read Only: true
	UUID string `json:"Uuid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemoryPersistentMemoryNamespace) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		Capacity string `json:"Capacity,omitempty"`

		HealthState string `json:"HealthState,omitempty"`

		LabelVersion string `json:"LabelVersion,omitempty"`

		MemoryPersistentMemoryRegion *MemoryPersistentMemoryRegionRef `json:"MemoryPersistentMemoryRegion,omitempty"`

		Mode string `json:"Mode,omitempty"`

		Name string `json:"Name,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		UUID string `json:"Uuid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Capacity = dataAO1.Capacity

	m.HealthState = dataAO1.HealthState

	m.LabelVersion = dataAO1.LabelVersion

	m.MemoryPersistentMemoryRegion = dataAO1.MemoryPersistentMemoryRegion

	m.Mode = dataAO1.Mode

	m.Name = dataAO1.Name

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.UUID = dataAO1.UUID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemoryPersistentMemoryNamespace) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Capacity string `json:"Capacity,omitempty"`

		HealthState string `json:"HealthState,omitempty"`

		LabelVersion string `json:"LabelVersion,omitempty"`

		MemoryPersistentMemoryRegion *MemoryPersistentMemoryRegionRef `json:"MemoryPersistentMemoryRegion,omitempty"`

		Mode string `json:"Mode,omitempty"`

		Name string `json:"Name,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		UUID string `json:"Uuid,omitempty"`
	}

	dataAO1.Capacity = m.Capacity

	dataAO1.HealthState = m.HealthState

	dataAO1.LabelVersion = m.LabelVersion

	dataAO1.MemoryPersistentMemoryRegion = m.MemoryPersistentMemoryRegion

	dataAO1.Mode = m.Mode

	dataAO1.Name = m.Name

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.UUID = m.UUID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this memory persistent memory namespace
func (m *MemoryPersistentMemoryNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryPersistentMemoryRegion(formats); err != nil {
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

func (m *MemoryPersistentMemoryNamespace) validateMemoryPersistentMemoryRegion(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryPersistentMemoryRegion) { // not required
		return nil
	}

	if m.MemoryPersistentMemoryRegion != nil {
		if err := m.MemoryPersistentMemoryRegion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MemoryPersistentMemoryRegion")
			}
			return err
		}
	}

	return nil
}

func (m *MemoryPersistentMemoryNamespace) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *MemoryPersistentMemoryNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoryPersistentMemoryNamespace) UnmarshalBinary(b []byte) error {
	var res MemoryPersistentMemoryNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}