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

// VirtualizationBaseVirtualMachineAllOf Definition of the list of properties defined in 'virtualization.BaseVirtualMachine', excluding properties defined in parent classes.
type VirtualizationBaseVirtualMachineAllOf struct {
	Capacity  *InfraHardwareInfo       `json:"Capacity,omitempty"`
	GuestInfo *VirtualizationGuestInfo `json:"GuestInfo,omitempty"`
	// Type of hypervisor where the virtual machine is hosted, for example VMware ESXi.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).
	Identity       *string                       `json:"Identity,omitempty"`
	IpAddress      *[]string                     `json:"IpAddress,omitempty"`
	MemoryCapacity *VirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
	// User-provided name to identify the virtual machine.
	Name *string `json:"Name,omitempty"`
	// Power state of the virtual machine.
	PowerState        *string                        `json:"PowerState,omitempty"`
	ProcessorCapacity *VirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
	// The uuid of this virtual machine. The uuid is internally generated and not user specified.
	Uuid *string `json:"Uuid,omitempty"`
}

// NewVirtualizationBaseVirtualMachineAllOf instantiates a new VirtualizationBaseVirtualMachineAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualMachineAllOf() *VirtualizationBaseVirtualMachineAllOf {
	this := VirtualizationBaseVirtualMachineAllOf{}
	var powerState string = "Unknown"
	this.PowerState = &powerState
	return &this
}

// NewVirtualizationBaseVirtualMachineAllOfWithDefaults instantiates a new VirtualizationBaseVirtualMachineAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualMachineAllOfWithDefaults() *VirtualizationBaseVirtualMachineAllOf {
	this := VirtualizationBaseVirtualMachineAllOf{}
	var powerState string = "Unknown"
	this.PowerState = &powerState
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetCapacity() InfraHardwareInfo {
	if o == nil || o.Capacity == nil {
		var ret InfraHardwareInfo
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetCapacityOk() (*InfraHardwareInfo, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given InfraHardwareInfo and assigns it to the Capacity field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetCapacity(v InfraHardwareInfo) {
	o.Capacity = &v
}

// GetGuestInfo returns the GuestInfo field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetGuestInfo() VirtualizationGuestInfo {
	if o == nil || o.GuestInfo == nil {
		var ret VirtualizationGuestInfo
		return ret
	}
	return *o.GuestInfo
}

// GetGuestInfoOk returns a tuple with the GuestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetGuestInfoOk() (*VirtualizationGuestInfo, bool) {
	if o == nil || o.GuestInfo == nil {
		return nil, false
	}
	return o.GuestInfo, true
}

// HasGuestInfo returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasGuestInfo() bool {
	if o != nil && o.GuestInfo != nil {
		return true
	}

	return false
}

// SetGuestInfo gets a reference to the given VirtualizationGuestInfo and assigns it to the GuestInfo field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetGuestInfo(v VirtualizationGuestInfo) {
	o.GuestInfo = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetIpAddress() []string {
	if o == nil || o.IpAddress == nil {
		var ret []string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetIpAddress(v []string) {
	o.IpAddress = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetMemoryCapacity() VirtualizationMemoryCapacity {
	if o == nil || o.MemoryCapacity == nil {
		var ret VirtualizationMemoryCapacity
		return ret
	}
	return *o.MemoryCapacity
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool) {
	if o == nil || o.MemoryCapacity == nil {
		return nil, false
	}
	return o.MemoryCapacity, true
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity != nil {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given VirtualizationMemoryCapacity and assigns it to the MemoryCapacity field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetMemoryCapacity(v VirtualizationMemoryCapacity) {
	o.MemoryCapacity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetName(v string) {
	o.Name = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetPowerState(v string) {
	o.PowerState = &v
}

// GetProcessorCapacity returns the ProcessorCapacity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetProcessorCapacity() VirtualizationComputeCapacity {
	if o == nil || o.ProcessorCapacity == nil {
		var ret VirtualizationComputeCapacity
		return ret
	}
	return *o.ProcessorCapacity
}

// GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool) {
	if o == nil || o.ProcessorCapacity == nil {
		return nil, false
	}
	return o.ProcessorCapacity, true
}

// HasProcessorCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasProcessorCapacity() bool {
	if o != nil && o.ProcessorCapacity != nil {
		return true
	}

	return false
}

// SetProcessorCapacity gets a reference to the given VirtualizationComputeCapacity and assigns it to the ProcessorCapacity field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetProcessorCapacity(v VirtualizationComputeCapacity) {
	o.ProcessorCapacity = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationBaseVirtualMachineAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o VirtualizationBaseVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.GuestInfo != nil {
		toSerialize["GuestInfo"] = o.GuestInfo
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.MemoryCapacity != nil {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.ProcessorCapacity != nil {
		toSerialize["ProcessorCapacity"] = o.ProcessorCapacity
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationBaseVirtualMachineAllOf struct {
	value *VirtualizationBaseVirtualMachineAllOf
	isSet bool
}

func (v NullableVirtualizationBaseVirtualMachineAllOf) Get() *VirtualizationBaseVirtualMachineAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualMachineAllOf) Set(val *VirtualizationBaseVirtualMachineAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualMachineAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualMachineAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualMachineAllOf(val *VirtualizationBaseVirtualMachineAllOf) *NullableVirtualizationBaseVirtualMachineAllOf {
	return &NullableVirtualizationBaseVirtualMachineAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualMachineAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
