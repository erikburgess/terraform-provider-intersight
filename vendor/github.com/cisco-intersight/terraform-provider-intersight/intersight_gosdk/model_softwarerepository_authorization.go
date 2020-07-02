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

// SoftwarerepositoryAuthorization User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user.
type SoftwarerepositoryAuthorization struct {
	MoBaseMo
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Indicates whether the value of the 'userId' property has been set.
	IsUserIdSet *bool `json:"IsUserIdSet,omitempty"`
	// The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	Password *string `json:"Password,omitempty"`
	// The external repository for which this authorization has been provided. The only supported repository today is cisco.com.
	RepositoryType *string `json:"RepositoryType,omitempty"`
	// The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	UserId  *string                 `json:"UserId,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
}

// NewSoftwarerepositoryAuthorization instantiates a new SoftwarerepositoryAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryAuthorization() *SoftwarerepositoryAuthorization {
	this := SoftwarerepositoryAuthorization{}
	var repositoryType string = "Cisco"
	this.RepositoryType = &repositoryType
	return &this
}

// NewSoftwarerepositoryAuthorizationWithDefaults instantiates a new SoftwarerepositoryAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryAuthorizationWithDefaults() *SoftwarerepositoryAuthorization {
	this := SoftwarerepositoryAuthorization{}
	var repositoryType string = "Cisco"
	this.RepositoryType = &repositoryType
	return &this
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryAuthorization) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetIsUserIdSet returns the IsUserIdSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetIsUserIdSet() bool {
	if o == nil || o.IsUserIdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsUserIdSet
}

// GetIsUserIdSetOk returns a tuple with the IsUserIdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetIsUserIdSetOk() (*bool, bool) {
	if o == nil || o.IsUserIdSet == nil {
		return nil, false
	}
	return o.IsUserIdSet, true
}

// HasIsUserIdSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasIsUserIdSet() bool {
	if o != nil && o.IsUserIdSet != nil {
		return true
	}

	return false
}

// SetIsUserIdSet gets a reference to the given bool and assigns it to the IsUserIdSet field.
func (o *SoftwarerepositoryAuthorization) SetIsUserIdSet(v bool) {
	o.IsUserIdSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryAuthorization) SetPassword(v string) {
	o.Password = &v
}

// GetRepositoryType returns the RepositoryType field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetRepositoryType() string {
	if o == nil || o.RepositoryType == nil {
		var ret string
		return ret
	}
	return *o.RepositoryType
}

// GetRepositoryTypeOk returns a tuple with the RepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetRepositoryTypeOk() (*string, bool) {
	if o == nil || o.RepositoryType == nil {
		return nil, false
	}
	return o.RepositoryType, true
}

// HasRepositoryType returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasRepositoryType() bool {
	if o != nil && o.RepositoryType != nil {
		return true
	}

	return false
}

// SetRepositoryType gets a reference to the given string and assigns it to the RepositoryType field.
func (o *SoftwarerepositoryAuthorization) SetRepositoryType(v string) {
	o.RepositoryType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SoftwarerepositoryAuthorization) SetUserId(v string) {
	o.UserId = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *SoftwarerepositoryAuthorization) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o SoftwarerepositoryAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.IsUserIdSet != nil {
		toSerialize["IsUserIdSet"] = o.IsUserIdSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.RepositoryType != nil {
		toSerialize["RepositoryType"] = o.RepositoryType
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwarerepositoryAuthorization struct {
	value *SoftwarerepositoryAuthorization
	isSet bool
}

func (v NullableSoftwarerepositoryAuthorization) Get() *SoftwarerepositoryAuthorization {
	return v.value
}

func (v *NullableSoftwarerepositoryAuthorization) Set(val *SoftwarerepositoryAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryAuthorization(val *SoftwarerepositoryAuthorization) *NullableSoftwarerepositoryAuthorization {
	return &NullableSoftwarerepositoryAuthorization{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
