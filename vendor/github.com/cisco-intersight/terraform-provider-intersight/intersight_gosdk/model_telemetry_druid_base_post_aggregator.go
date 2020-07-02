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

// TelemetryDruidBasePostAggregator Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires.
type TelemetryDruidBasePostAggregator struct {
	// The post-aggregator type.
	Type string `json:"type"`
}

// NewTelemetryDruidBasePostAggregator instantiates a new TelemetryDruidBasePostAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidBasePostAggregator(type_ string) *TelemetryDruidBasePostAggregator {
	this := TelemetryDruidBasePostAggregator{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidBasePostAggregatorWithDefaults instantiates a new TelemetryDruidBasePostAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidBasePostAggregatorWithDefaults() *TelemetryDruidBasePostAggregator {
	this := TelemetryDruidBasePostAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidBasePostAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidBasePostAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidBasePostAggregator) SetType(v string) {
	o.Type = v
}

func (o TelemetryDruidBasePostAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryDruidBasePostAggregator struct {
	value *TelemetryDruidBasePostAggregator
	isSet bool
}

func (v NullableTelemetryDruidBasePostAggregator) Get() *TelemetryDruidBasePostAggregator {
	return v.value
}

func (v *NullableTelemetryDruidBasePostAggregator) Set(val *TelemetryDruidBasePostAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidBasePostAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidBasePostAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidBasePostAggregator(val *TelemetryDruidBasePostAggregator) *NullableTelemetryDruidBasePostAggregator {
	return &NullableTelemetryDruidBasePostAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidBasePostAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidBasePostAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
