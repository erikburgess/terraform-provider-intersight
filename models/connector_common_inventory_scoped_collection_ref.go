// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConnectorCommonInventoryScopedCollectionRef connector common inventory scoped collection ref
//
// swagger:model connectorCommonInventoryScopedCollectionRef
type ConnectorCommonInventoryScopedCollectionRef struct {
	MoMoRef
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ConnectorCommonInventoryScopedCollectionRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoMoRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoMoRef = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ConnectorCommonInventoryScopedCollectionRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MoMoRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this connector common inventory scoped collection ref
func (m *ConnectorCommonInventoryScopedCollectionRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoMoRef
	if err := m.MoMoRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorCommonInventoryScopedCollectionRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorCommonInventoryScopedCollectionRef) UnmarshalBinary(b []byte) error {
	var res ConnectorCommonInventoryScopedCollectionRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
