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

// CommIpV6InterfaceAllOf Definition of the list of properties defined in 'comm.IpV6Interface', excluding properties defined in parent classes.
type CommIpV6InterfaceAllOf struct {
	// The IPv6 address of the default gateway.
	Gateway *string `json:"Gateway,omitempty"`
	// The IPv6 Address, represented as eight groups of four hexadecimal digits, separated by colons.
	IpAddress *string `json:"IpAddress,omitempty"`
	// The IPv6 Prefix, represented as eight groups of four hexadecimal digits, separated by colons.
	Prefix *string `json:"Prefix,omitempty"`
}

// NewCommIpV6InterfaceAllOf instantiates a new CommIpV6InterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommIpV6InterfaceAllOf() *CommIpV6InterfaceAllOf {
	this := CommIpV6InterfaceAllOf{}
	return &this
}

// NewCommIpV6InterfaceAllOfWithDefaults instantiates a new CommIpV6InterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommIpV6InterfaceAllOfWithDefaults() *CommIpV6InterfaceAllOf {
	this := CommIpV6InterfaceAllOf{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CommIpV6InterfaceAllOf) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV6InterfaceAllOf) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CommIpV6InterfaceAllOf) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CommIpV6InterfaceAllOf) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *CommIpV6InterfaceAllOf) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV6InterfaceAllOf) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CommIpV6InterfaceAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *CommIpV6InterfaceAllOf) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CommIpV6InterfaceAllOf) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV6InterfaceAllOf) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CommIpV6InterfaceAllOf) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CommIpV6InterfaceAllOf) SetPrefix(v string) {
	o.Prefix = &v
}

func (o CommIpV6InterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableCommIpV6InterfaceAllOf struct {
	value *CommIpV6InterfaceAllOf
	isSet bool
}

func (v NullableCommIpV6InterfaceAllOf) Get() *CommIpV6InterfaceAllOf {
	return v.value
}

func (v *NullableCommIpV6InterfaceAllOf) Set(val *CommIpV6InterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommIpV6InterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommIpV6InterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommIpV6InterfaceAllOf(val *CommIpV6InterfaceAllOf) *NullableCommIpV6InterfaceAllOf {
	return &NullableCommIpV6InterfaceAllOf{value: val, isSet: true}
}

func (v NullableCommIpV6InterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommIpV6InterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
