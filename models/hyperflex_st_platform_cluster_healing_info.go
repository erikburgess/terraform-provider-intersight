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

// HyperflexStPlatformClusterHealingInfo Hyperflex:St Platform Cluster Healing Info
// swagger:model hyperflexStPlatformClusterHealingInfo
type HyperflexStPlatformClusterHealingInfo struct {
	HyperflexStPlatformClusterHealingInfoAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexStPlatformClusterHealingInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexStPlatformClusterHealingInfoAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexStPlatformClusterHealingInfoAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexStPlatformClusterHealingInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexStPlatformClusterHealingInfoAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex st platform cluster healing info
func (m *HyperflexStPlatformClusterHealingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexStPlatformClusterHealingInfoAO0P0
	if err := m.HyperflexStPlatformClusterHealingInfoAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexStPlatformClusterHealingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexStPlatformClusterHealingInfo) UnmarshalBinary(b []byte) error {
	var res HyperflexStPlatformClusterHealingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexStPlatformClusterHealingInfoAO0P0 hyperflex st platform cluster healing info a o0 p0
// swagger:model HyperflexStPlatformClusterHealingInfoAO0P0
type HyperflexStPlatformClusterHealingInfoAO0P0 struct {

	// estimated completion time in seconds
	// Read Only: true
	EstimatedCompletionTimeInSeconds int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`

	// in progress
	// Read Only: true
	InProgress *bool `json:"InProgress,omitempty"`

	// messages
	// Read Only: true
	Messages []string `json:"Messages"`

	// messages iterator
	// Read Only: true
	MessagesIterator interface{} `json:"MessagesIterator,omitempty"`

	// messages size
	// Read Only: true
	MessagesSize int64 `json:"MessagesSize,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// percent complete
	// Read Only: true
	PercentComplete int64 `json:"PercentComplete,omitempty"`

	// hyperflex st platform cluster healing info a o0 p0
	HyperflexStPlatformClusterHealingInfoAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexStPlatformClusterHealingInfoAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// estimated completion time in seconds
		// Read Only: true
		EstimatedCompletionTimeInSeconds int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`

		// in progress
		// Read Only: true
		InProgress *bool `json:"InProgress,omitempty"`

		// messages
		// Read Only: true
		Messages []string `json:"Messages"`

		// messages iterator
		// Read Only: true
		MessagesIterator interface{} `json:"MessagesIterator,omitempty"`

		// messages size
		// Read Only: true
		MessagesSize int64 `json:"MessagesSize,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// percent complete
		// Read Only: true
		PercentComplete int64 `json:"PercentComplete,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexStPlatformClusterHealingInfoAO0P0

	rcv.EstimatedCompletionTimeInSeconds = stage1.EstimatedCompletionTimeInSeconds

	rcv.InProgress = stage1.InProgress

	rcv.Messages = stage1.Messages

	rcv.MessagesIterator = stage1.MessagesIterator

	rcv.MessagesSize = stage1.MessagesSize

	rcv.ObjectType = stage1.ObjectType

	rcv.PercentComplete = stage1.PercentComplete

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "EstimatedCompletionTimeInSeconds")

	delete(stage2, "InProgress")

	delete(stage2, "Messages")

	delete(stage2, "MessagesIterator")

	delete(stage2, "MessagesSize")

	delete(stage2, "ObjectType")

	delete(stage2, "PercentComplete")

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
		m.HyperflexStPlatformClusterHealingInfoAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexStPlatformClusterHealingInfoAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// estimated completion time in seconds
		// Read Only: true
		EstimatedCompletionTimeInSeconds int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`

		// in progress
		// Read Only: true
		InProgress *bool `json:"InProgress,omitempty"`

		// messages
		// Read Only: true
		Messages []string `json:"Messages"`

		// messages iterator
		// Read Only: true
		MessagesIterator interface{} `json:"MessagesIterator,omitempty"`

		// messages size
		// Read Only: true
		MessagesSize int64 `json:"MessagesSize,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// percent complete
		// Read Only: true
		PercentComplete int64 `json:"PercentComplete,omitempty"`
	}

	stage1.EstimatedCompletionTimeInSeconds = m.EstimatedCompletionTimeInSeconds

	stage1.InProgress = m.InProgress

	stage1.Messages = m.Messages

	stage1.MessagesIterator = m.MessagesIterator

	stage1.MessagesSize = m.MessagesSize

	stage1.ObjectType = m.ObjectType

	stage1.PercentComplete = m.PercentComplete

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexStPlatformClusterHealingInfoAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexStPlatformClusterHealingInfoAO0P0)
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

// Validate validates this hyperflex st platform cluster healing info a o0 p0
func (m *HyperflexStPlatformClusterHealingInfoAO0P0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexStPlatformClusterHealingInfoAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexStPlatformClusterHealingInfoAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexStPlatformClusterHealingInfoAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
