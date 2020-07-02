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

// VirtualizationVmwareDatastoreRelationship - A relationship to the 'virtualization.VmwareDatastore' resource, or the expanded 'virtualization.VmwareDatastore' resource, or the 'null' value.
type VirtualizationVmwareDatastoreRelationship struct {
	MoMoRef                       *MoMoRef
	VirtualizationVmwareDatastore *VirtualizationVmwareDatastore
}

// MoMoRefAsVirtualizationVmwareDatastoreRelationship is a convenience function that returns MoMoRef wrapped in VirtualizationVmwareDatastoreRelationship
func MoMoRefAsVirtualizationVmwareDatastoreRelationship(v *MoMoRef) VirtualizationVmwareDatastoreRelationship {
	return VirtualizationVmwareDatastoreRelationship{MoMoRef: v}
}

// VirtualizationVmwareDatastoreAsVirtualizationVmwareDatastoreRelationship is a convenience function that returns VirtualizationVmwareDatastore wrapped in VirtualizationVmwareDatastoreRelationship
func VirtualizationVmwareDatastoreAsVirtualizationVmwareDatastoreRelationship(v *VirtualizationVmwareDatastore) VirtualizationVmwareDatastoreRelationship {
	return VirtualizationVmwareDatastoreRelationship{VirtualizationVmwareDatastore: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VirtualizationVmwareDatastoreRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'virtualization.VmwareDatastore'
	if jsonDict["ClassId"] == "virtualization.VmwareDatastore" {
		// try to unmarshal JSON data into VirtualizationVmwareDatastore
		err = json.Unmarshal(data, &dst.VirtualizationVmwareDatastore)
		if err == nil {
			jsonVirtualizationVmwareDatastore, _ := json.Marshal(dst.VirtualizationVmwareDatastore)
			if string(jsonVirtualizationVmwareDatastore) == "{}" { // empty struct
				dst.VirtualizationVmwareDatastore = nil
			} else {
				return nil // data stored in dst.VirtualizationVmwareDatastore, return on the first match
			}
		} else {
			dst.VirtualizationVmwareDatastore = nil
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

	// try to unmarshal data into VirtualizationVmwareDatastore
	err = json.Unmarshal(data, &dst.VirtualizationVmwareDatastore)
	if err == nil {
		jsonVirtualizationVmwareDatastore, _ := json.Marshal(dst.VirtualizationVmwareDatastore)
		if string(jsonVirtualizationVmwareDatastore) == "{}" { // empty struct
			dst.VirtualizationVmwareDatastore = nil
		} else {
			match++
		}
	} else {
		dst.VirtualizationVmwareDatastore = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MoMoRef = nil
		dst.VirtualizationVmwareDatastore = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(VirtualizationVmwareDatastoreRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(VirtualizationVmwareDatastoreRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VirtualizationVmwareDatastoreRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VirtualizationVmwareDatastore != nil {
		return json.Marshal(&src.VirtualizationVmwareDatastore)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VirtualizationVmwareDatastoreRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VirtualizationVmwareDatastore != nil {
		return obj.VirtualizationVmwareDatastore
	}

	// all schemas are nil
	return nil
}

type NullableVirtualizationVmwareDatastoreRelationship struct {
	value *VirtualizationVmwareDatastoreRelationship
	isSet bool
}

func (v NullableVirtualizationVmwareDatastoreRelationship) Get() *VirtualizationVmwareDatastoreRelationship {
	return v.value
}

func (v *NullableVirtualizationVmwareDatastoreRelationship) Set(val *VirtualizationVmwareDatastoreRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatastoreRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatastoreRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatastoreRelationship(val *VirtualizationVmwareDatastoreRelationship) *NullableVirtualizationVmwareDatastoreRelationship {
	return &NullableVirtualizationVmwareDatastoreRelationship{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatastoreRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatastoreRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
