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

// InventoryJobInfo Inventory:Job Info
//
// Complex type representing the Job Information of a device.
//
// swagger:model inventoryJobInfo
type InventoryJobInfo struct {
	InventoryJobInfoAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InventoryJobInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryJobInfoAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryJobInfoAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InventoryJobInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.InventoryJobInfoAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this inventory job info
func (m *InventoryJobInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryJobInfoAO0P0
	if err := m.InventoryJobInfoAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryJobInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryJobInfo) UnmarshalBinary(b []byte) error {
	var res InventoryJobInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryJobInfoAO0P0 inventory job info a o0 p0
// swagger:model InventoryJobInfoAO0P0
type InventoryJobInfoAO0P0 struct {

	// The execution status of the job scheduled.
	//
	// Read Only: true
	// Enum: [Unknown Scheduled Completed Error]
	ExecutionStatus string `json:"ExecutionStatus,omitempty"`

	// The name of the job scheduled.
	//
	// Read Only: true
	JobName string `json:"JobName,omitempty"`

	// The last processed time of the job scheduled.
	//
	// Read Only: true
	// Format: date-time
	LastProcessedTime strfmt.DateTime `json:"LastProcessedTime,omitempty"`

	// The last scheduled time of the job scheduled.
	//
	// Read Only: true
	// Format: date-time
	LastScheduledTime strfmt.DateTime `json:"LastScheduledTime,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The propertie of the job scheduled.
	//
	// Read Only: true
	Properties []string `json:"Properties"`

	// regex
	// Read Only: true
	Regex []string `json:"Regex"`

	// inventory job info a o0 p0
	InventoryJobInfoAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *InventoryJobInfoAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The execution status of the job scheduled.
		//
		// Read Only: true
		// Enum: [Unknown Scheduled Completed Error]
		ExecutionStatus string `json:"ExecutionStatus,omitempty"`

		// The name of the job scheduled.
		//
		// Read Only: true
		JobName string `json:"JobName,omitempty"`

		// The last processed time of the job scheduled.
		//
		// Read Only: true
		// Format: date-time
		LastProcessedTime strfmt.DateTime `json:"LastProcessedTime,omitempty"`

		// The last scheduled time of the job scheduled.
		//
		// Read Only: true
		// Format: date-time
		LastScheduledTime strfmt.DateTime `json:"LastScheduledTime,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The propertie of the job scheduled.
		//
		// Read Only: true
		Properties []string `json:"Properties"`

		// regex
		// Read Only: true
		Regex []string `json:"Regex"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv InventoryJobInfoAO0P0

	rcv.ExecutionStatus = stage1.ExecutionStatus

	rcv.JobName = stage1.JobName

	rcv.LastProcessedTime = stage1.LastProcessedTime

	rcv.LastScheduledTime = stage1.LastScheduledTime

	rcv.ObjectType = stage1.ObjectType

	rcv.Properties = stage1.Properties

	rcv.Regex = stage1.Regex

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ExecutionStatus")

	delete(stage2, "JobName")

	delete(stage2, "LastProcessedTime")

	delete(stage2, "LastScheduledTime")

	delete(stage2, "ObjectType")

	delete(stage2, "Properties")

	delete(stage2, "Regex")

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
		m.InventoryJobInfoAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m InventoryJobInfoAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The execution status of the job scheduled.
		//
		// Read Only: true
		// Enum: [Unknown Scheduled Completed Error]
		ExecutionStatus string `json:"ExecutionStatus,omitempty"`

		// The name of the job scheduled.
		//
		// Read Only: true
		JobName string `json:"JobName,omitempty"`

		// The last processed time of the job scheduled.
		//
		// Read Only: true
		// Format: date-time
		LastProcessedTime strfmt.DateTime `json:"LastProcessedTime,omitempty"`

		// The last scheduled time of the job scheduled.
		//
		// Read Only: true
		// Format: date-time
		LastScheduledTime strfmt.DateTime `json:"LastScheduledTime,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The propertie of the job scheduled.
		//
		// Read Only: true
		Properties []string `json:"Properties"`

		// regex
		// Read Only: true
		Regex []string `json:"Regex"`
	}

	stage1.ExecutionStatus = m.ExecutionStatus

	stage1.JobName = m.JobName

	stage1.LastProcessedTime = m.LastProcessedTime

	stage1.LastScheduledTime = m.LastScheduledTime

	stage1.ObjectType = m.ObjectType

	stage1.Properties = m.Properties

	stage1.Regex = m.Regex

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.InventoryJobInfoAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.InventoryJobInfoAO0P0)
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

// Validate validates this inventory job info a o0 p0
func (m *InventoryJobInfoAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastProcessedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastScheduledTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var inventoryJobInfoAO0P0TypeExecutionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Scheduled","Completed","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryJobInfoAO0P0TypeExecutionStatusPropEnum = append(inventoryJobInfoAO0P0TypeExecutionStatusPropEnum, v)
	}
}

const (

	// InventoryJobInfoAO0P0ExecutionStatusUnknown captures enum value "Unknown"
	InventoryJobInfoAO0P0ExecutionStatusUnknown string = "Unknown"

	// InventoryJobInfoAO0P0ExecutionStatusScheduled captures enum value "Scheduled"
	InventoryJobInfoAO0P0ExecutionStatusScheduled string = "Scheduled"

	// InventoryJobInfoAO0P0ExecutionStatusCompleted captures enum value "Completed"
	InventoryJobInfoAO0P0ExecutionStatusCompleted string = "Completed"

	// InventoryJobInfoAO0P0ExecutionStatusError captures enum value "Error"
	InventoryJobInfoAO0P0ExecutionStatusError string = "Error"
)

// prop value enum
func (m *InventoryJobInfoAO0P0) validateExecutionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, inventoryJobInfoAO0P0TypeExecutionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InventoryJobInfoAO0P0) validateExecutionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutionStatusEnum("ExecutionStatus", "body", m.ExecutionStatus); err != nil {
		return err
	}

	return nil
}

func (m *InventoryJobInfoAO0P0) validateLastProcessedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastProcessedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("LastProcessedTime", "body", "date-time", m.LastProcessedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InventoryJobInfoAO0P0) validateLastScheduledTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastScheduledTime) { // not required
		return nil
	}

	if err := validate.FormatOf("LastScheduledTime", "body", "date-time", m.LastScheduledTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryJobInfoAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryJobInfoAO0P0) UnmarshalBinary(b []byte) error {
	var res InventoryJobInfoAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
