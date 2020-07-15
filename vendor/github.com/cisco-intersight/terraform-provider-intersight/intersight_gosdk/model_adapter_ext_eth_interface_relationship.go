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

// AdapterExtEthInterfaceRelationship - A relationship to the 'adapter.ExtEthInterface' resource, or the expanded 'adapter.ExtEthInterface' resource, or the 'null' value.
type AdapterExtEthInterfaceRelationship struct {
	AdapterExtEthInterface *AdapterExtEthInterface
	MoMoRef                *MoMoRef
}

// AdapterExtEthInterfaceAsAdapterExtEthInterfaceRelationship is a convenience function that returns AdapterExtEthInterface wrapped in AdapterExtEthInterfaceRelationship
func AdapterExtEthInterfaceAsAdapterExtEthInterfaceRelationship(v *AdapterExtEthInterface) AdapterExtEthInterfaceRelationship {
	return AdapterExtEthInterfaceRelationship{AdapterExtEthInterface: v}
}

// MoMoRefAsAdapterExtEthInterfaceRelationship is a convenience function that returns MoMoRef wrapped in AdapterExtEthInterfaceRelationship
func MoMoRefAsAdapterExtEthInterfaceRelationship(v *MoMoRef) AdapterExtEthInterfaceRelationship {
	return AdapterExtEthInterfaceRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AdapterExtEthInterfaceRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'adapter.ExtEthInterface'
	if jsonDict["ClassId"] == "adapter.ExtEthInterface" {
		// try to unmarshal JSON data into AdapterExtEthInterface
		err = json.Unmarshal(data, &dst.AdapterExtEthInterface)
		if err == nil {
			return nil // data stored in dst.AdapterExtEthInterface, return on the first match
		} else {
			dst.AdapterExtEthInterface = nil
			return fmt.Errorf("Failed to unmarshal AdapterExtEthInterfaceRelationship as AdapterExtEthInterface: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal AdapterExtEthInterfaceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AdapterExtEthInterfaceRelationship) MarshalJSON() ([]byte, error) {
	if src.AdapterExtEthInterface != nil {
		return json.Marshal(&src.AdapterExtEthInterface)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AdapterExtEthInterfaceRelationship) GetActualInstance() interface{} {
	if obj.AdapterExtEthInterface != nil {
		return obj.AdapterExtEthInterface
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableAdapterExtEthInterfaceRelationship struct {
	value *AdapterExtEthInterfaceRelationship
	isSet bool
}

func (v NullableAdapterExtEthInterfaceRelationship) Get() *AdapterExtEthInterfaceRelationship {
	return v.value
}

func (v *NullableAdapterExtEthInterfaceRelationship) Set(val *AdapterExtEthInterfaceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterExtEthInterfaceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterExtEthInterfaceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterExtEthInterfaceRelationship(val *AdapterExtEthInterfaceRelationship) *NullableAdapterExtEthInterfaceRelationship {
	return &NullableAdapterExtEthInterfaceRelationship{value: val, isSet: true}
}

func (v NullableAdapterExtEthInterfaceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterExtEthInterfaceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
