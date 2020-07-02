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

// AssetCloudConnection CloudConnection provides the necessary details for Intersight to connect to and authenticate with a target at a well-known service address. The service address is inferred based upon the target type. For example Amazon Web Services.
type AssetCloudConnection struct {
	AssetConnection
	Credential *AssetCredential `json:"Credential,omitempty"`
}

// NewAssetCloudConnection instantiates a new AssetCloudConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCloudConnection() *AssetCloudConnection {
	this := AssetCloudConnection{}
	return &this
}

// NewAssetCloudConnectionWithDefaults instantiates a new AssetCloudConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCloudConnectionWithDefaults() *AssetCloudConnection {
	this := AssetCloudConnection{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *AssetCloudConnection) GetCredential() AssetCredential {
	if o == nil || o.Credential == nil {
		var ret AssetCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCloudConnection) GetCredentialOk() (*AssetCredential, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *AssetCloudConnection) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given AssetCredential and assigns it to the Credential field.
func (o *AssetCloudConnection) SetCredential(v AssetCredential) {
	o.Credential = &v
}

func (o AssetCloudConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetConnection, errAssetConnection := json.Marshal(o.AssetConnection)
	if errAssetConnection != nil {
		return []byte{}, errAssetConnection
	}
	errAssetConnection = json.Unmarshal([]byte(serializedAssetConnection), &toSerialize)
	if errAssetConnection != nil {
		return []byte{}, errAssetConnection
	}
	if o.Credential != nil {
		toSerialize["Credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableAssetCloudConnection struct {
	value *AssetCloudConnection
	isSet bool
}

func (v NullableAssetCloudConnection) Get() *AssetCloudConnection {
	return v.value
}

func (v *NullableAssetCloudConnection) Set(val *AssetCloudConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCloudConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCloudConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCloudConnection(val *AssetCloudConnection) *NullableAssetCloudConnection {
	return &NullableAssetCloudConnection{value: val, isSet: true}
}

func (v NullableAssetCloudConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCloudConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
