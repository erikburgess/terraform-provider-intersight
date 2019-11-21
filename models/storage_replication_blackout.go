// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StorageReplicationBlackout Storage:Replication Blackout
//
// Range of time at which to suspend replication. System disables replication during this interval.
//
// swagger:model storageReplicationBlackout
type StorageReplicationBlackout struct {
	StorageReplicationBlackoutAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageReplicationBlackout) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StorageReplicationBlackoutAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StorageReplicationBlackoutAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageReplicationBlackout) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.StorageReplicationBlackoutAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage replication blackout
func (m *StorageReplicationBlackout) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StorageReplicationBlackoutAO0P0
	if err := m.StorageReplicationBlackoutAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageReplicationBlackout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageReplicationBlackout) UnmarshalBinary(b []byte) error {
	var res StorageReplicationBlackout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StorageReplicationBlackoutAO0P0 storage replication blackout a o0 p0
// swagger:model StorageReplicationBlackoutAO0P0
type StorageReplicationBlackoutAO0P0 struct {

	// The end time of day for replication blackout window.
	// Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
	//
	//
	// Read Only: true
	End string `json:"End,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication.
	// Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds.
	// The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
	//
	//
	// Read Only: true
	Start string `json:"Start,omitempty"`

	// storage replication blackout a o0 p0
	StorageReplicationBlackoutAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *StorageReplicationBlackoutAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The end time of day for replication blackout window.
		// Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
		//
		//
		// Read Only: true
		End string `json:"End,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication.
		// Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds.
		// The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
		//
		//
		// Read Only: true
		Start string `json:"Start,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StorageReplicationBlackoutAO0P0

	rcv.End = stage1.End

	rcv.ObjectType = stage1.ObjectType

	rcv.Start = stage1.Start

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "End")

	delete(stage2, "ObjectType")

	delete(stage2, "Start")

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
		m.StorageReplicationBlackoutAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m StorageReplicationBlackoutAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The end time of day for replication blackout window.
		// Example: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.
		//
		//
		// Read Only: true
		End string `json:"End,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication.
		// Example: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds.
		// The fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.
		//
		//
		// Read Only: true
		Start string `json:"Start,omitempty"`
	}

	stage1.End = m.End

	stage1.ObjectType = m.ObjectType

	stage1.Start = m.Start

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.StorageReplicationBlackoutAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.StorageReplicationBlackoutAO0P0)
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

// Validate validates this storage replication blackout a o0 p0
func (m *StorageReplicationBlackoutAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageReplicationBlackoutAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageReplicationBlackoutAO0P0) UnmarshalBinary(b []byte) error {
	var res StorageReplicationBlackoutAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
