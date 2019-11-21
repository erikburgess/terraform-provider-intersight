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

// HyperflexNode Hyperflex:Node
// swagger:model hyperflexNode
type HyperflexNode struct {
	MoBaseMo

	// build number
	// Read Only: true
	BuildNumber string `json:"BuildNumber,omitempty"`

	// A collection of references to the [hyperflex.Cluster](mo://hyperflex.Cluster) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.Cluster](mo://hyperflex.Cluster) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	Cluster *HyperflexClusterRef `json:"Cluster,omitempty"`

	// A relationship to the ClusterMember that represents this node's registration to Intersight.
	//
	// Read Only: true
	ClusterMember *AssetClusterMemberRef `json:"ClusterMember,omitempty"`

	// display version
	// Read Only: true
	DisplayVersion string `json:"DisplayVersion,omitempty"`

	// host name
	// Read Only: true
	HostName string `json:"HostName,omitempty"`

	// hypervisor
	// Read Only: true
	Hypervisor string `json:"Hypervisor,omitempty"`

	// identity
	// Read Only: true
	Identity *HyperflexHxUuIDDt `json:"Identity,omitempty"`

	// Ip
	// Read Only: true
	IP *HyperflexHxNetworkAddressDt `json:"Ip,omitempty"`

	// lockdown
	// Read Only: true
	Lockdown *bool `json:"Lockdown,omitempty"`

	// model number
	// Read Only: true
	ModelNumber string `json:"ModelNumber,omitempty"`

	// A relationship to the UCS server associated with this node.
	//
	// Read Only: true
	PhysicalServer *ComputePhysicalRef `json:"PhysicalServer,omitempty"`

	// role
	// Read Only: true
	// Enum: [UNKNOWN STORAGE COMPUTE]
	Role string `json:"Role,omitempty"`

	// serial number
	// Read Only: true
	SerialNumber string `json:"SerialNumber,omitempty"`

	// status
	// Read Only: true
	// Enum: [UNKNOWN ONLINE OFFLINE INMAINTENANCE DEGRADED]
	Status string `json:"Status,omitempty"`

	// version
	// Read Only: true
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexNode) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		BuildNumber string `json:"BuildNumber,omitempty"`

		Cluster *HyperflexClusterRef `json:"Cluster,omitempty"`

		ClusterMember *AssetClusterMemberRef `json:"ClusterMember,omitempty"`

		DisplayVersion string `json:"DisplayVersion,omitempty"`

		HostName string `json:"HostName,omitempty"`

		Hypervisor string `json:"Hypervisor,omitempty"`

		Identity *HyperflexHxUuIDDt `json:"Identity,omitempty"`

		IP *HyperflexHxNetworkAddressDt `json:"Ip,omitempty"`

		Lockdown *bool `json:"Lockdown,omitempty"`

		ModelNumber string `json:"ModelNumber,omitempty"`

		PhysicalServer *ComputePhysicalRef `json:"PhysicalServer,omitempty"`

		Role string `json:"Role,omitempty"`

		SerialNumber string `json:"SerialNumber,omitempty"`

		Status string `json:"Status,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BuildNumber = dataAO1.BuildNumber

	m.Cluster = dataAO1.Cluster

	m.ClusterMember = dataAO1.ClusterMember

	m.DisplayVersion = dataAO1.DisplayVersion

	m.HostName = dataAO1.HostName

	m.Hypervisor = dataAO1.Hypervisor

	m.Identity = dataAO1.Identity

	m.IP = dataAO1.IP

	m.Lockdown = dataAO1.Lockdown

	m.ModelNumber = dataAO1.ModelNumber

	m.PhysicalServer = dataAO1.PhysicalServer

	m.Role = dataAO1.Role

	m.SerialNumber = dataAO1.SerialNumber

	m.Status = dataAO1.Status

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexNode) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		BuildNumber string `json:"BuildNumber,omitempty"`

		Cluster *HyperflexClusterRef `json:"Cluster,omitempty"`

		ClusterMember *AssetClusterMemberRef `json:"ClusterMember,omitempty"`

		DisplayVersion string `json:"DisplayVersion,omitempty"`

		HostName string `json:"HostName,omitempty"`

		Hypervisor string `json:"Hypervisor,omitempty"`

		Identity *HyperflexHxUuIDDt `json:"Identity,omitempty"`

		IP *HyperflexHxNetworkAddressDt `json:"Ip,omitempty"`

		Lockdown *bool `json:"Lockdown,omitempty"`

		ModelNumber string `json:"ModelNumber,omitempty"`

		PhysicalServer *ComputePhysicalRef `json:"PhysicalServer,omitempty"`

		Role string `json:"Role,omitempty"`

		SerialNumber string `json:"SerialNumber,omitempty"`

		Status string `json:"Status,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.BuildNumber = m.BuildNumber

	dataAO1.Cluster = m.Cluster

	dataAO1.ClusterMember = m.ClusterMember

	dataAO1.DisplayVersion = m.DisplayVersion

	dataAO1.HostName = m.HostName

	dataAO1.Hypervisor = m.Hypervisor

	dataAO1.Identity = m.Identity

	dataAO1.IP = m.IP

	dataAO1.Lockdown = m.Lockdown

	dataAO1.ModelNumber = m.ModelNumber

	dataAO1.PhysicalServer = m.PhysicalServer

	dataAO1.Role = m.Role

	dataAO1.SerialNumber = m.SerialNumber

	dataAO1.Status = m.Status

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex node
func (m *HyperflexNode) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexNode) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cluster")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexNode) validateClusterMember(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterMember) { // not required
		return nil
	}

	if m.ClusterMember != nil {
		if err := m.ClusterMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterMember")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexNode) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identity")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexNode) validateIP(formats strfmt.Registry) error {

	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Ip")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexNode) validatePhysicalServer(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalServer) { // not required
		return nil
	}

	if m.PhysicalServer != nil {
		if err := m.PhysicalServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PhysicalServer")
			}
			return err
		}
	}

	return nil
}

var hyperflexNodeTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","STORAGE","COMPUTE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexNodeTypeRolePropEnum = append(hyperflexNodeTypeRolePropEnum, v)
	}
}

// property enum
func (m *HyperflexNode) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexNodeTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexNode) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("Role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

var hyperflexNodeTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","ONLINE","OFFLINE","INMAINTENANCE","DEGRADED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexNodeTypeStatusPropEnum = append(hyperflexNodeTypeStatusPropEnum, v)
	}
}

// property enum
func (m *HyperflexNode) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexNodeTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexNode) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexNode) UnmarshalBinary(b []byte) error {
	var res HyperflexNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
