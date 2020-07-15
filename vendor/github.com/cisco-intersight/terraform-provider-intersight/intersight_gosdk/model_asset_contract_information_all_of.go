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

// AssetContractInformationAllOf Definition of the list of properties defined in 'asset.ContractInformation', excluding properties defined in parent classes.
type AssetContractInformationAllOf struct {
	BillTo               *AssetAddressInformation `json:"BillTo,omitempty"`
	BillToGlobalUltimate *AssetGlobalUltimate     `json:"BillToGlobalUltimate,omitempty"`
	// Contract number for the Cisco support contract purchased for the Cisco device.
	ContractNumber *string `json:"ContractNumber,omitempty"`
	// Contract status as per the Cisco Contract APIx.
	LineStatus           *string `json:"LineStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetContractInformationAllOf AssetContractInformationAllOf

// NewAssetContractInformationAllOf instantiates a new AssetContractInformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetContractInformationAllOf() *AssetContractInformationAllOf {
	this := AssetContractInformationAllOf{}
	return &this
}

// NewAssetContractInformationAllOfWithDefaults instantiates a new AssetContractInformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetContractInformationAllOfWithDefaults() *AssetContractInformationAllOf {
	this := AssetContractInformationAllOf{}
	return &this
}

// GetBillTo returns the BillTo field value if set, zero value otherwise.
func (o *AssetContractInformationAllOf) GetBillTo() AssetAddressInformation {
	if o == nil || o.BillTo == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.BillTo
}

// GetBillToOk returns a tuple with the BillTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetContractInformationAllOf) GetBillToOk() (*AssetAddressInformation, bool) {
	if o == nil || o.BillTo == nil {
		return nil, false
	}
	return o.BillTo, true
}

// HasBillTo returns a boolean if a field has been set.
func (o *AssetContractInformationAllOf) HasBillTo() bool {
	if o != nil && o.BillTo != nil {
		return true
	}

	return false
}

// SetBillTo gets a reference to the given AssetAddressInformation and assigns it to the BillTo field.
func (o *AssetContractInformationAllOf) SetBillTo(v AssetAddressInformation) {
	o.BillTo = &v
}

// GetBillToGlobalUltimate returns the BillToGlobalUltimate field value if set, zero value otherwise.
func (o *AssetContractInformationAllOf) GetBillToGlobalUltimate() AssetGlobalUltimate {
	if o == nil || o.BillToGlobalUltimate == nil {
		var ret AssetGlobalUltimate
		return ret
	}
	return *o.BillToGlobalUltimate
}

// GetBillToGlobalUltimateOk returns a tuple with the BillToGlobalUltimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetContractInformationAllOf) GetBillToGlobalUltimateOk() (*AssetGlobalUltimate, bool) {
	if o == nil || o.BillToGlobalUltimate == nil {
		return nil, false
	}
	return o.BillToGlobalUltimate, true
}

// HasBillToGlobalUltimate returns a boolean if a field has been set.
func (o *AssetContractInformationAllOf) HasBillToGlobalUltimate() bool {
	if o != nil && o.BillToGlobalUltimate != nil {
		return true
	}

	return false
}

// SetBillToGlobalUltimate gets a reference to the given AssetGlobalUltimate and assigns it to the BillToGlobalUltimate field.
func (o *AssetContractInformationAllOf) SetBillToGlobalUltimate(v AssetGlobalUltimate) {
	o.BillToGlobalUltimate = &v
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *AssetContractInformationAllOf) GetContractNumber() string {
	if o == nil || o.ContractNumber == nil {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetContractInformationAllOf) GetContractNumberOk() (*string, bool) {
	if o == nil || o.ContractNumber == nil {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *AssetContractInformationAllOf) HasContractNumber() bool {
	if o != nil && o.ContractNumber != nil {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *AssetContractInformationAllOf) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetLineStatus returns the LineStatus field value if set, zero value otherwise.
func (o *AssetContractInformationAllOf) GetLineStatus() string {
	if o == nil || o.LineStatus == nil {
		var ret string
		return ret
	}
	return *o.LineStatus
}

// GetLineStatusOk returns a tuple with the LineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetContractInformationAllOf) GetLineStatusOk() (*string, bool) {
	if o == nil || o.LineStatus == nil {
		return nil, false
	}
	return o.LineStatus, true
}

// HasLineStatus returns a boolean if a field has been set.
func (o *AssetContractInformationAllOf) HasLineStatus() bool {
	if o != nil && o.LineStatus != nil {
		return true
	}

	return false
}

// SetLineStatus gets a reference to the given string and assigns it to the LineStatus field.
func (o *AssetContractInformationAllOf) SetLineStatus(v string) {
	o.LineStatus = &v
}

func (o AssetContractInformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillTo != nil {
		toSerialize["BillTo"] = o.BillTo
	}
	if o.BillToGlobalUltimate != nil {
		toSerialize["BillToGlobalUltimate"] = o.BillToGlobalUltimate
	}
	if o.ContractNumber != nil {
		toSerialize["ContractNumber"] = o.ContractNumber
	}
	if o.LineStatus != nil {
		toSerialize["LineStatus"] = o.LineStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetContractInformationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetContractInformationAllOf := _AssetContractInformationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetContractInformationAllOf); err == nil {
		*o = AssetContractInformationAllOf(varAssetContractInformationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "BillTo")
		delete(additionalProperties, "BillToGlobalUltimate")
		delete(additionalProperties, "ContractNumber")
		delete(additionalProperties, "LineStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetContractInformationAllOf struct {
	value *AssetContractInformationAllOf
	isSet bool
}

func (v NullableAssetContractInformationAllOf) Get() *AssetContractInformationAllOf {
	return v.value
}

func (v *NullableAssetContractInformationAllOf) Set(val *AssetContractInformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetContractInformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetContractInformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetContractInformationAllOf(val *AssetContractInformationAllOf) *NullableAssetContractInformationAllOf {
	return &NullableAssetContractInformationAllOf{value: val, isSet: true}
}

func (v NullableAssetContractInformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetContractInformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
