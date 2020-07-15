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

// AssetCustomerInformationAllOf Definition of the list of properties defined in 'asset.CustomerInformation', excluding properties defined in parent classes.
type AssetCustomerInformationAllOf struct {
	Address *AssetAddressInformation `json:"Address,omitempty"`
	// Unique identifier for an end customer. This identifier is allocated by Cisco.
	Id *string `json:"Id,omitempty"`
	// Name as per the information provided by the user.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetCustomerInformationAllOf AssetCustomerInformationAllOf

// NewAssetCustomerInformationAllOf instantiates a new AssetCustomerInformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCustomerInformationAllOf() *AssetCustomerInformationAllOf {
	this := AssetCustomerInformationAllOf{}
	return &this
}

// NewAssetCustomerInformationAllOfWithDefaults instantiates a new AssetCustomerInformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCustomerInformationAllOfWithDefaults() *AssetCustomerInformationAllOf {
	this := AssetCustomerInformationAllOf{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AssetCustomerInformationAllOf) GetAddress() AssetAddressInformation {
	if o == nil || o.Address == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformationAllOf) GetAddressOk() (*AssetAddressInformation, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AssetCustomerInformationAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AssetAddressInformation and assigns it to the Address field.
func (o *AssetCustomerInformationAllOf) SetAddress(v AssetAddressInformation) {
	o.Address = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetCustomerInformationAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformationAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetCustomerInformationAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetCustomerInformationAllOf) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetCustomerInformationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCustomerInformationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetCustomerInformationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetCustomerInformationAllOf) SetName(v string) {
	o.Name = &v
}

func (o AssetCustomerInformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetCustomerInformationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetCustomerInformationAllOf := _AssetCustomerInformationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetCustomerInformationAllOf); err == nil {
		*o = AssetCustomerInformationAllOf(varAssetCustomerInformationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetCustomerInformationAllOf struct {
	value *AssetCustomerInformationAllOf
	isSet bool
}

func (v NullableAssetCustomerInformationAllOf) Get() *AssetCustomerInformationAllOf {
	return v.value
}

func (v *NullableAssetCustomerInformationAllOf) Set(val *AssetCustomerInformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCustomerInformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCustomerInformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCustomerInformationAllOf(val *AssetCustomerInformationAllOf) *NullableAssetCustomerInformationAllOf {
	return &NullableAssetCustomerInformationAllOf{value: val, isSet: true}
}

func (v NullableAssetCustomerInformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCustomerInformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
