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

// IamLdapDNSParameters Iam:Ldap Dns Parameters
//
// DNS settings used to access LDAP Providers.
//
// swagger:model iamLdapDnsParameters
type IamLdapDNSParameters struct {
	IamLdapDNSParametersAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamLdapDNSParameters) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IamLdapDNSParametersAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IamLdapDNSParametersAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamLdapDNSParameters) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.IamLdapDNSParametersAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam ldap Dns parameters
func (m *IamLdapDNSParameters) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IamLdapDNSParametersAO0P0
	if err := m.IamLdapDNSParametersAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IamLdapDNSParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLdapDNSParameters) UnmarshalBinary(b []byte) error {
	var res IamLdapDNSParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IamLdapDNSParametersAO0P0 iam ldap DNS parameters a o0 p0
// swagger:model IamLdapDNSParametersAO0P0
type IamLdapDNSParametersAO0P0 struct {

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// Domain name that acts as a source for a DNS query.
	//
	SearchDomain string `json:"SearchDomain,omitempty"`

	// Forest name that acts as a source for a DNS query.
	//
	SearchForest string `json:"SearchForest,omitempty"`

	// Source of the domain name used for the DNS SRV request.
	//
	// Enum: [Extracted Configured ConfiguredExtracted]
	Source *string `json:"Source,omitempty"`

	// iam ldap DNS parameters a o0 p0
	IamLdapDNSParametersAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IamLdapDNSParametersAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Domain name that acts as a source for a DNS query.
		//
		SearchDomain string `json:"SearchDomain,omitempty"`

		// Forest name that acts as a source for a DNS query.
		//
		SearchForest string `json:"SearchForest,omitempty"`

		// Source of the domain name used for the DNS SRV request.
		//
		// Enum: [Extracted Configured ConfiguredExtracted]
		Source *string `json:"Source,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IamLdapDNSParametersAO0P0

	rcv.ObjectType = stage1.ObjectType

	rcv.SearchDomain = stage1.SearchDomain

	rcv.SearchForest = stage1.SearchForest

	rcv.Source = stage1.Source

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "ObjectType")

	delete(stage2, "SearchDomain")

	delete(stage2, "SearchForest")

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
		m.IamLdapDNSParametersAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IamLdapDNSParametersAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// Domain name that acts as a source for a DNS query.
		//
		SearchDomain string `json:"SearchDomain,omitempty"`

		// Forest name that acts as a source for a DNS query.
		//
		SearchForest string `json:"SearchForest,omitempty"`

		// Source of the domain name used for the DNS SRV request.
		//
		// Enum: [Extracted Configured ConfiguredExtracted]
		Source *string `json:"Source,omitempty"`
	}

	stage1.ObjectType = m.ObjectType

	stage1.SearchDomain = m.SearchDomain

	stage1.SearchForest = m.SearchForest

	stage1.Source = m.Source

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IamLdapDNSParametersAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IamLdapDNSParametersAO0P0)
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

// Validate validates this iam ldap DNS parameters a o0 p0
func (m *IamLdapDNSParametersAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var iamLdapDnsParametersAO0P0TypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Extracted","Configured","ConfiguredExtracted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iamLdapDnsParametersAO0P0TypeSourcePropEnum = append(iamLdapDnsParametersAO0P0TypeSourcePropEnum, v)
	}
}

const (

	// IamLdapDNSParametersAO0P0SourceExtracted captures enum value "Extracted"
	IamLdapDNSParametersAO0P0SourceExtracted string = "Extracted"

	// IamLdapDNSParametersAO0P0SourceConfigured captures enum value "Configured"
	IamLdapDNSParametersAO0P0SourceConfigured string = "Configured"

	// IamLdapDNSParametersAO0P0SourceConfiguredExtracted captures enum value "ConfiguredExtracted"
	IamLdapDNSParametersAO0P0SourceConfiguredExtracted string = "ConfiguredExtracted"
)

// prop value enum
func (m *IamLdapDNSParametersAO0P0) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iamLdapDnsParametersAO0P0TypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IamLdapDNSParametersAO0P0) validateSource(formats strfmt.Registry) error {

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
func (m *IamLdapDNSParametersAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamLdapDNSParametersAO0P0) UnmarshalBinary(b []byte) error {
	var res IamLdapDNSParametersAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
