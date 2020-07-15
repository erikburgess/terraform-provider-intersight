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
	"reflect"
	"strings"
)

// AssetOrchestrationService OrchestrationService provides the necessary configuration details to enable Intersight Orchestration on the selected managed target. Subject to licensing.
type AssetOrchestrationService struct {
	AssetService
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOrchestrationService AssetOrchestrationService

// NewAssetOrchestrationService instantiates a new AssetOrchestrationService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrchestrationService() *AssetOrchestrationService {
	this := AssetOrchestrationService{}
	return &this
}

// NewAssetOrchestrationServiceWithDefaults instantiates a new AssetOrchestrationService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrchestrationServiceWithDefaults() *AssetOrchestrationService {
	this := AssetOrchestrationService{}
	return &this
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetOrchestrationService) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationService) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetOrchestrationService) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetOrchestrationService) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AssetOrchestrationService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetService, errAssetService := json.Marshal(o.AssetService)
	if errAssetService != nil {
		return []byte{}, errAssetService
	}
	errAssetService = json.Unmarshal([]byte(serializedAssetService), &toSerialize)
	if errAssetService != nil {
		return []byte{}, errAssetService
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOrchestrationService) UnmarshalJSON(bytes []byte) (err error) {
	type AssetOrchestrationServiceWithoutEmbeddedStruct struct {
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varAssetOrchestrationServiceWithoutEmbeddedStruct := AssetOrchestrationServiceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetOrchestrationServiceWithoutEmbeddedStruct)
	if err == nil {
		varAssetOrchestrationService := _AssetOrchestrationService{}
		varAssetOrchestrationService.RegisteredDevice = varAssetOrchestrationServiceWithoutEmbeddedStruct.RegisteredDevice
		*o = AssetOrchestrationService(varAssetOrchestrationService)
	} else {
		return err
	}

	varAssetOrchestrationService := _AssetOrchestrationService{}

	err = json.Unmarshal(bytes, &varAssetOrchestrationService)
	if err == nil {
		o.AssetService = varAssetOrchestrationService.AssetService
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectAssetService := reflect.ValueOf(o.AssetService)
		for i := 0; i < reflectAssetService.Type().NumField(); i++ {
			t := reflectAssetService.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetOrchestrationService struct {
	value *AssetOrchestrationService
	isSet bool
}

func (v NullableAssetOrchestrationService) Get() *AssetOrchestrationService {
	return v.value
}

func (v *NullableAssetOrchestrationService) Set(val *AssetOrchestrationService) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrchestrationService) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrchestrationService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrchestrationService(val *AssetOrchestrationService) *NullableAssetOrchestrationService {
	return &NullableAssetOrchestrationService{value: val, isSet: true}
}

func (v NullableAssetOrchestrationService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrchestrationService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
