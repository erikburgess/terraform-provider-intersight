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

// StorageBaseSnapshotAllOf Definition of the list of properties defined in 'storage.BaseSnapshot', excluding properties defined in parent classes.
type StorageBaseSnapshotAllOf struct {
	// Exact date and time at which snapshot was created.
	CreatedTime *time.Time `json:"CreatedTime,omitempty"`
	// Name of the snapshot which represents point-in-time copy of volume.
	Name *string `json:"Name,omitempty"`
	// Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume.
	ProtectionGroupName *string `json:"ProtectionGroupName,omitempty"`
	// Snapshot size represented in bytes.
	Size *int64 `json:"Size,omitempty"`
	// Source object on which the snapshot is created. It is the name of the originating volume.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseSnapshotAllOf StorageBaseSnapshotAllOf

// NewStorageBaseSnapshotAllOf instantiates a new StorageBaseSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseSnapshotAllOf() *StorageBaseSnapshotAllOf {
	this := StorageBaseSnapshotAllOf{}
	return &this
}

// NewStorageBaseSnapshotAllOfWithDefaults instantiates a new StorageBaseSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseSnapshotAllOfWithDefaults() *StorageBaseSnapshotAllOf {
	this := StorageBaseSnapshotAllOf{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *StorageBaseSnapshotAllOf) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotAllOf) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *StorageBaseSnapshotAllOf) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *StorageBaseSnapshotAllOf) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseSnapshotAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseSnapshotAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseSnapshotAllOf) SetName(v string) {
	o.Name = &v
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise.
func (o *StorageBaseSnapshotAllOf) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotAllOf) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil || o.ProtectionGroupName == nil {
		return nil, false
	}
	return o.ProtectionGroupName, true
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *StorageBaseSnapshotAllOf) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName != nil {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given string and assigns it to the ProtectionGroupName field.
func (o *StorageBaseSnapshotAllOf) SetProtectionGroupName(v string) {
	o.ProtectionGroupName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseSnapshotAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseSnapshotAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseSnapshotAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageBaseSnapshotAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageBaseSnapshotAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StorageBaseSnapshotAllOf) SetSource(v string) {
	o.Source = &v
}

func (o StorageBaseSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedTime != nil {
		toSerialize["CreatedTime"] = o.CreatedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProtectionGroupName != nil {
		toSerialize["ProtectionGroupName"] = o.ProtectionGroupName
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseSnapshotAllOf := _StorageBaseSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseSnapshotAllOf); err == nil {
		*o = StorageBaseSnapshotAllOf(varStorageBaseSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CreatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProtectionGroupName")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseSnapshotAllOf struct {
	value *StorageBaseSnapshotAllOf
	isSet bool
}

func (v NullableStorageBaseSnapshotAllOf) Get() *StorageBaseSnapshotAllOf {
	return v.value
}

func (v *NullableStorageBaseSnapshotAllOf) Set(val *StorageBaseSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseSnapshotAllOf(val *StorageBaseSnapshotAllOf) *NullableStorageBaseSnapshotAllOf {
	return &NullableStorageBaseSnapshotAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
