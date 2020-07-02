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

// StoragePureHostGroup A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
type StoragePureHostGroup struct {
	StorageBaseHostGroup
	HostNames          *[]string                     `json:"HostNames,omitempty"`
	StorageUtilization *StoragePureHostUtilization   `json:"StorageUtilization,omitempty"`
	Array              *StoragePureArrayRelationship `json:"Array,omitempty"`
	// An array of relationships to storagePureHost resources.
	Hosts            []StoragePureHostRelationship           `json:"Hosts,omitempty"`
	ProtectionGroup  *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
}

// NewStoragePureHostGroup instantiates a new StoragePureHostGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostGroup() *StoragePureHostGroup {
	this := StoragePureHostGroup{}
	return &this
}

// NewStoragePureHostGroupWithDefaults instantiates a new StoragePureHostGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostGroupWithDefaults() *StoragePureHostGroup {
	this := StoragePureHostGroup{}
	return &this
}

// GetHostNames returns the HostNames field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetHostNames() []string {
	if o == nil || o.HostNames == nil {
		var ret []string
		return ret
	}
	return *o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetHostNamesOk() (*[]string, bool) {
	if o == nil || o.HostNames == nil {
		return nil, false
	}
	return o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasHostNames() bool {
	if o != nil && o.HostNames != nil {
		return true
	}

	return false
}

// SetHostNames gets a reference to the given []string and assigns it to the HostNames field.
func (o *StoragePureHostGroup) SetHostNames(v []string) {
	o.HostNames = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetStorageUtilization() StoragePureHostUtilization {
	if o == nil || o.StorageUtilization == nil {
		var ret StoragePureHostUtilization
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetStorageUtilizationOk() (*StoragePureHostUtilization, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StoragePureHostUtilization and assigns it to the StorageUtilization field.
func (o *StoragePureHostGroup) SetStorageUtilization(v StoragePureHostUtilization) {
	o.StorageUtilization = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHostGroup) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostGroup) GetHosts() []StoragePureHostRelationship {
	if o == nil {
		var ret []StoragePureHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostGroup) GetHostsOk() (*[]StoragePureHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []StoragePureHostRelationship and assigns it to the Hosts field.
func (o *StoragePureHostGroup) SetHosts(v []StoragePureHostRelationship) {
	o.Hosts = v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureHostGroup) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHostGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureHostGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostGroup, errStorageBaseHostGroup := json.Marshal(o.StorageBaseHostGroup)
	if errStorageBaseHostGroup != nil {
		return []byte{}, errStorageBaseHostGroup
	}
	errStorageBaseHostGroup = json.Unmarshal([]byte(serializedStorageBaseHostGroup), &toSerialize)
	if errStorageBaseHostGroup != nil {
		return []byte{}, errStorageBaseHostGroup
	}
	if o.HostNames != nil {
		toSerialize["HostNames"] = o.HostNames
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePureHostGroup struct {
	value *StoragePureHostGroup
	isSet bool
}

func (v NullableStoragePureHostGroup) Get() *StoragePureHostGroup {
	return v.value
}

func (v *NullableStoragePureHostGroup) Set(val *StoragePureHostGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostGroup(val *StoragePureHostGroup) *NullableStoragePureHostGroup {
	return &NullableStoragePureHostGroup{value: val, isSet: true}
}

func (v NullableStoragePureHostGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
