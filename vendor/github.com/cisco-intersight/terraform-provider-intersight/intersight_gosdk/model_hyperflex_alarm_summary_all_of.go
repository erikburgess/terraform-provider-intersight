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
)

// HyperflexAlarmSummaryAllOf Definition of the list of properties defined in 'hyperflex.AlarmSummary', excluding properties defined in parent classes.
type HyperflexAlarmSummaryAllOf struct {
	// The count of alarms that have severity type Critical.
	Critical *int64 `json:"Critical,omitempty"`
	// The count of alarms that have severity type Warning.
	Warning              *int64 `json:"Warning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAlarmSummaryAllOf HyperflexAlarmSummaryAllOf

// NewHyperflexAlarmSummaryAllOf instantiates a new HyperflexAlarmSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAlarmSummaryAllOf() *HyperflexAlarmSummaryAllOf {
	this := HyperflexAlarmSummaryAllOf{}
	return &this
}

// NewHyperflexAlarmSummaryAllOfWithDefaults instantiates a new HyperflexAlarmSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAlarmSummaryAllOfWithDefaults() *HyperflexAlarmSummaryAllOf {
	this := HyperflexAlarmSummaryAllOf{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *HyperflexAlarmSummaryAllOf) GetCritical() int64 {
	if o == nil || o.Critical == nil {
		var ret int64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmSummaryAllOf) GetCriticalOk() (*int64, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *HyperflexAlarmSummaryAllOf) HasCritical() bool {
	if o != nil && o.Critical != nil {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int64 and assigns it to the Critical field.
func (o *HyperflexAlarmSummaryAllOf) SetCritical(v int64) {
	o.Critical = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *HyperflexAlarmSummaryAllOf) GetWarning() int64 {
	if o == nil || o.Warning == nil {
		var ret int64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmSummaryAllOf) GetWarningOk() (*int64, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *HyperflexAlarmSummaryAllOf) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int64 and assigns it to the Warning field.
func (o *HyperflexAlarmSummaryAllOf) SetWarning(v int64) {
	o.Warning = &v
}

func (o HyperflexAlarmSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Critical != nil {
		toSerialize["Critical"] = o.Critical
	}
	if o.Warning != nil {
		toSerialize["Warning"] = o.Warning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAlarmSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexAlarmSummaryAllOf := _HyperflexAlarmSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexAlarmSummaryAllOf); err == nil {
		*o = HyperflexAlarmSummaryAllOf(varHyperflexAlarmSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Critical")
		delete(additionalProperties, "Warning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAlarmSummaryAllOf struct {
	value *HyperflexAlarmSummaryAllOf
	isSet bool
}

func (v NullableHyperflexAlarmSummaryAllOf) Get() *HyperflexAlarmSummaryAllOf {
	return v.value
}

func (v *NullableHyperflexAlarmSummaryAllOf) Set(val *HyperflexAlarmSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAlarmSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAlarmSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAlarmSummaryAllOf(val *HyperflexAlarmSummaryAllOf) *NullableHyperflexAlarmSummaryAllOf {
	return &NullableHyperflexAlarmSummaryAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAlarmSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAlarmSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
