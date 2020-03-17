// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoragePureVolumeSnapshot Storage:Pure Volume Snapshot
//
// Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
//
// swagger:model storagePureVolumeSnapshot
type StoragePureVolumeSnapshot struct {
	StorageSnapshot

	// Device registration managed object that represents this storage array connection to Intersight.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Unique serial number of the snapshot allocated by the storage array.
	// Read Only: true
	Serial string `json:"Serial,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StoragePureVolumeSnapshot) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageSnapshot
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageSnapshot = aO0

	// AO1
	var dataAO1 struct {
		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Serial string `json:"Serial,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Serial = dataAO1.Serial

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StoragePureVolumeSnapshot) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageSnapshot)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Serial string `json:"Serial,omitempty"`
	}

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Serial = m.Serial

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage pure volume snapshot
func (m *StoragePureVolumeSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageSnapshot
	if err := m.StorageSnapshot.Validate(formats); err != nil {
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

func (m *StoragePureVolumeSnapshot) validateRegisteredDevice(formats strfmt.Registry) error {

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
func (m *StoragePureVolumeSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePureVolumeSnapshot) UnmarshalBinary(b []byte) error {
	var res StoragePureVolumeSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
