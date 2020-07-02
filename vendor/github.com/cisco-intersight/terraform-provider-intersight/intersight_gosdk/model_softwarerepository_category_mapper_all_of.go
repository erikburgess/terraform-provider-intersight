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

// SoftwarerepositoryCategoryMapperAllOf Definition of the list of properties defined in 'softwarerepository.CategoryMapper', excluding properties defined in parent classes.
type SoftwarerepositoryCategoryMapperAllOf struct {
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// The type of distributable image, example huu, scu, driver, os.
	FileType *string `json:"FileType,omitempty"`
	// Cisco software repository image category identifier.
	MdfId *string `json:"MdfId,omitempty"`
	// The regex that all images of this category follow.
	RegexPattern *string `json:"RegexPattern,omitempty"`
	// The image can be downloaded from cisco.com or external cloud store.
	Source          *string   `json:"Source,omitempty"`
	SupportedModels *[]string `json:"SupportedModels,omitempty"`
	// The software type id provided by cisco.com.
	SwId     *string   `json:"SwId,omitempty"`
	TagTypes *[]string `json:"TagTypes,omitempty"`
	// The version from which user can download images from amazon store, if source is external cloud store.
	Version *string `json:"Version,omitempty"`
}

// NewSoftwarerepositoryCategoryMapperAllOf instantiates a new SoftwarerepositoryCategoryMapperAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapperAllOf() *SoftwarerepositoryCategoryMapperAllOf {
	this := SoftwarerepositoryCategoryMapperAllOf{}
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// NewSoftwarerepositoryCategoryMapperAllOfWithDefaults instantiates a new SoftwarerepositoryCategoryMapperAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperAllOfWithDefaults() *SoftwarerepositoryCategoryMapperAllOf {
	this := SoftwarerepositoryCategoryMapperAllOf{}
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetFileType(v string) {
	o.FileType = &v
}

// GetMdfId returns the MdfId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetMdfId() string {
	if o == nil || o.MdfId == nil {
		var ret string
		return ret
	}
	return *o.MdfId
}

// GetMdfIdOk returns a tuple with the MdfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetMdfIdOk() (*string, bool) {
	if o == nil || o.MdfId == nil {
		return nil, false
	}
	return o.MdfId, true
}

// HasMdfId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasMdfId() bool {
	if o != nil && o.MdfId != nil {
		return true
	}

	return false
}

// SetMdfId gets a reference to the given string and assigns it to the MdfId field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetMdfId(v string) {
	o.MdfId = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetRegexPattern() string {
	if o == nil || o.RegexPattern == nil {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetRegexPatternOk() (*string, bool) {
	if o == nil || o.RegexPattern == nil {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasRegexPattern() bool {
	if o != nil && o.RegexPattern != nil {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetSource(v string) {
	o.Source = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSupportedModels() []string {
	if o == nil || o.SupportedModels == nil {
		var ret []string
		return ret
	}
	return *o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = &v
}

// GetSwId returns the SwId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSwId() string {
	if o == nil || o.SwId == nil {
		var ret string
		return ret
	}
	return *o.SwId
}

// GetSwIdOk returns a tuple with the SwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSwIdOk() (*string, bool) {
	if o == nil || o.SwId == nil {
		return nil, false
	}
	return o.SwId, true
}

// HasSwId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasSwId() bool {
	if o != nil && o.SwId != nil {
		return true
	}

	return false
}

// SetSwId gets a reference to the given string and assigns it to the SwId field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetSwId(v string) {
	o.SwId = &v
}

// GetTagTypes returns the TagTypes field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetTagTypes() []string {
	if o == nil || o.TagTypes == nil {
		var ret []string
		return ret
	}
	return *o.TagTypes
}

// GetTagTypesOk returns a tuple with the TagTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetTagTypesOk() (*[]string, bool) {
	if o == nil || o.TagTypes == nil {
		return nil, false
	}
	return o.TagTypes, true
}

// HasTagTypes returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasTagTypes() bool {
	if o != nil && o.TagTypes != nil {
		return true
	}

	return false
}

// SetTagTypes gets a reference to the given []string and assigns it to the TagTypes field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetTagTypes(v []string) {
	o.TagTypes = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o SoftwarerepositoryCategoryMapperAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.FileType != nil {
		toSerialize["FileType"] = o.FileType
	}
	if o.MdfId != nil {
		toSerialize["MdfId"] = o.MdfId
	}
	if o.RegexPattern != nil {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if o.SwId != nil {
		toSerialize["SwId"] = o.SwId
	}
	if o.TagTypes != nil {
		toSerialize["TagTypes"] = o.TagTypes
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwarerepositoryCategoryMapperAllOf struct {
	value *SoftwarerepositoryCategoryMapperAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapperAllOf) Get() *SoftwarerepositoryCategoryMapperAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapperAllOf) Set(val *SoftwarerepositoryCategoryMapperAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapperAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapperAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapperAllOf(val *SoftwarerepositoryCategoryMapperAllOf) *NullableSoftwarerepositoryCategoryMapperAllOf {
	return &NullableSoftwarerepositoryCategoryMapperAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapperAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapperAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
