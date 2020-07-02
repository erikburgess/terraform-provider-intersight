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

// OsBaseInstallConfigAllOf Definition of the list of properties defined in 'os.BaseInstallConfig', excluding properties defined in parent classes.
type OsBaseInstallConfigAllOf struct {
	AdditionalParameters *[]OsPlaceHolder `json:"AdditionalParameters,omitempty"`
	Answers              *OsAnswers       `json:"Answers,omitempty"`
	// User provided description about the OS install configuration.
	Description *string `json:"Description,omitempty"`
	// The install method to be used for OS installation - vMedia, iPXE.  Only vMedia is supported as of now.
	InstallMethod             *string                      `json:"InstallMethod,omitempty"`
	OperatingSystemParameters *OsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
}

// NewOsBaseInstallConfigAllOf instantiates a new OsBaseInstallConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsBaseInstallConfigAllOf() *OsBaseInstallConfigAllOf {
	this := OsBaseInstallConfigAllOf{}
	var installMethod string = "vMedia"
	this.InstallMethod = &installMethod
	return &this
}

// NewOsBaseInstallConfigAllOfWithDefaults instantiates a new OsBaseInstallConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsBaseInstallConfigAllOfWithDefaults() *OsBaseInstallConfigAllOf {
	this := OsBaseInstallConfigAllOf{}
	var installMethod string = "vMedia"
	this.InstallMethod = &installMethod
	return &this
}

// GetAdditionalParameters returns the AdditionalParameters field value if set, zero value otherwise.
func (o *OsBaseInstallConfigAllOf) GetAdditionalParameters() []OsPlaceHolder {
	if o == nil || o.AdditionalParameters == nil {
		var ret []OsPlaceHolder
		return ret
	}
	return *o.AdditionalParameters
}

// GetAdditionalParametersOk returns a tuple with the AdditionalParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfigAllOf) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool) {
	if o == nil || o.AdditionalParameters == nil {
		return nil, false
	}
	return o.AdditionalParameters, true
}

// HasAdditionalParameters returns a boolean if a field has been set.
func (o *OsBaseInstallConfigAllOf) HasAdditionalParameters() bool {
	if o != nil && o.AdditionalParameters != nil {
		return true
	}

	return false
}

// SetAdditionalParameters gets a reference to the given []OsPlaceHolder and assigns it to the AdditionalParameters field.
func (o *OsBaseInstallConfigAllOf) SetAdditionalParameters(v []OsPlaceHolder) {
	o.AdditionalParameters = &v
}

// GetAnswers returns the Answers field value if set, zero value otherwise.
func (o *OsBaseInstallConfigAllOf) GetAnswers() OsAnswers {
	if o == nil || o.Answers == nil {
		var ret OsAnswers
		return ret
	}
	return *o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfigAllOf) GetAnswersOk() (*OsAnswers, bool) {
	if o == nil || o.Answers == nil {
		return nil, false
	}
	return o.Answers, true
}

// HasAnswers returns a boolean if a field has been set.
func (o *OsBaseInstallConfigAllOf) HasAnswers() bool {
	if o != nil && o.Answers != nil {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given OsAnswers and assigns it to the Answers field.
func (o *OsBaseInstallConfigAllOf) SetAnswers(v OsAnswers) {
	o.Answers = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OsBaseInstallConfigAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfigAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OsBaseInstallConfigAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OsBaseInstallConfigAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInstallMethod returns the InstallMethod field value if set, zero value otherwise.
func (o *OsBaseInstallConfigAllOf) GetInstallMethod() string {
	if o == nil || o.InstallMethod == nil {
		var ret string
		return ret
	}
	return *o.InstallMethod
}

// GetInstallMethodOk returns a tuple with the InstallMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfigAllOf) GetInstallMethodOk() (*string, bool) {
	if o == nil || o.InstallMethod == nil {
		return nil, false
	}
	return o.InstallMethod, true
}

// HasInstallMethod returns a boolean if a field has been set.
func (o *OsBaseInstallConfigAllOf) HasInstallMethod() bool {
	if o != nil && o.InstallMethod != nil {
		return true
	}

	return false
}

// SetInstallMethod gets a reference to the given string and assigns it to the InstallMethod field.
func (o *OsBaseInstallConfigAllOf) SetInstallMethod(v string) {
	o.InstallMethod = &v
}

// GetOperatingSystemParameters returns the OperatingSystemParameters field value if set, zero value otherwise.
func (o *OsBaseInstallConfigAllOf) GetOperatingSystemParameters() OsOperatingSystemParameters {
	if o == nil || o.OperatingSystemParameters == nil {
		var ret OsOperatingSystemParameters
		return ret
	}
	return *o.OperatingSystemParameters
}

// GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfigAllOf) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool) {
	if o == nil || o.OperatingSystemParameters == nil {
		return nil, false
	}
	return o.OperatingSystemParameters, true
}

// HasOperatingSystemParameters returns a boolean if a field has been set.
func (o *OsBaseInstallConfigAllOf) HasOperatingSystemParameters() bool {
	if o != nil && o.OperatingSystemParameters != nil {
		return true
	}

	return false
}

// SetOperatingSystemParameters gets a reference to the given OsOperatingSystemParameters and assigns it to the OperatingSystemParameters field.
func (o *OsBaseInstallConfigAllOf) SetOperatingSystemParameters(v OsOperatingSystemParameters) {
	o.OperatingSystemParameters = &v
}

func (o OsBaseInstallConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalParameters != nil {
		toSerialize["AdditionalParameters"] = o.AdditionalParameters
	}
	if o.Answers != nil {
		toSerialize["Answers"] = o.Answers
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InstallMethod != nil {
		toSerialize["InstallMethod"] = o.InstallMethod
	}
	if o.OperatingSystemParameters != nil {
		toSerialize["OperatingSystemParameters"] = o.OperatingSystemParameters
	}
	return json.Marshal(toSerialize)
}

type NullableOsBaseInstallConfigAllOf struct {
	value *OsBaseInstallConfigAllOf
	isSet bool
}

func (v NullableOsBaseInstallConfigAllOf) Get() *OsBaseInstallConfigAllOf {
	return v.value
}

func (v *NullableOsBaseInstallConfigAllOf) Set(val *OsBaseInstallConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsBaseInstallConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsBaseInstallConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsBaseInstallConfigAllOf(val *OsBaseInstallConfigAllOf) *NullableOsBaseInstallConfigAllOf {
	return &NullableOsBaseInstallConfigAllOf{value: val, isSet: true}
}

func (v NullableOsBaseInstallConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsBaseInstallConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
