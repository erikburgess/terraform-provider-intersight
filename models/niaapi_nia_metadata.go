// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NiaapiNiaMetadata Niaapi:Nia Metadata
//
// Contains the latest Metadata available for download from server.
//
// swagger:model niaapiNiaMetadata
type NiaapiNiaMetadata struct {
	MoBaseMo

	// The detail of child file in this package.
	//
	Content []*NiaapiDetail `json:"Content"`

	// The date when this package is generated.
	//
	// Format: date-time
	Date strfmt.DateTime `json:"Date,omitempty"`

	// Chksum used to check the integrity of the Metadata file downloaded.
	//
	MetadataChksum string `json:"MetadataChksum,omitempty"`

	// The Filename of this Metadata package.
	//
	MetadataFilename string `json:"MetadataFilename,omitempty"`

	// The version number of the Metadata package.
	//
	Version int64 `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NiaapiNiaMetadata) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Content []*NiaapiDetail `json:"Content"`

		Date strfmt.DateTime `json:"Date,omitempty"`

		MetadataChksum string `json:"MetadataChksum,omitempty"`

		MetadataFilename string `json:"MetadataFilename,omitempty"`

		Version int64 `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Content = dataAO1.Content

	m.Date = dataAO1.Date

	m.MetadataChksum = dataAO1.MetadataChksum

	m.MetadataFilename = dataAO1.MetadataFilename

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NiaapiNiaMetadata) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Content []*NiaapiDetail `json:"Content"`

		Date strfmt.DateTime `json:"Date,omitempty"`

		MetadataChksum string `json:"MetadataChksum,omitempty"`

		MetadataFilename string `json:"MetadataFilename,omitempty"`

		Version int64 `json:"Version,omitempty"`
	}

	dataAO1.Content = m.Content

	dataAO1.Date = m.Date

	dataAO1.MetadataChksum = m.MetadataChksum

	dataAO1.MetadataFilename = m.MetadataFilename

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this niaapi nia metadata
func (m *NiaapiNiaMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NiaapiNiaMetadata) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NiaapiNiaMetadata) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("Date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiNiaMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiNiaMetadata) UnmarshalBinary(b []byte) error {
	var res NiaapiNiaMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
