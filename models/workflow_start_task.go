// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowStartTask Workflow:Start Task
//
// A StartTask is the starting point for a workflow.  There can only be one StartTask in a workflow.
//
// swagger:model workflowStartTask
type WorkflowStartTask struct {
	WorkflowControlTask

	// The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node.
	//
	NextTask string `json:"NextTask,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowStartTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowControlTask
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowControlTask = aO0

	// AO1
	var dataAO1 struct {
		NextTask string `json:"NextTask,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.NextTask = dataAO1.NextTask

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowStartTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowControlTask)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		NextTask string `json:"NextTask,omitempty"`
	}

	dataAO1.NextTask = m.NextTask

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow start task
func (m *WorkflowStartTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowControlTask
	if err := m.WorkflowControlTask.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowStartTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowStartTask) UnmarshalBinary(b []byte) error {
	var res WorkflowStartTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
