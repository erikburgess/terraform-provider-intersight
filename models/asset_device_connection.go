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

// AssetDeviceConnection Asset:Device Connection
//
// An abstract class that holds the details of a device connectors connection to Intersight.
//
// swagger:model assetDeviceConnection
type AssetDeviceConnection struct {
	MoBaseMo

	// The version of the connector API, describes the capability of the connector's framework.
	// If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.
	//
	//
	// Read Only: true
	APIVersion int64 `json:"ApiVersion,omitempty"`

	// The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.
	//
	// Read Only: true
	AppPartitionNumber int64 `json:"AppPartitionNumber,omitempty"`

	// The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.
	//
	// Read Only: true
	ConnectionID string `json:"ConnectionId,omitempty"`

	// If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with the cloud.
	//
	// Read Only: true
	ConnectionReason string `json:"ConnectionReason,omitempty"`

	// The status of the persistent connection between the device connector and Intersight.
	//
	// Read Only: true
	// Enum: [ Connected NotConnected Unclaimed]
	ConnectionStatus string `json:"ConnectionStatus,omitempty"`

	// The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between the cloud and device connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with the cloud.
	//
	// Read Only: true
	// Format: date-time
	ConnectionStatusLastChangeTime strfmt.DateTime `json:"ConnectionStatusLastChangeTime,omitempty"`

	// The version of the device connector running on the managed device.
	//
	// Read Only: true
	ConnectorVersion string `json:"ConnectorVersion,omitempty"`

	// The IP Address of the managed device as seen from the cloud at the time of registration.
	// This could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.
	//
	// Read Only: true
	DeviceExternalIPAddress string `json:"DeviceExternalIpAddress,omitempty"`

	// The name of the app which will proxy the messages to the device connector.
	//
	// Read Only: true
	ProxyApp string `json:"ProxyApp,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetDeviceConnection) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		APIVersion int64 `json:"ApiVersion,omitempty"`

		AppPartitionNumber int64 `json:"AppPartitionNumber,omitempty"`

		ConnectionID string `json:"ConnectionId,omitempty"`

		ConnectionReason string `json:"ConnectionReason,omitempty"`

		ConnectionStatus string `json:"ConnectionStatus,omitempty"`

		ConnectionStatusLastChangeTime strfmt.DateTime `json:"ConnectionStatusLastChangeTime,omitempty"`

		ConnectorVersion string `json:"ConnectorVersion,omitempty"`

		DeviceExternalIPAddress string `json:"DeviceExternalIpAddress,omitempty"`

		ProxyApp string `json:"ProxyApp,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.APIVersion = dataAO1.APIVersion

	m.AppPartitionNumber = dataAO1.AppPartitionNumber

	m.ConnectionID = dataAO1.ConnectionID

	m.ConnectionReason = dataAO1.ConnectionReason

	m.ConnectionStatus = dataAO1.ConnectionStatus

	m.ConnectionStatusLastChangeTime = dataAO1.ConnectionStatusLastChangeTime

	m.ConnectorVersion = dataAO1.ConnectorVersion

	m.DeviceExternalIPAddress = dataAO1.DeviceExternalIPAddress

	m.ProxyApp = dataAO1.ProxyApp

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetDeviceConnection) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		APIVersion int64 `json:"ApiVersion,omitempty"`

		AppPartitionNumber int64 `json:"AppPartitionNumber,omitempty"`

		ConnectionID string `json:"ConnectionId,omitempty"`

		ConnectionReason string `json:"ConnectionReason,omitempty"`

		ConnectionStatus string `json:"ConnectionStatus,omitempty"`

		ConnectionStatusLastChangeTime strfmt.DateTime `json:"ConnectionStatusLastChangeTime,omitempty"`

		ConnectorVersion string `json:"ConnectorVersion,omitempty"`

		DeviceExternalIPAddress string `json:"DeviceExternalIpAddress,omitempty"`

		ProxyApp string `json:"ProxyApp,omitempty"`
	}

	dataAO1.APIVersion = m.APIVersion

	dataAO1.AppPartitionNumber = m.AppPartitionNumber

	dataAO1.ConnectionID = m.ConnectionID

	dataAO1.ConnectionReason = m.ConnectionReason

	dataAO1.ConnectionStatus = m.ConnectionStatus

	dataAO1.ConnectionStatusLastChangeTime = m.ConnectionStatusLastChangeTime

	dataAO1.ConnectorVersion = m.ConnectorVersion

	dataAO1.DeviceExternalIPAddress = m.DeviceExternalIPAddress

	dataAO1.ProxyApp = m.ProxyApp

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset device connection
func (m *AssetDeviceConnection) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatusLastChangeTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var assetDeviceConnectionTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","Connected","NotConnected","Unclaimed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetDeviceConnectionTypeConnectionStatusPropEnum = append(assetDeviceConnectionTypeConnectionStatusPropEnum, v)
	}
}

// property enum
func (m *AssetDeviceConnection) validateConnectionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetDeviceConnectionTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetDeviceConnection) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("ConnectionStatus", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *AssetDeviceConnection) validateConnectionStatusLastChangeTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatusLastChangeTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ConnectionStatusLastChangeTime", "body", "date-time", m.ConnectionStatusLastChangeTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetDeviceConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetDeviceConnection) UnmarshalBinary(b []byte) error {
	var res AssetDeviceConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
