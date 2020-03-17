// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommCredential Comm:Credential
//
// Base Credential for all the Managed Devices.
//
// swagger:model commCredential
type CommCredential struct {
	MoBaseComplexType

	CommCredentialAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommCredential) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 CommCredentialAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommCredentialAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommCredential) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.CommCredentialAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comm credential
func (m *CommCredential) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CommCredentialAO1P1
	if err := m.CommCredentialAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CommCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommCredential) UnmarshalBinary(b []byte) error {
	var res CommCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommCredentialAO1P1 comm credential a o1 p1
//
// swagger:model CommCredentialAO1P1
type CommCredentialAO1P1 struct {

	// is password set
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

	// Password for the Managed Device.
	Password string `json:"Password,omitempty"`

	// Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements.
	Username string `json:"Username,omitempty"`

	// comm credential a o1 p1
	CommCredentialAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *CommCredentialAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// is password set
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		// Password for the Managed Device.
		Password string `json:"Password,omitempty"`

		// Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements.
		Username string `json:"Username,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv CommCredentialAO1P1

	rcv.IsPasswordSet = stage1.IsPasswordSet
	rcv.Password = stage1.Password
	rcv.Username = stage1.Username
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "IsPasswordSet")
	delete(stage2, "Password")
	delete(stage2, "Username")
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
		m.CommCredentialAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m CommCredentialAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// is password set
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		// Password for the Managed Device.
		Password string `json:"Password,omitempty"`

		// Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements.
		Username string `json:"Username,omitempty"`
	}

	stage1.IsPasswordSet = m.IsPasswordSet
	stage1.Password = m.Password
	stage1.Username = m.Username

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.CommCredentialAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.CommCredentialAO1P1)
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

// Validate validates this comm credential a o1 p1
func (m *CommCredentialAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommCredentialAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommCredentialAO1P1) UnmarshalBinary(b []byte) error {
	var res CommCredentialAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
