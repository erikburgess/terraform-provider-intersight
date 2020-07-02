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

// WorkflowDecisionTaskAllOf Definition of the list of properties defined in 'workflow.DecisionTask', excluding properties defined in parent classes.
type WorkflowDecisionTaskAllOf struct {
	// The condition to evaluate for this decision task. The condition can be a workflow or task variable or an expression based on the input parameters. Example value for condition if its Workflow/task variable is -  \"${task1.output.var1} or ${workflow.input.var2}\" which evaluates to a value matching any of the decision case values. Example value for condition if its an expression is - \"if ( $.element ! = null && $.element > 0 ) 'true'; else 'false'; \" which evaluates to 'true' or 'false' and will match one of the decision case values. Here \"element\" is a variable in decisiontask's inputParameters JSON formatted map. You can also use javascript like functions indexOf, toUpperCase in the expression which will be evaluated by the expression evaluator.
	Condition     *string                 `json:"Condition,omitempty"`
	DecisionCases *[]WorkflowDecisionCase `json:"DecisionCases,omitempty"`
	// The default next Task to execute if the decision cannot be evaluated to any of the DecisionCases.
	DefaultTask *string `json:"DefaultTask,omitempty"`
	// JSON formatted map that defines the input given to the decision task. The inputs are used as variables in the condition property of decision task. The input variables can be static values like \"hello\" , \"24\", \"true\" OR previous task outputs/workflow inputs like \"${task2.output.var1}}\". The input variables are referrenced as $.inputVariableName in the condition property.
	InputParameters interface{} `json:"InputParameters,omitempty"`
}

// NewWorkflowDecisionTaskAllOf instantiates a new WorkflowDecisionTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDecisionTaskAllOf() *WorkflowDecisionTaskAllOf {
	this := WorkflowDecisionTaskAllOf{}
	return &this
}

// NewWorkflowDecisionTaskAllOfWithDefaults instantiates a new WorkflowDecisionTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDecisionTaskAllOfWithDefaults() *WorkflowDecisionTaskAllOf {
	this := WorkflowDecisionTaskAllOf{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WorkflowDecisionTaskAllOf) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WorkflowDecisionTaskAllOf) SetCondition(v string) {
	o.Condition = &v
}

// GetDecisionCases returns the DecisionCases field value if set, zero value otherwise.
func (o *WorkflowDecisionTaskAllOf) GetDecisionCases() []WorkflowDecisionCase {
	if o == nil || o.DecisionCases == nil {
		var ret []WorkflowDecisionCase
		return ret
	}
	return *o.DecisionCases
}

// GetDecisionCasesOk returns a tuple with the DecisionCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetDecisionCasesOk() (*[]WorkflowDecisionCase, bool) {
	if o == nil || o.DecisionCases == nil {
		return nil, false
	}
	return o.DecisionCases, true
}

// HasDecisionCases returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasDecisionCases() bool {
	if o != nil && o.DecisionCases != nil {
		return true
	}

	return false
}

// SetDecisionCases gets a reference to the given []WorkflowDecisionCase and assigns it to the DecisionCases field.
func (o *WorkflowDecisionTaskAllOf) SetDecisionCases(v []WorkflowDecisionCase) {
	o.DecisionCases = &v
}

// GetDefaultTask returns the DefaultTask field value if set, zero value otherwise.
func (o *WorkflowDecisionTaskAllOf) GetDefaultTask() string {
	if o == nil || o.DefaultTask == nil {
		var ret string
		return ret
	}
	return *o.DefaultTask
}

// GetDefaultTaskOk returns a tuple with the DefaultTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetDefaultTaskOk() (*string, bool) {
	if o == nil || o.DefaultTask == nil {
		return nil, false
	}
	return o.DefaultTask, true
}

// HasDefaultTask returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasDefaultTask() bool {
	if o != nil && o.DefaultTask != nil {
		return true
	}

	return false
}

// SetDefaultTask gets a reference to the given string and assigns it to the DefaultTask field.
func (o *WorkflowDecisionTaskAllOf) SetDefaultTask(v string) {
	o.DefaultTask = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDecisionTaskAllOf) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDecisionTaskAllOf) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowDecisionTaskAllOf) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

func (o WorkflowDecisionTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}
	if o.DecisionCases != nil {
		toSerialize["DecisionCases"] = o.DecisionCases
	}
	if o.DefaultTask != nil {
		toSerialize["DefaultTask"] = o.DefaultTask
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowDecisionTaskAllOf struct {
	value *WorkflowDecisionTaskAllOf
	isSet bool
}

func (v NullableWorkflowDecisionTaskAllOf) Get() *WorkflowDecisionTaskAllOf {
	return v.value
}

func (v *NullableWorkflowDecisionTaskAllOf) Set(val *WorkflowDecisionTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDecisionTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDecisionTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDecisionTaskAllOf(val *WorkflowDecisionTaskAllOf) *NullableWorkflowDecisionTaskAllOf {
	return &NullableWorkflowDecisionTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDecisionTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDecisionTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
