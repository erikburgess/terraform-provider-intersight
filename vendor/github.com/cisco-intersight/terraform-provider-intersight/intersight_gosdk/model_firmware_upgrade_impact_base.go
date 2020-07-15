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

// FirmwareUpgradeImpactBase Before submitting firmware upgrade operation, the estimate impact helps to know the list of components be impacted and require host reboot. This cannot be used for network share based upgrade.
type FirmwareUpgradeImpactBase struct {
	MoBaseMo
	Components *[]string `json:"Components,omitempty"`
	// Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint.
	ComputationState  *string               `json:"ComputationState,omitempty"`
	ExcludeComponents *[]string             `json:"ExcludeComponents,omitempty"`
	Impacts           *[]FirmwareBaseImpact `json:"Impacts,omitempty"`
	// The summary on the component or components getting impacted by the upgrade.
	Summary              *string `json:"Summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeImpactBase FirmwareUpgradeImpactBase

// NewFirmwareUpgradeImpactBase instantiates a new FirmwareUpgradeImpactBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeImpactBase() *FirmwareUpgradeImpactBase {
	this := FirmwareUpgradeImpactBase{}
	var computationState string = "Inprogress"
	this.ComputationState = &computationState
	var summary string = "Basic"
	this.Summary = &summary
	return &this
}

// NewFirmwareUpgradeImpactBaseWithDefaults instantiates a new FirmwareUpgradeImpactBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeImpactBaseWithDefaults() *FirmwareUpgradeImpactBase {
	this := FirmwareUpgradeImpactBase{}
	var computationState string = "Inprogress"
	this.ComputationState = &computationState
	var summary string = "Basic"
	this.Summary = &summary
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactBase) GetComponents() []string {
	if o == nil || o.Components == nil {
		var ret []string
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactBase) GetComponentsOk() (*[]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactBase) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *FirmwareUpgradeImpactBase) SetComponents(v []string) {
	o.Components = &v
}

// GetComputationState returns the ComputationState field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactBase) GetComputationState() string {
	if o == nil || o.ComputationState == nil {
		var ret string
		return ret
	}
	return *o.ComputationState
}

// GetComputationStateOk returns a tuple with the ComputationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactBase) GetComputationStateOk() (*string, bool) {
	if o == nil || o.ComputationState == nil {
		return nil, false
	}
	return o.ComputationState, true
}

// HasComputationState returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactBase) HasComputationState() bool {
	if o != nil && o.ComputationState != nil {
		return true
	}

	return false
}

// SetComputationState gets a reference to the given string and assigns it to the ComputationState field.
func (o *FirmwareUpgradeImpactBase) SetComputationState(v string) {
	o.ComputationState = &v
}

// GetExcludeComponents returns the ExcludeComponents field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactBase) GetExcludeComponents() []string {
	if o == nil || o.ExcludeComponents == nil {
		var ret []string
		return ret
	}
	return *o.ExcludeComponents
}

// GetExcludeComponentsOk returns a tuple with the ExcludeComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactBase) GetExcludeComponentsOk() (*[]string, bool) {
	if o == nil || o.ExcludeComponents == nil {
		return nil, false
	}
	return o.ExcludeComponents, true
}

// HasExcludeComponents returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactBase) HasExcludeComponents() bool {
	if o != nil && o.ExcludeComponents != nil {
		return true
	}

	return false
}

// SetExcludeComponents gets a reference to the given []string and assigns it to the ExcludeComponents field.
func (o *FirmwareUpgradeImpactBase) SetExcludeComponents(v []string) {
	o.ExcludeComponents = &v
}

// GetImpacts returns the Impacts field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactBase) GetImpacts() []FirmwareBaseImpact {
	if o == nil || o.Impacts == nil {
		var ret []FirmwareBaseImpact
		return ret
	}
	return *o.Impacts
}

// GetImpactsOk returns a tuple with the Impacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactBase) GetImpactsOk() (*[]FirmwareBaseImpact, bool) {
	if o == nil || o.Impacts == nil {
		return nil, false
	}
	return o.Impacts, true
}

// HasImpacts returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactBase) HasImpacts() bool {
	if o != nil && o.Impacts != nil {
		return true
	}

	return false
}

// SetImpacts gets a reference to the given []FirmwareBaseImpact and assigns it to the Impacts field.
func (o *FirmwareUpgradeImpactBase) SetImpacts(v []FirmwareBaseImpact) {
	o.Impacts = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactBase) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactBase) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactBase) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *FirmwareUpgradeImpactBase) SetSummary(v string) {
	o.Summary = &v
}

func (o FirmwareUpgradeImpactBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Components != nil {
		toSerialize["Components"] = o.Components
	}
	if o.ComputationState != nil {
		toSerialize["ComputationState"] = o.ComputationState
	}
	if o.ExcludeComponents != nil {
		toSerialize["ExcludeComponents"] = o.ExcludeComponents
	}
	if o.Impacts != nil {
		toSerialize["Impacts"] = o.Impacts
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeImpactBase) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUpgradeImpactBaseWithoutEmbeddedStruct struct {
		Components *[]string `json:"Components,omitempty"`
		// Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint.
		ComputationState  *string               `json:"ComputationState,omitempty"`
		ExcludeComponents *[]string             `json:"ExcludeComponents,omitempty"`
		Impacts           *[]FirmwareBaseImpact `json:"Impacts,omitempty"`
		// The summary on the component or components getting impacted by the upgrade.
		Summary *string `json:"Summary,omitempty"`
	}

	varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct := FirmwareUpgradeImpactBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUpgradeImpactBase := _FirmwareUpgradeImpactBase{}
		varFirmwareUpgradeImpactBase.Components = varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct.Components
		varFirmwareUpgradeImpactBase.ComputationState = varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct.ComputationState
		varFirmwareUpgradeImpactBase.ExcludeComponents = varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct.ExcludeComponents
		varFirmwareUpgradeImpactBase.Impacts = varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct.Impacts
		varFirmwareUpgradeImpactBase.Summary = varFirmwareUpgradeImpactBaseWithoutEmbeddedStruct.Summary
		*o = FirmwareUpgradeImpactBase(varFirmwareUpgradeImpactBase)
	} else {
		return err
	}

	varFirmwareUpgradeImpactBase := _FirmwareUpgradeImpactBase{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeImpactBase)
	if err == nil {
		o.MoBaseMo = varFirmwareUpgradeImpactBase.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Components")
		delete(additionalProperties, "ComputationState")
		delete(additionalProperties, "ExcludeComponents")
		delete(additionalProperties, "Impacts")
		delete(additionalProperties, "Summary")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableFirmwareUpgradeImpactBase struct {
	value *FirmwareUpgradeImpactBase
	isSet bool
}

func (v NullableFirmwareUpgradeImpactBase) Get() *FirmwareUpgradeImpactBase {
	return v.value
}

func (v *NullableFirmwareUpgradeImpactBase) Set(val *FirmwareUpgradeImpactBase) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeImpactBase) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeImpactBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeImpactBase(val *FirmwareUpgradeImpactBase) *NullableFirmwareUpgradeImpactBase {
	return &NullableFirmwareUpgradeImpactBase{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeImpactBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeImpactBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
