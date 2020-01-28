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

// VnicFcErrorRecoverySettings Error Recovery Settings
//
// Fibre Channel Error Recovery Settings.
//
// swagger:model vnicFcErrorRecoverySettings
type VnicFcErrorRecoverySettings struct {
	MoBaseComplexType

	VnicFcErrorRecoverySettingsAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicFcErrorRecoverySettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VnicFcErrorRecoverySettingsAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VnicFcErrorRecoverySettingsAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicFcErrorRecoverySettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VnicFcErrorRecoverySettingsAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic fc error recovery settings
func (m *VnicFcErrorRecoverySettings) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VnicFcErrorRecoverySettingsAO1P1
	if err := m.VnicFcErrorRecoverySettingsAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcErrorRecoverySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcErrorRecoverySettings) UnmarshalBinary(b []byte) error {
	var res VnicFcErrorRecoverySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VnicFcErrorRecoverySettingsAO1P1 vnic fc error recovery settings a o1 p1
// swagger:model VnicFcErrorRecoverySettingsAO1P1
type VnicFcErrorRecoverySettingsAO1P1 struct {

	// Enables Fibre Channel Error recovery.
	//
	Enabled *bool `json:"Enabled,omitempty"`

	// The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable.
	//
	IoRetryCount int64 `json:"IoRetryCount,omitempty"`

	// The number of seconds the adapter waits before aborting the pending command and resending the same IO request.
	//
	IoRetryTimeout int64 `json:"IoRetryTimeout,omitempty"`

	// The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost.
	//
	LinkDownTimeout int64 `json:"LinkDownTimeout,omitempty"`

	// The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds.
	//
	PortDownTimeout int64 `json:"PortDownTimeout,omitempty"`

	// vnic fc error recovery settings a o1 p1
	VnicFcErrorRecoverySettingsAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VnicFcErrorRecoverySettingsAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Enables Fibre Channel Error recovery.
		//
		Enabled *bool `json:"Enabled,omitempty"`

		// The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable.
		//
		IoRetryCount int64 `json:"IoRetryCount,omitempty"`

		// The number of seconds the adapter waits before aborting the pending command and resending the same IO request.
		//
		IoRetryTimeout int64 `json:"IoRetryTimeout,omitempty"`

		// The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost.
		//
		LinkDownTimeout int64 `json:"LinkDownTimeout,omitempty"`

		// The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds.
		//
		PortDownTimeout int64 `json:"PortDownTimeout,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VnicFcErrorRecoverySettingsAO1P1

	rcv.Enabled = stage1.Enabled

	rcv.IoRetryCount = stage1.IoRetryCount

	rcv.IoRetryTimeout = stage1.IoRetryTimeout

	rcv.LinkDownTimeout = stage1.LinkDownTimeout

	rcv.PortDownTimeout = stage1.PortDownTimeout

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "Enabled")

	delete(stage2, "IoRetryCount")

	delete(stage2, "IoRetryTimeout")

	delete(stage2, "LinkDownTimeout")

	delete(stage2, "PortDownTimeout")

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
		m.VnicFcErrorRecoverySettingsAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VnicFcErrorRecoverySettingsAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Enables Fibre Channel Error recovery.
		//
		Enabled *bool `json:"Enabled,omitempty"`

		// The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable.
		//
		IoRetryCount int64 `json:"IoRetryCount,omitempty"`

		// The number of seconds the adapter waits before aborting the pending command and resending the same IO request.
		//
		IoRetryTimeout int64 `json:"IoRetryTimeout,omitempty"`

		// The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost.
		//
		LinkDownTimeout int64 `json:"LinkDownTimeout,omitempty"`

		// The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds.
		//
		PortDownTimeout int64 `json:"PortDownTimeout,omitempty"`
	}

	stage1.Enabled = m.Enabled

	stage1.IoRetryCount = m.IoRetryCount

	stage1.IoRetryTimeout = m.IoRetryTimeout

	stage1.LinkDownTimeout = m.LinkDownTimeout

	stage1.PortDownTimeout = m.PortDownTimeout

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VnicFcErrorRecoverySettingsAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VnicFcErrorRecoverySettingsAO1P1)
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

// Validate validates this vnic fc error recovery settings a o1 p1
func (m *VnicFcErrorRecoverySettingsAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcErrorRecoverySettingsAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcErrorRecoverySettingsAO1P1) UnmarshalBinary(b []byte) error {
	var res VnicFcErrorRecoverySettingsAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
