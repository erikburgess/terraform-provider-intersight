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

// StorageSasPortAllOf Definition of the list of properties defined in 'storage.SasPort', excluding properties defined in parent classes.
type StorageSasPortAllOf struct {
	// The SAS Address assigned to storage port.
	Address *string `json:"Address,omitempty"`
	// The unique disk identifier.
	DiskId *int64 `json:"DiskId,omitempty"`
	// The end-point Id assigned to storage port.
	EndPointId *int64 `json:"EndPointId,omitempty"`
	// The description for the link.
	LinkDescription *string `json:"LinkDescription,omitempty"`
	// The link speed negotiated for communication.
	LinkSpeed           *string                              `json:"LinkSpeed,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StoragePhysicalDisk *StoragePhysicalDiskRelationship     `json:"StoragePhysicalDisk,omitempty"`
}

// NewStorageSasPortAllOf instantiates a new StorageSasPortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSasPortAllOf() *StorageSasPortAllOf {
	this := StorageSasPortAllOf{}
	return &this
}

// NewStorageSasPortAllOfWithDefaults instantiates a new StorageSasPortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSasPortAllOfWithDefaults() *StorageSasPortAllOf {
	this := StorageSasPortAllOf{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *StorageSasPortAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetDiskId() int64 {
	if o == nil || o.DiskId == nil {
		var ret int64
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetDiskIdOk() (*int64, bool) {
	if o == nil || o.DiskId == nil {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasDiskId() bool {
	if o != nil && o.DiskId != nil {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given int64 and assigns it to the DiskId field.
func (o *StorageSasPortAllOf) SetDiskId(v int64) {
	o.DiskId = &v
}

// GetEndPointId returns the EndPointId field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetEndPointId() int64 {
	if o == nil || o.EndPointId == nil {
		var ret int64
		return ret
	}
	return *o.EndPointId
}

// GetEndPointIdOk returns a tuple with the EndPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetEndPointIdOk() (*int64, bool) {
	if o == nil || o.EndPointId == nil {
		return nil, false
	}
	return o.EndPointId, true
}

// HasEndPointId returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasEndPointId() bool {
	if o != nil && o.EndPointId != nil {
		return true
	}

	return false
}

// SetEndPointId gets a reference to the given int64 and assigns it to the EndPointId field.
func (o *StorageSasPortAllOf) SetEndPointId(v int64) {
	o.EndPointId = &v
}

// GetLinkDescription returns the LinkDescription field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetLinkDescription() string {
	if o == nil || o.LinkDescription == nil {
		var ret string
		return ret
	}
	return *o.LinkDescription
}

// GetLinkDescriptionOk returns a tuple with the LinkDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetLinkDescriptionOk() (*string, bool) {
	if o == nil || o.LinkDescription == nil {
		return nil, false
	}
	return o.LinkDescription, true
}

// HasLinkDescription returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasLinkDescription() bool {
	if o != nil && o.LinkDescription != nil {
		return true
	}

	return false
}

// SetLinkDescription gets a reference to the given string and assigns it to the LinkDescription field.
func (o *StorageSasPortAllOf) SetLinkDescription(v string) {
	o.LinkDescription = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetLinkSpeed() string {
	if o == nil || o.LinkSpeed == nil {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetLinkSpeedOk() (*string, bool) {
	if o == nil || o.LinkSpeed == nil {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasLinkSpeed() bool {
	if o != nil && o.LinkSpeed != nil {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *StorageSasPortAllOf) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageSasPortAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageSasPortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStoragePhysicalDisk returns the StoragePhysicalDisk field value if set, zero value otherwise.
func (o *StorageSasPortAllOf) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.StoragePhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.StoragePhysicalDisk
}

// GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPortAllOf) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.StoragePhysicalDisk == nil {
		return nil, false
	}
	return o.StoragePhysicalDisk, true
}

// HasStoragePhysicalDisk returns a boolean if a field has been set.
func (o *StorageSasPortAllOf) HasStoragePhysicalDisk() bool {
	if o != nil && o.StoragePhysicalDisk != nil {
		return true
	}

	return false
}

// SetStoragePhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the StoragePhysicalDisk field.
func (o *StorageSasPortAllOf) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.StoragePhysicalDisk = &v
}

func (o StorageSasPortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.DiskId != nil {
		toSerialize["DiskId"] = o.DiskId
	}
	if o.EndPointId != nil {
		toSerialize["EndPointId"] = o.EndPointId
	}
	if o.LinkDescription != nil {
		toSerialize["LinkDescription"] = o.LinkDescription
	}
	if o.LinkSpeed != nil {
		toSerialize["LinkSpeed"] = o.LinkSpeed
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StoragePhysicalDisk != nil {
		toSerialize["StoragePhysicalDisk"] = o.StoragePhysicalDisk
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSasPortAllOf struct {
	value *StorageSasPortAllOf
	isSet bool
}

func (v NullableStorageSasPortAllOf) Get() *StorageSasPortAllOf {
	return v.value
}

func (v *NullableStorageSasPortAllOf) Set(val *StorageSasPortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSasPortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSasPortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSasPortAllOf(val *StorageSasPortAllOf) *NullableStorageSasPortAllOf {
	return &NullableStorageSasPortAllOf{value: val, isSet: true}
}

func (v NullableStorageSasPortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSasPortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
