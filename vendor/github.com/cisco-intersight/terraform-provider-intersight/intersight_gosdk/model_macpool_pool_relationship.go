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
	"fmt"
)

// MacpoolPoolRelationship - A relationship to the 'macpool.Pool' resource, or the expanded 'macpool.Pool' resource, or the 'null' value.
type MacpoolPoolRelationship struct {
	MacpoolPool *MacpoolPool
	MoMoRef     *MoMoRef
}

// MacpoolPoolAsMacpoolPoolRelationship is a convenience function that returns MacpoolPool wrapped in MacpoolPoolRelationship
func MacpoolPoolAsMacpoolPoolRelationship(v *MacpoolPool) MacpoolPoolRelationship {
	return MacpoolPoolRelationship{MacpoolPool: v}
}

// MoMoRefAsMacpoolPoolRelationship is a convenience function that returns MoMoRef wrapped in MacpoolPoolRelationship
func MoMoRefAsMacpoolPoolRelationship(v *MoMoRef) MacpoolPoolRelationship {
	return MacpoolPoolRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MacpoolPoolRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'macpool.Pool'
	if jsonDict["ClassId"] == "macpool.Pool" {
		// try to unmarshal JSON data into MacpoolPool
		err = json.Unmarshal(data, &dst.MacpoolPool)
		if err == nil {
			return nil // data stored in dst.MacpoolPool, return on the first match
		} else {
			dst.MacpoolPool = nil
			return fmt.Errorf("Failed to unmarshal MacpoolPoolRelationship as MacpoolPool: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal MacpoolPoolRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MacpoolPoolRelationship) MarshalJSON() ([]byte, error) {
	if src.MacpoolPool != nil {
		return json.Marshal(&src.MacpoolPool)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MacpoolPoolRelationship) GetActualInstance() interface{} {
	if obj.MacpoolPool != nil {
		return obj.MacpoolPool
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableMacpoolPoolRelationship struct {
	value *MacpoolPoolRelationship
	isSet bool
}

func (v NullableMacpoolPoolRelationship) Get() *MacpoolPoolRelationship {
	return v.value
}

func (v *NullableMacpoolPoolRelationship) Set(val *MacpoolPoolRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolPoolRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolPoolRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolPoolRelationship(val *MacpoolPoolRelationship) *NullableMacpoolPoolRelationship {
	return &NullableMacpoolPoolRelationship{value: val, isSet: true}
}

func (v NullableMacpoolPoolRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolPoolRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
