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

// FirmwareBiosDescriptorResponse - The response body of a HTTP GET request for the 'firmware.BiosDescriptor' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'firmware.BiosDescriptor' resources.
type FirmwareBiosDescriptorResponse struct {
	FirmwareBiosDescriptorList *FirmwareBiosDescriptorList
	MoAggregateTransform       *MoAggregateTransform
	MoDocumentCount            *MoDocumentCount
	MoTagSummary               *MoTagSummary
}

// FirmwareBiosDescriptorListAsFirmwareBiosDescriptorResponse is a convenience function that returns FirmwareBiosDescriptorList wrapped in FirmwareBiosDescriptorResponse
func FirmwareBiosDescriptorListAsFirmwareBiosDescriptorResponse(v *FirmwareBiosDescriptorList) FirmwareBiosDescriptorResponse {
	return FirmwareBiosDescriptorResponse{FirmwareBiosDescriptorList: v}
}

// MoAggregateTransformAsFirmwareBiosDescriptorResponse is a convenience function that returns MoAggregateTransform wrapped in FirmwareBiosDescriptorResponse
func MoAggregateTransformAsFirmwareBiosDescriptorResponse(v *MoAggregateTransform) FirmwareBiosDescriptorResponse {
	return FirmwareBiosDescriptorResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsFirmwareBiosDescriptorResponse is a convenience function that returns MoDocumentCount wrapped in FirmwareBiosDescriptorResponse
func MoDocumentCountAsFirmwareBiosDescriptorResponse(v *MoDocumentCount) FirmwareBiosDescriptorResponse {
	return FirmwareBiosDescriptorResponse{MoDocumentCount: v}
}

// MoTagSummaryAsFirmwareBiosDescriptorResponse is a convenience function that returns MoTagSummary wrapped in FirmwareBiosDescriptorResponse
func MoTagSummaryAsFirmwareBiosDescriptorResponse(v *MoTagSummary) FirmwareBiosDescriptorResponse {
	return FirmwareBiosDescriptorResponse{MoTagSummary: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FirmwareBiosDescriptorResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'firmware.BiosDescriptor.List'
	if jsonDict["ObjectType"] == "firmware.BiosDescriptor.List" {
		// try to unmarshal JSON data into FirmwareBiosDescriptorList
		err = json.Unmarshal(data, &dst.FirmwareBiosDescriptorList)
		if err == nil {
			return nil // data stored in dst.FirmwareBiosDescriptorList, return on the first match
		} else {
			dst.FirmwareBiosDescriptorList = nil
			return fmt.Errorf("Failed to unmarshal FirmwareBiosDescriptorResponse as FirmwareBiosDescriptorList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal FirmwareBiosDescriptorResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal FirmwareBiosDescriptorResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal FirmwareBiosDescriptorResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FirmwareBiosDescriptorResponse) MarshalJSON() ([]byte, error) {
	if src.FirmwareBiosDescriptorList != nil {
		return json.Marshal(&src.FirmwareBiosDescriptorList)
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
func (obj *FirmwareBiosDescriptorResponse) GetActualInstance() interface{} {
	if obj.FirmwareBiosDescriptorList != nil {
		return obj.FirmwareBiosDescriptorList
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

type NullableFirmwareBiosDescriptorResponse struct {
	value *FirmwareBiosDescriptorResponse
	isSet bool
}

func (v NullableFirmwareBiosDescriptorResponse) Get() *FirmwareBiosDescriptorResponse {
	return v.value
}

func (v *NullableFirmwareBiosDescriptorResponse) Set(val *FirmwareBiosDescriptorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareBiosDescriptorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareBiosDescriptorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareBiosDescriptorResponse(val *FirmwareBiosDescriptorResponse) *NullableFirmwareBiosDescriptorResponse {
	return &NullableFirmwareBiosDescriptorResponse{value: val, isSet: true}
}

func (v NullableFirmwareBiosDescriptorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareBiosDescriptorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
