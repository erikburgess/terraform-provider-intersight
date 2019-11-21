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

// HyperflexAppSettingConstraint Hyperflex:App Setting Constraint
//
// A constraint that can be used to qualify an application setting.
//
// swagger:model hyperflexAppSettingConstraint
type HyperflexAppSettingConstraint struct {
	HyperflexAppSettingConstraintAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexAppSettingConstraint) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexAppSettingConstraintAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexAppSettingConstraintAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexAppSettingConstraint) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexAppSettingConstraintAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex app setting constraint
func (m *HyperflexAppSettingConstraint) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexAppSettingConstraintAO0P0
	if err := m.HyperflexAppSettingConstraintAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexAppSettingConstraint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexAppSettingConstraint) UnmarshalBinary(b []byte) error {
	var res HyperflexAppSettingConstraint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexAppSettingConstraintAO0P0 hyperflex app setting constraint a o0 p0
// swagger:model HyperflexAppSettingConstraintAO0P0
type HyperflexAppSettingConstraintAO0P0 struct {

	// The supported HyperFlex Data Platform version in regex format.
	//
	HxdpVersion string `json:"HxdpVersion,omitempty"`

	// The hypervisor type for the HyperFlex cluster.
	//
	// Enum: [ESXi]
	HypervisorType *string `json:"HypervisorType,omitempty"`

	// The supported management platform for the HyperFlex Cluster.
	//
	// Enum: [FI EDGE]
	MgmtPlatform *string `json:"MgmtPlatform,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The supported server models in regex format.
	//
	ServerModel string `json:"ServerModel,omitempty"`

	// hyperflex app setting constraint a o0 p0
	HyperflexAppSettingConstraintAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexAppSettingConstraintAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The supported HyperFlex Data Platform version in regex format.
		//
		HxdpVersion string `json:"HxdpVersion,omitempty"`

		// The hypervisor type for the HyperFlex cluster.
		//
		// Enum: [ESXi]
		HypervisorType *string `json:"HypervisorType,omitempty"`

		// The supported management platform for the HyperFlex Cluster.
		//
		// Enum: [FI EDGE]
		MgmtPlatform *string `json:"MgmtPlatform,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The supported server models in regex format.
		//
		ServerModel string `json:"ServerModel,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexAppSettingConstraintAO0P0

	rcv.HxdpVersion = stage1.HxdpVersion

	rcv.HypervisorType = stage1.HypervisorType

	rcv.MgmtPlatform = stage1.MgmtPlatform

	rcv.ObjectType = stage1.ObjectType

	rcv.ServerModel = stage1.ServerModel

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "HxdpVersion")

	delete(stage2, "HypervisorType")

	delete(stage2, "MgmtPlatform")

	delete(stage2, "ObjectType")

	delete(stage2, "ServerModel")

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
		m.HyperflexAppSettingConstraintAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexAppSettingConstraintAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The supported HyperFlex Data Platform version in regex format.
		//
		HxdpVersion string `json:"HxdpVersion,omitempty"`

		// The hypervisor type for the HyperFlex cluster.
		//
		// Enum: [ESXi]
		HypervisorType *string `json:"HypervisorType,omitempty"`

		// The supported management platform for the HyperFlex Cluster.
		//
		// Enum: [FI EDGE]
		MgmtPlatform *string `json:"MgmtPlatform,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The supported server models in regex format.
		//
		ServerModel string `json:"ServerModel,omitempty"`
	}

	stage1.HxdpVersion = m.HxdpVersion

	stage1.HypervisorType = m.HypervisorType

	stage1.MgmtPlatform = m.MgmtPlatform

	stage1.ObjectType = m.ObjectType

	stage1.ServerModel = m.ServerModel

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexAppSettingConstraintAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexAppSettingConstraintAO0P0)
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

// Validate validates this hyperflex app setting constraint a o0 p0
func (m *HyperflexAppSettingConstraintAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHypervisorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMgmtPlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hyperflexAppSettingConstraintAO0P0TypeHypervisorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ESXi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexAppSettingConstraintAO0P0TypeHypervisorTypePropEnum = append(hyperflexAppSettingConstraintAO0P0TypeHypervisorTypePropEnum, v)
	}
}

const (

	// HyperflexAppSettingConstraintAO0P0HypervisorTypeESXi captures enum value "ESXi"
	HyperflexAppSettingConstraintAO0P0HypervisorTypeESXi string = "ESXi"
)

// prop value enum
func (m *HyperflexAppSettingConstraintAO0P0) validateHypervisorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexAppSettingConstraintAO0P0TypeHypervisorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexAppSettingConstraintAO0P0) validateHypervisorType(formats strfmt.Registry) error {

	if swag.IsZero(m.HypervisorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateHypervisorTypeEnum("HypervisorType", "body", *m.HypervisorType); err != nil {
		return err
	}

	return nil
}

var hyperflexAppSettingConstraintAO0P0TypeMgmtPlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FI","EDGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexAppSettingConstraintAO0P0TypeMgmtPlatformPropEnum = append(hyperflexAppSettingConstraintAO0P0TypeMgmtPlatformPropEnum, v)
	}
}

const (

	// HyperflexAppSettingConstraintAO0P0MgmtPlatformFI captures enum value "FI"
	HyperflexAppSettingConstraintAO0P0MgmtPlatformFI string = "FI"

	// HyperflexAppSettingConstraintAO0P0MgmtPlatformEDGE captures enum value "EDGE"
	HyperflexAppSettingConstraintAO0P0MgmtPlatformEDGE string = "EDGE"
)

// prop value enum
func (m *HyperflexAppSettingConstraintAO0P0) validateMgmtPlatformEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexAppSettingConstraintAO0P0TypeMgmtPlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexAppSettingConstraintAO0P0) validateMgmtPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.MgmtPlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateMgmtPlatformEnum("MgmtPlatform", "body", *m.MgmtPlatform); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexAppSettingConstraintAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexAppSettingConstraintAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexAppSettingConstraintAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
