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

// VnicEthInterruptSettings Interrupt settings for the virtual ethernet interface.
type VnicEthInterruptSettings struct {
	MoBaseComplexType
	// The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.
	CoalescingTime *int64 `json:"CoalescingTime,omitempty"`
	// Interrupt Coalescing Type. This can be one of the following:- MIN  — The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE — The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.
	CoalescingType *string `json:"CoalescingType,omitempty"`
	// The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources.
	Count *int64 `json:"Count,omitempty"`
	// Preferred driver interrupt mode. This can be one of the following:- MSIx — Message Signaled Interrupts (MSI) with the optional extension. MSI   — MSI only. INTx  — PCI INTx interrupts. MSIx is the recommended option.
	Mode *string `json:"Mode,omitempty"`
}

// NewVnicEthInterruptSettings instantiates a new VnicEthInterruptSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthInterruptSettings() *VnicEthInterruptSettings {
	this := VnicEthInterruptSettings{}
	var coalescingType string = "MIN"
	this.CoalescingType = &coalescingType
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// NewVnicEthInterruptSettingsWithDefaults instantiates a new VnicEthInterruptSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthInterruptSettingsWithDefaults() *VnicEthInterruptSettings {
	this := VnicEthInterruptSettings{}
	var coalescingType string = "MIN"
	this.CoalescingType = &coalescingType
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// GetCoalescingTime returns the CoalescingTime field value if set, zero value otherwise.
func (o *VnicEthInterruptSettings) GetCoalescingTime() int64 {
	if o == nil || o.CoalescingTime == nil {
		var ret int64
		return ret
	}
	return *o.CoalescingTime
}

// GetCoalescingTimeOk returns a tuple with the CoalescingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettings) GetCoalescingTimeOk() (*int64, bool) {
	if o == nil || o.CoalescingTime == nil {
		return nil, false
	}
	return o.CoalescingTime, true
}

// HasCoalescingTime returns a boolean if a field has been set.
func (o *VnicEthInterruptSettings) HasCoalescingTime() bool {
	if o != nil && o.CoalescingTime != nil {
		return true
	}

	return false
}

// SetCoalescingTime gets a reference to the given int64 and assigns it to the CoalescingTime field.
func (o *VnicEthInterruptSettings) SetCoalescingTime(v int64) {
	o.CoalescingTime = &v
}

// GetCoalescingType returns the CoalescingType field value if set, zero value otherwise.
func (o *VnicEthInterruptSettings) GetCoalescingType() string {
	if o == nil || o.CoalescingType == nil {
		var ret string
		return ret
	}
	return *o.CoalescingType
}

// GetCoalescingTypeOk returns a tuple with the CoalescingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettings) GetCoalescingTypeOk() (*string, bool) {
	if o == nil || o.CoalescingType == nil {
		return nil, false
	}
	return o.CoalescingType, true
}

// HasCoalescingType returns a boolean if a field has been set.
func (o *VnicEthInterruptSettings) HasCoalescingType() bool {
	if o != nil && o.CoalescingType != nil {
		return true
	}

	return false
}

// SetCoalescingType gets a reference to the given string and assigns it to the CoalescingType field.
func (o *VnicEthInterruptSettings) SetCoalescingType(v string) {
	o.CoalescingType = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicEthInterruptSettings) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettings) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicEthInterruptSettings) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicEthInterruptSettings) SetCount(v int64) {
	o.Count = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicEthInterruptSettings) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthInterruptSettings) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicEthInterruptSettings) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicEthInterruptSettings) SetMode(v string) {
	o.Mode = &v
}

func (o VnicEthInterruptSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.CoalescingTime != nil {
		toSerialize["CoalescingTime"] = o.CoalescingTime
	}
	if o.CoalescingType != nil {
		toSerialize["CoalescingType"] = o.CoalescingType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableVnicEthInterruptSettings struct {
	value *VnicEthInterruptSettings
	isSet bool
}

func (v NullableVnicEthInterruptSettings) Get() *VnicEthInterruptSettings {
	return v.value
}

func (v *NullableVnicEthInterruptSettings) Set(val *VnicEthInterruptSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthInterruptSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthInterruptSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthInterruptSettings(val *VnicEthInterruptSettings) *NullableVnicEthInterruptSettings {
	return &NullableVnicEthInterruptSettings{value: val, isSet: true}
}

func (v NullableVnicEthInterruptSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthInterruptSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
