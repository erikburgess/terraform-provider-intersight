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

// ConnectorExpectPromptAllOf Definition of the list of properties defined in 'connector.ExpectPrompt', excluding properties defined in parent classes.
type ConnectorExpectPromptAllOf struct {
	// The regex of the expect prompt of the interactive command.
	Expect *string `json:"Expect,omitempty"`
	// The timeout for the expect prompt while executing interactive command. If timeout is not set a default of 60 seconds will be used.
	ExpectTimeout *int64 `json:"ExpectTimeout,omitempty"`
	// The answer string to the expect prompt.
	Send                 *string `json:"Send,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorExpectPromptAllOf ConnectorExpectPromptAllOf

// NewConnectorExpectPromptAllOf instantiates a new ConnectorExpectPromptAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorExpectPromptAllOf() *ConnectorExpectPromptAllOf {
	this := ConnectorExpectPromptAllOf{}
	return &this
}

// NewConnectorExpectPromptAllOfWithDefaults instantiates a new ConnectorExpectPromptAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorExpectPromptAllOfWithDefaults() *ConnectorExpectPromptAllOf {
	this := ConnectorExpectPromptAllOf{}
	return &this
}

// GetExpect returns the Expect field value if set, zero value otherwise.
func (o *ConnectorExpectPromptAllOf) GetExpect() string {
	if o == nil || o.Expect == nil {
		var ret string
		return ret
	}
	return *o.Expect
}

// GetExpectOk returns a tuple with the Expect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPromptAllOf) GetExpectOk() (*string, bool) {
	if o == nil || o.Expect == nil {
		return nil, false
	}
	return o.Expect, true
}

// HasExpect returns a boolean if a field has been set.
func (o *ConnectorExpectPromptAllOf) HasExpect() bool {
	if o != nil && o.Expect != nil {
		return true
	}

	return false
}

// SetExpect gets a reference to the given string and assigns it to the Expect field.
func (o *ConnectorExpectPromptAllOf) SetExpect(v string) {
	o.Expect = &v
}

// GetExpectTimeout returns the ExpectTimeout field value if set, zero value otherwise.
func (o *ConnectorExpectPromptAllOf) GetExpectTimeout() int64 {
	if o == nil || o.ExpectTimeout == nil {
		var ret int64
		return ret
	}
	return *o.ExpectTimeout
}

// GetExpectTimeoutOk returns a tuple with the ExpectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPromptAllOf) GetExpectTimeoutOk() (*int64, bool) {
	if o == nil || o.ExpectTimeout == nil {
		return nil, false
	}
	return o.ExpectTimeout, true
}

// HasExpectTimeout returns a boolean if a field has been set.
func (o *ConnectorExpectPromptAllOf) HasExpectTimeout() bool {
	if o != nil && o.ExpectTimeout != nil {
		return true
	}

	return false
}

// SetExpectTimeout gets a reference to the given int64 and assigns it to the ExpectTimeout field.
func (o *ConnectorExpectPromptAllOf) SetExpectTimeout(v int64) {
	o.ExpectTimeout = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *ConnectorExpectPromptAllOf) GetSend() string {
	if o == nil || o.Send == nil {
		var ret string
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorExpectPromptAllOf) GetSendOk() (*string, bool) {
	if o == nil || o.Send == nil {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *ConnectorExpectPromptAllOf) HasSend() bool {
	if o != nil && o.Send != nil {
		return true
	}

	return false
}

// SetSend gets a reference to the given string and assigns it to the Send field.
func (o *ConnectorExpectPromptAllOf) SetSend(v string) {
	o.Send = &v
}

func (o ConnectorExpectPromptAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expect != nil {
		toSerialize["Expect"] = o.Expect
	}
	if o.ExpectTimeout != nil {
		toSerialize["ExpectTimeout"] = o.ExpectTimeout
	}
	if o.Send != nil {
		toSerialize["Send"] = o.Send
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorExpectPromptAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorExpectPromptAllOf := _ConnectorExpectPromptAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorExpectPromptAllOf); err == nil {
		*o = ConnectorExpectPromptAllOf(varConnectorExpectPromptAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Expect")
		delete(additionalProperties, "ExpectTimeout")
		delete(additionalProperties, "Send")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorExpectPromptAllOf struct {
	value *ConnectorExpectPromptAllOf
	isSet bool
}

func (v NullableConnectorExpectPromptAllOf) Get() *ConnectorExpectPromptAllOf {
	return v.value
}

func (v *NullableConnectorExpectPromptAllOf) Set(val *ConnectorExpectPromptAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorExpectPromptAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorExpectPromptAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorExpectPromptAllOf(val *ConnectorExpectPromptAllOf) *NullableConnectorExpectPromptAllOf {
	return &NullableConnectorExpectPromptAllOf{value: val, isSet: true}
}

func (v NullableConnectorExpectPromptAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorExpectPromptAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
