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

// HyperflexNamedVlan Hyperflex:Named Vlan
//
// A VLAN with a name and ID.
//
// Named VLANs are used for defining the network and iSCSI external storage policies for the cluster.
//
// swagger:model hyperflexNamedVlan
type HyperflexNamedVlan struct {
	HyperflexNamedVlanAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexNamedVlan) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexNamedVlanAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexNamedVlanAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexNamedVlan) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexNamedVlanAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex named vlan
func (m *HyperflexNamedVlan) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexNamedVlanAO0P0
	if err := m.HyperflexNamedVlanAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexNamedVlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexNamedVlan) UnmarshalBinary(b []byte) error {
	var res HyperflexNamedVlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexNamedVlanAO0P0 hyperflex named vlan a o0 p0
// swagger:model HyperflexNamedVlanAO0P0
type HyperflexNamedVlanAO0P0 struct {

	// The name of the VLAN.
	//
	// The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
	//
	//
	Name string `json:"Name,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The ID of the named VLAN. An ID of 0 means the traffic is untagged.
	//
	// The ID can be any number between 0 and 4095, inclusive.
	//
	//
	VlanID int64 `json:"VlanId,omitempty"`

	// hyperflex named vlan a o0 p0
	HyperflexNamedVlanAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexNamedVlanAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The name of the VLAN.
		//
		// The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
		//
		//
		Name string `json:"Name,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The ID of the named VLAN. An ID of 0 means the traffic is untagged.
		//
		// The ID can be any number between 0 and 4095, inclusive.
		//
		//
		VlanID int64 `json:"VlanId,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexNamedVlanAO0P0

	rcv.Name = stage1.Name

	rcv.ObjectType = stage1.ObjectType

	rcv.VlanID = stage1.VlanID

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Name")

	delete(stage2, "ObjectType")

	delete(stage2, "VlanId")

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
		m.HyperflexNamedVlanAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexNamedVlanAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The name of the VLAN.
		//
		// The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
		//
		//
		Name string `json:"Name,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The ID of the named VLAN. An ID of 0 means the traffic is untagged.
		//
		// The ID can be any number between 0 and 4095, inclusive.
		//
		//
		VlanID int64 `json:"VlanId,omitempty"`
	}

	stage1.Name = m.Name

	stage1.ObjectType = m.ObjectType

	stage1.VlanID = m.VlanID

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexNamedVlanAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexNamedVlanAO0P0)
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

// Validate validates this hyperflex named vlan a o0 p0
func (m *HyperflexNamedVlanAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexNamedVlanAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexNamedVlanAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexNamedVlanAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
