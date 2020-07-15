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

// VirtualizationVmwareCluster A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
type VirtualizationVmwareCluster struct {
	VirtualizationBaseCluster
	// Count of all datastores associated with this cluster.
	DatastoreCount       *int64                                      `json:"DatastoreCount,omitempty"`
	Datacenter           *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareCluster VirtualizationVmwareCluster

// NewVirtualizationVmwareCluster instantiates a new VirtualizationVmwareCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareCluster() *VirtualizationVmwareCluster {
	this := VirtualizationVmwareCluster{}
	return &this
}

// NewVirtualizationVmwareClusterWithDefaults instantiates a new VirtualizationVmwareCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareClusterWithDefaults() *VirtualizationVmwareCluster {
	this := VirtualizationVmwareCluster{}
	return &this
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetDatastoreCount() int64 {
	if o == nil || o.DatastoreCount == nil {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || o.DatastoreCount == nil {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasDatastoreCount() bool {
	if o != nil && o.DatastoreCount != nil {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareCluster) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareCluster) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

func (o VirtualizationVmwareCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseCluster, errVirtualizationBaseCluster := json.Marshal(o.VirtualizationBaseCluster)
	if errVirtualizationBaseCluster != nil {
		return []byte{}, errVirtualizationBaseCluster
	}
	errVirtualizationBaseCluster = json.Unmarshal([]byte(serializedVirtualizationBaseCluster), &toSerialize)
	if errVirtualizationBaseCluster != nil {
		return []byte{}, errVirtualizationBaseCluster
	}
	if o.DatastoreCount != nil {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareCluster) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareClusterWithoutEmbeddedStruct struct {
		// Count of all datastores associated with this cluster.
		DatastoreCount *int64                                      `json:"DatastoreCount,omitempty"`
		Datacenter     *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	}

	varVirtualizationVmwareClusterWithoutEmbeddedStruct := VirtualizationVmwareClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareClusterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareCluster := _VirtualizationVmwareCluster{}
		varVirtualizationVmwareCluster.DatastoreCount = varVirtualizationVmwareClusterWithoutEmbeddedStruct.DatastoreCount
		varVirtualizationVmwareCluster.Datacenter = varVirtualizationVmwareClusterWithoutEmbeddedStruct.Datacenter
		*o = VirtualizationVmwareCluster(varVirtualizationVmwareCluster)
	} else {
		return err
	}

	varVirtualizationVmwareCluster := _VirtualizationVmwareCluster{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareCluster)
	if err == nil {
		o.VirtualizationBaseCluster = varVirtualizationVmwareCluster.VirtualizationBaseCluster
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DatastoreCount")
		delete(additionalProperties, "Datacenter")

		// remove fields from embedded structs
		reflectVirtualizationBaseCluster := reflect.ValueOf(o.VirtualizationBaseCluster)
		for i := 0; i < reflectVirtualizationBaseCluster.Type().NumField(); i++ {
			t := reflectVirtualizationBaseCluster.Type().Field(i)

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

type NullableVirtualizationVmwareCluster struct {
	value *VirtualizationVmwareCluster
	isSet bool
}

func (v NullableVirtualizationVmwareCluster) Get() *VirtualizationVmwareCluster {
	return v.value
}

func (v *NullableVirtualizationVmwareCluster) Set(val *VirtualizationVmwareCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareCluster(val *VirtualizationVmwareCluster) *NullableVirtualizationVmwareCluster {
	return &NullableVirtualizationVmwareCluster{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
