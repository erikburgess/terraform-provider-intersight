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

// WorkflowCustomDataType This data type represents a custom data object.
type WorkflowCustomDataType struct {
	WorkflowBaseDataType
	Properties *WorkflowCustomDataProperty `json:"Properties,omitempty"`
}

// NewWorkflowCustomDataType instantiates a new WorkflowCustomDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataType() *WorkflowCustomDataType {
	this := WorkflowCustomDataType{}
	return &this
}

// NewWorkflowCustomDataTypeWithDefaults instantiates a new WorkflowCustomDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataTypeWithDefaults() *WorkflowCustomDataType {
	this := WorkflowCustomDataType{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowCustomDataType) GetProperties() WorkflowCustomDataProperty {
	if o == nil || o.Properties == nil {
		var ret WorkflowCustomDataProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataType) GetPropertiesOk() (*WorkflowCustomDataProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowCustomDataType) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given WorkflowCustomDataProperty and assigns it to the Properties field.
func (o *WorkflowCustomDataType) SetProperties(v WorkflowCustomDataProperty) {
	o.Properties = &v
}

func (o WorkflowCustomDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseDataType, errWorkflowBaseDataType := json.Marshal(o.WorkflowBaseDataType)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	errWorkflowBaseDataType = json.Unmarshal([]byte(serializedWorkflowBaseDataType), &toSerialize)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowCustomDataType struct {
	value *WorkflowCustomDataType
	isSet bool
}

func (v NullableWorkflowCustomDataType) Get() *WorkflowCustomDataType {
	return v.value
}

func (v *NullableWorkflowCustomDataType) Set(val *WorkflowCustomDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataType(val *WorkflowCustomDataType) *NullableWorkflowCustomDataType {
	return &NullableWorkflowCustomDataType{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
