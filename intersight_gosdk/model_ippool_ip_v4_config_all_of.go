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

// IppoolIpV4ConfigAllOf Definition of the list of properties defined in 'ippool.IpV4Config', excluding properties defined in parent classes.
type IppoolIpV4ConfigAllOf struct {
	// IP address of the default IPv4 gateway.
	Gateway *string `json:"Gateway,omitempty"`
	// A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.
	Netmask *string `json:"Netmask,omitempty"`
	// IP Address of the primary Domain Name System (DNS) server.
	PrimaryDns *string `json:"PrimaryDns,omitempty"`
	// IP Address of the secondary Domain Name System (DNS) server.
	SecondaryDns         *string `json:"SecondaryDns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolIpV4ConfigAllOf IppoolIpV4ConfigAllOf

// NewIppoolIpV4ConfigAllOf instantiates a new IppoolIpV4ConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpV4ConfigAllOf() *IppoolIpV4ConfigAllOf {
	this := IppoolIpV4ConfigAllOf{}
	return &this
}

// NewIppoolIpV4ConfigAllOfWithDefaults instantiates a new IppoolIpV4ConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpV4ConfigAllOfWithDefaults() *IppoolIpV4ConfigAllOf {
	this := IppoolIpV4ConfigAllOf{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *IppoolIpV4ConfigAllOf) SetGateway(v string) {
	o.Gateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *IppoolIpV4ConfigAllOf) SetNetmask(v string) {
	o.Netmask = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetPrimaryDns() string {
	if o == nil || o.PrimaryDns == nil {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || o.PrimaryDns == nil {
		return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasPrimaryDns() bool {
	if o != nil && o.PrimaryDns != nil {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *IppoolIpV4ConfigAllOf) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetSecondaryDns() string {
	if o == nil || o.SecondaryDns == nil {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || o.SecondaryDns == nil {
		return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasSecondaryDns() bool {
	if o != nil && o.SecondaryDns != nil {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *IppoolIpV4ConfigAllOf) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

func (o IppoolIpV4ConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.PrimaryDns != nil {
		toSerialize["PrimaryDns"] = o.PrimaryDns
	}
	if o.SecondaryDns != nil {
		toSerialize["SecondaryDns"] = o.SecondaryDns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolIpV4ConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIppoolIpV4ConfigAllOf := _IppoolIpV4ConfigAllOf{}

	if err = json.Unmarshal(bytes, &varIppoolIpV4ConfigAllOf); err == nil {
		*o = IppoolIpV4ConfigAllOf(varIppoolIpV4ConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "PrimaryDns")
		delete(additionalProperties, "SecondaryDns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolIpV4ConfigAllOf struct {
	value *IppoolIpV4ConfigAllOf
	isSet bool
}

func (v NullableIppoolIpV4ConfigAllOf) Get() *IppoolIpV4ConfigAllOf {
	return v.value
}

func (v *NullableIppoolIpV4ConfigAllOf) Set(val *IppoolIpV4ConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpV4ConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpV4ConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpV4ConfigAllOf(val *IppoolIpV4ConfigAllOf) *NullableIppoolIpV4ConfigAllOf {
	return &NullableIppoolIpV4ConfigAllOf{value: val, isSet: true}
}

func (v NullableIppoolIpV4ConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpV4ConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
