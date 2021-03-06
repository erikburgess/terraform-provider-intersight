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

// IamApiKeyAllOf Definition of the list of properties defined in 'iam.ApiKey', excluding properties defined in parent classes.
type IamApiKeyAllOf struct {
	// The cryptographic hash algorithm to calculate the message digest.
	HashAlgorithm *string                `json:"HashAlgorithm,omitempty"`
	KeySpec       *PkixKeyGenerationSpec `json:"KeySpec,omitempty"`
	// Holds the private key for the API key.
	PrivateKey *string `json:"PrivateKey,omitempty"`
	// The purpose of the API Key.
	Purpose *string `json:"Purpose,omitempty"`
	// The signing algorithm used by the client to authenticate API requests to Intersight. The signing algorithm must be compatible with the key generation specification.
	SigningAlgorithm     *string                    `json:"SigningAlgorithm,omitempty"`
	Permission           *IamPermissionRelationship `json:"Permission,omitempty"`
	User                 *IamUserRelationship       `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamApiKeyAllOf IamApiKeyAllOf

// NewIamApiKeyAllOf instantiates a new IamApiKeyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamApiKeyAllOf() *IamApiKeyAllOf {
	this := IamApiKeyAllOf{}
	var hashAlgorithm string = "SHA256"
	this.HashAlgorithm = &hashAlgorithm
	var signingAlgorithm string = "RSASSA-PKCS1-v1_5"
	this.SigningAlgorithm = &signingAlgorithm
	return &this
}

// NewIamApiKeyAllOfWithDefaults instantiates a new IamApiKeyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamApiKeyAllOfWithDefaults() *IamApiKeyAllOf {
	this := IamApiKeyAllOf{}
	var hashAlgorithm string = "SHA256"
	this.HashAlgorithm = &hashAlgorithm
	var signingAlgorithm string = "RSASSA-PKCS1-v1_5"
	this.SigningAlgorithm = &signingAlgorithm
	return &this
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetHashAlgorithm() string {
	if o == nil || o.HashAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || o.HashAlgorithm == nil {
		return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm != nil {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *IamApiKeyAllOf) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetKeySpec() PkixKeyGenerationSpec {
	if o == nil || o.KeySpec == nil {
		var ret PkixKeyGenerationSpec
		return ret
	}
	return *o.KeySpec
}

// GetKeySpecOk returns a tuple with the KeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetKeySpecOk() (*PkixKeyGenerationSpec, bool) {
	if o == nil || o.KeySpec == nil {
		return nil, false
	}
	return o.KeySpec, true
}

// HasKeySpec returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasKeySpec() bool {
	if o != nil && o.KeySpec != nil {
		return true
	}

	return false
}

// SetKeySpec gets a reference to the given PkixKeyGenerationSpec and assigns it to the KeySpec field.
func (o *IamApiKeyAllOf) SetKeySpec(v PkixKeyGenerationSpec) {
	o.KeySpec = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *IamApiKeyAllOf) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *IamApiKeyAllOf) SetPurpose(v string) {
	o.Purpose = &v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetSigningAlgorithm() string {
	if o == nil || o.SigningAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil || o.SigningAlgorithm == nil {
		return nil, false
	}
	return o.SigningAlgorithm, true
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasSigningAlgorithm() bool {
	if o != nil && o.SigningAlgorithm != nil {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given string and assigns it to the SigningAlgorithm field.
func (o *IamApiKeyAllOf) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamApiKeyAllOf) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamApiKeyAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKeyAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamApiKeyAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamApiKeyAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamApiKeyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HashAlgorithm != nil {
		toSerialize["HashAlgorithm"] = o.HashAlgorithm
	}
	if o.KeySpec != nil {
		toSerialize["KeySpec"] = o.KeySpec
	}
	if o.PrivateKey != nil {
		toSerialize["PrivateKey"] = o.PrivateKey
	}
	if o.Purpose != nil {
		toSerialize["Purpose"] = o.Purpose
	}
	if o.SigningAlgorithm != nil {
		toSerialize["SigningAlgorithm"] = o.SigningAlgorithm
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamApiKeyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamApiKeyAllOf := _IamApiKeyAllOf{}

	if err = json.Unmarshal(bytes, &varIamApiKeyAllOf); err == nil {
		*o = IamApiKeyAllOf(varIamApiKeyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "HashAlgorithm")
		delete(additionalProperties, "KeySpec")
		delete(additionalProperties, "PrivateKey")
		delete(additionalProperties, "Purpose")
		delete(additionalProperties, "SigningAlgorithm")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamApiKeyAllOf struct {
	value *IamApiKeyAllOf
	isSet bool
}

func (v NullableIamApiKeyAllOf) Get() *IamApiKeyAllOf {
	return v.value
}

func (v *NullableIamApiKeyAllOf) Set(val *IamApiKeyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamApiKeyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamApiKeyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamApiKeyAllOf(val *IamApiKeyAllOf) *NullableIamApiKeyAllOf {
	return &NullableIamApiKeyAllOf{value: val, isSet: true}
}

func (v NullableIamApiKeyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamApiKeyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
