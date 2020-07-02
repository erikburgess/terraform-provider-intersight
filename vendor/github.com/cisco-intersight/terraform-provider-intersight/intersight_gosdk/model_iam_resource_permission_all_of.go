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

// IamResourcePermissionAllOf Definition of the list of properties defined in 'iam.ResourcePermission', excluding properties defined in parent classes.
type IamResourcePermissionAllOf struct {
	PermissionRoles *[]IamPermissionToRoles `json:"PermissionRoles,omitempty"`
	// Name of the service owning the resource.
	TargetApp *string                        `json:"TargetApp,omitempty"`
	Holder    *IamSecurityHolderRelationship `json:"Holder,omitempty"`
	Resource  *MoBaseMoRelationship          `json:"Resource,omitempty"`
}

// NewIamResourcePermissionAllOf instantiates a new IamResourcePermissionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamResourcePermissionAllOf() *IamResourcePermissionAllOf {
	this := IamResourcePermissionAllOf{}
	return &this
}

// NewIamResourcePermissionAllOfWithDefaults instantiates a new IamResourcePermissionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamResourcePermissionAllOfWithDefaults() *IamResourcePermissionAllOf {
	this := IamResourcePermissionAllOf{}
	return &this
}

// GetPermissionRoles returns the PermissionRoles field value if set, zero value otherwise.
func (o *IamResourcePermissionAllOf) GetPermissionRoles() []IamPermissionToRoles {
	if o == nil || o.PermissionRoles == nil {
		var ret []IamPermissionToRoles
		return ret
	}
	return *o.PermissionRoles
}

// GetPermissionRolesOk returns a tuple with the PermissionRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermissionAllOf) GetPermissionRolesOk() (*[]IamPermissionToRoles, bool) {
	if o == nil || o.PermissionRoles == nil {
		return nil, false
	}
	return o.PermissionRoles, true
}

// HasPermissionRoles returns a boolean if a field has been set.
func (o *IamResourcePermissionAllOf) HasPermissionRoles() bool {
	if o != nil && o.PermissionRoles != nil {
		return true
	}

	return false
}

// SetPermissionRoles gets a reference to the given []IamPermissionToRoles and assigns it to the PermissionRoles field.
func (o *IamResourcePermissionAllOf) SetPermissionRoles(v []IamPermissionToRoles) {
	o.PermissionRoles = &v
}

// GetTargetApp returns the TargetApp field value if set, zero value otherwise.
func (o *IamResourcePermissionAllOf) GetTargetApp() string {
	if o == nil || o.TargetApp == nil {
		var ret string
		return ret
	}
	return *o.TargetApp
}

// GetTargetAppOk returns a tuple with the TargetApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermissionAllOf) GetTargetAppOk() (*string, bool) {
	if o == nil || o.TargetApp == nil {
		return nil, false
	}
	return o.TargetApp, true
}

// HasTargetApp returns a boolean if a field has been set.
func (o *IamResourcePermissionAllOf) HasTargetApp() bool {
	if o != nil && o.TargetApp != nil {
		return true
	}

	return false
}

// SetTargetApp gets a reference to the given string and assigns it to the TargetApp field.
func (o *IamResourcePermissionAllOf) SetTargetApp(v string) {
	o.TargetApp = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *IamResourcePermissionAllOf) GetHolder() IamSecurityHolderRelationship {
	if o == nil || o.Holder == nil {
		var ret IamSecurityHolderRelationship
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermissionAllOf) GetHolderOk() (*IamSecurityHolderRelationship, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *IamResourcePermissionAllOf) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given IamSecurityHolderRelationship and assigns it to the Holder field.
func (o *IamResourcePermissionAllOf) SetHolder(v IamSecurityHolderRelationship) {
	o.Holder = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IamResourcePermissionAllOf) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourcePermissionAllOf) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IamResourcePermissionAllOf) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *IamResourcePermissionAllOf) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o IamResourcePermissionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionRoles != nil {
		toSerialize["PermissionRoles"] = o.PermissionRoles
	}
	if o.TargetApp != nil {
		toSerialize["TargetApp"] = o.TargetApp
	}
	if o.Holder != nil {
		toSerialize["Holder"] = o.Holder
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableIamResourcePermissionAllOf struct {
	value *IamResourcePermissionAllOf
	isSet bool
}

func (v NullableIamResourcePermissionAllOf) Get() *IamResourcePermissionAllOf {
	return v.value
}

func (v *NullableIamResourcePermissionAllOf) Set(val *IamResourcePermissionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourcePermissionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourcePermissionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourcePermissionAllOf(val *IamResourcePermissionAllOf) *NullableIamResourcePermissionAllOf {
	return &NullableIamResourcePermissionAllOf{value: val, isSet: true}
}

func (v NullableIamResourcePermissionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourcePermissionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
