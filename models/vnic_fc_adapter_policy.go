// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicFcAdapterPolicy Fibre Channel Adapter
//
// A Fibre Channel Adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. You can enable FCP Error Recovery, change the default settings of Queues and Interrupt handling for performance enhancement.
//
// swagger:model vnicFcAdapterPolicy
type VnicFcAdapterPolicy struct {
	PolicyAbstractPolicy

	// Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred.
	//
	ErrorDetectionTimeout int64 `json:"ErrorDetectionTimeout,omitempty"`

	// Fibre Channel Error Recovery Settings.
	//
	ErrorRecoverySettings *VnicFcErrorRecoverySettings `json:"ErrorRecoverySettings,omitempty"`

	// Fibre Channel Flogi Settings.
	//
	FlogiSettings *VnicFlogiSettings `json:"FlogiSettings,omitempty"`

	// Interrupt Settings for the virtual fibre channel interface.
	//
	InterruptSettings *VnicFcInterruptSettings `json:"InterruptSettings,omitempty"`

	// The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed.
	//
	IoThrottleCount int64 `json:"IoThrottleCount,omitempty"`

	// The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server.
	//
	LunCount int64 `json:"LunCount,omitempty"`

	// The number of commands that the HBA can send and receive in a single transmission per LUN.
	//
	LunQueueDepth int64 `json:"LunQueueDepth,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Fibre Channel Plogi Settings.
	//
	PlogiSettings *VnicPlogiSettings `json:"PlogiSettings,omitempty"`

	// Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated.
	//
	ResourceAllocationTimeout int64 `json:"ResourceAllocationTimeout,omitempty"`

	// Fibre Channel Receive Queue Settings.
	//
	RxQueueSettings *VnicFcQueueSettings `json:"RxQueueSettings,omitempty"`

	// SCSI Input/Output Queue Settings.
	//
	ScsiQueueSettings *VnicScsiQueueSettings `json:"ScsiQueueSettings,omitempty"`

	// Fibre Channel Transmit Queue Settings.
	//
	TxQueueSettings *VnicFcQueueSettings `json:"TxQueueSettings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicFcAdapterPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		ErrorDetectionTimeout int64 `json:"ErrorDetectionTimeout,omitempty"`

		ErrorRecoverySettings *VnicFcErrorRecoverySettings `json:"ErrorRecoverySettings,omitempty"`

		FlogiSettings *VnicFlogiSettings `json:"FlogiSettings,omitempty"`

		InterruptSettings *VnicFcInterruptSettings `json:"InterruptSettings,omitempty"`

		IoThrottleCount int64 `json:"IoThrottleCount,omitempty"`

		LunCount int64 `json:"LunCount,omitempty"`

		LunQueueDepth int64 `json:"LunQueueDepth,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		PlogiSettings *VnicPlogiSettings `json:"PlogiSettings,omitempty"`

		ResourceAllocationTimeout int64 `json:"ResourceAllocationTimeout,omitempty"`

		RxQueueSettings *VnicFcQueueSettings `json:"RxQueueSettings,omitempty"`

		ScsiQueueSettings *VnicScsiQueueSettings `json:"ScsiQueueSettings,omitempty"`

		TxQueueSettings *VnicFcQueueSettings `json:"TxQueueSettings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ErrorDetectionTimeout = dataAO1.ErrorDetectionTimeout

	m.ErrorRecoverySettings = dataAO1.ErrorRecoverySettings

	m.FlogiSettings = dataAO1.FlogiSettings

	m.InterruptSettings = dataAO1.InterruptSettings

	m.IoThrottleCount = dataAO1.IoThrottleCount

	m.LunCount = dataAO1.LunCount

	m.LunQueueDepth = dataAO1.LunQueueDepth

	m.Organization = dataAO1.Organization

	m.PlogiSettings = dataAO1.PlogiSettings

	m.ResourceAllocationTimeout = dataAO1.ResourceAllocationTimeout

	m.RxQueueSettings = dataAO1.RxQueueSettings

	m.ScsiQueueSettings = dataAO1.ScsiQueueSettings

	m.TxQueueSettings = dataAO1.TxQueueSettings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicFcAdapterPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ErrorDetectionTimeout int64 `json:"ErrorDetectionTimeout,omitempty"`

		ErrorRecoverySettings *VnicFcErrorRecoverySettings `json:"ErrorRecoverySettings,omitempty"`

		FlogiSettings *VnicFlogiSettings `json:"FlogiSettings,omitempty"`

		InterruptSettings *VnicFcInterruptSettings `json:"InterruptSettings,omitempty"`

		IoThrottleCount int64 `json:"IoThrottleCount,omitempty"`

		LunCount int64 `json:"LunCount,omitempty"`

		LunQueueDepth int64 `json:"LunQueueDepth,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		PlogiSettings *VnicPlogiSettings `json:"PlogiSettings,omitempty"`

		ResourceAllocationTimeout int64 `json:"ResourceAllocationTimeout,omitempty"`

		RxQueueSettings *VnicFcQueueSettings `json:"RxQueueSettings,omitempty"`

		ScsiQueueSettings *VnicScsiQueueSettings `json:"ScsiQueueSettings,omitempty"`

		TxQueueSettings *VnicFcQueueSettings `json:"TxQueueSettings,omitempty"`
	}

	dataAO1.ErrorDetectionTimeout = m.ErrorDetectionTimeout

	dataAO1.ErrorRecoverySettings = m.ErrorRecoverySettings

	dataAO1.FlogiSettings = m.FlogiSettings

	dataAO1.InterruptSettings = m.InterruptSettings

	dataAO1.IoThrottleCount = m.IoThrottleCount

	dataAO1.LunCount = m.LunCount

	dataAO1.LunQueueDepth = m.LunQueueDepth

	dataAO1.Organization = m.Organization

	dataAO1.PlogiSettings = m.PlogiSettings

	dataAO1.ResourceAllocationTimeout = m.ResourceAllocationTimeout

	dataAO1.RxQueueSettings = m.RxQueueSettings

	dataAO1.ScsiQueueSettings = m.ScsiQueueSettings

	dataAO1.TxQueueSettings = m.TxQueueSettings

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic fc adapter policy
func (m *VnicFcAdapterPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorRecoverySettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlogiSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterruptSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlogiSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRxQueueSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScsiQueueSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxQueueSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VnicFcAdapterPolicy) validateErrorRecoverySettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorRecoverySettings) { // not required
		return nil
	}

	if m.ErrorRecoverySettings != nil {
		if err := m.ErrorRecoverySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ErrorRecoverySettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validateFlogiSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.FlogiSettings) { // not required
		return nil
	}

	if m.FlogiSettings != nil {
		if err := m.FlogiSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FlogiSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validateInterruptSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.InterruptSettings) { // not required
		return nil
	}

	if m.InterruptSettings != nil {
		if err := m.InterruptSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InterruptSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validatePlogiSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.PlogiSettings) { // not required
		return nil
	}

	if m.PlogiSettings != nil {
		if err := m.PlogiSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PlogiSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validateRxQueueSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.RxQueueSettings) { // not required
		return nil
	}

	if m.RxQueueSettings != nil {
		if err := m.RxQueueSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RxQueueSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validateScsiQueueSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ScsiQueueSettings) { // not required
		return nil
	}

	if m.ScsiQueueSettings != nil {
		if err := m.ScsiQueueSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ScsiQueueSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicFcAdapterPolicy) validateTxQueueSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.TxQueueSettings) { // not required
		return nil
	}

	if m.TxQueueSettings != nil {
		if err := m.TxQueueSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TxQueueSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicFcAdapterPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicFcAdapterPolicy) UnmarshalBinary(b []byte) error {
	var res VnicFcAdapterPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
