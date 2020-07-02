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

// IamRole A role is a collection of privilege sets that are assigned to a user using a permission object.
type IamRole struct {
	MoBaseMo
	// Informative description about each role.
	Description *string `json:"Description,omitempty"`
	// The name of the role which has to be granted to user.
	Name           *string                 `json:"Name,omitempty"`
	PrivilegeNames *[]string               `json:"PrivilegeNames,omitempty"`
	Account        *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
	System        *IamSystemRelationship        `json:"System,omitempty"`
}

// NewIamRole instantiates a new IamRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRole() *IamRole {
	this := IamRole{}
	return &this
}

// NewIamRoleWithDefaults instantiates a new IamRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRoleWithDefaults() *IamRole {
	this := IamRole{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRole) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamRole) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamRole) SetName(v string) {
	o.Name = &v
}

// GetPrivilegeNames returns the PrivilegeNames field value if set, zero value otherwise.
func (o *IamRole) GetPrivilegeNames() []string {
	if o == nil || o.PrivilegeNames == nil {
		var ret []string
		return ret
	}
	return *o.PrivilegeNames
}

// GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRole) GetPrivilegeNamesOk() (*[]string, bool) {
	if o == nil || o.PrivilegeNames == nil {
		return nil, false
	}
	return o.PrivilegeNames, true
}

// HasPrivilegeNames returns a boolean if a field has been set.
func (o *IamRole) HasPrivilegeNames() bool {
	if o != nil && o.PrivilegeNames != nil {
		return true
	}

	return false
}

// SetPrivilegeNames gets a reference to the given []string and assigns it to the PrivilegeNames field.
func (o *IamRole) SetPrivilegeNames(v []string) {
	o.PrivilegeNames = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamRole) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRole) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamRole) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamRole) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamRole) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamRole) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.PrivilegeSets == nil {
		return nil, false
	}
	return &o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamRole) HasPrivilegeSets() bool {
	if o != nil && o.PrivilegeSets != nil {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamRole) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamRole) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRole) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamRole) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamRole) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PrivilegeNames != nil {
		toSerialize["PrivilegeNames"] = o.PrivilegeNames
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}
	return json.Marshal(toSerialize)
}

type NullableIamRole struct {
	value *IamRole
	isSet bool
}

func (v NullableIamRole) Get() *IamRole {
	return v.value
}

func (v *NullableIamRole) Set(val *IamRole) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRole) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRole(val *IamRole) *NullableIamRole {
	return &NullableIamRole{value: val, isSet: true}
}

func (v NullableIamRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
