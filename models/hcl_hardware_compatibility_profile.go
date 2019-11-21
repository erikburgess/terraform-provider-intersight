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

// HclHardwareCompatibilityProfile Hcl:Hardware Compatibility Profile
//
// Profile giving server hardware details, OS details and UCS Version details.
//
// swagger:model hclHardwareCompatibilityProfile
type HclHardwareCompatibilityProfile struct {
	HclHardwareCompatibilityProfileAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HclHardwareCompatibilityProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HclHardwareCompatibilityProfileAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HclHardwareCompatibilityProfileAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HclHardwareCompatibilityProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.HclHardwareCompatibilityProfileAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hcl hardware compatibility profile
func (m *HclHardwareCompatibilityProfile) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HclHardwareCompatibilityProfileAO0P0
	if err := m.HclHardwareCompatibilityProfileAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HclHardwareCompatibilityProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HclHardwareCompatibilityProfile) UnmarshalBinary(b []byte) error {
	var res HclHardwareCompatibilityProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HclHardwareCompatibilityProfileAO0P0 hcl hardware compatibility profile a o0 p0
// swagger:model HclHardwareCompatibilityProfileAO0P0
type HclHardwareCompatibilityProfileAO0P0 struct {

	// Url for the ISO with the drivers supported for the server.
	//
	DriverIsoURL string `json:"DriverIsoUrl,omitempty"`

	// Error code indicating the compatibility status.
	//
	// Read Only: true
	// Enum: [Success Unknown UnknownServer InvalidUcsVersion ProcessorNotSupported OSNotSupported OSUnknown UCSVersionNotSupported UcsVersionServerOSCombinationNotSupported ProductUnknown ProductNotSupported DriverNameNotSupported FirmwareVersionNotSupported DriverVersionNotSupported FirmwareVersionDriverVersionCombinationNotSupported FirmwareVersionAndDriverVersionNotSupported FirmwareVersionAndDriverNameNotSupported InternalError MarshallingError Exempted]
	ErrorCode string `json:"ErrorCode,omitempty"`

	// Identifier of the hardware compatibility profile.
	//
	ID string `json:"Id,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Vendor of the Operating System running on the server.
	//
	OsVendor string `json:"OsVendor,omitempty"`

	// Version of the Operating System running on the server.
	//
	OsVersion string `json:"OsVersion,omitempty"`

	// Model of the processor present in the server.
	//
	ProcessorModel string `json:"ProcessorModel,omitempty"`

	// List of the products (adapters/storage controllers) for which compatibility status needs to be checked.
	//
	Products []*HclProduct `json:"Products"`

	// Model of the server as returned by UCSM/CIMC XML API.
	//
	ServerModel string `json:"ServerModel,omitempty"`

	// Revision of the server model.
	//
	ServerRevision string `json:"ServerRevision,omitempty"`

	// Version of the UCS software.
	//
	UcsVersion string `json:"UcsVersion,omitempty"`

	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	//
	// Enum: [UCSM IMC]
	VersionType *string `json:"VersionType,omitempty"`

	// hcl hardware compatibility profile a o0 p0
	HclHardwareCompatibilityProfileAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *HclHardwareCompatibilityProfileAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Url for the ISO with the drivers supported for the server.
		//
		DriverIsoURL string `json:"DriverIsoUrl,omitempty"`

		// Error code indicating the compatibility status.
		//
		// Read Only: true
		// Enum: [Success Unknown UnknownServer InvalidUcsVersion ProcessorNotSupported OSNotSupported OSUnknown UCSVersionNotSupported UcsVersionServerOSCombinationNotSupported ProductUnknown ProductNotSupported DriverNameNotSupported FirmwareVersionNotSupported DriverVersionNotSupported FirmwareVersionDriverVersionCombinationNotSupported FirmwareVersionAndDriverVersionNotSupported FirmwareVersionAndDriverNameNotSupported InternalError MarshallingError Exempted]
		ErrorCode string `json:"ErrorCode,omitempty"`

		// Identifier of the hardware compatibility profile.
		//
		ID string `json:"Id,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Vendor of the Operating System running on the server.
		//
		OsVendor string `json:"OsVendor,omitempty"`

		// Version of the Operating System running on the server.
		//
		OsVersion string `json:"OsVersion,omitempty"`

		// Model of the processor present in the server.
		//
		ProcessorModel string `json:"ProcessorModel,omitempty"`

		// List of the products (adapters/storage controllers) for which compatibility status needs to be checked.
		//
		Products []*HclProduct `json:"Products"`

		// Model of the server as returned by UCSM/CIMC XML API.
		//
		ServerModel string `json:"ServerModel,omitempty"`

		// Revision of the server model.
		//
		ServerRevision string `json:"ServerRevision,omitempty"`

		// Version of the UCS software.
		//
		UcsVersion string `json:"UcsVersion,omitempty"`

		// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
		//
		// Enum: [UCSM IMC]
		VersionType *string `json:"VersionType,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv HclHardwareCompatibilityProfileAO0P0

	rcv.DriverIsoURL = stage1.DriverIsoURL

	rcv.ErrorCode = stage1.ErrorCode

	rcv.ID = stage1.ID

	rcv.ObjectType = stage1.ObjectType

	rcv.OsVendor = stage1.OsVendor

	rcv.OsVersion = stage1.OsVersion

	rcv.ProcessorModel = stage1.ProcessorModel

	rcv.Products = stage1.Products

	rcv.ServerModel = stage1.ServerModel

	rcv.ServerRevision = stage1.ServerRevision

	rcv.UcsVersion = stage1.UcsVersion

	rcv.VersionType = stage1.VersionType

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DriverIsoUrl")

	delete(stage2, "ErrorCode")

	delete(stage2, "Id")

	delete(stage2, "ObjectType")

	delete(stage2, "OsVendor")

	delete(stage2, "OsVersion")

	delete(stage2, "ProcessorModel")

	delete(stage2, "Products")

	delete(stage2, "ServerModel")

	delete(stage2, "ServerRevision")

	delete(stage2, "UcsVersion")

	delete(stage2, "VersionType")

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
		m.HclHardwareCompatibilityProfileAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m HclHardwareCompatibilityProfileAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Url for the ISO with the drivers supported for the server.
		//
		DriverIsoURL string `json:"DriverIsoUrl,omitempty"`

		// Error code indicating the compatibility status.
		//
		// Read Only: true
		// Enum: [Success Unknown UnknownServer InvalidUcsVersion ProcessorNotSupported OSNotSupported OSUnknown UCSVersionNotSupported UcsVersionServerOSCombinationNotSupported ProductUnknown ProductNotSupported DriverNameNotSupported FirmwareVersionNotSupported DriverVersionNotSupported FirmwareVersionDriverVersionCombinationNotSupported FirmwareVersionAndDriverVersionNotSupported FirmwareVersionAndDriverNameNotSupported InternalError MarshallingError Exempted]
		ErrorCode string `json:"ErrorCode,omitempty"`

		// Identifier of the hardware compatibility profile.
		//
		ID string `json:"Id,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Vendor of the Operating System running on the server.
		//
		OsVendor string `json:"OsVendor,omitempty"`

		// Version of the Operating System running on the server.
		//
		OsVersion string `json:"OsVersion,omitempty"`

		// Model of the processor present in the server.
		//
		ProcessorModel string `json:"ProcessorModel,omitempty"`

		// List of the products (adapters/storage controllers) for which compatibility status needs to be checked.
		//
		Products []*HclProduct `json:"Products"`

		// Model of the server as returned by UCSM/CIMC XML API.
		//
		ServerModel string `json:"ServerModel,omitempty"`

		// Revision of the server model.
		//
		ServerRevision string `json:"ServerRevision,omitempty"`

		// Version of the UCS software.
		//
		UcsVersion string `json:"UcsVersion,omitempty"`

		// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
		//
		// Enum: [UCSM IMC]
		VersionType *string `json:"VersionType,omitempty"`
	}

	stage1.DriverIsoURL = m.DriverIsoURL

	stage1.ErrorCode = m.ErrorCode

	stage1.ID = m.ID

	stage1.ObjectType = m.ObjectType

	stage1.OsVendor = m.OsVendor

	stage1.OsVersion = m.OsVersion

	stage1.ProcessorModel = m.ProcessorModel

	stage1.Products = m.Products

	stage1.ServerModel = m.ServerModel

	stage1.ServerRevision = m.ServerRevision

	stage1.UcsVersion = m.UcsVersion

	stage1.VersionType = m.VersionType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.HclHardwareCompatibilityProfileAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.HclHardwareCompatibilityProfileAO0P0)
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

// Validate validates this hcl hardware compatibility profile a o0 p0
func (m *HclHardwareCompatibilityProfileAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hclHardwareCompatibilityProfileAO0P0TypeErrorCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Success","Unknown","UnknownServer","InvalidUcsVersion","ProcessorNotSupported","OSNotSupported","OSUnknown","UCSVersionNotSupported","UcsVersionServerOSCombinationNotSupported","ProductUnknown","ProductNotSupported","DriverNameNotSupported","FirmwareVersionNotSupported","DriverVersionNotSupported","FirmwareVersionDriverVersionCombinationNotSupported","FirmwareVersionAndDriverVersionNotSupported","FirmwareVersionAndDriverNameNotSupported","InternalError","MarshallingError","Exempted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hclHardwareCompatibilityProfileAO0P0TypeErrorCodePropEnum = append(hclHardwareCompatibilityProfileAO0P0TypeErrorCodePropEnum, v)
	}
}

const (

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeSuccess captures enum value "Success"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeSuccess string = "Success"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeUnknown captures enum value "Unknown"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeUnknown string = "Unknown"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeUnknownServer captures enum value "UnknownServer"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeUnknownServer string = "UnknownServer"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeInvalidUcsVersion captures enum value "InvalidUcsVersion"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeInvalidUcsVersion string = "InvalidUcsVersion"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeProcessorNotSupported captures enum value "ProcessorNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeProcessorNotSupported string = "ProcessorNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeOSNotSupported captures enum value "OSNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeOSNotSupported string = "OSNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeOSUnknown captures enum value "OSUnknown"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeOSUnknown string = "OSUnknown"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeUCSVersionNotSupported captures enum value "UCSVersionNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeUCSVersionNotSupported string = "UCSVersionNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeUcsVersionServerOSCombinationNotSupported captures enum value "UcsVersionServerOSCombinationNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeUcsVersionServerOSCombinationNotSupported string = "UcsVersionServerOSCombinationNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeProductUnknown captures enum value "ProductUnknown"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeProductUnknown string = "ProductUnknown"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeProductNotSupported captures enum value "ProductNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeProductNotSupported string = "ProductNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeDriverNameNotSupported captures enum value "DriverNameNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeDriverNameNotSupported string = "DriverNameNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionNotSupported captures enum value "FirmwareVersionNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionNotSupported string = "FirmwareVersionNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeDriverVersionNotSupported captures enum value "DriverVersionNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeDriverVersionNotSupported string = "DriverVersionNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionDriverVersionCombinationNotSupported captures enum value "FirmwareVersionDriverVersionCombinationNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionDriverVersionCombinationNotSupported string = "FirmwareVersionDriverVersionCombinationNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionAndDriverVersionNotSupported captures enum value "FirmwareVersionAndDriverVersionNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionAndDriverVersionNotSupported string = "FirmwareVersionAndDriverVersionNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionAndDriverNameNotSupported captures enum value "FirmwareVersionAndDriverNameNotSupported"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeFirmwareVersionAndDriverNameNotSupported string = "FirmwareVersionAndDriverNameNotSupported"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeInternalError captures enum value "InternalError"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeInternalError string = "InternalError"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeMarshallingError captures enum value "MarshallingError"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeMarshallingError string = "MarshallingError"

	// HclHardwareCompatibilityProfileAO0P0ErrorCodeExempted captures enum value "Exempted"
	HclHardwareCompatibilityProfileAO0P0ErrorCodeExempted string = "Exempted"
)

// prop value enum
func (m *HclHardwareCompatibilityProfileAO0P0) validateErrorCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hclHardwareCompatibilityProfileAO0P0TypeErrorCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HclHardwareCompatibilityProfileAO0P0) validateErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorCodeEnum("ErrorCode", "body", m.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (m *HclHardwareCompatibilityProfileAO0P0) validateProducts(formats strfmt.Registry) error {

	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var hclHardwareCompatibilityProfileAO0P0TypeVersionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UCSM","IMC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hclHardwareCompatibilityProfileAO0P0TypeVersionTypePropEnum = append(hclHardwareCompatibilityProfileAO0P0TypeVersionTypePropEnum, v)
	}
}

const (

	// HclHardwareCompatibilityProfileAO0P0VersionTypeUCSM captures enum value "UCSM"
	HclHardwareCompatibilityProfileAO0P0VersionTypeUCSM string = "UCSM"

	// HclHardwareCompatibilityProfileAO0P0VersionTypeIMC captures enum value "IMC"
	HclHardwareCompatibilityProfileAO0P0VersionTypeIMC string = "IMC"
)

// prop value enum
func (m *HclHardwareCompatibilityProfileAO0P0) validateVersionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hclHardwareCompatibilityProfileAO0P0TypeVersionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HclHardwareCompatibilityProfileAO0P0) validateVersionType(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionTypeEnum("VersionType", "body", *m.VersionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HclHardwareCompatibilityProfileAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HclHardwareCompatibilityProfileAO0P0) UnmarshalBinary(b []byte) error {
	var res HclHardwareCompatibilityProfileAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
