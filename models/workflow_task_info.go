// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowTaskInfo Workflow:Task Info
//
// Task instance which represents the run time instance of a task within a workflow.
//
// swagger:model workflowTaskInfo
type WorkflowTaskInfo struct {
	MoBaseMo

	// The task description and this is the description that was added when the task was included into the workflow.
	//
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// The time stamp when the task reached a final state.
	//
	// Read Only: true
	// Format: date-time
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	// Description of the reason why the task failed.
	//
	// Read Only: true
	FailureReason string `json:"FailureReason,omitempty"`

	// The input data that was sent to the task at the start of execution.
	//
	// Read Only: true
	Input interface{} `json:"Input,omitempty"`

	// The current running task instance Id.
	//
	// Read Only: true
	InstID string `json:"InstId,omitempty"`

	// Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.
	//
	// Read Only: true
	Internal *bool `json:"Internal,omitempty"`

	// User friendly short label to describe this task instance in the workflow.
	//
	// Read Only: true
	Label string `json:"Label,omitempty"`

	// Collection of Task execution messages with severity.
	//
	// Read Only: true
	Message []*WorkflowMessage `json:"Message"`

	// Task definition name which specifies the task type.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// The output data that was generated by this task.
	//
	// Read Only: true
	Output interface{} `json:"Output,omitempty"`

	// The task reference name to ensure its unique inside a workflow.
	//
	// Read Only: true
	RefName string `json:"RefName,omitempty"`

	// A counter for number of retries.
	//
	// Read Only: true
	RetryCount int64 `json:"RetryCount,omitempty"`

	// The time stamp when the task started execution.
	//
	// Read Only: true
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// The status of the task and this will specify if the task is running or has reached a final state.
	//
	Status string `json:"Status,omitempty"`

	// A collection of references to the [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) Managed Object.
	//
	// When this managed object is deleted, the referenced [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	SubWorkflowInfo *WorkflowWorkflowInfoRef `json:"SubWorkflowInfo,omitempty"`

	// The task definition that was used to launch this task execution instance.
	//
	TaskDefinition *WorkflowTaskDefinitionRef `json:"TaskDefinition,omitempty"`

	// The list keeps a list of retried task instances.
	//
	// Read Only: true
	TaskInstIDList []*WorkflowTaskRetryInfo `json:"TaskInstIdList"`

	// A collection of references to the [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) Managed Object.
	//
	// When this managed object is deleted, the referenced [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) MO unsets its reference to this deleted MO.
	//
	// Read Only: true
	WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowTaskInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		FailureReason string `json:"FailureReason,omitempty"`

		Input interface{} `json:"Input,omitempty"`

		InstID string `json:"InstId,omitempty"`

		Internal *bool `json:"Internal,omitempty"`

		Label string `json:"Label,omitempty"`

		Message []*WorkflowMessage `json:"Message"`

		Name string `json:"Name,omitempty"`

		Output interface{} `json:"Output,omitempty"`

		RefName string `json:"RefName,omitempty"`

		RetryCount int64 `json:"RetryCount,omitempty"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`

		SubWorkflowInfo *WorkflowWorkflowInfoRef `json:"SubWorkflowInfo,omitempty"`

		TaskDefinition *WorkflowTaskDefinitionRef `json:"TaskDefinition,omitempty"`

		TaskInstIDList []*WorkflowTaskRetryInfo `json:"TaskInstIdList"`

		WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.EndTime = dataAO1.EndTime

	m.FailureReason = dataAO1.FailureReason

	m.Input = dataAO1.Input

	m.InstID = dataAO1.InstID

	m.Internal = dataAO1.Internal

	m.Label = dataAO1.Label

	m.Message = dataAO1.Message

	m.Name = dataAO1.Name

	m.Output = dataAO1.Output

	m.RefName = dataAO1.RefName

	m.RetryCount = dataAO1.RetryCount

	m.StartTime = dataAO1.StartTime

	m.Status = dataAO1.Status

	m.SubWorkflowInfo = dataAO1.SubWorkflowInfo

	m.TaskDefinition = dataAO1.TaskDefinition

	m.TaskInstIDList = dataAO1.TaskInstIDList

	m.WorkflowInfo = dataAO1.WorkflowInfo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowTaskInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		FailureReason string `json:"FailureReason,omitempty"`

		Input interface{} `json:"Input,omitempty"`

		InstID string `json:"InstId,omitempty"`

		Internal *bool `json:"Internal,omitempty"`

		Label string `json:"Label,omitempty"`

		Message []*WorkflowMessage `json:"Message"`

		Name string `json:"Name,omitempty"`

		Output interface{} `json:"Output,omitempty"`

		RefName string `json:"RefName,omitempty"`

		RetryCount int64 `json:"RetryCount,omitempty"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`

		SubWorkflowInfo *WorkflowWorkflowInfoRef `json:"SubWorkflowInfo,omitempty"`

		TaskDefinition *WorkflowTaskDefinitionRef `json:"TaskDefinition,omitempty"`

		TaskInstIDList []*WorkflowTaskRetryInfo `json:"TaskInstIdList"`

		WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`
	}

	dataAO1.Description = m.Description

	dataAO1.EndTime = m.EndTime

	dataAO1.FailureReason = m.FailureReason

	dataAO1.Input = m.Input

	dataAO1.InstID = m.InstID

	dataAO1.Internal = m.Internal

	dataAO1.Label = m.Label

	dataAO1.Message = m.Message

	dataAO1.Name = m.Name

	dataAO1.Output = m.Output

	dataAO1.RefName = m.RefName

	dataAO1.RetryCount = m.RetryCount

	dataAO1.StartTime = m.StartTime

	dataAO1.Status = m.Status

	dataAO1.SubWorkflowInfo = m.SubWorkflowInfo

	dataAO1.TaskDefinition = m.TaskDefinition

	dataAO1.TaskInstIDList = m.TaskInstIDList

	dataAO1.WorkflowInfo = m.WorkflowInfo

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow task info
func (m *WorkflowTaskInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubWorkflowInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskInstIDList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowTaskInfo) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowTaskInfo) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	for i := 0; i < len(m.Message); i++ {
		if swag.IsZero(m.Message[i]) { // not required
			continue
		}

		if m.Message[i] != nil {
			if err := m.Message[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Message" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowTaskInfo) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowTaskInfo) validateSubWorkflowInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.SubWorkflowInfo) { // not required
		return nil
	}

	if m.SubWorkflowInfo != nil {
		if err := m.SubWorkflowInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubWorkflowInfo")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowTaskInfo) validateTaskDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskDefinition) { // not required
		return nil
	}

	if m.TaskDefinition != nil {
		if err := m.TaskDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaskDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowTaskInfo) validateTaskInstIDList(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskInstIDList) { // not required
		return nil
	}

	for i := 0; i < len(m.TaskInstIDList); i++ {
		if swag.IsZero(m.TaskInstIDList[i]) { // not required
			continue
		}

		if m.TaskInstIDList[i] != nil {
			if err := m.TaskInstIDList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TaskInstIdList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowTaskInfo) validateWorkflowInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowInfo) { // not required
		return nil
	}

	if m.WorkflowInfo != nil {
		if err := m.WorkflowInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkflowInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowTaskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowTaskInfo) UnmarshalBinary(b []byte) error {
	var res WorkflowTaskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
