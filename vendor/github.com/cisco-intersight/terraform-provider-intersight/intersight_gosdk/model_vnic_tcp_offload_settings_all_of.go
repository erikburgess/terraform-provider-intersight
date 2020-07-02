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

// VnicTcpOffloadSettingsAllOf Definition of the list of properties defined in 'vnic.TcpOffloadSettings', excluding properties defined in parent classes.
type VnicTcpOffloadSettingsAllOf struct {
	// Enables the reassembly of segmented packets in hardware before sending them to the CPU.
	LargeReceive *bool `json:"LargeReceive,omitempty"`
	// Enables the CPU to send large packets to the hardware for segmentation.
	LargeSend *bool `json:"LargeSend,omitempty"`
	// When enabled, the CPU sends all packet checksums to the hardware for validation.
	RxChecksum *bool `json:"RxChecksum,omitempty"`
	// When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.
	TxChecksum *bool `json:"TxChecksum,omitempty"`
}

// NewVnicTcpOffloadSettingsAllOf instantiates a new VnicTcpOffloadSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicTcpOffloadSettingsAllOf() *VnicTcpOffloadSettingsAllOf {
	this := VnicTcpOffloadSettingsAllOf{}
	return &this
}

// NewVnicTcpOffloadSettingsAllOfWithDefaults instantiates a new VnicTcpOffloadSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicTcpOffloadSettingsAllOfWithDefaults() *VnicTcpOffloadSettingsAllOf {
	this := VnicTcpOffloadSettingsAllOf{}
	return &this
}

// GetLargeReceive returns the LargeReceive field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettingsAllOf) GetLargeReceive() bool {
	if o == nil || o.LargeReceive == nil {
		var ret bool
		return ret
	}
	return *o.LargeReceive
}

// GetLargeReceiveOk returns a tuple with the LargeReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettingsAllOf) GetLargeReceiveOk() (*bool, bool) {
	if o == nil || o.LargeReceive == nil {
		return nil, false
	}
	return o.LargeReceive, true
}

// HasLargeReceive returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettingsAllOf) HasLargeReceive() bool {
	if o != nil && o.LargeReceive != nil {
		return true
	}

	return false
}

// SetLargeReceive gets a reference to the given bool and assigns it to the LargeReceive field.
func (o *VnicTcpOffloadSettingsAllOf) SetLargeReceive(v bool) {
	o.LargeReceive = &v
}

// GetLargeSend returns the LargeSend field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettingsAllOf) GetLargeSend() bool {
	if o == nil || o.LargeSend == nil {
		var ret bool
		return ret
	}
	return *o.LargeSend
}

// GetLargeSendOk returns a tuple with the LargeSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettingsAllOf) GetLargeSendOk() (*bool, bool) {
	if o == nil || o.LargeSend == nil {
		return nil, false
	}
	return o.LargeSend, true
}

// HasLargeSend returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettingsAllOf) HasLargeSend() bool {
	if o != nil && o.LargeSend != nil {
		return true
	}

	return false
}

// SetLargeSend gets a reference to the given bool and assigns it to the LargeSend field.
func (o *VnicTcpOffloadSettingsAllOf) SetLargeSend(v bool) {
	o.LargeSend = &v
}

// GetRxChecksum returns the RxChecksum field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettingsAllOf) GetRxChecksum() bool {
	if o == nil || o.RxChecksum == nil {
		var ret bool
		return ret
	}
	return *o.RxChecksum
}

// GetRxChecksumOk returns a tuple with the RxChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettingsAllOf) GetRxChecksumOk() (*bool, bool) {
	if o == nil || o.RxChecksum == nil {
		return nil, false
	}
	return o.RxChecksum, true
}

// HasRxChecksum returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettingsAllOf) HasRxChecksum() bool {
	if o != nil && o.RxChecksum != nil {
		return true
	}

	return false
}

// SetRxChecksum gets a reference to the given bool and assigns it to the RxChecksum field.
func (o *VnicTcpOffloadSettingsAllOf) SetRxChecksum(v bool) {
	o.RxChecksum = &v
}

// GetTxChecksum returns the TxChecksum field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettingsAllOf) GetTxChecksum() bool {
	if o == nil || o.TxChecksum == nil {
		var ret bool
		return ret
	}
	return *o.TxChecksum
}

// GetTxChecksumOk returns a tuple with the TxChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettingsAllOf) GetTxChecksumOk() (*bool, bool) {
	if o == nil || o.TxChecksum == nil {
		return nil, false
	}
	return o.TxChecksum, true
}

// HasTxChecksum returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettingsAllOf) HasTxChecksum() bool {
	if o != nil && o.TxChecksum != nil {
		return true
	}

	return false
}

// SetTxChecksum gets a reference to the given bool and assigns it to the TxChecksum field.
func (o *VnicTcpOffloadSettingsAllOf) SetTxChecksum(v bool) {
	o.TxChecksum = &v
}

func (o VnicTcpOffloadSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LargeReceive != nil {
		toSerialize["LargeReceive"] = o.LargeReceive
	}
	if o.LargeSend != nil {
		toSerialize["LargeSend"] = o.LargeSend
	}
	if o.RxChecksum != nil {
		toSerialize["RxChecksum"] = o.RxChecksum
	}
	if o.TxChecksum != nil {
		toSerialize["TxChecksum"] = o.TxChecksum
	}
	return json.Marshal(toSerialize)
}

type NullableVnicTcpOffloadSettingsAllOf struct {
	value *VnicTcpOffloadSettingsAllOf
	isSet bool
}

func (v NullableVnicTcpOffloadSettingsAllOf) Get() *VnicTcpOffloadSettingsAllOf {
	return v.value
}

func (v *NullableVnicTcpOffloadSettingsAllOf) Set(val *VnicTcpOffloadSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicTcpOffloadSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicTcpOffloadSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicTcpOffloadSettingsAllOf(val *VnicTcpOffloadSettingsAllOf) *NullableVnicTcpOffloadSettingsAllOf {
	return &NullableVnicTcpOffloadSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicTcpOffloadSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicTcpOffloadSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}