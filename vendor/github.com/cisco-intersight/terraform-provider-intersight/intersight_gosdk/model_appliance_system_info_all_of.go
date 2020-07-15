/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ApplianceSystemInfoAllOf Definition of the list of properties defined in 'appliance.SystemInfo', excluding properties defined in parent classes.
type ApplianceSystemInfoAllOf struct {
	// Connection state of the Intersight Appliance to the Intersight.
	CloudConnStatus *string `json:"CloudConnStatus,omitempty"`
	// Current running deployment size for the Intersight Appliance cluster. Eg. small, medium, large etc.
	DeploymentSize *string `json:"DeploymentSize,omitempty"`
	// Publicly accessible FQDN or IP address of the Intersight Appliance.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates that the setup initialization process has been completed.
	InitDone *bool `json:"InitDone,omitempty"`
	// Operational status of the Intersight Appliance cluster.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
	// SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance.
	SerialId *string `json:"SerialId,omitempty"`
	// Timezone of the Intersight Appliance.
	TimeZone *string `json:"TimeZone,omitempty"`
	// Current software version of the Intersight Appliance.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceSystemInfoAllOf ApplianceSystemInfoAllOf

// NewApplianceSystemInfoAllOf instantiates a new ApplianceSystemInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceSystemInfoAllOf() *ApplianceSystemInfoAllOf {
	this := ApplianceSystemInfoAllOf{}
	var cloudConnStatus string = ""
	this.CloudConnStatus = &cloudConnStatus
	var operationalStatus string = "Unknown"
	this.OperationalStatus = &operationalStatus
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	return &this
}

// NewApplianceSystemInfoAllOfWithDefaults instantiates a new ApplianceSystemInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceSystemInfoAllOfWithDefaults() *ApplianceSystemInfoAllOf {
	this := ApplianceSystemInfoAllOf{}
	var cloudConnStatus string = ""
	this.CloudConnStatus = &cloudConnStatus
	var operationalStatus string = "Unknown"
	this.OperationalStatus = &operationalStatus
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	return &this
}

// GetCloudConnStatus returns the CloudConnStatus field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetCloudConnStatus() string {
	if o == nil || o.CloudConnStatus == nil {
		var ret string
		return ret
	}
	return *o.CloudConnStatus
}

// GetCloudConnStatusOk returns a tuple with the CloudConnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetCloudConnStatusOk() (*string, bool) {
	if o == nil || o.CloudConnStatus == nil {
		return nil, false
	}
	return o.CloudConnStatus, true
}

// HasCloudConnStatus returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasCloudConnStatus() bool {
	if o != nil && o.CloudConnStatus != nil {
		return true
	}

	return false
}

// SetCloudConnStatus gets a reference to the given string and assigns it to the CloudConnStatus field.
func (o *ApplianceSystemInfoAllOf) SetCloudConnStatus(v string) {
	o.CloudConnStatus = &v
}

// GetDeploymentSize returns the DeploymentSize field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetDeploymentSize() string {
	if o == nil || o.DeploymentSize == nil {
		var ret string
		return ret
	}
	return *o.DeploymentSize
}

// GetDeploymentSizeOk returns a tuple with the DeploymentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetDeploymentSizeOk() (*string, bool) {
	if o == nil || o.DeploymentSize == nil {
		return nil, false
	}
	return o.DeploymentSize, true
}

// HasDeploymentSize returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasDeploymentSize() bool {
	if o != nil && o.DeploymentSize != nil {
		return true
	}

	return false
}

// SetDeploymentSize gets a reference to the given string and assigns it to the DeploymentSize field.
func (o *ApplianceSystemInfoAllOf) SetDeploymentSize(v string) {
	o.DeploymentSize = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceSystemInfoAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetInitDone returns the InitDone field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetInitDone() bool {
	if o == nil || o.InitDone == nil {
		var ret bool
		return ret
	}
	return *o.InitDone
}

// GetInitDoneOk returns a tuple with the InitDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetInitDoneOk() (*bool, bool) {
	if o == nil || o.InitDone == nil {
		return nil, false
	}
	return o.InitDone, true
}

// HasInitDone returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasInitDone() bool {
	if o != nil && o.InitDone != nil {
		return true
	}

	return false
}

// SetInitDone gets a reference to the given bool and assigns it to the InitDone field.
func (o *ApplianceSystemInfoAllOf) SetInitDone(v bool) {
	o.InitDone = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceSystemInfoAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetSerialId returns the SerialId field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetSerialId() string {
	if o == nil || o.SerialId == nil {
		var ret string
		return ret
	}
	return *o.SerialId
}

// GetSerialIdOk returns a tuple with the SerialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetSerialIdOk() (*string, bool) {
	if o == nil || o.SerialId == nil {
		return nil, false
	}
	return o.SerialId, true
}

// HasSerialId returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasSerialId() bool {
	if o != nil && o.SerialId != nil {
		return true
	}

	return false
}

// SetSerialId gets a reference to the given string and assigns it to the SerialId field.
func (o *ApplianceSystemInfoAllOf) SetSerialId(v string) {
	o.SerialId = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ApplianceSystemInfoAllOf) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceSystemInfoAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemInfoAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceSystemInfoAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceSystemInfoAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o ApplianceSystemInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudConnStatus != nil {
		toSerialize["CloudConnStatus"] = o.CloudConnStatus
	}
	if o.DeploymentSize != nil {
		toSerialize["DeploymentSize"] = o.DeploymentSize
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.InitDone != nil {
		toSerialize["InitDone"] = o.InitDone
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.SerialId != nil {
		toSerialize["SerialId"] = o.SerialId
	}
	if o.TimeZone != nil {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceSystemInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceSystemInfoAllOf := _ApplianceSystemInfoAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceSystemInfoAllOf); err == nil {
		*o = ApplianceSystemInfoAllOf(varApplianceSystemInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CloudConnStatus")
		delete(additionalProperties, "DeploymentSize")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "InitDone")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "SerialId")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceSystemInfoAllOf struct {
	value *ApplianceSystemInfoAllOf
	isSet bool
}

func (v NullableApplianceSystemInfoAllOf) Get() *ApplianceSystemInfoAllOf {
	return v.value
}

func (v *NullableApplianceSystemInfoAllOf) Set(val *ApplianceSystemInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceSystemInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceSystemInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceSystemInfoAllOf(val *ApplianceSystemInfoAllOf) *NullableApplianceSystemInfoAllOf {
	return &NullableApplianceSystemInfoAllOf{value: val, isSet: true}
}

func (v NullableApplianceSystemInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceSystemInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
