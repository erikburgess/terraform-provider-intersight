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

// CapabilityPortRange Type to represent ranges of ports supporting certain speeds or features on a switch.
type CapabilityPortRange struct {
	MoBaseComplexType
	// Ending Port ID in this range of ports.
	EndPortId *int64 `json:"EndPortId,omitempty"`
	// Ending Slot ID in this range of ports.
	EndSlotId *int64 `json:"EndSlotId,omitempty"`
	// Starting Port ID in this range of ports.
	StartPortId *int64 `json:"StartPortId,omitempty"`
	// Starting Slot ID in this range of ports.
	StartSlotId *int64 `json:"StartSlotId,omitempty"`
}

// NewCapabilityPortRange instantiates a new CapabilityPortRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityPortRange() *CapabilityPortRange {
	this := CapabilityPortRange{}
	return &this
}

// NewCapabilityPortRangeWithDefaults instantiates a new CapabilityPortRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityPortRangeWithDefaults() *CapabilityPortRange {
	this := CapabilityPortRange{}
	return &this
}

// GetEndPortId returns the EndPortId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetEndPortId() int64 {
	if o == nil || o.EndPortId == nil {
		var ret int64
		return ret
	}
	return *o.EndPortId
}

// GetEndPortIdOk returns a tuple with the EndPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetEndPortIdOk() (*int64, bool) {
	if o == nil || o.EndPortId == nil {
		return nil, false
	}
	return o.EndPortId, true
}

// HasEndPortId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasEndPortId() bool {
	if o != nil && o.EndPortId != nil {
		return true
	}

	return false
}

// SetEndPortId gets a reference to the given int64 and assigns it to the EndPortId field.
func (o *CapabilityPortRange) SetEndPortId(v int64) {
	o.EndPortId = &v
}

// GetEndSlotId returns the EndSlotId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetEndSlotId() int64 {
	if o == nil || o.EndSlotId == nil {
		var ret int64
		return ret
	}
	return *o.EndSlotId
}

// GetEndSlotIdOk returns a tuple with the EndSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetEndSlotIdOk() (*int64, bool) {
	if o == nil || o.EndSlotId == nil {
		return nil, false
	}
	return o.EndSlotId, true
}

// HasEndSlotId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasEndSlotId() bool {
	if o != nil && o.EndSlotId != nil {
		return true
	}

	return false
}

// SetEndSlotId gets a reference to the given int64 and assigns it to the EndSlotId field.
func (o *CapabilityPortRange) SetEndSlotId(v int64) {
	o.EndSlotId = &v
}

// GetStartPortId returns the StartPortId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetStartPortId() int64 {
	if o == nil || o.StartPortId == nil {
		var ret int64
		return ret
	}
	return *o.StartPortId
}

// GetStartPortIdOk returns a tuple with the StartPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetStartPortIdOk() (*int64, bool) {
	if o == nil || o.StartPortId == nil {
		return nil, false
	}
	return o.StartPortId, true
}

// HasStartPortId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasStartPortId() bool {
	if o != nil && o.StartPortId != nil {
		return true
	}

	return false
}

// SetStartPortId gets a reference to the given int64 and assigns it to the StartPortId field.
func (o *CapabilityPortRange) SetStartPortId(v int64) {
	o.StartPortId = &v
}

// GetStartSlotId returns the StartSlotId field value if set, zero value otherwise.
func (o *CapabilityPortRange) GetStartSlotId() int64 {
	if o == nil || o.StartSlotId == nil {
		var ret int64
		return ret
	}
	return *o.StartSlotId
}

// GetStartSlotIdOk returns a tuple with the StartSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortRange) GetStartSlotIdOk() (*int64, bool) {
	if o == nil || o.StartSlotId == nil {
		return nil, false
	}
	return o.StartSlotId, true
}

// HasStartSlotId returns a boolean if a field has been set.
func (o *CapabilityPortRange) HasStartSlotId() bool {
	if o != nil && o.StartSlotId != nil {
		return true
	}

	return false
}

// SetStartSlotId gets a reference to the given int64 and assigns it to the StartSlotId field.
func (o *CapabilityPortRange) SetStartSlotId(v int64) {
	o.StartSlotId = &v
}

func (o CapabilityPortRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.EndPortId != nil {
		toSerialize["EndPortId"] = o.EndPortId
	}
	if o.EndSlotId != nil {
		toSerialize["EndSlotId"] = o.EndSlotId
	}
	if o.StartPortId != nil {
		toSerialize["StartPortId"] = o.StartPortId
	}
	if o.StartSlotId != nil {
		toSerialize["StartSlotId"] = o.StartSlotId
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilityPortRange struct {
	value *CapabilityPortRange
	isSet bool
}

func (v NullableCapabilityPortRange) Get() *CapabilityPortRange {
	return v.value
}

func (v *NullableCapabilityPortRange) Set(val *CapabilityPortRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityPortRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityPortRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityPortRange(val *CapabilityPortRange) *NullableCapabilityPortRange {
	return &NullableCapabilityPortRange{value: val, isSet: true}
}

func (v NullableCapabilityPortRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityPortRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
