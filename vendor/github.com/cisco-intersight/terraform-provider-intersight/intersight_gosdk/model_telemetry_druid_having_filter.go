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

// TelemetryDruidHavingFilter - struct for TelemetryDruidHavingFilter
type TelemetryDruidHavingFilter struct {
	TelemetryDruidHavingDimensionSelectorFilter *TelemetryDruidHavingDimensionSelectorFilter
	TelemetryDruidHavingNumericFilter           *TelemetryDruidHavingNumericFilter
	TelemetryDruidHavingQueryFilter             *TelemetryDruidHavingQueryFilter
}

// TelemetryDruidHavingDimensionSelectorFilterAsTelemetryDruidHavingFilter is a convenience function that returns TelemetryDruidHavingDimensionSelectorFilter wrapped in TelemetryDruidHavingFilter
func TelemetryDruidHavingDimensionSelectorFilterAsTelemetryDruidHavingFilter(v *TelemetryDruidHavingDimensionSelectorFilter) TelemetryDruidHavingFilter {
	return TelemetryDruidHavingFilter{TelemetryDruidHavingDimensionSelectorFilter: v}
}

// TelemetryDruidHavingNumericFilterAsTelemetryDruidHavingFilter is a convenience function that returns TelemetryDruidHavingNumericFilter wrapped in TelemetryDruidHavingFilter
func TelemetryDruidHavingNumericFilterAsTelemetryDruidHavingFilter(v *TelemetryDruidHavingNumericFilter) TelemetryDruidHavingFilter {
	return TelemetryDruidHavingFilter{TelemetryDruidHavingNumericFilter: v}
}

// TelemetryDruidHavingQueryFilterAsTelemetryDruidHavingFilter is a convenience function that returns TelemetryDruidHavingQueryFilter wrapped in TelemetryDruidHavingFilter
func TelemetryDruidHavingQueryFilterAsTelemetryDruidHavingFilter(v *TelemetryDruidHavingQueryFilter) TelemetryDruidHavingFilter {
	return TelemetryDruidHavingFilter{TelemetryDruidHavingQueryFilter: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidHavingFilter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'dimSelector'
	if jsonDict["type"] == "dimSelector" {
		// try to unmarshal JSON data into TelemetryDruidHavingDimensionSelectorFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingDimensionSelectorFilter)
		if err == nil {
			jsonTelemetryDruidHavingDimensionSelectorFilter, _ := json.Marshal(dst.TelemetryDruidHavingDimensionSelectorFilter)
			if string(jsonTelemetryDruidHavingDimensionSelectorFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingDimensionSelectorFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingDimensionSelectorFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingDimensionSelectorFilter = nil
		}
	}

	// check if the discriminator value is 'equalTo'
	if jsonDict["type"] == "equalTo" {
		// try to unmarshal JSON data into TelemetryDruidHavingNumericFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingNumericFilter)
		if err == nil {
			jsonTelemetryDruidHavingNumericFilter, _ := json.Marshal(dst.TelemetryDruidHavingNumericFilter)
			if string(jsonTelemetryDruidHavingNumericFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingNumericFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingNumericFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingNumericFilter = nil
		}
	}

	// check if the discriminator value is 'greaterThan'
	if jsonDict["type"] == "greaterThan" {
		// try to unmarshal JSON data into TelemetryDruidHavingNumericFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingNumericFilter)
		if err == nil {
			jsonTelemetryDruidHavingNumericFilter, _ := json.Marshal(dst.TelemetryDruidHavingNumericFilter)
			if string(jsonTelemetryDruidHavingNumericFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingNumericFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingNumericFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingNumericFilter = nil
		}
	}

	// check if the discriminator value is 'having'
	if jsonDict["type"] == "having" {
		// try to unmarshal JSON data into TelemetryDruidHavingQueryFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingQueryFilter)
		if err == nil {
			jsonTelemetryDruidHavingQueryFilter, _ := json.Marshal(dst.TelemetryDruidHavingQueryFilter)
			if string(jsonTelemetryDruidHavingQueryFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingQueryFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingQueryFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingQueryFilter = nil
		}
	}

	// check if the discriminator value is 'lessThan'
	if jsonDict["type"] == "lessThan" {
		// try to unmarshal JSON data into TelemetryDruidHavingNumericFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingNumericFilter)
		if err == nil {
			jsonTelemetryDruidHavingNumericFilter, _ := json.Marshal(dst.TelemetryDruidHavingNumericFilter)
			if string(jsonTelemetryDruidHavingNumericFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingNumericFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingNumericFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingNumericFilter = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidHavingDimensionSelectorFilter'
	if jsonDict["type"] == "telemetry.DruidHavingDimensionSelectorFilter" {
		// try to unmarshal JSON data into TelemetryDruidHavingDimensionSelectorFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingDimensionSelectorFilter)
		if err == nil {
			jsonTelemetryDruidHavingDimensionSelectorFilter, _ := json.Marshal(dst.TelemetryDruidHavingDimensionSelectorFilter)
			if string(jsonTelemetryDruidHavingDimensionSelectorFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingDimensionSelectorFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingDimensionSelectorFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingDimensionSelectorFilter = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidHavingNumericFilter'
	if jsonDict["type"] == "telemetry.DruidHavingNumericFilter" {
		// try to unmarshal JSON data into TelemetryDruidHavingNumericFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingNumericFilter)
		if err == nil {
			jsonTelemetryDruidHavingNumericFilter, _ := json.Marshal(dst.TelemetryDruidHavingNumericFilter)
			if string(jsonTelemetryDruidHavingNumericFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingNumericFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingNumericFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingNumericFilter = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidHavingQueryFilter'
	if jsonDict["type"] == "telemetry.DruidHavingQueryFilter" {
		// try to unmarshal JSON data into TelemetryDruidHavingQueryFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidHavingQueryFilter)
		if err == nil {
			jsonTelemetryDruidHavingQueryFilter, _ := json.Marshal(dst.TelemetryDruidHavingQueryFilter)
			if string(jsonTelemetryDruidHavingQueryFilter) == "{}" { // empty struct
				dst.TelemetryDruidHavingQueryFilter = nil
			} else {
				return nil // data stored in dst.TelemetryDruidHavingQueryFilter, return on the first match
			}
		} else {
			dst.TelemetryDruidHavingQueryFilter = nil
		}
	}

	// try to unmarshal data into TelemetryDruidHavingDimensionSelectorFilter
	err = json.Unmarshal(data, &dst.TelemetryDruidHavingDimensionSelectorFilter)
	if err == nil {
		jsonTelemetryDruidHavingDimensionSelectorFilter, _ := json.Marshal(dst.TelemetryDruidHavingDimensionSelectorFilter)
		if string(jsonTelemetryDruidHavingDimensionSelectorFilter) == "{}" { // empty struct
			dst.TelemetryDruidHavingDimensionSelectorFilter = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidHavingDimensionSelectorFilter = nil
	}

	// try to unmarshal data into TelemetryDruidHavingNumericFilter
	err = json.Unmarshal(data, &dst.TelemetryDruidHavingNumericFilter)
	if err == nil {
		jsonTelemetryDruidHavingNumericFilter, _ := json.Marshal(dst.TelemetryDruidHavingNumericFilter)
		if string(jsonTelemetryDruidHavingNumericFilter) == "{}" { // empty struct
			dst.TelemetryDruidHavingNumericFilter = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidHavingNumericFilter = nil
	}

	// try to unmarshal data into TelemetryDruidHavingQueryFilter
	err = json.Unmarshal(data, &dst.TelemetryDruidHavingQueryFilter)
	if err == nil {
		jsonTelemetryDruidHavingQueryFilter, _ := json.Marshal(dst.TelemetryDruidHavingQueryFilter)
		if string(jsonTelemetryDruidHavingQueryFilter) == "{}" { // empty struct
			dst.TelemetryDruidHavingQueryFilter = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidHavingQueryFilter = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TelemetryDruidHavingDimensionSelectorFilter = nil
		dst.TelemetryDruidHavingNumericFilter = nil
		dst.TelemetryDruidHavingQueryFilter = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TelemetryDruidHavingFilter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TelemetryDruidHavingFilter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidHavingFilter) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidHavingDimensionSelectorFilter != nil {
		return json.Marshal(&src.TelemetryDruidHavingDimensionSelectorFilter)
	}

	if src.TelemetryDruidHavingNumericFilter != nil {
		return json.Marshal(&src.TelemetryDruidHavingNumericFilter)
	}

	if src.TelemetryDruidHavingQueryFilter != nil {
		return json.Marshal(&src.TelemetryDruidHavingQueryFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidHavingFilter) GetActualInstance() interface{} {
	if obj.TelemetryDruidHavingDimensionSelectorFilter != nil {
		return obj.TelemetryDruidHavingDimensionSelectorFilter
	}

	if obj.TelemetryDruidHavingNumericFilter != nil {
		return obj.TelemetryDruidHavingNumericFilter
	}

	if obj.TelemetryDruidHavingQueryFilter != nil {
		return obj.TelemetryDruidHavingQueryFilter
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidHavingFilter struct {
	value *TelemetryDruidHavingFilter
	isSet bool
}

func (v NullableTelemetryDruidHavingFilter) Get() *TelemetryDruidHavingFilter {
	return v.value
}

func (v *NullableTelemetryDruidHavingFilter) Set(val *TelemetryDruidHavingFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidHavingFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidHavingFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidHavingFilter(val *TelemetryDruidHavingFilter) *NullableTelemetryDruidHavingFilter {
	return &NullableTelemetryDruidHavingFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidHavingFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidHavingFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}