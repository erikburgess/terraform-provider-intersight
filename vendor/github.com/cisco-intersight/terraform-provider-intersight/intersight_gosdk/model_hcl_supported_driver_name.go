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

// HclSupportedDriverName Supported driver names for a given product for the given operating system.
type HclSupportedDriverName struct {
	MoBaseMo
	// Vendor distributing the Operating System.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating System.
	OsVersion   *string       `json:"OsVersion,omitempty"`
	ProductList *[]HclProduct `json:"ProductList,omitempty"`
}

// NewHclSupportedDriverName instantiates a new HclSupportedDriverName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclSupportedDriverName() *HclSupportedDriverName {
	this := HclSupportedDriverName{}
	return &this
}

// NewHclSupportedDriverNameWithDefaults instantiates a new HclSupportedDriverName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclSupportedDriverNameWithDefaults() *HclSupportedDriverName {
	this := HclSupportedDriverName{}
	return &this
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclSupportedDriverName) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclSupportedDriverName) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclSupportedDriverName) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclSupportedDriverName) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclSupportedDriverName) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclSupportedDriverName) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProductList returns the ProductList field value if set, zero value otherwise.
func (o *HclSupportedDriverName) GetProductList() []HclProduct {
	if o == nil || o.ProductList == nil {
		var ret []HclProduct
		return ret
	}
	return *o.ProductList
}

// GetProductListOk returns a tuple with the ProductList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclSupportedDriverName) GetProductListOk() (*[]HclProduct, bool) {
	if o == nil || o.ProductList == nil {
		return nil, false
	}
	return o.ProductList, true
}

// HasProductList returns a boolean if a field has been set.
func (o *HclSupportedDriverName) HasProductList() bool {
	if o != nil && o.ProductList != nil {
		return true
	}

	return false
}

// SetProductList gets a reference to the given []HclProduct and assigns it to the ProductList field.
func (o *HclSupportedDriverName) SetProductList(v []HclProduct) {
	o.ProductList = &v
}

func (o HclSupportedDriverName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.OsVendor != nil {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.ProductList != nil {
		toSerialize["ProductList"] = o.ProductList
	}
	return json.Marshal(toSerialize)
}

type NullableHclSupportedDriverName struct {
	value *HclSupportedDriverName
	isSet bool
}

func (v NullableHclSupportedDriverName) Get() *HclSupportedDriverName {
	return v.value
}

func (v *NullableHclSupportedDriverName) Set(val *HclSupportedDriverName) {
	v.value = val
	v.isSet = true
}

func (v NullableHclSupportedDriverName) IsSet() bool {
	return v.isSet
}

func (v *NullableHclSupportedDriverName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclSupportedDriverName(val *HclSupportedDriverName) *NullableHclSupportedDriverName {
	return &NullableHclSupportedDriverName{value: val, isSet: true}
}

func (v NullableHclSupportedDriverName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclSupportedDriverName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
