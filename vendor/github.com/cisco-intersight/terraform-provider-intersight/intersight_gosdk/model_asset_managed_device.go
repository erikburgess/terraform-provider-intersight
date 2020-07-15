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
	"reflect"
	"strings"
)

// AssetManagedDevice Attributes for Managed Device in Intersight and it maintains the relationship to the Intersight Assist Device. Once added, Device Connector for the Managed Device type is started on the Intersight Assist and status related to it is maintained.
type AssetManagedDevice struct {
	MoBaseMo
	Credential *CommCredential `json:"Credential,omitempty"`
	// Type of the Device such as VMware, Pure Storage supported by Intersight Assist.
	DeviceType *string `json:"DeviceType,omitempty"`
	// Ignore Certificates with protocol https for connecting to the Managed Device. It is not used for other protocols.
	IgnoreCert *bool `json:"IgnoreCert,omitempty"`
	// Device is Enabled/Disabled.
	IsEnabled *bool `json:"IsEnabled,omitempty"`
	// Management address of the device. It can be IPv4, IPv6 or FQDN.
	ManagementAddress *string `json:"ManagementAddress,omitempty"`
	// Port to use for connecting to the Managed Device. Port is optional. If not specified, default port for protocol is used.
	Port *int64 `json:"Port,omitempty"`
	// Protocol to use for connecting to the Managed Device.
	Protocol               *string                              `json:"Protocol,omitempty"`
	Status                 *AssetManagedDeviceStatus            `json:"Status,omitempty"`
	Account                *IamAccountRelationship              `json:"Account,omitempty"`
	DeviceConnectorManager *AssetDeviceRegistrationRelationship `json:"DeviceConnectorManager,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	WorkflowInfo           *WorkflowWorkflowInfoRelationship    `json:"WorkflowInfo,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _AssetManagedDevice AssetManagedDevice

// NewAssetManagedDevice instantiates a new AssetManagedDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetManagedDevice() *AssetManagedDevice {
	this := AssetManagedDevice{}
	var deviceType string = ""
	this.DeviceType = &deviceType
	var protocol string = "https"
	this.Protocol = &protocol
	return &this
}

// NewAssetManagedDeviceWithDefaults instantiates a new AssetManagedDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetManagedDeviceWithDefaults() *AssetManagedDevice {
	this := AssetManagedDevice{}
	var deviceType string = ""
	this.DeviceType = &deviceType
	var protocol string = "https"
	this.Protocol = &protocol
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetCredential() CommCredential {
	if o == nil || o.Credential == nil {
		var ret CommCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetCredentialOk() (*CommCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given CommCredential and assigns it to the Credential field.
func (o *AssetManagedDevice) SetCredential(v CommCredential) {
	o.Credential = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *AssetManagedDevice) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetIgnoreCert returns the IgnoreCert field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetIgnoreCert() bool {
	if o == nil || o.IgnoreCert == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreCert
}

// GetIgnoreCertOk returns a tuple with the IgnoreCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetIgnoreCertOk() (*bool, bool) {
	if o == nil || o.IgnoreCert == nil {
		return nil, false
	}
	return o.IgnoreCert, true
}

// HasIgnoreCert returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasIgnoreCert() bool {
	if o != nil && o.IgnoreCert != nil {
		return true
	}

	return false
}

// SetIgnoreCert gets a reference to the given bool and assigns it to the IgnoreCert field.
func (o *AssetManagedDevice) SetIgnoreCert(v bool) {
	o.IgnoreCert = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *AssetManagedDevice) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetManagementAddress() string {
	if o == nil || o.ManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetManagementAddressOk() (*string, bool) {
	if o == nil || o.ManagementAddress == nil {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasManagementAddress() bool {
	if o != nil && o.ManagementAddress != nil {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *AssetManagedDevice) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *AssetManagedDevice) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AssetManagedDevice) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetStatus() AssetManagedDeviceStatus {
	if o == nil || o.Status == nil {
		var ret AssetManagedDeviceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetStatusOk() (*AssetManagedDeviceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AssetManagedDeviceStatus and assigns it to the Status field.
func (o *AssetManagedDevice) SetStatus(v AssetManagedDeviceStatus) {
	o.Status = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetManagedDevice) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetDeviceConnectorManager returns the DeviceConnectorManager field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetDeviceConnectorManager() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceConnectorManager == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceConnectorManager
}

// GetDeviceConnectorManagerOk returns a tuple with the DeviceConnectorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetDeviceConnectorManagerOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceConnectorManager == nil {
		return nil, false
	}
	return o.DeviceConnectorManager, true
}

// HasDeviceConnectorManager returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasDeviceConnectorManager() bool {
	if o != nil && o.DeviceConnectorManager != nil {
		return true
	}

	return false
}

// SetDeviceConnectorManager gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceConnectorManager field.
func (o *AssetManagedDevice) SetDeviceConnectorManager(v AssetDeviceRegistrationRelationship) {
	o.DeviceConnectorManager = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetManagedDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *AssetManagedDevice) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetManagedDevice) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *AssetManagedDevice) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *AssetManagedDevice) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o AssetManagedDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Credential != nil {
		toSerialize["Credential"] = o.Credential
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.IgnoreCert != nil {
		toSerialize["IgnoreCert"] = o.IgnoreCert
	}
	if o.IsEnabled != nil {
		toSerialize["IsEnabled"] = o.IsEnabled
	}
	if o.ManagementAddress != nil {
		toSerialize["ManagementAddress"] = o.ManagementAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.DeviceConnectorManager != nil {
		toSerialize["DeviceConnectorManager"] = o.DeviceConnectorManager
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetManagedDevice) UnmarshalJSON(bytes []byte) (err error) {
	type AssetManagedDeviceWithoutEmbeddedStruct struct {
		Credential *CommCredential `json:"Credential,omitempty"`
		// Type of the Device such as VMware, Pure Storage supported by Intersight Assist.
		DeviceType *string `json:"DeviceType,omitempty"`
		// Ignore Certificates with protocol https for connecting to the Managed Device. It is not used for other protocols.
		IgnoreCert *bool `json:"IgnoreCert,omitempty"`
		// Device is Enabled/Disabled.
		IsEnabled *bool `json:"IsEnabled,omitempty"`
		// Management address of the device. It can be IPv4, IPv6 or FQDN.
		ManagementAddress *string `json:"ManagementAddress,omitempty"`
		// Port to use for connecting to the Managed Device. Port is optional. If not specified, default port for protocol is used.
		Port *int64 `json:"Port,omitempty"`
		// Protocol to use for connecting to the Managed Device.
		Protocol               *string                              `json:"Protocol,omitempty"`
		Status                 *AssetManagedDeviceStatus            `json:"Status,omitempty"`
		Account                *IamAccountRelationship              `json:"Account,omitempty"`
		DeviceConnectorManager *AssetDeviceRegistrationRelationship `json:"DeviceConnectorManager,omitempty"`
		RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		WorkflowInfo           *WorkflowWorkflowInfoRelationship    `json:"WorkflowInfo,omitempty"`
	}

	varAssetManagedDeviceWithoutEmbeddedStruct := AssetManagedDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetManagedDeviceWithoutEmbeddedStruct)
	if err == nil {
		varAssetManagedDevice := _AssetManagedDevice{}
		varAssetManagedDevice.Credential = varAssetManagedDeviceWithoutEmbeddedStruct.Credential
		varAssetManagedDevice.DeviceType = varAssetManagedDeviceWithoutEmbeddedStruct.DeviceType
		varAssetManagedDevice.IgnoreCert = varAssetManagedDeviceWithoutEmbeddedStruct.IgnoreCert
		varAssetManagedDevice.IsEnabled = varAssetManagedDeviceWithoutEmbeddedStruct.IsEnabled
		varAssetManagedDevice.ManagementAddress = varAssetManagedDeviceWithoutEmbeddedStruct.ManagementAddress
		varAssetManagedDevice.Port = varAssetManagedDeviceWithoutEmbeddedStruct.Port
		varAssetManagedDevice.Protocol = varAssetManagedDeviceWithoutEmbeddedStruct.Protocol
		varAssetManagedDevice.Status = varAssetManagedDeviceWithoutEmbeddedStruct.Status
		varAssetManagedDevice.Account = varAssetManagedDeviceWithoutEmbeddedStruct.Account
		varAssetManagedDevice.DeviceConnectorManager = varAssetManagedDeviceWithoutEmbeddedStruct.DeviceConnectorManager
		varAssetManagedDevice.RegisteredDevice = varAssetManagedDeviceWithoutEmbeddedStruct.RegisteredDevice
		varAssetManagedDevice.WorkflowInfo = varAssetManagedDeviceWithoutEmbeddedStruct.WorkflowInfo
		*o = AssetManagedDevice(varAssetManagedDevice)
	} else {
		return err
	}

	varAssetManagedDevice := _AssetManagedDevice{}

	err = json.Unmarshal(bytes, &varAssetManagedDevice)
	if err == nil {
		o.MoBaseMo = varAssetManagedDevice.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Credential")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "IgnoreCert")
		delete(additionalProperties, "IsEnabled")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "DeviceConnectorManager")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "WorkflowInfo")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetManagedDevice struct {
	value *AssetManagedDevice
	isSet bool
}

func (v NullableAssetManagedDevice) Get() *AssetManagedDevice {
	return v.value
}

func (v *NullableAssetManagedDevice) Set(val *AssetManagedDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetManagedDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetManagedDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetManagedDevice(val *AssetManagedDevice) *NullableAssetManagedDevice {
	return &NullableAssetManagedDevice{value: val, isSet: true}
}

func (v NullableAssetManagedDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetManagedDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
