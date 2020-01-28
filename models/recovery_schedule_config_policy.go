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

// RecoveryScheduleConfigPolicy Schedule Policy
//
// Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
//
// swagger:model recoveryScheduleConfigPolicy
type RecoveryScheduleConfigPolicy struct {
	PolicyAbstractPolicy

	// List of Backup profiles using this policy.
	//
	BackupProfiles []*RecoveryBackupProfileRef `json:"BackupProfiles"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Schedule to create a backup on the target device. Minimum is 4 hours and Max is 1440 hours (30 Days).
	//
	Schedule *RecoveryBackupSchedule `json:"Schedule,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoveryScheduleConfigPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		BackupProfiles []*RecoveryBackupProfileRef `json:"BackupProfiles"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Schedule *RecoveryBackupSchedule `json:"Schedule,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BackupProfiles = dataAO1.BackupProfiles

	m.Organization = dataAO1.Organization

	m.Schedule = dataAO1.Schedule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoveryScheduleConfigPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BackupProfiles []*RecoveryBackupProfileRef `json:"BackupProfiles"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Schedule *RecoveryBackupSchedule `json:"Schedule,omitempty"`
	}

	dataAO1.BackupProfiles = m.BackupProfiles

	dataAO1.Organization = m.Organization

	dataAO1.Schedule = m.Schedule

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recovery schedule config policy
func (m *RecoveryScheduleConfigPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryScheduleConfigPolicy) validateBackupProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.BackupProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupProfiles); i++ {
		if swag.IsZero(m.BackupProfiles[i]) { // not required
			continue
		}

		if m.BackupProfiles[i] != nil {
			if err := m.BackupProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BackupProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryScheduleConfigPolicy) validateOrganization(formats strfmt.Registry) error {

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

func (m *RecoveryScheduleConfigPolicy) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryScheduleConfigPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryScheduleConfigPolicy) UnmarshalBinary(b []byte) error {
	var res RecoveryScheduleConfigPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
