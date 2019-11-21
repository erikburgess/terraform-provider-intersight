// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NiaapiDcnmCcoPost Niaapi:Dcnm Cco Post
//
// The post reporting a new release is available for DCNM.
//
// swagger:model niaapiDcnmCcoPost
type NiaapiDcnmCcoPost struct {
	NiaapiNewReleasePost

	NiaapiDcnmCcoPostAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NiaapiDcnmCcoPost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NiaapiNewReleasePost
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NiaapiNewReleasePost = aO0

	// AO1
	var aO1 NiaapiDcnmCcoPostAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NiaapiDcnmCcoPostAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NiaapiDcnmCcoPost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NiaapiNewReleasePost)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.NiaapiDcnmCcoPostAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this niaapi dcnm cco post
func (m *NiaapiDcnmCcoPost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NiaapiNewReleasePost
	if err := m.NiaapiNewReleasePost.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with NiaapiDcnmCcoPostAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiDcnmCcoPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiDcnmCcoPost) UnmarshalBinary(b []byte) error {
	var res NiaapiDcnmCcoPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NiaapiDcnmCcoPostAllOf1 niaapi dcnm cco post all of1
// swagger:model NiaapiDcnmCcoPostAllOf1
type NiaapiDcnmCcoPostAllOf1 interface{}
