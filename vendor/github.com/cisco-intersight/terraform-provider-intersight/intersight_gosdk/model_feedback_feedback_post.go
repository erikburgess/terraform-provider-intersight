/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FeedbackFeedbackPost Initial feedback submitted by the user from UI.
type FeedbackFeedbackPost struct {
	MoBaseMo
	FeedbackData         *FeedbackFeedbackData   `json:"FeedbackData,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeedbackFeedbackPost FeedbackFeedbackPost

// NewFeedbackFeedbackPost instantiates a new FeedbackFeedbackPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackFeedbackPost() *FeedbackFeedbackPost {
	this := FeedbackFeedbackPost{}
	return &this
}

// NewFeedbackFeedbackPostWithDefaults instantiates a new FeedbackFeedbackPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackFeedbackPostWithDefaults() *FeedbackFeedbackPost {
	this := FeedbackFeedbackPost{}
	return &this
}

// GetFeedbackData returns the FeedbackData field value if set, zero value otherwise.
func (o *FeedbackFeedbackPost) GetFeedbackData() FeedbackFeedbackData {
	if o == nil || o.FeedbackData == nil {
		var ret FeedbackFeedbackData
		return ret
	}
	return *o.FeedbackData
}

// GetFeedbackDataOk returns a tuple with the FeedbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPost) GetFeedbackDataOk() (*FeedbackFeedbackData, bool) {
	if o == nil || o.FeedbackData == nil {
		return nil, false
	}
	return o.FeedbackData, true
}

// HasFeedbackData returns a boolean if a field has been set.
func (o *FeedbackFeedbackPost) HasFeedbackData() bool {
	if o != nil && o.FeedbackData != nil {
		return true
	}

	return false
}

// SetFeedbackData gets a reference to the given FeedbackFeedbackData and assigns it to the FeedbackData field.
func (o *FeedbackFeedbackPost) SetFeedbackData(v FeedbackFeedbackData) {
	o.FeedbackData = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *FeedbackFeedbackPost) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPost) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *FeedbackFeedbackPost) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *FeedbackFeedbackPost) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o FeedbackFeedbackPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.FeedbackData != nil {
		toSerialize["FeedbackData"] = o.FeedbackData
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FeedbackFeedbackPost) UnmarshalJSON(bytes []byte) (err error) {
	type FeedbackFeedbackPostWithoutEmbeddedStruct struct {
		FeedbackData *FeedbackFeedbackData   `json:"FeedbackData,omitempty"`
		Account      *IamAccountRelationship `json:"Account,omitempty"`
	}

	varFeedbackFeedbackPostWithoutEmbeddedStruct := FeedbackFeedbackPostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFeedbackFeedbackPostWithoutEmbeddedStruct)
	if err == nil {
		varFeedbackFeedbackPost := _FeedbackFeedbackPost{}
		varFeedbackFeedbackPost.FeedbackData = varFeedbackFeedbackPostWithoutEmbeddedStruct.FeedbackData
		varFeedbackFeedbackPost.Account = varFeedbackFeedbackPostWithoutEmbeddedStruct.Account
		*o = FeedbackFeedbackPost(varFeedbackFeedbackPost)
	} else {
		return err
	}

	varFeedbackFeedbackPost := _FeedbackFeedbackPost{}

	err = json.Unmarshal(bytes, &varFeedbackFeedbackPost)
	if err == nil {
		o.MoBaseMo = varFeedbackFeedbackPost.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FeedbackData")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeedbackFeedbackPost struct {
	value *FeedbackFeedbackPost
	isSet bool
}

func (v NullableFeedbackFeedbackPost) Get() *FeedbackFeedbackPost {
	return v.value
}

func (v *NullableFeedbackFeedbackPost) Set(val *FeedbackFeedbackPost) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackFeedbackPost) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackFeedbackPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackFeedbackPost(val *FeedbackFeedbackPost) *NullableFeedbackFeedbackPost {
	return &NullableFeedbackFeedbackPost{value: val, isSet: true}
}

func (v NullableFeedbackFeedbackPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackFeedbackPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
