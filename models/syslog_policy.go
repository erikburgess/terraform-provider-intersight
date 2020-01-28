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

// SyslogPolicy Syslog
//
// The syslog policy configure the syslog server to receive CIMC log entries.
//
// swagger:model syslogPolicy
type SyslogPolicy struct {
	PolicyAbstractPolicy

	// Set of local logging clients on the endpoint.
	//
	LocalClients []*SyslogLocalClientBase `json:"LocalClients"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Relationship to the profile object.
	//
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

	// Set of remote logging clients on the endpoint.
	//
	RemoteClients []*SyslogRemoteClientBase `json:"RemoteClients"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SyslogPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		LocalClients []*SyslogLocalClientBase `json:"LocalClients"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		RemoteClients []*SyslogRemoteClientBase `json:"RemoteClients"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LocalClients = dataAO1.LocalClients

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	m.RemoteClients = dataAO1.RemoteClients

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SyslogPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		LocalClients []*SyslogLocalClientBase `json:"LocalClients"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		RemoteClients []*SyslogRemoteClientBase `json:"RemoteClients"`
	}

	dataAO1.LocalClients = m.LocalClients

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	dataAO1.RemoteClients = m.RemoteClients

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this syslog policy
func (m *SyslogPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteClients(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyslogPolicy) validateLocalClients(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalClients) { // not required
		return nil
	}

	for i := 0; i < len(m.LocalClients); i++ {
		if swag.IsZero(m.LocalClients[i]) { // not required
			continue
		}

		if m.LocalClients[i] != nil {
			if err := m.LocalClients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LocalClients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SyslogPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *SyslogPolicy) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SyslogPolicy) validateRemoteClients(formats strfmt.Registry) error {

	if swag.IsZero(m.RemoteClients) { // not required
		return nil
	}

	for i := 0; i < len(m.RemoteClients); i++ {
		if swag.IsZero(m.RemoteClients[i]) { // not required
			continue
		}

		if m.RemoteClients[i] != nil {
			if err := m.RemoteClients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoteClients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyslogPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyslogPolicy) UnmarshalBinary(b []byte) error {
	var res SyslogPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
