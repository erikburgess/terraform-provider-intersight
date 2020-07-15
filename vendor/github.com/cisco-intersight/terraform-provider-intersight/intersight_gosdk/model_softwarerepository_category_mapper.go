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
	"reflect"
	"strings"
)

// SoftwarerepositoryCategoryMapper Maps a Cisco software repository image category identifier to its applicable hardware models.
type SoftwarerepositoryCategoryMapper struct {
	CapabilityCapability
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// Cisco software repository image category identifier.
	MdfId *string `json:"MdfId,omitempty"`
	// The regex that all images of this category follow.
	RegexPattern         *string   `json:"RegexPattern,omitempty"`
	SupportedModels      *[]string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategoryMapper SoftwarerepositoryCategoryMapper

// NewSoftwarerepositoryCategoryMapper instantiates a new SoftwarerepositoryCategoryMapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapper() *SoftwarerepositoryCategoryMapper {
	this := SoftwarerepositoryCategoryMapper{}
	return &this
}

// NewSoftwarerepositoryCategoryMapperWithDefaults instantiates a new SoftwarerepositoryCategoryMapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperWithDefaults() *SoftwarerepositoryCategoryMapper {
	this := SoftwarerepositoryCategoryMapper{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapper) SetCategory(v string) {
	o.Category = &v
}

// GetMdfId returns the MdfId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetMdfId() string {
	if o == nil || o.MdfId == nil {
		var ret string
		return ret
	}
	return *o.MdfId
}

// GetMdfIdOk returns a tuple with the MdfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetMdfIdOk() (*string, bool) {
	if o == nil || o.MdfId == nil {
		return nil, false
	}
	return o.MdfId, true
}

// HasMdfId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasMdfId() bool {
	if o != nil && o.MdfId != nil {
		return true
	}

	return false
}

// SetMdfId gets a reference to the given string and assigns it to the MdfId field.
func (o *SoftwarerepositoryCategoryMapper) SetMdfId(v string) {
	o.MdfId = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetRegexPattern() string {
	if o == nil || o.RegexPattern == nil {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetRegexPatternOk() (*string, bool) {
	if o == nil || o.RegexPattern == nil {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasRegexPattern() bool {
	if o != nil && o.RegexPattern != nil {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapper) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapper) GetSupportedModels() []string {
	if o == nil || o.SupportedModels == nil {
		var ret []string
		return ret
	}
	return *o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapper) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapper) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapper) SetSupportedModels(v []string) {
	o.SupportedModels = &v
}

func (o SoftwarerepositoryCategoryMapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.MdfId != nil {
		toSerialize["MdfId"] = o.MdfId
	}
	if o.RegexPattern != nil {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCategoryMapper) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryCategoryMapperWithoutEmbeddedStruct struct {
		// The category of the model series.
		Category *string `json:"Category,omitempty"`
		// Cisco software repository image category identifier.
		MdfId *string `json:"MdfId,omitempty"`
		// The regex that all images of this category follow.
		RegexPattern    *string   `json:"RegexPattern,omitempty"`
		SupportedModels *[]string `json:"SupportedModels,omitempty"`
	}

	varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct := SoftwarerepositoryCategoryMapperWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryCategoryMapper := _SoftwarerepositoryCategoryMapper{}
		varSoftwarerepositoryCategoryMapper.Category = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.Category
		varSoftwarerepositoryCategoryMapper.MdfId = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.MdfId
		varSoftwarerepositoryCategoryMapper.RegexPattern = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.RegexPattern
		varSoftwarerepositoryCategoryMapper.SupportedModels = varSoftwarerepositoryCategoryMapperWithoutEmbeddedStruct.SupportedModels
		*o = SoftwarerepositoryCategoryMapper(varSoftwarerepositoryCategoryMapper)
	} else {
		return err
	}

	varSoftwarerepositoryCategoryMapper := _SoftwarerepositoryCategoryMapper{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCategoryMapper)
	if err == nil {
		o.CapabilityCapability = varSoftwarerepositoryCategoryMapper.CapabilityCapability
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Category")
		delete(additionalProperties, "MdfId")
		delete(additionalProperties, "RegexPattern")
		delete(additionalProperties, "SupportedModels")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableSoftwarerepositoryCategoryMapper struct {
	value *SoftwarerepositoryCategoryMapper
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapper) Get() *SoftwarerepositoryCategoryMapper {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapper) Set(val *SoftwarerepositoryCategoryMapper) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapper) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapper(val *SoftwarerepositoryCategoryMapper) *NullableSoftwarerepositoryCategoryMapper {
	return &NullableSoftwarerepositoryCategoryMapper{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
