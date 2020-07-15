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

// StorageBaseInitiatorAllOf Definition of the list of properties defined in 'storage.BaseInitiator', excluding properties defined in parent classes.
type StorageBaseInitiatorAllOf struct {
	// IQN (iSCSI qualified name). Can be up to 255 characters long and has the format iqn.yyyy-mm.naming-authority:unique name.
	Iqn *string `json:"Iqn,omitempty"`
	// Name of the initiator represented in the storage array.
	Name *string `json:"Name,omitempty"`
	// Initiator type, it can be FC or iSCSI.
	Type *string `json:"Type,omitempty"`
	// World wide name, 128 bit address represented in hexadecimal notation. For example, 51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01.
	Wwn                  *string `json:"Wwn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseInitiatorAllOf StorageBaseInitiatorAllOf

// NewStorageBaseInitiatorAllOf instantiates a new StorageBaseInitiatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseInitiatorAllOf() *StorageBaseInitiatorAllOf {
	this := StorageBaseInitiatorAllOf{}
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// NewStorageBaseInitiatorAllOfWithDefaults instantiates a new StorageBaseInitiatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseInitiatorAllOfWithDefaults() *StorageBaseInitiatorAllOf {
	this := StorageBaseInitiatorAllOf{}
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// GetIqn returns the Iqn field value if set, zero value otherwise.
func (o *StorageBaseInitiatorAllOf) GetIqn() string {
	if o == nil || o.Iqn == nil {
		var ret string
		return ret
	}
	return *o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseInitiatorAllOf) GetIqnOk() (*string, bool) {
	if o == nil || o.Iqn == nil {
		return nil, false
	}
	return o.Iqn, true
}

// HasIqn returns a boolean if a field has been set.
func (o *StorageBaseInitiatorAllOf) HasIqn() bool {
	if o != nil && o.Iqn != nil {
		return true
	}

	return false
}

// SetIqn gets a reference to the given string and assigns it to the Iqn field.
func (o *StorageBaseInitiatorAllOf) SetIqn(v string) {
	o.Iqn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseInitiatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseInitiatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseInitiatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseInitiatorAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageBaseInitiatorAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseInitiatorAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageBaseInitiatorAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageBaseInitiatorAllOf) SetType(v string) {
	o.Type = &v
}

// GetWwn returns the Wwn field value if set, zero value otherwise.
func (o *StorageBaseInitiatorAllOf) GetWwn() string {
	if o == nil || o.Wwn == nil {
		var ret string
		return ret
	}
	return *o.Wwn
}

// GetWwnOk returns a tuple with the Wwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseInitiatorAllOf) GetWwnOk() (*string, bool) {
	if o == nil || o.Wwn == nil {
		return nil, false
	}
	return o.Wwn, true
}

// HasWwn returns a boolean if a field has been set.
func (o *StorageBaseInitiatorAllOf) HasWwn() bool {
	if o != nil && o.Wwn != nil {
		return true
	}

	return false
}

// SetWwn gets a reference to the given string and assigns it to the Wwn field.
func (o *StorageBaseInitiatorAllOf) SetWwn(v string) {
	o.Wwn = &v
}

func (o StorageBaseInitiatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iqn != nil {
		toSerialize["Iqn"] = o.Iqn
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Wwn != nil {
		toSerialize["Wwn"] = o.Wwn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseInitiatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseInitiatorAllOf := _StorageBaseInitiatorAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseInitiatorAllOf); err == nil {
		*o = StorageBaseInitiatorAllOf(varStorageBaseInitiatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Iqn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Wwn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseInitiatorAllOf struct {
	value *StorageBaseInitiatorAllOf
	isSet bool
}

func (v NullableStorageBaseInitiatorAllOf) Get() *StorageBaseInitiatorAllOf {
	return v.value
}

func (v *NullableStorageBaseInitiatorAllOf) Set(val *StorageBaseInitiatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseInitiatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseInitiatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseInitiatorAllOf(val *StorageBaseInitiatorAllOf) *NullableStorageBaseInitiatorAllOf {
	return &NullableStorageBaseInitiatorAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseInitiatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseInitiatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
