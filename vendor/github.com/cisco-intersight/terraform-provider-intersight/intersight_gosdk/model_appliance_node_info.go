/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ApplianceNodeInfo NodeInfo managed object stores the Intersight Appliance's cluster node information. NodeInfo managed objects are created during the Intersight Appliance setup. The Intersight Appliance updates the NodeInfo managed objects with status information periodically.
type ApplianceNodeInfo struct {
	MoBaseMo
	// Cluster node's FQDN or IP address.
	Hostname *string `json:"Hostname,omitempty"`
	// System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
	NodeId         *int64             `json:"NodeId,omitempty"`
	NodeIpV4Config *CommIpV4Interface `json:"NodeIpV4Config,omitempty"`
	// Operational status of the Intersight Appliance node.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
}

// NewApplianceNodeInfo instantiates a new ApplianceNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNodeInfo() *ApplianceNodeInfo {
	this := ApplianceNodeInfo{}
	var operationalStatus string = "Unknown"
	this.OperationalStatus = &operationalStatus
	return &this
}

// NewApplianceNodeInfoWithDefaults instantiates a new ApplianceNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNodeInfoWithDefaults() *ApplianceNodeInfo {
	this := ApplianceNodeInfo{}
	var operationalStatus string = "Unknown"
	this.OperationalStatus = &operationalStatus
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceNodeInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceNodeInfo) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceNodeInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceNodeInfo) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfo) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceNodeInfo) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceNodeInfo) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeIpV4Config returns the NodeIpV4Config field value if set, zero value otherwise.
func (o *ApplianceNodeInfo) GetNodeIpV4Config() CommIpV4Interface {
	if o == nil || o.NodeIpV4Config == nil {
		var ret CommIpV4Interface
		return ret
	}
	return *o.NodeIpV4Config
}

// GetNodeIpV4ConfigOk returns a tuple with the NodeIpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfo) GetNodeIpV4ConfigOk() (*CommIpV4Interface, bool) {
	if o == nil || o.NodeIpV4Config == nil {
		return nil, false
	}
	return o.NodeIpV4Config, true
}

// HasNodeIpV4Config returns a boolean if a field has been set.
func (o *ApplianceNodeInfo) HasNodeIpV4Config() bool {
	if o != nil && o.NodeIpV4Config != nil {
		return true
	}

	return false
}

// SetNodeIpV4Config gets a reference to the given CommIpV4Interface and assigns it to the NodeIpV4Config field.
func (o *ApplianceNodeInfo) SetNodeIpV4Config(v CommIpV4Interface) {
	o.NodeIpV4Config = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceNodeInfo) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfo) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceNodeInfo) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceNodeInfo) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

func (o ApplianceNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.NodeIpV4Config != nil {
		toSerialize["NodeIpV4Config"] = o.NodeIpV4Config
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceNodeInfo struct {
	value *ApplianceNodeInfo
	isSet bool
}

func (v NullableApplianceNodeInfo) Get() *ApplianceNodeInfo {
	return v.value
}

func (v *NullableApplianceNodeInfo) Set(val *ApplianceNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNodeInfo(val *ApplianceNodeInfo) *NullableApplianceNodeInfo {
	return &NullableApplianceNodeInfo{value: val, isSet: true}
}

func (v NullableApplianceNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
