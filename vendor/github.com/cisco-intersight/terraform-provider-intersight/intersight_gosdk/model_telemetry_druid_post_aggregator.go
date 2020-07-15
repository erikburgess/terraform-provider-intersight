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

// TelemetryDruidPostAggregator - struct for TelemetryDruidPostAggregator
type TelemetryDruidPostAggregator struct {
	TelemetryDruidArithmeticPostAggregator            *TelemetryDruidArithmeticPostAggregator
	TelemetryDruidConstantPostAggregator              *TelemetryDruidConstantPostAggregator
	TelemetryDruidFieldAccessorPostAggregator         *TelemetryDruidFieldAccessorPostAggregator
	TelemetryDruidGreatestLeastPostAggregator         *TelemetryDruidGreatestLeastPostAggregator
	TelemetryDruidHyperUniquePostAggregator           *TelemetryDruidHyperUniquePostAggregator
	TelemetryDruidThetaSketchEstimatePostAggregator   *TelemetryDruidThetaSketchEstimatePostAggregator
	TelemetryDruidThetaSketchOperationsPostAggregator *TelemetryDruidThetaSketchOperationsPostAggregator
}

// TelemetryDruidArithmeticPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidArithmeticPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidArithmeticPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidArithmeticPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidArithmeticPostAggregator: v}
}

// TelemetryDruidConstantPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidConstantPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidConstantPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidConstantPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidConstantPostAggregator: v}
}

// TelemetryDruidFieldAccessorPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidFieldAccessorPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidFieldAccessorPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidFieldAccessorPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidFieldAccessorPostAggregator: v}
}

// TelemetryDruidGreatestLeastPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidGreatestLeastPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidGreatestLeastPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidGreatestLeastPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidGreatestLeastPostAggregator: v}
}

// TelemetryDruidHyperUniquePostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidHyperUniquePostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidHyperUniquePostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidHyperUniquePostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidHyperUniquePostAggregator: v}
}

// TelemetryDruidThetaSketchEstimatePostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidThetaSketchEstimatePostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidThetaSketchEstimatePostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidThetaSketchEstimatePostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidThetaSketchEstimatePostAggregator: v}
}

// TelemetryDruidThetaSketchOperationsPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidThetaSketchOperationsPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidThetaSketchOperationsPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidThetaSketchOperationsPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidThetaSketchOperationsPostAggregator: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidPostAggregator) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'arithmetic'
	if jsonDict["type"] == "arithmetic" {
		// try to unmarshal JSON data into TelemetryDruidArithmeticPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidArithmeticPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidArithmeticPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidArithmeticPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidArithmeticPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'constant'
	if jsonDict["type"] == "constant" {
		// try to unmarshal JSON data into TelemetryDruidConstantPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidConstantPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidConstantPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidConstantPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidConstantPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleGreatest'
	if jsonDict["type"] == "doubleGreatest" {
		// try to unmarshal JSON data into TelemetryDruidGreatestLeastPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidGreatestLeastPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidGreatestLeastPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidGreatestLeastPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidGreatestLeastPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleLeast'
	if jsonDict["type"] == "doubleLeast" {
		// try to unmarshal JSON data into TelemetryDruidGreatestLeastPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidGreatestLeastPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidGreatestLeastPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidGreatestLeastPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidGreatestLeastPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'fieldAccess'
	if jsonDict["type"] == "fieldAccess" {
		// try to unmarshal JSON data into TelemetryDruidFieldAccessorPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFieldAccessorPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFieldAccessorPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidFieldAccessorPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidFieldAccessorPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'finalizingFieldAccess'
	if jsonDict["type"] == "finalizingFieldAccess" {
		// try to unmarshal JSON data into TelemetryDruidFieldAccessorPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFieldAccessorPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFieldAccessorPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidFieldAccessorPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidFieldAccessorPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'hyperUniqueCardinality'
	if jsonDict["type"] == "hyperUniqueCardinality" {
		// try to unmarshal JSON data into TelemetryDruidHyperUniquePostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidHyperUniquePostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidHyperUniquePostAggregator, return on the first match
		} else {
			dst.TelemetryDruidHyperUniquePostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidHyperUniquePostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longGreatest'
	if jsonDict["type"] == "longGreatest" {
		// try to unmarshal JSON data into TelemetryDruidGreatestLeastPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidGreatestLeastPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidGreatestLeastPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidGreatestLeastPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidGreatestLeastPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longLeast'
	if jsonDict["type"] == "longLeast" {
		// try to unmarshal JSON data into TelemetryDruidGreatestLeastPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidGreatestLeastPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidGreatestLeastPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidGreatestLeastPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidGreatestLeastPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidArithmeticPostAggregator'
	if jsonDict["type"] == "telemetry.DruidArithmeticPostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidArithmeticPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidArithmeticPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidArithmeticPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidArithmeticPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidArithmeticPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidConstantPostAggregator'
	if jsonDict["type"] == "telemetry.DruidConstantPostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidConstantPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidConstantPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidConstantPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidConstantPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidConstantPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidFieldAccessorPostAggregator'
	if jsonDict["type"] == "telemetry.DruidFieldAccessorPostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidFieldAccessorPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFieldAccessorPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFieldAccessorPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidFieldAccessorPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidFieldAccessorPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidGreatestLeastPostAggregator'
	if jsonDict["type"] == "telemetry.DruidGreatestLeastPostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidGreatestLeastPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidGreatestLeastPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidGreatestLeastPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidGreatestLeastPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidGreatestLeastPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidHyperUniquePostAggregator'
	if jsonDict["type"] == "telemetry.DruidHyperUniquePostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidHyperUniquePostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidHyperUniquePostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidHyperUniquePostAggregator, return on the first match
		} else {
			dst.TelemetryDruidHyperUniquePostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidHyperUniquePostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidThetaSketchEstimatePostAggregator'
	if jsonDict["type"] == "telemetry.DruidThetaSketchEstimatePostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidThetaSketchEstimatePostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidThetaSketchEstimatePostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidThetaSketchEstimatePostAggregator, return on the first match
		} else {
			dst.TelemetryDruidThetaSketchEstimatePostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidThetaSketchEstimatePostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidThetaSketchOperationsPostAggregator'
	if jsonDict["type"] == "telemetry.DruidThetaSketchOperationsPostAggregator" {
		// try to unmarshal JSON data into TelemetryDruidThetaSketchOperationsPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidThetaSketchOperationsPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidThetaSketchOperationsPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidThetaSketchOperationsPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidThetaSketchOperationsPostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'thetaSketchEstimate'
	if jsonDict["type"] == "thetaSketchEstimate" {
		// try to unmarshal JSON data into TelemetryDruidThetaSketchEstimatePostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidThetaSketchEstimatePostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidThetaSketchEstimatePostAggregator, return on the first match
		} else {
			dst.TelemetryDruidThetaSketchEstimatePostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidThetaSketchEstimatePostAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'thetaSketchSetOp'
	if jsonDict["type"] == "thetaSketchSetOp" {
		// try to unmarshal JSON data into TelemetryDruidThetaSketchOperationsPostAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidThetaSketchOperationsPostAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidThetaSketchOperationsPostAggregator, return on the first match
		} else {
			dst.TelemetryDruidThetaSketchOperationsPostAggregator = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidPostAggregator as TelemetryDruidThetaSketchOperationsPostAggregator: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidPostAggregator) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidArithmeticPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidArithmeticPostAggregator)
	}

	if src.TelemetryDruidConstantPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidConstantPostAggregator)
	}

	if src.TelemetryDruidFieldAccessorPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidFieldAccessorPostAggregator)
	}

	if src.TelemetryDruidGreatestLeastPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidGreatestLeastPostAggregator)
	}

	if src.TelemetryDruidHyperUniquePostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidHyperUniquePostAggregator)
	}

	if src.TelemetryDruidThetaSketchEstimatePostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidThetaSketchEstimatePostAggregator)
	}

	if src.TelemetryDruidThetaSketchOperationsPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidThetaSketchOperationsPostAggregator)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidPostAggregator) GetActualInstance() interface{} {
	if obj.TelemetryDruidArithmeticPostAggregator != nil {
		return obj.TelemetryDruidArithmeticPostAggregator
	}

	if obj.TelemetryDruidConstantPostAggregator != nil {
		return obj.TelemetryDruidConstantPostAggregator
	}

	if obj.TelemetryDruidFieldAccessorPostAggregator != nil {
		return obj.TelemetryDruidFieldAccessorPostAggregator
	}

	if obj.TelemetryDruidGreatestLeastPostAggregator != nil {
		return obj.TelemetryDruidGreatestLeastPostAggregator
	}

	if obj.TelemetryDruidHyperUniquePostAggregator != nil {
		return obj.TelemetryDruidHyperUniquePostAggregator
	}

	if obj.TelemetryDruidThetaSketchEstimatePostAggregator != nil {
		return obj.TelemetryDruidThetaSketchEstimatePostAggregator
	}

	if obj.TelemetryDruidThetaSketchOperationsPostAggregator != nil {
		return obj.TelemetryDruidThetaSketchOperationsPostAggregator
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidPostAggregator struct {
	value *TelemetryDruidPostAggregator
	isSet bool
}

func (v NullableTelemetryDruidPostAggregator) Get() *TelemetryDruidPostAggregator {
	return v.value
}

func (v *NullableTelemetryDruidPostAggregator) Set(val *TelemetryDruidPostAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidPostAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidPostAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidPostAggregator(val *TelemetryDruidPostAggregator) *NullableTelemetryDruidPostAggregator {
	return &NullableTelemetryDruidPostAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidPostAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidPostAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
