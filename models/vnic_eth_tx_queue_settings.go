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

// VnicEthTxQueueSettings Queue Settings
//
// Transmit Queue resource settings.
//
// swagger:model vnicEthTxQueueSettings
type VnicEthTxQueueSettings struct {
	VnicEthTxQueueSettingsAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicEthTxQueueSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VnicEthTxQueueSettingsAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VnicEthTxQueueSettingsAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicEthTxQueueSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.VnicEthTxQueueSettingsAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic eth tx queue settings
func (m *VnicEthTxQueueSettings) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VnicEthTxQueueSettingsAO0P0
	if err := m.VnicEthTxQueueSettingsAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicEthTxQueueSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicEthTxQueueSettings) UnmarshalBinary(b []byte) error {
	var res VnicEthTxQueueSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VnicEthTxQueueSettingsAO0P0 vnic eth tx queue settings a o0 p0
// swagger:model VnicEthTxQueueSettingsAO0P0
type VnicEthTxQueueSettingsAO0P0 struct {

	// The number of queue resources to allocate.
	//
	Count int64 `json:"Count,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The number of descriptors in each queue.
	//
	RingSize int64 `json:"RingSize,omitempty"`

	// vnic eth tx queue settings a o0 p0
	VnicEthTxQueueSettingsAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VnicEthTxQueueSettingsAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The number of queue resources to allocate.
		//
		Count int64 `json:"Count,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The number of descriptors in each queue.
		//
		RingSize int64 `json:"RingSize,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VnicEthTxQueueSettingsAO0P0

	rcv.Count = stage1.Count

	rcv.ObjectType = stage1.ObjectType

	rcv.RingSize = stage1.RingSize

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Count")

	delete(stage2, "ObjectType")

	delete(stage2, "RingSize")

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
		m.VnicEthTxQueueSettingsAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VnicEthTxQueueSettingsAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The number of queue resources to allocate.
		//
		Count int64 `json:"Count,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The number of descriptors in each queue.
		//
		RingSize int64 `json:"RingSize,omitempty"`
	}

	stage1.Count = m.Count

	stage1.ObjectType = m.ObjectType

	stage1.RingSize = m.RingSize

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VnicEthTxQueueSettingsAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VnicEthTxQueueSettingsAO0P0)
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

// Validate validates this vnic eth tx queue settings a o0 p0
func (m *VnicEthTxQueueSettingsAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VnicEthTxQueueSettingsAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicEthTxQueueSettingsAO0P0) UnmarshalBinary(b []byte) error {
	var res VnicEthTxQueueSettingsAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
