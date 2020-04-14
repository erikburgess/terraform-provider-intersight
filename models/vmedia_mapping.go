// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VmediaMapping Vmedia:Mapping
//
// Virtual Media mapping settings required to map images from remote server.
//
// swagger:model vmediaMapping
type VmediaMapping struct {
	MoBaseComplexType

	VmediaMappingAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VmediaMapping) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VmediaMappingAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VmediaMappingAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VmediaMapping) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VmediaMappingAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vmedia mapping
func (m *VmediaMapping) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VmediaMappingAO1P1
	if err := m.VmediaMappingAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VmediaMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmediaMapping) UnmarshalBinary(b []byte) error {
	var res VmediaMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VmediaMappingAO1P1 vmedia mapping a o1 p1
//
// swagger:model VmediaMappingAO1P1
type VmediaMappingAO1P1 struct {

	// Type of Authentication protocol when CIFS is used for communication with the remote server.
	// Enum: [none ntlm ntlmi ntlmv2 ntlmv2i ntlmssp ntlmsspi]
	AuthenticationProtocol *string `json:"AuthenticationProtocol,omitempty"`

	// Type of remote Virtual Media device.
	// Enum: [cdd hdd]
	DeviceType *string `json:"DeviceType,omitempty"`

	// IP address or hostname of the remote server.
	HostName string `json:"HostName,omitempty"`

	// Indicates whether the value of the 'password' property has been set.
	// Read Only: true
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

	// Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\n For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\n For CIFS, supported options are soft, nounix, noserverino, guest.\n For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\n For HTTP/HTTPS, the only supported option is noauto.
	MountOptions string `json:"MountOptions,omitempty"`

	// Protocol to use to communicate with the remote server.
	// Enum: [nfs cifs http https]
	MountProtocol *string `json:"MountProtocol,omitempty"`

	// Password associated with the username.
	Password string `json:"Password,omitempty"`

	// Name of the remote file. It should be of .img format for HDD Virtual Media mapping and .iso format for CDD Virtual Media mapping.
	RemoteFile string `json:"RemoteFile,omitempty"`

	// URL path to the location of the image on the remote server. The preferred format is '/path'.
	RemotePath string `json:"RemotePath,omitempty"`

	// Username to log in to the remote server.
	Username string `json:"Username,omitempty"`

	// Identity of the image for Virtual Media mapping.
	VolumeName string `json:"VolumeName,omitempty"`

	// vmedia mapping a o1 p1
	VmediaMappingAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VmediaMappingAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Type of Authentication protocol when CIFS is used for communication with the remote server.
		// Enum: [none ntlm ntlmi ntlmv2 ntlmv2i ntlmssp ntlmsspi]
		AuthenticationProtocol *string `json:"AuthenticationProtocol,omitempty"`

		// Type of remote Virtual Media device.
		// Enum: [cdd hdd]
		DeviceType *string `json:"DeviceType,omitempty"`

		// IP address or hostname of the remote server.
		HostName string `json:"HostName,omitempty"`

		// Indicates whether the value of the 'password' property has been set.
		// Read Only: true
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		// Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\n For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\n For CIFS, supported options are soft, nounix, noserverino, guest.\n For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\n For HTTP/HTTPS, the only supported option is noauto.
		MountOptions string `json:"MountOptions,omitempty"`

		// Protocol to use to communicate with the remote server.
		// Enum: [nfs cifs http https]
		MountProtocol *string `json:"MountProtocol,omitempty"`

		// Password associated with the username.
		Password string `json:"Password,omitempty"`

		// Name of the remote file. It should be of .img format for HDD Virtual Media mapping and .iso format for CDD Virtual Media mapping.
		RemoteFile string `json:"RemoteFile,omitempty"`

		// URL path to the location of the image on the remote server. The preferred format is '/path'.
		RemotePath string `json:"RemotePath,omitempty"`

		// Username to log in to the remote server.
		Username string `json:"Username,omitempty"`

		// Identity of the image for Virtual Media mapping.
		VolumeName string `json:"VolumeName,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VmediaMappingAO1P1

	rcv.AuthenticationProtocol = stage1.AuthenticationProtocol
	rcv.DeviceType = stage1.DeviceType
	rcv.HostName = stage1.HostName
	rcv.IsPasswordSet = stage1.IsPasswordSet
	rcv.MountOptions = stage1.MountOptions
	rcv.MountProtocol = stage1.MountProtocol
	rcv.Password = stage1.Password
	rcv.RemoteFile = stage1.RemoteFile
	rcv.RemotePath = stage1.RemotePath
	rcv.Username = stage1.Username
	rcv.VolumeName = stage1.VolumeName
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "AuthenticationProtocol")
	delete(stage2, "DeviceType")
	delete(stage2, "HostName")
	delete(stage2, "IsPasswordSet")
	delete(stage2, "MountOptions")
	delete(stage2, "MountProtocol")
	delete(stage2, "Password")
	delete(stage2, "RemoteFile")
	delete(stage2, "RemotePath")
	delete(stage2, "Username")
	delete(stage2, "VolumeName")
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
		m.VmediaMappingAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VmediaMappingAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Type of Authentication protocol when CIFS is used for communication with the remote server.
		// Enum: [none ntlm ntlmi ntlmv2 ntlmv2i ntlmssp ntlmsspi]
		AuthenticationProtocol *string `json:"AuthenticationProtocol,omitempty"`

		// Type of remote Virtual Media device.
		// Enum: [cdd hdd]
		DeviceType *string `json:"DeviceType,omitempty"`

		// IP address or hostname of the remote server.
		HostName string `json:"HostName,omitempty"`

		// Indicates whether the value of the 'password' property has been set.
		// Read Only: true
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`

		// Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\n For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\n For CIFS, supported options are soft, nounix, noserverino, guest.\n For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\n For HTTP/HTTPS, the only supported option is noauto.
		MountOptions string `json:"MountOptions,omitempty"`

		// Protocol to use to communicate with the remote server.
		// Enum: [nfs cifs http https]
		MountProtocol *string `json:"MountProtocol,omitempty"`

		// Password associated with the username.
		Password string `json:"Password,omitempty"`

		// Name of the remote file. It should be of .img format for HDD Virtual Media mapping and .iso format for CDD Virtual Media mapping.
		RemoteFile string `json:"RemoteFile,omitempty"`

		// URL path to the location of the image on the remote server. The preferred format is '/path'.
		RemotePath string `json:"RemotePath,omitempty"`

		// Username to log in to the remote server.
		Username string `json:"Username,omitempty"`

		// Identity of the image for Virtual Media mapping.
		VolumeName string `json:"VolumeName,omitempty"`
	}

	stage1.AuthenticationProtocol = m.AuthenticationProtocol
	stage1.DeviceType = m.DeviceType
	stage1.HostName = m.HostName
	stage1.IsPasswordSet = m.IsPasswordSet
	stage1.MountOptions = m.MountOptions
	stage1.MountProtocol = m.MountProtocol
	stage1.Password = m.Password
	stage1.RemoteFile = m.RemoteFile
	stage1.RemotePath = m.RemotePath
	stage1.Username = m.Username
	stage1.VolumeName = m.VolumeName

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VmediaMappingAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VmediaMappingAO1P1)
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

// Validate validates this vmedia mapping a o1 p1
func (m *VmediaMappingAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vmediaMappingAO1P1TypeAuthenticationProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","ntlm","ntlmi","ntlmv2","ntlmv2i","ntlmssp","ntlmsspi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmediaMappingAO1P1TypeAuthenticationProtocolPropEnum = append(vmediaMappingAO1P1TypeAuthenticationProtocolPropEnum, v)
	}
}

const (

	// VmediaMappingAO1P1AuthenticationProtocolNone captures enum value "none"
	VmediaMappingAO1P1AuthenticationProtocolNone string = "none"

	// VmediaMappingAO1P1AuthenticationProtocolNtlm captures enum value "ntlm"
	VmediaMappingAO1P1AuthenticationProtocolNtlm string = "ntlm"

	// VmediaMappingAO1P1AuthenticationProtocolNtlmi captures enum value "ntlmi"
	VmediaMappingAO1P1AuthenticationProtocolNtlmi string = "ntlmi"

	// VmediaMappingAO1P1AuthenticationProtocolNtlmv2 captures enum value "ntlmv2"
	VmediaMappingAO1P1AuthenticationProtocolNtlmv2 string = "ntlmv2"

	// VmediaMappingAO1P1AuthenticationProtocolNtlmv2i captures enum value "ntlmv2i"
	VmediaMappingAO1P1AuthenticationProtocolNtlmv2i string = "ntlmv2i"

	// VmediaMappingAO1P1AuthenticationProtocolNtlmssp captures enum value "ntlmssp"
	VmediaMappingAO1P1AuthenticationProtocolNtlmssp string = "ntlmssp"

	// VmediaMappingAO1P1AuthenticationProtocolNtlmsspi captures enum value "ntlmsspi"
	VmediaMappingAO1P1AuthenticationProtocolNtlmsspi string = "ntlmsspi"
)

// prop value enum
func (m *VmediaMappingAO1P1) validateAuthenticationProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vmediaMappingAO1P1TypeAuthenticationProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VmediaMappingAO1P1) validateAuthenticationProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationProtocolEnum("AuthenticationProtocol", "body", *m.AuthenticationProtocol); err != nil {
		return err
	}

	return nil
}

var vmediaMappingAO1P1TypeDeviceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cdd","hdd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmediaMappingAO1P1TypeDeviceTypePropEnum = append(vmediaMappingAO1P1TypeDeviceTypePropEnum, v)
	}
}

const (

	// VmediaMappingAO1P1DeviceTypeCdd captures enum value "cdd"
	VmediaMappingAO1P1DeviceTypeCdd string = "cdd"

	// VmediaMappingAO1P1DeviceTypeHdd captures enum value "hdd"
	VmediaMappingAO1P1DeviceTypeHdd string = "hdd"
)

// prop value enum
func (m *VmediaMappingAO1P1) validateDeviceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vmediaMappingAO1P1TypeDeviceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VmediaMappingAO1P1) validateDeviceType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeviceTypeEnum("DeviceType", "body", *m.DeviceType); err != nil {
		return err
	}

	return nil
}

var vmediaMappingAO1P1TypeMountProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nfs","cifs","http","https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmediaMappingAO1P1TypeMountProtocolPropEnum = append(vmediaMappingAO1P1TypeMountProtocolPropEnum, v)
	}
}

const (

	// VmediaMappingAO1P1MountProtocolNfs captures enum value "nfs"
	VmediaMappingAO1P1MountProtocolNfs string = "nfs"

	// VmediaMappingAO1P1MountProtocolCifs captures enum value "cifs"
	VmediaMappingAO1P1MountProtocolCifs string = "cifs"

	// VmediaMappingAO1P1MountProtocolHTTP captures enum value "http"
	VmediaMappingAO1P1MountProtocolHTTP string = "http"

	// VmediaMappingAO1P1MountProtocolHTTPS captures enum value "https"
	VmediaMappingAO1P1MountProtocolHTTPS string = "https"
)

// prop value enum
func (m *VmediaMappingAO1P1) validateMountProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vmediaMappingAO1P1TypeMountProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VmediaMappingAO1P1) validateMountProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.MountProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateMountProtocolEnum("MountProtocol", "body", *m.MountProtocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmediaMappingAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmediaMappingAO1P1) UnmarshalBinary(b []byte) error {
	var res VmediaMappingAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
