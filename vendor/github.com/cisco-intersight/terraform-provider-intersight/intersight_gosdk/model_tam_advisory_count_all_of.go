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

// TamAdvisoryCountAllOf Definition of the list of properties defined in 'tam.AdvisoryCount', excluding properties defined in parent classes.
type TamAdvisoryCountAllOf struct {
	// Total number of advisories affecting the account.
	AdvisoryCount *int64                  `json:"AdvisoryCount,omitempty"`
	Account       *IamAccountRelationship `json:"Account,omitempty"`
}

// NewTamAdvisoryCountAllOf instantiates a new TamAdvisoryCountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryCountAllOf() *TamAdvisoryCountAllOf {
	this := TamAdvisoryCountAllOf{}
	return &this
}

// NewTamAdvisoryCountAllOfWithDefaults instantiates a new TamAdvisoryCountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryCountAllOfWithDefaults() *TamAdvisoryCountAllOf {
	this := TamAdvisoryCountAllOf{}
	return &this
}

// GetAdvisoryCount returns the AdvisoryCount field value if set, zero value otherwise.
func (o *TamAdvisoryCountAllOf) GetAdvisoryCount() int64 {
	if o == nil || o.AdvisoryCount == nil {
		var ret int64
		return ret
	}
	return *o.AdvisoryCount
}

// GetAdvisoryCountOk returns a tuple with the AdvisoryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCountAllOf) GetAdvisoryCountOk() (*int64, bool) {
	if o == nil || o.AdvisoryCount == nil {
		return nil, false
	}
	return o.AdvisoryCount, true
}

// HasAdvisoryCount returns a boolean if a field has been set.
func (o *TamAdvisoryCountAllOf) HasAdvisoryCount() bool {
	if o != nil && o.AdvisoryCount != nil {
		return true
	}

	return false
}

// SetAdvisoryCount gets a reference to the given int64 and assigns it to the AdvisoryCount field.
func (o *TamAdvisoryCountAllOf) SetAdvisoryCount(v int64) {
	o.AdvisoryCount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TamAdvisoryCountAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCountAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TamAdvisoryCountAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TamAdvisoryCountAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TamAdvisoryCountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdvisoryCount != nil {
		toSerialize["AdvisoryCount"] = o.AdvisoryCount
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableTamAdvisoryCountAllOf struct {
	value *TamAdvisoryCountAllOf
	isSet bool
}

func (v NullableTamAdvisoryCountAllOf) Get() *TamAdvisoryCountAllOf {
	return v.value
}

func (v *NullableTamAdvisoryCountAllOf) Set(val *TamAdvisoryCountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryCountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryCountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryCountAllOf(val *TamAdvisoryCountAllOf) *NullableTamAdvisoryCountAllOf {
	return &NullableTamAdvisoryCountAllOf{value: val, isSet: true}
}

func (v NullableTamAdvisoryCountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryCountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
