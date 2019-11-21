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

// HyperflexWwxnPrefixRange Hyperflex:Wwxn Prefix Range
//
// A range of WWxN prefixes.
//
// The range is inclusive and comprised of a start and end WWxN addresses.
// A single address can be specified by setting it as the start and end of the range.
//
// swagger:model hyperflexWwxnPrefixRange
type HyperflexWwxnPrefixRange struct {
	HyperflexWwxnPrefixRangeAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexWwxnPrefixRange) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexWwxnPrefixRangeAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexWwxnPrefixRangeAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexWwxnPrefixRange) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexWwxnPrefixRangeAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex wwxn prefix range
func (m *HyperflexWwxnPrefixRange) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexWwxnPrefixRangeAO0P0
	if err := m.HyperflexWwxnPrefixRangeAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexWwxnPrefixRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexWwxnPrefixRange) UnmarshalBinary(b []byte) error {
	var res HyperflexWwxnPrefixRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexWwxnPrefixRangeAO0P0 hyperflex wwxn prefix range a o0 p0
// swagger:model HyperflexWwxnPrefixRangeAO0P0
type HyperflexWwxnPrefixRangeAO0P0 struct {

	// The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
	//
	EndAddr string `json:"EndAddr,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
	//
	StartAddr string `json:"StartAddr,omitempty"`

	// hyperflex wwxn prefix range a o0 p0
	HyperflexWwxnPrefixRangeAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexWwxnPrefixRangeAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
		//
		EndAddr string `json:"EndAddr,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
		//
		StartAddr string `json:"StartAddr,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexWwxnPrefixRangeAO0P0

	rcv.EndAddr = stage1.EndAddr

	rcv.ObjectType = stage1.ObjectType

	rcv.StartAddr = stage1.StartAddr

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "EndAddr")

	delete(stage2, "ObjectType")

	delete(stage2, "StartAddr")

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
		m.HyperflexWwxnPrefixRangeAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexWwxnPrefixRangeAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
		//
		EndAddr string `json:"EndAddr,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX.
		//
		StartAddr string `json:"StartAddr,omitempty"`
	}

	stage1.EndAddr = m.EndAddr

	stage1.ObjectType = m.ObjectType

	stage1.StartAddr = m.StartAddr

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexWwxnPrefixRangeAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexWwxnPrefixRangeAO0P0)
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

// Validate validates this hyperflex wwxn prefix range a o0 p0
func (m *HyperflexWwxnPrefixRangeAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexWwxnPrefixRangeAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexWwxnPrefixRangeAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexWwxnPrefixRangeAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
