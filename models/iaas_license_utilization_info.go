// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IaasLicenseUtilizationInfo Iaas:License Utilization Info
//
// License list with the Utilization info for UCSD.
//
// swagger:model iaasLicenseUtilizationInfo
type IaasLicenseUtilizationInfo struct {
	MoBaseComplexType

	IaasLicenseUtilizationInfoAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IaasLicenseUtilizationInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 IaasLicenseUtilizationInfoAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IaasLicenseUtilizationInfoAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IaasLicenseUtilizationInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IaasLicenseUtilizationInfoAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iaas license utilization info
func (m *IaasLicenseUtilizationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IaasLicenseUtilizationInfoAO1P1
	if err := m.IaasLicenseUtilizationInfoAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IaasLicenseUtilizationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasLicenseUtilizationInfo) UnmarshalBinary(b []byte) error {
	var res IaasLicenseUtilizationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IaasLicenseUtilizationInfoAO1P1 iaas license utilization info a o1 p1
// swagger:model IaasLicenseUtilizationInfoAO1P1
type IaasLicenseUtilizationInfoAO1P1 struct {

	// Number of licenses actually used for this feature.
	//
	// Read Only: true
	ActualUsed int64 `json:"ActualUsed,omitempty"`

	// License Label.
	//
	// Read Only: true
	Label string `json:"Label,omitempty"`

	// License limit for this license feature.
	//
	// Read Only: true
	LicensedLimit string `json:"LicensedLimit,omitempty"`

	// SKU for the license.
	//
	// Read Only: true
	Sku string `json:"Sku,omitempty"`

	// iaas license utilization info a o1 p1
	IaasLicenseUtilizationInfoAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IaasLicenseUtilizationInfoAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Number of licenses actually used for this feature.
		//
		// Read Only: true
		ActualUsed int64 `json:"ActualUsed,omitempty"`

		// License Label.
		//
		// Read Only: true
		Label string `json:"Label,omitempty"`

		// License limit for this license feature.
		//
		// Read Only: true
		LicensedLimit string `json:"LicensedLimit,omitempty"`

		// SKU for the license.
		//
		// Read Only: true
		Sku string `json:"Sku,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IaasLicenseUtilizationInfoAO1P1

	rcv.ActualUsed = stage1.ActualUsed

	rcv.Label = stage1.Label

	rcv.LicensedLimit = stage1.LicensedLimit

	rcv.Sku = stage1.Sku

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ActualUsed")

	delete(stage2, "Label")

	delete(stage2, "LicensedLimit")

	delete(stage2, "Sku")

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
		m.IaasLicenseUtilizationInfoAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IaasLicenseUtilizationInfoAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Number of licenses actually used for this feature.
		//
		// Read Only: true
		ActualUsed int64 `json:"ActualUsed,omitempty"`

		// License Label.
		//
		// Read Only: true
		Label string `json:"Label,omitempty"`

		// License limit for this license feature.
		//
		// Read Only: true
		LicensedLimit string `json:"LicensedLimit,omitempty"`

		// SKU for the license.
		//
		// Read Only: true
		Sku string `json:"Sku,omitempty"`
	}

	stage1.ActualUsed = m.ActualUsed

	stage1.Label = m.Label

	stage1.LicensedLimit = m.LicensedLimit

	stage1.Sku = m.Sku

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IaasLicenseUtilizationInfoAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IaasLicenseUtilizationInfoAO1P1)
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

// Validate validates this iaas license utilization info a o1 p1
func (m *IaasLicenseUtilizationInfoAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IaasLicenseUtilizationInfoAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasLicenseUtilizationInfoAO1P1) UnmarshalBinary(b []byte) error {
	var res IaasLicenseUtilizationInfoAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
