// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualizationDatacenter Virtualization:Datacenter
//
// Common attributes of any datacenter. Serves as a base class for all concrete hypervisor types. Datacenter (DC) is the container that brings together all other elements (Host, Datastore, Virtual Machine) A typical DC has datastores for storage for Virtual Machines, and a handful of hosts to run the Virtual Machines. In addition, there could be virtual switches and portgroups in those switches.
//
// swagger:model virtualizationDatacenter
type VirtualizationDatacenter struct {
	VirtualizationSourceDevice

	// Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference).
	// Read Only: true
	Identity string `json:"Identity,omitempty"`

	// User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter.
	Name string `json:"Name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VirtualizationDatacenter) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VirtualizationSourceDevice
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VirtualizationSourceDevice = aO0

	// AO1
	var dataAO1 struct {
		Identity string `json:"Identity,omitempty"`

		Name string `json:"Name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Identity = dataAO1.Identity

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VirtualizationDatacenter) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.VirtualizationSourceDevice)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Identity string `json:"Identity,omitempty"`

		Name string `json:"Name,omitempty"`
	}

	dataAO1.Identity = m.Identity

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this virtualization datacenter
func (m *VirtualizationDatacenter) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VirtualizationSourceDevice
	if err := m.VirtualizationSourceDevice.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationDatacenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationDatacenter) UnmarshalBinary(b []byte) error {
	var res VirtualizationDatacenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
