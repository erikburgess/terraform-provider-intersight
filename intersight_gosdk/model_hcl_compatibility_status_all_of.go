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

// HclCompatibilityStatusAllOf Definition of the list of properties defined in 'hcl.CompatibilityStatus', excluding properties defined in parent classes.
type HclCompatibilityStatusAllOf struct {
	ProfileList *[]HclHardwareCompatibilityProfile `json:"ProfileList,omitempty"`
	// Type of the request to be served.
	RequestType          *string `json:"RequestType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclCompatibilityStatusAllOf HclCompatibilityStatusAllOf

// NewHclCompatibilityStatusAllOf instantiates a new HclCompatibilityStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclCompatibilityStatusAllOf() *HclCompatibilityStatusAllOf {
	this := HclCompatibilityStatusAllOf{}
	var requestType string = "FillSupportedVersions"
	this.RequestType = &requestType
	return &this
}

// NewHclCompatibilityStatusAllOfWithDefaults instantiates a new HclCompatibilityStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclCompatibilityStatusAllOfWithDefaults() *HclCompatibilityStatusAllOf {
	this := HclCompatibilityStatusAllOf{}
	var requestType string = "FillSupportedVersions"
	this.RequestType = &requestType
	return &this
}

// GetProfileList returns the ProfileList field value if set, zero value otherwise.
func (o *HclCompatibilityStatusAllOf) GetProfileList() []HclHardwareCompatibilityProfile {
	if o == nil || o.ProfileList == nil {
		var ret []HclHardwareCompatibilityProfile
		return ret
	}
	return *o.ProfileList
}

// GetProfileListOk returns a tuple with the ProfileList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclCompatibilityStatusAllOf) GetProfileListOk() (*[]HclHardwareCompatibilityProfile, bool) {
	if o == nil || o.ProfileList == nil {
		return nil, false
	}
	return o.ProfileList, true
}

// HasProfileList returns a boolean if a field has been set.
func (o *HclCompatibilityStatusAllOf) HasProfileList() bool {
	if o != nil && o.ProfileList != nil {
		return true
	}

	return false
}

// SetProfileList gets a reference to the given []HclHardwareCompatibilityProfile and assigns it to the ProfileList field.
func (o *HclCompatibilityStatusAllOf) SetProfileList(v []HclHardwareCompatibilityProfile) {
	o.ProfileList = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *HclCompatibilityStatusAllOf) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclCompatibilityStatusAllOf) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *HclCompatibilityStatusAllOf) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *HclCompatibilityStatusAllOf) SetRequestType(v string) {
	o.RequestType = &v
}

func (o HclCompatibilityStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProfileList != nil {
		toSerialize["ProfileList"] = o.ProfileList
	}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclCompatibilityStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclCompatibilityStatusAllOf := _HclCompatibilityStatusAllOf{}

	if err = json.Unmarshal(bytes, &varHclCompatibilityStatusAllOf); err == nil {
		*o = HclCompatibilityStatusAllOf(varHclCompatibilityStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ProfileList")
		delete(additionalProperties, "RequestType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclCompatibilityStatusAllOf struct {
	value *HclCompatibilityStatusAllOf
	isSet bool
}

func (v NullableHclCompatibilityStatusAllOf) Get() *HclCompatibilityStatusAllOf {
	return v.value
}

func (v *NullableHclCompatibilityStatusAllOf) Set(val *HclCompatibilityStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclCompatibilityStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclCompatibilityStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclCompatibilityStatusAllOf(val *HclCompatibilityStatusAllOf) *NullableHclCompatibilityStatusAllOf {
	return &NullableHclCompatibilityStatusAllOf{value: val, isSet: true}
}

func (v NullableHclCompatibilityStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclCompatibilityStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
