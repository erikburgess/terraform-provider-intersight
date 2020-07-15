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

// StoragePureReplicationBlackout Range of time at which to suspend replication. System disables replication during this interval.
type StoragePureReplicationBlackout struct {
	StorageBaseReplicationBlackout
	AdditionalProperties map[string]interface{}
}

type _StoragePureReplicationBlackout StoragePureReplicationBlackout

// NewStoragePureReplicationBlackout instantiates a new StoragePureReplicationBlackout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureReplicationBlackout() *StoragePureReplicationBlackout {
	this := StoragePureReplicationBlackout{}
	return &this
}

// NewStoragePureReplicationBlackoutWithDefaults instantiates a new StoragePureReplicationBlackout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureReplicationBlackoutWithDefaults() *StoragePureReplicationBlackout {
	this := StoragePureReplicationBlackout{}
	return &this
}

func (o StoragePureReplicationBlackout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseReplicationBlackout, errStorageBaseReplicationBlackout := json.Marshal(o.StorageBaseReplicationBlackout)
	if errStorageBaseReplicationBlackout != nil {
		return []byte{}, errStorageBaseReplicationBlackout
	}
	errStorageBaseReplicationBlackout = json.Unmarshal([]byte(serializedStorageBaseReplicationBlackout), &toSerialize)
	if errStorageBaseReplicationBlackout != nil {
		return []byte{}, errStorageBaseReplicationBlackout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureReplicationBlackout) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureReplicationBlackoutWithoutEmbeddedStruct struct {
	}

	varStoragePureReplicationBlackoutWithoutEmbeddedStruct := StoragePureReplicationBlackoutWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureReplicationBlackoutWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureReplicationBlackout := _StoragePureReplicationBlackout{}
		*o = StoragePureReplicationBlackout(varStoragePureReplicationBlackout)
	} else {
		return err
	}

	varStoragePureReplicationBlackout := _StoragePureReplicationBlackout{}

	err = json.Unmarshal(bytes, &varStoragePureReplicationBlackout)
	if err == nil {
		o.StorageBaseReplicationBlackout = varStoragePureReplicationBlackout.StorageBaseReplicationBlackout
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectStorageBaseReplicationBlackout := reflect.ValueOf(o.StorageBaseReplicationBlackout)
		for i := 0; i < reflectStorageBaseReplicationBlackout.Type().NumField(); i++ {
			t := reflectStorageBaseReplicationBlackout.Type().Field(i)

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

type NullableStoragePureReplicationBlackout struct {
	value *StoragePureReplicationBlackout
	isSet bool
}

func (v NullableStoragePureReplicationBlackout) Get() *StoragePureReplicationBlackout {
	return v.value
}

func (v *NullableStoragePureReplicationBlackout) Set(val *StoragePureReplicationBlackout) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureReplicationBlackout) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureReplicationBlackout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureReplicationBlackout(val *StoragePureReplicationBlackout) *NullableStoragePureReplicationBlackout {
	return &NullableStoragePureReplicationBlackout{value: val, isSet: true}
}

func (v NullableStoragePureReplicationBlackout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureReplicationBlackout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
