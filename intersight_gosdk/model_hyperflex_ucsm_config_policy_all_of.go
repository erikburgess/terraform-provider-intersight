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
)

// HyperflexUcsmConfigPolicyAllOf Definition of the list of properties defined in 'hyperflex.UcsmConfigPolicy', excluding properties defined in parent classes.
type HyperflexUcsmConfigPolicyAllOf struct {
	KvmIpRange     *HyperflexIpAddrRange        `json:"KvmIpRange,omitempty"`
	MacPrefixRange *HyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty"`
	// The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc.
	ServerFirmwareVersion *string `json:"ServerFirmwareVersion,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
}

// NewHyperflexUcsmConfigPolicyAllOf instantiates a new HyperflexUcsmConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexUcsmConfigPolicyAllOf() *HyperflexUcsmConfigPolicyAllOf {
	this := HyperflexUcsmConfigPolicyAllOf{}
	return &this
}

// NewHyperflexUcsmConfigPolicyAllOfWithDefaults instantiates a new HyperflexUcsmConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexUcsmConfigPolicyAllOfWithDefaults() *HyperflexUcsmConfigPolicyAllOf {
	this := HyperflexUcsmConfigPolicyAllOf{}
	return &this
}

// GetKvmIpRange returns the KvmIpRange field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicyAllOf) GetKvmIpRange() HyperflexIpAddrRange {
	if o == nil || o.KvmIpRange == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.KvmIpRange
}

// GetKvmIpRangeOk returns a tuple with the KvmIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil || o.KvmIpRange == nil {
		return nil, false
	}
	return o.KvmIpRange, true
}

// HasKvmIpRange returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasKvmIpRange() bool {
	if o != nil && o.KvmIpRange != nil {
		return true
	}

	return false
}

// SetKvmIpRange gets a reference to the given HyperflexIpAddrRange and assigns it to the KvmIpRange field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetKvmIpRange(v HyperflexIpAddrRange) {
	o.KvmIpRange = &v
}

// GetMacPrefixRange returns the MacPrefixRange field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicyAllOf) GetMacPrefixRange() HyperflexMacAddrPrefixRange {
	if o == nil || o.MacPrefixRange == nil {
		var ret HyperflexMacAddrPrefixRange
		return ret
	}
	return *o.MacPrefixRange
}

// GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool) {
	if o == nil || o.MacPrefixRange == nil {
		return nil, false
	}
	return o.MacPrefixRange, true
}

// HasMacPrefixRange returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasMacPrefixRange() bool {
	if o != nil && o.MacPrefixRange != nil {
		return true
	}

	return false
}

// SetMacPrefixRange gets a reference to the given HyperflexMacAddrPrefixRange and assigns it to the MacPrefixRange field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetMacPrefixRange(v HyperflexMacAddrPrefixRange) {
	o.MacPrefixRange = &v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicyAllOf) GetServerFirmwareVersion() string {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetServerFirmwareVersionOk() (*string, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given string and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetServerFirmwareVersion(v string) {
	o.ServerFirmwareVersion = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexUcsmConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KvmIpRange != nil {
		toSerialize["KvmIpRange"] = o.KvmIpRange
	}
	if o.MacPrefixRange != nil {
		toSerialize["MacPrefixRange"] = o.MacPrefixRange
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexUcsmConfigPolicyAllOf struct {
	value *HyperflexUcsmConfigPolicyAllOf
	isSet bool
}

func (v NullableHyperflexUcsmConfigPolicyAllOf) Get() *HyperflexUcsmConfigPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexUcsmConfigPolicyAllOf) Set(val *HyperflexUcsmConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexUcsmConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexUcsmConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexUcsmConfigPolicyAllOf(val *HyperflexUcsmConfigPolicyAllOf) *NullableHyperflexUcsmConfigPolicyAllOf {
	return &NullableHyperflexUcsmConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexUcsmConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexUcsmConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}