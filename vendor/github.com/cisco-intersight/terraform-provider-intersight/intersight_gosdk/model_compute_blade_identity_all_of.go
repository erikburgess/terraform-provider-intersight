/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ComputeBladeIdentityAllOf Definition of the list of properties defined in 'compute.BladeIdentity', excluding properties defined in parent classes.
type ComputeBladeIdentityAllOf struct {
	// Chassis Identifier of a blade server.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// FI Device registration Mo ID.
	DeviceMoId *string `json:"DeviceMoId,omitempty"`
	// Indicates pending discovery flag.
	PendingDiscovery *string `json:"PendingDiscovery,omitempty"`
	// Chassis slot number of a blade server.
	SlotId               *int64 `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBladeIdentityAllOf ComputeBladeIdentityAllOf

// NewComputeBladeIdentityAllOf instantiates a new ComputeBladeIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBladeIdentityAllOf() *ComputeBladeIdentityAllOf {
	this := ComputeBladeIdentityAllOf{}
	return &this
}

// NewComputeBladeIdentityAllOfWithDefaults instantiates a new ComputeBladeIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBladeIdentityAllOfWithDefaults() *ComputeBladeIdentityAllOf {
	this := ComputeBladeIdentityAllOf{}
	return &this
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ComputeBladeIdentityAllOf) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentityAllOf) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ComputeBladeIdentityAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *ComputeBladeIdentityAllOf) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetDeviceMoId returns the DeviceMoId field value if set, zero value otherwise.
func (o *ComputeBladeIdentityAllOf) GetDeviceMoId() string {
	if o == nil || o.DeviceMoId == nil {
		var ret string
		return ret
	}
	return *o.DeviceMoId
}

// GetDeviceMoIdOk returns a tuple with the DeviceMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentityAllOf) GetDeviceMoIdOk() (*string, bool) {
	if o == nil || o.DeviceMoId == nil {
		return nil, false
	}
	return o.DeviceMoId, true
}

// HasDeviceMoId returns a boolean if a field has been set.
func (o *ComputeBladeIdentityAllOf) HasDeviceMoId() bool {
	if o != nil && o.DeviceMoId != nil {
		return true
	}

	return false
}

// SetDeviceMoId gets a reference to the given string and assigns it to the DeviceMoId field.
func (o *ComputeBladeIdentityAllOf) SetDeviceMoId(v string) {
	o.DeviceMoId = &v
}

// GetPendingDiscovery returns the PendingDiscovery field value if set, zero value otherwise.
func (o *ComputeBladeIdentityAllOf) GetPendingDiscovery() string {
	if o == nil || o.PendingDiscovery == nil {
		var ret string
		return ret
	}
	return *o.PendingDiscovery
}

// GetPendingDiscoveryOk returns a tuple with the PendingDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentityAllOf) GetPendingDiscoveryOk() (*string, bool) {
	if o == nil || o.PendingDiscovery == nil {
		return nil, false
	}
	return o.PendingDiscovery, true
}

// HasPendingDiscovery returns a boolean if a field has been set.
func (o *ComputeBladeIdentityAllOf) HasPendingDiscovery() bool {
	if o != nil && o.PendingDiscovery != nil {
		return true
	}

	return false
}

// SetPendingDiscovery gets a reference to the given string and assigns it to the PendingDiscovery field.
func (o *ComputeBladeIdentityAllOf) SetPendingDiscovery(v string) {
	o.PendingDiscovery = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ComputeBladeIdentityAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentityAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ComputeBladeIdentityAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ComputeBladeIdentityAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o ComputeBladeIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.DeviceMoId != nil {
		toSerialize["DeviceMoId"] = o.DeviceMoId
	}
	if o.PendingDiscovery != nil {
		toSerialize["PendingDiscovery"] = o.PendingDiscovery
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBladeIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeBladeIdentityAllOf := _ComputeBladeIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varComputeBladeIdentityAllOf); err == nil {
		*o = ComputeBladeIdentityAllOf(varComputeBladeIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "DeviceMoId")
		delete(additionalProperties, "PendingDiscovery")
		delete(additionalProperties, "SlotId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeBladeIdentityAllOf struct {
	value *ComputeBladeIdentityAllOf
	isSet bool
}

func (v NullableComputeBladeIdentityAllOf) Get() *ComputeBladeIdentityAllOf {
	return v.value
}

func (v *NullableComputeBladeIdentityAllOf) Set(val *ComputeBladeIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBladeIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBladeIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBladeIdentityAllOf(val *ComputeBladeIdentityAllOf) *NullableComputeBladeIdentityAllOf {
	return &NullableComputeBladeIdentityAllOf{value: val, isSet: true}
}

func (v NullableComputeBladeIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBladeIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
