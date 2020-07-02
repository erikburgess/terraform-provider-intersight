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

// OsIpv4Configuration In case of static IPv4 configuration, IP address, netmask and gateway details are provided.
type OsIpv4Configuration struct {
	OsIpConfiguration
	IpV4Config *CommIpV4Interface `json:"IpV4Config,omitempty"`
}

// NewOsIpv4Configuration instantiates a new OsIpv4Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsIpv4Configuration() *OsIpv4Configuration {
	this := OsIpv4Configuration{}
	return &this
}

// NewOsIpv4ConfigurationWithDefaults instantiates a new OsIpv4Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsIpv4ConfigurationWithDefaults() *OsIpv4Configuration {
	this := OsIpv4Configuration{}
	return &this
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise.
func (o *OsIpv4Configuration) GetIpV4Config() CommIpV4Interface {
	if o == nil || o.IpV4Config == nil {
		var ret CommIpV4Interface
		return ret
	}
	return *o.IpV4Config
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsIpv4Configuration) GetIpV4ConfigOk() (*CommIpV4Interface, bool) {
	if o == nil || o.IpV4Config == nil {
		return nil, false
	}
	return o.IpV4Config, true
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *OsIpv4Configuration) HasIpV4Config() bool {
	if o != nil && o.IpV4Config != nil {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given CommIpV4Interface and assigns it to the IpV4Config field.
func (o *OsIpv4Configuration) SetIpV4Config(v CommIpV4Interface) {
	o.IpV4Config = &v
}

func (o OsIpv4Configuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsIpConfiguration, errOsIpConfiguration := json.Marshal(o.OsIpConfiguration)
	if errOsIpConfiguration != nil {
		return []byte{}, errOsIpConfiguration
	}
	errOsIpConfiguration = json.Unmarshal([]byte(serializedOsIpConfiguration), &toSerialize)
	if errOsIpConfiguration != nil {
		return []byte{}, errOsIpConfiguration
	}
	if o.IpV4Config != nil {
		toSerialize["IpV4Config"] = o.IpV4Config
	}
	return json.Marshal(toSerialize)
}

type NullableOsIpv4Configuration struct {
	value *OsIpv4Configuration
	isSet bool
}

func (v NullableOsIpv4Configuration) Get() *OsIpv4Configuration {
	return v.value
}

func (v *NullableOsIpv4Configuration) Set(val *OsIpv4Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableOsIpv4Configuration) IsSet() bool {
	return v.isSet
}

func (v *NullableOsIpv4Configuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsIpv4Configuration(val *OsIpv4Configuration) *NullableOsIpv4Configuration {
	return &NullableOsIpv4Configuration{value: val, isSet: true}
}

func (v NullableOsIpv4Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsIpv4Configuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
