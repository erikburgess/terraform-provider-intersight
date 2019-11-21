// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FirmwareEula Firmware:Eula
//
// EULA (End User License Agreement) acceptance status for an account to use cisco.com, to download software.
//
// swagger:model firmwareEula
type FirmwareEula struct {
	MoBaseMo

	// EULA acceptance status for the account.
	//
	// Read Only: true
	Accepted *bool `json:"Accepted,omitempty"`

	// The account associated with this EULA acceptance status.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// EULA acceptance form content provided by cisco.com.
	//
	// Read Only: true
	Content string `json:"Content,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FirmwareEula) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Accepted *bool `json:"Accepted,omitempty"`

		Account *IamAccountRef `json:"Account,omitempty"`

		Content string `json:"Content,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Accepted = dataAO1.Accepted

	m.Account = dataAO1.Account

	m.Content = dataAO1.Content

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FirmwareEula) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Accepted *bool `json:"Accepted,omitempty"`

		Account *IamAccountRef `json:"Account,omitempty"`

		Content string `json:"Content,omitempty"`
	}

	dataAO1.Accepted = m.Accepted

	dataAO1.Account = m.Account

	dataAO1.Content = m.Content

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this firmware eula
func (m *FirmwareEula) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareEula) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareEula) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareEula) UnmarshalBinary(b []byte) error {
	var res FirmwareEula
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
