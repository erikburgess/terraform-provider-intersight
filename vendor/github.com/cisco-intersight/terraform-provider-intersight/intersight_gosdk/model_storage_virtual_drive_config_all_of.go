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

// StorageVirtualDriveConfigAllOf Definition of the list of properties defined in 'storage.VirtualDriveConfig', excluding properties defined in parent classes.
type StorageVirtualDriveConfigAllOf struct {
	// Access policy that host has on this virtual drive.
	AccessPolicy *string `json:"AccessPolicy,omitempty"`
	// The flag enables the use of this virtual drive as a boot drive.
	BootDrive *bool `json:"BootDrive,omitempty"`
	// Disk group policy that has the disk group in which this virtual drive needs to be created.
	DiskGroupName *string `json:"DiskGroupName,omitempty"`
	// Disk group policy that has the disk group in which this virtual drive needs to be created.
	DiskGroupPolicy *string `json:"DiskGroupPolicy,omitempty"`
	// The property expect disk cache policy.
	DriveCache *string `json:"DriveCache,omitempty"`
	// The flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.
	ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`
	// Desired IO mode - direct IO or cached IO.
	IoPolicy *string `json:"IoPolicy,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
	Name *string `json:"Name,omitempty"`
	// Read ahead mode to be used to read data from this virtual drive.
	ReadPolicy *string `json:"ReadPolicy,omitempty"`
	// Virtual drive size in MB. Size is mandatory field unless the 'Expand to Available' option is enabled.
	Size *int64 `json:"Size,omitempty"`
	// Write mode to be used to write data to this virtual drive.
	WritePolicy *string `json:"WritePolicy,omitempty"`
}

// NewStorageVirtualDriveConfigAllOf instantiates a new StorageVirtualDriveConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveConfigAllOf() *StorageVirtualDriveConfigAllOf {
	this := StorageVirtualDriveConfigAllOf{}
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var ioPolicy string = "Default"
	this.IoPolicy = &ioPolicy
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// NewStorageVirtualDriveConfigAllOfWithDefaults instantiates a new StorageVirtualDriveConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveConfigAllOfWithDefaults() *StorageVirtualDriveConfigAllOf {
	this := StorageVirtualDriveConfigAllOf{}
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var ioPolicy string = "Default"
	this.IoPolicy = &ioPolicy
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetAccessPolicy() string {
	if o == nil || o.AccessPolicy == nil {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetAccessPolicyOk() (*string, bool) {
	if o == nil || o.AccessPolicy == nil {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasAccessPolicy() bool {
	if o != nil && o.AccessPolicy != nil {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *StorageVirtualDriveConfigAllOf) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

// GetBootDrive returns the BootDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetBootDrive() bool {
	if o == nil || o.BootDrive == nil {
		var ret bool
		return ret
	}
	return *o.BootDrive
}

// GetBootDriveOk returns a tuple with the BootDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetBootDriveOk() (*bool, bool) {
	if o == nil || o.BootDrive == nil {
		return nil, false
	}
	return o.BootDrive, true
}

// HasBootDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasBootDrive() bool {
	if o != nil && o.BootDrive != nil {
		return true
	}

	return false
}

// SetBootDrive gets a reference to the given bool and assigns it to the BootDrive field.
func (o *StorageVirtualDriveConfigAllOf) SetBootDrive(v bool) {
	o.BootDrive = &v
}

// GetDiskGroupName returns the DiskGroupName field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetDiskGroupName() string {
	if o == nil || o.DiskGroupName == nil {
		var ret string
		return ret
	}
	return *o.DiskGroupName
}

// GetDiskGroupNameOk returns a tuple with the DiskGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetDiskGroupNameOk() (*string, bool) {
	if o == nil || o.DiskGroupName == nil {
		return nil, false
	}
	return o.DiskGroupName, true
}

// HasDiskGroupName returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasDiskGroupName() bool {
	if o != nil && o.DiskGroupName != nil {
		return true
	}

	return false
}

// SetDiskGroupName gets a reference to the given string and assigns it to the DiskGroupName field.
func (o *StorageVirtualDriveConfigAllOf) SetDiskGroupName(v string) {
	o.DiskGroupName = &v
}

// GetDiskGroupPolicy returns the DiskGroupPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetDiskGroupPolicy() string {
	if o == nil || o.DiskGroupPolicy == nil {
		var ret string
		return ret
	}
	return *o.DiskGroupPolicy
}

// GetDiskGroupPolicyOk returns a tuple with the DiskGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetDiskGroupPolicyOk() (*string, bool) {
	if o == nil || o.DiskGroupPolicy == nil {
		return nil, false
	}
	return o.DiskGroupPolicy, true
}

// HasDiskGroupPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasDiskGroupPolicy() bool {
	if o != nil && o.DiskGroupPolicy != nil {
		return true
	}

	return false
}

// SetDiskGroupPolicy gets a reference to the given string and assigns it to the DiskGroupPolicy field.
func (o *StorageVirtualDriveConfigAllOf) SetDiskGroupPolicy(v string) {
	o.DiskGroupPolicy = &v
}

// GetDriveCache returns the DriveCache field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetDriveCache() string {
	if o == nil || o.DriveCache == nil {
		var ret string
		return ret
	}
	return *o.DriveCache
}

// GetDriveCacheOk returns a tuple with the DriveCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetDriveCacheOk() (*string, bool) {
	if o == nil || o.DriveCache == nil {
		return nil, false
	}
	return o.DriveCache, true
}

// HasDriveCache returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasDriveCache() bool {
	if o != nil && o.DriveCache != nil {
		return true
	}

	return false
}

// SetDriveCache gets a reference to the given string and assigns it to the DriveCache field.
func (o *StorageVirtualDriveConfigAllOf) SetDriveCache(v string) {
	o.DriveCache = &v
}

// GetExpandToAvailable returns the ExpandToAvailable field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetExpandToAvailable() bool {
	if o == nil || o.ExpandToAvailable == nil {
		var ret bool
		return ret
	}
	return *o.ExpandToAvailable
}

// GetExpandToAvailableOk returns a tuple with the ExpandToAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetExpandToAvailableOk() (*bool, bool) {
	if o == nil || o.ExpandToAvailable == nil {
		return nil, false
	}
	return o.ExpandToAvailable, true
}

// HasExpandToAvailable returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasExpandToAvailable() bool {
	if o != nil && o.ExpandToAvailable != nil {
		return true
	}

	return false
}

// SetExpandToAvailable gets a reference to the given bool and assigns it to the ExpandToAvailable field.
func (o *StorageVirtualDriveConfigAllOf) SetExpandToAvailable(v bool) {
	o.ExpandToAvailable = &v
}

// GetIoPolicy returns the IoPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetIoPolicy() string {
	if o == nil || o.IoPolicy == nil {
		var ret string
		return ret
	}
	return *o.IoPolicy
}

// GetIoPolicyOk returns a tuple with the IoPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetIoPolicyOk() (*string, bool) {
	if o == nil || o.IoPolicy == nil {
		return nil, false
	}
	return o.IoPolicy, true
}

// HasIoPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasIoPolicy() bool {
	if o != nil && o.IoPolicy != nil {
		return true
	}

	return false
}

// SetIoPolicy gets a reference to the given string and assigns it to the IoPolicy field.
func (o *StorageVirtualDriveConfigAllOf) SetIoPolicy(v string) {
	o.IoPolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveConfigAllOf) SetName(v string) {
	o.Name = &v
}

// GetReadPolicy returns the ReadPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetReadPolicy() string {
	if o == nil || o.ReadPolicy == nil {
		var ret string
		return ret
	}
	return *o.ReadPolicy
}

// GetReadPolicyOk returns a tuple with the ReadPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetReadPolicyOk() (*string, bool) {
	if o == nil || o.ReadPolicy == nil {
		return nil, false
	}
	return o.ReadPolicy, true
}

// HasReadPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasReadPolicy() bool {
	if o != nil && o.ReadPolicy != nil {
		return true
	}

	return false
}

// SetReadPolicy gets a reference to the given string and assigns it to the ReadPolicy field.
func (o *StorageVirtualDriveConfigAllOf) SetReadPolicy(v string) {
	o.ReadPolicy = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageVirtualDriveConfigAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetWritePolicy returns the WritePolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfigAllOf) GetWritePolicy() string {
	if o == nil || o.WritePolicy == nil {
		var ret string
		return ret
	}
	return *o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfigAllOf) GetWritePolicyOk() (*string, bool) {
	if o == nil || o.WritePolicy == nil {
		return nil, false
	}
	return o.WritePolicy, true
}

// HasWritePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfigAllOf) HasWritePolicy() bool {
	if o != nil && o.WritePolicy != nil {
		return true
	}

	return false
}

// SetWritePolicy gets a reference to the given string and assigns it to the WritePolicy field.
func (o *StorageVirtualDriveConfigAllOf) SetWritePolicy(v string) {
	o.WritePolicy = &v
}

func (o StorageVirtualDriveConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessPolicy != nil {
		toSerialize["AccessPolicy"] = o.AccessPolicy
	}
	if o.BootDrive != nil {
		toSerialize["BootDrive"] = o.BootDrive
	}
	if o.DiskGroupName != nil {
		toSerialize["DiskGroupName"] = o.DiskGroupName
	}
	if o.DiskGroupPolicy != nil {
		toSerialize["DiskGroupPolicy"] = o.DiskGroupPolicy
	}
	if o.DriveCache != nil {
		toSerialize["DriveCache"] = o.DriveCache
	}
	if o.ExpandToAvailable != nil {
		toSerialize["ExpandToAvailable"] = o.ExpandToAvailable
	}
	if o.IoPolicy != nil {
		toSerialize["IoPolicy"] = o.IoPolicy
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ReadPolicy != nil {
		toSerialize["ReadPolicy"] = o.ReadPolicy
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.WritePolicy != nil {
		toSerialize["WritePolicy"] = o.WritePolicy
	}
	return json.Marshal(toSerialize)
}

type NullableStorageVirtualDriveConfigAllOf struct {
	value *StorageVirtualDriveConfigAllOf
	isSet bool
}

func (v NullableStorageVirtualDriveConfigAllOf) Get() *StorageVirtualDriveConfigAllOf {
	return v.value
}

func (v *NullableStorageVirtualDriveConfigAllOf) Set(val *StorageVirtualDriveConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveConfigAllOf(val *StorageVirtualDriveConfigAllOf) *NullableStorageVirtualDriveConfigAllOf {
	return &NullableStorageVirtualDriveConfigAllOf{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
