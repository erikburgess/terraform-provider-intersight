// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestcryptReadOnlyUser Testcrypt:Read Only User
//
// Mo with collection of complextype with secure properties.
//
// swagger:model testcryptReadOnlyUser
type TestcryptReadOnlyUser struct {
	MoBaseMo

	// The account linked to ReadOnlyUser.
	//
	Account *IamAccountRef `json:"Account,omitempty"`

	// The property, rouser is a collection of complex type, only for testing.
	//
	Rouser []*TestcryptUser `json:"Rouser"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TestcryptReadOnlyUser) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Rouser []*TestcryptUser `json:"Rouser"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.Rouser = dataAO1.Rouser

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TestcryptReadOnlyUser) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Rouser []*TestcryptUser `json:"Rouser"`
	}

	dataAO1.Account = m.Account

	dataAO1.Rouser = m.Rouser

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this testcrypt read only user
func (m *TestcryptReadOnlyUser) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestcryptReadOnlyUser) validateAccount(formats strfmt.Registry) error {

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

func (m *TestcryptReadOnlyUser) validateRouser(formats strfmt.Registry) error {

	if swag.IsZero(m.Rouser) { // not required
		return nil
	}

	for i := 0; i < len(m.Rouser); i++ {
		if swag.IsZero(m.Rouser[i]) { // not required
			continue
		}

		if m.Rouser[i] != nil {
			if err := m.Rouser[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Rouser" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestcryptReadOnlyUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestcryptReadOnlyUser) UnmarshalBinary(b []byte) error {
	var res TestcryptReadOnlyUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}