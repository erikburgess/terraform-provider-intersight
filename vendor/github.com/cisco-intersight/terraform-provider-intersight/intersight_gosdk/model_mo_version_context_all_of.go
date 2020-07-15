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
	"time"
)

// MoVersionContextAllOf Definition of the list of properties defined in 'mo.VersionContext', excluding properties defined in parent classes.
type MoVersionContextAllOf struct {
	InterestedMos *[]MoMoRef `json:"InterestedMos,omitempty"`
	RefMo         *MoMoRef   `json:"RefMo,omitempty"`
	// The time this versioned Managed Object was created.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	Version *string `json:"Version,omitempty"`
	// Specifies type of version. Currently the only supported value is \"Configured\" that is used to keep track of snapshots of policies and profiles that are intended to be configured to target endpoints.
	VersionType          *string `json:"VersionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoVersionContextAllOf MoVersionContextAllOf

// NewMoVersionContextAllOf instantiates a new MoVersionContextAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoVersionContextAllOf() *MoVersionContextAllOf {
	this := MoVersionContextAllOf{}
	var versionType string = "Modified"
	this.VersionType = &versionType
	return &this
}

// NewMoVersionContextAllOfWithDefaults instantiates a new MoVersionContextAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoVersionContextAllOfWithDefaults() *MoVersionContextAllOf {
	this := MoVersionContextAllOf{}
	var versionType string = "Modified"
	this.VersionType = &versionType
	return &this
}

// GetInterestedMos returns the InterestedMos field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetInterestedMos() []MoMoRef {
	if o == nil || o.InterestedMos == nil {
		var ret []MoMoRef
		return ret
	}
	return *o.InterestedMos
}

// GetInterestedMosOk returns a tuple with the InterestedMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetInterestedMosOk() (*[]MoMoRef, bool) {
	if o == nil || o.InterestedMos == nil {
		return nil, false
	}
	return o.InterestedMos, true
}

// HasInterestedMos returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasInterestedMos() bool {
	if o != nil && o.InterestedMos != nil {
		return true
	}

	return false
}

// SetInterestedMos gets a reference to the given []MoMoRef and assigns it to the InterestedMos field.
func (o *MoVersionContextAllOf) SetInterestedMos(v []MoMoRef) {
	o.InterestedMos = &v
}

// GetRefMo returns the RefMo field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetRefMo() MoMoRef {
	if o == nil || o.RefMo == nil {
		var ret MoMoRef
		return ret
	}
	return *o.RefMo
}

// GetRefMoOk returns a tuple with the RefMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetRefMoOk() (*MoMoRef, bool) {
	if o == nil || o.RefMo == nil {
		return nil, false
	}
	return o.RefMo, true
}

// HasRefMo returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasRefMo() bool {
	if o != nil && o.RefMo != nil {
		return true
	}

	return false
}

// SetRefMo gets a reference to the given MoMoRef and assigns it to the RefMo field.
func (o *MoVersionContextAllOf) SetRefMo(v MoMoRef) {
	o.RefMo = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MoVersionContextAllOf) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MoVersionContextAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *MoVersionContextAllOf) SetVersionType(v string) {
	o.VersionType = &v
}

func (o MoVersionContextAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InterestedMos != nil {
		toSerialize["InterestedMos"] = o.InterestedMos
	}
	if o.RefMo != nil {
		toSerialize["RefMo"] = o.RefMo
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.VersionType != nil {
		toSerialize["VersionType"] = o.VersionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoVersionContextAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMoVersionContextAllOf := _MoVersionContextAllOf{}

	if err = json.Unmarshal(bytes, &varMoVersionContextAllOf); err == nil {
		*o = MoVersionContextAllOf(varMoVersionContextAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "InterestedMos")
		delete(additionalProperties, "RefMo")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "VersionType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoVersionContextAllOf struct {
	value *MoVersionContextAllOf
	isSet bool
}

func (v NullableMoVersionContextAllOf) Get() *MoVersionContextAllOf {
	return v.value
}

func (v *NullableMoVersionContextAllOf) Set(val *MoVersionContextAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoVersionContextAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoVersionContextAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoVersionContextAllOf(val *MoVersionContextAllOf) *NullableMoVersionContextAllOf {
	return &NullableMoVersionContextAllOf{value: val, isSet: true}
}

func (v NullableMoVersionContextAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoVersionContextAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
