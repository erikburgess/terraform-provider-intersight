// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowWorkflowMeta Workflow:Workflow Meta
//
// Contains a workflow definition which is a sequence of tasks to execute.
//
// swagger:model workflowWorkflowMeta
type WorkflowWorkflowMeta struct {
	MoBaseMo

	// The description for the workflow.
	//
	Description string `json:"Description,omitempty"`

	// The workflow input parameters.
	//
	InputParameters []string `json:"InputParameters"`

	// The name given to the workflow.
	//
	Name string `json:"Name,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// The workflow output parameters.
	//
	OutputParameters interface{} `json:"OutputParameters,omitempty"`

	// The Conductor schema version that decides what attribute can be supported.
	//
	SchemaVersion int64 `json:"SchemaVersion,omitempty"`

	// The src is workflow owner service.
	//
	Src string `json:"Src,omitempty"`

	// The tasks contained inside of the workflow.
	//
	Tasks interface{} `json:"Tasks,omitempty"`

	// The type of workflow definition.
	//
	// Enum: [SystemDefined UserDefined Dynamic]
	Type *string `json:"Type,omitempty"`

	// The version for the workflow so we can support multiple versions for the same workflow name.
	//
	Version int64 `json:"Version,omitempty"`

	// Parameter decides if workflows will wait for a duplicate to finish before starting a new one.
	//
	WaitOnDuplicate *bool `json:"WaitOnDuplicate,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowWorkflowMeta) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		InputParameters []string `json:"InputParameters"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		OutputParameters interface{} `json:"OutputParameters,omitempty"`

		SchemaVersion int64 `json:"SchemaVersion,omitempty"`

		Src string `json:"Src,omitempty"`

		Tasks interface{} `json:"Tasks,omitempty"`

		Type *string `json:"Type,omitempty"`

		Version int64 `json:"Version,omitempty"`

		WaitOnDuplicate *bool `json:"WaitOnDuplicate,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.InputParameters = dataAO1.InputParameters

	m.Name = dataAO1.Name

	m.Organization = dataAO1.Organization

	m.OutputParameters = dataAO1.OutputParameters

	m.SchemaVersion = dataAO1.SchemaVersion

	m.Src = dataAO1.Src

	m.Tasks = dataAO1.Tasks

	m.Type = dataAO1.Type

	m.Version = dataAO1.Version

	m.WaitOnDuplicate = dataAO1.WaitOnDuplicate

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowWorkflowMeta) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"Description,omitempty"`

		InputParameters []string `json:"InputParameters"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		OutputParameters interface{} `json:"OutputParameters,omitempty"`

		SchemaVersion int64 `json:"SchemaVersion,omitempty"`

		Src string `json:"Src,omitempty"`

		Tasks interface{} `json:"Tasks,omitempty"`

		Type *string `json:"Type,omitempty"`

		Version int64 `json:"Version,omitempty"`

		WaitOnDuplicate *bool `json:"WaitOnDuplicate,omitempty"`
	}

	dataAO1.Description = m.Description

	dataAO1.InputParameters = m.InputParameters

	dataAO1.Name = m.Name

	dataAO1.Organization = m.Organization

	dataAO1.OutputParameters = m.OutputParameters

	dataAO1.SchemaVersion = m.SchemaVersion

	dataAO1.Src = m.Src

	dataAO1.Tasks = m.Tasks

	dataAO1.Type = m.Type

	dataAO1.Version = m.Version

	dataAO1.WaitOnDuplicate = m.WaitOnDuplicate

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow workflow meta
func (m *WorkflowWorkflowMeta) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowWorkflowMeta) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

var workflowWorkflowMetaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SystemDefined","UserDefined","Dynamic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWorkflowMetaTypeTypePropEnum = append(workflowWorkflowMetaTypeTypePropEnum, v)
	}
}

// property enum
func (m *WorkflowWorkflowMeta) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWorkflowMetaTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWorkflowMeta) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowWorkflowMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowWorkflowMeta) UnmarshalBinary(b []byte) error {
	var res WorkflowWorkflowMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
