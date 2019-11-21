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

// SdcardPartition Sdcard:Partition
//
// This adds and configures operating system partitions.
//
// swagger:model sdcardPartition
type SdcardPartition struct {
	SdcardPartitionAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SdcardPartition) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SdcardPartitionAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SdcardPartitionAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SdcardPartition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SdcardPartitionAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sdcard partition
func (m *SdcardPartition) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SdcardPartitionAO0P0
	if err := m.SdcardPartitionAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SdcardPartition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdcardPartition) UnmarshalBinary(b []byte) error {
	var res SdcardPartition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SdcardPartitionAO0P0 sdcard partition a o0 p0
// swagger:model SdcardPartitionAO0P0
type SdcardPartitionAO0P0 struct {

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// This specifies the type of the partition. Allowed values are OS, Utility.
	//
	// Enum: [OS Utility]
	Type *string `json:"Type,omitempty"`

	// Collection of available virtual drives.
	//
	VirtualDrives []*SdcardVirtualDrive `json:"VirtualDrives"`

	// sdcard partition a o0 p0
	SdcardPartitionAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *SdcardPartitionAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// This specifies the type of the partition. Allowed values are OS, Utility.
		//
		// Enum: [OS Utility]
		Type *string `json:"Type,omitempty"`

		// Collection of available virtual drives.
		//
		VirtualDrives []*SdcardVirtualDrive `json:"VirtualDrives"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv SdcardPartitionAO0P0

	rcv.ObjectType = stage1.ObjectType

	rcv.Type = stage1.Type

	rcv.VirtualDrives = stage1.VirtualDrives

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ObjectType")

	delete(stage2, "Type")

	delete(stage2, "VirtualDrives")

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
		m.SdcardPartitionAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m SdcardPartitionAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// This specifies the type of the partition. Allowed values are OS, Utility.
		//
		// Enum: [OS Utility]
		Type *string `json:"Type,omitempty"`

		// Collection of available virtual drives.
		//
		VirtualDrives []*SdcardVirtualDrive `json:"VirtualDrives"`
	}

	stage1.ObjectType = m.ObjectType

	stage1.Type = m.Type

	stage1.VirtualDrives = m.VirtualDrives

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.SdcardPartitionAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.SdcardPartitionAO0P0)
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

// Validate validates this sdcard partition a o0 p0
func (m *SdcardPartitionAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
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

var sdcardPartitionAO0P0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OS","Utility"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdcardPartitionAO0P0TypeTypePropEnum = append(sdcardPartitionAO0P0TypeTypePropEnum, v)
	}
}

const (

	// SdcardPartitionAO0P0TypeOS captures enum value "OS"
	SdcardPartitionAO0P0TypeOS string = "OS"

	// SdcardPartitionAO0P0TypeUtility captures enum value "Utility"
	SdcardPartitionAO0P0TypeUtility string = "Utility"
)

// prop value enum
func (m *SdcardPartitionAO0P0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdcardPartitionAO0P0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdcardPartitionAO0P0) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SdcardPartitionAO0P0) validateVirtualDrives(formats strfmt.Registry) error {

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
func (m *SdcardPartitionAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdcardPartitionAO0P0) UnmarshalBinary(b []byte) error {
	var res SdcardPartitionAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
