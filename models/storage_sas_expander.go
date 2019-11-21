// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageSasExpander Storage:Sas Expander
//
// SAS Expander present in a server.
//
// swagger:model storageSasExpander
type StorageSasExpander struct {
	EquipmentBase

	// controller
	// Read Only: true
	Controller *ManagementControllerRef `json:"Controller,omitempty"`

	// A collection of references to the [equipment.Chassis](mo://equipment.Chassis) Managed Object.
	//
	// When this managed object is deleted, the referenced [equipment.Chassis](mo://equipment.Chassis) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	EquipmentChassis *EquipmentChassisRef `json:"EquipmentChassis,omitempty"`

	// expander Id
	// Read Only: true
	ExpanderID int64 `json:"ExpanderId,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

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

	// sas address
	// Read Only: true
	SasAddress string `json:"SasAddress,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageSasExpander) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		Controller *ManagementControllerRef `json:"Controller,omitempty"`

		EquipmentChassis *EquipmentChassisRef `json:"EquipmentChassis,omitempty"`

		ExpanderID int64 `json:"ExpanderId,omitempty"`

		Name string `json:"Name,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SasAddress string `json:"SasAddress,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Controller = dataAO1.Controller

	m.EquipmentChassis = dataAO1.EquipmentChassis

	m.ExpanderID = dataAO1.ExpanderID

	m.Name = dataAO1.Name

	m.OperState = dataAO1.OperState

	m.Operability = dataAO1.Operability

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.SasAddress = dataAO1.SasAddress

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageSasExpander) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Controller *ManagementControllerRef `json:"Controller,omitempty"`

		EquipmentChassis *EquipmentChassisRef `json:"EquipmentChassis,omitempty"`

		ExpanderID int64 `json:"ExpanderId,omitempty"`

		Name string `json:"Name,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Operability string `json:"Operability,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SasAddress string `json:"SasAddress,omitempty"`
	}

	dataAO1.Controller = m.Controller

	dataAO1.EquipmentChassis = m.EquipmentChassis

	dataAO1.ExpanderID = m.ExpanderID

	dataAO1.Name = m.Name

	dataAO1.OperState = m.OperState

	dataAO1.Operability = m.Operability

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.SasAddress = m.SasAddress

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage sas expander
func (m *StorageSasExpander) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEquipmentChassis(formats); err != nil {
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

func (m *StorageSasExpander) validateController(formats strfmt.Registry) error {

	if swag.IsZero(m.Controller) { // not required
		return nil
	}

	if m.Controller != nil {
		if err := m.Controller.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Controller")
			}
			return err
		}
	}

	return nil
}

func (m *StorageSasExpander) validateEquipmentChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.EquipmentChassis) { // not required
		return nil
	}

	if m.EquipmentChassis != nil {
		if err := m.EquipmentChassis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EquipmentChassis")
			}
			return err
		}
	}

	return nil
}

func (m *StorageSasExpander) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StorageSasExpander) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSasExpander) UnmarshalBinary(b []byte) error {
	var res StorageSasExpander
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
