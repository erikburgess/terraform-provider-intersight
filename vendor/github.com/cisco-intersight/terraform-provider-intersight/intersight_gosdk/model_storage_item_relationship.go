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
	"fmt"
)

// StorageItemRelationship - A relationship to the 'storage.Item' resource, or the expanded 'storage.Item' resource, or the 'null' value.
type StorageItemRelationship struct {
	MoMoRef     *MoMoRef
	StorageItem *StorageItem
}

// MoMoRefAsStorageItemRelationship is a convenience function that returns MoMoRef wrapped in StorageItemRelationship
func MoMoRefAsStorageItemRelationship(v *MoMoRef) StorageItemRelationship {
	return StorageItemRelationship{MoMoRef: v}
}

// StorageItemAsStorageItemRelationship is a convenience function that returns StorageItem wrapped in StorageItemRelationship
func StorageItemAsStorageItemRelationship(v *StorageItem) StorageItemRelationship {
	return StorageItemRelationship{StorageItem: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageItemRelationship) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			jsonMoMoRef, _ := json.Marshal(dst.MoMoRef)
			if string(jsonMoMoRef) == "{}" { // empty struct
				dst.MoMoRef = nil
			} else {
				return nil // data stored in dst.MoMoRef, return on the first match
			}
		} else {
			dst.MoMoRef = nil
		}
	}

	// check if the discriminator value is 'storage.Item'
	if jsonDict["ClassId"] == "storage.Item" {
		// try to unmarshal JSON data into StorageItem
		err = json.Unmarshal(data, &dst.StorageItem)
		if err == nil {
			jsonStorageItem, _ := json.Marshal(dst.StorageItem)
			if string(jsonStorageItem) == "{}" { // empty struct
				dst.StorageItem = nil
			} else {
				return nil // data stored in dst.StorageItem, return on the first match
			}
		} else {
			dst.StorageItem = nil
		}
	}

	// try to unmarshal data into MoMoRef
	err = json.Unmarshal(data, &dst.MoMoRef)
	if err == nil {
		jsonMoMoRef, _ := json.Marshal(dst.MoMoRef)
		if string(jsonMoMoRef) == "{}" { // empty struct
			dst.MoMoRef = nil
		} else {
			match++
		}
	} else {
		dst.MoMoRef = nil
	}

	// try to unmarshal data into StorageItem
	err = json.Unmarshal(data, &dst.StorageItem)
	if err == nil {
		jsonStorageItem, _ := json.Marshal(dst.StorageItem)
		if string(jsonStorageItem) == "{}" { // empty struct
			dst.StorageItem = nil
		} else {
			match++
		}
	} else {
		dst.StorageItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MoMoRef = nil
		dst.StorageItem = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(StorageItemRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(StorageItemRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageItemRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageItem != nil {
		return json.Marshal(&src.StorageItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageItemRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageItem != nil {
		return obj.StorageItem
	}

	// all schemas are nil
	return nil
}

type NullableStorageItemRelationship struct {
	value *StorageItemRelationship
	isSet bool
}

func (v NullableStorageItemRelationship) Get() *StorageItemRelationship {
	return v.value
}

func (v *NullableStorageItemRelationship) Set(val *StorageItemRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageItemRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageItemRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageItemRelationship(val *StorageItemRelationship) *NullableStorageItemRelationship {
	return &NullableStorageItemRelationship{value: val, isSet: true}
}

func (v NullableStorageItemRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageItemRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
