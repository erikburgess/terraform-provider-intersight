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

// FabricPcOperationAllOf Definition of the list of properties defined in 'fabric.PcOperation', excluding properties defined in parent classes.
type FabricPcOperationAllOf struct {
	// Admin configured state to disable the port channel.
	AdminState *string `json:"AdminState,omitempty"`
	// Port Channel Identifier for the collection of ports.
	PcId           *int64                      `json:"PcId,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
}

// NewFabricPcOperationAllOf instantiates a new FabricPcOperationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPcOperationAllOf() *FabricPcOperationAllOf {
	this := FabricPcOperationAllOf{}
	var adminState string = "Enabled"
	this.AdminState = &adminState
	return &this
}

// NewFabricPcOperationAllOfWithDefaults instantiates a new FabricPcOperationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPcOperationAllOfWithDefaults() *FabricPcOperationAllOf {
	this := FabricPcOperationAllOf{}
	var adminState string = "Enabled"
	this.AdminState = &adminState
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricPcOperationAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperationAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricPcOperationAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricPcOperationAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricPcOperationAllOf) GetPcId() int64 {
	if o == nil || o.PcId == nil {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperationAllOf) GetPcIdOk() (*int64, bool) {
	if o == nil || o.PcId == nil {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricPcOperationAllOf) HasPcId() bool {
	if o != nil && o.PcId != nil {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricPcOperationAllOf) SetPcId(v int64) {
	o.PcId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *FabricPcOperationAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperationAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricPcOperationAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricPcOperationAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

func (o FabricPcOperationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.PcId != nil {
		toSerialize["PcId"] = o.PcId
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	return json.Marshal(toSerialize)
}

type NullableFabricPcOperationAllOf struct {
	value *FabricPcOperationAllOf
	isSet bool
}

func (v NullableFabricPcOperationAllOf) Get() *FabricPcOperationAllOf {
	return v.value
}

func (v *NullableFabricPcOperationAllOf) Set(val *FabricPcOperationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPcOperationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPcOperationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPcOperationAllOf(val *FabricPcOperationAllOf) *NullableFabricPcOperationAllOf {
	return &NullableFabricPcOperationAllOf{value: val, isSet: true}
}

func (v NullableFabricPcOperationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPcOperationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
