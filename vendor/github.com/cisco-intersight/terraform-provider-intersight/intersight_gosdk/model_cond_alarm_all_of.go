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

// CondAlarmAllOf Definition of the list of properties defined in 'cond.Alarm', excluding properties defined in parent classes.
type CondAlarmAllOf struct {
	// MoId of the affected object from the managed system's point of view.
	AffectedMoId *string `json:"AffectedMoId,omitempty"`
	// Managed system affected object type. For example Adaptor, FI, CIMC.
	AffectedMoType *string `json:"AffectedMoType,omitempty"`
	// A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject.
	AffectedObject *string `json:"AffectedObject,omitempty"`
	// Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.
	AncestorMoId *string `json:"AncestorMoId,omitempty"`
	// Parent MO type of the fault from managed system. For example, Blade for adaptor fault.
	AncestorMoType *string `json:"AncestorMoType,omitempty"`
	// A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.
	Code *string `json:"Code,omitempty"`
	// The time the alarm was created.
	CreationTime *time.Time `json:"CreationTime,omitempty"`
	// A longer description of the alarm than the name. The description contains details of the component reporting the issue.
	Description *string `json:"Description,omitempty"`
	// The time the alarm last had a change in severity.
	LastTransitionTime *time.Time `json:"LastTransitionTime,omitempty"`
	// A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.
	MsAffectedObject *string `json:"MsAffectedObject,omitempty"`
	// Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.
	Name *string `json:"Name,omitempty"`
	// The original severity when the alarm was first created.
	OrigSeverity *string `json:"OrigSeverity,omitempty"`
	// The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared.
	Severity             *string                              `json:"Severity,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmAllOf CondAlarmAllOf

// NewCondAlarmAllOf instantiates a new CondAlarmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmAllOf() *CondAlarmAllOf {
	this := CondAlarmAllOf{}
	var origSeverity string = "None"
	this.OrigSeverity = &origSeverity
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// NewCondAlarmAllOfWithDefaults instantiates a new CondAlarmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmAllOfWithDefaults() *CondAlarmAllOf {
	this := CondAlarmAllOf{}
	var origSeverity string = "None"
	this.OrigSeverity = &origSeverity
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// GetAffectedMoId returns the AffectedMoId field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAffectedMoId() string {
	if o == nil || o.AffectedMoId == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoId
}

// GetAffectedMoIdOk returns a tuple with the AffectedMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAffectedMoIdOk() (*string, bool) {
	if o == nil || o.AffectedMoId == nil {
		return nil, false
	}
	return o.AffectedMoId, true
}

// HasAffectedMoId returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedMoId() bool {
	if o != nil && o.AffectedMoId != nil {
		return true
	}

	return false
}

// SetAffectedMoId gets a reference to the given string and assigns it to the AffectedMoId field.
func (o *CondAlarmAllOf) SetAffectedMoId(v string) {
	o.AffectedMoId = &v
}

// GetAffectedMoType returns the AffectedMoType field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAffectedMoType() string {
	if o == nil || o.AffectedMoType == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoType
}

// GetAffectedMoTypeOk returns a tuple with the AffectedMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAffectedMoTypeOk() (*string, bool) {
	if o == nil || o.AffectedMoType == nil {
		return nil, false
	}
	return o.AffectedMoType, true
}

// HasAffectedMoType returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedMoType() bool {
	if o != nil && o.AffectedMoType != nil {
		return true
	}

	return false
}

// SetAffectedMoType gets a reference to the given string and assigns it to the AffectedMoType field.
func (o *CondAlarmAllOf) SetAffectedMoType(v string) {
	o.AffectedMoType = &v
}

// GetAffectedObject returns the AffectedObject field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAffectedObject() string {
	if o == nil || o.AffectedObject == nil {
		var ret string
		return ret
	}
	return *o.AffectedObject
}

// GetAffectedObjectOk returns a tuple with the AffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAffectedObjectOk() (*string, bool) {
	if o == nil || o.AffectedObject == nil {
		return nil, false
	}
	return o.AffectedObject, true
}

// HasAffectedObject returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedObject() bool {
	if o != nil && o.AffectedObject != nil {
		return true
	}

	return false
}

// SetAffectedObject gets a reference to the given string and assigns it to the AffectedObject field.
func (o *CondAlarmAllOf) SetAffectedObject(v string) {
	o.AffectedObject = &v
}

// GetAncestorMoId returns the AncestorMoId field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAncestorMoId() string {
	if o == nil || o.AncestorMoId == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoId
}

// GetAncestorMoIdOk returns a tuple with the AncestorMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAncestorMoIdOk() (*string, bool) {
	if o == nil || o.AncestorMoId == nil {
		return nil, false
	}
	return o.AncestorMoId, true
}

// HasAncestorMoId returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAncestorMoId() bool {
	if o != nil && o.AncestorMoId != nil {
		return true
	}

	return false
}

// SetAncestorMoId gets a reference to the given string and assigns it to the AncestorMoId field.
func (o *CondAlarmAllOf) SetAncestorMoId(v string) {
	o.AncestorMoId = &v
}

// GetAncestorMoType returns the AncestorMoType field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAncestorMoType() string {
	if o == nil || o.AncestorMoType == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoType
}

// GetAncestorMoTypeOk returns a tuple with the AncestorMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAncestorMoTypeOk() (*string, bool) {
	if o == nil || o.AncestorMoType == nil {
		return nil, false
	}
	return o.AncestorMoType, true
}

// HasAncestorMoType returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAncestorMoType() bool {
	if o != nil && o.AncestorMoType != nil {
		return true
	}

	return false
}

// SetAncestorMoType gets a reference to the given string and assigns it to the AncestorMoType field.
func (o *CondAlarmAllOf) SetAncestorMoType(v string) {
	o.AncestorMoType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CondAlarmAllOf) SetCode(v string) {
	o.Code = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CondAlarmAllOf) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CondAlarmAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *CondAlarmAllOf) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetMsAffectedObject returns the MsAffectedObject field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetMsAffectedObject() string {
	if o == nil || o.MsAffectedObject == nil {
		var ret string
		return ret
	}
	return *o.MsAffectedObject
}

// GetMsAffectedObjectOk returns a tuple with the MsAffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetMsAffectedObjectOk() (*string, bool) {
	if o == nil || o.MsAffectedObject == nil {
		return nil, false
	}
	return o.MsAffectedObject, true
}

// HasMsAffectedObject returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasMsAffectedObject() bool {
	if o != nil && o.MsAffectedObject != nil {
		return true
	}

	return false
}

// SetMsAffectedObject gets a reference to the given string and assigns it to the MsAffectedObject field.
func (o *CondAlarmAllOf) SetMsAffectedObject(v string) {
	o.MsAffectedObject = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CondAlarmAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrigSeverity returns the OrigSeverity field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetOrigSeverity() string {
	if o == nil || o.OrigSeverity == nil {
		var ret string
		return ret
	}
	return *o.OrigSeverity
}

// GetOrigSeverityOk returns a tuple with the OrigSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetOrigSeverityOk() (*string, bool) {
	if o == nil || o.OrigSeverity == nil {
		return nil, false
	}
	return o.OrigSeverity, true
}

// HasOrigSeverity returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasOrigSeverity() bool {
	if o != nil && o.OrigSeverity != nil {
		return true
	}

	return false
}

// SetOrigSeverity gets a reference to the given string and assigns it to the OrigSeverity field.
func (o *CondAlarmAllOf) SetOrigSeverity(v string) {
	o.OrigSeverity = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CondAlarmAllOf) SetSeverity(v string) {
	o.Severity = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CondAlarmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CondAlarmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AffectedMoId != nil {
		toSerialize["AffectedMoId"] = o.AffectedMoId
	}
	if o.AffectedMoType != nil {
		toSerialize["AffectedMoType"] = o.AffectedMoType
	}
	if o.AffectedObject != nil {
		toSerialize["AffectedObject"] = o.AffectedObject
	}
	if o.AncestorMoId != nil {
		toSerialize["AncestorMoId"] = o.AncestorMoId
	}
	if o.AncestorMoType != nil {
		toSerialize["AncestorMoType"] = o.AncestorMoType
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.CreationTime != nil {
		toSerialize["CreationTime"] = o.CreationTime
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.LastTransitionTime != nil {
		toSerialize["LastTransitionTime"] = o.LastTransitionTime
	}
	if o.MsAffectedObject != nil {
		toSerialize["MsAffectedObject"] = o.MsAffectedObject
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OrigSeverity != nil {
		toSerialize["OrigSeverity"] = o.OrigSeverity
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondAlarmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCondAlarmAllOf := _CondAlarmAllOf{}

	if err = json.Unmarshal(bytes, &varCondAlarmAllOf); err == nil {
		*o = CondAlarmAllOf(varCondAlarmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AffectedMoId")
		delete(additionalProperties, "AffectedMoType")
		delete(additionalProperties, "AffectedObject")
		delete(additionalProperties, "AncestorMoId")
		delete(additionalProperties, "AncestorMoType")
		delete(additionalProperties, "Code")
		delete(additionalProperties, "CreationTime")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "LastTransitionTime")
		delete(additionalProperties, "MsAffectedObject")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OrigSeverity")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondAlarmAllOf struct {
	value *CondAlarmAllOf
	isSet bool
}

func (v NullableCondAlarmAllOf) Get() *CondAlarmAllOf {
	return v.value
}

func (v *NullableCondAlarmAllOf) Set(val *CondAlarmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmAllOf(val *CondAlarmAllOf) *NullableCondAlarmAllOf {
	return &NullableCondAlarmAllOf{value: val, isSet: true}
}

func (v NullableCondAlarmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
