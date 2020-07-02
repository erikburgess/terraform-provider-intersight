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

// FirmwareFirmwareInventory Firmware inventory for server component.
type FirmwareFirmwareInventory struct {
	MoBaseComplexType
	// Component category. For example, MRAID is under storage controller, CIMC is under management controller.
	Category *string `json:"Category,omitempty"`
	// The name of the component to reflect on UI.
	Label *string `json:"Label,omitempty"`
	// Model deatils of component.
	Model *string `json:"Model,omitempty"`
	// The redfish URI to get the firmware inventory of server components.
	UpdateUri *string `json:"UpdateUri,omitempty"`
	// The vendor of the component.
	Vendor *string `json:"Vendor,omitempty"`
	// The firmware running version on the component.
	Version *string `json:"Version,omitempty"`
}

// NewFirmwareFirmwareInventory instantiates a new FirmwareFirmwareInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareFirmwareInventory() *FirmwareFirmwareInventory {
	this := FirmwareFirmwareInventory{}
	return &this
}

// NewFirmwareFirmwareInventoryWithDefaults instantiates a new FirmwareFirmwareInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareFirmwareInventoryWithDefaults() *FirmwareFirmwareInventory {
	this := FirmwareFirmwareInventory{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *FirmwareFirmwareInventory) SetCategory(v string) {
	o.Category = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FirmwareFirmwareInventory) SetLabel(v string) {
	o.Label = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FirmwareFirmwareInventory) SetModel(v string) {
	o.Model = &v
}

// GetUpdateUri returns the UpdateUri field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetUpdateUri() string {
	if o == nil || o.UpdateUri == nil {
		var ret string
		return ret
	}
	return *o.UpdateUri
}

// GetUpdateUriOk returns a tuple with the UpdateUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetUpdateUriOk() (*string, bool) {
	if o == nil || o.UpdateUri == nil {
		return nil, false
	}
	return o.UpdateUri, true
}

// HasUpdateUri returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasUpdateUri() bool {
	if o != nil && o.UpdateUri != nil {
		return true
	}

	return false
}

// SetUpdateUri gets a reference to the given string and assigns it to the UpdateUri field.
func (o *FirmwareFirmwareInventory) SetUpdateUri(v string) {
	o.UpdateUri = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *FirmwareFirmwareInventory) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FirmwareFirmwareInventory) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareInventory) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FirmwareFirmwareInventory) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FirmwareFirmwareInventory) SetVersion(v string) {
	o.Version = &v
}

func (o FirmwareFirmwareInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.UpdateUri != nil {
		toSerialize["UpdateUri"] = o.UpdateUri
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareFirmwareInventory struct {
	value *FirmwareFirmwareInventory
	isSet bool
}

func (v NullableFirmwareFirmwareInventory) Get() *FirmwareFirmwareInventory {
	return v.value
}

func (v *NullableFirmwareFirmwareInventory) Set(val *FirmwareFirmwareInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareFirmwareInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareFirmwareInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareFirmwareInventory(val *FirmwareFirmwareInventory) *NullableFirmwareFirmwareInventory {
	return &NullableFirmwareFirmwareInventory{value: val, isSet: true}
}

func (v NullableFirmwareFirmwareInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareFirmwareInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
