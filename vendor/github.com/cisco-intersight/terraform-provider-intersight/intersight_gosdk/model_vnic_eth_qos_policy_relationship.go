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

// VnicEthQosPolicyRelationship - A relationship to the 'vnic.EthQosPolicy' resource, or the expanded 'vnic.EthQosPolicy' resource, or the 'null' value.
type VnicEthQosPolicyRelationship struct {
	MoMoRef          *MoMoRef
	VnicEthQosPolicy *VnicEthQosPolicy
}

// MoMoRefAsVnicEthQosPolicyRelationship is a convenience function that returns MoMoRef wrapped in VnicEthQosPolicyRelationship
func MoMoRefAsVnicEthQosPolicyRelationship(v *MoMoRef) VnicEthQosPolicyRelationship {
	return VnicEthQosPolicyRelationship{MoMoRef: v}
}

// VnicEthQosPolicyAsVnicEthQosPolicyRelationship is a convenience function that returns VnicEthQosPolicy wrapped in VnicEthQosPolicyRelationship
func VnicEthQosPolicyAsVnicEthQosPolicyRelationship(v *VnicEthQosPolicy) VnicEthQosPolicyRelationship {
	return VnicEthQosPolicyRelationship{VnicEthQosPolicy: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VnicEthQosPolicyRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'vnic.EthQosPolicy'
	if jsonDict["ClassId"] == "vnic.EthQosPolicy" {
		// try to unmarshal JSON data into VnicEthQosPolicy
		err = json.Unmarshal(data, &dst.VnicEthQosPolicy)
		if err == nil {
			jsonVnicEthQosPolicy, _ := json.Marshal(dst.VnicEthQosPolicy)
			if string(jsonVnicEthQosPolicy) == "{}" { // empty struct
				dst.VnicEthQosPolicy = nil
			} else {
				return nil // data stored in dst.VnicEthQosPolicy, return on the first match
			}
		} else {
			dst.VnicEthQosPolicy = nil
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

	// try to unmarshal data into VnicEthQosPolicy
	err = json.Unmarshal(data, &dst.VnicEthQosPolicy)
	if err == nil {
		jsonVnicEthQosPolicy, _ := json.Marshal(dst.VnicEthQosPolicy)
		if string(jsonVnicEthQosPolicy) == "{}" { // empty struct
			dst.VnicEthQosPolicy = nil
		} else {
			match++
		}
	} else {
		dst.VnicEthQosPolicy = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MoMoRef = nil
		dst.VnicEthQosPolicy = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(VnicEthQosPolicyRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(VnicEthQosPolicyRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VnicEthQosPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VnicEthQosPolicy != nil {
		return json.Marshal(&src.VnicEthQosPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VnicEthQosPolicyRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VnicEthQosPolicy != nil {
		return obj.VnicEthQosPolicy
	}

	// all schemas are nil
	return nil
}

type NullableVnicEthQosPolicyRelationship struct {
	value *VnicEthQosPolicyRelationship
	isSet bool
}

func (v NullableVnicEthQosPolicyRelationship) Get() *VnicEthQosPolicyRelationship {
	return v.value
}

func (v *NullableVnicEthQosPolicyRelationship) Set(val *VnicEthQosPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthQosPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthQosPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthQosPolicyRelationship(val *VnicEthQosPolicyRelationship) *NullableVnicEthQosPolicyRelationship {
	return &NullableVnicEthQosPolicyRelationship{value: val, isSet: true}
}

func (v NullableVnicEthQosPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthQosPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
