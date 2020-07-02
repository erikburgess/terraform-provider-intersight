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

// ComputeRackUnitIdentity Identity object that uniquely represents a server object under a DR.
type ComputeRackUnitIdentity struct {
	EquipmentPhysicalIdentity
	// Serial Identifier of an adapter participating in SWM.
	AdapterSerial *string `json:"AdapterSerial,omitempty"`
}

// NewComputeRackUnitIdentity instantiates a new ComputeRackUnitIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeRackUnitIdentity() *ComputeRackUnitIdentity {
	this := ComputeRackUnitIdentity{}
	return &this
}

// NewComputeRackUnitIdentityWithDefaults instantiates a new ComputeRackUnitIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeRackUnitIdentityWithDefaults() *ComputeRackUnitIdentity {
	this := ComputeRackUnitIdentity{}
	return &this
}

// GetAdapterSerial returns the AdapterSerial field value if set, zero value otherwise.
func (o *ComputeRackUnitIdentity) GetAdapterSerial() string {
	if o == nil || o.AdapterSerial == nil {
		var ret string
		return ret
	}
	return *o.AdapterSerial
}

// GetAdapterSerialOk returns a tuple with the AdapterSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitIdentity) GetAdapterSerialOk() (*string, bool) {
	if o == nil || o.AdapterSerial == nil {
		return nil, false
	}
	return o.AdapterSerial, true
}

// HasAdapterSerial returns a boolean if a field has been set.
func (o *ComputeRackUnitIdentity) HasAdapterSerial() bool {
	if o != nil && o.AdapterSerial != nil {
		return true
	}

	return false
}

// SetAdapterSerial gets a reference to the given string and assigns it to the AdapterSerial field.
func (o *ComputeRackUnitIdentity) SetAdapterSerial(v string) {
	o.AdapterSerial = &v
}

func (o ComputeRackUnitIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentPhysicalIdentity, errEquipmentPhysicalIdentity := json.Marshal(o.EquipmentPhysicalIdentity)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	errEquipmentPhysicalIdentity = json.Unmarshal([]byte(serializedEquipmentPhysicalIdentity), &toSerialize)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	if o.AdapterSerial != nil {
		toSerialize["AdapterSerial"] = o.AdapterSerial
	}
	return json.Marshal(toSerialize)
}

type NullableComputeRackUnitIdentity struct {
	value *ComputeRackUnitIdentity
	isSet bool
}

func (v NullableComputeRackUnitIdentity) Get() *ComputeRackUnitIdentity {
	return v.value
}

func (v *NullableComputeRackUnitIdentity) Set(val *ComputeRackUnitIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeRackUnitIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeRackUnitIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeRackUnitIdentity(val *ComputeRackUnitIdentity) *NullableComputeRackUnitIdentity {
	return &NullableComputeRackUnitIdentity{value: val, isSet: true}
}

func (v NullableComputeRackUnitIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeRackUnitIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
