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

// FcpoolPoolResponse - The response body of a HTTP GET request for the 'fcpool.Pool' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'fcpool.Pool' resources.
type FcpoolPoolResponse struct {
	FcpoolPoolList       *FcpoolPoolList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// FcpoolPoolListAsFcpoolPoolResponse is a convenience function that returns FcpoolPoolList wrapped in FcpoolPoolResponse
func FcpoolPoolListAsFcpoolPoolResponse(v *FcpoolPoolList) FcpoolPoolResponse {
	return FcpoolPoolResponse{FcpoolPoolList: v}
}

// MoAggregateTransformAsFcpoolPoolResponse is a convenience function that returns MoAggregateTransform wrapped in FcpoolPoolResponse
func MoAggregateTransformAsFcpoolPoolResponse(v *MoAggregateTransform) FcpoolPoolResponse {
	return FcpoolPoolResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsFcpoolPoolResponse is a convenience function that returns MoDocumentCount wrapped in FcpoolPoolResponse
func MoDocumentCountAsFcpoolPoolResponse(v *MoDocumentCount) FcpoolPoolResponse {
	return FcpoolPoolResponse{MoDocumentCount: v}
}

// MoTagSummaryAsFcpoolPoolResponse is a convenience function that returns MoTagSummary wrapped in FcpoolPoolResponse
func MoTagSummaryAsFcpoolPoolResponse(v *MoTagSummary) FcpoolPoolResponse {
	return FcpoolPoolResponse{MoTagSummary: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FcpoolPoolResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'fcpool.Pool.List'
	if jsonDict["ObjectType"] == "fcpool.Pool.List" {
		// try to unmarshal JSON data into FcpoolPoolList
		err = json.Unmarshal(data, &dst.FcpoolPoolList)
		if err == nil {
			jsonFcpoolPoolList, _ := json.Marshal(dst.FcpoolPoolList)
			if string(jsonFcpoolPoolList) == "{}" { // empty struct
				dst.FcpoolPoolList = nil
			} else {
				return nil // data stored in dst.FcpoolPoolList, return on the first match
			}
		} else {
			dst.FcpoolPoolList = nil
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

	// try to unmarshal data into FcpoolPoolList
	err = json.Unmarshal(data, &dst.FcpoolPoolList)
	if err == nil {
		jsonFcpoolPoolList, _ := json.Marshal(dst.FcpoolPoolList)
		if string(jsonFcpoolPoolList) == "{}" { // empty struct
			dst.FcpoolPoolList = nil
		} else {
			match++
		}
	} else {
		dst.FcpoolPoolList = nil
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
		dst.FcpoolPoolList = nil
		dst.MoAggregateTransform = nil
		dst.MoDocumentCount = nil
		dst.MoTagSummary = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(FcpoolPoolResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(FcpoolPoolResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FcpoolPoolResponse) MarshalJSON() ([]byte, error) {
	if src.FcpoolPoolList != nil {
		return json.Marshal(&src.FcpoolPoolList)
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
func (obj *FcpoolPoolResponse) GetActualInstance() interface{} {
	if obj.FcpoolPoolList != nil {
		return obj.FcpoolPoolList
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

type NullableFcpoolPoolResponse struct {
	value *FcpoolPoolResponse
	isSet bool
}

func (v NullableFcpoolPoolResponse) Get() *FcpoolPoolResponse {
	return v.value
}

func (v *NullableFcpoolPoolResponse) Set(val *FcpoolPoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolPoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolPoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolPoolResponse(val *FcpoolPoolResponse) *NullableFcpoolPoolResponse {
	return &NullableFcpoolPoolResponse{value: val, isSet: true}
}

func (v NullableFcpoolPoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolPoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
