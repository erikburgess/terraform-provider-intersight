// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GraphicsCard Graphics:Card
//
// Graphics Card present in a server.
//
// swagger:model graphicsCard
type GraphicsCard struct {
	EquipmentBase

	// It shows the id of graphics card.
	// Read Only: true
	CardID int64 `json:"CardId,omitempty"`

	// A collection of references to the [compute.Board](mo://compute.Board) Managed Object.
	// When this managed object is deleted, the referenced [compute.Board](mo://compute.Board) MO unsets its reference to this deleted MO.
	// Read Only: true
	ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

	// It shows the device id of grphics card.
	// Read Only: true
	DeviceID int64 `json:"DeviceId,omitempty"`

	// It shows the expander slot inforamtion for the card.
	// Read Only: true
	ExpanderSlot string `json:"ExpanderSlot,omitempty"`

	// It shows current firmware version of graphics card.
	// Read Only: true
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// It shows the controllers under each graphics card.
	// Read Only: true
	GraphicsControllers []*GraphicsControllerRef `json:"GraphicsControllers"`

	// It shows the current mode of graphics card.
	// Read Only: true
	Mode string `json:"Mode,omitempty"`

	// It shows number of controllers under each card.
	NumGpus string `json:"NumGpus,omitempty"`

	// It shows the current operational state of graphics card.
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// It shows the pci address of graphics card.
	// Read Only: true
	PciAddress string `json:"PciAddress,omitempty"`

	// This list contains the pci address of all controllers for corresponding card.
	// Read Only: true
	PciAddressList string `json:"PciAddressList,omitempty"`

	// It shows the pci slot name for grapchics card.
	// Read Only: true
	PciSlot string `json:"PciSlot,omitempty"`

	// The Device to which this Managed Object is associated.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GraphicsCard) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		CardID int64 `json:"CardId,omitempty"`

		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		DeviceID int64 `json:"DeviceId,omitempty"`

		ExpanderSlot string `json:"ExpanderSlot,omitempty"`

		FirmwareVersion string `json:"FirmwareVersion,omitempty"`

		GraphicsControllers []*GraphicsControllerRef `json:"GraphicsControllers"`

		Mode string `json:"Mode,omitempty"`

		NumGpus string `json:"NumGpus,omitempty"`

		OperState string `json:"OperState,omitempty"`

		PciAddress string `json:"PciAddress,omitempty"`

		PciAddressList string `json:"PciAddressList,omitempty"`

		PciSlot string `json:"PciSlot,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CardID = dataAO1.CardID

	m.ComputeBoard = dataAO1.ComputeBoard

	m.DeviceID = dataAO1.DeviceID

	m.ExpanderSlot = dataAO1.ExpanderSlot

	m.FirmwareVersion = dataAO1.FirmwareVersion

	m.GraphicsControllers = dataAO1.GraphicsControllers

	m.Mode = dataAO1.Mode

	m.NumGpus = dataAO1.NumGpus

	m.OperState = dataAO1.OperState

	m.PciAddress = dataAO1.PciAddress

	m.PciAddressList = dataAO1.PciAddressList

	m.PciSlot = dataAO1.PciSlot

	m.RegisteredDevice = dataAO1.RegisteredDevice

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GraphicsCard) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CardID int64 `json:"CardId,omitempty"`

		ComputeBoard *ComputeBoardRef `json:"ComputeBoard,omitempty"`

		DeviceID int64 `json:"DeviceId,omitempty"`

		ExpanderSlot string `json:"ExpanderSlot,omitempty"`

		FirmwareVersion string `json:"FirmwareVersion,omitempty"`

		GraphicsControllers []*GraphicsControllerRef `json:"GraphicsControllers"`

		Mode string `json:"Mode,omitempty"`

		NumGpus string `json:"NumGpus,omitempty"`

		OperState string `json:"OperState,omitempty"`

		PciAddress string `json:"PciAddress,omitempty"`

		PciAddressList string `json:"PciAddressList,omitempty"`

		PciSlot string `json:"PciSlot,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`
	}

	dataAO1.CardID = m.CardID

	dataAO1.ComputeBoard = m.ComputeBoard

	dataAO1.DeviceID = m.DeviceID

	dataAO1.ExpanderSlot = m.ExpanderSlot

	dataAO1.FirmwareVersion = m.FirmwareVersion

	dataAO1.GraphicsControllers = m.GraphicsControllers

	dataAO1.Mode = m.Mode

	dataAO1.NumGpus = m.NumGpus

	dataAO1.OperState = m.OperState

	dataAO1.PciAddress = m.PciAddress

	dataAO1.PciAddressList = m.PciAddressList

	dataAO1.PciSlot = m.PciSlot

	dataAO1.RegisteredDevice = m.RegisteredDevice

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this graphics card
func (m *GraphicsCard) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComputeBoard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraphicsControllers(formats); err != nil {
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

func (m *GraphicsCard) validateComputeBoard(formats strfmt.Registry) error {

	if swag.IsZero(m.ComputeBoard) { // not required
		return nil
	}

	if m.ComputeBoard != nil {
		if err := m.ComputeBoard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ComputeBoard")
			}
			return err
		}
	}

	return nil
}

func (m *GraphicsCard) validateGraphicsControllers(formats strfmt.Registry) error {

	if swag.IsZero(m.GraphicsControllers) { // not required
		return nil
	}

	for i := 0; i < len(m.GraphicsControllers); i++ {
		if swag.IsZero(m.GraphicsControllers[i]) { // not required
			continue
		}

		if m.GraphicsControllers[i] != nil {
			if err := m.GraphicsControllers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GraphicsControllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GraphicsCard) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *GraphicsCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphicsCard) UnmarshalBinary(b []byte) error {
	var res GraphicsCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
