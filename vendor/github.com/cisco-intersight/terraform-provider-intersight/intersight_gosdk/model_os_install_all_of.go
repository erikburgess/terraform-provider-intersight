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

// OsInstallAllOf Definition of the list of properties defined in 'os.Install', excluding properties defined in parent classes.
type OsInstallAllOf struct {
	// The name of the OS install configuration.
	Name              *string                                                      `json:"Name,omitempty"`
	ConfigurationFile *OsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
	Image             *SoftwarerepositoryOperatingSystemFileRelationship           `json:"Image,omitempty"`
	Organization      *OrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
	OsduImage         *FirmwareServerConfigurationUtilityDistributableRelationship `json:"OsduImage,omitempty"`
	Server            *ComputePhysicalRelationship                                 `json:"Server,omitempty"`
	WorkflowInfo      *WorkflowWorkflowInfoRelationship                            `json:"WorkflowInfo,omitempty"`
}

// NewOsInstallAllOf instantiates a new OsInstallAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsInstallAllOf() *OsInstallAllOf {
	this := OsInstallAllOf{}
	return &this
}

// NewOsInstallAllOfWithDefaults instantiates a new OsInstallAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsInstallAllOfWithDefaults() *OsInstallAllOf {
	this := OsInstallAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsInstallAllOf) SetName(v string) {
	o.Name = &v
}

// GetConfigurationFile returns the ConfigurationFile field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetConfigurationFile() OsConfigurationFileRelationship {
	if o == nil || o.ConfigurationFile == nil {
		var ret OsConfigurationFileRelationship
		return ret
	}
	return *o.ConfigurationFile
}

// GetConfigurationFileOk returns a tuple with the ConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool) {
	if o == nil || o.ConfigurationFile == nil {
		return nil, false
	}
	return o.ConfigurationFile, true
}

// HasConfigurationFile returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasConfigurationFile() bool {
	if o != nil && o.ConfigurationFile != nil {
		return true
	}

	return false
}

// SetConfigurationFile gets a reference to the given OsConfigurationFileRelationship and assigns it to the ConfigurationFile field.
func (o *OsInstallAllOf) SetConfigurationFile(v OsConfigurationFileRelationship) {
	o.ConfigurationFile = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetImage() SoftwarerepositoryOperatingSystemFileRelationship {
	if o == nil || o.Image == nil {
		var ret SoftwarerepositoryOperatingSystemFileRelationship
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given SoftwarerepositoryOperatingSystemFileRelationship and assigns it to the Image field.
func (o *OsInstallAllOf) SetImage(v SoftwarerepositoryOperatingSystemFileRelationship) {
	o.Image = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OsInstallAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetOsduImage returns the OsduImage field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetOsduImage() FirmwareServerConfigurationUtilityDistributableRelationship {
	if o == nil || o.OsduImage == nil {
		var ret FirmwareServerConfigurationUtilityDistributableRelationship
		return ret
	}
	return *o.OsduImage
}

// GetOsduImageOk returns a tuple with the OsduImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetOsduImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool) {
	if o == nil || o.OsduImage == nil {
		return nil, false
	}
	return o.OsduImage, true
}

// HasOsduImage returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasOsduImage() bool {
	if o != nil && o.OsduImage != nil {
		return true
	}

	return false
}

// SetOsduImage gets a reference to the given FirmwareServerConfigurationUtilityDistributableRelationship and assigns it to the OsduImage field.
func (o *OsInstallAllOf) SetOsduImage(v FirmwareServerConfigurationUtilityDistributableRelationship) {
	o.OsduImage = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *OsInstallAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *OsInstallAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o OsInstallAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ConfigurationFile != nil {
		toSerialize["ConfigurationFile"] = o.ConfigurationFile
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OsduImage != nil {
		toSerialize["OsduImage"] = o.OsduImage
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}
	return json.Marshal(toSerialize)
}

type NullableOsInstallAllOf struct {
	value *OsInstallAllOf
	isSet bool
}

func (v NullableOsInstallAllOf) Get() *OsInstallAllOf {
	return v.value
}

func (v *NullableOsInstallAllOf) Set(val *OsInstallAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsInstallAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsInstallAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsInstallAllOf(val *OsInstallAllOf) *NullableOsInstallAllOf {
	return &NullableOsInstallAllOf{value: val, isSet: true}
}

func (v NullableOsInstallAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsInstallAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
