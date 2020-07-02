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

// WorkflowWaitTaskPromptAllOf Definition of the list of properties defined in 'workflow.WaitTaskPrompt', excluding properties defined in parent classes.
type WorkflowWaitTaskPromptAllOf struct {
	// Description that give more details about what it means to pick this prompt option for the wait task.
	Description *string `json:"Description,omitempty"`
	// User friendly label for the prompt. This label will be shown to the user as one of available options for the wait task.
	Label *string `json:"Label,omitempty"`
	// Name for the wait prompt.
	Name *string `json:"Name,omitempty"`
	// Task status for the wait task when this prompt option is selected.
	TaskStatus *string `json:"TaskStatus,omitempty"`
}

// NewWorkflowWaitTaskPromptAllOf instantiates a new WorkflowWaitTaskPromptAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWaitTaskPromptAllOf() *WorkflowWaitTaskPromptAllOf {
	this := WorkflowWaitTaskPromptAllOf{}
	var taskStatus string = "Scheduled"
	this.TaskStatus = &taskStatus
	return &this
}

// NewWorkflowWaitTaskPromptAllOfWithDefaults instantiates a new WorkflowWaitTaskPromptAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWaitTaskPromptAllOfWithDefaults() *WorkflowWaitTaskPromptAllOf {
	this := WorkflowWaitTaskPromptAllOf{}
	var taskStatus string = "Scheduled"
	this.TaskStatus = &taskStatus
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWaitTaskPromptAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowWaitTaskPromptAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWaitTaskPromptAllOf) SetName(v string) {
	o.Name = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetTaskStatus() string {
	if o == nil || o.TaskStatus == nil {
		var ret string
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetTaskStatusOk() (*string, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasTaskStatus() bool {
	if o != nil && o.TaskStatus != nil {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given string and assigns it to the TaskStatus field.
func (o *WorkflowWaitTaskPromptAllOf) SetTaskStatus(v string) {
	o.TaskStatus = &v
}

func (o WorkflowWaitTaskPromptAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TaskStatus != nil {
		toSerialize["TaskStatus"] = o.TaskStatus
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowWaitTaskPromptAllOf struct {
	value *WorkflowWaitTaskPromptAllOf
	isSet bool
}

func (v NullableWorkflowWaitTaskPromptAllOf) Get() *WorkflowWaitTaskPromptAllOf {
	return v.value
}

func (v *NullableWorkflowWaitTaskPromptAllOf) Set(val *WorkflowWaitTaskPromptAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWaitTaskPromptAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWaitTaskPromptAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWaitTaskPromptAllOf(val *WorkflowWaitTaskPromptAllOf) *NullableWorkflowWaitTaskPromptAllOf {
	return &NullableWorkflowWaitTaskPromptAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWaitTaskPromptAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWaitTaskPromptAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
