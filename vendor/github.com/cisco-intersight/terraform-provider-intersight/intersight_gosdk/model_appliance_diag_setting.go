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

// ApplianceDiagSetting DiagSetting model is used for changing the password of the operating system's diagnostic user account. The diagnostic user account can be used to login to the Intersight Appliance virtual machine. The diagnostic user account is protected by two separate authentication mechanisms: user's password and Cisco CT-engine generated key. Only the Intersight Appliance's local account administrator has the privileges to use this REST API.
type ApplianceDiagSetting struct {
	MoBaseMo
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Status message of the password change operation.
	Message *string `json:"Message,omitempty"`
	// Password of the Intersight Appliance's OS diagnostic user account.
	Password *string                 `json:"Password,omitempty"`
	Account  *IamAccountRelationship `json:"Account,omitempty"`
}

// NewApplianceDiagSetting instantiates a new ApplianceDiagSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDiagSetting() *ApplianceDiagSetting {
	this := ApplianceDiagSetting{}
	return &this
}

// NewApplianceDiagSettingWithDefaults instantiates a new ApplianceDiagSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDiagSettingWithDefaults() *ApplianceDiagSetting {
	this := ApplianceDiagSetting{}
	return &this
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceDiagSetting) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceDiagSetting) SetMessage(v string) {
	o.Message = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceDiagSetting) SetPassword(v string) {
	o.Password = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceDiagSetting) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDiagSetting) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceDiagSetting) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceDiagSetting) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceDiagSetting) MarshalJSON() ([]byte, error) {
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
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceDiagSetting struct {
	value *ApplianceDiagSetting
	isSet bool
}

func (v NullableApplianceDiagSetting) Get() *ApplianceDiagSetting {
	return v.value
}

func (v *NullableApplianceDiagSetting) Set(val *ApplianceDiagSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDiagSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDiagSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDiagSetting(val *ApplianceDiagSetting) *NullableApplianceDiagSetting {
	return &NullableApplianceDiagSetting{value: val, isSet: true}
}

func (v NullableApplianceDiagSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDiagSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
