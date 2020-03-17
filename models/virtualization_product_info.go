// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualizationProductInfo Virtualization:Product Info
//
// The details of the product offering.
//
// swagger:model virtualizationProductInfo
type VirtualizationProductInfo struct {
	MoBaseComplexType

	VirtualizationProductInfoAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VirtualizationProductInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VirtualizationProductInfoAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VirtualizationProductInfoAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VirtualizationProductInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VirtualizationProductInfoAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this virtualization product info
func (m *VirtualizationProductInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VirtualizationProductInfoAO1P1
	if err := m.VirtualizationProductInfoAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationProductInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationProductInfo) UnmarshalBinary(b []byte) error {
	var res VirtualizationProductInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualizationProductInfoAO1P1 virtualization product info a o1 p1
//
// swagger:model VirtualizationProductInfoAO1P1
type VirtualizationProductInfoAO1P1 struct {

	// Commercial product name, example, VMware ESXi.
	ProductName string `json:"ProductName,omitempty"`

	// Indication of product type by the vendor, example, embeddedEsx.
	ProductType string `json:"ProductType,omitempty"`

	// Commercial vendor name, example, VMware, Inc'.
	ProductVendor string `json:"ProductVendor,omitempty"`

	// Hypervisor version running on the system.
	Version string `json:"Version,omitempty"`

	// virtualization product info a o1 p1
	VirtualizationProductInfoAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VirtualizationProductInfoAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Commercial product name, example, VMware ESXi.
		ProductName string `json:"ProductName,omitempty"`

		// Indication of product type by the vendor, example, embeddedEsx.
		ProductType string `json:"ProductType,omitempty"`

		// Commercial vendor name, example, VMware, Inc'.
		ProductVendor string `json:"ProductVendor,omitempty"`

		// Hypervisor version running on the system.
		Version string `json:"Version,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VirtualizationProductInfoAO1P1

	rcv.ProductName = stage1.ProductName
	rcv.ProductType = stage1.ProductType
	rcv.ProductVendor = stage1.ProductVendor
	rcv.Version = stage1.Version
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ProductName")
	delete(stage2, "ProductType")
	delete(stage2, "ProductVendor")
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
		m.VirtualizationProductInfoAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VirtualizationProductInfoAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Commercial product name, example, VMware ESXi.
		ProductName string `json:"ProductName,omitempty"`

		// Indication of product type by the vendor, example, embeddedEsx.
		ProductType string `json:"ProductType,omitempty"`

		// Commercial vendor name, example, VMware, Inc'.
		ProductVendor string `json:"ProductVendor,omitempty"`

		// Hypervisor version running on the system.
		Version string `json:"Version,omitempty"`
	}

	stage1.ProductName = m.ProductName
	stage1.ProductType = m.ProductType
	stage1.ProductVendor = m.ProductVendor
	stage1.Version = m.Version

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VirtualizationProductInfoAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VirtualizationProductInfoAO1P1)
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

// Validate validates this virtualization product info a o1 p1
func (m *VirtualizationProductInfoAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationProductInfoAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationProductInfoAO1P1) UnmarshalBinary(b []byte) error {
	var res VirtualizationProductInfoAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
