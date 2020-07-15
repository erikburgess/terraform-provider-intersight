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

// WorkflowTargetDataTypeAllOf Definition of the list of properties defined in 'workflow.TargetDataType', excluding properties defined in parent classes.
type WorkflowTargetDataTypeAllOf struct {
	Properties           *[]WorkflowTargetProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetDataTypeAllOf WorkflowTargetDataTypeAllOf

// NewWorkflowTargetDataTypeAllOf instantiates a new WorkflowTargetDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetDataTypeAllOf() *WorkflowTargetDataTypeAllOf {
	this := WorkflowTargetDataTypeAllOf{}
	return &this
}

// NewWorkflowTargetDataTypeAllOfWithDefaults instantiates a new WorkflowTargetDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetDataTypeAllOfWithDefaults() *WorkflowTargetDataTypeAllOf {
	this := WorkflowTargetDataTypeAllOf{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowTargetDataTypeAllOf) GetProperties() []WorkflowTargetProperty {
	if o == nil || o.Properties == nil {
		var ret []WorkflowTargetProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataTypeAllOf) GetPropertiesOk() (*[]WorkflowTargetProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowTargetDataTypeAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []WorkflowTargetProperty and assigns it to the Properties field.
func (o *WorkflowTargetDataTypeAllOf) SetProperties(v []WorkflowTargetProperty) {
	o.Properties = &v
}

func (o WorkflowTargetDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetDataTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTargetDataTypeAllOf := _WorkflowTargetDataTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTargetDataTypeAllOf); err == nil {
		*o = WorkflowTargetDataTypeAllOf(varWorkflowTargetDataTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTargetDataTypeAllOf struct {
	value *WorkflowTargetDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowTargetDataTypeAllOf) Get() *WorkflowTargetDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowTargetDataTypeAllOf) Set(val *WorkflowTargetDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetDataTypeAllOf(val *WorkflowTargetDataTypeAllOf) *NullableWorkflowTargetDataTypeAllOf {
	return &NullableWorkflowTargetDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTargetDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
