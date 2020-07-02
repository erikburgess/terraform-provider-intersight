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

// ConnectorBaseMessage A base abstract message that handles the common request properties applicable for all the message types.
type ConnectorBaseMessage struct {
	MoBaseComplexType
	// The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector public key and passed as the value for this property. The secure properties that are encrypted using the AES key are mapped against the property name with prefix 'AES' in SecureProperties dictionary.
	EncryptedAesKey *string `json:"EncryptedAesKey,omitempty"`
	// The public key that was used to encrypt the values present in SecureProperties dictionary. If the given public key is not same as device connector's public key, an error reponse with appropriate error message is thrown back.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using AES key must be mapped against the secure property name with a 'AES' prefix Device connector expects the message body to be a golang template and the template can use the secure property names as placeholders.
	SecureProperties interface{} `json:"SecureProperties,omitempty"`
}

// NewConnectorBaseMessage instantiates a new ConnectorBaseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorBaseMessage() *ConnectorBaseMessage {
	this := ConnectorBaseMessage{}
	return &this
}

// NewConnectorBaseMessageWithDefaults instantiates a new ConnectorBaseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorBaseMessageWithDefaults() *ConnectorBaseMessage {
	this := ConnectorBaseMessage{}
	return &this
}

// GetEncryptedAesKey returns the EncryptedAesKey field value if set, zero value otherwise.
func (o *ConnectorBaseMessage) GetEncryptedAesKey() string {
	if o == nil || o.EncryptedAesKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptedAesKey
}

// GetEncryptedAesKeyOk returns a tuple with the EncryptedAesKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorBaseMessage) GetEncryptedAesKeyOk() (*string, bool) {
	if o == nil || o.EncryptedAesKey == nil {
		return nil, false
	}
	return o.EncryptedAesKey, true
}

// HasEncryptedAesKey returns a boolean if a field has been set.
func (o *ConnectorBaseMessage) HasEncryptedAesKey() bool {
	if o != nil && o.EncryptedAesKey != nil {
		return true
	}

	return false
}

// SetEncryptedAesKey gets a reference to the given string and assigns it to the EncryptedAesKey field.
func (o *ConnectorBaseMessage) SetEncryptedAesKey(v string) {
	o.EncryptedAesKey = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *ConnectorBaseMessage) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorBaseMessage) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *ConnectorBaseMessage) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *ConnectorBaseMessage) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetSecureProperties returns the SecureProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorBaseMessage) GetSecureProperties() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SecureProperties
}

// GetSecurePropertiesOk returns a tuple with the SecureProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorBaseMessage) GetSecurePropertiesOk() (*interface{}, bool) {
	if o == nil || o.SecureProperties == nil {
		return nil, false
	}
	return &o.SecureProperties, true
}

// HasSecureProperties returns a boolean if a field has been set.
func (o *ConnectorBaseMessage) HasSecureProperties() bool {
	if o != nil && o.SecureProperties != nil {
		return true
	}

	return false
}

// SetSecureProperties gets a reference to the given interface{} and assigns it to the SecureProperties field.
func (o *ConnectorBaseMessage) SetSecureProperties(v interface{}) {
	o.SecureProperties = v
}

func (o ConnectorBaseMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.EncryptedAesKey != nil {
		toSerialize["EncryptedAesKey"] = o.EncryptedAesKey
	}
	if o.EncryptionKey != nil {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.SecureProperties != nil {
		toSerialize["SecureProperties"] = o.SecureProperties
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorBaseMessage struct {
	value *ConnectorBaseMessage
	isSet bool
}

func (v NullableConnectorBaseMessage) Get() *ConnectorBaseMessage {
	return v.value
}

func (v *NullableConnectorBaseMessage) Set(val *ConnectorBaseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorBaseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorBaseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorBaseMessage(val *ConnectorBaseMessage) *NullableConnectorBaseMessage {
	return &NullableConnectorBaseMessage{value: val, isSet: true}
}

func (v NullableConnectorBaseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorBaseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
