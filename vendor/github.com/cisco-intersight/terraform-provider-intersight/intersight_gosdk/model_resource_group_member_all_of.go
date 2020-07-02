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

// ResourceGroupMemberAllOf Definition of the list of properties defined in 'resource.GroupMember', excluding properties defined in parent classes.
type ResourceGroupMemberAllOf struct {
	Group    *ResourceGroupRelationship `json:"Group,omitempty"`
	Resource *MoBaseMoRelationship      `json:"Resource,omitempty"`
}

// NewResourceGroupMemberAllOf instantiates a new ResourceGroupMemberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceGroupMemberAllOf() *ResourceGroupMemberAllOf {
	this := ResourceGroupMemberAllOf{}
	return &this
}

// NewResourceGroupMemberAllOfWithDefaults instantiates a new ResourceGroupMemberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceGroupMemberAllOfWithDefaults() *ResourceGroupMemberAllOf {
	this := ResourceGroupMemberAllOf{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ResourceGroupMemberAllOf) GetGroup() ResourceGroupRelationship {
	if o == nil || o.Group == nil {
		var ret ResourceGroupRelationship
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroupMemberAllOf) GetGroupOk() (*ResourceGroupRelationship, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ResourceGroupMemberAllOf) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given ResourceGroupRelationship and assigns it to the Group field.
func (o *ResourceGroupMemberAllOf) SetGroup(v ResourceGroupRelationship) {
	o.Group = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceGroupMemberAllOf) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroupMemberAllOf) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceGroupMemberAllOf) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *ResourceGroupMemberAllOf) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o ResourceGroupMemberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableResourceGroupMemberAllOf struct {
	value *ResourceGroupMemberAllOf
	isSet bool
}

func (v NullableResourceGroupMemberAllOf) Get() *ResourceGroupMemberAllOf {
	return v.value
}

func (v *NullableResourceGroupMemberAllOf) Set(val *ResourceGroupMemberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceGroupMemberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceGroupMemberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceGroupMemberAllOf(val *ResourceGroupMemberAllOf) *NullableResourceGroupMemberAllOf {
	return &NullableResourceGroupMemberAllOf{value: val, isSet: true}
}

func (v NullableResourceGroupMemberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceGroupMemberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
