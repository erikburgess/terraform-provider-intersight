/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-08T07:46:03Z.
 *
 * API version: 0.0.1-15983
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// HyperflexVmDiskAllOf Definition of the list of properties defined in 'hyperflex.VmDisk', excluding properties defined in parent classes.
type HyperflexVmDiskAllOf struct {
	// Boot order for this disk.
	BootOrder *int64 `json:"BootOrder,omitempty"`
	// Virtual machine brige name of interface.
	Bus *string `json:"Bus,omitempty"`
	// Disk name associated with virtual machine.
	Name *string `json:"Name,omitempty"`
	// Type of the Disk (hdd, cdrom).
	Type        *string               `json:"Type,omitempty"`
	VirtualDisk *HyperflexVdiskConfig `json:"VirtualDisk,omitempty"`
	// Virtual disk reference name.
	VirtualDiskReference *string `json:"VirtualDiskReference,omitempty"`
}

// NewHyperflexVmDiskAllOf instantiates a new HyperflexVmDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmDiskAllOf() *HyperflexVmDiskAllOf {
	this := HyperflexVmDiskAllOf{}
	var bus string = "virtio"
	this.Bus = &bus
	var type_ string = "hdd"
	this.Type = &type_
	return &this
}

// NewHyperflexVmDiskAllOfWithDefaults instantiates a new HyperflexVmDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmDiskAllOfWithDefaults() *HyperflexVmDiskAllOf {
	this := HyperflexVmDiskAllOf{}
	var bus string = "virtio"
	this.Bus = &bus
	var type_ string = "hdd"
	this.Type = &type_
	return &this
}

// GetBootOrder returns the BootOrder field value if set, zero value otherwise.
func (o *HyperflexVmDiskAllOf) GetBootOrder() int64 {
	if o == nil || o.BootOrder == nil {
		var ret int64
		return ret
	}
	return *o.BootOrder
}

// GetBootOrderOk returns a tuple with the BootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmDiskAllOf) GetBootOrderOk() (*int64, bool) {
	if o == nil || o.BootOrder == nil {
		return nil, false
	}
	return o.BootOrder, true
}

// HasBootOrder returns a boolean if a field has been set.
func (o *HyperflexVmDiskAllOf) HasBootOrder() bool {
	if o != nil && o.BootOrder != nil {
		return true
	}

	return false
}

// SetBootOrder gets a reference to the given int64 and assigns it to the BootOrder field.
func (o *HyperflexVmDiskAllOf) SetBootOrder(v int64) {
	o.BootOrder = &v
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *HyperflexVmDiskAllOf) GetBus() string {
	if o == nil || o.Bus == nil {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmDiskAllOf) GetBusOk() (*string, bool) {
	if o == nil || o.Bus == nil {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *HyperflexVmDiskAllOf) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *HyperflexVmDiskAllOf) SetBus(v string) {
	o.Bus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexVmDiskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmDiskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexVmDiskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexVmDiskAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HyperflexVmDiskAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmDiskAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HyperflexVmDiskAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HyperflexVmDiskAllOf) SetType(v string) {
	o.Type = &v
}

// GetVirtualDisk returns the VirtualDisk field value if set, zero value otherwise.
func (o *HyperflexVmDiskAllOf) GetVirtualDisk() HyperflexVdiskConfig {
	if o == nil || o.VirtualDisk == nil {
		var ret HyperflexVdiskConfig
		return ret
	}
	return *o.VirtualDisk
}

// GetVirtualDiskOk returns a tuple with the VirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmDiskAllOf) GetVirtualDiskOk() (*HyperflexVdiskConfig, bool) {
	if o == nil || o.VirtualDisk == nil {
		return nil, false
	}
	return o.VirtualDisk, true
}

// HasVirtualDisk returns a boolean if a field has been set.
func (o *HyperflexVmDiskAllOf) HasVirtualDisk() bool {
	if o != nil && o.VirtualDisk != nil {
		return true
	}

	return false
}

// SetVirtualDisk gets a reference to the given HyperflexVdiskConfig and assigns it to the VirtualDisk field.
func (o *HyperflexVmDiskAllOf) SetVirtualDisk(v HyperflexVdiskConfig) {
	o.VirtualDisk = &v
}

// GetVirtualDiskReference returns the VirtualDiskReference field value if set, zero value otherwise.
func (o *HyperflexVmDiskAllOf) GetVirtualDiskReference() string {
	if o == nil || o.VirtualDiskReference == nil {
		var ret string
		return ret
	}
	return *o.VirtualDiskReference
}

// GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmDiskAllOf) GetVirtualDiskReferenceOk() (*string, bool) {
	if o == nil || o.VirtualDiskReference == nil {
		return nil, false
	}
	return o.VirtualDiskReference, true
}

// HasVirtualDiskReference returns a boolean if a field has been set.
func (o *HyperflexVmDiskAllOf) HasVirtualDiskReference() bool {
	if o != nil && o.VirtualDiskReference != nil {
		return true
	}

	return false
}

// SetVirtualDiskReference gets a reference to the given string and assigns it to the VirtualDiskReference field.
func (o *HyperflexVmDiskAllOf) SetVirtualDiskReference(v string) {
	o.VirtualDiskReference = &v
}

func (o HyperflexVmDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootOrder != nil {
		toSerialize["BootOrder"] = o.BootOrder
	}
	if o.Bus != nil {
		toSerialize["Bus"] = o.Bus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDisk != nil {
		toSerialize["VirtualDisk"] = o.VirtualDisk
	}
	if o.VirtualDiskReference != nil {
		toSerialize["VirtualDiskReference"] = o.VirtualDiskReference
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexVmDiskAllOf struct {
	value *HyperflexVmDiskAllOf
	isSet bool
}

func (v NullableHyperflexVmDiskAllOf) Get() *HyperflexVmDiskAllOf {
	return v.value
}

func (v *NullableHyperflexVmDiskAllOf) Set(val *HyperflexVmDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmDiskAllOf(val *HyperflexVmDiskAllOf) *NullableHyperflexVmDiskAllOf {
	return &NullableHyperflexVmDiskAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVmDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}