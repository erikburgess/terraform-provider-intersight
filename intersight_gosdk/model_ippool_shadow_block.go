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

// IppoolShadowBlock A block of Contiguous IP addresses that are part of a shadow pool.
type IppoolShadowBlock struct {
	PoolAbstractBlock
	IpV4Block            *IppoolIpBlock                `json:"IpV4Block,omitempty"`
	Pool                 *IppoolShadowPoolRelationship `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolShadowBlock IppoolShadowBlock

// NewIppoolShadowBlock instantiates a new IppoolShadowBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolShadowBlock() *IppoolShadowBlock {
	this := IppoolShadowBlock{}
	return &this
}

// NewIppoolShadowBlockWithDefaults instantiates a new IppoolShadowBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolShadowBlockWithDefaults() *IppoolShadowBlock {
	this := IppoolShadowBlock{}
	return &this
}

// GetIpV4Block returns the IpV4Block field value if set, zero value otherwise.
func (o *IppoolShadowBlock) GetIpV4Block() IppoolIpBlock {
	if o == nil || o.IpV4Block == nil {
		var ret IppoolIpBlock
		return ret
	}
	return *o.IpV4Block
}

// GetIpV4BlockOk returns a tuple with the IpV4Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetIpV4BlockOk() (*IppoolIpBlock, bool) {
	if o == nil || o.IpV4Block == nil {
		return nil, false
	}
	return o.IpV4Block, true
}

// HasIpV4Block returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasIpV4Block() bool {
	if o != nil && o.IpV4Block != nil {
		return true
	}

	return false
}

// SetIpV4Block gets a reference to the given IppoolIpBlock and assigns it to the IpV4Block field.
func (o *IppoolShadowBlock) SetIpV4Block(v IppoolIpBlock) {
	o.IpV4Block = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolShadowBlock) GetPool() IppoolShadowPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolShadowPoolRelationship and assigns it to the Pool field.
func (o *IppoolShadowBlock) SetPool(v IppoolShadowPoolRelationship) {
	o.Pool = &v
}

func (o IppoolShadowBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlock, errPoolAbstractBlock := json.Marshal(o.PoolAbstractBlock)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	errPoolAbstractBlock = json.Unmarshal([]byte(serializedPoolAbstractBlock), &toSerialize)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	if o.IpV4Block != nil {
		toSerialize["IpV4Block"] = o.IpV4Block
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolShadowBlock) UnmarshalJSON(bytes []byte) (err error) {
	type IppoolShadowBlockWithoutEmbeddedStruct struct {
		IpV4Block *IppoolIpBlock                `json:"IpV4Block,omitempty"`
		Pool      *IppoolShadowPoolRelationship `json:"Pool,omitempty"`
	}

	varIppoolShadowBlockWithoutEmbeddedStruct := IppoolShadowBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIppoolShadowBlockWithoutEmbeddedStruct)
	if err == nil {
		varIppoolShadowBlock := _IppoolShadowBlock{}
		varIppoolShadowBlock.IpV4Block = varIppoolShadowBlockWithoutEmbeddedStruct.IpV4Block
		varIppoolShadowBlock.Pool = varIppoolShadowBlockWithoutEmbeddedStruct.Pool
		*o = IppoolShadowBlock(varIppoolShadowBlock)
	} else {
		return err
	}

	varIppoolShadowBlock := _IppoolShadowBlock{}

	err = json.Unmarshal(bytes, &varIppoolShadowBlock)
	if err == nil {
		o.PoolAbstractBlock = varIppoolShadowBlock.PoolAbstractBlock
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IpV4Block")
		delete(additionalProperties, "Pool")

		// remove fields from embedded structs
		reflectPoolAbstractBlock := reflect.ValueOf(o.PoolAbstractBlock)
		for i := 0; i < reflectPoolAbstractBlock.Type().NumField(); i++ {
			t := reflectPoolAbstractBlock.Type().Field(i)

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

type NullableIppoolShadowBlock struct {
	value *IppoolShadowBlock
	isSet bool
}

func (v NullableIppoolShadowBlock) Get() *IppoolShadowBlock {
	return v.value
}

func (v *NullableIppoolShadowBlock) Set(val *IppoolShadowBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolShadowBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolShadowBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolShadowBlock(val *IppoolShadowBlock) *NullableIppoolShadowBlock {
	return &NullableIppoolShadowBlock{value: val, isSet: true}
}

func (v NullableIppoolShadowBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolShadowBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
