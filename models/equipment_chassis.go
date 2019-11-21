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

// EquipmentChassis Equipment:Chassis
//
// A physical holder housing blade servers.
//
// swagger:model equipmentChassis
type EquipmentChassis struct {
	EquipmentBase

	// blades
	// Read Only: true
	Blades []*ComputeBladeRef `json:"Blades"`

	// The assigned identifier for a chassis.
	//
	// Read Only: true
	ChassisID int64 `json:"ChassisId,omitempty"`

	// This field identifies the connectivity path for the chassis enclosure.
	//
	// Read Only: true
	ConnectionPath string `json:"ConnectionPath,omitempty"`

	// This field identifies the connectivity status for the chassis enclosure.
	//
	// Read Only: true
	ConnectionStatus string `json:"ConnectionStatus,omitempty"`

	// This field is to provide description for chassis model.
	//
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// fanmodules
	// Read Only: true
	Fanmodules []*EquipmentFanModuleRef `json:"Fanmodules"`

	// fault summary
	FaultSummary int64 `json:"FaultSummary,omitempty"`

	// ioms
	// Read Only: true
	Ioms []*EquipmentIoCardRef `json:"Ioms"`

	// This field identifies the name for the chassis enclosure.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// oper state
	// Read Only: true
	OperState string `json:"OperState,omitempty"`

	// Part Number identifier for the chassis enclosure.
	//
	// Read Only: true
	PartNumber string `json:"PartNumber,omitempty"`

	// This field identifies the Product ID for the chassis enclosure.
	//
	// Read Only: true
	Pid string `json:"Pid,omitempty"`

	// platform type
	PlatformType string `json:"PlatformType,omitempty"`

	// This field identifies the Product Name for the chassis enclosure.
	//
	// Read Only: true
	ProductName string `json:"ProductName,omitempty"`

	// psus
	// Read Only: true
	Psus []*EquipmentPsuRef `json:"Psus"`

	// The Device to which this Managed Object is associated.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// sasexpanders
	// Read Only: true
	Sasexpanders []*StorageSasExpanderRef `json:"Sasexpanders"`

	// siocs
	// Read Only: true
	Siocs []*EquipmentSystemIoControllerRef `json:"Siocs"`

	// This field identifies the Stock Keeping Unit for the chassis enclosure.
	//
	// Read Only: true
	Sku string `json:"Sku,omitempty"`

	// This field identifies the chassis enclosures.
	//
	// Read Only: true
	StorageEnclosures []*StorageEnclosureRef `json:"StorageEnclosures"`

	// This field identifies the Vendor ID for the chassis enclosure.
	//
	// Read Only: true
	Vid string `json:"Vid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EquipmentChassis) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EquipmentBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EquipmentBase = aO0

	// AO1
	var dataAO1 struct {
		Blades []*ComputeBladeRef `json:"Blades"`

		ChassisID int64 `json:"ChassisId,omitempty"`

		ConnectionPath string `json:"ConnectionPath,omitempty"`

		ConnectionStatus string `json:"ConnectionStatus,omitempty"`

		Description string `json:"Description,omitempty"`

		Fanmodules []*EquipmentFanModuleRef `json:"Fanmodules"`

		FaultSummary int64 `json:"FaultSummary,omitempty"`

		Ioms []*EquipmentIoCardRef `json:"Ioms"`

		Name string `json:"Name,omitempty"`

		OperState string `json:"OperState,omitempty"`

		PartNumber string `json:"PartNumber,omitempty"`

		Pid string `json:"Pid,omitempty"`

		PlatformType string `json:"PlatformType,omitempty"`

		ProductName string `json:"ProductName,omitempty"`

		Psus []*EquipmentPsuRef `json:"Psus"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Sasexpanders []*StorageSasExpanderRef `json:"Sasexpanders"`

		Siocs []*EquipmentSystemIoControllerRef `json:"Siocs"`

		Sku string `json:"Sku,omitempty"`

		StorageEnclosures []*StorageEnclosureRef `json:"StorageEnclosures"`

		Vid string `json:"Vid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Blades = dataAO1.Blades

	m.ChassisID = dataAO1.ChassisID

	m.ConnectionPath = dataAO1.ConnectionPath

	m.ConnectionStatus = dataAO1.ConnectionStatus

	m.Description = dataAO1.Description

	m.Fanmodules = dataAO1.Fanmodules

	m.FaultSummary = dataAO1.FaultSummary

	m.Ioms = dataAO1.Ioms

	m.Name = dataAO1.Name

	m.OperState = dataAO1.OperState

	m.PartNumber = dataAO1.PartNumber

	m.Pid = dataAO1.Pid

	m.PlatformType = dataAO1.PlatformType

	m.ProductName = dataAO1.ProductName

	m.Psus = dataAO1.Psus

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Sasexpanders = dataAO1.Sasexpanders

	m.Siocs = dataAO1.Siocs

	m.Sku = dataAO1.Sku

	m.StorageEnclosures = dataAO1.StorageEnclosures

	m.Vid = dataAO1.Vid

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EquipmentChassis) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EquipmentBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Blades []*ComputeBladeRef `json:"Blades"`

		ChassisID int64 `json:"ChassisId,omitempty"`

		ConnectionPath string `json:"ConnectionPath,omitempty"`

		ConnectionStatus string `json:"ConnectionStatus,omitempty"`

		Description string `json:"Description,omitempty"`

		Fanmodules []*EquipmentFanModuleRef `json:"Fanmodules"`

		FaultSummary int64 `json:"FaultSummary,omitempty"`

		Ioms []*EquipmentIoCardRef `json:"Ioms"`

		Name string `json:"Name,omitempty"`

		OperState string `json:"OperState,omitempty"`

		PartNumber string `json:"PartNumber,omitempty"`

		Pid string `json:"Pid,omitempty"`

		PlatformType string `json:"PlatformType,omitempty"`

		ProductName string `json:"ProductName,omitempty"`

		Psus []*EquipmentPsuRef `json:"Psus"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Sasexpanders []*StorageSasExpanderRef `json:"Sasexpanders"`

		Siocs []*EquipmentSystemIoControllerRef `json:"Siocs"`

		Sku string `json:"Sku,omitempty"`

		StorageEnclosures []*StorageEnclosureRef `json:"StorageEnclosures"`

		Vid string `json:"Vid,omitempty"`
	}

	dataAO1.Blades = m.Blades

	dataAO1.ChassisID = m.ChassisID

	dataAO1.ConnectionPath = m.ConnectionPath

	dataAO1.ConnectionStatus = m.ConnectionStatus

	dataAO1.Description = m.Description

	dataAO1.Fanmodules = m.Fanmodules

	dataAO1.FaultSummary = m.FaultSummary

	dataAO1.Ioms = m.Ioms

	dataAO1.Name = m.Name

	dataAO1.OperState = m.OperState

	dataAO1.PartNumber = m.PartNumber

	dataAO1.Pid = m.Pid

	dataAO1.PlatformType = m.PlatformType

	dataAO1.ProductName = m.ProductName

	dataAO1.Psus = m.Psus

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Sasexpanders = m.Sasexpanders

	dataAO1.Siocs = m.Siocs

	dataAO1.Sku = m.Sku

	dataAO1.StorageEnclosures = m.StorageEnclosures

	dataAO1.Vid = m.Vid

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this equipment chassis
func (m *EquipmentChassis) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EquipmentBase
	if err := m.EquipmentBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlades(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFanmodules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePsus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSasexpanders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiocs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageEnclosures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EquipmentChassis) validateBlades(formats strfmt.Registry) error {

	if swag.IsZero(m.Blades) { // not required
		return nil
	}

	for i := 0; i < len(m.Blades); i++ {
		if swag.IsZero(m.Blades[i]) { // not required
			continue
		}

		if m.Blades[i] != nil {
			if err := m.Blades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Blades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EquipmentChassis) validateFanmodules(formats strfmt.Registry) error {

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

func (m *EquipmentChassis) validateIoms(formats strfmt.Registry) error {

	if swag.IsZero(m.Ioms) { // not required
		return nil
	}

	for i := 0; i < len(m.Ioms); i++ {
		if swag.IsZero(m.Ioms[i]) { // not required
			continue
		}

		if m.Ioms[i] != nil {
			if err := m.Ioms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ioms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EquipmentChassis) validatePsus(formats strfmt.Registry) error {

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

func (m *EquipmentChassis) validateRegisteredDevice(formats strfmt.Registry) error {

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

func (m *EquipmentChassis) validateSasexpanders(formats strfmt.Registry) error {

	if swag.IsZero(m.Sasexpanders) { // not required
		return nil
	}

	for i := 0; i < len(m.Sasexpanders); i++ {
		if swag.IsZero(m.Sasexpanders[i]) { // not required
			continue
		}

		if m.Sasexpanders[i] != nil {
			if err := m.Sasexpanders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sasexpanders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EquipmentChassis) validateSiocs(formats strfmt.Registry) error {

	if swag.IsZero(m.Siocs) { // not required
		return nil
	}

	for i := 0; i < len(m.Siocs); i++ {
		if swag.IsZero(m.Siocs[i]) { // not required
			continue
		}

		if m.Siocs[i] != nil {
			if err := m.Siocs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Siocs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EquipmentChassis) validateStorageEnclosures(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *EquipmentChassis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EquipmentChassis) UnmarshalBinary(b []byte) error {
	var res EquipmentChassis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
