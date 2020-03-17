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

// RecoveryConfigResult Recovery:Config Result
//
// Profile configuration (deploy, validation) results with the overall state and detailed result messages.
//
// swagger:model recoveryConfigResult
type RecoveryConfigResult struct {
	PolicyAbstractConfigResult

	// A collection of references to the [recovery.BackupProfile](mo://recovery.BackupProfile) Managed Object.
	// When this managed object is deleted, the referenced [recovery.BackupProfile](mo://recovery.BackupProfile) MO unsets its reference to this deleted MO.
	// Read Only: true
	BackupProfile *RecoveryBackupProfileRef `json:"BackupProfile,omitempty"`

	// Detailed result entries for both validation & configration. Each result entry can be error/warning/info messages and the context.
	ResultEntries []*RecoveryConfigResultEntryRef `json:"ResultEntries"`

	// A collection of references to the [recovery.OnDemandBackup](mo://recovery.OnDemandBackup) Managed Object.
	// When this managed object is deleted, the referenced [recovery.OnDemandBackup](mo://recovery.OnDemandBackup) MO unsets its reference to this deleted MO.
	// Read Only: true
	Nr0OnDemandBackup *RecoveryOnDemandBackupRef `json:"_0_OnDemandBackup,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoveryConfigResult) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractConfigResult
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractConfigResult = aO0

	// AO1
	var dataAO1 struct {
		BackupProfile *RecoveryBackupProfileRef `json:"BackupProfile,omitempty"`

		ResultEntries []*RecoveryConfigResultEntryRef `json:"ResultEntries"`

		Nr0OnDemandBackup *RecoveryOnDemandBackupRef `json:"_0_OnDemandBackup,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BackupProfile = dataAO1.BackupProfile

	m.ResultEntries = dataAO1.ResultEntries

	m.Nr0OnDemandBackup = dataAO1.Nr0OnDemandBackup

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoveryConfigResult) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractConfigResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		BackupProfile *RecoveryBackupProfileRef `json:"BackupProfile,omitempty"`

		ResultEntries []*RecoveryConfigResultEntryRef `json:"ResultEntries"`

		Nr0OnDemandBackup *RecoveryOnDemandBackupRef `json:"_0_OnDemandBackup,omitempty"`
	}

	dataAO1.BackupProfile = m.BackupProfile

	dataAO1.ResultEntries = m.ResultEntries

	dataAO1.Nr0OnDemandBackup = m.Nr0OnDemandBackup

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recovery config result
func (m *RecoveryConfigResult) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractConfigResult
	if err := m.PolicyAbstractConfigResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNr0OnDemandBackup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryConfigResult) validateBackupProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.BackupProfile) { // not required
		return nil
	}

	if m.BackupProfile != nil {
		if err := m.BackupProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BackupProfile")
			}
			return err
		}
	}

	return nil
}

func (m *RecoveryConfigResult) validateResultEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.ResultEntries); i++ {
		if swag.IsZero(m.ResultEntries[i]) { // not required
			continue
		}

		if m.ResultEntries[i] != nil {
			if err := m.ResultEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ResultEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryConfigResult) validateNr0OnDemandBackup(formats strfmt.Registry) error {

	if swag.IsZero(m.Nr0OnDemandBackup) { // not required
		return nil
	}

	if m.Nr0OnDemandBackup != nil {
		if err := m.Nr0OnDemandBackup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_0_OnDemandBackup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryConfigResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryConfigResult) UnmarshalBinary(b []byte) error {
	var res RecoveryConfigResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
