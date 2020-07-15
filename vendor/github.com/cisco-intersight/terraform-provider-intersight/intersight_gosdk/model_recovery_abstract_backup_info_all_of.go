/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// RecoveryAbstractBackupInfoAllOf Definition of the list of properties defined in 'recovery.AbstractBackupInfo', excluding properties defined in parent classes.
type RecoveryAbstractBackupInfoAllOf struct {
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryAbstractBackupInfoAllOf RecoveryAbstractBackupInfoAllOf

// NewRecoveryAbstractBackupInfoAllOf instantiates a new RecoveryAbstractBackupInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryAbstractBackupInfoAllOf() *RecoveryAbstractBackupInfoAllOf {
	this := RecoveryAbstractBackupInfoAllOf{}
	return &this
}

// NewRecoveryAbstractBackupInfoAllOfWithDefaults instantiates a new RecoveryAbstractBackupInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryAbstractBackupInfoAllOfWithDefaults() *RecoveryAbstractBackupInfoAllOf {
	this := RecoveryAbstractBackupInfoAllOf{}
	return &this
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupInfoAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RecoveryAbstractBackupInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o RecoveryAbstractBackupInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryAbstractBackupInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecoveryAbstractBackupInfoAllOf := _RecoveryAbstractBackupInfoAllOf{}

	if err = json.Unmarshal(bytes, &varRecoveryAbstractBackupInfoAllOf); err == nil {
		*o = RecoveryAbstractBackupInfoAllOf(varRecoveryAbstractBackupInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryAbstractBackupInfoAllOf struct {
	value *RecoveryAbstractBackupInfoAllOf
	isSet bool
}

func (v NullableRecoveryAbstractBackupInfoAllOf) Get() *RecoveryAbstractBackupInfoAllOf {
	return v.value
}

func (v *NullableRecoveryAbstractBackupInfoAllOf) Set(val *RecoveryAbstractBackupInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryAbstractBackupInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryAbstractBackupInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryAbstractBackupInfoAllOf(val *RecoveryAbstractBackupInfoAllOf) *NullableRecoveryAbstractBackupInfoAllOf {
	return &NullableRecoveryAbstractBackupInfoAllOf{value: val, isSet: true}
}

func (v NullableRecoveryAbstractBackupInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryAbstractBackupInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
