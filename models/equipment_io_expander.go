// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EquipmentIoExpander Equipment:Io Expander
//
// I/O expander card which is used as an extension for servers in a Colusa Chassis.
//
// swagger:model equipmentIoExpander
type EquipmentIoExpander struct {
	EquipmentBase

	// A collection of references to the [compute.Blade](mo://compute.Blade) Managed Object.
	//
	// When this managed object is deleted, the referenced [compute.Blade](mo://compute.Blade) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

	// Identifies the operational state of I/O expander.
	//
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// Identifies the availability of I/O expander.
	//
	// Read Only: true
	Presence string `json:"Presence,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EquipmentIoExpander) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ComputeBlade = dataAO1.ComputeBlade

	m.OperState = dataAO1.OperState

	m.Presence = dataAO1.Presence

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EquipmentIoExpander) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ComputeBlade *ComputeBladeRef `json:"ComputeBlade,omitempty"`

		OperState string `json:"OperState,omitempty"`

		Presence string `json:"Presence,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.ComputeBlade = m.ComputeBlade

	dataAO1.OperState = m.OperState

	dataAO1.Presence = m.Presence

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this equipment io expander
func (m *EquipmentIoExpander) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBlade(formats); err != nil {
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

func (m *EquipmentIoExpander) validateComputeBlade(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBlade) { // not required
		return nil
	}

	if m.ComputeBlade != nil {
		if err := m.ComputeBlade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBlade")
			}
			return err
		}
	}

	return nil
}

func (m *EquipmentIoExpander) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *EquipmentIoExpander) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EquipmentIoExpander) UnmarshalBinary(b []byte) error {
	var res EquipmentIoExpander
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
