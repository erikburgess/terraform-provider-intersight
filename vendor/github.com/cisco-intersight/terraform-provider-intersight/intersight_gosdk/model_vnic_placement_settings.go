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

// VnicPlacementSettings Placement Settings for the virtual interface.
type VnicPlacementSettings struct {
	MoBaseComplexType
	// PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
	Id *string `json:"Id,omitempty"`
	// The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two.
	PciLink *int64 `json:"PciLink,omitempty"`
	// The fabric port to which the vnics will be associated.
	SwitchId *string `json:"SwitchId,omitempty"`
	// Adapter port on which the virtual interface will be created.
	Uplink               *int64 `json:"Uplink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicPlacementSettings VnicPlacementSettings

// NewVnicPlacementSettings instantiates a new VnicPlacementSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicPlacementSettings() *VnicPlacementSettings {
	this := VnicPlacementSettings{}
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// NewVnicPlacementSettingsWithDefaults instantiates a new VnicPlacementSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicPlacementSettingsWithDefaults() *VnicPlacementSettings {
	this := VnicPlacementSettings{}
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VnicPlacementSettings) SetId(v string) {
	o.Id = &v
}

// GetPciLink returns the PciLink field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetPciLink() int64 {
	if o == nil || o.PciLink == nil {
		var ret int64
		return ret
	}
	return *o.PciLink
}

// GetPciLinkOk returns a tuple with the PciLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetPciLinkOk() (*int64, bool) {
	if o == nil || o.PciLink == nil {
		return nil, false
	}
	return o.PciLink, true
}

// HasPciLink returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasPciLink() bool {
	if o != nil && o.PciLink != nil {
		return true
	}

	return false
}

// SetPciLink gets a reference to the given int64 and assigns it to the PciLink field.
func (o *VnicPlacementSettings) SetPciLink(v int64) {
	o.PciLink = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *VnicPlacementSettings) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *VnicPlacementSettings) GetUplink() int64 {
	if o == nil || o.Uplink == nil {
		var ret int64
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettings) GetUplinkOk() (*int64, bool) {
	if o == nil || o.Uplink == nil {
		return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *VnicPlacementSettings) HasUplink() bool {
	if o != nil && o.Uplink != nil {
		return true
	}

	return false
}

// SetUplink gets a reference to the given int64 and assigns it to the Uplink field.
func (o *VnicPlacementSettings) SetUplink(v int64) {
	o.Uplink = &v
}

func (o VnicPlacementSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.PciLink != nil {
		toSerialize["PciLink"] = o.PciLink
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Uplink != nil {
		toSerialize["Uplink"] = o.Uplink
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicPlacementSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicPlacementSettingsWithoutEmbeddedStruct struct {
		// PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
		Id *string `json:"Id,omitempty"`
		// The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two.
		PciLink *int64 `json:"PciLink,omitempty"`
		// The fabric port to which the vnics will be associated.
		SwitchId *string `json:"SwitchId,omitempty"`
		// Adapter port on which the virtual interface will be created.
		Uplink *int64 `json:"Uplink,omitempty"`
	}

	varVnicPlacementSettingsWithoutEmbeddedStruct := VnicPlacementSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicPlacementSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicPlacementSettings := _VnicPlacementSettings{}
		varVnicPlacementSettings.Id = varVnicPlacementSettingsWithoutEmbeddedStruct.Id
		varVnicPlacementSettings.PciLink = varVnicPlacementSettingsWithoutEmbeddedStruct.PciLink
		varVnicPlacementSettings.SwitchId = varVnicPlacementSettingsWithoutEmbeddedStruct.SwitchId
		varVnicPlacementSettings.Uplink = varVnicPlacementSettingsWithoutEmbeddedStruct.Uplink
		*o = VnicPlacementSettings(varVnicPlacementSettings)
	} else {
		return err
	}

	varVnicPlacementSettings := _VnicPlacementSettings{}

	err = json.Unmarshal(bytes, &varVnicPlacementSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicPlacementSettings.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Id")
		delete(additionalProperties, "PciLink")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Uplink")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicPlacementSettings struct {
	value *VnicPlacementSettings
	isSet bool
}

func (v NullableVnicPlacementSettings) Get() *VnicPlacementSettings {
	return v.value
}

func (v *NullableVnicPlacementSettings) Set(val *VnicPlacementSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicPlacementSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicPlacementSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicPlacementSettings(val *VnicPlacementSettings) *NullableVnicPlacementSettings {
	return &NullableVnicPlacementSettings{value: val, isSet: true}
}

func (v NullableVnicPlacementSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicPlacementSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
