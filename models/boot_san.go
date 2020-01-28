// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BootSan SAN Boot
//
// Device type used when booting from SAN Boot device.
//
// swagger:model bootSan
type BootSan struct {
	BootDeviceBase

	// Details of the bootloader to be used during SAN boot.
	//
	Bootloader *BootBootloader `json:"Bootloader,omitempty"`

	// The Logical Unit Number (LUN) of the device.
	//
	Lun int64 `json:"Lun,omitempty"`

	// Slot ID of the device. Supported values are ( 1 - 255, "MLOM", "L1", "L2" ).
	//
	Slot string `json:"Slot,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BootSan) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BootDeviceBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BootDeviceBase = aO0

	// AO1
	var dataAO1 struct {
		Bootloader *BootBootloader `json:"Bootloader,omitempty"`

		Lun int64 `json:"Lun,omitempty"`

		Slot string `json:"Slot,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Bootloader = dataAO1.Bootloader

	m.Lun = dataAO1.Lun

	m.Slot = dataAO1.Slot

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BootSan) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BootDeviceBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Bootloader *BootBootloader `json:"Bootloader,omitempty"`

		Lun int64 `json:"Lun,omitempty"`

		Slot string `json:"Slot,omitempty"`
	}

	dataAO1.Bootloader = m.Bootloader

	dataAO1.Lun = m.Lun

	dataAO1.Slot = m.Slot

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this boot san
func (m *BootSan) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BootDeviceBase
	if err := m.BootDeviceBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootloader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BootSan) validateBootloader(formats strfmt.Registry) error {

	if swag.IsZero(m.Bootloader) { // not required
		return nil
	}

	if m.Bootloader != nil {
		if err := m.Bootloader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Bootloader")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BootSan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BootSan) UnmarshalBinary(b []byte) error {
	var res BootSan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
