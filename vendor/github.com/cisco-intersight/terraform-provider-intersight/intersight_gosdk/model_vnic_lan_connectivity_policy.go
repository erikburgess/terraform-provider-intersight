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

// VnicLanConnectivityPolicy A LAN Connectivity Policy determines the network resources and the connections between the server and the LAN on the network. This policy uses Consistent Device Naming to configure the vNIC. You can configure a usNIC or VMQ connection for the vNIC to improve network performance.
type VnicLanConnectivityPolicy struct {
	PolicyAbstractPolicy
	// The mode used for placement of vnics on network adapters. It can either be Auto or Custom.
	PlacementMode *string `json:"PlacementMode,omitempty"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// An array of relationships to vnicEthIf resources.
	EthIfs       []VnicEthIfRelationship               `json:"EthIfs,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicLanConnectivityPolicy VnicLanConnectivityPolicy

// NewVnicLanConnectivityPolicy instantiates a new VnicLanConnectivityPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicLanConnectivityPolicy() *VnicLanConnectivityPolicy {
	this := VnicLanConnectivityPolicy{}
	var placementMode string = "custom"
	this.PlacementMode = &placementMode
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewVnicLanConnectivityPolicyWithDefaults instantiates a new VnicLanConnectivityPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicLanConnectivityPolicyWithDefaults() *VnicLanConnectivityPolicy {
	this := VnicLanConnectivityPolicy{}
	var placementMode string = "custom"
	this.PlacementMode = &placementMode
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetPlacementMode returns the PlacementMode field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetPlacementMode() string {
	if o == nil || o.PlacementMode == nil {
		var ret string
		return ret
	}
	return *o.PlacementMode
}

// GetPlacementModeOk returns a tuple with the PlacementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetPlacementModeOk() (*string, bool) {
	if o == nil || o.PlacementMode == nil {
		return nil, false
	}
	return o.PlacementMode, true
}

// HasPlacementMode returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasPlacementMode() bool {
	if o != nil && o.PlacementMode != nil {
		return true
	}

	return false
}

// SetPlacementMode gets a reference to the given string and assigns it to the PlacementMode field.
func (o *VnicLanConnectivityPolicy) SetPlacementMode(v string) {
	o.PlacementMode = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *VnicLanConnectivityPolicy) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetEthIfs returns the EthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicLanConnectivityPolicy) GetEthIfs() []VnicEthIfRelationship {
	if o == nil {
		var ret []VnicEthIfRelationship
		return ret
	}
	return o.EthIfs
}

// GetEthIfsOk returns a tuple with the EthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicLanConnectivityPolicy) GetEthIfsOk() (*[]VnicEthIfRelationship, bool) {
	if o == nil || o.EthIfs == nil {
		return nil, false
	}
	return &o.EthIfs, true
}

// HasEthIfs returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasEthIfs() bool {
	if o != nil && o.EthIfs != nil {
		return true
	}

	return false
}

// SetEthIfs gets a reference to the given []VnicEthIfRelationship and assigns it to the EthIfs field.
func (o *VnicLanConnectivityPolicy) SetEthIfs(v []VnicEthIfRelationship) {
	o.EthIfs = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicLanConnectivityPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicLanConnectivityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicLanConnectivityPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *VnicLanConnectivityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o VnicLanConnectivityPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.PlacementMode != nil {
		toSerialize["PlacementMode"] = o.PlacementMode
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.EthIfs != nil {
		toSerialize["EthIfs"] = o.EthIfs
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicLanConnectivityPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicLanConnectivityPolicyWithoutEmbeddedStruct struct {
		// The mode used for placement of vnics on network adapters. It can either be Auto or Custom.
		PlacementMode *string `json:"PlacementMode,omitempty"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.
		TargetPlatform *string `json:"TargetPlatform,omitempty"`
		// An array of relationships to vnicEthIf resources.
		EthIfs       []VnicEthIfRelationship               `json:"EthIfs,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varVnicLanConnectivityPolicyWithoutEmbeddedStruct := VnicLanConnectivityPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicLanConnectivityPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicLanConnectivityPolicy := _VnicLanConnectivityPolicy{}
		varVnicLanConnectivityPolicy.PlacementMode = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.PlacementMode
		varVnicLanConnectivityPolicy.TargetPlatform = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.TargetPlatform
		varVnicLanConnectivityPolicy.EthIfs = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.EthIfs
		varVnicLanConnectivityPolicy.Organization = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.Organization
		varVnicLanConnectivityPolicy.Profiles = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.Profiles
		*o = VnicLanConnectivityPolicy(varVnicLanConnectivityPolicy)
	} else {
		return err
	}

	varVnicLanConnectivityPolicy := _VnicLanConnectivityPolicy{}

	err = json.Unmarshal(bytes, &varVnicLanConnectivityPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicLanConnectivityPolicy.PolicyAbstractPolicy
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "PlacementMode")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "EthIfs")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableVnicLanConnectivityPolicy struct {
	value *VnicLanConnectivityPolicy
	isSet bool
}

func (v NullableVnicLanConnectivityPolicy) Get() *VnicLanConnectivityPolicy {
	return v.value
}

func (v *NullableVnicLanConnectivityPolicy) Set(val *VnicLanConnectivityPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLanConnectivityPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLanConnectivityPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLanConnectivityPolicy(val *VnicLanConnectivityPolicy) *NullableVnicLanConnectivityPolicy {
	return &NullableVnicLanConnectivityPolicy{value: val, isSet: true}
}

func (v NullableVnicLanConnectivityPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLanConnectivityPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
