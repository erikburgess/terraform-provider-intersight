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

// ComputeIpAddress Complex type representing an IP address in UCSM.
type ComputeIpAddress struct {
	MoBaseComplexType
	// IP Address to be used for KVM.
	Address *string `json:"Address,omitempty"`
	// Category of the Kvm IP Address.
	Category *string `json:"Category,omitempty"`
	// Default gateway property of KVM IP Address.
	DefaultGateway *string `json:"DefaultGateway,omitempty"`
	// The distinguished name for this managed object.
	Dn *string `json:"Dn,omitempty"`
	// HTTP port of an IP Address.
	HttpPort *int64 `json:"HttpPort,omitempty"`
	// Secured HTTP port of an IP Address.
	HttpsPort *int64 `json:"HttpsPort,omitempty"`
	// Port number on which the KVM is running and used for connecting to KVM console.
	KvmPort *int64 `json:"KvmPort,omitempty"`
	// VLAN Identifier of Inband IP Address.
	KvmVlan *int64 `json:"KvmVlan,omitempty"`
	// Name to identify the KVM IP Address.
	Name *string `json:"Name,omitempty"`
	// Subnet detail of a KVM IP Address.
	Subnet *string `json:"Subnet,omitempty"`
	// Type of the KVM IP Address.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeIpAddress ComputeIpAddress

// NewComputeIpAddress instantiates a new ComputeIpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeIpAddress() *ComputeIpAddress {
	this := ComputeIpAddress{}
	var category string = "Equipment"
	this.Category = &category
	var name string = "Outband"
	this.Name = &name
	var type_ string = "MgmtInterface"
	this.Type = &type_
	return &this
}

// NewComputeIpAddressWithDefaults instantiates a new ComputeIpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeIpAddressWithDefaults() *ComputeIpAddress {
	this := ComputeIpAddress{}
	var category string = "Equipment"
	this.Category = &category
	var name string = "Outband"
	this.Name = &name
	var type_ string = "MgmtInterface"
	this.Type = &type_
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ComputeIpAddress) SetAddress(v string) {
	o.Address = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ComputeIpAddress) SetCategory(v string) {
	o.Category = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetDefaultGateway() string {
	if o == nil || o.DefaultGateway == nil {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *ComputeIpAddress) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *ComputeIpAddress) SetDn(v string) {
	o.Dn = &v
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetHttpPort() int64 {
	if o == nil || o.HttpPort == nil {
		var ret int64
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetHttpPortOk() (*int64, bool) {
	if o == nil || o.HttpPort == nil {
		return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasHttpPort() bool {
	if o != nil && o.HttpPort != nil {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given int64 and assigns it to the HttpPort field.
func (o *ComputeIpAddress) SetHttpPort(v int64) {
	o.HttpPort = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetHttpsPort() int64 {
	if o == nil || o.HttpsPort == nil {
		var ret int64
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetHttpsPortOk() (*int64, bool) {
	if o == nil || o.HttpsPort == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasHttpsPort() bool {
	if o != nil && o.HttpsPort != nil {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int64 and assigns it to the HttpsPort field.
func (o *ComputeIpAddress) SetHttpsPort(v int64) {
	o.HttpsPort = &v
}

// GetKvmPort returns the KvmPort field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetKvmPort() int64 {
	if o == nil || o.KvmPort == nil {
		var ret int64
		return ret
	}
	return *o.KvmPort
}

// GetKvmPortOk returns a tuple with the KvmPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetKvmPortOk() (*int64, bool) {
	if o == nil || o.KvmPort == nil {
		return nil, false
	}
	return o.KvmPort, true
}

// HasKvmPort returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasKvmPort() bool {
	if o != nil && o.KvmPort != nil {
		return true
	}

	return false
}

// SetKvmPort gets a reference to the given int64 and assigns it to the KvmPort field.
func (o *ComputeIpAddress) SetKvmPort(v int64) {
	o.KvmPort = &v
}

// GetKvmVlan returns the KvmVlan field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetKvmVlan() int64 {
	if o == nil || o.KvmVlan == nil {
		var ret int64
		return ret
	}
	return *o.KvmVlan
}

// GetKvmVlanOk returns a tuple with the KvmVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetKvmVlanOk() (*int64, bool) {
	if o == nil || o.KvmVlan == nil {
		return nil, false
	}
	return o.KvmVlan, true
}

// HasKvmVlan returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasKvmVlan() bool {
	if o != nil && o.KvmVlan != nil {
		return true
	}

	return false
}

// SetKvmVlan gets a reference to the given int64 and assigns it to the KvmVlan field.
func (o *ComputeIpAddress) SetKvmVlan(v int64) {
	o.KvmVlan = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeIpAddress) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *ComputeIpAddress) SetSubnet(v string) {
	o.Subnet = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ComputeIpAddress) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddress) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ComputeIpAddress) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ComputeIpAddress) SetType(v string) {
	o.Type = &v
}

func (o ComputeIpAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.DefaultGateway != nil {
		toSerialize["DefaultGateway"] = o.DefaultGateway
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.HttpPort != nil {
		toSerialize["HttpPort"] = o.HttpPort
	}
	if o.HttpsPort != nil {
		toSerialize["HttpsPort"] = o.HttpsPort
	}
	if o.KvmPort != nil {
		toSerialize["KvmPort"] = o.KvmPort
	}
	if o.KvmVlan != nil {
		toSerialize["KvmVlan"] = o.KvmVlan
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["Subnet"] = o.Subnet
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeIpAddress) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeIpAddressWithoutEmbeddedStruct struct {
		// IP Address to be used for KVM.
		Address *string `json:"Address,omitempty"`
		// Category of the Kvm IP Address.
		Category *string `json:"Category,omitempty"`
		// Default gateway property of KVM IP Address.
		DefaultGateway *string `json:"DefaultGateway,omitempty"`
		// The distinguished name for this managed object.
		Dn *string `json:"Dn,omitempty"`
		// HTTP port of an IP Address.
		HttpPort *int64 `json:"HttpPort,omitempty"`
		// Secured HTTP port of an IP Address.
		HttpsPort *int64 `json:"HttpsPort,omitempty"`
		// Port number on which the KVM is running and used for connecting to KVM console.
		KvmPort *int64 `json:"KvmPort,omitempty"`
		// VLAN Identifier of Inband IP Address.
		KvmVlan *int64 `json:"KvmVlan,omitempty"`
		// Name to identify the KVM IP Address.
		Name *string `json:"Name,omitempty"`
		// Subnet detail of a KVM IP Address.
		Subnet *string `json:"Subnet,omitempty"`
		// Type of the KVM IP Address.
		Type *string `json:"Type,omitempty"`
	}

	varComputeIpAddressWithoutEmbeddedStruct := ComputeIpAddressWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeIpAddressWithoutEmbeddedStruct)
	if err == nil {
		varComputeIpAddress := _ComputeIpAddress{}
		varComputeIpAddress.Address = varComputeIpAddressWithoutEmbeddedStruct.Address
		varComputeIpAddress.Category = varComputeIpAddressWithoutEmbeddedStruct.Category
		varComputeIpAddress.DefaultGateway = varComputeIpAddressWithoutEmbeddedStruct.DefaultGateway
		varComputeIpAddress.Dn = varComputeIpAddressWithoutEmbeddedStruct.Dn
		varComputeIpAddress.HttpPort = varComputeIpAddressWithoutEmbeddedStruct.HttpPort
		varComputeIpAddress.HttpsPort = varComputeIpAddressWithoutEmbeddedStruct.HttpsPort
		varComputeIpAddress.KvmPort = varComputeIpAddressWithoutEmbeddedStruct.KvmPort
		varComputeIpAddress.KvmVlan = varComputeIpAddressWithoutEmbeddedStruct.KvmVlan
		varComputeIpAddress.Name = varComputeIpAddressWithoutEmbeddedStruct.Name
		varComputeIpAddress.Subnet = varComputeIpAddressWithoutEmbeddedStruct.Subnet
		varComputeIpAddress.Type = varComputeIpAddressWithoutEmbeddedStruct.Type
		*o = ComputeIpAddress(varComputeIpAddress)
	} else {
		return err
	}

	varComputeIpAddress := _ComputeIpAddress{}

	err = json.Unmarshal(bytes, &varComputeIpAddress)
	if err == nil {
		o.MoBaseComplexType = varComputeIpAddress.MoBaseComplexType
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "DefaultGateway")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "HttpPort")
		delete(additionalProperties, "HttpsPort")
		delete(additionalProperties, "KvmPort")
		delete(additionalProperties, "KvmVlan")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Subnet")
		delete(additionalProperties, "Type")

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

type NullableComputeIpAddress struct {
	value *ComputeIpAddress
	isSet bool
}

func (v NullableComputeIpAddress) Get() *ComputeIpAddress {
	return v.value
}

func (v *NullableComputeIpAddress) Set(val *ComputeIpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeIpAddress(val *ComputeIpAddress) *NullableComputeIpAddress {
	return &NullableComputeIpAddress{value: val, isSet: true}
}

func (v NullableComputeIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
