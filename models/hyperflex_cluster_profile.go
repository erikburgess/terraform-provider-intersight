// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HyperflexClusterProfile HyperFlex Cluster Profile
//
// A profile specifying configuration settings for a HyperFlex cluster.
//
// swagger:model hyperflexClusterProfile
type HyperflexClusterProfile struct {
	PolicyAbstractConfigProfile

	// HyperFlex cluster associated with this cluster profile.
	//
	// Read Only: true
	AssociatedCluster *HyperflexClusterRef `json:"AssociatedCluster,omitempty"`

	// A collection of references to the [hyperflex.AutoSupportPolicy](mo://hyperflex.AutoSupportPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.AutoSupportPolicy](mo://hyperflex.AutoSupportPolicy) MO unsets its reference to this deleted MO.
	//
	AutoSupport *HyperflexAutoSupportPolicyRef `json:"AutoSupport,omitempty"`

	// A collection of references to the [hyperflex.ClusterNetworkPolicy](mo://hyperflex.ClusterNetworkPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.ClusterNetworkPolicy](mo://hyperflex.ClusterNetworkPolicy) MO unsets its reference to this deleted MO.
	//
	ClusterNetwork *HyperflexClusterNetworkPolicyRef `json:"ClusterNetwork,omitempty"`

	// A collection of references to the [hyperflex.ClusterStoragePolicy](mo://hyperflex.ClusterStoragePolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.ClusterStoragePolicy](mo://hyperflex.ClusterStoragePolicy) MO unsets its reference to this deleted MO.
	//
	ClusterStorage *HyperflexClusterStoragePolicyRef `json:"ClusterStorage,omitempty"`

	// The profile configuration (deploy, validation) results with the overall state and detailed result messages.
	//
	// Read Only: true
	ConfigResult *HyperflexConfigResultRef `json:"ConfigResult,omitempty"`

	// The storage data IP address for the HyperFlex cluster.
	//
	DataIPAddress string `json:"DataIpAddress,omitempty"`

	// A collection of references to the [hyperflex.ExtFcStoragePolicy](mo://hyperflex.ExtFcStoragePolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.ExtFcStoragePolicy](mo://hyperflex.ExtFcStoragePolicy) MO unsets its reference to this deleted MO.
	//
	ExtFcStorage *HyperflexExtFcStoragePolicyRef `json:"ExtFcStorage,omitempty"`

	// A collection of references to the [hyperflex.ExtIscsiStoragePolicy](mo://hyperflex.ExtIscsiStoragePolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.ExtIscsiStoragePolicy](mo://hyperflex.ExtIscsiStoragePolicy) MO unsets its reference to this deleted MO.
	//
	ExtIscsiStorage *HyperflexExtIscsiStoragePolicyRef `json:"ExtIscsiStorage,omitempty"`

	// The hypervisor type for the HyperFlex cluster.
	//
	// Enum: [Unknown Hyper-V ESXi]
	HypervisorType *string `json:"HypervisorType,omitempty"`

	// A collection of references to the [hyperflex.LocalCredentialPolicy](mo://hyperflex.LocalCredentialPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.LocalCredentialPolicy](mo://hyperflex.LocalCredentialPolicy) MO unsets its reference to this deleted MO.
	//
	LocalCredential *HyperflexLocalCredentialPolicyRef `json:"LocalCredential,omitempty"`

	// The MAC address prefix in the form of 00:25:B5:XX.
	//
	MacAddressPrefix string `json:"MacAddressPrefix,omitempty"`

	// The management IP address for the HyperFlex cluster.
	//
	MgmtIPAddress string `json:"MgmtIpAddress,omitempty"`

	// The management platform for the HyperFlex cluster.
	//
	// Enum: [FI EDGE]
	MgmtPlatform *string `json:"MgmtPlatform,omitempty"`

	// A collection of references to the [hyperflex.NodeConfigPolicy](mo://hyperflex.NodeConfigPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.NodeConfigPolicy](mo://hyperflex.NodeConfigPolicy) MO unsets its reference to this deleted MO.
	//
	NodeConfig *HyperflexNodeConfigPolicyRef `json:"NodeConfig,omitempty"`

	// List of node profiles representing the configuraion of the corresponding HX cluster nodes.
	//
	NodeProfileConfig []*HyperflexNodeProfileRef `json:"NodeProfileConfig"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// A collection of references to the [hyperflex.ProxySettingPolicy](mo://hyperflex.ProxySettingPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.ProxySettingPolicy](mo://hyperflex.ProxySettingPolicy) MO unsets its reference to this deleted MO.
	//
	ProxySetting *HyperflexProxySettingPolicyRef `json:"ProxySetting,omitempty"`

	// The number of copies of each data block written.
	//
	Replication int64 `json:"Replication,omitempty"`

	// List of workflows that are currently running for this cluster profile.
	//
	// Read Only: true
	RunningWorkflows []*WorkflowWorkflowInfoRef `json:"RunningWorkflows"`

	// A collection of references to the [hyperflex.SoftwareVersionPolicy](mo://hyperflex.SoftwareVersionPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.SoftwareVersionPolicy](mo://hyperflex.SoftwareVersionPolicy) MO unsets its reference to this deleted MO.
	//
	SoftwareVersion *HyperflexSoftwareVersionPolicyRef `json:"SoftwareVersion,omitempty"`

	// The VLAN for the HyperFlex storage data traffic.
	//
	StorageDataVlan *HyperflexNamedVlan `json:"StorageDataVlan,omitempty"`

	// A collection of references to the [hyperflex.SysConfigPolicy](mo://hyperflex.SysConfigPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.SysConfigPolicy](mo://hyperflex.SysConfigPolicy) MO unsets its reference to this deleted MO.
	//
	SysConfig *HyperflexSysConfigPolicyRef `json:"SysConfig,omitempty"`

	// A collection of references to the [hyperflex.UcsmConfigPolicy](mo://hyperflex.UcsmConfigPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.UcsmConfigPolicy](mo://hyperflex.UcsmConfigPolicy) MO unsets its reference to this deleted MO.
	//
	UcsmConfig *HyperflexUcsmConfigPolicyRef `json:"UcsmConfig,omitempty"`

	// A collection of references to the [hyperflex.VcenterConfigPolicy](mo://hyperflex.VcenterConfigPolicy) Managed Object.
	//
	// When this managed object is deleted, the referenced [hyperflex.VcenterConfigPolicy](mo://hyperflex.VcenterConfigPolicy) MO unsets its reference to this deleted MO.
	//
	VcenterConfig *HyperflexVcenterConfigPolicyRef `json:"VcenterConfig,omitempty"`

	// The WWxN prefix in the form of 20:00:00:25:B5:XX.
	//
	WwxnPrefix string `json:"WwxnPrefix,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexClusterProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractConfigProfile
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractConfigProfile = aO0

	// AO1
	var dataAO1 struct {
		AssociatedCluster *HyperflexClusterRef `json:"AssociatedCluster,omitempty"`

		AutoSupport *HyperflexAutoSupportPolicyRef `json:"AutoSupport,omitempty"`

		ClusterNetwork *HyperflexClusterNetworkPolicyRef `json:"ClusterNetwork,omitempty"`

		ClusterStorage *HyperflexClusterStoragePolicyRef `json:"ClusterStorage,omitempty"`

		ConfigResult *HyperflexConfigResultRef `json:"ConfigResult,omitempty"`

		DataIPAddress string `json:"DataIpAddress,omitempty"`

		ExtFcStorage *HyperflexExtFcStoragePolicyRef `json:"ExtFcStorage,omitempty"`

		ExtIscsiStorage *HyperflexExtIscsiStoragePolicyRef `json:"ExtIscsiStorage,omitempty"`

		HypervisorType *string `json:"HypervisorType,omitempty"`

		LocalCredential *HyperflexLocalCredentialPolicyRef `json:"LocalCredential,omitempty"`

		MacAddressPrefix string `json:"MacAddressPrefix,omitempty"`

		MgmtIPAddress string `json:"MgmtIpAddress,omitempty"`

		MgmtPlatform *string `json:"MgmtPlatform,omitempty"`

		NodeConfig *HyperflexNodeConfigPolicyRef `json:"NodeConfig,omitempty"`

		NodeProfileConfig []*HyperflexNodeProfileRef `json:"NodeProfileConfig"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		ProxySetting *HyperflexProxySettingPolicyRef `json:"ProxySetting,omitempty"`

		Replication int64 `json:"Replication,omitempty"`

		RunningWorkflows []*WorkflowWorkflowInfoRef `json:"RunningWorkflows"`

		SoftwareVersion *HyperflexSoftwareVersionPolicyRef `json:"SoftwareVersion,omitempty"`

		StorageDataVlan *HyperflexNamedVlan `json:"StorageDataVlan,omitempty"`

		SysConfig *HyperflexSysConfigPolicyRef `json:"SysConfig,omitempty"`

		UcsmConfig *HyperflexUcsmConfigPolicyRef `json:"UcsmConfig,omitempty"`

		VcenterConfig *HyperflexVcenterConfigPolicyRef `json:"VcenterConfig,omitempty"`

		WwxnPrefix string `json:"WwxnPrefix,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AssociatedCluster = dataAO1.AssociatedCluster

	m.AutoSupport = dataAO1.AutoSupport

	m.ClusterNetwork = dataAO1.ClusterNetwork

	m.ClusterStorage = dataAO1.ClusterStorage

	m.ConfigResult = dataAO1.ConfigResult

	m.DataIPAddress = dataAO1.DataIPAddress

	m.ExtFcStorage = dataAO1.ExtFcStorage

	m.ExtIscsiStorage = dataAO1.ExtIscsiStorage

	m.HypervisorType = dataAO1.HypervisorType

	m.LocalCredential = dataAO1.LocalCredential

	m.MacAddressPrefix = dataAO1.MacAddressPrefix

	m.MgmtIPAddress = dataAO1.MgmtIPAddress

	m.MgmtPlatform = dataAO1.MgmtPlatform

	m.NodeConfig = dataAO1.NodeConfig

	m.NodeProfileConfig = dataAO1.NodeProfileConfig

	m.Organization = dataAO1.Organization

	m.ProxySetting = dataAO1.ProxySetting

	m.Replication = dataAO1.Replication

	m.RunningWorkflows = dataAO1.RunningWorkflows

	m.SoftwareVersion = dataAO1.SoftwareVersion

	m.StorageDataVlan = dataAO1.StorageDataVlan

	m.SysConfig = dataAO1.SysConfig

	m.UcsmConfig = dataAO1.UcsmConfig

	m.VcenterConfig = dataAO1.VcenterConfig

	m.WwxnPrefix = dataAO1.WwxnPrefix

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexClusterProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractConfigProfile)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AssociatedCluster *HyperflexClusterRef `json:"AssociatedCluster,omitempty"`

		AutoSupport *HyperflexAutoSupportPolicyRef `json:"AutoSupport,omitempty"`

		ClusterNetwork *HyperflexClusterNetworkPolicyRef `json:"ClusterNetwork,omitempty"`

		ClusterStorage *HyperflexClusterStoragePolicyRef `json:"ClusterStorage,omitempty"`

		ConfigResult *HyperflexConfigResultRef `json:"ConfigResult,omitempty"`

		DataIPAddress string `json:"DataIpAddress,omitempty"`

		ExtFcStorage *HyperflexExtFcStoragePolicyRef `json:"ExtFcStorage,omitempty"`

		ExtIscsiStorage *HyperflexExtIscsiStoragePolicyRef `json:"ExtIscsiStorage,omitempty"`

		HypervisorType *string `json:"HypervisorType,omitempty"`

		LocalCredential *HyperflexLocalCredentialPolicyRef `json:"LocalCredential,omitempty"`

		MacAddressPrefix string `json:"MacAddressPrefix,omitempty"`

		MgmtIPAddress string `json:"MgmtIpAddress,omitempty"`

		MgmtPlatform *string `json:"MgmtPlatform,omitempty"`

		NodeConfig *HyperflexNodeConfigPolicyRef `json:"NodeConfig,omitempty"`

		NodeProfileConfig []*HyperflexNodeProfileRef `json:"NodeProfileConfig"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		ProxySetting *HyperflexProxySettingPolicyRef `json:"ProxySetting,omitempty"`

		Replication int64 `json:"Replication,omitempty"`

		RunningWorkflows []*WorkflowWorkflowInfoRef `json:"RunningWorkflows"`

		SoftwareVersion *HyperflexSoftwareVersionPolicyRef `json:"SoftwareVersion,omitempty"`

		StorageDataVlan *HyperflexNamedVlan `json:"StorageDataVlan,omitempty"`

		SysConfig *HyperflexSysConfigPolicyRef `json:"SysConfig,omitempty"`

		UcsmConfig *HyperflexUcsmConfigPolicyRef `json:"UcsmConfig,omitempty"`

		VcenterConfig *HyperflexVcenterConfigPolicyRef `json:"VcenterConfig,omitempty"`

		WwxnPrefix string `json:"WwxnPrefix,omitempty"`
	}

	dataAO1.AssociatedCluster = m.AssociatedCluster

	dataAO1.AutoSupport = m.AutoSupport

	dataAO1.ClusterNetwork = m.ClusterNetwork

	dataAO1.ClusterStorage = m.ClusterStorage

	dataAO1.ConfigResult = m.ConfigResult

	dataAO1.DataIPAddress = m.DataIPAddress

	dataAO1.ExtFcStorage = m.ExtFcStorage

	dataAO1.ExtIscsiStorage = m.ExtIscsiStorage

	dataAO1.HypervisorType = m.HypervisorType

	dataAO1.LocalCredential = m.LocalCredential

	dataAO1.MacAddressPrefix = m.MacAddressPrefix

	dataAO1.MgmtIPAddress = m.MgmtIPAddress

	dataAO1.MgmtPlatform = m.MgmtPlatform

	dataAO1.NodeConfig = m.NodeConfig

	dataAO1.NodeProfileConfig = m.NodeProfileConfig

	dataAO1.Organization = m.Organization

	dataAO1.ProxySetting = m.ProxySetting

	dataAO1.Replication = m.Replication

	dataAO1.RunningWorkflows = m.RunningWorkflows

	dataAO1.SoftwareVersion = m.SoftwareVersion

	dataAO1.StorageDataVlan = m.StorageDataVlan

	dataAO1.SysConfig = m.SysConfig

	dataAO1.UcsmConfig = m.UcsmConfig

	dataAO1.VcenterConfig = m.VcenterConfig

	dataAO1.WwxnPrefix = m.WwxnPrefix

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex cluster profile
func (m *HyperflexClusterProfile) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractConfigProfile
	if err := m.PolicyAbstractConfigProfile.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoSupport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtFcStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtIscsiStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervisorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMgmtPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeProfileConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxySetting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunningWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageDataVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSysConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUcsmConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperflexClusterProfile) validateAssociatedCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.AssociatedCluster) { // not required
		return nil
	}

	if m.AssociatedCluster != nil {
		if err := m.AssociatedCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AssociatedCluster")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateAutoSupport(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoSupport) { // not required
		return nil
	}

	if m.AutoSupport != nil {
		if err := m.AutoSupport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AutoSupport")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateClusterNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetwork) { // not required
		return nil
	}

	if m.ClusterNetwork != nil {
		if err := m.ClusterNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateClusterStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterStorage) { // not required
		return nil
	}

	if m.ClusterStorage != nil {
		if err := m.ClusterStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ClusterStorage")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateConfigResult(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigResult) { // not required
		return nil
	}

	if m.ConfigResult != nil {
		if err := m.ConfigResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigResult")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateExtFcStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtFcStorage) { // not required
		return nil
	}

	if m.ExtFcStorage != nil {
		if err := m.ExtFcStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExtFcStorage")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateExtIscsiStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtIscsiStorage) { // not required
		return nil
	}

	if m.ExtIscsiStorage != nil {
		if err := m.ExtIscsiStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExtIscsiStorage")
			}
			return err
		}
	}

	return nil
}

var hyperflexClusterProfileTypeHypervisorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Hyper-V","ESXi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexClusterProfileTypeHypervisorTypePropEnum = append(hyperflexClusterProfileTypeHypervisorTypePropEnum, v)
	}
}

// property enum
func (m *HyperflexClusterProfile) validateHypervisorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexClusterProfileTypeHypervisorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexClusterProfile) validateHypervisorType(formats strfmt.Registry) error {

	if swag.IsZero(m.HypervisorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateHypervisorTypeEnum("HypervisorType", "body", *m.HypervisorType); err != nil {
		return err
	}

	return nil
}

func (m *HyperflexClusterProfile) validateLocalCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalCredential) { // not required
		return nil
	}

	if m.LocalCredential != nil {
		if err := m.LocalCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LocalCredential")
			}
			return err
		}
	}

	return nil
}

var hyperflexClusterProfileTypeMgmtPlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FI","EDGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexClusterProfileTypeMgmtPlatformPropEnum = append(hyperflexClusterProfileTypeMgmtPlatformPropEnum, v)
	}
}

// property enum
func (m *HyperflexClusterProfile) validateMgmtPlatformEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexClusterProfileTypeMgmtPlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexClusterProfile) validateMgmtPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.MgmtPlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateMgmtPlatformEnum("MgmtPlatform", "body", *m.MgmtPlatform); err != nil {
		return err
	}

	return nil
}

func (m *HyperflexClusterProfile) validateNodeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeConfig) { // not required
		return nil
	}

	if m.NodeConfig != nil {
		if err := m.NodeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodeConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateNodeProfileConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeProfileConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeProfileConfig); i++ {
		if swag.IsZero(m.NodeProfileConfig[i]) { // not required
			continue
		}

		if m.NodeProfileConfig[i] != nil {
			if err := m.NodeProfileConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NodeProfileConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexClusterProfile) validateOrganization(formats strfmt.Registry) error {

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

func (m *HyperflexClusterProfile) validateProxySetting(formats strfmt.Registry) error {

	if swag.IsZero(m.ProxySetting) { // not required
		return nil
	}

	if m.ProxySetting != nil {
		if err := m.ProxySetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProxySetting")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateRunningWorkflows(formats strfmt.Registry) error {

	if swag.IsZero(m.RunningWorkflows) { // not required
		return nil
	}

	for i := 0; i < len(m.RunningWorkflows); i++ {
		if swag.IsZero(m.RunningWorkflows[i]) { // not required
			continue
		}

		if m.RunningWorkflows[i] != nil {
			if err := m.RunningWorkflows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RunningWorkflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperflexClusterProfile) validateSoftwareVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.SoftwareVersion) { // not required
		return nil
	}

	if m.SoftwareVersion != nil {
		if err := m.SoftwareVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SoftwareVersion")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateStorageDataVlan(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageDataVlan) { // not required
		return nil
	}

	if m.StorageDataVlan != nil {
		if err := m.StorageDataVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageDataVlan")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateSysConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SysConfig) { // not required
		return nil
	}

	if m.SysConfig != nil {
		if err := m.SysConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SysConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateUcsmConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.UcsmConfig) { // not required
		return nil
	}

	if m.UcsmConfig != nil {
		if err := m.UcsmConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UcsmConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HyperflexClusterProfile) validateVcenterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.VcenterConfig) { // not required
		return nil
	}

	if m.VcenterConfig != nil {
		if err := m.VcenterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VcenterConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexClusterProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexClusterProfile) UnmarshalBinary(b []byte) error {
	var res HyperflexClusterProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
