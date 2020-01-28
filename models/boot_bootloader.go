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

// BootBootloader Boot:Bootloader
//
// Lists the bootloader properties that can be overriden for boot from Local disk and SAN boot.
//
// swagger:model bootBootloader
type BootBootloader struct {
	MoBaseComplexType

	BootBootloaderAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BootBootloader) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 BootBootloaderAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.BootBootloaderAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BootBootloader) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.BootBootloaderAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this boot bootloader
func (m *BootBootloader) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BootBootloaderAO1P1
	if err := m.BootBootloaderAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BootBootloader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BootBootloader) UnmarshalBinary(b []byte) error {
	var res BootBootloader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BootBootloaderAO1P1 boot bootloader a o1 p1
// swagger:model BootBootloaderAO1P1
type BootBootloaderAO1P1 struct {

	// Carries more information about the bootloader.
	//
	Description string `json:"Description,omitempty"`

	// Name of the bootloader.
	//
	Name string `json:"Name,omitempty"`

	// Path to the bootloader image.
	//
	Path string `json:"Path,omitempty"`

	// boot bootloader a o1 p1
	BootBootloaderAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *BootBootloaderAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Carries more information about the bootloader.
		//
		Description string `json:"Description,omitempty"`

		// Name of the bootloader.
		//
		Name string `json:"Name,omitempty"`

		// Path to the bootloader image.
		//
		Path string `json:"Path,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv BootBootloaderAO1P1

	rcv.Description = stage1.Description

	rcv.Name = stage1.Name

	rcv.Path = stage1.Path

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Description")

	delete(stage2, "Name")

	delete(stage2, "Path")

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
		m.BootBootloaderAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m BootBootloaderAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Carries more information about the bootloader.
		//
		Description string `json:"Description,omitempty"`

		// Name of the bootloader.
		//
		Name string `json:"Name,omitempty"`

		// Path to the bootloader image.
		//
		Path string `json:"Path,omitempty"`
	}

	stage1.Description = m.Description

	stage1.Name = m.Name

	stage1.Path = m.Path

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.BootBootloaderAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.BootBootloaderAO1P1)
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

// Validate validates this boot bootloader a o1 p1
func (m *BootBootloaderAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BootBootloaderAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BootBootloaderAO1P1) UnmarshalBinary(b []byte) error {
	var res BootBootloaderAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
