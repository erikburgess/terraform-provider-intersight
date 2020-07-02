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

// TelemetryDruidDefaultLimitSpecAllOf struct for TelemetryDruidDefaultLimitSpecAllOf
type TelemetryDruidDefaultLimitSpecAllOf struct {
	// How many rows to return. If not specified, all rows will be returned.
	Limit int32 `json:"limit"`
	// null
	Columns []TelemetryDruidOrderByColumnSpec `json:"columns"`
}

// NewTelemetryDruidDefaultLimitSpecAllOf instantiates a new TelemetryDruidDefaultLimitSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDefaultLimitSpecAllOf(limit int32, columns []TelemetryDruidOrderByColumnSpec) *TelemetryDruidDefaultLimitSpecAllOf {
	this := TelemetryDruidDefaultLimitSpecAllOf{}
	this.Limit = limit
	this.Columns = columns
	return &this
}

// NewTelemetryDruidDefaultLimitSpecAllOfWithDefaults instantiates a new TelemetryDruidDefaultLimitSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDefaultLimitSpecAllOfWithDefaults() *TelemetryDruidDefaultLimitSpecAllOf {
	this := TelemetryDruidDefaultLimitSpecAllOf{}
	return &this
}

// GetLimit returns the Limit field value
func (o *TelemetryDruidDefaultLimitSpecAllOf) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpecAllOf) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *TelemetryDruidDefaultLimitSpecAllOf) SetLimit(v int32) {
	o.Limit = v
}

// GetColumns returns the Columns field value
func (o *TelemetryDruidDefaultLimitSpecAllOf) GetColumns() []TelemetryDruidOrderByColumnSpec {
	if o == nil {
		var ret []TelemetryDruidOrderByColumnSpec
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDefaultLimitSpecAllOf) GetColumnsOk() (*[]TelemetryDruidOrderByColumnSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *TelemetryDruidDefaultLimitSpecAllOf) SetColumns(v []TelemetryDruidOrderByColumnSpec) {
	o.Columns = v
}

func (o TelemetryDruidDefaultLimitSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["columns"] = o.Columns
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryDruidDefaultLimitSpecAllOf struct {
	value *TelemetryDruidDefaultLimitSpecAllOf
	isSet bool
}

func (v NullableTelemetryDruidDefaultLimitSpecAllOf) Get() *TelemetryDruidDefaultLimitSpecAllOf {
	return v.value
}

func (v *NullableTelemetryDruidDefaultLimitSpecAllOf) Set(val *TelemetryDruidDefaultLimitSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDefaultLimitSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDefaultLimitSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDefaultLimitSpecAllOf(val *TelemetryDruidDefaultLimitSpecAllOf) *NullableTelemetryDruidDefaultLimitSpecAllOf {
	return &NullableTelemetryDruidDefaultLimitSpecAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidDefaultLimitSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDefaultLimitSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
