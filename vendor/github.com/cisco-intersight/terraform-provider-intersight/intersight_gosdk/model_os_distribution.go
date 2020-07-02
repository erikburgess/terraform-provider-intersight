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

// OsDistribution Intersight has the distribution details for all the Intersight supported OS distributions. There will be a Distribution object for each supported OS.
type OsDistribution struct {
	MoBaseMo
	// The name of the OS distribution such as ESXi, CentOS.
	Name              *string                         `json:"Name,omitempty"`
	SupportedEditions *[]string                       `json:"SupportedEditions,omitempty"`
	Catalog           *OsCatalogRelationship          `json:"Catalog,omitempty"`
	Version           *HclOperatingSystemRelationship `json:"Version,omitempty"`
}

// NewOsDistribution instantiates a new OsDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsDistribution() *OsDistribution {
	this := OsDistribution{}
	return &this
}

// NewOsDistributionWithDefaults instantiates a new OsDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsDistributionWithDefaults() *OsDistribution {
	this := OsDistribution{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsDistribution) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsDistribution) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsDistribution) SetName(v string) {
	o.Name = &v
}

// GetSupportedEditions returns the SupportedEditions field value if set, zero value otherwise.
func (o *OsDistribution) GetSupportedEditions() []string {
	if o == nil || o.SupportedEditions == nil {
		var ret []string
		return ret
	}
	return *o.SupportedEditions
}

// GetSupportedEditionsOk returns a tuple with the SupportedEditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetSupportedEditionsOk() (*[]string, bool) {
	if o == nil || o.SupportedEditions == nil {
		return nil, false
	}
	return o.SupportedEditions, true
}

// HasSupportedEditions returns a boolean if a field has been set.
func (o *OsDistribution) HasSupportedEditions() bool {
	if o != nil && o.SupportedEditions != nil {
		return true
	}

	return false
}

// SetSupportedEditions gets a reference to the given []string and assigns it to the SupportedEditions field.
func (o *OsDistribution) SetSupportedEditions(v []string) {
	o.SupportedEditions = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *OsDistribution) GetCatalog() OsCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret OsCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetCatalogOk() (*OsCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *OsDistribution) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given OsCatalogRelationship and assigns it to the Catalog field.
func (o *OsDistribution) SetCatalog(v OsCatalogRelationship) {
	o.Catalog = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OsDistribution) GetVersion() HclOperatingSystemRelationship {
	if o == nil || o.Version == nil {
		var ret HclOperatingSystemRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsDistribution) GetVersionOk() (*HclOperatingSystemRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OsDistribution) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given HclOperatingSystemRelationship and assigns it to the Version field.
func (o *OsDistribution) SetVersion(v HclOperatingSystemRelationship) {
	o.Version = &v
}

func (o OsDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SupportedEditions != nil {
		toSerialize["SupportedEditions"] = o.SupportedEditions
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableOsDistribution struct {
	value *OsDistribution
	isSet bool
}

func (v NullableOsDistribution) Get() *OsDistribution {
	return v.value
}

func (v *NullableOsDistribution) Set(val *OsDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableOsDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableOsDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsDistribution(val *OsDistribution) *NullableOsDistribution {
	return &NullableOsDistribution{value: val, isSet: true}
}

func (v NullableOsDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
