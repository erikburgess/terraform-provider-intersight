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
	"fmt"
)

// ConfigExportedItemRelationship - A relationship to the 'config.ExportedItem' resource, or the expanded 'config.ExportedItem' resource, or the 'null' value.
type ConfigExportedItemRelationship struct {
	ConfigExportedItem *ConfigExportedItem
	MoMoRef            *MoMoRef
}

// ConfigExportedItemAsConfigExportedItemRelationship is a convenience function that returns ConfigExportedItem wrapped in ConfigExportedItemRelationship
func ConfigExportedItemAsConfigExportedItemRelationship(v *ConfigExportedItem) ConfigExportedItemRelationship {
	return ConfigExportedItemRelationship{ConfigExportedItem: v}
}

// MoMoRefAsConfigExportedItemRelationship is a convenience function that returns MoMoRef wrapped in ConfigExportedItemRelationship
func MoMoRefAsConfigExportedItemRelationship(v *MoMoRef) ConfigExportedItemRelationship {
	return ConfigExportedItemRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConfigExportedItemRelationship) UnmarshalJSON(data []byte) error {
	var err error
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

	// check if the discriminator value is 'config.ExportedItem'
	if jsonDict["ClassId"] == "config.ExportedItem" {
		// try to unmarshal JSON data into ConfigExportedItem
		err = json.Unmarshal(data, &dst.ConfigExportedItem)
		if err == nil {
			return nil // data stored in dst.ConfigExportedItem, return on the first match
		} else {
			dst.ConfigExportedItem = nil
			return fmt.Errorf("Failed to unmarshal ConfigExportedItemRelationship as ConfigExportedItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal ConfigExportedItemRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConfigExportedItemRelationship) MarshalJSON() ([]byte, error) {
	if src.ConfigExportedItem != nil {
		return json.Marshal(&src.ConfigExportedItem)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConfigExportedItemRelationship) GetActualInstance() interface{} {
	if obj.ConfigExportedItem != nil {
		return obj.ConfigExportedItem
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableConfigExportedItemRelationship struct {
	value *ConfigExportedItemRelationship
	isSet bool
}

func (v NullableConfigExportedItemRelationship) Get() *ConfigExportedItemRelationship {
	return v.value
}

func (v *NullableConfigExportedItemRelationship) Set(val *ConfigExportedItemRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigExportedItemRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigExportedItemRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigExportedItemRelationship(val *ConfigExportedItemRelationship) *NullableConfigExportedItemRelationship {
	return &NullableConfigExportedItemRelationship{value: val, isSet: true}
}

func (v NullableConfigExportedItemRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigExportedItemRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
