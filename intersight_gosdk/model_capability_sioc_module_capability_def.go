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

// CapabilitySiocModuleCapabilityDef Chassis SIOC module capabilities.
type CapabilitySiocModuleCapabilityDef struct {
	CapabilityCapability
	// Device connector support on SIOC.
	DcSupported          *bool `json:"DcSupported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySiocModuleCapabilityDef CapabilitySiocModuleCapabilityDef

// NewCapabilitySiocModuleCapabilityDef instantiates a new CapabilitySiocModuleCapabilityDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySiocModuleCapabilityDef() *CapabilitySiocModuleCapabilityDef {
	this := CapabilitySiocModuleCapabilityDef{}
	return &this
}

// NewCapabilitySiocModuleCapabilityDefWithDefaults instantiates a new CapabilitySiocModuleCapabilityDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySiocModuleCapabilityDefWithDefaults() *CapabilitySiocModuleCapabilityDef {
	this := CapabilitySiocModuleCapabilityDef{}
	return &this
}

// GetDcSupported returns the DcSupported field value if set, zero value otherwise.
func (o *CapabilitySiocModuleCapabilityDef) GetDcSupported() bool {
	if o == nil || o.DcSupported == nil {
		var ret bool
		return ret
	}
	return *o.DcSupported
}

// GetDcSupportedOk returns a tuple with the DcSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleCapabilityDef) GetDcSupportedOk() (*bool, bool) {
	if o == nil || o.DcSupported == nil {
		return nil, false
	}
	return o.DcSupported, true
}

// HasDcSupported returns a boolean if a field has been set.
func (o *CapabilitySiocModuleCapabilityDef) HasDcSupported() bool {
	if o != nil && o.DcSupported != nil {
		return true
	}

	return false
}

// SetDcSupported gets a reference to the given bool and assigns it to the DcSupported field.
func (o *CapabilitySiocModuleCapabilityDef) SetDcSupported(v bool) {
	o.DcSupported = &v
}

func (o CapabilitySiocModuleCapabilityDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if o.DcSupported != nil {
		toSerialize["DcSupported"] = o.DcSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySiocModuleCapabilityDef) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySiocModuleCapabilityDefWithoutEmbeddedStruct struct {
		// Device connector support on SIOC.
		DcSupported *bool `json:"DcSupported,omitempty"`
	}

	varCapabilitySiocModuleCapabilityDefWithoutEmbeddedStruct := CapabilitySiocModuleCapabilityDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySiocModuleCapabilityDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySiocModuleCapabilityDef := _CapabilitySiocModuleCapabilityDef{}
		varCapabilitySiocModuleCapabilityDef.DcSupported = varCapabilitySiocModuleCapabilityDefWithoutEmbeddedStruct.DcSupported
		*o = CapabilitySiocModuleCapabilityDef(varCapabilitySiocModuleCapabilityDef)
	} else {
		return err
	}

	varCapabilitySiocModuleCapabilityDef := _CapabilitySiocModuleCapabilityDef{}

	err = json.Unmarshal(bytes, &varCapabilitySiocModuleCapabilityDef)
	if err == nil {
		o.CapabilityCapability = varCapabilitySiocModuleCapabilityDef.CapabilityCapability
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DcSupported")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilitySiocModuleCapabilityDef struct {
	value *CapabilitySiocModuleCapabilityDef
	isSet bool
}

func (v NullableCapabilitySiocModuleCapabilityDef) Get() *CapabilitySiocModuleCapabilityDef {
	return v.value
}

func (v *NullableCapabilitySiocModuleCapabilityDef) Set(val *CapabilitySiocModuleCapabilityDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySiocModuleCapabilityDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySiocModuleCapabilityDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySiocModuleCapabilityDef(val *CapabilitySiocModuleCapabilityDef) *NullableCapabilitySiocModuleCapabilityDef {
	return &NullableCapabilitySiocModuleCapabilityDef{value: val, isSet: true}
}

func (v NullableCapabilitySiocModuleCapabilityDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySiocModuleCapabilityDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
