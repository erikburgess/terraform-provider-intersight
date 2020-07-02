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

// HclHardwareCompatibilityProfile Profile giving server hardware details, OS details and UCS Version details.
type HclHardwareCompatibilityProfile struct {
	MoBaseComplexType
	// Url for the ISO with the drivers supported for the server.
	DriverIsoUrl *string `json:"DriverIsoUrl,omitempty"`
	// Error code indicating the compatibility status.
	ErrorCode *string `json:"ErrorCode,omitempty"`
	// Identifier of the hardware compatibility profile.
	Id *string `json:"Id,omitempty"`
	// Vendor of the Operating System running on the server.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating System running on the server.
	OsVersion *string `json:"OsVersion,omitempty"`
	// Model of the processor present in the server.
	ProcessorModel *string       `json:"ProcessorModel,omitempty"`
	Products       *[]HclProduct `json:"Products,omitempty"`
	// Model of the server as returned by UCSM/CIMC XML API.
	ServerModel *string `json:"ServerModel,omitempty"`
	// Revision of the server model.
	ServerRevision *string `json:"ServerRevision,omitempty"`
	// Version of the UCS software.
	UcsVersion *string `json:"UcsVersion,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	VersionType *string `json:"VersionType,omitempty"`
}

// NewHclHardwareCompatibilityProfile instantiates a new HclHardwareCompatibilityProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclHardwareCompatibilityProfile() *HclHardwareCompatibilityProfile {
	this := HclHardwareCompatibilityProfile{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	var versionType string = "UCSM"
	this.VersionType = &versionType
	return &this
}

// NewHclHardwareCompatibilityProfileWithDefaults instantiates a new HclHardwareCompatibilityProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclHardwareCompatibilityProfileWithDefaults() *HclHardwareCompatibilityProfile {
	this := HclHardwareCompatibilityProfile{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	var versionType string = "UCSM"
	this.VersionType = &versionType
	return &this
}

// GetDriverIsoUrl returns the DriverIsoUrl field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrl() string {
	if o == nil || o.DriverIsoUrl == nil {
		var ret string
		return ret
	}
	return *o.DriverIsoUrl
}

// GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrlOk() (*string, bool) {
	if o == nil || o.DriverIsoUrl == nil {
		return nil, false
	}
	return o.DriverIsoUrl, true
}

// HasDriverIsoUrl returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasDriverIsoUrl() bool {
	if o != nil && o.DriverIsoUrl != nil {
		return true
	}

	return false
}

// SetDriverIsoUrl gets a reference to the given string and assigns it to the DriverIsoUrl field.
func (o *HclHardwareCompatibilityProfile) SetDriverIsoUrl(v string) {
	o.DriverIsoUrl = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclHardwareCompatibilityProfile) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclHardwareCompatibilityProfile) SetId(v string) {
	o.Id = &v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclHardwareCompatibilityProfile) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclHardwareCompatibilityProfile) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProcessorModel returns the ProcessorModel field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetProcessorModel() string {
	if o == nil || o.ProcessorModel == nil {
		var ret string
		return ret
	}
	return *o.ProcessorModel
}

// GetProcessorModelOk returns a tuple with the ProcessorModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetProcessorModelOk() (*string, bool) {
	if o == nil || o.ProcessorModel == nil {
		return nil, false
	}
	return o.ProcessorModel, true
}

// HasProcessorModel returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasProcessorModel() bool {
	if o != nil && o.ProcessorModel != nil {
		return true
	}

	return false
}

// SetProcessorModel gets a reference to the given string and assigns it to the ProcessorModel field.
func (o *HclHardwareCompatibilityProfile) SetProcessorModel(v string) {
	o.ProcessorModel = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetProducts() []HclProduct {
	if o == nil || o.Products == nil {
		var ret []HclProduct
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetProductsOk() (*[]HclProduct, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []HclProduct and assigns it to the Products field.
func (o *HclHardwareCompatibilityProfile) SetProducts(v []HclProduct) {
	o.Products = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HclHardwareCompatibilityProfile) SetServerModel(v string) {
	o.ServerModel = &v
}

// GetServerRevision returns the ServerRevision field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetServerRevision() string {
	if o == nil || o.ServerRevision == nil {
		var ret string
		return ret
	}
	return *o.ServerRevision
}

// GetServerRevisionOk returns a tuple with the ServerRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetServerRevisionOk() (*string, bool) {
	if o == nil || o.ServerRevision == nil {
		return nil, false
	}
	return o.ServerRevision, true
}

// HasServerRevision returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasServerRevision() bool {
	if o != nil && o.ServerRevision != nil {
		return true
	}

	return false
}

// SetServerRevision gets a reference to the given string and assigns it to the ServerRevision field.
func (o *HclHardwareCompatibilityProfile) SetServerRevision(v string) {
	o.ServerRevision = &v
}

// GetUcsVersion returns the UcsVersion field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetUcsVersion() string {
	if o == nil || o.UcsVersion == nil {
		var ret string
		return ret
	}
	return *o.UcsVersion
}

// GetUcsVersionOk returns a tuple with the UcsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetUcsVersionOk() (*string, bool) {
	if o == nil || o.UcsVersion == nil {
		return nil, false
	}
	return o.UcsVersion, true
}

// HasUcsVersion returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasUcsVersion() bool {
	if o != nil && o.UcsVersion != nil {
		return true
	}

	return false
}

// SetUcsVersion gets a reference to the given string and assigns it to the UcsVersion field.
func (o *HclHardwareCompatibilityProfile) SetUcsVersion(v string) {
	o.UcsVersion = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *HclHardwareCompatibilityProfile) SetVersionType(v string) {
	o.VersionType = &v
}

func (o HclHardwareCompatibilityProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.DriverIsoUrl != nil {
		toSerialize["DriverIsoUrl"] = o.DriverIsoUrl
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.OsVendor != nil {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.ProcessorModel != nil {
		toSerialize["ProcessorModel"] = o.ProcessorModel
	}
	if o.Products != nil {
		toSerialize["Products"] = o.Products
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	if o.ServerRevision != nil {
		toSerialize["ServerRevision"] = o.ServerRevision
	}
	if o.UcsVersion != nil {
		toSerialize["UcsVersion"] = o.UcsVersion
	}
	if o.VersionType != nil {
		toSerialize["VersionType"] = o.VersionType
	}
	return json.Marshal(toSerialize)
}

type NullableHclHardwareCompatibilityProfile struct {
	value *HclHardwareCompatibilityProfile
	isSet bool
}

func (v NullableHclHardwareCompatibilityProfile) Get() *HclHardwareCompatibilityProfile {
	return v.value
}

func (v *NullableHclHardwareCompatibilityProfile) Set(val *HclHardwareCompatibilityProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHardwareCompatibilityProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHardwareCompatibilityProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHardwareCompatibilityProfile(val *HclHardwareCompatibilityProfile) *NullableHclHardwareCompatibilityProfile {
	return &NullableHclHardwareCompatibilityProfile{value: val, isSet: true}
}

func (v NullableHclHardwareCompatibilityProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHardwareCompatibilityProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}