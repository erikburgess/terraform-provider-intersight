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
)

// MetaDefinitionAllOf Definition of the list of properties defined in 'meta.Definition', excluding properties defined in parent classes.
type MetaDefinitionAllOf struct {
	AccessPrivileges *[]MetaAccessPrivilege `json:"AccessPrivileges,omitempty"`
	AncestorClasses  *[]string              `json:"AncestorClasses,omitempty"`
	// Boolean flag to specify whether the meta class is a concrete class or not.
	IsConcrete *bool `json:"IsConcrete,omitempty"`
	// Indicates whether the meta class is a complex type or managed object.
	MetaType *string `json:"MetaType,omitempty"`
	// The fully-qualified class name of the Managed Object or complex type. For example, \"compute:Blade\" where the Managed Object is \"Blade\" and the package is 'compute'.
	Name *string `json:"Name,omitempty"`
	// The namespace of the meta.
	Namespace *string `json:"Namespace,omitempty"`
	// The fully-qualified name of the parent metaclass in the class inheritance hierarchy.
	ParentClass *string `json:"ParentClass,omitempty"`
	// Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations.
	PermissionSupported *bool                 `json:"PermissionSupported,omitempty"`
	Properties          *[]MetaPropDefinition `json:"Properties,omitempty"`
	// Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource=yes. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set.
	RbacResource  *bool                         `json:"RbacResource,omitempty"`
	Relationships *[]MetaRelationshipDefinition `json:"Relationships,omitempty"`
	// Restful URL path for the meta.
	RestPath *string `json:"RestPath,omitempty"`
	// The version of the service that defines the meta-data.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaDefinitionAllOf MetaDefinitionAllOf

// NewMetaDefinitionAllOf instantiates a new MetaDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaDefinitionAllOf() *MetaDefinitionAllOf {
	this := MetaDefinitionAllOf{}
	var metaType string = "ManagedObject"
	this.MetaType = &metaType
	return &this
}

// NewMetaDefinitionAllOfWithDefaults instantiates a new MetaDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaDefinitionAllOfWithDefaults() *MetaDefinitionAllOf {
	this := MetaDefinitionAllOf{}
	var metaType string = "ManagedObject"
	this.MetaType = &metaType
	return &this
}

// GetAccessPrivileges returns the AccessPrivileges field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetAccessPrivileges() []MetaAccessPrivilege {
	if o == nil || o.AccessPrivileges == nil {
		var ret []MetaAccessPrivilege
		return ret
	}
	return *o.AccessPrivileges
}

// GetAccessPrivilegesOk returns a tuple with the AccessPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetAccessPrivilegesOk() (*[]MetaAccessPrivilege, bool) {
	if o == nil || o.AccessPrivileges == nil {
		return nil, false
	}
	return o.AccessPrivileges, true
}

// HasAccessPrivileges returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasAccessPrivileges() bool {
	if o != nil && o.AccessPrivileges != nil {
		return true
	}

	return false
}

// SetAccessPrivileges gets a reference to the given []MetaAccessPrivilege and assigns it to the AccessPrivileges field.
func (o *MetaDefinitionAllOf) SetAccessPrivileges(v []MetaAccessPrivilege) {
	o.AccessPrivileges = &v
}

// GetAncestorClasses returns the AncestorClasses field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetAncestorClasses() []string {
	if o == nil || o.AncestorClasses == nil {
		var ret []string
		return ret
	}
	return *o.AncestorClasses
}

// GetAncestorClassesOk returns a tuple with the AncestorClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetAncestorClassesOk() (*[]string, bool) {
	if o == nil || o.AncestorClasses == nil {
		return nil, false
	}
	return o.AncestorClasses, true
}

// HasAncestorClasses returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasAncestorClasses() bool {
	if o != nil && o.AncestorClasses != nil {
		return true
	}

	return false
}

// SetAncestorClasses gets a reference to the given []string and assigns it to the AncestorClasses field.
func (o *MetaDefinitionAllOf) SetAncestorClasses(v []string) {
	o.AncestorClasses = &v
}

// GetIsConcrete returns the IsConcrete field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetIsConcrete() bool {
	if o == nil || o.IsConcrete == nil {
		var ret bool
		return ret
	}
	return *o.IsConcrete
}

// GetIsConcreteOk returns a tuple with the IsConcrete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetIsConcreteOk() (*bool, bool) {
	if o == nil || o.IsConcrete == nil {
		return nil, false
	}
	return o.IsConcrete, true
}

// HasIsConcrete returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasIsConcrete() bool {
	if o != nil && o.IsConcrete != nil {
		return true
	}

	return false
}

// SetIsConcrete gets a reference to the given bool and assigns it to the IsConcrete field.
func (o *MetaDefinitionAllOf) SetIsConcrete(v bool) {
	o.IsConcrete = &v
}

// GetMetaType returns the MetaType field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetMetaType() string {
	if o == nil || o.MetaType == nil {
		var ret string
		return ret
	}
	return *o.MetaType
}

// GetMetaTypeOk returns a tuple with the MetaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetMetaTypeOk() (*string, bool) {
	if o == nil || o.MetaType == nil {
		return nil, false
	}
	return o.MetaType, true
}

// HasMetaType returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasMetaType() bool {
	if o != nil && o.MetaType != nil {
		return true
	}

	return false
}

// SetMetaType gets a reference to the given string and assigns it to the MetaType field.
func (o *MetaDefinitionAllOf) SetMetaType(v string) {
	o.MetaType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *MetaDefinitionAllOf) SetNamespace(v string) {
	o.Namespace = &v
}

// GetParentClass returns the ParentClass field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetParentClass() string {
	if o == nil || o.ParentClass == nil {
		var ret string
		return ret
	}
	return *o.ParentClass
}

// GetParentClassOk returns a tuple with the ParentClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetParentClassOk() (*string, bool) {
	if o == nil || o.ParentClass == nil {
		return nil, false
	}
	return o.ParentClass, true
}

// HasParentClass returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasParentClass() bool {
	if o != nil && o.ParentClass != nil {
		return true
	}

	return false
}

// SetParentClass gets a reference to the given string and assigns it to the ParentClass field.
func (o *MetaDefinitionAllOf) SetParentClass(v string) {
	o.ParentClass = &v
}

// GetPermissionSupported returns the PermissionSupported field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetPermissionSupported() bool {
	if o == nil || o.PermissionSupported == nil {
		var ret bool
		return ret
	}
	return *o.PermissionSupported
}

// GetPermissionSupportedOk returns a tuple with the PermissionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetPermissionSupportedOk() (*bool, bool) {
	if o == nil || o.PermissionSupported == nil {
		return nil, false
	}
	return o.PermissionSupported, true
}

// HasPermissionSupported returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasPermissionSupported() bool {
	if o != nil && o.PermissionSupported != nil {
		return true
	}

	return false
}

// SetPermissionSupported gets a reference to the given bool and assigns it to the PermissionSupported field.
func (o *MetaDefinitionAllOf) SetPermissionSupported(v bool) {
	o.PermissionSupported = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetProperties() []MetaPropDefinition {
	if o == nil || o.Properties == nil {
		var ret []MetaPropDefinition
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetPropertiesOk() (*[]MetaPropDefinition, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []MetaPropDefinition and assigns it to the Properties field.
func (o *MetaDefinitionAllOf) SetProperties(v []MetaPropDefinition) {
	o.Properties = &v
}

// GetRbacResource returns the RbacResource field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetRbacResource() bool {
	if o == nil || o.RbacResource == nil {
		var ret bool
		return ret
	}
	return *o.RbacResource
}

// GetRbacResourceOk returns a tuple with the RbacResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetRbacResourceOk() (*bool, bool) {
	if o == nil || o.RbacResource == nil {
		return nil, false
	}
	return o.RbacResource, true
}

// HasRbacResource returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasRbacResource() bool {
	if o != nil && o.RbacResource != nil {
		return true
	}

	return false
}

// SetRbacResource gets a reference to the given bool and assigns it to the RbacResource field.
func (o *MetaDefinitionAllOf) SetRbacResource(v bool) {
	o.RbacResource = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetRelationships() []MetaRelationshipDefinition {
	if o == nil || o.Relationships == nil {
		var ret []MetaRelationshipDefinition
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetRelationshipsOk() (*[]MetaRelationshipDefinition, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []MetaRelationshipDefinition and assigns it to the Relationships field.
func (o *MetaDefinitionAllOf) SetRelationships(v []MetaRelationshipDefinition) {
	o.Relationships = &v
}

// GetRestPath returns the RestPath field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetRestPath() string {
	if o == nil || o.RestPath == nil {
		var ret string
		return ret
	}
	return *o.RestPath
}

// GetRestPathOk returns a tuple with the RestPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetRestPathOk() (*string, bool) {
	if o == nil || o.RestPath == nil {
		return nil, false
	}
	return o.RestPath, true
}

// HasRestPath returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasRestPath() bool {
	if o != nil && o.RestPath != nil {
		return true
	}

	return false
}

// SetRestPath gets a reference to the given string and assigns it to the RestPath field.
func (o *MetaDefinitionAllOf) SetRestPath(v string) {
	o.RestPath = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MetaDefinitionAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDefinitionAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MetaDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MetaDefinitionAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o MetaDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessPrivileges != nil {
		toSerialize["AccessPrivileges"] = o.AccessPrivileges
	}
	if o.AncestorClasses != nil {
		toSerialize["AncestorClasses"] = o.AncestorClasses
	}
	if o.IsConcrete != nil {
		toSerialize["IsConcrete"] = o.IsConcrete
	}
	if o.MetaType != nil {
		toSerialize["MetaType"] = o.MetaType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.ParentClass != nil {
		toSerialize["ParentClass"] = o.ParentClass
	}
	if o.PermissionSupported != nil {
		toSerialize["PermissionSupported"] = o.PermissionSupported
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	if o.RbacResource != nil {
		toSerialize["RbacResource"] = o.RbacResource
	}
	if o.Relationships != nil {
		toSerialize["Relationships"] = o.Relationships
	}
	if o.RestPath != nil {
		toSerialize["RestPath"] = o.RestPath
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMetaDefinitionAllOf := _MetaDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varMetaDefinitionAllOf); err == nil {
		*o = MetaDefinitionAllOf(varMetaDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AccessPrivileges")
		delete(additionalProperties, "AncestorClasses")
		delete(additionalProperties, "IsConcrete")
		delete(additionalProperties, "MetaType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "ParentClass")
		delete(additionalProperties, "PermissionSupported")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "RbacResource")
		delete(additionalProperties, "Relationships")
		delete(additionalProperties, "RestPath")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetaDefinitionAllOf struct {
	value *MetaDefinitionAllOf
	isSet bool
}

func (v NullableMetaDefinitionAllOf) Get() *MetaDefinitionAllOf {
	return v.value
}

func (v *NullableMetaDefinitionAllOf) Set(val *MetaDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaDefinitionAllOf(val *MetaDefinitionAllOf) *NullableMetaDefinitionAllOf {
	return &NullableMetaDefinitionAllOf{value: val, isSet: true}
}

func (v NullableMetaDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
