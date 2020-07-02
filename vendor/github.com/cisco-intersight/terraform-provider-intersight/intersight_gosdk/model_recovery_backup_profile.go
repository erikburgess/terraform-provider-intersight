/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// RecoveryBackupProfile Backup profile to initiate on-demand or scheduled backups at end points.
type RecoveryBackupProfile struct {
	PolicyAbstractConfigProfile
	// Enables/Disables the schedule on the endpoint.
	Enabled        *bool                                     `json:"Enabled,omitempty"`
	BackupConfig   *RecoveryBackupConfigPolicyRelationship   `json:"BackupConfig,omitempty"`
	ConfigResult   *RecoveryConfigResultRelationship         `json:"ConfigResult,omitempty"`
	DeviceId       *AssetDeviceRegistrationRelationship      `json:"DeviceId,omitempty"`
	Organization   *OrganizationOrganizationRelationship     `json:"Organization,omitempty"`
	ScheduleConfig *RecoveryScheduleConfigPolicyRelationship `json:"ScheduleConfig,omitempty"`
}

// NewRecoveryBackupProfile instantiates a new RecoveryBackupProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryBackupProfile() *RecoveryBackupProfile {
	this := RecoveryBackupProfile{}
	return &this
}

// NewRecoveryBackupProfileWithDefaults instantiates a new RecoveryBackupProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryBackupProfileWithDefaults() *RecoveryBackupProfile {
	this := RecoveryBackupProfile{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RecoveryBackupProfile) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupProfile) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RecoveryBackupProfile) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RecoveryBackupProfile) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetBackupConfig returns the BackupConfig field value if set, zero value otherwise.
func (o *RecoveryBackupProfile) GetBackupConfig() RecoveryBackupConfigPolicyRelationship {
	if o == nil || o.BackupConfig == nil {
		var ret RecoveryBackupConfigPolicyRelationship
		return ret
	}
	return *o.BackupConfig
}

// GetBackupConfigOk returns a tuple with the BackupConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupProfile) GetBackupConfigOk() (*RecoveryBackupConfigPolicyRelationship, bool) {
	if o == nil || o.BackupConfig == nil {
		return nil, false
	}
	return o.BackupConfig, true
}

// HasBackupConfig returns a boolean if a field has been set.
func (o *RecoveryBackupProfile) HasBackupConfig() bool {
	if o != nil && o.BackupConfig != nil {
		return true
	}

	return false
}

// SetBackupConfig gets a reference to the given RecoveryBackupConfigPolicyRelationship and assigns it to the BackupConfig field.
func (o *RecoveryBackupProfile) SetBackupConfig(v RecoveryBackupConfigPolicyRelationship) {
	o.BackupConfig = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *RecoveryBackupProfile) GetConfigResult() RecoveryConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret RecoveryConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupProfile) GetConfigResultOk() (*RecoveryConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *RecoveryBackupProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given RecoveryConfigResultRelationship and assigns it to the ConfigResult field.
func (o *RecoveryBackupProfile) SetConfigResult(v RecoveryConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *RecoveryBackupProfile) GetDeviceId() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceId == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupProfile) GetDeviceIdOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *RecoveryBackupProfile) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceId field.
func (o *RecoveryBackupProfile) SetDeviceId(v AssetDeviceRegistrationRelationship) {
	o.DeviceId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryBackupProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryBackupProfile) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryBackupProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetScheduleConfig returns the ScheduleConfig field value if set, zero value otherwise.
func (o *RecoveryBackupProfile) GetScheduleConfig() RecoveryScheduleConfigPolicyRelationship {
	if o == nil || o.ScheduleConfig == nil {
		var ret RecoveryScheduleConfigPolicyRelationship
		return ret
	}
	return *o.ScheduleConfig
}

// GetScheduleConfigOk returns a tuple with the ScheduleConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupProfile) GetScheduleConfigOk() (*RecoveryScheduleConfigPolicyRelationship, bool) {
	if o == nil || o.ScheduleConfig == nil {
		return nil, false
	}
	return o.ScheduleConfig, true
}

// HasScheduleConfig returns a boolean if a field has been set.
func (o *RecoveryBackupProfile) HasScheduleConfig() bool {
	if o != nil && o.ScheduleConfig != nil {
		return true
	}

	return false
}

// SetScheduleConfig gets a reference to the given RecoveryScheduleConfigPolicyRelationship and assigns it to the ScheduleConfig field.
func (o *RecoveryBackupProfile) SetScheduleConfig(v RecoveryScheduleConfigPolicyRelationship) {
	o.ScheduleConfig = &v
}

func (o RecoveryBackupProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.BackupConfig != nil {
		toSerialize["BackupConfig"] = o.BackupConfig
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ScheduleConfig != nil {
		toSerialize["ScheduleConfig"] = o.ScheduleConfig
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryBackupProfile struct {
	value *RecoveryBackupProfile
	isSet bool
}

func (v NullableRecoveryBackupProfile) Get() *RecoveryBackupProfile {
	return v.value
}

func (v *NullableRecoveryBackupProfile) Set(val *RecoveryBackupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryBackupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryBackupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryBackupProfile(val *RecoveryBackupProfile) *NullableRecoveryBackupProfile {
	return &NullableRecoveryBackupProfile{value: val, isSet: true}
}

func (v NullableRecoveryBackupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryBackupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
