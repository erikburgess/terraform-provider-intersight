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

// HclProduct Model giving the details of product.
type HclProduct struct {
	MoBaseComplexType
	DriverNames *[]string `json:"DriverNames,omitempty"`
	// Error code indicating the support status.
	ErrorCode *string        `json:"ErrorCode,omitempty"`
	Firmwares *[]HclFirmware `json:"Firmwares,omitempty"`
	// Identifier of the product.
	Id *string `json:"Id,omitempty"`
	// Model/PID of the product/adapter.
	Model *string `json:"Model,omitempty"`
	// Revision of the adapter model.
	Revision *string `json:"Revision,omitempty"`
	// Type of the product/adapter say OCP, PT, GPU.
	Type *string `json:"Type,omitempty"`
	// Vendor of the product or adapter.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclProduct HclProduct

// NewHclProduct instantiates a new HclProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclProduct() *HclProduct {
	this := HclProduct{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	return &this
}

// NewHclProductWithDefaults instantiates a new HclProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclProductWithDefaults() *HclProduct {
	this := HclProduct{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	return &this
}

// GetDriverNames returns the DriverNames field value if set, zero value otherwise.
func (o *HclProduct) GetDriverNames() []string {
	if o == nil || o.DriverNames == nil {
		var ret []string
		return ret
	}
	return *o.DriverNames
}

// GetDriverNamesOk returns a tuple with the DriverNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetDriverNamesOk() (*[]string, bool) {
	if o == nil || o.DriverNames == nil {
		return nil, false
	}
	return o.DriverNames, true
}

// HasDriverNames returns a boolean if a field has been set.
func (o *HclProduct) HasDriverNames() bool {
	if o != nil && o.DriverNames != nil {
		return true
	}

	return false
}

// SetDriverNames gets a reference to the given []string and assigns it to the DriverNames field.
func (o *HclProduct) SetDriverNames(v []string) {
	o.DriverNames = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclProduct) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclProduct) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclProduct) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFirmwares returns the Firmwares field value if set, zero value otherwise.
func (o *HclProduct) GetFirmwares() []HclFirmware {
	if o == nil || o.Firmwares == nil {
		var ret []HclFirmware
		return ret
	}
	return *o.Firmwares
}

// GetFirmwaresOk returns a tuple with the Firmwares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetFirmwaresOk() (*[]HclFirmware, bool) {
	if o == nil || o.Firmwares == nil {
		return nil, false
	}
	return o.Firmwares, true
}

// HasFirmwares returns a boolean if a field has been set.
func (o *HclProduct) HasFirmwares() bool {
	if o != nil && o.Firmwares != nil {
		return true
	}

	return false
}

// SetFirmwares gets a reference to the given []HclFirmware and assigns it to the Firmwares field.
func (o *HclProduct) SetFirmwares(v []HclFirmware) {
	o.Firmwares = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclProduct) SetId(v string) {
	o.Id = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HclProduct) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HclProduct) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HclProduct) SetModel(v string) {
	o.Model = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *HclProduct) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *HclProduct) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *HclProduct) SetRevision(v string) {
	o.Revision = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HclProduct) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HclProduct) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HclProduct) SetType(v string) {
	o.Type = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *HclProduct) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProduct) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *HclProduct) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *HclProduct) SetVendor(v string) {
	o.Vendor = &v
}

func (o HclProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.DriverNames != nil {
		toSerialize["DriverNames"] = o.DriverNames
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Firmwares != nil {
		toSerialize["Firmwares"] = o.Firmwares
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclProduct) UnmarshalJSON(bytes []byte) (err error) {
	type HclProductWithoutEmbeddedStruct struct {
		DriverNames *[]string `json:"DriverNames,omitempty"`
		// Error code indicating the support status.
		ErrorCode *string        `json:"ErrorCode,omitempty"`
		Firmwares *[]HclFirmware `json:"Firmwares,omitempty"`
		// Identifier of the product.
		Id *string `json:"Id,omitempty"`
		// Model/PID of the product/adapter.
		Model *string `json:"Model,omitempty"`
		// Revision of the adapter model.
		Revision *string `json:"Revision,omitempty"`
		// Type of the product/adapter say OCP, PT, GPU.
		Type *string `json:"Type,omitempty"`
		// Vendor of the product or adapter.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varHclProductWithoutEmbeddedStruct := HclProductWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclProductWithoutEmbeddedStruct)
	if err == nil {
		varHclProduct := _HclProduct{}
		varHclProduct.DriverNames = varHclProductWithoutEmbeddedStruct.DriverNames
		varHclProduct.ErrorCode = varHclProductWithoutEmbeddedStruct.ErrorCode
		varHclProduct.Firmwares = varHclProductWithoutEmbeddedStruct.Firmwares
		varHclProduct.Id = varHclProductWithoutEmbeddedStruct.Id
		varHclProduct.Model = varHclProductWithoutEmbeddedStruct.Model
		varHclProduct.Revision = varHclProductWithoutEmbeddedStruct.Revision
		varHclProduct.Type = varHclProductWithoutEmbeddedStruct.Type
		varHclProduct.Vendor = varHclProductWithoutEmbeddedStruct.Vendor
		*o = HclProduct(varHclProduct)
	} else {
		return err
	}

	varHclProduct := _HclProduct{}

	err = json.Unmarshal(bytes, &varHclProduct)
	if err == nil {
		o.MoBaseComplexType = varHclProduct.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DriverNames")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "Firmwares")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Revision")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableHclProduct struct {
	value *HclProduct
	isSet bool
}

func (v NullableHclProduct) Get() *HclProduct {
	return v.value
}

func (v *NullableHclProduct) Set(val *HclProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableHclProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableHclProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclProduct(val *HclProduct) *NullableHclProduct {
	return &NullableHclProduct{value: val, isSet: true}
}

func (v NullableHclProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
