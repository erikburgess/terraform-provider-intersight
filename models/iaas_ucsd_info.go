// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IaasUcsdInfo Iaas:Ucsd Info
//
// UCS Director accounts managed by Intersight.
//
// swagger:model iaasUcsdInfo
type IaasUcsdInfo struct {
	MoBaseMo

	// connector pack
	// Read Only: true
	ConnectorPack []*IaasConnectorPackRef `json:"ConnectorPack"`

	// Moid of the UCSD device connector's asset.DeviceRegistration.
	//
	// Read Only: true
	DeviceID string `json:"DeviceId,omitempty"`

	// device status
	// Read Only: true
	DeviceStatus []*IaasDeviceStatusRef `json:"DeviceStatus"`

	// Unique ID of UCSD getting registerd with Intersight.
	//
	// Read Only: true
	GUID string `json:"Guid,omitempty"`

	// The UCSD host name.
	//
	// Read Only: true
	HostName string `json:"HostName,omitempty"`

	// The UCSD IP address.
	//
	// Read Only: true
	IP string `json:"Ip,omitempty"`

	// license info
	// Read Only: true
	LicenseInfo *IaasLicenseInfoRef `json:"LicenseInfo,omitempty"`

	// Relationship to collection of MostRunTasks objects with cascade on delete of UcsdInfo object.
	//
	// Read Only: true
	MostRunTasks []*IaasMostRunTasksRef `json:"MostRunTasks"`

	// NodeType specifies if UCSD is deployed in Stand-alone or Multi Node.
	//
	// Read Only: true
	NodeType string `json:"NodeType,omitempty"`

	// The UCSD product name.
	//
	// Read Only: true
	ProductName string `json:"ProductName,omitempty"`

	// The UCSD product vendor.
	//
	// Read Only: true
	ProductVendor string `json:"ProductVendor,omitempty"`

	// The UCSD product/platform version.
	//
	// Read Only: true
	ProductVersion string `json:"ProductVersion,omitempty"`

	// registered device
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// The UCSD status. Possible values are Active, In-Active, Unknown.
	//
	// Read Only: true
	Status string `json:"Status,omitempty"`

	// ucsd managed infra
	// Read Only: true
	UcsdManagedInfra *IaasUcsdManagedInfraRef `json:"UcsdManagedInfra,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IaasUcsdInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ConnectorPack []*IaasConnectorPackRef `json:"ConnectorPack"`

		DeviceID string `json:"DeviceId,omitempty"`

		DeviceStatus []*IaasDeviceStatusRef `json:"DeviceStatus"`

		GUID string `json:"Guid,omitempty"`

		HostName string `json:"HostName,omitempty"`

		IP string `json:"Ip,omitempty"`

		LicenseInfo *IaasLicenseInfoRef `json:"LicenseInfo,omitempty"`

		MostRunTasks []*IaasMostRunTasksRef `json:"MostRunTasks"`

		NodeType string `json:"NodeType,omitempty"`

		ProductName string `json:"ProductName,omitempty"`

		ProductVendor string `json:"ProductVendor,omitempty"`

		ProductVersion string `json:"ProductVersion,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Status string `json:"Status,omitempty"`

		UcsdManagedInfra *IaasUcsdManagedInfraRef `json:"UcsdManagedInfra,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConnectorPack = dataAO1.ConnectorPack

	m.DeviceID = dataAO1.DeviceID

	m.DeviceStatus = dataAO1.DeviceStatus

	m.GUID = dataAO1.GUID

	m.HostName = dataAO1.HostName

	m.IP = dataAO1.IP

	m.LicenseInfo = dataAO1.LicenseInfo

	m.MostRunTasks = dataAO1.MostRunTasks

	m.NodeType = dataAO1.NodeType

	m.ProductName = dataAO1.ProductName

	m.ProductVendor = dataAO1.ProductVendor

	m.ProductVersion = dataAO1.ProductVersion

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Status = dataAO1.Status

	m.UcsdManagedInfra = dataAO1.UcsdManagedInfra

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IaasUcsdInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ConnectorPack []*IaasConnectorPackRef `json:"ConnectorPack"`

		DeviceID string `json:"DeviceId,omitempty"`

		DeviceStatus []*IaasDeviceStatusRef `json:"DeviceStatus"`

		GUID string `json:"Guid,omitempty"`

		HostName string `json:"HostName,omitempty"`

		IP string `json:"Ip,omitempty"`

		LicenseInfo *IaasLicenseInfoRef `json:"LicenseInfo,omitempty"`

		MostRunTasks []*IaasMostRunTasksRef `json:"MostRunTasks"`

		NodeType string `json:"NodeType,omitempty"`

		ProductName string `json:"ProductName,omitempty"`

		ProductVendor string `json:"ProductVendor,omitempty"`

		ProductVersion string `json:"ProductVersion,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Status string `json:"Status,omitempty"`

		UcsdManagedInfra *IaasUcsdManagedInfraRef `json:"UcsdManagedInfra,omitempty"`
	}

	dataAO1.ConnectorPack = m.ConnectorPack

	dataAO1.DeviceID = m.DeviceID

	dataAO1.DeviceStatus = m.DeviceStatus

	dataAO1.GUID = m.GUID

	dataAO1.HostName = m.HostName

	dataAO1.IP = m.IP

	dataAO1.LicenseInfo = m.LicenseInfo

	dataAO1.MostRunTasks = m.MostRunTasks

	dataAO1.NodeType = m.NodeType

	dataAO1.ProductName = m.ProductName

	dataAO1.ProductVendor = m.ProductVendor

	dataAO1.ProductVersion = m.ProductVersion

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Status = m.Status

	dataAO1.UcsdManagedInfra = m.UcsdManagedInfra

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iaas ucsd info
func (m *IaasUcsdInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectorPack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMostRunTasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUcsdManagedInfra(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IaasUcsdInfo) validateConnectorPack(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorPack) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectorPack); i++ {
		if swag.IsZero(m.ConnectorPack[i]) { // not required
			continue
		}

		if m.ConnectorPack[i] != nil {
			if err := m.ConnectorPack[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ConnectorPack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaasUcsdInfo) validateDeviceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceStatus); i++ {
		if swag.IsZero(m.DeviceStatus[i]) { // not required
			continue
		}

		if m.DeviceStatus[i] != nil {
			if err := m.DeviceStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeviceStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaasUcsdInfo) validateLicenseInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.LicenseInfo) { // not required
		return nil
	}

	if m.LicenseInfo != nil {
		if err := m.LicenseInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LicenseInfo")
			}
			return err
		}
	}

	return nil
}

func (m *IaasUcsdInfo) validateMostRunTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.MostRunTasks) { // not required
		return nil
	}

	for i := 0; i < len(m.MostRunTasks); i++ {
		if swag.IsZero(m.MostRunTasks[i]) { // not required
			continue
		}

		if m.MostRunTasks[i] != nil {
			if err := m.MostRunTasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("MostRunTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaasUcsdInfo) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

func (m *IaasUcsdInfo) validateUcsdManagedInfra(formats strfmt.Registry) error {

	if swag.IsZero(m.UcsdManagedInfra) { // not required
		return nil
	}

	if m.UcsdManagedInfra != nil {
		if err := m.UcsdManagedInfra.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UcsdManagedInfra")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IaasUcsdInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaasUcsdInfo) UnmarshalBinary(b []byte) error {
	var res IaasUcsdInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
