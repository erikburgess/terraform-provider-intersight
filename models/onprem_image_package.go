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

// OnpremImagePackage Onprem:Image Package
//
// ImagePackage encapsulates a software image package. ImagePackage can be
// a docker image, a UI web image, an endpoint (e.g. UCSM) image, a device
// connector image or an ansible scripts package.
//
// swagger:model onpremImagePackage
type OnpremImagePackage struct {
	MoBaseComplexType

	OnpremImagePackageAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OnpremImagePackage) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 OnpremImagePackageAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.OnpremImagePackageAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OnpremImagePackage) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.OnpremImagePackageAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this onprem image package
func (m *OnpremImagePackage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OnpremImagePackageAO1P1
	if err := m.OnpremImagePackageAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OnpremImagePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnpremImagePackage) UnmarshalBinary(b []byte) error {
	var res OnpremImagePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OnpremImagePackageAO1P1 onprem image package a o1 p1
// swagger:model OnpremImagePackageAO1P1
type OnpremImagePackageAO1P1 struct {

	// Optional file path of the image package.
	//
	// Read Only: true
	FilePath string `json:"FilePath,omitempty"`

	// Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.
	//
	// Read Only: true
	FileSha string `json:"FileSha,omitempty"`

	// Image file size in bytes.
	//
	// Read Only: true
	FileSize int64 `json:"FileSize,omitempty"`

	// Image file's last modified date and time.
	//
	// Read Only: true
	// Format: date-time
	FileTime strfmt.DateTime `json:"FileTime,omitempty"`

	// Filename of the image package.
	//
	// Read Only: true
	Filename string `json:"Filename,omitempty"`

	// Name of the software image package.
	//
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Image package type (e.g. service, system etc.).
	//
	// Read Only: true
	PackageType string `json:"PackageType,omitempty"`

	// Image package version string.
	//
	// Read Only: true
	Version string `json:"Version,omitempty"`

	// onprem image package a o1 p1
	OnpremImagePackageAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *OnpremImagePackageAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Optional file path of the image package.
		//
		// Read Only: true
		FilePath string `json:"FilePath,omitempty"`

		// Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.
		//
		// Read Only: true
		FileSha string `json:"FileSha,omitempty"`

		// Image file size in bytes.
		//
		// Read Only: true
		FileSize int64 `json:"FileSize,omitempty"`

		// Image file's last modified date and time.
		//
		// Read Only: true
		// Format: date-time
		FileTime strfmt.DateTime `json:"FileTime,omitempty"`

		// Filename of the image package.
		//
		// Read Only: true
		Filename string `json:"Filename,omitempty"`

		// Name of the software image package.
		//
		// Read Only: true
		Name string `json:"Name,omitempty"`

		// Image package type (e.g. service, system etc.).
		//
		// Read Only: true
		PackageType string `json:"PackageType,omitempty"`

		// Image package version string.
		//
		// Read Only: true
		Version string `json:"Version,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv OnpremImagePackageAO1P1

	rcv.FilePath = stage1.FilePath

	rcv.FileSha = stage1.FileSha

	rcv.FileSize = stage1.FileSize

	rcv.FileTime = stage1.FileTime

	rcv.Filename = stage1.Filename

	rcv.Name = stage1.Name

	rcv.PackageType = stage1.PackageType

	rcv.Version = stage1.Version

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "FilePath")

	delete(stage2, "FileSha")

	delete(stage2, "FileSize")

	delete(stage2, "FileTime")

	delete(stage2, "Filename")

	delete(stage2, "Name")

	delete(stage2, "PackageType")

	delete(stage2, "Version")

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
		m.OnpremImagePackageAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m OnpremImagePackageAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Optional file path of the image package.
		//
		// Read Only: true
		FilePath string `json:"FilePath,omitempty"`

		// Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.
		//
		// Read Only: true
		FileSha string `json:"FileSha,omitempty"`

		// Image file size in bytes.
		//
		// Read Only: true
		FileSize int64 `json:"FileSize,omitempty"`

		// Image file's last modified date and time.
		//
		// Read Only: true
		// Format: date-time
		FileTime strfmt.DateTime `json:"FileTime,omitempty"`

		// Filename of the image package.
		//
		// Read Only: true
		Filename string `json:"Filename,omitempty"`

		// Name of the software image package.
		//
		// Read Only: true
		Name string `json:"Name,omitempty"`

		// Image package type (e.g. service, system etc.).
		//
		// Read Only: true
		PackageType string `json:"PackageType,omitempty"`

		// Image package version string.
		//
		// Read Only: true
		Version string `json:"Version,omitempty"`
	}

	stage1.FilePath = m.FilePath

	stage1.FileSha = m.FileSha

	stage1.FileSize = m.FileSize

	stage1.FileTime = m.FileTime

	stage1.Filename = m.Filename

	stage1.Name = m.Name

	stage1.PackageType = m.PackageType

	stage1.Version = m.Version

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.OnpremImagePackageAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.OnpremImagePackageAO1P1)
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

// Validate validates this onprem image package a o1 p1
func (m *OnpremImagePackageAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnpremImagePackageAO1P1) validateFileTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FileTime) { // not required
		return nil
	}

	if err := validate.FormatOf("FileTime", "body", "date-time", m.FileTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnpremImagePackageAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnpremImagePackageAO1P1) UnmarshalBinary(b []byte) error {
	var res OnpremImagePackageAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
