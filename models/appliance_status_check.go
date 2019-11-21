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

// ApplianceStatusCheck Appliance:Status Check
//
// Status check on an Intersight Appliance entity.
//
// swagger:model applianceStatusCheck
type ApplianceStatusCheck struct {
	ApplianceStatusCheckAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceStatusCheck) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ApplianceStatusCheckAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ApplianceStatusCheckAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceStatusCheck) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ApplianceStatusCheckAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance status check
func (m *ApplianceStatusCheck) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ApplianceStatusCheckAO0P0
	if err := m.ApplianceStatusCheckAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceStatusCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceStatusCheck) UnmarshalBinary(b []byte) error {
	var res ApplianceStatusCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplianceStatusCheckAO0P0 appliance status check a o0 p0
// swagger:model ApplianceStatusCheckAO0P0
type ApplianceStatusCheckAO0P0 struct {

	// Unique identifier of the status check.
	//
	Code string `json:"Code,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Result of this status check.
	//
	// Enum: [OK Warning Critical Info]
	Result *string `json:"Result,omitempty"`

	// appliance status check a o0 p0
	ApplianceStatusCheckAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *ApplianceStatusCheckAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Unique identifier of the status check.
		//
		Code string `json:"Code,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Result of this status check.
		//
		// Enum: [OK Warning Critical Info]
		Result *string `json:"Result,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ApplianceStatusCheckAO0P0

	rcv.Code = stage1.Code

	rcv.ObjectType = stage1.ObjectType

	rcv.Result = stage1.Result

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Code")

	delete(stage2, "ObjectType")

	delete(stage2, "Result")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.ApplianceStatusCheckAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m ApplianceStatusCheckAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Unique identifier of the status check.
		//
		Code string `json:"Code,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Result of this status check.
		//
		// Enum: [OK Warning Critical Info]
		Result *string `json:"Result,omitempty"`
	}

	stage1.Code = m.Code

	stage1.ObjectType = m.ObjectType

	stage1.Result = m.Result

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.ApplianceStatusCheckAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.ApplianceStatusCheckAO0P0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this appliance status check a o0 p0
func (m *ApplianceStatusCheckAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applianceStatusCheckAO0P0TypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","Warning","Critical","Info"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceStatusCheckAO0P0TypeResultPropEnum = append(applianceStatusCheckAO0P0TypeResultPropEnum, v)
	}
}

const (

	// ApplianceStatusCheckAO0P0ResultOK captures enum value "OK"
	ApplianceStatusCheckAO0P0ResultOK string = "OK"

	// ApplianceStatusCheckAO0P0ResultWarning captures enum value "Warning"
	ApplianceStatusCheckAO0P0ResultWarning string = "Warning"

	// ApplianceStatusCheckAO0P0ResultCritical captures enum value "Critical"
	ApplianceStatusCheckAO0P0ResultCritical string = "Critical"

	// ApplianceStatusCheckAO0P0ResultInfo captures enum value "Info"
	ApplianceStatusCheckAO0P0ResultInfo string = "Info"
)

// prop value enum
func (m *ApplianceStatusCheckAO0P0) validateResultEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceStatusCheckAO0P0TypeResultPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceStatusCheckAO0P0) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("Result", "body", *m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceStatusCheckAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceStatusCheckAO0P0) UnmarshalBinary(b []byte) error {
	var res ApplianceStatusCheckAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
