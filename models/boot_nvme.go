// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BootNvme NVMe
//
// Device type used when booting from a NVMe device.
//
// swagger:model bootNvme
type BootNvme struct {
	BootDeviceBase

	BootNvmeAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BootNvme) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BootDeviceBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BootDeviceBase = aO0

	// AO1
	var aO1 BootNvmeAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.BootNvmeAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BootNvme) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BootDeviceBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.BootNvmeAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this boot nvme
func (m *BootNvme) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BootDeviceBase
	if err := m.BootDeviceBase.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BootNvmeAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BootNvme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BootNvme) UnmarshalBinary(b []byte) error {
	var res BootNvme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BootNvmeAllOf1 boot nvme all of1
// swagger:model BootNvmeAllOf1
type BootNvmeAllOf1 interface{}
