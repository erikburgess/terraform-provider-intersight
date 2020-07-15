/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricPortBase Abstract representation of Switch/FEX/Chassis Port.
type FabricPortBase struct {
	MoBaseMo
	// Breakout port Identifier of the Switch Interface.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Port Identifier of the Switch/FEX/Chassis Interface.
	PortId *int64 `json:"PortId,omitempty"`
	// Slot Identifier of the Switch/FEX/Chassis Interface.
	SlotId               *int64 `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPortBase FabricPortBase

// NewFabricPortBase instantiates a new FabricPortBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortBase() *FabricPortBase {
	this := FabricPortBase{}
	return &this
}

// NewFabricPortBaseWithDefaults instantiates a new FabricPortBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortBaseWithDefaults() *FabricPortBase {
	this := FabricPortBase{}
	return &this
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *FabricPortBase) GetAggregatePortId() int64 {
	if o == nil || o.AggregatePortId == nil {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortBase) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || o.AggregatePortId == nil {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *FabricPortBase) HasAggregatePortId() bool {
	if o != nil && o.AggregatePortId != nil {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *FabricPortBase) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FabricPortBase) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortBase) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FabricPortBase) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *FabricPortBase) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *FabricPortBase) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortBase) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *FabricPortBase) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *FabricPortBase) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o FabricPortBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.AggregatePortId != nil {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricPortBase) UnmarshalJSON(bytes []byte) (err error) {
	type FabricPortBaseWithoutEmbeddedStruct struct {
		// Breakout port Identifier of the Switch Interface.
		AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
		// Port Identifier of the Switch/FEX/Chassis Interface.
		PortId *int64 `json:"PortId,omitempty"`
		// Slot Identifier of the Switch/FEX/Chassis Interface.
		SlotId *int64 `json:"SlotId,omitempty"`
	}

	varFabricPortBaseWithoutEmbeddedStruct := FabricPortBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricPortBaseWithoutEmbeddedStruct)
	if err == nil {
		varFabricPortBase := _FabricPortBase{}
		varFabricPortBase.AggregatePortId = varFabricPortBaseWithoutEmbeddedStruct.AggregatePortId
		varFabricPortBase.PortId = varFabricPortBaseWithoutEmbeddedStruct.PortId
		varFabricPortBase.SlotId = varFabricPortBaseWithoutEmbeddedStruct.SlotId
		*o = FabricPortBase(varFabricPortBase)
	} else {
		return err
	}

	varFabricPortBase := _FabricPortBase{}

	err = json.Unmarshal(bytes, &varFabricPortBase)
	if err == nil {
		o.MoBaseMo = varFabricPortBase.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "SlotId")

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

type NullableFabricPortBase struct {
	value *FabricPortBase
	isSet bool
}

func (v NullableFabricPortBase) Get() *FabricPortBase {
	return v.value
}

func (v *NullableFabricPortBase) Set(val *FabricPortBase) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortBase) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortBase(val *FabricPortBase) *NullableFabricPortBase {
	return &NullableFabricPortBase{value: val, isSet: true}
}

func (v NullableFabricPortBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
