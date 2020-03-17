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
	"github.com/go-openapi/validate"
)

// MoVersionContext Mo:Version Context
//
// VersionContext contains the versioning info for an object.
//
// swagger:model moVersionContext
type MoVersionContext struct {
	MoBaseComplexType

	MoVersionContextAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MoVersionContext) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 MoVersionContextAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.MoVersionContextAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MoVersionContext) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.MoVersionContextAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this mo version context
func (m *MoVersionContext) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with MoVersionContextAO1P1
	if err := m.MoVersionContextAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MoVersionContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoVersionContext) UnmarshalBinary(b []byte) error {
	var res MoVersionContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MoVersionContextAO1P1 mo version context a o1 p1
//
// swagger:model MoVersionContextAO1P1
type MoVersionContextAO1P1 struct {

	// A collection of objects that have reference to this versioned object.
	// The lifecycle of the versioned object is based on the interestedMos list;
	// the versioned object will be purged when interestedMos is empty.
	// Read Only: true
	InterestedMos []*MoMoRef `json:"InterestedMos"`

	// A reference to the original Managed Object.
	// Read Only: true
	RefMo *MoMoRef `json:"RefMo,omitempty"`

	// The time this versioned Managed Object was created.
	// Read Only: true
	// Format: date-time
	Timestamp strfmt.DateTime `json:"Timestamp,omitempty"`

	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	// Read Only: true
	Version string `json:"Version,omitempty"`

	// Specifies type of version. Currently the only supported value is "Configured"
	// that is used to keep track of snapshots of policies and profiles that are intended
	// to be configured to target endpoints.
	// Read Only: true
	// Enum: [Modified Configured Deployed]
	VersionType string `json:"VersionType,omitempty"`

	// mo version context a o1 p1
	MoVersionContextAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *MoVersionContextAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// A collection of objects that have reference to this versioned object.
		// The lifecycle of the versioned object is based on the interestedMos list;
		// the versioned object will be purged when interestedMos is empty.
		// Read Only: true
		InterestedMos []*MoMoRef `json:"InterestedMos"`

		// A reference to the original Managed Object.
		// Read Only: true
		RefMo *MoMoRef `json:"RefMo,omitempty"`

		// The time this versioned Managed Object was created.
		// Read Only: true
		// Format: date-time
		Timestamp strfmt.DateTime `json:"Timestamp,omitempty"`

		// The version of the Managed Object, e.g. an incrementing number or a hash id.
		// Read Only: true
		Version string `json:"Version,omitempty"`

		// Specifies type of version. Currently the only supported value is "Configured"
		// that is used to keep track of snapshots of policies and profiles that are intended
		// to be configured to target endpoints.
		// Read Only: true
		// Enum: [Modified Configured Deployed]
		VersionType string `json:"VersionType,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv MoVersionContextAO1P1

	rcv.InterestedMos = stage1.InterestedMos
	rcv.RefMo = stage1.RefMo
	rcv.Timestamp = stage1.Timestamp
	rcv.Version = stage1.Version
	rcv.VersionType = stage1.VersionType
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "InterestedMos")
	delete(stage2, "RefMo")
	delete(stage2, "Timestamp")
	delete(stage2, "Version")
	delete(stage2, "VersionType")
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
		m.MoVersionContextAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m MoVersionContextAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// A collection of objects that have reference to this versioned object.
		// The lifecycle of the versioned object is based on the interestedMos list;
		// the versioned object will be purged when interestedMos is empty.
		// Read Only: true
		InterestedMos []*MoMoRef `json:"InterestedMos"`

		// A reference to the original Managed Object.
		// Read Only: true
		RefMo *MoMoRef `json:"RefMo,omitempty"`

		// The time this versioned Managed Object was created.
		// Read Only: true
		// Format: date-time
		Timestamp strfmt.DateTime `json:"Timestamp,omitempty"`

		// The version of the Managed Object, e.g. an incrementing number or a hash id.
		// Read Only: true
		Version string `json:"Version,omitempty"`

		// Specifies type of version. Currently the only supported value is "Configured"
		// that is used to keep track of snapshots of policies and profiles that are intended
		// to be configured to target endpoints.
		// Read Only: true
		// Enum: [Modified Configured Deployed]
		VersionType string `json:"VersionType,omitempty"`
	}

	stage1.InterestedMos = m.InterestedMos
	stage1.RefMo = m.RefMo
	stage1.Timestamp = m.Timestamp
	stage1.Version = m.Version
	stage1.VersionType = m.VersionType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.MoVersionContextAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.MoVersionContextAO1P1)
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

// Validate validates this mo version context a o1 p1
func (m *MoVersionContextAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterestedMos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefMo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoVersionContextAO1P1) validateInterestedMos(formats strfmt.Registry) error {

	if swag.IsZero(m.InterestedMos) { // not required
		return nil
	}

	for i := 0; i < len(m.InterestedMos); i++ {
		if swag.IsZero(m.InterestedMos[i]) { // not required
			continue
		}

		if m.InterestedMos[i] != nil {
			if err := m.InterestedMos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InterestedMos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoVersionContextAO1P1) validateRefMo(formats strfmt.Registry) error {

	if swag.IsZero(m.RefMo) { // not required
		return nil
	}

	if m.RefMo != nil {
		if err := m.RefMo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RefMo")
			}
			return err
		}
	}

	return nil
}

func (m *MoVersionContextAO1P1) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("Timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var moVersionContextAO1P1TypeVersionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Modified","Configured","Deployed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moVersionContextAO1P1TypeVersionTypePropEnum = append(moVersionContextAO1P1TypeVersionTypePropEnum, v)
	}
}

const (

	// MoVersionContextAO1P1VersionTypeModified captures enum value "Modified"
	MoVersionContextAO1P1VersionTypeModified string = "Modified"

	// MoVersionContextAO1P1VersionTypeConfigured captures enum value "Configured"
	MoVersionContextAO1P1VersionTypeConfigured string = "Configured"

	// MoVersionContextAO1P1VersionTypeDeployed captures enum value "Deployed"
	MoVersionContextAO1P1VersionTypeDeployed string = "Deployed"
)

// prop value enum
func (m *MoVersionContextAO1P1) validateVersionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moVersionContextAO1P1TypeVersionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoVersionContextAO1P1) validateVersionType(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionTypeEnum("VersionType", "body", m.VersionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoVersionContextAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoVersionContextAO1P1) UnmarshalBinary(b []byte) error {
	var res MoVersionContextAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
