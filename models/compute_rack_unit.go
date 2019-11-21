// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ComputeRackUnit Compute:Rack Unit
//
// Rack-mounted server.
//
// swagger:model computeRackUnit
type ComputeRackUnit struct {
	ComputePhysical

	// adapters
	// Read Only: true
	Adapters []*AdapterUnitRef `json:"Adapters"`

	// bios bootmode
	BiosBootmode *BiosBootModeRef `json:"BiosBootmode,omitempty"`

	// biosunits
	// Read Only: true
	Biosunits []*BiosUnitRef `json:"Biosunits"`

	// bmc
	// Read Only: true
	Bmc *ManagementControllerRef `json:"Bmc,omitempty"`

	// board
	// Read Only: true
	Board *ComputeBoardRef `json:"Board,omitempty"`

	// boot device bootmode
	// Read Only: true
	BootDeviceBootmode *BootDeviceBootModeRef `json:"BootDeviceBootmode,omitempty"`

	// fanmodules
	// Read Only: true
	Fanmodules []*EquipmentFanModuleRef `json:"Fanmodules"`

	// generic inventory holders
	// Read Only: true
	GenericInventoryHolders []*InventoryGenericInventoryHolderRef `json:"GenericInventoryHolders"`

	// locator led
	// Read Only: true
	LocatorLed *EquipmentLocatorLedRef `json:"LocatorLed,omitempty"`

	// pci devices
	// Read Only: true
	PciDevices []*PciDeviceRef `json:"PciDevices"`

	// psus
	// Read Only: true
	Psus []*EquipmentPsuRef `json:"Psus"`

	// A collection of references to the [equipment.RackEnclosureSlot](mo://equipment.RackEnclosureSlot) Managed Object.
	//
	// When this managed object is deleted, the referenced [equipment.RackEnclosureSlot](mo://equipment.RackEnclosureSlot) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	RackEnclosureSlot *EquipmentRackEnclosureSlotRef `json:"RackEnclosureSlot,omitempty"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// sas expanders
	SasExpanders []*StorageSasExpanderRef `json:"SasExpanders"`

	// server Id
	// Read Only: true
	ServerID int64 `json:"ServerId,omitempty"`

	// storage enclosures
	// Read Only: true
	StorageEnclosures []*StorageEnclosureRef `json:"StorageEnclosures"`

	// A collection of references to the [top.System](mo://top.System) Managed Object.
	//
	// When this managed object is deleted, the referenced [top.System](mo://top.System) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	TopSystem *TopSystemRef `json:"TopSystem,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ComputeRackUnit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ComputePhysical
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ComputePhysical = aO0

	// AO1
	var dataAO1 struct {
		Adapters []*AdapterUnitRef `json:"Adapters"`

		BiosBootmode *BiosBootModeRef `json:"BiosBootmode,omitempty"`

		Biosunits []*BiosUnitRef `json:"Biosunits"`

		Bmc *ManagementControllerRef `json:"Bmc,omitempty"`

		Board *ComputeBoardRef `json:"Board,omitempty"`

		BootDeviceBootmode *BootDeviceBootModeRef `json:"BootDeviceBootmode,omitempty"`

		Fanmodules []*EquipmentFanModuleRef `json:"Fanmodules"`

		GenericInventoryHolders []*InventoryGenericInventoryHolderRef `json:"GenericInventoryHolders"`

		LocatorLed *EquipmentLocatorLedRef `json:"LocatorLed,omitempty"`

		PciDevices []*PciDeviceRef `json:"PciDevices"`

		Psus []*EquipmentPsuRef `json:"Psus"`

		RackEnclosureSlot *EquipmentRackEnclosureSlotRef `json:"RackEnclosureSlot,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SasExpanders []*StorageSasExpanderRef `json:"SasExpanders"`

		ServerID int64 `json:"ServerId,omitempty"`

		StorageEnclosures []*StorageEnclosureRef `json:"StorageEnclosures"`

		TopSystem *TopSystemRef `json:"TopSystem,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Adapters = dataAO1.Adapters

	m.BiosBootmode = dataAO1.BiosBootmode

	m.Biosunits = dataAO1.Biosunits

	m.Bmc = dataAO1.Bmc

	m.Board = dataAO1.Board

	m.BootDeviceBootmode = dataAO1.BootDeviceBootmode

	m.Fanmodules = dataAO1.Fanmodules

	m.GenericInventoryHolders = dataAO1.GenericInventoryHolders

	m.LocatorLed = dataAO1.LocatorLed

	m.PciDevices = dataAO1.PciDevices

	m.Psus = dataAO1.Psus

	m.RackEnclosureSlot = dataAO1.RackEnclosureSlot

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.SasExpanders = dataAO1.SasExpanders

	m.ServerID = dataAO1.ServerID

	m.StorageEnclosures = dataAO1.StorageEnclosures

	m.TopSystem = dataAO1.TopSystem

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ComputeRackUnit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ComputePhysical)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Adapters []*AdapterUnitRef `json:"Adapters"`

		BiosBootmode *BiosBootModeRef `json:"BiosBootmode,omitempty"`

		Biosunits []*BiosUnitRef `json:"Biosunits"`

		Bmc *ManagementControllerRef `json:"Bmc,omitempty"`

		Board *ComputeBoardRef `json:"Board,omitempty"`

		BootDeviceBootmode *BootDeviceBootModeRef `json:"BootDeviceBootmode,omitempty"`

		Fanmodules []*EquipmentFanModuleRef `json:"Fanmodules"`

		GenericInventoryHolders []*InventoryGenericInventoryHolderRef `json:"GenericInventoryHolders"`

		LocatorLed *EquipmentLocatorLedRef `json:"LocatorLed,omitempty"`

		PciDevices []*PciDeviceRef `json:"PciDevices"`

		Psus []*EquipmentPsuRef `json:"Psus"`

		RackEnclosureSlot *EquipmentRackEnclosureSlotRef `json:"RackEnclosureSlot,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		SasExpanders []*StorageSasExpanderRef `json:"SasExpanders"`

		ServerID int64 `json:"ServerId,omitempty"`

		StorageEnclosures []*StorageEnclosureRef `json:"StorageEnclosures"`

		TopSystem *TopSystemRef `json:"TopSystem,omitempty"`
	}

	dataAO1.Adapters = m.Adapters

	dataAO1.BiosBootmode = m.BiosBootmode

	dataAO1.Biosunits = m.Biosunits

	dataAO1.Bmc = m.Bmc

	dataAO1.Board = m.Board

	dataAO1.BootDeviceBootmode = m.BootDeviceBootmode

	dataAO1.Fanmodules = m.Fanmodules

	dataAO1.GenericInventoryHolders = m.GenericInventoryHolders

	dataAO1.LocatorLed = m.LocatorLed

	dataAO1.PciDevices = m.PciDevices

	dataAO1.Psus = m.Psus

	dataAO1.RackEnclosureSlot = m.RackEnclosureSlot

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.SasExpanders = m.SasExpanders

	dataAO1.ServerID = m.ServerID

	dataAO1.StorageEnclosures = m.StorageEnclosures

	dataAO1.TopSystem = m.TopSystem

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this compute rack unit
func (m *ComputeRackUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ComputePhysical
	if err := m.ComputePhysical.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdapters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBiosBootmode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBiosunits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBmc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootDeviceBootmode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFanmodules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenericInventoryHolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocatorLed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePciDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePsus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRackEnclosureSlot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSasExpanders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageEnclosures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeRackUnit) validateAdapters(formats strfmt.Registry) error {

	if swag.IsZero(m.Adapters) { // not required
		return nil
	}

	for i := 0; i < len(m.Adapters); i++ {
		if swag.IsZero(m.Adapters[i]) { // not required
			continue
		}

		if m.Adapters[i] != nil {
			if err := m.Adapters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Adapters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateBiosBootmode(formats strfmt.Registry) error {

	if swag.IsZero(m.BiosBootmode) { // not required
		return nil
	}

	if m.BiosBootmode != nil {
		if err := m.BiosBootmode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BiosBootmode")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeRackUnit) validateBiosunits(formats strfmt.Registry) error {

	if swag.IsZero(m.Biosunits) { // not required
		return nil
	}

	for i := 0; i < len(m.Biosunits); i++ {
		if swag.IsZero(m.Biosunits[i]) { // not required
			continue
		}

		if m.Biosunits[i] != nil {
			if err := m.Biosunits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Biosunits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateBmc(formats strfmt.Registry) error {

	if swag.IsZero(m.Bmc) { // not required
		return nil
	}

	if m.Bmc != nil {
		if err := m.Bmc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Bmc")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeRackUnit) validateBoard(formats strfmt.Registry) error {

	if swag.IsZero(m.Board) { // not required
		return nil
	}

	if m.Board != nil {
		if err := m.Board.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Board")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeRackUnit) validateBootDeviceBootmode(formats strfmt.Registry) error {

	if swag.IsZero(m.BootDeviceBootmode) { // not required
		return nil
	}

	if m.BootDeviceBootmode != nil {
		if err := m.BootDeviceBootmode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BootDeviceBootmode")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeRackUnit) validateFanmodules(formats strfmt.Registry) error {

	if swag.IsZero(m.Fanmodules) { // not required
		return nil
	}

	for i := 0; i < len(m.Fanmodules); i++ {
		if swag.IsZero(m.Fanmodules[i]) { // not required
			continue
		}

		if m.Fanmodules[i] != nil {
			if err := m.Fanmodules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fanmodules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateGenericInventoryHolders(formats strfmt.Registry) error {

	if swag.IsZero(m.GenericInventoryHolders) { // not required
		return nil
	}

	for i := 0; i < len(m.GenericInventoryHolders); i++ {
		if swag.IsZero(m.GenericInventoryHolders[i]) { // not required
			continue
		}

		if m.GenericInventoryHolders[i] != nil {
			if err := m.GenericInventoryHolders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GenericInventoryHolders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateLocatorLed(formats strfmt.Registry) error {

	if swag.IsZero(m.LocatorLed) { // not required
		return nil
	}

	if m.LocatorLed != nil {
		if err := m.LocatorLed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LocatorLed")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeRackUnit) validatePciDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.PciDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.PciDevices); i++ {
		if swag.IsZero(m.PciDevices[i]) { // not required
			continue
		}

		if m.PciDevices[i] != nil {
			if err := m.PciDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PciDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validatePsus(formats strfmt.Registry) error {

	if swag.IsZero(m.Psus) { // not required
		return nil
	}

	for i := 0; i < len(m.Psus); i++ {
		if swag.IsZero(m.Psus[i]) { // not required
			continue
		}

		if m.Psus[i] != nil {
			if err := m.Psus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Psus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateRackEnclosureSlot(formats strfmt.Registry) error {

	if swag.IsZero(m.RackEnclosureSlot) { // not required
		return nil
	}

	if m.RackEnclosureSlot != nil {
		if err := m.RackEnclosureSlot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RackEnclosureSlot")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeRackUnit) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *ComputeRackUnit) validateSasExpanders(formats strfmt.Registry) error {

	if swag.IsZero(m.SasExpanders) { // not required
		return nil
	}

	for i := 0; i < len(m.SasExpanders); i++ {
		if swag.IsZero(m.SasExpanders[i]) { // not required
			continue
		}

		if m.SasExpanders[i] != nil {
			if err := m.SasExpanders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SasExpanders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateStorageEnclosures(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageEnclosures) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageEnclosures); i++ {
		if swag.IsZero(m.StorageEnclosures[i]) { // not required
			continue
		}

		if m.StorageEnclosures[i] != nil {
			if err := m.StorageEnclosures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StorageEnclosures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeRackUnit) validateTopSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.TopSystem) { // not required
		return nil
	}

	if m.TopSystem != nil {
		if err := m.TopSystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TopSystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeRackUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeRackUnit) UnmarshalBinary(b []byte) error {
	var res ComputeRackUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
