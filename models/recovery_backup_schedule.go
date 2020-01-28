// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecoveryBackupSchedule Recovery:Backup Schedule
//
// Encapsulates the various backup settings available to the user for scheduling a backup on the endpoint.
//
// swagger:model recoveryBackupSchedule
type RecoveryBackupSchedule struct {
	MoBaseComplexType

	RecoveryBackupScheduleAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoveryBackupSchedule) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 RecoveryBackupScheduleAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.RecoveryBackupScheduleAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoveryBackupSchedule) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.RecoveryBackupScheduleAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recovery backup schedule
func (m *RecoveryBackupSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RecoveryBackupScheduleAO1P1
	if err := m.RecoveryBackupScheduleAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryBackupSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryBackupSchedule) UnmarshalBinary(b []byte) error {
	var res RecoveryBackupSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RecoveryBackupScheduleAO1P1 recovery backup schedule a o1 p1
// swagger:model RecoveryBackupScheduleAO1P1
type RecoveryBackupScheduleAO1P1 struct {

	// The time at which the backup is to be run on a given day. This is used when the frequency unit is daily.
	//
	// Format: date-time
	ExecutionTime strfmt.DateTime `json:"ExecutionTime,omitempty"`

	// The frequency at which the backup schedule must run.
	//
	// Enum: [Daily Periodic]
	FrequencyUnit *string `json:"FrequencyUnit,omitempty"`

	// The frequency, in hours, at which the backup schedule runs.
	//
	// Enum: [8 4 12 16 20]
	Hours *int64 `json:"Hours,omitempty"`

	// recovery backup schedule a o1 p1
	RecoveryBackupScheduleAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *RecoveryBackupScheduleAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The time at which the backup is to be run on a given day. This is used when the frequency unit is daily.
		//
		// Format: date-time
		ExecutionTime strfmt.DateTime `json:"ExecutionTime,omitempty"`

		// The frequency at which the backup schedule must run.
		//
		// Enum: [Daily Periodic]
		FrequencyUnit *string `json:"FrequencyUnit,omitempty"`

		// The frequency, in hours, at which the backup schedule runs.
		//
		// Enum: [8 4 12 16 20]
		Hours *int64 `json:"Hours,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv RecoveryBackupScheduleAO1P1

	rcv.ExecutionTime = stage1.ExecutionTime

	rcv.FrequencyUnit = stage1.FrequencyUnit

	rcv.Hours = stage1.Hours

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ExecutionTime")

	delete(stage2, "FrequencyUnit")

	delete(stage2, "Hours")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.RecoveryBackupScheduleAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m RecoveryBackupScheduleAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The time at which the backup is to be run on a given day. This is used when the frequency unit is daily.
		//
		// Format: date-time
		ExecutionTime strfmt.DateTime `json:"ExecutionTime,omitempty"`

		// The frequency at which the backup schedule must run.
		//
		// Enum: [Daily Periodic]
		FrequencyUnit *string `json:"FrequencyUnit,omitempty"`

		// The frequency, in hours, at which the backup schedule runs.
		//
		// Enum: [8 4 12 16 20]
		Hours *int64 `json:"Hours,omitempty"`
	}

	stage1.ExecutionTime = m.ExecutionTime

	stage1.FrequencyUnit = m.FrequencyUnit

	stage1.Hours = m.Hours

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.RecoveryBackupScheduleAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.RecoveryBackupScheduleAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this recovery backup schedule a o1 p1
func (m *RecoveryBackupScheduleAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHours(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryBackupScheduleAO1P1) validateExecutionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ExecutionTime", "body", "date-time", m.ExecutionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var recoveryBackupScheduleAO1P1TypeFrequencyUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Daily","Periodic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoveryBackupScheduleAO1P1TypeFrequencyUnitPropEnum = append(recoveryBackupScheduleAO1P1TypeFrequencyUnitPropEnum, v)
	}
}

const (

	// RecoveryBackupScheduleAO1P1FrequencyUnitDaily captures enum value "Daily"
	RecoveryBackupScheduleAO1P1FrequencyUnitDaily string = "Daily"

	// RecoveryBackupScheduleAO1P1FrequencyUnitPeriodic captures enum value "Periodic"
	RecoveryBackupScheduleAO1P1FrequencyUnitPeriodic string = "Periodic"
)

// prop value enum
func (m *RecoveryBackupScheduleAO1P1) validateFrequencyUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recoveryBackupScheduleAO1P1TypeFrequencyUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RecoveryBackupScheduleAO1P1) validateFrequencyUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.FrequencyUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateFrequencyUnitEnum("FrequencyUnit", "body", *m.FrequencyUnit); err != nil {
		return err
	}

	return nil
}

var recoveryBackupScheduleAO1P1TypeHoursPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[8,4,12,16,20]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoveryBackupScheduleAO1P1TypeHoursPropEnum = append(recoveryBackupScheduleAO1P1TypeHoursPropEnum, v)
	}
}

// prop value enum
func (m *RecoveryBackupScheduleAO1P1) validateHoursEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, recoveryBackupScheduleAO1P1TypeHoursPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RecoveryBackupScheduleAO1P1) validateHours(formats strfmt.Registry) error {

	if swag.IsZero(m.Hours) { // not required
		return nil
	}

	// value enum
	if err := m.validateHoursEnum("Hours", "body", *m.Hours); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryBackupScheduleAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryBackupScheduleAO1P1) UnmarshalBinary(b []byte) error {
	var res RecoveryBackupScheduleAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
