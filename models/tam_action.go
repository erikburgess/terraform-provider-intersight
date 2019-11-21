// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TamAction Tam:Action
//
// An action is used to react when an object satifies the condition specified in alert query. For e.g. an action in case of an object matching a fieldNotice query would be to create an alert instance of type 'fieldNotice' for the affected object.
//
// swagger:model tamAction
type TamAction struct {
	TamActionAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamAction) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TamActionAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TamActionAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamAction) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TamActionAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam action
func (m *TamAction) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TamActionAO0P0
	if err := m.TamActionAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TamAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamAction) UnmarshalBinary(b []byte) error {
	var res TamAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TamActionAO0P0 tam action a o0 p0
// swagger:model TamActionAO0P0
type TamActionAO0P0 struct {

	// Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).
	//
	AffectedObjectType string `json:"AffectedObjectType,omitempty"`

	// Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).
	//
	// Enum: [psirt fieldNotice]
	AlertType *string `json:"AlertType,omitempty"`

	// Identifiers represents the filter criteria (property names and values) used to identify an Intersight managed object of type specified in affectedObjectType property. An instance of an alert is then create on (or removed from) the identified managed object.
	//
	Identifiers []*TamIdentifiers `json:"Identifiers"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Operation type for the alert action. An action is used to carry out the process of "reacting" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.
	//
	// Enum: [create remove]
	OperationType *string `json:"OperationType,omitempty"`

	// Set of SparkSQL queries used determine if a given alert is applicable or not. Refer to https://spark.apache.org/sql/ for more details.
	//
	Queries []*TamQueryEntry `json:"Queries"`

	// Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.
	//
	// Enum: [restApi]
	Type *string `json:"Type,omitempty"`

	// tam action a o0 p0
	TamActionAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *TamActionAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).
		//
		AffectedObjectType string `json:"AffectedObjectType,omitempty"`

		// Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).
		//
		// Enum: [psirt fieldNotice]
		AlertType *string `json:"AlertType,omitempty"`

		// Identifiers represents the filter criteria (property names and values) used to identify an Intersight managed object of type specified in affectedObjectType property. An instance of an alert is then create on (or removed from) the identified managed object.
		//
		Identifiers []*TamIdentifiers `json:"Identifiers"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Operation type for the alert action. An action is used to carry out the process of "reacting" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.
		//
		// Enum: [create remove]
		OperationType *string `json:"OperationType,omitempty"`

		// Set of SparkSQL queries used determine if a given alert is applicable or not. Refer to https://spark.apache.org/sql/ for more details.
		//
		Queries []*TamQueryEntry `json:"Queries"`

		// Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.
		//
		// Enum: [restApi]
		Type *string `json:"Type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv TamActionAO0P0

	rcv.AffectedObjectType = stage1.AffectedObjectType

	rcv.AlertType = stage1.AlertType

	rcv.Identifiers = stage1.Identifiers

	rcv.ObjectType = stage1.ObjectType

	rcv.OperationType = stage1.OperationType

	rcv.Queries = stage1.Queries

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "AffectedObjectType")

	delete(stage2, "AlertType")

	delete(stage2, "Identifiers")

	delete(stage2, "ObjectType")

	delete(stage2, "OperationType")

	delete(stage2, "Queries")

	delete(stage2, "Type")

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
		m.TamActionAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m TamActionAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).
		//
		AffectedObjectType string `json:"AffectedObjectType,omitempty"`

		// Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).
		//
		// Enum: [psirt fieldNotice]
		AlertType *string `json:"AlertType,omitempty"`

		// Identifiers represents the filter criteria (property names and values) used to identify an Intersight managed object of type specified in affectedObjectType property. An instance of an alert is then create on (or removed from) the identified managed object.
		//
		Identifiers []*TamIdentifiers `json:"Identifiers"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Operation type for the alert action. An action is used to carry out the process of "reacting" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.
		//
		// Enum: [create remove]
		OperationType *string `json:"OperationType,omitempty"`

		// Set of SparkSQL queries used determine if a given alert is applicable or not. Refer to https://spark.apache.org/sql/ for more details.
		//
		Queries []*TamQueryEntry `json:"Queries"`

		// Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.
		//
		// Enum: [restApi]
		Type *string `json:"Type,omitempty"`
	}

	stage1.AffectedObjectType = m.AffectedObjectType

	stage1.AlertType = m.AlertType

	stage1.Identifiers = m.Identifiers

	stage1.ObjectType = m.ObjectType

	stage1.OperationType = m.OperationType

	stage1.Queries = m.Queries

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.TamActionAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.TamActionAO0P0)
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

// Validate validates this tam action a o0 p0
func (m *TamActionAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueries(formats); err != nil {
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

var tamActionAO0P0TypeAlertTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["psirt","fieldNotice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tamActionAO0P0TypeAlertTypePropEnum = append(tamActionAO0P0TypeAlertTypePropEnum, v)
	}
}

const (

	// TamActionAO0P0AlertTypePsirt captures enum value "psirt"
	TamActionAO0P0AlertTypePsirt string = "psirt"

	// TamActionAO0P0AlertTypeFieldNotice captures enum value "fieldNotice"
	TamActionAO0P0AlertTypeFieldNotice string = "fieldNotice"
)

// prop value enum
func (m *TamActionAO0P0) validateAlertTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tamActionAO0P0TypeAlertTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TamActionAO0P0) validateAlertType(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlertTypeEnum("AlertType", "body", *m.AlertType); err != nil {
		return err
	}

	return nil
}

func (m *TamActionAO0P0) validateIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {
		if swag.IsZero(m.Identifiers[i]) { // not required
			continue
		}

		if m.Identifiers[i] != nil {
			if err := m.Identifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Identifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var tamActionAO0P0TypeOperationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create","remove"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tamActionAO0P0TypeOperationTypePropEnum = append(tamActionAO0P0TypeOperationTypePropEnum, v)
	}
}

const (

	// TamActionAO0P0OperationTypeCreate captures enum value "create"
	TamActionAO0P0OperationTypeCreate string = "create"

	// TamActionAO0P0OperationTypeRemove captures enum value "remove"
	TamActionAO0P0OperationTypeRemove string = "remove"
)

// prop value enum
func (m *TamActionAO0P0) validateOperationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tamActionAO0P0TypeOperationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TamActionAO0P0) validateOperationType(formats strfmt.Registry) error {

	if swag.IsZero(m.OperationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationTypeEnum("OperationType", "body", *m.OperationType); err != nil {
		return err
	}

	return nil
}

func (m *TamActionAO0P0) validateQueries(formats strfmt.Registry) error {

	if swag.IsZero(m.Queries) { // not required
		return nil
	}

	for i := 0; i < len(m.Queries); i++ {
		if swag.IsZero(m.Queries[i]) { // not required
			continue
		}

		if m.Queries[i] != nil {
			if err := m.Queries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Queries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var tamActionAO0P0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["restApi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tamActionAO0P0TypeTypePropEnum = append(tamActionAO0P0TypeTypePropEnum, v)
	}
}

const (

	// TamActionAO0P0TypeRestAPI captures enum value "restApi"
	TamActionAO0P0TypeRestAPI string = "restApi"
)

// prop value enum
func (m *TamActionAO0P0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tamActionAO0P0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TamActionAO0P0) validateType(formats strfmt.Registry) error {

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
func (m *TamActionAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamActionAO0P0) UnmarshalBinary(b []byte) error {
	var res TamActionAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
