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

// StoragePureHostGroupAllOf Definition of the list of properties defined in 'storage.PureHostGroup', excluding properties defined in parent classes.
type StoragePureHostGroupAllOf struct {
	HostNames          *[]string                     `json:"HostNames,omitempty"`
	StorageUtilization *StoragePureHostUtilization   `json:"StorageUtilization,omitempty"`
	Array              *StoragePureArrayRelationship `json:"Array,omitempty"`
	// An array of relationships to storagePureHost resources.
	Hosts            []StoragePureHostRelationship           `json:"Hosts,omitempty"`
	ProtectionGroup  *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
}

// NewStoragePureHostGroupAllOf instantiates a new StoragePureHostGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostGroupAllOf() *StoragePureHostGroupAllOf {
	this := StoragePureHostGroupAllOf{}
	return &this
}

// NewStoragePureHostGroupAllOfWithDefaults instantiates a new StoragePureHostGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostGroupAllOfWithDefaults() *StoragePureHostGroupAllOf {
	this := StoragePureHostGroupAllOf{}
	return &this
}

// GetHostNames returns the HostNames field value if set, zero value otherwise.
func (o *StoragePureHostGroupAllOf) GetHostNames() []string {
	if o == nil || o.HostNames == nil {
		var ret []string
		return ret
	}
	return *o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroupAllOf) GetHostNamesOk() (*[]string, bool) {
	if o == nil || o.HostNames == nil {
		return nil, false
	}
	return o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *StoragePureHostGroupAllOf) HasHostNames() bool {
	if o != nil && o.HostNames != nil {
		return true
	}

	return false
}

// SetHostNames gets a reference to the given []string and assigns it to the HostNames field.
func (o *StoragePureHostGroupAllOf) SetHostNames(v []string) {
	o.HostNames = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StoragePureHostGroupAllOf) GetStorageUtilization() StoragePureHostUtilization {
	if o == nil || o.StorageUtilization == nil {
		var ret StoragePureHostUtilization
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroupAllOf) GetStorageUtilizationOk() (*StoragePureHostUtilization, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StoragePureHostGroupAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StoragePureHostUtilization and assigns it to the StorageUtilization field.
func (o *StoragePureHostGroupAllOf) SetStorageUtilization(v StoragePureHostUtilization) {
	o.StorageUtilization = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureHostGroupAllOf) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroupAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHostGroupAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHostGroupAllOf) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostGroupAllOf) GetHosts() []StoragePureHostRelationship {
	if o == nil {
		var ret []StoragePureHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostGroupAllOf) GetHostsOk() (*[]StoragePureHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *StoragePureHostGroupAllOf) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []StoragePureHostRelationship and assigns it to the Hosts field.
func (o *StoragePureHostGroupAllOf) SetHosts(v []StoragePureHostRelationship) {
	o.Hosts = v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureHostGroupAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroupAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureHostGroupAllOf) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureHostGroupAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureHostGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHostGroupAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHostGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureHostGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableStoragePureHostGroupAllOf struct {
	value *StoragePureHostGroupAllOf
	isSet bool
}

func (v NullableStoragePureHostGroupAllOf) Get() *StoragePureHostGroupAllOf {
	return v.value
}

func (v *NullableStoragePureHostGroupAllOf) Set(val *StoragePureHostGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostGroupAllOf(val *StoragePureHostGroupAllOf) *NullableStoragePureHostGroupAllOf {
	return &NullableStoragePureHostGroupAllOf{value: val, isSet: true}
}

func (v NullableStoragePureHostGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
