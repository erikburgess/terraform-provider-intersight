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
	"reflect"
	"strings"
)

// InventoryDnMoBinding DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
type InventoryDnMoBinding struct {
	MoBaseMo
	// The Distinguished Name for this object, used to uniquely identify this object.
	Dn *string `json:"Dn,omitempty"`
	// The MO ID of the target MO for this particular Distinguished Name (dn).
	TargetMoId *string `json:"TargetMoId,omitempty"`
	// The type of the target MO for this particular Distinguished Name (dn).
	TargetMoType         *string                              `json:"TargetMoType,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryDnMoBinding InventoryDnMoBinding

// NewInventoryDnMoBinding instantiates a new InventoryDnMoBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDnMoBinding() *InventoryDnMoBinding {
	this := InventoryDnMoBinding{}
	return &this
}

// NewInventoryDnMoBindingWithDefaults instantiates a new InventoryDnMoBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDnMoBindingWithDefaults() *InventoryDnMoBinding {
	this := InventoryDnMoBinding{}
	return &this
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *InventoryDnMoBinding) SetDn(v string) {
	o.Dn = &v
}

// GetTargetMoId returns the TargetMoId field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetTargetMoId() string {
	if o == nil || o.TargetMoId == nil {
		var ret string
		return ret
	}
	return *o.TargetMoId
}

// GetTargetMoIdOk returns a tuple with the TargetMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetTargetMoIdOk() (*string, bool) {
	if o == nil || o.TargetMoId == nil {
		return nil, false
	}
	return o.TargetMoId, true
}

// HasTargetMoId returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasTargetMoId() bool {
	if o != nil && o.TargetMoId != nil {
		return true
	}

	return false
}

// SetTargetMoId gets a reference to the given string and assigns it to the TargetMoId field.
func (o *InventoryDnMoBinding) SetTargetMoId(v string) {
	o.TargetMoId = &v
}

// GetTargetMoType returns the TargetMoType field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetTargetMoType() string {
	if o == nil || o.TargetMoType == nil {
		var ret string
		return ret
	}
	return *o.TargetMoType
}

// GetTargetMoTypeOk returns a tuple with the TargetMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetTargetMoTypeOk() (*string, bool) {
	if o == nil || o.TargetMoType == nil {
		return nil, false
	}
	return o.TargetMoType, true
}

// HasTargetMoType returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasTargetMoType() bool {
	if o != nil && o.TargetMoType != nil {
		return true
	}

	return false
}

// SetTargetMoType gets a reference to the given string and assigns it to the TargetMoType field.
func (o *InventoryDnMoBinding) SetTargetMoType(v string) {
	o.TargetMoType = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryDnMoBinding) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBinding) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryDnMoBinding) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryDnMoBinding) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryDnMoBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.TargetMoId != nil {
		toSerialize["TargetMoId"] = o.TargetMoId
	}
	if o.TargetMoType != nil {
		toSerialize["TargetMoType"] = o.TargetMoType
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryDnMoBinding) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryDnMoBindingWithoutEmbeddedStruct struct {
		// The Distinguished Name for this object, used to uniquely identify this object.
		Dn *string `json:"Dn,omitempty"`
		// The MO ID of the target MO for this particular Distinguished Name (dn).
		TargetMoId *string `json:"TargetMoId,omitempty"`
		// The type of the target MO for this particular Distinguished Name (dn).
		TargetMoType     *string                              `json:"TargetMoType,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varInventoryDnMoBindingWithoutEmbeddedStruct := InventoryDnMoBindingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryDnMoBindingWithoutEmbeddedStruct)
	if err == nil {
		varInventoryDnMoBinding := _InventoryDnMoBinding{}
		varInventoryDnMoBinding.Dn = varInventoryDnMoBindingWithoutEmbeddedStruct.Dn
		varInventoryDnMoBinding.TargetMoId = varInventoryDnMoBindingWithoutEmbeddedStruct.TargetMoId
		varInventoryDnMoBinding.TargetMoType = varInventoryDnMoBindingWithoutEmbeddedStruct.TargetMoType
		varInventoryDnMoBinding.RegisteredDevice = varInventoryDnMoBindingWithoutEmbeddedStruct.RegisteredDevice
		*o = InventoryDnMoBinding(varInventoryDnMoBinding)
	} else {
		return err
	}

	varInventoryDnMoBinding := _InventoryDnMoBinding{}

	err = json.Unmarshal(bytes, &varInventoryDnMoBinding)
	if err == nil {
		o.MoBaseMo = varInventoryDnMoBinding.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "TargetMoId")
		delete(additionalProperties, "TargetMoType")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryDnMoBinding struct {
	value *InventoryDnMoBinding
	isSet bool
}

func (v NullableInventoryDnMoBinding) Get() *InventoryDnMoBinding {
	return v.value
}

func (v *NullableInventoryDnMoBinding) Set(val *InventoryDnMoBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDnMoBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDnMoBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDnMoBinding(val *InventoryDnMoBinding) *NullableInventoryDnMoBinding {
	return &NullableInventoryDnMoBinding{value: val, isSet: true}
}

func (v NullableInventoryDnMoBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDnMoBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
