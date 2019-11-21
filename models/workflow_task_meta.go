// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowTaskMeta Workflow:Task Meta
//
// This MO contains a task definition.
//
// swagger:model workflowTaskMeta
type WorkflowTaskMeta struct {
	MoBaseMo

	// A task execution type to indicate if it is a system task.
	//
	ActionType string `json:"ActionType,omitempty"`

	// A description of the task.
	//
	Description string `json:"Description,omitempty"`

	// Input keys for the task which specifies parameters the task can take in as inputs.
	//
	InputKeys []string `json:"InputKeys"`

	// Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.
	//
	Internal *bool `json:"Internal,omitempty"`

	// A task name that should be unique in Conductor DB.
	//
	Name string `json:"Name,omitempty"`

	// Output keys for the task which specifies parameters the task will output at the end of execution.
	//
	OutputKeys []string `json:"OutputKeys"`

	// The worker respnose timeout value.
	//
	ResponseTimeoutSec int64 `json:"ResponseTimeoutSec,omitempty"`

	// A number of reties for this task.
	//
	RetryCount int64 `json:"RetryCount,omitempty"`

	// The time on which the retry will be delayed.
	//
	RetryDelaySec int64 `json:"RetryDelaySec,omitempty"`

	// A logic which defines the way to handle retry (FIXED, EXPONENTIAL_BACKOFF).
	//
	RetryLogic string `json:"RetryLogic,omitempty"`

	// A service owns the task metadata.
	//
	Src string `json:"Src,omitempty"`

	// A policy which defines the way to handle timeout (RETRY, TIME_OUT_WF, ALERT_ONLY).
	//
	TimeoutPolicy string `json:"TimeoutPolicy,omitempty"`

	// A timeout value for the task in seconds.
	//
	TimeoutSec int64 `json:"TimeoutSec,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowTaskMeta) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ActionType string `json:"ActionType,omitempty"`

		Description string `json:"Description,omitempty"`

		InputKeys []string `json:"InputKeys"`

		Internal *bool `json:"Internal,omitempty"`

		Name string `json:"Name,omitempty"`

		OutputKeys []string `json:"OutputKeys"`

		ResponseTimeoutSec int64 `json:"ResponseTimeoutSec,omitempty"`

		RetryCount int64 `json:"RetryCount,omitempty"`

		RetryDelaySec int64 `json:"RetryDelaySec,omitempty"`

		RetryLogic string `json:"RetryLogic,omitempty"`

		Src string `json:"Src,omitempty"`

		TimeoutPolicy string `json:"TimeoutPolicy,omitempty"`

		TimeoutSec int64 `json:"TimeoutSec,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ActionType = dataAO1.ActionType

	m.Description = dataAO1.Description

	m.InputKeys = dataAO1.InputKeys

	m.Internal = dataAO1.Internal

	m.Name = dataAO1.Name

	m.OutputKeys = dataAO1.OutputKeys

	m.ResponseTimeoutSec = dataAO1.ResponseTimeoutSec

	m.RetryCount = dataAO1.RetryCount

	m.RetryDelaySec = dataAO1.RetryDelaySec

	m.RetryLogic = dataAO1.RetryLogic

	m.Src = dataAO1.Src

	m.TimeoutPolicy = dataAO1.TimeoutPolicy

	m.TimeoutSec = dataAO1.TimeoutSec

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowTaskMeta) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ActionType string `json:"ActionType,omitempty"`

		Description string `json:"Description,omitempty"`

		InputKeys []string `json:"InputKeys"`

		Internal *bool `json:"Internal,omitempty"`

		Name string `json:"Name,omitempty"`

		OutputKeys []string `json:"OutputKeys"`

		ResponseTimeoutSec int64 `json:"ResponseTimeoutSec,omitempty"`

		RetryCount int64 `json:"RetryCount,omitempty"`

		RetryDelaySec int64 `json:"RetryDelaySec,omitempty"`

		RetryLogic string `json:"RetryLogic,omitempty"`

		Src string `json:"Src,omitempty"`

		TimeoutPolicy string `json:"TimeoutPolicy,omitempty"`

		TimeoutSec int64 `json:"TimeoutSec,omitempty"`
	}

	dataAO1.ActionType = m.ActionType

	dataAO1.Description = m.Description

	dataAO1.InputKeys = m.InputKeys

	dataAO1.Internal = m.Internal

	dataAO1.Name = m.Name

	dataAO1.OutputKeys = m.OutputKeys

	dataAO1.ResponseTimeoutSec = m.ResponseTimeoutSec

	dataAO1.RetryCount = m.RetryCount

	dataAO1.RetryDelaySec = m.RetryDelaySec

	dataAO1.RetryLogic = m.RetryLogic

	dataAO1.Src = m.Src

	dataAO1.TimeoutPolicy = m.TimeoutPolicy

	dataAO1.TimeoutSec = m.TimeoutSec

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow task meta
func (m *WorkflowTaskMeta) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowTaskMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowTaskMeta) UnmarshalBinary(b []byte) error {
	var res WorkflowTaskMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
