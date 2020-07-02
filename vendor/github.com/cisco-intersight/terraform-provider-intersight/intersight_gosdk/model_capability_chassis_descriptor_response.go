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

// CapabilityChassisDescriptorResponse - The response body of a HTTP GET request for the 'capability.ChassisDescriptor' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'capability.ChassisDescriptor' resources.
type CapabilityChassisDescriptorResponse struct {
	CapabilityChassisDescriptorList *CapabilityChassisDescriptorList
	MoAggregateTransform            *MoAggregateTransform
	MoDocumentCount                 *MoDocumentCount
	MoTagSummary                    *MoTagSummary
}

// CapabilityChassisDescriptorListAsCapabilityChassisDescriptorResponse is a convenience function that returns CapabilityChassisDescriptorList wrapped in CapabilityChassisDescriptorResponse
func CapabilityChassisDescriptorListAsCapabilityChassisDescriptorResponse(v *CapabilityChassisDescriptorList) CapabilityChassisDescriptorResponse {
	return CapabilityChassisDescriptorResponse{CapabilityChassisDescriptorList: v}
}

// MoAggregateTransformAsCapabilityChassisDescriptorResponse is a convenience function that returns MoAggregateTransform wrapped in CapabilityChassisDescriptorResponse
func MoAggregateTransformAsCapabilityChassisDescriptorResponse(v *MoAggregateTransform) CapabilityChassisDescriptorResponse {
	return CapabilityChassisDescriptorResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsCapabilityChassisDescriptorResponse is a convenience function that returns MoDocumentCount wrapped in CapabilityChassisDescriptorResponse
func MoDocumentCountAsCapabilityChassisDescriptorResponse(v *MoDocumentCount) CapabilityChassisDescriptorResponse {
	return CapabilityChassisDescriptorResponse{MoDocumentCount: v}
}

// MoTagSummaryAsCapabilityChassisDescriptorResponse is a convenience function that returns MoTagSummary wrapped in CapabilityChassisDescriptorResponse
func MoTagSummaryAsCapabilityChassisDescriptorResponse(v *MoTagSummary) CapabilityChassisDescriptorResponse {
	return CapabilityChassisDescriptorResponse{MoTagSummary: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CapabilityChassisDescriptorResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'capability.ChassisDescriptor.List'
	if jsonDict["ObjectType"] == "capability.ChassisDescriptor.List" {
		// try to unmarshal JSON data into CapabilityChassisDescriptorList
		err = json.Unmarshal(data, &dst.CapabilityChassisDescriptorList)
		if err == nil {
			jsonCapabilityChassisDescriptorList, _ := json.Marshal(dst.CapabilityChassisDescriptorList)
			if string(jsonCapabilityChassisDescriptorList) == "{}" { // empty struct
				dst.CapabilityChassisDescriptorList = nil
			} else {
				return nil // data stored in dst.CapabilityChassisDescriptorList, return on the first match
			}
		} else {
			dst.CapabilityChassisDescriptorList = nil
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			jsonMoAggregateTransform, _ := json.Marshal(dst.MoAggregateTransform)
			if string(jsonMoAggregateTransform) == "{}" { // empty struct
				dst.MoAggregateTransform = nil
			} else {
				return nil // data stored in dst.MoAggregateTransform, return on the first match
			}
		} else {
			dst.MoAggregateTransform = nil
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			jsonMoDocumentCount, _ := json.Marshal(dst.MoDocumentCount)
			if string(jsonMoDocumentCount) == "{}" { // empty struct
				dst.MoDocumentCount = nil
			} else {
				return nil // data stored in dst.MoDocumentCount, return on the first match
			}
		} else {
			dst.MoDocumentCount = nil
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			jsonMoTagSummary, _ := json.Marshal(dst.MoTagSummary)
			if string(jsonMoTagSummary) == "{}" { // empty struct
				dst.MoTagSummary = nil
			} else {
				return nil // data stored in dst.MoTagSummary, return on the first match
			}
		} else {
			dst.MoTagSummary = nil
		}
	}

	// try to unmarshal data into CapabilityChassisDescriptorList
	err = json.Unmarshal(data, &dst.CapabilityChassisDescriptorList)
	if err == nil {
		jsonCapabilityChassisDescriptorList, _ := json.Marshal(dst.CapabilityChassisDescriptorList)
		if string(jsonCapabilityChassisDescriptorList) == "{}" { // empty struct
			dst.CapabilityChassisDescriptorList = nil
		} else {
			match++
		}
	} else {
		dst.CapabilityChassisDescriptorList = nil
	}

	// try to unmarshal data into MoAggregateTransform
	err = json.Unmarshal(data, &dst.MoAggregateTransform)
	if err == nil {
		jsonMoAggregateTransform, _ := json.Marshal(dst.MoAggregateTransform)
		if string(jsonMoAggregateTransform) == "{}" { // empty struct
			dst.MoAggregateTransform = nil
		} else {
			match++
		}
	} else {
		dst.MoAggregateTransform = nil
	}

	// try to unmarshal data into MoDocumentCount
	err = json.Unmarshal(data, &dst.MoDocumentCount)
	if err == nil {
		jsonMoDocumentCount, _ := json.Marshal(dst.MoDocumentCount)
		if string(jsonMoDocumentCount) == "{}" { // empty struct
			dst.MoDocumentCount = nil
		} else {
			match++
		}
	} else {
		dst.MoDocumentCount = nil
	}

	// try to unmarshal data into MoTagSummary
	err = json.Unmarshal(data, &dst.MoTagSummary)
	if err == nil {
		jsonMoTagSummary, _ := json.Marshal(dst.MoTagSummary)
		if string(jsonMoTagSummary) == "{}" { // empty struct
			dst.MoTagSummary = nil
		} else {
			match++
		}
	} else {
		dst.MoTagSummary = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CapabilityChassisDescriptorList = nil
		dst.MoAggregateTransform = nil
		dst.MoDocumentCount = nil
		dst.MoTagSummary = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CapabilityChassisDescriptorResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CapabilityChassisDescriptorResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CapabilityChassisDescriptorResponse) MarshalJSON() ([]byte, error) {
	if src.CapabilityChassisDescriptorList != nil {
		return json.Marshal(&src.CapabilityChassisDescriptorList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CapabilityChassisDescriptorResponse) GetActualInstance() interface{} {
	if obj.CapabilityChassisDescriptorList != nil {
		return obj.CapabilityChassisDescriptorList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableCapabilityChassisDescriptorResponse struct {
	value *CapabilityChassisDescriptorResponse
	isSet bool
}

func (v NullableCapabilityChassisDescriptorResponse) Get() *CapabilityChassisDescriptorResponse {
	return v.value
}

func (v *NullableCapabilityChassisDescriptorResponse) Set(val *CapabilityChassisDescriptorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityChassisDescriptorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityChassisDescriptorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityChassisDescriptorResponse(val *CapabilityChassisDescriptorResponse) *NullableCapabilityChassisDescriptorResponse {
	return &NullableCapabilityChassisDescriptorResponse{value: val, isSet: true}
}

func (v NullableCapabilityChassisDescriptorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityChassisDescriptorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}