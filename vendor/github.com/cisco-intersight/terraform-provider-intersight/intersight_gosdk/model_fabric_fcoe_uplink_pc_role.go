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
	"reflect"
	"strings"
)

// FabricFcoeUplinkPcRole Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
type FabricFcoeUplinkPcRole struct {
	FabricPortChannelRole
	// Admin configured speed for the port.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Admin configured state for UDLD for this port.
	UdldAdminState       *string `json:"UdldAdminState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcoeUplinkPcRole FabricFcoeUplinkPcRole

// NewFabricFcoeUplinkPcRole instantiates a new FabricFcoeUplinkPcRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcoeUplinkPcRole() *FabricFcoeUplinkPcRole {
	this := FabricFcoeUplinkPcRole{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// NewFabricFcoeUplinkPcRoleWithDefaults instantiates a new FabricFcoeUplinkPcRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcoeUplinkPcRoleWithDefaults() *FabricFcoeUplinkPcRole {
	this := FabricFcoeUplinkPcRole{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcoeUplinkPcRole) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcoeUplinkPcRole) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcoeUplinkPcRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetUdldAdminState returns the UdldAdminState field value if set, zero value otherwise.
func (o *FabricFcoeUplinkPcRole) GetUdldAdminState() string {
	if o == nil || o.UdldAdminState == nil {
		var ret string
		return ret
	}
	return *o.UdldAdminState
}

// GetUdldAdminStateOk returns a tuple with the UdldAdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcoeUplinkPcRole) GetUdldAdminStateOk() (*string, bool) {
	if o == nil || o.UdldAdminState == nil {
		return nil, false
	}
	return o.UdldAdminState, true
}

// HasUdldAdminState returns a boolean if a field has been set.
func (o *FabricFcoeUplinkPcRole) HasUdldAdminState() bool {
	if o != nil && o.UdldAdminState != nil {
		return true
	}

	return false
}

// SetUdldAdminState gets a reference to the given string and assigns it to the UdldAdminState field.
func (o *FabricFcoeUplinkPcRole) SetUdldAdminState(v string) {
	o.UdldAdminState = &v
}

func (o FabricFcoeUplinkPcRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortChannelRole, errFabricPortChannelRole := json.Marshal(o.FabricPortChannelRole)
	if errFabricPortChannelRole != nil {
		return []byte{}, errFabricPortChannelRole
	}
	errFabricPortChannelRole = json.Unmarshal([]byte(serializedFabricPortChannelRole), &toSerialize)
	if errFabricPortChannelRole != nil {
		return []byte{}, errFabricPortChannelRole
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.UdldAdminState != nil {
		toSerialize["UdldAdminState"] = o.UdldAdminState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcoeUplinkPcRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcoeUplinkPcRoleWithoutEmbeddedStruct struct {
		// Admin configured speed for the port.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Admin configured state for UDLD for this port.
		UdldAdminState *string `json:"UdldAdminState,omitempty"`
	}

	varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct := FabricFcoeUplinkPcRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcoeUplinkPcRole := _FabricFcoeUplinkPcRole{}
		varFabricFcoeUplinkPcRole.AdminSpeed = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricFcoeUplinkPcRole.UdldAdminState = varFabricFcoeUplinkPcRoleWithoutEmbeddedStruct.UdldAdminState
		*o = FabricFcoeUplinkPcRole(varFabricFcoeUplinkPcRole)
	} else {
		return err
	}

	varFabricFcoeUplinkPcRole := _FabricFcoeUplinkPcRole{}

	err = json.Unmarshal(bytes, &varFabricFcoeUplinkPcRole)
	if err == nil {
		o.FabricPortChannelRole = varFabricFcoeUplinkPcRole.FabricPortChannelRole
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "UdldAdminState")

		// remove fields from embedded structs
		reflectFabricPortChannelRole := reflect.ValueOf(o.FabricPortChannelRole)
		for i := 0; i < reflectFabricPortChannelRole.Type().NumField(); i++ {
			t := reflectFabricPortChannelRole.Type().Field(i)

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

type NullableFabricFcoeUplinkPcRole struct {
	value *FabricFcoeUplinkPcRole
	isSet bool
}

func (v NullableFabricFcoeUplinkPcRole) Get() *FabricFcoeUplinkPcRole {
	return v.value
}

func (v *NullableFabricFcoeUplinkPcRole) Set(val *FabricFcoeUplinkPcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcoeUplinkPcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcoeUplinkPcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcoeUplinkPcRole(val *FabricFcoeUplinkPcRole) *NullableFabricFcoeUplinkPcRole {
	return &NullableFabricFcoeUplinkPcRole{value: val, isSet: true}
}

func (v NullableFabricFcoeUplinkPcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcoeUplinkPcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
