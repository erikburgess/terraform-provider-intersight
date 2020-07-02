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

// EquipmentFanModuleRelationship - A relationship to the 'equipment.FanModule' resource, or the expanded 'equipment.FanModule' resource, or the 'null' value.
type EquipmentFanModuleRelationship struct {
	EquipmentFanModule *EquipmentFanModule
	MoMoRef            *MoMoRef
}

// EquipmentFanModuleAsEquipmentFanModuleRelationship is a convenience function that returns EquipmentFanModule wrapped in EquipmentFanModuleRelationship
func EquipmentFanModuleAsEquipmentFanModuleRelationship(v *EquipmentFanModule) EquipmentFanModuleRelationship {
	return EquipmentFanModuleRelationship{EquipmentFanModule: v}
}

// MoMoRefAsEquipmentFanModuleRelationship is a convenience function that returns MoMoRef wrapped in EquipmentFanModuleRelationship
func MoMoRefAsEquipmentFanModuleRelationship(v *MoMoRef) EquipmentFanModuleRelationship {
	return EquipmentFanModuleRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentFanModuleRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'equipment.FanModule'
	if jsonDict["ClassId"] == "equipment.FanModule" {
		// try to unmarshal JSON data into EquipmentFanModule
		err = json.Unmarshal(data, &dst.EquipmentFanModule)
		if err == nil {
			jsonEquipmentFanModule, _ := json.Marshal(dst.EquipmentFanModule)
			if string(jsonEquipmentFanModule) == "{}" { // empty struct
				dst.EquipmentFanModule = nil
			} else {
				return nil // data stored in dst.EquipmentFanModule, return on the first match
			}
		} else {
			dst.EquipmentFanModule = nil
		}
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

	// try to unmarshal data into EquipmentFanModule
	err = json.Unmarshal(data, &dst.EquipmentFanModule)
	if err == nil {
		jsonEquipmentFanModule, _ := json.Marshal(dst.EquipmentFanModule)
		if string(jsonEquipmentFanModule) == "{}" { // empty struct
			dst.EquipmentFanModule = nil
		} else {
			match++
		}
	} else {
		dst.EquipmentFanModule = nil
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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EquipmentFanModule = nil
		dst.MoMoRef = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(EquipmentFanModuleRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(EquipmentFanModuleRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentFanModuleRelationship) MarshalJSON() ([]byte, error) {
	if src.EquipmentFanModule != nil {
		return json.Marshal(&src.EquipmentFanModule)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EquipmentFanModuleRelationship) GetActualInstance() interface{} {
	if obj.EquipmentFanModule != nil {
		return obj.EquipmentFanModule
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableEquipmentFanModuleRelationship struct {
	value *EquipmentFanModuleRelationship
	isSet bool
}

func (v NullableEquipmentFanModuleRelationship) Get() *EquipmentFanModuleRelationship {
	return v.value
}

func (v *NullableEquipmentFanModuleRelationship) Set(val *EquipmentFanModuleRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFanModuleRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFanModuleRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFanModuleRelationship(val *EquipmentFanModuleRelationship) *NullableEquipmentFanModuleRelationship {
	return &NullableEquipmentFanModuleRelationship{value: val, isSet: true}
}

func (v NullableEquipmentFanModuleRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFanModuleRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
