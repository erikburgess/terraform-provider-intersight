// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TamAPIDataSource Tam:Api Data Source
//
// Data source using Intersight API to collect data regarding managed devices.
//
// swagger:model tamApiDataSource
type TamAPIDataSource struct {
	TamBaseDataSource

	// Type of Intersight managed object used as data source.
	//
	MoType string `json:"MoType,omitempty"`

	// Optional set of Queries to filter the output for Api datasource. the queries are executed in the order specified.
	//
	Queries []*TamQueryEntry `json:"Queries"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TamAPIDataSource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TamBaseDataSource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TamBaseDataSource = aO0

	// AO1
	var dataAO1 struct {
		MoType string `json:"MoType,omitempty"`

		Queries []*TamQueryEntry `json:"Queries"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.MoType = dataAO1.MoType

	m.Queries = dataAO1.Queries

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TamAPIDataSource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TamBaseDataSource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		MoType string `json:"MoType,omitempty"`

		Queries []*TamQueryEntry `json:"Queries"`
	}

	dataAO1.MoType = m.MoType

	dataAO1.Queries = m.Queries

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tam Api data source
func (m *TamAPIDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TamBaseDataSource
	if err := m.TamBaseDataSource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TamAPIDataSource) validateQueries(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *TamAPIDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TamAPIDataSource) UnmarshalBinary(b []byte) error {
	var res TamAPIDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
