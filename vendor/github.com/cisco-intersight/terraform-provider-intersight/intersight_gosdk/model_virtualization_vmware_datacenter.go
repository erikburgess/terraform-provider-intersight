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

// VirtualizationVmwareDatacenter Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
type VirtualizationVmwareDatacenter struct {
	VirtualizationBaseDatacenter
	// Count of all clusters associated with this DC.
	ClusterCount *int64 `json:"ClusterCount,omitempty"`
	// Count of all datastores associated with this DC.
	DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
	// Count of all hosts associated with this DC.
	HostCount *int64 `json:"HostCount,omitempty"`
	// Count of all networks associated with this datacenter (DC).
	NetworkCount *int64 `json:"NetworkCount,omitempty"`
	// Count of all virtual machines (VMs) associated with this DC.
	VmCount           *int64                                   `json:"VmCount,omitempty"`
	HypervisorManager *VirtualizationVmwareVcenterRelationship `json:"HypervisorManager,omitempty"`
}

// NewVirtualizationVmwareDatacenter instantiates a new VirtualizationVmwareDatacenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDatacenter() *VirtualizationVmwareDatacenter {
	this := VirtualizationVmwareDatacenter{}
	return &this
}

// NewVirtualizationVmwareDatacenterWithDefaults instantiates a new VirtualizationVmwareDatacenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDatacenterWithDefaults() *VirtualizationVmwareDatacenter {
	this := VirtualizationVmwareDatacenter{}
	return &this
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetClusterCount() int64 {
	if o == nil || o.ClusterCount == nil {
		var ret int64
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetClusterCountOk() (*int64, bool) {
	if o == nil || o.ClusterCount == nil {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasClusterCount() bool {
	if o != nil && o.ClusterCount != nil {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int64 and assigns it to the ClusterCount field.
func (o *VirtualizationVmwareDatacenter) SetClusterCount(v int64) {
	o.ClusterCount = &v
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetDatastoreCount() int64 {
	if o == nil || o.DatastoreCount == nil {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || o.DatastoreCount == nil {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasDatastoreCount() bool {
	if o != nil && o.DatastoreCount != nil {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareDatacenter) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationVmwareDatacenter) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetNetworkCount returns the NetworkCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetNetworkCount() int64 {
	if o == nil || o.NetworkCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkCount
}

// GetNetworkCountOk returns a tuple with the NetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetNetworkCountOk() (*int64, bool) {
	if o == nil || o.NetworkCount == nil {
		return nil, false
	}
	return o.NetworkCount, true
}

// HasNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasNetworkCount() bool {
	if o != nil && o.NetworkCount != nil {
		return true
	}

	return false
}

// SetNetworkCount gets a reference to the given int64 and assigns it to the NetworkCount field.
func (o *VirtualizationVmwareDatacenter) SetNetworkCount(v int64) {
	o.NetworkCount = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationVmwareDatacenter) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetHypervisorManager() VirtualizationVmwareVcenterRelationship {
	if o == nil || o.HypervisorManager == nil {
		var ret VirtualizationVmwareVcenterRelationship
		return ret
	}
	return *o.HypervisorManager
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool) {
	if o == nil || o.HypervisorManager == nil {
		return nil, false
	}
	return o.HypervisorManager, true
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager != nil {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given VirtualizationVmwareVcenterRelationship and assigns it to the HypervisorManager field.
func (o *VirtualizationVmwareDatacenter) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship) {
	o.HypervisorManager = &v
}

func (o VirtualizationVmwareDatacenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDatacenter, errVirtualizationBaseDatacenter := json.Marshal(o.VirtualizationBaseDatacenter)
	if errVirtualizationBaseDatacenter != nil {
		return []byte{}, errVirtualizationBaseDatacenter
	}
	errVirtualizationBaseDatacenter = json.Unmarshal([]byte(serializedVirtualizationBaseDatacenter), &toSerialize)
	if errVirtualizationBaseDatacenter != nil {
		return []byte{}, errVirtualizationBaseDatacenter
	}
	if o.ClusterCount != nil {
		toSerialize["ClusterCount"] = o.ClusterCount
	}
	if o.DatastoreCount != nil {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if o.HostCount != nil {
		toSerialize["HostCount"] = o.HostCount
	}
	if o.NetworkCount != nil {
		toSerialize["NetworkCount"] = o.NetworkCount
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.HypervisorManager != nil {
		toSerialize["HypervisorManager"] = o.HypervisorManager
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationVmwareDatacenter struct {
	value *VirtualizationVmwareDatacenter
	isSet bool
}

func (v NullableVirtualizationVmwareDatacenter) Get() *VirtualizationVmwareDatacenter {
	return v.value
}

func (v *NullableVirtualizationVmwareDatacenter) Set(val *VirtualizationVmwareDatacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatacenter(val *VirtualizationVmwareDatacenter) *NullableVirtualizationVmwareDatacenter {
	return &NullableVirtualizationVmwareDatacenter{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
