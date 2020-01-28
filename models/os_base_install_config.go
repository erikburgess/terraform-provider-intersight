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

// OsBaseInstallConfig Base OS Install Configuration
//
// BaseInstallConfig models the configuration required to install OS.
//
// swagger:model osBaseInstallConfig
type OsBaseInstallConfig struct {
	MoBaseMo

	// If the os.ConfigurationFile MO selected is a template that uses additional
	// placeholders other than the ones provided in standard os.Answers MO, the values
	// for those additional placeholders are provided here.
	//
	//
	AdditionalParameters []*OsPlaceHolder `json:"AdditionalParameters"`

	// Answers provided by user for the unattended OS installation.
	//
	//
	Answers *OsAnswers `json:"Answers,omitempty"`

	// User provided description about the OS install configuration.
	//
	//
	Description string `json:"Description,omitempty"`

	// The install method to be used for OS installation - vMedia, iPXE.
	// Only vMedia is supported as of now.
	//
	//
	// Enum: [vMedia]
	InstallMethod *string `json:"InstallMethod,omitempty"`

	// Parameters specific to selected OS.
	//
	//
	OperatingSystemParameters *OsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OsBaseInstallConfig) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AdditionalParameters []*OsPlaceHolder `json:"AdditionalParameters"`

		Answers *OsAnswers `json:"Answers,omitempty"`

		Description string `json:"Description,omitempty"`

		InstallMethod *string `json:"InstallMethod,omitempty"`

		OperatingSystemParameters *OsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdditionalParameters = dataAO1.AdditionalParameters

	m.Answers = dataAO1.Answers

	m.Description = dataAO1.Description

	m.InstallMethod = dataAO1.InstallMethod

	m.OperatingSystemParameters = dataAO1.OperatingSystemParameters

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OsBaseInstallConfig) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AdditionalParameters []*OsPlaceHolder `json:"AdditionalParameters"`

		Answers *OsAnswers `json:"Answers,omitempty"`

		Description string `json:"Description,omitempty"`

		InstallMethod *string `json:"InstallMethod,omitempty"`

		OperatingSystemParameters *OsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	}

	dataAO1.AdditionalParameters = m.AdditionalParameters

	dataAO1.Answers = m.Answers

	dataAO1.Description = m.Description

	dataAO1.InstallMethod = m.InstallMethod

	dataAO1.OperatingSystemParameters = m.OperatingSystemParameters

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this os base install config
func (m *OsBaseInstallConfig) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystemParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsBaseInstallConfig) validateAdditionalParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalParameters); i++ {
		if swag.IsZero(m.AdditionalParameters[i]) { // not required
			continue
		}

		if m.AdditionalParameters[i] != nil {
			if err := m.AdditionalParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AdditionalParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OsBaseInstallConfig) validateAnswers(formats strfmt.Registry) error {

	if swag.IsZero(m.Answers) { // not required
		return nil
	}

	if m.Answers != nil {
		if err := m.Answers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Answers")
			}
			return err
		}
	}

	return nil
}

var osBaseInstallConfigTypeInstallMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vMedia"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		osBaseInstallConfigTypeInstallMethodPropEnum = append(osBaseInstallConfigTypeInstallMethodPropEnum, v)
	}
}

// property enum
func (m *OsBaseInstallConfig) validateInstallMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, osBaseInstallConfigTypeInstallMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OsBaseInstallConfig) validateInstallMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateInstallMethodEnum("InstallMethod", "body", *m.InstallMethod); err != nil {
		return err
	}

	return nil
}

func (m *OsBaseInstallConfig) validateOperatingSystemParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.OperatingSystemParameters) { // not required
		return nil
	}

	if m.OperatingSystemParameters != nil {
		if err := m.OperatingSystemParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OperatingSystemParameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OsBaseInstallConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsBaseInstallConfig) UnmarshalBinary(b []byte) error {
	var res OsBaseInstallConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
