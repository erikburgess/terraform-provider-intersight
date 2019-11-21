// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageStorageUtilization Storage:Storage Utilization
//
// Storage space utilized by Pure storage entity.
//
// swagger:model storageStorageUtilization
type StorageStorageUtilization struct {
	StorageCapacity

	// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
	//
	// Read Only: true
	DataReduction float32 `json:"DataReduction,omitempty"`

	// Physical space occupied by the snapshots, represented in bytes.
	//
	// Read Only: true
	Snapshot int64 `json:"Snapshot,omitempty"`

	// Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed.
	//
	// Read Only: true
	ThinProvisioned float32 `json:"ThinProvisioned,omitempty"`

	// Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array's flash modules.
	//
	// Read Only: true
	TotalReduction float32 `json:"TotalReduction,omitempty"`

	// Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size id represented in bytes.
	//
	// Read Only: true
	Volume int64 `json:"Volume,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageStorageUtilization) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageCapacity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageCapacity = aO0

	// AO1
	var dataAO1 struct {
		DataReduction float32 `json:"DataReduction,omitempty"`

		Snapshot int64 `json:"Snapshot,omitempty"`

		ThinProvisioned float32 `json:"ThinProvisioned,omitempty"`

		TotalReduction float32 `json:"TotalReduction,omitempty"`

		Volume int64 `json:"Volume,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DataReduction = dataAO1.DataReduction

	m.Snapshot = dataAO1.Snapshot

	m.ThinProvisioned = dataAO1.ThinProvisioned

	m.TotalReduction = dataAO1.TotalReduction

	m.Volume = dataAO1.Volume

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageStorageUtilization) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.StorageCapacity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DataReduction float32 `json:"DataReduction,omitempty"`

		Snapshot int64 `json:"Snapshot,omitempty"`

		ThinProvisioned float32 `json:"ThinProvisioned,omitempty"`

		TotalReduction float32 `json:"TotalReduction,omitempty"`

		Volume int64 `json:"Volume,omitempty"`
	}

	dataAO1.DataReduction = m.DataReduction

	dataAO1.Snapshot = m.Snapshot

	dataAO1.ThinProvisioned = m.ThinProvisioned

	dataAO1.TotalReduction = m.TotalReduction

	dataAO1.Volume = m.Volume

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage storage utilization
func (m *StorageStorageUtilization) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageCapacity
	if err := m.StorageCapacity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageStorageUtilization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageStorageUtilization) UnmarshalBinary(b []byte) error {
	var res StorageStorageUtilization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
