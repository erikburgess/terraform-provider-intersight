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

// KvmPolicy Virtual KVM
//
// Policy to configure KVM Launch settings.
//
// swagger:model kvmPolicy
type KvmPolicy struct {
	PolicyAbstractPolicy

	// If enabled, displays KVM session on any monitor attached to the server.
	//
	EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`

	// If enabled, encrypts all video information sent through KVM.
	//
	EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`

	// State of the vKVM service on the endpoint.
	//
	Enabled *bool `json:"Enabled,omitempty"`

	// The maximum number of concurrent KVM sessions allowed.
	//
	MaximumSessions *int64 `json:"MaximumSessions,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Relationship to the profile object.
	//
	Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

	// The port used for KVM communication.
	//
	RemotePort int64 `json:"RemotePort,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *KvmPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`

		EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`

		Enabled *bool `json:"Enabled,omitempty"`

		MaximumSessions *int64 `json:"MaximumSessions,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		RemotePort int64 `json:"RemotePort,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EnableLocalServerVideo = dataAO1.EnableLocalServerVideo

	m.EnableVideoEncryption = dataAO1.EnableVideoEncryption

	m.Enabled = dataAO1.Enabled

	m.MaximumSessions = dataAO1.MaximumSessions

	m.Organization = dataAO1.Organization

	m.Profiles = dataAO1.Profiles

	m.RemotePort = dataAO1.RemotePort

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m KvmPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`

		EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`

		Enabled *bool `json:"Enabled,omitempty"`

		MaximumSessions *int64 `json:"MaximumSessions,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Profiles []*PolicyAbstractConfigProfileRef `json:"Profiles"`

		RemotePort int64 `json:"RemotePort,omitempty"`
	}

	dataAO1.EnableLocalServerVideo = m.EnableLocalServerVideo

	dataAO1.EnableVideoEncryption = m.EnableVideoEncryption

	dataAO1.Enabled = m.Enabled

	dataAO1.MaximumSessions = m.MaximumSessions

	dataAO1.Organization = m.Organization

	dataAO1.Profiles = m.Profiles

	dataAO1.RemotePort = m.RemotePort

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this kvm policy
func (m *KvmPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KvmPolicy) validateOrganization(formats strfmt.Registry) error {

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

func (m *KvmPolicy) validateProfiles(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *KvmPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KvmPolicy) UnmarshalBinary(b []byte) error {
	var res KvmPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}