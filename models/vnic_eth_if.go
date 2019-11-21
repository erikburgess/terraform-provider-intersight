// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicEthIf Virtual Ethernet Interface
//
// Virtual Ethernet Interface.
//
// swagger:model vnicEthIf
type VnicEthIf struct {
	MoBaseMo

	// Consistent Device Naming configuration for the virtual NIC.
	//
	Cdn *VnicCdn `json:"Cdn,omitempty"`

	// Relationship to the the Ethernet Adapter Policy.
	//
	EthAdapterPolicy *VnicEthAdapterPolicyRef `json:"EthAdapterPolicy,omitempty"`

	// Relationship to the Ethernet Network Policy.
	//
	EthNetworkPolicy *VnicEthNetworkPolicyRef `json:"EthNetworkPolicy,omitempty"`

	// Relationship to the Ethernet QoS Policy.
	//
	EthQosPolicy *VnicEthQosPolicyRef `json:"EthQosPolicy,omitempty"`

	// Relationship to the LAN Connectivity Policy.
	//
	LanConnectivityPolicy *VnicLanConnectivityPolicyRef `json:"LanConnectivityPolicy,omitempty"`

	// Name of the virtual ethernet interface.
	//
	Name string `json:"Name,omitempty"`

	// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
	//
	Order int64 `json:"Order,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *IamAccountRef `json:"Organization,omitempty"`

	// Placement Settings for the virtual interface.
	//
	Placement *VnicPlacementSettings `json:"Placement,omitempty"`

	// User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets.
	//
	UsnicSettings *VnicUsnicSettings `json:"UsnicSettings,omitempty"`

	// Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS.
	//
	VmqSettings *VnicVmqSettings `json:"VmqSettings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicEthIf) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Cdn *VnicCdn `json:"Cdn,omitempty"`

		EthAdapterPolicy *VnicEthAdapterPolicyRef `json:"EthAdapterPolicy,omitempty"`

		EthNetworkPolicy *VnicEthNetworkPolicyRef `json:"EthNetworkPolicy,omitempty"`

		EthQosPolicy *VnicEthQosPolicyRef `json:"EthQosPolicy,omitempty"`

		LanConnectivityPolicy *VnicLanConnectivityPolicyRef `json:"LanConnectivityPolicy,omitempty"`

		Name string `json:"Name,omitempty"`

		Order int64 `json:"Order,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Placement *VnicPlacementSettings `json:"Placement,omitempty"`

		UsnicSettings *VnicUsnicSettings `json:"UsnicSettings,omitempty"`

		VmqSettings *VnicVmqSettings `json:"VmqSettings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Cdn = dataAO1.Cdn

	m.EthAdapterPolicy = dataAO1.EthAdapterPolicy

	m.EthNetworkPolicy = dataAO1.EthNetworkPolicy

	m.EthQosPolicy = dataAO1.EthQosPolicy

	m.LanConnectivityPolicy = dataAO1.LanConnectivityPolicy

	m.Name = dataAO1.Name

	m.Order = dataAO1.Order

	m.Organization = dataAO1.Organization

	m.Placement = dataAO1.Placement

	m.UsnicSettings = dataAO1.UsnicSettings

	m.VmqSettings = dataAO1.VmqSettings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicEthIf) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Cdn *VnicCdn `json:"Cdn,omitempty"`

		EthAdapterPolicy *VnicEthAdapterPolicyRef `json:"EthAdapterPolicy,omitempty"`

		EthNetworkPolicy *VnicEthNetworkPolicyRef `json:"EthNetworkPolicy,omitempty"`

		EthQosPolicy *VnicEthQosPolicyRef `json:"EthQosPolicy,omitempty"`

		LanConnectivityPolicy *VnicLanConnectivityPolicyRef `json:"LanConnectivityPolicy,omitempty"`

		Name string `json:"Name,omitempty"`

		Order int64 `json:"Order,omitempty"`

		Organization *IamAccountRef `json:"Organization,omitempty"`

		Placement *VnicPlacementSettings `json:"Placement,omitempty"`

		UsnicSettings *VnicUsnicSettings `json:"UsnicSettings,omitempty"`

		VmqSettings *VnicVmqSettings `json:"VmqSettings,omitempty"`
	}

	dataAO1.Cdn = m.Cdn

	dataAO1.EthAdapterPolicy = m.EthAdapterPolicy

	dataAO1.EthNetworkPolicy = m.EthNetworkPolicy

	dataAO1.EthQosPolicy = m.EthQosPolicy

	dataAO1.LanConnectivityPolicy = m.LanConnectivityPolicy

	dataAO1.Name = m.Name

	dataAO1.Order = m.Order

	dataAO1.Organization = m.Organization

	dataAO1.Placement = m.Placement

	dataAO1.UsnicSettings = m.UsnicSettings

	dataAO1.VmqSettings = m.VmqSettings

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic eth if
func (m *VnicEthIf) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthAdapterPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthNetworkPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEthQosPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanConnectivityPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsnicSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmqSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VnicEthIf) validateCdn(formats strfmt.Registry) error {

	if swag.IsZero(m.Cdn) { // not required
		return nil
	}

	if m.Cdn != nil {
		if err := m.Cdn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cdn")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateEthAdapterPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.EthAdapterPolicy) { // not required
		return nil
	}

	if m.EthAdapterPolicy != nil {
		if err := m.EthAdapterPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EthAdapterPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateEthNetworkPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.EthNetworkPolicy) { // not required
		return nil
	}

	if m.EthNetworkPolicy != nil {
		if err := m.EthNetworkPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EthNetworkPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateEthQosPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.EthQosPolicy) { // not required
		return nil
	}

	if m.EthQosPolicy != nil {
		if err := m.EthQosPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EthQosPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateLanConnectivityPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.LanConnectivityPolicy) { // not required
		return nil
	}

	if m.LanConnectivityPolicy != nil {
		if err := m.LanConnectivityPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LanConnectivityPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateOrganization(formats strfmt.Registry) error {

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

func (m *VnicEthIf) validatePlacement(formats strfmt.Registry) error {

	if swag.IsZero(m.Placement) { // not required
		return nil
	}

	if m.Placement != nil {
		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Placement")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateUsnicSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.UsnicSettings) { // not required
		return nil
	}

	if m.UsnicSettings != nil {
		if err := m.UsnicSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsnicSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthIf) validateVmqSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.VmqSettings) { // not required
		return nil
	}

	if m.VmqSettings != nil {
		if err := m.VmqSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VmqSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicEthIf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicEthIf) UnmarshalBinary(b []byte) error {
	var res VnicEthIf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
