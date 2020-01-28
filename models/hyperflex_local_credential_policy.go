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

// HyperflexLocalCredentialPolicy Security
//
// A policy specifying credentials for HyperFlex cluster such as controller VM password, hypervisor username, and password.
//
// swagger:model hyperflexLocalCredentialPolicy
type HyperflexLocalCredentialPolicy struct {
	PolicyAbstractPolicy

	// List of cluster profiles using this policy.
	//
	ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

	// Indicates if Hypervisor password is the factory set default password. For HyperFlex Data Platform versions 3.0 or higher, enable this if the default password was not changed on the Hypervisor. It is required to supply a new custom Hypervisor password that will be applied to the Hypervisor during deployment. For HyperFlex Data Platform versions prior to 3.0 release, this setting has no effect and the default password will be used for initial install. The Hypervisor password should be changed after deployment.
	//
	FactoryHypervisorPassword *bool `json:"FactoryHypervisorPassword,omitempty"`

	// HyperFlex storage controller VM password must contain a minimum of 10 characters, with at least 1 lowercase, 1 uppercase, 1 numeric, and 1 of these -_@#$%^&*! special characters.
	//
	HxdpRootPwd string `json:"HxdpRootPwd,omitempty"`

	// Hypervisor administrator username must contain only alphanumeric characters. Use the root account for ESXi deployments.
	//
	HypervisorAdmin string `json:"HypervisorAdmin,omitempty"`

	// The ESXi root password. For HyperFlex Data Platform 3.0 or later, if the factory default password was not manually changed, you must set a new custom password. If the password was manually changed, you must not enable the factory default password property and provide the current hypervisor password. Note - All HyperFlex nodes require the same hypervisor password for installation. For HyperFlex Data Platform prior to 3.0, use the default password "Cisco123" for newly manufactured HyperFlex servers. A custom password should only be entered if hypervisor credentials were manually changed prior to deployment.
	//
	HypervisorAdminPwd string `json:"HypervisorAdminPwd,omitempty"`

	// Indicates whether the value of the 'hxdpRootPwd' property has been set.
	//
	// Read Only: true
	IsHxdpRootPwdSet *bool `json:"IsHxdpRootPwdSet,omitempty"`

	// Indicates whether the value of the 'hypervisorAdminPwd' property has been set.
	//
	// Read Only: true
	IsHypervisorAdminPwdSet *bool `json:"IsHypervisorAdminPwdSet,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexLocalCredentialPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		FactoryHypervisorPassword *bool `json:"FactoryHypervisorPassword,omitempty"`

		HxdpRootPwd string `json:"HxdpRootPwd,omitempty"`

		HypervisorAdmin string `json:"HypervisorAdmin,omitempty"`

		HypervisorAdminPwd string `json:"HypervisorAdminPwd,omitempty"`

		IsHxdpRootPwdSet *bool `json:"IsHxdpRootPwdSet,omitempty"`

		IsHypervisorAdminPwdSet *bool `json:"IsHypervisorAdminPwdSet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ClusterProfiles = dataAO1.ClusterProfiles

	m.FactoryHypervisorPassword = dataAO1.FactoryHypervisorPassword

	m.HxdpRootPwd = dataAO1.HxdpRootPwd

	m.HypervisorAdmin = dataAO1.HypervisorAdmin

	m.HypervisorAdminPwd = dataAO1.HypervisorAdminPwd

	m.IsHxdpRootPwdSet = dataAO1.IsHxdpRootPwdSet

	m.IsHypervisorAdminPwdSet = dataAO1.IsHypervisorAdminPwdSet

	m.Organization = dataAO1.Organization

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexLocalCredentialPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ClusterProfiles []*HyperflexClusterProfileRef `json:"ClusterProfiles"`

		FactoryHypervisorPassword *bool `json:"FactoryHypervisorPassword,omitempty"`

		HxdpRootPwd string `json:"HxdpRootPwd,omitempty"`

		HypervisorAdmin string `json:"HypervisorAdmin,omitempty"`

		HypervisorAdminPwd string `json:"HypervisorAdminPwd,omitempty"`

		IsHxdpRootPwdSet *bool `json:"IsHxdpRootPwdSet,omitempty"`

		IsHypervisorAdminPwdSet *bool `json:"IsHypervisorAdminPwdSet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}

	dataAO1.ClusterProfiles = m.ClusterProfiles

	dataAO1.FactoryHypervisorPassword = m.FactoryHypervisorPassword

	dataAO1.HxdpRootPwd = m.HxdpRootPwd

	dataAO1.HypervisorAdmin = m.HypervisorAdmin

	dataAO1.HypervisorAdminPwd = m.HypervisorAdminPwd

	dataAO1.IsHxdpRootPwdSet = m.IsHxdpRootPwdSet

	dataAO1.IsHypervisorAdminPwdSet = m.IsHypervisorAdminPwdSet

	dataAO1.Organization = m.Organization

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex local credential policy
func (m *HyperflexLocalCredentialPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexLocalCredentialPolicy) validateClusterProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfiles); i++ {
		if swag.IsZero(m.ClusterProfiles[i]) { // not required
			continue
		}

		if m.ClusterProfiles[i] != nil {
			if err := m.ClusterProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ClusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexLocalCredentialPolicy) validateOrganization(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *HyperflexLocalCredentialPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexLocalCredentialPolicy) UnmarshalBinary(b []byte) error {
	var res HyperflexLocalCredentialPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
