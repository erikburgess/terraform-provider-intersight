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
)

// FabricEthNetworkControlPolicyAllOf Definition of the list of properties defined in 'fabric.EthNetworkControlPolicy', excluding properties defined in parent classes.
type FabricEthNetworkControlPolicyAllOf struct {
	// Enables the CDP on an interface.
	CdpEnabled *bool `json:"CdpEnabled,omitempty"`
	// Determines if the MAC forging is allowed or denied on an interface.
	ForgeMac     *string             `json:"ForgeMac,omitempty"`
	LldpSettings *FabricLldpSettings `json:"LldpSettings,omitempty"`
	// Determines the MAC addresses that have to be registered with the switch.
	MacRegisterMode *string `json:"MacRegisterMode,omitempty"`
	// Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned.
	UplinkFailAction *string `json:"UplinkFailAction,omitempty"`
	// An array of relationships to vnicEthNetworkPolicy resources.
	NetworkPolicy        []VnicEthNetworkPolicyRelationship    `json:"NetworkPolicy,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEthNetworkControlPolicyAllOf FabricEthNetworkControlPolicyAllOf

// NewFabricEthNetworkControlPolicyAllOf instantiates a new FabricEthNetworkControlPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEthNetworkControlPolicyAllOf() *FabricEthNetworkControlPolicyAllOf {
	this := FabricEthNetworkControlPolicyAllOf{}
	var forgeMac string = "allow"
	this.ForgeMac = &forgeMac
	var macRegisterMode string = "nativeVlanOnly"
	this.MacRegisterMode = &macRegisterMode
	var uplinkFailAction string = "linkDown"
	this.UplinkFailAction = &uplinkFailAction
	return &this
}

// NewFabricEthNetworkControlPolicyAllOfWithDefaults instantiates a new FabricEthNetworkControlPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEthNetworkControlPolicyAllOfWithDefaults() *FabricEthNetworkControlPolicyAllOf {
	this := FabricEthNetworkControlPolicyAllOf{}
	var forgeMac string = "allow"
	this.ForgeMac = &forgeMac
	var macRegisterMode string = "nativeVlanOnly"
	this.MacRegisterMode = &macRegisterMode
	var uplinkFailAction string = "linkDown"
	this.UplinkFailAction = &uplinkFailAction
	return &this
}

// GetCdpEnabled returns the CdpEnabled field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetCdpEnabled() bool {
	if o == nil || o.CdpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CdpEnabled
}

// GetCdpEnabledOk returns a tuple with the CdpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetCdpEnabledOk() (*bool, bool) {
	if o == nil || o.CdpEnabled == nil {
		return nil, false
	}
	return o.CdpEnabled, true
}

// HasCdpEnabled returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasCdpEnabled() bool {
	if o != nil && o.CdpEnabled != nil {
		return true
	}

	return false
}

// SetCdpEnabled gets a reference to the given bool and assigns it to the CdpEnabled field.
func (o *FabricEthNetworkControlPolicyAllOf) SetCdpEnabled(v bool) {
	o.CdpEnabled = &v
}

// GetForgeMac returns the ForgeMac field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetForgeMac() string {
	if o == nil || o.ForgeMac == nil {
		var ret string
		return ret
	}
	return *o.ForgeMac
}

// GetForgeMacOk returns a tuple with the ForgeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetForgeMacOk() (*string, bool) {
	if o == nil || o.ForgeMac == nil {
		return nil, false
	}
	return o.ForgeMac, true
}

// HasForgeMac returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasForgeMac() bool {
	if o != nil && o.ForgeMac != nil {
		return true
	}

	return false
}

// SetForgeMac gets a reference to the given string and assigns it to the ForgeMac field.
func (o *FabricEthNetworkControlPolicyAllOf) SetForgeMac(v string) {
	o.ForgeMac = &v
}

// GetLldpSettings returns the LldpSettings field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetLldpSettings() FabricLldpSettings {
	if o == nil || o.LldpSettings == nil {
		var ret FabricLldpSettings
		return ret
	}
	return *o.LldpSettings
}

// GetLldpSettingsOk returns a tuple with the LldpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetLldpSettingsOk() (*FabricLldpSettings, bool) {
	if o == nil || o.LldpSettings == nil {
		return nil, false
	}
	return o.LldpSettings, true
}

// HasLldpSettings returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasLldpSettings() bool {
	if o != nil && o.LldpSettings != nil {
		return true
	}

	return false
}

// SetLldpSettings gets a reference to the given FabricLldpSettings and assigns it to the LldpSettings field.
func (o *FabricEthNetworkControlPolicyAllOf) SetLldpSettings(v FabricLldpSettings) {
	o.LldpSettings = &v
}

// GetMacRegisterMode returns the MacRegisterMode field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetMacRegisterMode() string {
	if o == nil || o.MacRegisterMode == nil {
		var ret string
		return ret
	}
	return *o.MacRegisterMode
}

// GetMacRegisterModeOk returns a tuple with the MacRegisterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetMacRegisterModeOk() (*string, bool) {
	if o == nil || o.MacRegisterMode == nil {
		return nil, false
	}
	return o.MacRegisterMode, true
}

// HasMacRegisterMode returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasMacRegisterMode() bool {
	if o != nil && o.MacRegisterMode != nil {
		return true
	}

	return false
}

// SetMacRegisterMode gets a reference to the given string and assigns it to the MacRegisterMode field.
func (o *FabricEthNetworkControlPolicyAllOf) SetMacRegisterMode(v string) {
	o.MacRegisterMode = &v
}

// GetUplinkFailAction returns the UplinkFailAction field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetUplinkFailAction() string {
	if o == nil || o.UplinkFailAction == nil {
		var ret string
		return ret
	}
	return *o.UplinkFailAction
}

// GetUplinkFailActionOk returns a tuple with the UplinkFailAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetUplinkFailActionOk() (*string, bool) {
	if o == nil || o.UplinkFailAction == nil {
		return nil, false
	}
	return o.UplinkFailAction, true
}

// HasUplinkFailAction returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasUplinkFailAction() bool {
	if o != nil && o.UplinkFailAction != nil {
		return true
	}

	return false
}

// SetUplinkFailAction gets a reference to the given string and assigns it to the UplinkFailAction field.
func (o *FabricEthNetworkControlPolicyAllOf) SetUplinkFailAction(v string) {
	o.UplinkFailAction = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkControlPolicyAllOf) GetNetworkPolicy() []VnicEthNetworkPolicyRelationship {
	if o == nil {
		var ret []VnicEthNetworkPolicyRelationship
		return ret
	}
	return o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkControlPolicyAllOf) GetNetworkPolicyOk() (*[]VnicEthNetworkPolicyRelationship, bool) {
	if o == nil || o.NetworkPolicy == nil {
		return nil, false
	}
	return &o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasNetworkPolicy() bool {
	if o != nil && o.NetworkPolicy != nil {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given []VnicEthNetworkPolicyRelationship and assigns it to the NetworkPolicy field.
func (o *FabricEthNetworkControlPolicyAllOf) SetNetworkPolicy(v []VnicEthNetworkPolicyRelationship) {
	o.NetworkPolicy = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricEthNetworkControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricEthNetworkControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CdpEnabled != nil {
		toSerialize["CdpEnabled"] = o.CdpEnabled
	}
	if o.ForgeMac != nil {
		toSerialize["ForgeMac"] = o.ForgeMac
	}
	if o.LldpSettings != nil {
		toSerialize["LldpSettings"] = o.LldpSettings
	}
	if o.MacRegisterMode != nil {
		toSerialize["MacRegisterMode"] = o.MacRegisterMode
	}
	if o.UplinkFailAction != nil {
		toSerialize["UplinkFailAction"] = o.UplinkFailAction
	}
	if o.NetworkPolicy != nil {
		toSerialize["NetworkPolicy"] = o.NetworkPolicy
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEthNetworkControlPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricEthNetworkControlPolicyAllOf := _FabricEthNetworkControlPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricEthNetworkControlPolicyAllOf); err == nil {
		*o = FabricEthNetworkControlPolicyAllOf(varFabricEthNetworkControlPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CdpEnabled")
		delete(additionalProperties, "ForgeMac")
		delete(additionalProperties, "LldpSettings")
		delete(additionalProperties, "MacRegisterMode")
		delete(additionalProperties, "UplinkFailAction")
		delete(additionalProperties, "NetworkPolicy")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricEthNetworkControlPolicyAllOf struct {
	value *FabricEthNetworkControlPolicyAllOf
	isSet bool
}

func (v NullableFabricEthNetworkControlPolicyAllOf) Get() *FabricEthNetworkControlPolicyAllOf {
	return v.value
}

func (v *NullableFabricEthNetworkControlPolicyAllOf) Set(val *FabricEthNetworkControlPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEthNetworkControlPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEthNetworkControlPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEthNetworkControlPolicyAllOf(val *FabricEthNetworkControlPolicyAllOf) *NullableFabricEthNetworkControlPolicyAllOf {
	return &NullableFabricEthNetworkControlPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricEthNetworkControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEthNetworkControlPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
