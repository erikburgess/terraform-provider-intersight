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
	"time"
)

// SoftwarerepositoryRelease A Cisco release containing one or more firmware images. Cisco releases images for rack server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
type SoftwarerepositoryRelease struct {
	MoBaseMo
	// The date when the file was released or distributed by its vendor.
	ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`
	// The URL for the release notes of this image.
	ReleaseNotesUrl *string   `json:"ReleaseNotesUrl,omitempty"`
	SupportedModels *[]string `json:"SupportedModels,omitempty"`
	// The platform release type for which the images are released. This can be a fabric switch or compute server hardware.
	Type *string `json:"Type,omitempty"`
	// Cisco provided release version.
	Version *string                                `json:"Version,omitempty"`
	Catalog *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
}

// NewSoftwarerepositoryRelease instantiates a new SoftwarerepositoryRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryRelease() *SoftwarerepositoryRelease {
	this := SoftwarerepositoryRelease{}
	var type_ string = "FabricSwitch"
	this.Type = &type_
	return &this
}

// NewSoftwarerepositoryReleaseWithDefaults instantiates a new SoftwarerepositoryRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryReleaseWithDefaults() *SoftwarerepositoryRelease {
	this := SoftwarerepositoryRelease{}
	var type_ string = "FabricSwitch"
	this.Type = &type_
	return &this
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetReleaseDate() time.Time {
	if o == nil || o.ReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *SoftwarerepositoryRelease) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetReleaseNotesUrl returns the ReleaseNotesUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetReleaseNotesUrl() string {
	if o == nil || o.ReleaseNotesUrl == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNotesUrl
}

// GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetReleaseNotesUrlOk() (*string, bool) {
	if o == nil || o.ReleaseNotesUrl == nil {
		return nil, false
	}
	return o.ReleaseNotesUrl, true
}

// HasReleaseNotesUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasReleaseNotesUrl() bool {
	if o != nil && o.ReleaseNotesUrl != nil {
		return true
	}

	return false
}

// SetReleaseNotesUrl gets a reference to the given string and assigns it to the ReleaseNotesUrl field.
func (o *SoftwarerepositoryRelease) SetReleaseNotesUrl(v string) {
	o.ReleaseNotesUrl = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetSupportedModels() []string {
	if o == nil || o.SupportedModels == nil {
		var ret []string
		return ret
	}
	return *o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryRelease) SetSupportedModels(v []string) {
	o.SupportedModels = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SoftwarerepositoryRelease) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryRelease) SetVersion(v string) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwarerepositoryRelease) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryRelease) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwarerepositoryRelease) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwarerepositoryRelease) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwarerepositoryRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.ReleaseDate != nil {
		toSerialize["ReleaseDate"] = o.ReleaseDate
	}
	if o.ReleaseNotesUrl != nil {
		toSerialize["ReleaseNotesUrl"] = o.ReleaseNotesUrl
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwarerepositoryRelease struct {
	value *SoftwarerepositoryRelease
	isSet bool
}

func (v NullableSoftwarerepositoryRelease) Get() *SoftwarerepositoryRelease {
	return v.value
}

func (v *NullableSoftwarerepositoryRelease) Set(val *SoftwarerepositoryRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryRelease(val *SoftwarerepositoryRelease) *NullableSoftwarerepositoryRelease {
	return &NullableSoftwarerepositoryRelease{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
