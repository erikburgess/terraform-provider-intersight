// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageSpanGroup Storage:Span Group
//
// This models a single span group of disks in a RAID group.
//
// swagger:model storageSpanGroup
type StorageSpanGroup struct {
	MoBaseComplexType

	StorageSpanGroupAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StorageSpanGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 StorageSpanGroupAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.StorageSpanGroupAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StorageSpanGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.StorageSpanGroupAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this storage span group
func (m *StorageSpanGroup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with StorageSpanGroupAO1P1
	if err := m.StorageSpanGroupAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StorageSpanGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSpanGroup) UnmarshalBinary(b []byte) error {
	var res StorageSpanGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StorageSpanGroupAO1P1 storage span group a o1 p1
//
// swagger:model StorageSpanGroupAO1P1
type StorageSpanGroupAO1P1 struct {

	// Collection of local disks that are part of this span group. The minimum number of disks needed in a span group varies based on RAID level. Raid0 requires at least one disk, Raid1 and Raid10 requires at least 2 and in multiples of 2, Raid5 Raid50 Raid6 and Raid60 require at least 3 disks in a span group.
	Disks []*StorageLocalDisk `json:"Disks"`

	// storage span group a o1 p1
	StorageSpanGroupAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *StorageSpanGroupAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Collection of local disks that are part of this span group. The minimum number of disks needed in a span group varies based on RAID level. Raid0 requires at least one disk, Raid1 and Raid10 requires at least 2 and in multiples of 2, Raid5 Raid50 Raid6 and Raid60 require at least 3 disks in a span group.
		Disks []*StorageLocalDisk `json:"Disks"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StorageSpanGroupAO1P1

	rcv.Disks = stage1.Disks
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Disks")
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
		m.StorageSpanGroupAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m StorageSpanGroupAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Collection of local disks that are part of this span group. The minimum number of disks needed in a span group varies based on RAID level. Raid0 requires at least one disk, Raid1 and Raid10 requires at least 2 and in multiples of 2, Raid5 Raid50 Raid6 and Raid60 require at least 3 disks in a span group.
		Disks []*StorageLocalDisk `json:"Disks"`
	}

	stage1.Disks = m.Disks

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.StorageSpanGroupAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.StorageSpanGroupAO1P1)
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

// Validate validates this storage span group a o1 p1
func (m *StorageSpanGroupAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageSpanGroupAO1P1) validateDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageSpanGroupAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSpanGroupAO1P1) UnmarshalBinary(b []byte) error {
	var res StorageSpanGroupAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
