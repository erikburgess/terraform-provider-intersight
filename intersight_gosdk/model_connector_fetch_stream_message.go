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
	"reflect"
	"strings"
)

// ConnectorFetchStreamMessage Retrieve a list of cached stream messages by stream id. Cloud services will request stream messages to be re-sent in case of dropped messages (the cloud service receieves an unexpected stream sequence number). On success the device connector will 'replay' the messages, publishing them to the streams response topic, they will not be returned in the response to this message. If any of the requested sequences are not present in the cache an error will be returned.
type ConnectorFetchStreamMessage struct {
	ConnectorStreamMessage
	Sequences            *[]int64 `json:"Sequences,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorFetchStreamMessage ConnectorFetchStreamMessage

// NewConnectorFetchStreamMessage instantiates a new ConnectorFetchStreamMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorFetchStreamMessage() *ConnectorFetchStreamMessage {
	this := ConnectorFetchStreamMessage{}
	return &this
}

// NewConnectorFetchStreamMessageWithDefaults instantiates a new ConnectorFetchStreamMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorFetchStreamMessageWithDefaults() *ConnectorFetchStreamMessage {
	this := ConnectorFetchStreamMessage{}
	return &this
}

// GetSequences returns the Sequences field value if set, zero value otherwise.
func (o *ConnectorFetchStreamMessage) GetSequences() []int64 {
	if o == nil || o.Sequences == nil {
		var ret []int64
		return ret
	}
	return *o.Sequences
}

// GetSequencesOk returns a tuple with the Sequences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFetchStreamMessage) GetSequencesOk() (*[]int64, bool) {
	if o == nil || o.Sequences == nil {
		return nil, false
	}
	return o.Sequences, true
}

// HasSequences returns a boolean if a field has been set.
func (o *ConnectorFetchStreamMessage) HasSequences() bool {
	if o != nil && o.Sequences != nil {
		return true
	}

	return false
}

// SetSequences gets a reference to the given []int64 and assigns it to the Sequences field.
func (o *ConnectorFetchStreamMessage) SetSequences(v []int64) {
	o.Sequences = &v
}

func (o ConnectorFetchStreamMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	if o.Sequences != nil {
		toSerialize["Sequences"] = o.Sequences
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorFetchStreamMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorFetchStreamMessageWithoutEmbeddedStruct struct {
		Sequences *[]int64 `json:"Sequences,omitempty"`
	}

	varConnectorFetchStreamMessageWithoutEmbeddedStruct := ConnectorFetchStreamMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorFetchStreamMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorFetchStreamMessage := _ConnectorFetchStreamMessage{}
		varConnectorFetchStreamMessage.Sequences = varConnectorFetchStreamMessageWithoutEmbeddedStruct.Sequences
		*o = ConnectorFetchStreamMessage(varConnectorFetchStreamMessage)
	} else {
		return err
	}

	varConnectorFetchStreamMessage := _ConnectorFetchStreamMessage{}

	err = json.Unmarshal(bytes, &varConnectorFetchStreamMessage)
	if err == nil {
		o.ConnectorStreamMessage = varConnectorFetchStreamMessage.ConnectorStreamMessage
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Sequences")

		// remove fields from embedded structs
		reflectConnectorStreamMessage := reflect.ValueOf(o.ConnectorStreamMessage)
		for i := 0; i < reflectConnectorStreamMessage.Type().NumField(); i++ {
			t := reflectConnectorStreamMessage.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorFetchStreamMessage struct {
	value *ConnectorFetchStreamMessage
	isSet bool
}

func (v NullableConnectorFetchStreamMessage) Get() *ConnectorFetchStreamMessage {
	return v.value
}

func (v *NullableConnectorFetchStreamMessage) Set(val *ConnectorFetchStreamMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorFetchStreamMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorFetchStreamMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorFetchStreamMessage(val *ConnectorFetchStreamMessage) *NullableConnectorFetchStreamMessage {
	return &NullableConnectorFetchStreamMessage{value: val, isSet: true}
}

func (v NullableConnectorFetchStreamMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorFetchStreamMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
