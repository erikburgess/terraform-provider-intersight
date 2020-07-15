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

// HyperflexFeatureLimitExternal The HyperFlex feature limits that are available to end users.
type HyperflexFeatureLimitExternal struct {
	MoBaseMo
	FeatureLimitEntries  *[]HyperflexFeatureLimitEntry    `json:"FeatureLimitEntries,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFeatureLimitExternal HyperflexFeatureLimitExternal

// NewHyperflexFeatureLimitExternal instantiates a new HyperflexFeatureLimitExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFeatureLimitExternal() *HyperflexFeatureLimitExternal {
	this := HyperflexFeatureLimitExternal{}
	return &this
}

// NewHyperflexFeatureLimitExternalWithDefaults instantiates a new HyperflexFeatureLimitExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFeatureLimitExternalWithDefaults() *HyperflexFeatureLimitExternal {
	this := HyperflexFeatureLimitExternal{}
	return &this
}

// GetFeatureLimitEntries returns the FeatureLimitEntries field value if set, zero value otherwise.
func (o *HyperflexFeatureLimitExternal) GetFeatureLimitEntries() []HyperflexFeatureLimitEntry {
	if o == nil || o.FeatureLimitEntries == nil {
		var ret []HyperflexFeatureLimitEntry
		return ret
	}
	return *o.FeatureLimitEntries
}

// GetFeatureLimitEntriesOk returns a tuple with the FeatureLimitEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitExternal) GetFeatureLimitEntriesOk() (*[]HyperflexFeatureLimitEntry, bool) {
	if o == nil || o.FeatureLimitEntries == nil {
		return nil, false
	}
	return o.FeatureLimitEntries, true
}

// HasFeatureLimitEntries returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitExternal) HasFeatureLimitEntries() bool {
	if o != nil && o.FeatureLimitEntries != nil {
		return true
	}

	return false
}

// SetFeatureLimitEntries gets a reference to the given []HyperflexFeatureLimitEntry and assigns it to the FeatureLimitEntries field.
func (o *HyperflexFeatureLimitExternal) SetFeatureLimitEntries(v []HyperflexFeatureLimitEntry) {
	o.FeatureLimitEntries = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexFeatureLimitExternal) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitExternal) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitExternal) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexFeatureLimitExternal) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexFeatureLimitExternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.FeatureLimitEntries != nil {
		toSerialize["FeatureLimitEntries"] = o.FeatureLimitEntries
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexFeatureLimitExternal) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexFeatureLimitExternalWithoutEmbeddedStruct struct {
		FeatureLimitEntries *[]HyperflexFeatureLimitEntry    `json:"FeatureLimitEntries,omitempty"`
		AppCatalog          *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHyperflexFeatureLimitExternalWithoutEmbeddedStruct := HyperflexFeatureLimitExternalWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexFeatureLimitExternalWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexFeatureLimitExternal := _HyperflexFeatureLimitExternal{}
		varHyperflexFeatureLimitExternal.FeatureLimitEntries = varHyperflexFeatureLimitExternalWithoutEmbeddedStruct.FeatureLimitEntries
		varHyperflexFeatureLimitExternal.AppCatalog = varHyperflexFeatureLimitExternalWithoutEmbeddedStruct.AppCatalog
		*o = HyperflexFeatureLimitExternal(varHyperflexFeatureLimitExternal)
	} else {
		return err
	}

	varHyperflexFeatureLimitExternal := _HyperflexFeatureLimitExternal{}

	err = json.Unmarshal(bytes, &varHyperflexFeatureLimitExternal)
	if err == nil {
		o.MoBaseMo = varHyperflexFeatureLimitExternal.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FeatureLimitEntries")
		delete(additionalProperties, "AppCatalog")

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

type NullableHyperflexFeatureLimitExternal struct {
	value *HyperflexFeatureLimitExternal
	isSet bool
}

func (v NullableHyperflexFeatureLimitExternal) Get() *HyperflexFeatureLimitExternal {
	return v.value
}

func (v *NullableHyperflexFeatureLimitExternal) Set(val *HyperflexFeatureLimitExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFeatureLimitExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFeatureLimitExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFeatureLimitExternal(val *HyperflexFeatureLimitExternal) *NullableHyperflexFeatureLimitExternal {
	return &NullableHyperflexFeatureLimitExternal{value: val, isSet: true}
}

func (v NullableHyperflexFeatureLimitExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFeatureLimitExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
