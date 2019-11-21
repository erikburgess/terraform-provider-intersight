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
)

// HyperflexHxUuIDDt Hyperflex:Hx Uu Id Dt
// swagger:model hyperflexHxUuIdDt
type HyperflexHxUuIDDt struct {
	HyperflexHxUuIDDtAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHxUuIDDt) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HyperflexHxUuIDDtAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HyperflexHxUuIDDtAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHxUuIDDt) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HyperflexHxUuIDDtAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex hx uu Id dt
func (m *HyperflexHxUuIDDt) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HyperflexHxUuIDDtAO0P0
	if err := m.HyperflexHxUuIDDtAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxUuIDDt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxUuIDDt) UnmarshalBinary(b []byte) error {
	var res HyperflexHxUuIDDt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HyperflexHxUuIDDtAO0P0 hyperflex hx uu ID dt a o0 p0
// swagger:model HyperflexHxUuIDDtAO0P0
type HyperflexHxUuIDDtAO0P0 struct {

	// links
	// Read Only: true
	Links []*HyperflexHxLinkDt `json:"Links"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Uuid
	// Read Only: true
	UUID string `json:"Uuid,omitempty"`

	// hyperflex hx uu ID dt a o0 p0
	HyperflexHxUuIDDtAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HyperflexHxUuIDDtAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// links
		// Read Only: true
		Links []*HyperflexHxLinkDt `json:"Links"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Uuid
		// Read Only: true
		UUID string `json:"Uuid,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HyperflexHxUuIDDtAO0P0

	rcv.Links = stage1.Links

	rcv.ObjectType = stage1.ObjectType

	rcv.UUID = stage1.UUID

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Links")

	delete(stage2, "ObjectType")

	delete(stage2, "Uuid")

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
		m.HyperflexHxUuIDDtAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HyperflexHxUuIDDtAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// links
		// Read Only: true
		Links []*HyperflexHxLinkDt `json:"Links"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Uuid
		// Read Only: true
		UUID string `json:"Uuid,omitempty"`
	}

	stage1.Links = m.Links

	stage1.ObjectType = m.ObjectType

	stage1.UUID = m.UUID

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HyperflexHxUuIDDtAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HyperflexHxUuIDDtAO0P0)
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

// Validate validates this hyperflex hx uu ID dt a o0 p0
func (m *HyperflexHxUuIDDtAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexHxUuIDDtAO0P0) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHxUuIDDtAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHxUuIDDtAO0P0) UnmarshalBinary(b []byte) error {
	var res HyperflexHxUuIDDtAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
