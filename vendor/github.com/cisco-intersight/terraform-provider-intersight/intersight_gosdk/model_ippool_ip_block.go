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

// IppoolIpBlock A block of IPv4 addresses.
type IppoolIpBlock struct {
	PoolAbstractBlockType
	// First IPv4 address of the block.
	From *string `json:"From,omitempty"`
	// Last IPv4 address of the block.
	To *string `json:"To,omitempty"`
}

// NewIppoolIpBlock instantiates a new IppoolIpBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpBlock() *IppoolIpBlock {
	this := IppoolIpBlock{}
	return &this
}

// NewIppoolIpBlockWithDefaults instantiates a new IppoolIpBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpBlockWithDefaults() *IppoolIpBlock {
	this := IppoolIpBlock{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *IppoolIpBlock) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpBlock) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *IppoolIpBlock) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *IppoolIpBlock) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *IppoolIpBlock) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpBlock) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *IppoolIpBlock) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *IppoolIpBlock) SetTo(v string) {
	o.To = &v
}

func (o IppoolIpBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlockType, errPoolAbstractBlockType := json.Marshal(o.PoolAbstractBlockType)
	if errPoolAbstractBlockType != nil {
		return []byte{}, errPoolAbstractBlockType
	}
	errPoolAbstractBlockType = json.Unmarshal([]byte(serializedPoolAbstractBlockType), &toSerialize)
	if errPoolAbstractBlockType != nil {
		return []byte{}, errPoolAbstractBlockType
	}
	if o.From != nil {
		toSerialize["From"] = o.From
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableIppoolIpBlock struct {
	value *IppoolIpBlock
	isSet bool
}

func (v NullableIppoolIpBlock) Get() *IppoolIpBlock {
	return v.value
}

func (v *NullableIppoolIpBlock) Set(val *IppoolIpBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpBlock(val *IppoolIpBlock) *NullableIppoolIpBlock {
	return &NullableIppoolIpBlock{value: val, isSet: true}
}

func (v NullableIppoolIpBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
