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
	"reflect"
	"strings"
)

// IamEndPointPrivilege The privilege defined at the end point which can be assigned to a user.
type IamEndPointPrivilege struct {
	MoBaseMo
	// The functionality of this privilege.
	Description *string `json:"Description,omitempty"`
	// The name of the end point privilege.
	Name *string `json:"Name,omitempty"`
	// The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
	Type                 *string                `json:"Type,omitempty"`
	System               *IamSystemRelationship `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointPrivilege IamEndPointPrivilege

// NewIamEndPointPrivilege instantiates a new IamEndPointPrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointPrivilege() *IamEndPointPrivilege {
	this := IamEndPointPrivilege{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewIamEndPointPrivilegeWithDefaults instantiates a new IamEndPointPrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointPrivilegeWithDefaults() *IamEndPointPrivilege {
	this := IamEndPointPrivilege{}
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamEndPointPrivilege) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPrivilege) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamEndPointPrivilege) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamEndPointPrivilege) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamEndPointPrivilege) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPrivilege) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamEndPointPrivilege) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamEndPointPrivilege) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IamEndPointPrivilege) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPrivilege) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IamEndPointPrivilege) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IamEndPointPrivilege) SetType(v string) {
	o.Type = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamEndPointPrivilege) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointPrivilege) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamEndPointPrivilege) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamEndPointPrivilege) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamEndPointPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointPrivilege) UnmarshalJSON(bytes []byte) (err error) {
	type IamEndPointPrivilegeWithoutEmbeddedStruct struct {
		// The functionality of this privilege.
		Description *string `json:"Description,omitempty"`
		// The name of the end point privilege.
		Name *string `json:"Name,omitempty"`
		// The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
		Type   *string                `json:"Type,omitempty"`
		System *IamSystemRelationship `json:"System,omitempty"`
	}

	varIamEndPointPrivilegeWithoutEmbeddedStruct := IamEndPointPrivilegeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamEndPointPrivilegeWithoutEmbeddedStruct)
	if err == nil {
		varIamEndPointPrivilege := _IamEndPointPrivilege{}
		varIamEndPointPrivilege.Description = varIamEndPointPrivilegeWithoutEmbeddedStruct.Description
		varIamEndPointPrivilege.Name = varIamEndPointPrivilegeWithoutEmbeddedStruct.Name
		varIamEndPointPrivilege.Type = varIamEndPointPrivilegeWithoutEmbeddedStruct.Type
		varIamEndPointPrivilege.System = varIamEndPointPrivilegeWithoutEmbeddedStruct.System
		*o = IamEndPointPrivilege(varIamEndPointPrivilege)
	} else {
		return err
	}

	varIamEndPointPrivilege := _IamEndPointPrivilege{}

	err = json.Unmarshal(bytes, &varIamEndPointPrivilege)
	if err == nil {
		o.MoBaseMo = varIamEndPointPrivilege.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "System")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamEndPointPrivilege struct {
	value *IamEndPointPrivilege
	isSet bool
}

func (v NullableIamEndPointPrivilege) Get() *IamEndPointPrivilege {
	return v.value
}

func (v *NullableIamEndPointPrivilege) Set(val *IamEndPointPrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointPrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointPrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointPrivilege(val *IamEndPointPrivilege) *NullableIamEndPointPrivilege {
	return &NullableIamEndPointPrivilege{value: val, isSet: true}
}

func (v NullableIamEndPointPrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointPrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
