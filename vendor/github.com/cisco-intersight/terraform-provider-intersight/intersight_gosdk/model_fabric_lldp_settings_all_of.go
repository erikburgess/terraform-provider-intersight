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

// FabricLldpSettingsAllOf Definition of the list of properties defined in 'fabric.LldpSettings', excluding properties defined in parent classes.
type FabricLldpSettingsAllOf struct {
	// Determines if the LLDP frames can be received by an interface on the switch.
	ReceiveEnabled *bool `json:"ReceiveEnabled,omitempty"`
	// Determines if the LLDP frames can be transmitted by an interface on the switch.
	TransmitEnabled *bool `json:"TransmitEnabled,omitempty"`
}

// NewFabricLldpSettingsAllOf instantiates a new FabricLldpSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLldpSettingsAllOf() *FabricLldpSettingsAllOf {
	this := FabricLldpSettingsAllOf{}
	return &this
}

// NewFabricLldpSettingsAllOfWithDefaults instantiates a new FabricLldpSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLldpSettingsAllOfWithDefaults() *FabricLldpSettingsAllOf {
	this := FabricLldpSettingsAllOf{}
	return &this
}

// GetReceiveEnabled returns the ReceiveEnabled field value if set, zero value otherwise.
func (o *FabricLldpSettingsAllOf) GetReceiveEnabled() bool {
	if o == nil || o.ReceiveEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReceiveEnabled
}

// GetReceiveEnabledOk returns a tuple with the ReceiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLldpSettingsAllOf) GetReceiveEnabledOk() (*bool, bool) {
	if o == nil || o.ReceiveEnabled == nil {
		return nil, false
	}
	return o.ReceiveEnabled, true
}

// HasReceiveEnabled returns a boolean if a field has been set.
func (o *FabricLldpSettingsAllOf) HasReceiveEnabled() bool {
	if o != nil && o.ReceiveEnabled != nil {
		return true
	}

	return false
}

// SetReceiveEnabled gets a reference to the given bool and assigns it to the ReceiveEnabled field.
func (o *FabricLldpSettingsAllOf) SetReceiveEnabled(v bool) {
	o.ReceiveEnabled = &v
}

// GetTransmitEnabled returns the TransmitEnabled field value if set, zero value otherwise.
func (o *FabricLldpSettingsAllOf) GetTransmitEnabled() bool {
	if o == nil || o.TransmitEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TransmitEnabled
}

// GetTransmitEnabledOk returns a tuple with the TransmitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLldpSettingsAllOf) GetTransmitEnabledOk() (*bool, bool) {
	if o == nil || o.TransmitEnabled == nil {
		return nil, false
	}
	return o.TransmitEnabled, true
}

// HasTransmitEnabled returns a boolean if a field has been set.
func (o *FabricLldpSettingsAllOf) HasTransmitEnabled() bool {
	if o != nil && o.TransmitEnabled != nil {
		return true
	}

	return false
}

// SetTransmitEnabled gets a reference to the given bool and assigns it to the TransmitEnabled field.
func (o *FabricLldpSettingsAllOf) SetTransmitEnabled(v bool) {
	o.TransmitEnabled = &v
}

func (o FabricLldpSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReceiveEnabled != nil {
		toSerialize["ReceiveEnabled"] = o.ReceiveEnabled
	}
	if o.TransmitEnabled != nil {
		toSerialize["TransmitEnabled"] = o.TransmitEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableFabricLldpSettingsAllOf struct {
	value *FabricLldpSettingsAllOf
	isSet bool
}

func (v NullableFabricLldpSettingsAllOf) Get() *FabricLldpSettingsAllOf {
	return v.value
}

func (v *NullableFabricLldpSettingsAllOf) Set(val *FabricLldpSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLldpSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLldpSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLldpSettingsAllOf(val *FabricLldpSettingsAllOf) *NullableFabricLldpSettingsAllOf {
	return &NullableFabricLldpSettingsAllOf{value: val, isSet: true}
}

func (v NullableFabricLldpSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLldpSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
