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

// WorkflowWebAPI Web API
//
// This models a single Web API request within a batch of requests that get
// executed within a single workflow task.
//
// swagger:model workflowWebApi
type WorkflowWebAPI struct {
	WorkflowAPI

	// Collection of key value pairs to set in the request header.
	//
	//
	Headers interface{} `json:"Headers,omitempty"`

	// The HTTP method to be executed in the given URL (GET, POST, PUT, etc).
	// If the value is not specified, GET will be used as default.
	//
	//
	Method string `json:"Method,omitempty"`

	// The accepted web protocol values are http and https.
	//
	//
	Protocol string `json:"Protocol,omitempty"`

	// If the web API is to be executed in a remote device connected to the
	// Intersight through device connector, 'Endpoint' is expected as the value
	// whereas if the API is an Intersight API, 'Local' is expected as the value.
	//
	//
	// Enum: [Endpoint Local]
	TargetType *string `json:"TargetType,omitempty"`

	// The URL of the resource in the target to which the API request is made.
	//
	//
	URL string `json:"Url,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowWebAPI) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowAPI
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowAPI = aO0

	// AO1
	var dataAO1 struct {
		Headers interface{} `json:"Headers,omitempty"`

		Method string `json:"Method,omitempty"`

		Protocol string `json:"Protocol,omitempty"`

		TargetType *string `json:"TargetType,omitempty"`

		URL string `json:"Url,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Headers = dataAO1.Headers

	m.Method = dataAO1.Method

	m.Protocol = dataAO1.Protocol

	m.TargetType = dataAO1.TargetType

	m.URL = dataAO1.URL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowWebAPI) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowAPI)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Headers interface{} `json:"Headers,omitempty"`

		Method string `json:"Method,omitempty"`

		Protocol string `json:"Protocol,omitempty"`

		TargetType *string `json:"TargetType,omitempty"`

		URL string `json:"Url,omitempty"`
	}

	dataAO1.Headers = m.Headers

	dataAO1.Method = m.Method

	dataAO1.Protocol = m.Protocol

	dataAO1.TargetType = m.TargetType

	dataAO1.URL = m.URL

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow web Api
func (m *WorkflowWebAPI) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowAPI
	if err := m.WorkflowAPI.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workflowWebApiTypeTargetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Endpoint","Local"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWebApiTypeTargetTypePropEnum = append(workflowWebApiTypeTargetTypePropEnum, v)
	}
}

// property enum
func (m *WorkflowWebAPI) validateTargetTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWebApiTypeTargetTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWebAPI) validateTargetType(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetTypeEnum("TargetType", "body", *m.TargetType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowWebAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowWebAPI) UnmarshalBinary(b []byte) error {
	var res WorkflowWebAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
