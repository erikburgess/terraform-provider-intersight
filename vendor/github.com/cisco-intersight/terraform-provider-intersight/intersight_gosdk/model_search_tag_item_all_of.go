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

// SearchTagItemAllOf Definition of the list of properties defined in 'search.TagItem', excluding properties defined in parent classes.
type SearchTagItemAllOf struct {
	// The number of times this tag key has been set across all resources.
	Count *int64 `json:"Count,omitempty"`
	// Key of the Tag from all the resources in Intersight.
	Key                  *string   `json:"Key,omitempty"`
	Values               *[]string `json:"Values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchTagItemAllOf SearchTagItemAllOf

// NewSearchTagItemAllOf instantiates a new SearchTagItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTagItemAllOf() *SearchTagItemAllOf {
	this := SearchTagItemAllOf{}
	return &this
}

// NewSearchTagItemAllOfWithDefaults instantiates a new SearchTagItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTagItemAllOfWithDefaults() *SearchTagItemAllOf {
	this := SearchTagItemAllOf{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchTagItemAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTagItemAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchTagItemAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SearchTagItemAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SearchTagItemAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTagItemAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SearchTagItemAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SearchTagItemAllOf) SetKey(v string) {
	o.Key = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SearchTagItemAllOf) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTagItemAllOf) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SearchTagItemAllOf) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SearchTagItemAllOf) SetValues(v []string) {
	o.Values = &v
}

func (o SearchTagItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchTagItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSearchTagItemAllOf := _SearchTagItemAllOf{}

	if err = json.Unmarshal(bytes, &varSearchTagItemAllOf); err == nil {
		*o = SearchTagItemAllOf(varSearchTagItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchTagItemAllOf struct {
	value *SearchTagItemAllOf
	isSet bool
}

func (v NullableSearchTagItemAllOf) Get() *SearchTagItemAllOf {
	return v.value
}

func (v *NullableSearchTagItemAllOf) Set(val *SearchTagItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTagItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTagItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTagItemAllOf(val *SearchTagItemAllOf) *NullableSearchTagItemAllOf {
	return &NullableSearchTagItemAllOf{value: val, isSet: true}
}

func (v NullableSearchTagItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTagItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
