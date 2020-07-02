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

// IamSsoSessionAttributes Session attributes required to maintain states of SP and IdP.
type IamSsoSessionAttributes struct {
	MoBaseComplexType
	// SAML SessionNotOnOrAfter attribute sent by IdP in the assertion. IdP uses this to control for how long SP session maybe. SP does not issue SLO if the session is not valid.
	IdpSessionExpiration *string `json:"IdpSessionExpiration,omitempty"`
	// SAML SessionIndex attribute sent by IdP in the assertion. This has to be sent back to IdP in LogoutRequest.
	IdpSessionIndex *string `json:"IdpSessionIndex,omitempty"`
}

// NewIamSsoSessionAttributes instantiates a new IamSsoSessionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSsoSessionAttributes() *IamSsoSessionAttributes {
	this := IamSsoSessionAttributes{}
	return &this
}

// NewIamSsoSessionAttributesWithDefaults instantiates a new IamSsoSessionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSsoSessionAttributesWithDefaults() *IamSsoSessionAttributes {
	this := IamSsoSessionAttributes{}
	return &this
}

// GetIdpSessionExpiration returns the IdpSessionExpiration field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetIdpSessionExpiration() string {
	if o == nil || o.IdpSessionExpiration == nil {
		var ret string
		return ret
	}
	return *o.IdpSessionExpiration
}

// GetIdpSessionExpirationOk returns a tuple with the IdpSessionExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetIdpSessionExpirationOk() (*string, bool) {
	if o == nil || o.IdpSessionExpiration == nil {
		return nil, false
	}
	return o.IdpSessionExpiration, true
}

// HasIdpSessionExpiration returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasIdpSessionExpiration() bool {
	if o != nil && o.IdpSessionExpiration != nil {
		return true
	}

	return false
}

// SetIdpSessionExpiration gets a reference to the given string and assigns it to the IdpSessionExpiration field.
func (o *IamSsoSessionAttributes) SetIdpSessionExpiration(v string) {
	o.IdpSessionExpiration = &v
}

// GetIdpSessionIndex returns the IdpSessionIndex field value if set, zero value otherwise.
func (o *IamSsoSessionAttributes) GetIdpSessionIndex() string {
	if o == nil || o.IdpSessionIndex == nil {
		var ret string
		return ret
	}
	return *o.IdpSessionIndex
}

// GetIdpSessionIndexOk returns a tuple with the IdpSessionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSsoSessionAttributes) GetIdpSessionIndexOk() (*string, bool) {
	if o == nil || o.IdpSessionIndex == nil {
		return nil, false
	}
	return o.IdpSessionIndex, true
}

// HasIdpSessionIndex returns a boolean if a field has been set.
func (o *IamSsoSessionAttributes) HasIdpSessionIndex() bool {
	if o != nil && o.IdpSessionIndex != nil {
		return true
	}

	return false
}

// SetIdpSessionIndex gets a reference to the given string and assigns it to the IdpSessionIndex field.
func (o *IamSsoSessionAttributes) SetIdpSessionIndex(v string) {
	o.IdpSessionIndex = &v
}

func (o IamSsoSessionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.IdpSessionExpiration != nil {
		toSerialize["IdpSessionExpiration"] = o.IdpSessionExpiration
	}
	if o.IdpSessionIndex != nil {
		toSerialize["IdpSessionIndex"] = o.IdpSessionIndex
	}
	return json.Marshal(toSerialize)
}

type NullableIamSsoSessionAttributes struct {
	value *IamSsoSessionAttributes
	isSet bool
}

func (v NullableIamSsoSessionAttributes) Get() *IamSsoSessionAttributes {
	return v.value
}

func (v *NullableIamSsoSessionAttributes) Set(val *IamSsoSessionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSsoSessionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSsoSessionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSsoSessionAttributes(val *IamSsoSessionAttributes) *NullableIamSsoSessionAttributes {
	return &NullableIamSsoSessionAttributes{value: val, isSet: true}
}

func (v NullableIamSsoSessionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSsoSessionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
