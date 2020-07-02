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

// IamEndPointUserPolicyAllOf Definition of the list of properties defined in 'iam.EndPointUserPolicy', excluding properties defined in parent classes.
type IamEndPointUserPolicyAllOf struct {
	PasswordProperties *IamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`
	// An array of relationships to iamEndPointUserRole resources.
	EndPointUserRoles []IamEndPointUserRoleRelationship     `json:"EndPointUserRoles,omitempty"`
	Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
}

// NewIamEndPointUserPolicyAllOf instantiates a new IamEndPointUserPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserPolicyAllOf() *IamEndPointUserPolicyAllOf {
	this := IamEndPointUserPolicyAllOf{}
	return &this
}

// NewIamEndPointUserPolicyAllOfWithDefaults instantiates a new IamEndPointUserPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserPolicyAllOfWithDefaults() *IamEndPointUserPolicyAllOf {
	this := IamEndPointUserPolicyAllOf{}
	return &this
}

// GetPasswordProperties returns the PasswordProperties field value if set, zero value otherwise.
func (o *IamEndPointUserPolicyAllOf) GetPasswordProperties() IamEndPointPasswordProperties {
	if o == nil || o.PasswordProperties == nil {
		var ret IamEndPointPasswordProperties
		return ret
	}
	return *o.PasswordProperties
}

// GetPasswordPropertiesOk returns a tuple with the PasswordProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyAllOf) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool) {
	if o == nil || o.PasswordProperties == nil {
		return nil, false
	}
	return o.PasswordProperties, true
}

// HasPasswordProperties returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyAllOf) HasPasswordProperties() bool {
	if o != nil && o.PasswordProperties != nil {
		return true
	}

	return false
}

// SetPasswordProperties gets a reference to the given IamEndPointPasswordProperties and assigns it to the PasswordProperties field.
func (o *IamEndPointUserPolicyAllOf) SetPasswordProperties(v IamEndPointPasswordProperties) {
	o.PasswordProperties = &v
}

// GetEndPointUserRoles returns the EndPointUserRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyAllOf) GetEndPointUserRoles() []IamEndPointUserRoleRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleRelationship
		return ret
	}
	return o.EndPointUserRoles
}

// GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyAllOf) GetEndPointUserRolesOk() (*[]IamEndPointUserRoleRelationship, bool) {
	if o == nil || o.EndPointUserRoles == nil {
		return nil, false
	}
	return &o.EndPointUserRoles, true
}

// HasEndPointUserRoles returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyAllOf) HasEndPointUserRoles() bool {
	if o != nil && o.EndPointUserRoles != nil {
		return true
	}

	return false
}

// SetEndPointUserRoles gets a reference to the given []IamEndPointUserRoleRelationship and assigns it to the EndPointUserRoles field.
func (o *IamEndPointUserPolicyAllOf) SetEndPointUserRoles(v []IamEndPointUserRoleRelationship) {
	o.EndPointUserRoles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IamEndPointUserPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IamEndPointUserPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *IamEndPointUserPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o IamEndPointUserPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordProperties != nil {
		toSerialize["PasswordProperties"] = o.PasswordProperties
	}
	if o.EndPointUserRoles != nil {
		toSerialize["EndPointUserRoles"] = o.EndPointUserRoles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullableIamEndPointUserPolicyAllOf struct {
	value *IamEndPointUserPolicyAllOf
	isSet bool
}

func (v NullableIamEndPointUserPolicyAllOf) Get() *IamEndPointUserPolicyAllOf {
	return v.value
}

func (v *NullableIamEndPointUserPolicyAllOf) Set(val *IamEndPointUserPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserPolicyAllOf(val *IamEndPointUserPolicyAllOf) *NullableIamEndPointUserPolicyAllOf {
	return &NullableIamEndPointUserPolicyAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointUserPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
