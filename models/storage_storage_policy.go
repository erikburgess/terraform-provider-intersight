// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageStoragePolicy Storage
//
// The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
//
// swagger:model storageStoragePolicy
type StorageStoragePolicy struct {
	PolicyAbstractPolicy

	// Relationship to the used disk group policies.
	//
	DiskGroupPolicies []*StorageDiskGroupPolicyRef `json:"DiskGroupPolicies"`

	// A collection of disks used as hot spares globally for all the RAID groups.
	//
	GlobalHotSpares []*StorageLocalDisk `json:"GlobalHotSpares"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Relationship to the profile objects.
	//
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

	// Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy.
	//
	RetainPolicyVirtualDrives *bool `json:"RetainPolicyVirtualDrives,omitempty"`

	// This is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved.
	//
	// Enum: [UnconfiguredGood Jbod]
	UnusedDisksState *string `json:"UnusedDisksState,omitempty"`

	// The list of virtual drives and the disk groups that need to be created through this policy.
	//
	VirtualDrives []*StorageVirtualDriveConfig `json:"VirtualDrives"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageStoragePolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		DiskGroupPolicies []*StorageDiskGroupPolicyRef `json:"DiskGroupPolicies"`

		GlobalHotSpares []*StorageLocalDisk `json:"GlobalHotSpares"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		RetainPolicyVirtualDrives *bool `json:"RetainPolicyVirtualDrives,omitempty"`

		UnusedDisksState *string `json:"UnusedDisksState,omitempty"`

		VirtualDrives []*StorageVirtualDriveConfig `json:"VirtualDrives"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DiskGroupPolicies = dataAO1.DiskGroupPolicies

	m.GlobalHotSpares = dataAO1.GlobalHotSpares

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	m.RetainPolicyVirtualDrives = dataAO1.RetainPolicyVirtualDrives

	m.UnusedDisksState = dataAO1.UnusedDisksState

	m.VirtualDrives = dataAO1.VirtualDrives

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageStoragePolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DiskGroupPolicies []*StorageDiskGroupPolicyRef `json:"DiskGroupPolicies"`

		GlobalHotSpares []*StorageLocalDisk `json:"GlobalHotSpares"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		RetainPolicyVirtualDrives *bool `json:"RetainPolicyVirtualDrives,omitempty"`

		UnusedDisksState *string `json:"UnusedDisksState,omitempty"`

		VirtualDrives []*StorageVirtualDriveConfig `json:"VirtualDrives"`
	}

	dataAO1.DiskGroupPolicies = m.DiskGroupPolicies

	dataAO1.GlobalHotSpares = m.GlobalHotSpares

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	dataAO1.RetainPolicyVirtualDrives = m.RetainPolicyVirtualDrives

	dataAO1.UnusedDisksState = m.UnusedDisksState

	dataAO1.VirtualDrives = m.VirtualDrives

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage storage policy
func (m *StorageStoragePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskGroupPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalHotSpares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnusedDisksState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualDrives(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageStoragePolicy) validateDiskGroupPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskGroupPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.DiskGroupPolicies); i++ {
		if swag.IsZero(m.DiskGroupPolicies[i]) { // not required
			continue
		}

		if m.DiskGroupPolicies[i] != nil {
			if err := m.DiskGroupPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DiskGroupPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageStoragePolicy) validateGlobalHotSpares(formats strfmt.Registry) error {

	if swag.IsZero(m.GlobalHotSpares) { // not required
		return nil
	}

	for i := 0; i < len(m.GlobalHotSpares); i++ {
		if swag.IsZero(m.GlobalHotSpares[i]) { // not required
			continue
		}

		if m.GlobalHotSpares[i] != nil {
			if err := m.GlobalHotSpares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GlobalHotSpares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StorageStoragePolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *StorageStoragePolicy) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var storageStoragePolicyTypeUnusedDisksStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UnconfiguredGood","Jbod"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageStoragePolicyTypeUnusedDisksStatePropEnum = append(storageStoragePolicyTypeUnusedDisksStatePropEnum, v)
	}
}

// property enum
func (m *StorageStoragePolicy) validateUnusedDisksStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageStoragePolicyTypeUnusedDisksStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageStoragePolicy) validateUnusedDisksState(formats strfmt.Registry) error {

	if swag.IsZero(m.UnusedDisksState) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnusedDisksStateEnum("UnusedDisksState", "body", *m.UnusedDisksState); err != nil {
		return err
	}

	return nil
}

func (m *StorageStoragePolicy) validateVirtualDrives(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualDrives) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualDrives); i++ {
		if swag.IsZero(m.VirtualDrives[i]) { // not required
			continue
		}

		if m.VirtualDrives[i] != nil {
			if err := m.VirtualDrives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("VirtualDrives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageStoragePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageStoragePolicy) UnmarshalBinary(b []byte) error {
	var res StorageStoragePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
