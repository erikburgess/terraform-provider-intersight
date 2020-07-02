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

// ConnectorpackConnectorPackUpdateAllOf Definition of the list of properties defined in 'connectorpack.ConnectorPackUpdate', excluding properties defined in parent classes.
type ConnectorpackConnectorPackUpdateAllOf struct {
	// Version of connector pack currently running in UCS Director.
	CurrentVersion *string `json:"CurrentVersion,omitempty"`
	// Name of the connector pack.
	Name *string `json:"Name,omitempty"`
	// Version of connector pack to be installed in the next upgrade cycle.
	NewVersion *string `json:"NewVersion,omitempty"`
}

// NewConnectorpackConnectorPackUpdateAllOf instantiates a new ConnectorpackConnectorPackUpdateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackConnectorPackUpdateAllOf() *ConnectorpackConnectorPackUpdateAllOf {
	this := ConnectorpackConnectorPackUpdateAllOf{}
	return &this
}

// NewConnectorpackConnectorPackUpdateAllOfWithDefaults instantiates a new ConnectorpackConnectorPackUpdateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackConnectorPackUpdateAllOfWithDefaults() *ConnectorpackConnectorPackUpdateAllOf {
	this := ConnectorpackConnectorPackUpdateAllOf{}
	return &this
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetCurrentVersion() string {
	if o == nil || o.CurrentVersion == nil {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetCurrentVersionOk() (*string, bool) {
	if o == nil || o.CurrentVersion == nil {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) HasCurrentVersion() bool {
	if o != nil && o.CurrentVersion != nil {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *ConnectorpackConnectorPackUpdateAllOf) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorpackConnectorPackUpdateAllOf) SetName(v string) {
	o.Name = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetNewVersion() string {
	if o == nil || o.NewVersion == nil {
		var ret string
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetNewVersionOk() (*string, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given string and assigns it to the NewVersion field.
func (o *ConnectorpackConnectorPackUpdateAllOf) SetNewVersion(v string) {
	o.NewVersion = &v
}

func (o ConnectorpackConnectorPackUpdateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentVersion != nil {
		toSerialize["CurrentVersion"] = o.CurrentVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NewVersion != nil {
		toSerialize["NewVersion"] = o.NewVersion
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorpackConnectorPackUpdateAllOf struct {
	value *ConnectorpackConnectorPackUpdateAllOf
	isSet bool
}

func (v NullableConnectorpackConnectorPackUpdateAllOf) Get() *ConnectorpackConnectorPackUpdateAllOf {
	return v.value
}

func (v *NullableConnectorpackConnectorPackUpdateAllOf) Set(val *ConnectorpackConnectorPackUpdateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackConnectorPackUpdateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackConnectorPackUpdateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackConnectorPackUpdateAllOf(val *ConnectorpackConnectorPackUpdateAllOf) *NullableConnectorpackConnectorPackUpdateAllOf {
	return &NullableConnectorpackConnectorPackUpdateAllOf{value: val, isSet: true}
}

func (v NullableConnectorpackConnectorPackUpdateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackConnectorPackUpdateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
