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

// ContentGrammarAllOf Definition of the list of properties defined in 'content.Grammar', excluding properties defined in parent classes.
type ContentGrammarAllOf struct {
	ErrorParameters      *[]ContentBaseParameter `json:"ErrorParameters,omitempty"`
	Parameters           *[]ContentBaseParameter `json:"Parameters,omitempty"`
	Types                *[]ContentComplexType   `json:"Types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentGrammarAllOf ContentGrammarAllOf

// NewContentGrammarAllOf instantiates a new ContentGrammarAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentGrammarAllOf() *ContentGrammarAllOf {
	this := ContentGrammarAllOf{}
	return &this
}

// NewContentGrammarAllOfWithDefaults instantiates a new ContentGrammarAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentGrammarAllOfWithDefaults() *ContentGrammarAllOf {
	this := ContentGrammarAllOf{}
	return &this
}

// GetErrorParameters returns the ErrorParameters field value if set, zero value otherwise.
func (o *ContentGrammarAllOf) GetErrorParameters() []ContentBaseParameter {
	if o == nil || o.ErrorParameters == nil {
		var ret []ContentBaseParameter
		return ret
	}
	return *o.ErrorParameters
}

// GetErrorParametersOk returns a tuple with the ErrorParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentGrammarAllOf) GetErrorParametersOk() (*[]ContentBaseParameter, bool) {
	if o == nil || o.ErrorParameters == nil {
		return nil, false
	}
	return o.ErrorParameters, true
}

// HasErrorParameters returns a boolean if a field has been set.
func (o *ContentGrammarAllOf) HasErrorParameters() bool {
	if o != nil && o.ErrorParameters != nil {
		return true
	}

	return false
}

// SetErrorParameters gets a reference to the given []ContentBaseParameter and assigns it to the ErrorParameters field.
func (o *ContentGrammarAllOf) SetErrorParameters(v []ContentBaseParameter) {
	o.ErrorParameters = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ContentGrammarAllOf) GetParameters() []ContentBaseParameter {
	if o == nil || o.Parameters == nil {
		var ret []ContentBaseParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentGrammarAllOf) GetParametersOk() (*[]ContentBaseParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentGrammarAllOf) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ContentBaseParameter and assigns it to the Parameters field.
func (o *ContentGrammarAllOf) SetParameters(v []ContentBaseParameter) {
	o.Parameters = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *ContentGrammarAllOf) GetTypes() []ContentComplexType {
	if o == nil || o.Types == nil {
		var ret []ContentComplexType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentGrammarAllOf) GetTypesOk() (*[]ContentComplexType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ContentGrammarAllOf) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []ContentComplexType and assigns it to the Types field.
func (o *ContentGrammarAllOf) SetTypes(v []ContentComplexType) {
	o.Types = &v
}

func (o ContentGrammarAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorParameters != nil {
		toSerialize["ErrorParameters"] = o.ErrorParameters
	}
	if o.Parameters != nil {
		toSerialize["Parameters"] = o.Parameters
	}
	if o.Types != nil {
		toSerialize["Types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContentGrammarAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varContentGrammarAllOf := _ContentGrammarAllOf{}

	if err = json.Unmarshal(bytes, &varContentGrammarAllOf); err == nil {
		*o = ContentGrammarAllOf(varContentGrammarAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ErrorParameters")
		delete(additionalProperties, "Parameters")
		delete(additionalProperties, "Types")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContentGrammarAllOf struct {
	value *ContentGrammarAllOf
	isSet bool
}

func (v NullableContentGrammarAllOf) Get() *ContentGrammarAllOf {
	return v.value
}

func (v *NullableContentGrammarAllOf) Set(val *ContentGrammarAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentGrammarAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentGrammarAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentGrammarAllOf(val *ContentGrammarAllOf) *NullableContentGrammarAllOf {
	return &NullableContentGrammarAllOf{value: val, isSet: true}
}

func (v NullableContentGrammarAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentGrammarAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
