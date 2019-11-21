// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IamSession Iam:Session
//
// The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
//
// swagger:model iamSession
type IamSession struct {
	MoBaseMo

	// The accounts and the permissions within each account which a user can select after authentication. After authentication if user has access to multiple permissions, then user and session object are created in onboarding user account and asked to select one of these permissions.
	//
	// Read Only: true
	AccountPermissions []*IamAccountPermissions `json:"AccountPermissions"`

	// The user agent IP address from which the session is launched.
	//
	// Read Only: true
	ClientIPAddress string `json:"ClientIpAddress,omitempty"`

	// Expiration time for the session.
	//
	// Read Only: true
	// Format: date-time
	Expiration strfmt.DateTime `json:"Expiration,omitempty"`

	// Idle time expiration for the session.
	//
	// Read Only: true
	// Format: date-time
	IdleTimeExpiration strfmt.DateTime `json:"IdleTimeExpiration,omitempty"`

	// The client address from which last login is initiated.
	//
	// Read Only: true
	LastLoginClient string `json:"LastLoginClient,omitempty"`

	// The last login time for user.
	//
	// Read Only: true
	// Format: date-time
	LastLoginTime strfmt.DateTime `json:"LastLoginTime,omitempty"`

	// Permissions associated with the web session. Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
	//
	// Read Only: true
	Permission *IamPermissionRef `json:"Permission,omitempty"`

	// A collection of references to the [iam.User](mo://iam.User) Managed Object.
	//
	// When this managed object is deleted, the referenced [iam.User](mo://iam.User) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	User *IamUserRef `json:"User,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamSession) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AccountPermissions []*IamAccountPermissions `json:"AccountPermissions"`

		ClientIPAddress string `json:"ClientIpAddress,omitempty"`

		Expiration strfmt.DateTime `json:"Expiration,omitempty"`

		IdleTimeExpiration strfmt.DateTime `json:"IdleTimeExpiration,omitempty"`

		LastLoginClient string `json:"LastLoginClient,omitempty"`

		LastLoginTime strfmt.DateTime `json:"LastLoginTime,omitempty"`

		Permission *IamPermissionRef `json:"Permission,omitempty"`

		User *IamUserRef `json:"User,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AccountPermissions = dataAO1.AccountPermissions

	m.ClientIPAddress = dataAO1.ClientIPAddress

	m.Expiration = dataAO1.Expiration

	m.IdleTimeExpiration = dataAO1.IdleTimeExpiration

	m.LastLoginClient = dataAO1.LastLoginClient

	m.LastLoginTime = dataAO1.LastLoginTime

	m.Permission = dataAO1.Permission

	m.User = dataAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamSession) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AccountPermissions []*IamAccountPermissions `json:"AccountPermissions"`

		ClientIPAddress string `json:"ClientIpAddress,omitempty"`

		Expiration strfmt.DateTime `json:"Expiration,omitempty"`

		IdleTimeExpiration strfmt.DateTime `json:"IdleTimeExpiration,omitempty"`

		LastLoginClient string `json:"LastLoginClient,omitempty"`

		LastLoginTime strfmt.DateTime `json:"LastLoginTime,omitempty"`

		Permission *IamPermissionRef `json:"Permission,omitempty"`

		User *IamUserRef `json:"User,omitempty"`
	}

	dataAO1.AccountPermissions = m.AccountPermissions

	dataAO1.ClientIPAddress = m.ClientIPAddress

	dataAO1.Expiration = m.Expiration

	dataAO1.IdleTimeExpiration = m.IdleTimeExpiration

	dataAO1.LastLoginClient = m.LastLoginClient

	dataAO1.LastLoginTime = m.LastLoginTime

	dataAO1.Permission = m.Permission

	dataAO1.User = m.User

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam session
func (m *IamSession) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdleTimeExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLoginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamSession) validateAccountPermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountPermissions) { // not required
		return nil
	}

	for i := 0; i < len(m.AccountPermissions); i++ {
		if swag.IsZero(m.AccountPermissions[i]) { // not required
			continue
		}

		if m.AccountPermissions[i] != nil {
			if err := m.AccountPermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AccountPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IamSession) validateExpiration(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("Expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IamSession) validateIdleTimeExpiration(formats strfmt.Registry) error {

	if swag.IsZero(m.IdleTimeExpiration) { // not required
		return nil
	}

	if err := validate.FormatOf("IdleTimeExpiration", "body", "date-time", m.IdleTimeExpiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IamSession) validateLastLoginTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastLoginTime) { // not required
		return nil
	}

	if err := validate.FormatOf("LastLoginTime", "body", "date-time", m.LastLoginTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IamSession) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if m.Permission != nil {
		if err := m.Permission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Permission")
			}
			return err
		}
	}

	return nil
}

func (m *IamSession) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("User")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IamSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamSession) UnmarshalBinary(b []byte) error {
	var res IamSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
