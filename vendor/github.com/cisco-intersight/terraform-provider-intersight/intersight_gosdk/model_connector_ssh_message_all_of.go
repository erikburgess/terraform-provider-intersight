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

// ConnectorSshMessageAllOf Definition of the list of properties defined in 'connector.SshMessage', excluding properties defined in parent classes.
type ConnectorSshMessageAllOf struct {
	ExpectPrompts *[]ConnectorExpectPrompt `json:"ExpectPrompts,omitempty"`
	// The operation to execute on a new or existing session.
	MsgType *int64 `json:"MsgType,omitempty"`
	// Unique id of session to route messages to.
	SessionId *string `json:"SessionId,omitempty"`
	// The regex of the secure shell prompt.
	ShellPrompt *string `json:"ShellPrompt,omitempty"`
	// Input to the SSH operation to be executed. e.g. file contents to write.
	Stream *string `json:"Stream,omitempty"`
	// The timeout for the ssh command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorSshMessageAllOf ConnectorSshMessageAllOf

// NewConnectorSshMessageAllOf instantiates a new ConnectorSshMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorSshMessageAllOf() *ConnectorSshMessageAllOf {
	this := ConnectorSshMessageAllOf{}
	return &this
}

// NewConnectorSshMessageAllOfWithDefaults instantiates a new ConnectorSshMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorSshMessageAllOfWithDefaults() *ConnectorSshMessageAllOf {
	this := ConnectorSshMessageAllOf{}
	return &this
}

// GetExpectPrompts returns the ExpectPrompts field value if set, zero value otherwise.
func (o *ConnectorSshMessageAllOf) GetExpectPrompts() []ConnectorExpectPrompt {
	if o == nil || o.ExpectPrompts == nil {
		var ret []ConnectorExpectPrompt
		return ret
	}
	return *o.ExpectPrompts
}

// GetExpectPromptsOk returns a tuple with the ExpectPrompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessageAllOf) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool) {
	if o == nil || o.ExpectPrompts == nil {
		return nil, false
	}
	return o.ExpectPrompts, true
}

// HasExpectPrompts returns a boolean if a field has been set.
func (o *ConnectorSshMessageAllOf) HasExpectPrompts() bool {
	if o != nil && o.ExpectPrompts != nil {
		return true
	}

	return false
}

// SetExpectPrompts gets a reference to the given []ConnectorExpectPrompt and assigns it to the ExpectPrompts field.
func (o *ConnectorSshMessageAllOf) SetExpectPrompts(v []ConnectorExpectPrompt) {
	o.ExpectPrompts = &v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *ConnectorSshMessageAllOf) GetMsgType() int64 {
	if o == nil || o.MsgType == nil {
		var ret int64
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessageAllOf) GetMsgTypeOk() (*int64, bool) {
	if o == nil || o.MsgType == nil {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *ConnectorSshMessageAllOf) HasMsgType() bool {
	if o != nil && o.MsgType != nil {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given int64 and assigns it to the MsgType field.
func (o *ConnectorSshMessageAllOf) SetMsgType(v int64) {
	o.MsgType = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ConnectorSshMessageAllOf) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessageAllOf) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ConnectorSshMessageAllOf) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ConnectorSshMessageAllOf) SetSessionId(v string) {
	o.SessionId = &v
}

// GetShellPrompt returns the ShellPrompt field value if set, zero value otherwise.
func (o *ConnectorSshMessageAllOf) GetShellPrompt() string {
	if o == nil || o.ShellPrompt == nil {
		var ret string
		return ret
	}
	return *o.ShellPrompt
}

// GetShellPromptOk returns a tuple with the ShellPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessageAllOf) GetShellPromptOk() (*string, bool) {
	if o == nil || o.ShellPrompt == nil {
		return nil, false
	}
	return o.ShellPrompt, true
}

// HasShellPrompt returns a boolean if a field has been set.
func (o *ConnectorSshMessageAllOf) HasShellPrompt() bool {
	if o != nil && o.ShellPrompt != nil {
		return true
	}

	return false
}

// SetShellPrompt gets a reference to the given string and assigns it to the ShellPrompt field.
func (o *ConnectorSshMessageAllOf) SetShellPrompt(v string) {
	o.ShellPrompt = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ConnectorSshMessageAllOf) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessageAllOf) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ConnectorSshMessageAllOf) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *ConnectorSshMessageAllOf) SetStream(v string) {
	o.Stream = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorSshMessageAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessageAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorSshMessageAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorSshMessageAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o ConnectorSshMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpectPrompts != nil {
		toSerialize["ExpectPrompts"] = o.ExpectPrompts
	}
	if o.MsgType != nil {
		toSerialize["MsgType"] = o.MsgType
	}
	if o.SessionId != nil {
		toSerialize["SessionId"] = o.SessionId
	}
	if o.ShellPrompt != nil {
		toSerialize["ShellPrompt"] = o.ShellPrompt
	}
	if o.Stream != nil {
		toSerialize["Stream"] = o.Stream
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorSshMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorSshMessageAllOf := _ConnectorSshMessageAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorSshMessageAllOf); err == nil {
		*o = ConnectorSshMessageAllOf(varConnectorSshMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ExpectPrompts")
		delete(additionalProperties, "MsgType")
		delete(additionalProperties, "SessionId")
		delete(additionalProperties, "ShellPrompt")
		delete(additionalProperties, "Stream")
		delete(additionalProperties, "Timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorSshMessageAllOf struct {
	value *ConnectorSshMessageAllOf
	isSet bool
}

func (v NullableConnectorSshMessageAllOf) Get() *ConnectorSshMessageAllOf {
	return v.value
}

func (v *NullableConnectorSshMessageAllOf) Set(val *ConnectorSshMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorSshMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorSshMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorSshMessageAllOf(val *ConnectorSshMessageAllOf) *NullableConnectorSshMessageAllOf {
	return &NullableConnectorSshMessageAllOf{value: val, isSet: true}
}

func (v NullableConnectorSshMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorSshMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
