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

// PkixEddsaKeySpec Pkix:Eddsa Key Spec
//
// The key pair is generated using Edwards-Curve Digital Signature Algorithm (EdDSA).
// The Edwards-curve Digital Signature Algorithm (EdDSA) is a variant of
// Schnorr's signature system with (possibly twisted) Edwards curves.
//
// swagger:model pkixEddsaKeySpec
type PkixEddsaKeySpec struct {
	PkixKeyGenerationSpec

	// The EdDSA algorithm, as defined in RFC 8032.
	//
	// Enum: [Ed25519 Ed25519ph Ed25519ctx]
	Algorithm *string `json:"Algorithm,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PkixEddsaKeySpec) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PkixKeyGenerationSpec
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PkixKeyGenerationSpec = aO0

	// AO1
	var dataAO1 struct {
		Algorithm *string `json:"Algorithm,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Algorithm = dataAO1.Algorithm

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PkixEddsaKeySpec) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PkixKeyGenerationSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Algorithm *string `json:"Algorithm,omitempty"`
	}

	dataAO1.Algorithm = m.Algorithm

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pkix eddsa key spec
func (m *PkixEddsaKeySpec) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PkixKeyGenerationSpec
	if err := m.PkixKeyGenerationSpec.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pkixEddsaKeySpecTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ed25519","Ed25519ph","Ed25519ctx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pkixEddsaKeySpecTypeAlgorithmPropEnum = append(pkixEddsaKeySpecTypeAlgorithmPropEnum, v)
	}
}

// property enum
func (m *PkixEddsaKeySpec) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pkixEddsaKeySpecTypeAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PkixEddsaKeySpec) validateAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.Algorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlgorithmEnum("Algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkixEddsaKeySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkixEddsaKeySpec) UnmarshalBinary(b []byte) error {
	var res PkixEddsaKeySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
