// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SoftwarerepositoryAuthorization Softwarerepository:Authorization
//
// Consent that a user has provided to Intersight for contacting an external software repository such as cisco.com, on the user account's behalf.
//
// swagger:model softwarerepositoryAuthorization
type SoftwarerepositoryAuthorization struct {
	MoBaseMo

	// The account, on behalf of which, the authorization is being provided.
	//
	// Read Only: true
	Account *IamAccountRef `json:"Account,omitempty"`

	// is password set
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

	// is user Id set
	IsUserIDSet *bool `json:"IsUserIdSet,omitempty"`

	// The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	//
	Password string `json:"Password,omitempty"`

	// The external repository for which this authorization has been provided. The only supported repository today is cisco.com.
	//
	// Enum: [Cisco IntersightCloud LocalMachine NetworkShare]
	RepositoryType *string `json:"RepositoryType,omitempty"`

	// The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	//
	UserID string `json:"UserId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SoftwarerepositoryAuthorization) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		IsUserIDSet *bool `json:"IsUserIdSet,omitempty"`

		Password string `json:"Password,omitempty"`

		RepositoryType *string `json:"RepositoryType,omitempty"`

		UserID string `json:"UserId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.IsPasswordSet = dataAO1.IsPasswordSet

	m.IsUserIDSet = dataAO1.IsUserIDSet

	m.Password = dataAO1.Password

	m.RepositoryType = dataAO1.RepositoryType

	m.UserID = dataAO1.UserID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SoftwarerepositoryAuthorization) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		IsUserIDSet *bool `json:"IsUserIdSet,omitempty"`

		Password string `json:"Password,omitempty"`

		RepositoryType *string `json:"RepositoryType,omitempty"`

		UserID string `json:"UserId,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.IsPasswordSet = m.IsPasswordSet

	dataAO1.IsUserIDSet = m.IsUserIDSet

	dataAO1.Password = m.Password

	dataAO1.RepositoryType = m.RepositoryType

	dataAO1.UserID = m.UserID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this softwarerepository authorization
func (m *SoftwarerepositoryAuthorization) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwarerepositoryAuthorization) validateAccount(formats strfmt.Registry) error {

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

var softwarerepositoryAuthorizationTypeRepositoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cisco","IntersightCloud","LocalMachine","NetworkShare"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwarerepositoryAuthorizationTypeRepositoryTypePropEnum = append(softwarerepositoryAuthorizationTypeRepositoryTypePropEnum, v)
	}
}

// property enum
func (m *SoftwarerepositoryAuthorization) validateRepositoryTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, softwarerepositoryAuthorizationTypeRepositoryTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SoftwarerepositoryAuthorization) validateRepositoryType(formats strfmt.Registry) error {

	if swag.IsZero(m.RepositoryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRepositoryTypeEnum("RepositoryType", "body", *m.RepositoryType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwarerepositoryAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwarerepositoryAuthorization) UnmarshalBinary(b []byte) error {
	var res SoftwarerepositoryAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
