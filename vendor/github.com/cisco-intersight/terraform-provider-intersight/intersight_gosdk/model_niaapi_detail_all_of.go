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

// NiaapiDetailAllOf Definition of the list of properties defined in 'niaapi.Detail', excluding properties defined in parent classes.
type NiaapiDetailAllOf struct {
	// Checksum of this part of Content.
	Chksum *string `json:"Chksum,omitempty"`
	// The file name within this Metadata file.
	Filename *string `json:"Filename,omitempty"`
	// The name of this Content.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiDetailAllOf NiaapiDetailAllOf

// NewNiaapiDetailAllOf instantiates a new NiaapiDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiDetailAllOf() *NiaapiDetailAllOf {
	this := NiaapiDetailAllOf{}
	return &this
}

// NewNiaapiDetailAllOfWithDefaults instantiates a new NiaapiDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiDetailAllOfWithDefaults() *NiaapiDetailAllOf {
	this := NiaapiDetailAllOf{}
	return &this
}

// GetChksum returns the Chksum field value if set, zero value otherwise.
func (o *NiaapiDetailAllOf) GetChksum() string {
	if o == nil || o.Chksum == nil {
		var ret string
		return ret
	}
	return *o.Chksum
}

// GetChksumOk returns a tuple with the Chksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDetailAllOf) GetChksumOk() (*string, bool) {
	if o == nil || o.Chksum == nil {
		return nil, false
	}
	return o.Chksum, true
}

// HasChksum returns a boolean if a field has been set.
func (o *NiaapiDetailAllOf) HasChksum() bool {
	if o != nil && o.Chksum != nil {
		return true
	}

	return false
}

// SetChksum gets a reference to the given string and assigns it to the Chksum field.
func (o *NiaapiDetailAllOf) SetChksum(v string) {
	o.Chksum = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *NiaapiDetailAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDetailAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *NiaapiDetailAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *NiaapiDetailAllOf) SetFilename(v string) {
	o.Filename = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiaapiDetailAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDetailAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiaapiDetailAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiaapiDetailAllOf) SetName(v string) {
	o.Name = &v
}

func (o NiaapiDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Chksum != nil {
		toSerialize["Chksum"] = o.Chksum
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiDetailAllOf := _NiaapiDetailAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiDetailAllOf); err == nil {
		*o = NiaapiDetailAllOf(varNiaapiDetailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Chksum")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiDetailAllOf struct {
	value *NiaapiDetailAllOf
	isSet bool
}

func (v NullableNiaapiDetailAllOf) Get() *NiaapiDetailAllOf {
	return v.value
}

func (v *NullableNiaapiDetailAllOf) Set(val *NiaapiDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiDetailAllOf(val *NiaapiDetailAllOf) *NullableNiaapiDetailAllOf {
	return &NullableNiaapiDetailAllOf{value: val, isSet: true}
}

func (v NullableNiaapiDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
