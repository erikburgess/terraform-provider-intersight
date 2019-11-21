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

// OsAnswers Os:Answers
//
// This MO captures the values for the most common set of fields in OS specific
// answer files. The values provided in this MO are used to construct the OS specific
// answer files (kickstart, seed, unattended xml) by replacing the fields/placeholders
// in selected os.ConfigurationFile content with the values of this MO properties.
//
// swagger:model osAnswers
type OsAnswers struct {
	OsAnswersAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OsAnswers) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OsAnswersAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OsAnswersAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OsAnswers) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.OsAnswersAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this os answers
func (m *OsAnswers) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OsAnswersAO0P0
	if err := m.OsAnswersAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OsAnswers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsAnswers) UnmarshalBinary(b []byte) error {
	var res OsAnswers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OsAnswersAO0P0 os answers a o0 p0
// swagger:model OsAnswersAO0P0
type OsAnswersAO0P0 struct {

	// If the source of the answers is a static file, the content of the file is stored as value
	// in this property.
	//
	// The value is mandatory only when the 'Source' property has been set to 'File'.
	//
	//
	AnswerFile string `json:"AnswerFile,omitempty"`

	// Hostname to be configured for the server in the OS. Hostname property is required
	// when the 'Source' property has been set to 'Template'.
	//
	// Please note that this property will be ignored for the answer instance used
	// inside the os.InstallTemplate MO since this property must have unique value for
	// each server and is irrelevant in a reusable os.InstallTemplate MO. The value of
	// this property specified in os.Install MO will be considered for the installation.
	//
	//
	Hostname string `json:"Hostname,omitempty"`

	// IP configuration type. Values are Static or Dynamic when the Source property is set to Template.
	//
	// In case of static IP configuration, IP address, gateway and other details need
	// to be populated. In case of dynamic the IP configuration is obtained dynamically
	// from DHCP.
	//
	//
	// Enum: [static dynamic]
	IPConfigType *string `json:"IpConfigType,omitempty"`

	// In case of static IP configuration, IP address, netmask and gateway details are
	// provided.
	//
	// Please note that this property will be ignored for the answer instance used
	// inside the os.InstallTemplate MO since this property must have unique value for
	// each server and is irrelevant in a reusable os.InstallTemplate MO. The value of
	// this property specified in os.Install MO will be considered for the installation.
	//
	//
	IPV4Config *CommIPV4Interface `json:"IpV4Config,omitempty"`

	// Indicates whether the value of the 'answerFile' property has been set.
	//
	// Read Only: true
	IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`

	// is root password set
	IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`

	// IP address of the name server to be configured in the OS in case the 'Source'
	// property has been set to 'Template'.
	//
	//
	Nameserver string `json:"Nameserver,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Password to be set for root/administrator user in case the 'Source'
	// property has been set to 'Template'.
	//
	//
	RootPassword string `json:"RootPassword,omitempty"`

	// Answer values can be provided from three sources - Embedded in OS image, static file,
	// or as placeholder values for an answer file template.
	//
	// Source of the answers is given as value, Embedded/File/Template.
	// 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'
	// option indicates that the answers are provided as a file. 'Template' indicates that the
	// placeholders in the selected os.ConfigurationFile MO are replaced with values provided
	// as os.Answers MO.
	//
	//
	// Enum: [Embedded File]
	Source *string `json:"Source,omitempty"`

	// os answers a o0 p0
	OsAnswersAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *OsAnswersAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// If the source of the answers is a static file, the content of the file is stored as value
		// in this property.
		//
		// The value is mandatory only when the 'Source' property has been set to 'File'.
		//
		//
		AnswerFile string `json:"AnswerFile,omitempty"`

		// Hostname to be configured for the server in the OS. Hostname property is required
		// when the 'Source' property has been set to 'Template'.
		//
		// Please note that this property will be ignored for the answer instance used
		// inside the os.InstallTemplate MO since this property must have unique value for
		// each server and is irrelevant in a reusable os.InstallTemplate MO. The value of
		// this property specified in os.Install MO will be considered for the installation.
		//
		//
		Hostname string `json:"Hostname,omitempty"`

		// IP configuration type. Values are Static or Dynamic when the Source property is set to Template.
		//
		// In case of static IP configuration, IP address, gateway and other details need
		// to be populated. In case of dynamic the IP configuration is obtained dynamically
		// from DHCP.
		//
		//
		// Enum: [static dynamic]
		IPConfigType *string `json:"IpConfigType,omitempty"`

		// In case of static IP configuration, IP address, netmask and gateway details are
		// provided.
		//
		// Please note that this property will be ignored for the answer instance used
		// inside the os.InstallTemplate MO since this property must have unique value for
		// each server and is irrelevant in a reusable os.InstallTemplate MO. The value of
		// this property specified in os.Install MO will be considered for the installation.
		//
		//
		IPV4Config *CommIPV4Interface `json:"IpV4Config,omitempty"`

		// Indicates whether the value of the 'answerFile' property has been set.
		//
		// Read Only: true
		IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`

		// is root password set
		IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`

		// IP address of the name server to be configured in the OS in case the 'Source'
		// property has been set to 'Template'.
		//
		//
		Nameserver string `json:"Nameserver,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Password to be set for root/administrator user in case the 'Source'
		// property has been set to 'Template'.
		//
		//
		RootPassword string `json:"RootPassword,omitempty"`

		// Answer values can be provided from three sources - Embedded in OS image, static file,
		// or as placeholder values for an answer file template.
		//
		// Source of the answers is given as value, Embedded/File/Template.
		// 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'
		// option indicates that the answers are provided as a file. 'Template' indicates that the
		// placeholders in the selected os.ConfigurationFile MO are replaced with values provided
		// as os.Answers MO.
		//
		//
		// Enum: [Embedded File]
		Source *string `json:"Source,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv OsAnswersAO0P0

	rcv.AnswerFile = stage1.AnswerFile

	rcv.Hostname = stage1.Hostname

	rcv.IPConfigType = stage1.IPConfigType

	rcv.IPV4Config = stage1.IPV4Config

	rcv.IsAnswerFileSet = stage1.IsAnswerFileSet

	rcv.IsRootPasswordSet = stage1.IsRootPasswordSet

	rcv.Nameserver = stage1.Nameserver

	rcv.ObjectType = stage1.ObjectType

	rcv.RootPassword = stage1.RootPassword

	rcv.Source = stage1.Source

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "AnswerFile")

	delete(stage2, "Hostname")

	delete(stage2, "IpConfigType")

	delete(stage2, "IpV4Config")

	delete(stage2, "IsAnswerFileSet")

	delete(stage2, "IsRootPasswordSet")

	delete(stage2, "Nameserver")

	delete(stage2, "ObjectType")

	delete(stage2, "RootPassword")

	delete(stage2, "Source")

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
		m.OsAnswersAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m OsAnswersAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// If the source of the answers is a static file, the content of the file is stored as value
		// in this property.
		//
		// The value is mandatory only when the 'Source' property has been set to 'File'.
		//
		//
		AnswerFile string `json:"AnswerFile,omitempty"`

		// Hostname to be configured for the server in the OS. Hostname property is required
		// when the 'Source' property has been set to 'Template'.
		//
		// Please note that this property will be ignored for the answer instance used
		// inside the os.InstallTemplate MO since this property must have unique value for
		// each server and is irrelevant in a reusable os.InstallTemplate MO. The value of
		// this property specified in os.Install MO will be considered for the installation.
		//
		//
		Hostname string `json:"Hostname,omitempty"`

		// IP configuration type. Values are Static or Dynamic when the Source property is set to Template.
		//
		// In case of static IP configuration, IP address, gateway and other details need
		// to be populated. In case of dynamic the IP configuration is obtained dynamically
		// from DHCP.
		//
		//
		// Enum: [static dynamic]
		IPConfigType *string `json:"IpConfigType,omitempty"`

		// In case of static IP configuration, IP address, netmask and gateway details are
		// provided.
		//
		// Please note that this property will be ignored for the answer instance used
		// inside the os.InstallTemplate MO since this property must have unique value for
		// each server and is irrelevant in a reusable os.InstallTemplate MO. The value of
		// this property specified in os.Install MO will be considered for the installation.
		//
		//
		IPV4Config *CommIPV4Interface `json:"IpV4Config,omitempty"`

		// Indicates whether the value of the 'answerFile' property has been set.
		//
		// Read Only: true
		IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`

		// is root password set
		IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`

		// IP address of the name server to be configured in the OS in case the 'Source'
		// property has been set to 'Template'.
		//
		//
		Nameserver string `json:"Nameserver,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Password to be set for root/administrator user in case the 'Source'
		// property has been set to 'Template'.
		//
		//
		RootPassword string `json:"RootPassword,omitempty"`

		// Answer values can be provided from three sources - Embedded in OS image, static file,
		// or as placeholder values for an answer file template.
		//
		// Source of the answers is given as value, Embedded/File/Template.
		// 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'
		// option indicates that the answers are provided as a file. 'Template' indicates that the
		// placeholders in the selected os.ConfigurationFile MO are replaced with values provided
		// as os.Answers MO.
		//
		//
		// Enum: [Embedded File]
		Source *string `json:"Source,omitempty"`
	}

	stage1.AnswerFile = m.AnswerFile

	stage1.Hostname = m.Hostname

	stage1.IPConfigType = m.IPConfigType

	stage1.IPV4Config = m.IPV4Config

	stage1.IsAnswerFileSet = m.IsAnswerFileSet

	stage1.IsRootPasswordSet = m.IsRootPasswordSet

	stage1.Nameserver = m.Nameserver

	stage1.ObjectType = m.ObjectType

	stage1.RootPassword = m.RootPassword

	stage1.Source = m.Source

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.OsAnswersAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.OsAnswersAO0P0)
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

// Validate validates this os answers a o0 p0
func (m *OsAnswersAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPConfigType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4Config(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var osAnswersAO0P0TypeIPConfigTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["static","dynamic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		osAnswersAO0P0TypeIPConfigTypePropEnum = append(osAnswersAO0P0TypeIPConfigTypePropEnum, v)
	}
}

const (

	// OsAnswersAO0P0IPConfigTypeStatic captures enum value "static"
	OsAnswersAO0P0IPConfigTypeStatic string = "static"

	// OsAnswersAO0P0IPConfigTypeDynamic captures enum value "dynamic"
	OsAnswersAO0P0IPConfigTypeDynamic string = "dynamic"
)

// prop value enum
func (m *OsAnswersAO0P0) validateIPConfigTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, osAnswersAO0P0TypeIPConfigTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OsAnswersAO0P0) validateIPConfigType(formats strfmt.Registry) error {

	if swag.IsZero(m.IPConfigType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPConfigTypeEnum("IpConfigType", "body", *m.IPConfigType); err != nil {
		return err
	}

	return nil
}

func (m *OsAnswersAO0P0) validateIPV4Config(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4Config) { // not required
		return nil
	}

	if m.IPV4Config != nil {
		if err := m.IPV4Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IpV4Config")
			}
			return err
		}
	}

	return nil
}

var osAnswersAO0P0TypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Embedded","File"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		osAnswersAO0P0TypeSourcePropEnum = append(osAnswersAO0P0TypeSourcePropEnum, v)
	}
}

const (

	// OsAnswersAO0P0SourceEmbedded captures enum value "Embedded"
	OsAnswersAO0P0SourceEmbedded string = "Embedded"

	// OsAnswersAO0P0SourceFile captures enum value "File"
	OsAnswersAO0P0SourceFile string = "File"
)

// prop value enum
func (m *OsAnswersAO0P0) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, osAnswersAO0P0TypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OsAnswersAO0P0) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("Source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OsAnswersAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsAnswersAO0P0) UnmarshalBinary(b []byte) error {
	var res OsAnswersAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
