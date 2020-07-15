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
)

// StorageBaseArrayAllOf Definition of the list of properties defined in 'storage.BaseArray', excluding properties defined in parent classes.
type StorageBaseArrayAllOf struct {
	StorageUtilization   *StorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseArrayAllOf StorageBaseArrayAllOf

// NewStorageBaseArrayAllOf instantiates a new StorageBaseArrayAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseArrayAllOf() *StorageBaseArrayAllOf {
	this := StorageBaseArrayAllOf{}
	return &this
}

// NewStorageBaseArrayAllOfWithDefaults instantiates a new StorageBaseArrayAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseArrayAllOfWithDefaults() *StorageBaseArrayAllOf {
	this := StorageBaseArrayAllOf{}
	return &this
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StorageBaseArrayAllOf) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseArrayAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseArrayAllOf) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization = &v
}

func (o StorageBaseArrayAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseArrayAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseArrayAllOf := _StorageBaseArrayAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseArrayAllOf); err == nil {
		*o = StorageBaseArrayAllOf(varStorageBaseArrayAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "StorageUtilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseArrayAllOf struct {
	value *StorageBaseArrayAllOf
	isSet bool
}

func (v NullableStorageBaseArrayAllOf) Get() *StorageBaseArrayAllOf {
	return v.value
}

func (v *NullableStorageBaseArrayAllOf) Set(val *StorageBaseArrayAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseArrayAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseArrayAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseArrayAllOf(val *StorageBaseArrayAllOf) *NullableStorageBaseArrayAllOf {
	return &NullableStorageBaseArrayAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseArrayAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseArrayAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
