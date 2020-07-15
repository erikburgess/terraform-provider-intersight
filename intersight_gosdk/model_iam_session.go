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
	"reflect"
	"strings"
	"time"
)

// IamSession The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
type IamSession struct {
	MoBaseMo
	AccountPermissions *[]IamAccountPermissions `json:"AccountPermissions,omitempty"`
	// The user agent IP address from which the session is launched.
	ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
	// Expiration time for the session.
	Expiration *time.Time `json:"Expiration,omitempty"`
	// Idle time expiration for the session.
	IdleTimeExpiration *time.Time `json:"IdleTimeExpiration,omitempty"`
	// The client address from which last login is initiated.
	LastLoginClient *string `json:"LastLoginClient,omitempty"`
	// The last login time for user.
	LastLoginTime        *time.Time                 `json:"LastLoginTime,omitempty"`
	Permission           *IamPermissionRelationship `json:"Permission,omitempty"`
	User                 *IamUserRelationship       `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSession IamSession

// NewIamSession instantiates a new IamSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSession() *IamSession {
	this := IamSession{}
	return &this
}

// NewIamSessionWithDefaults instantiates a new IamSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSessionWithDefaults() *IamSession {
	this := IamSession{}
	return &this
}

// GetAccountPermissions returns the AccountPermissions field value if set, zero value otherwise.
func (o *IamSession) GetAccountPermissions() []IamAccountPermissions {
	if o == nil || o.AccountPermissions == nil {
		var ret []IamAccountPermissions
		return ret
	}
	return *o.AccountPermissions
}

// GetAccountPermissionsOk returns a tuple with the AccountPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetAccountPermissionsOk() (*[]IamAccountPermissions, bool) {
	if o == nil || o.AccountPermissions == nil {
		return nil, false
	}
	return o.AccountPermissions, true
}

// HasAccountPermissions returns a boolean if a field has been set.
func (o *IamSession) HasAccountPermissions() bool {
	if o != nil && o.AccountPermissions != nil {
		return true
	}

	return false
}

// SetAccountPermissions gets a reference to the given []IamAccountPermissions and assigns it to the AccountPermissions field.
func (o *IamSession) SetAccountPermissions(v []IamAccountPermissions) {
	o.AccountPermissions = &v
}

// GetClientIpAddress returns the ClientIpAddress field value if set, zero value otherwise.
func (o *IamSession) GetClientIpAddress() string {
	if o == nil || o.ClientIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ClientIpAddress
}

// GetClientIpAddressOk returns a tuple with the ClientIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetClientIpAddressOk() (*string, bool) {
	if o == nil || o.ClientIpAddress == nil {
		return nil, false
	}
	return o.ClientIpAddress, true
}

// HasClientIpAddress returns a boolean if a field has been set.
func (o *IamSession) HasClientIpAddress() bool {
	if o != nil && o.ClientIpAddress != nil {
		return true
	}

	return false
}

// SetClientIpAddress gets a reference to the given string and assigns it to the ClientIpAddress field.
func (o *IamSession) SetClientIpAddress(v string) {
	o.ClientIpAddress = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *IamSession) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *IamSession) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *IamSession) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetIdleTimeExpiration returns the IdleTimeExpiration field value if set, zero value otherwise.
func (o *IamSession) GetIdleTimeExpiration() time.Time {
	if o == nil || o.IdleTimeExpiration == nil {
		var ret time.Time
		return ret
	}
	return *o.IdleTimeExpiration
}

// GetIdleTimeExpirationOk returns a tuple with the IdleTimeExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetIdleTimeExpirationOk() (*time.Time, bool) {
	if o == nil || o.IdleTimeExpiration == nil {
		return nil, false
	}
	return o.IdleTimeExpiration, true
}

// HasIdleTimeExpiration returns a boolean if a field has been set.
func (o *IamSession) HasIdleTimeExpiration() bool {
	if o != nil && o.IdleTimeExpiration != nil {
		return true
	}

	return false
}

// SetIdleTimeExpiration gets a reference to the given time.Time and assigns it to the IdleTimeExpiration field.
func (o *IamSession) SetIdleTimeExpiration(v time.Time) {
	o.IdleTimeExpiration = &v
}

// GetLastLoginClient returns the LastLoginClient field value if set, zero value otherwise.
func (o *IamSession) GetLastLoginClient() string {
	if o == nil || o.LastLoginClient == nil {
		var ret string
		return ret
	}
	return *o.LastLoginClient
}

// GetLastLoginClientOk returns a tuple with the LastLoginClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetLastLoginClientOk() (*string, bool) {
	if o == nil || o.LastLoginClient == nil {
		return nil, false
	}
	return o.LastLoginClient, true
}

// HasLastLoginClient returns a boolean if a field has been set.
func (o *IamSession) HasLastLoginClient() bool {
	if o != nil && o.LastLoginClient != nil {
		return true
	}

	return false
}

// SetLastLoginClient gets a reference to the given string and assigns it to the LastLoginClient field.
func (o *IamSession) SetLastLoginClient(v string) {
	o.LastLoginClient = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *IamSession) GetLastLoginTime() time.Time {
	if o == nil || o.LastLoginTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetLastLoginTimeOk() (*time.Time, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *IamSession) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given time.Time and assigns it to the LastLoginTime field.
func (o *IamSession) SetLastLoginTime(v time.Time) {
	o.LastLoginTime = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamSession) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamSession) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamSession) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamSession) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSession) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamSession) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamSession) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.AccountPermissions != nil {
		toSerialize["AccountPermissions"] = o.AccountPermissions
	}
	if o.ClientIpAddress != nil {
		toSerialize["ClientIpAddress"] = o.ClientIpAddress
	}
	if o.Expiration != nil {
		toSerialize["Expiration"] = o.Expiration
	}
	if o.IdleTimeExpiration != nil {
		toSerialize["IdleTimeExpiration"] = o.IdleTimeExpiration
	}
	if o.LastLoginClient != nil {
		toSerialize["LastLoginClient"] = o.LastLoginClient
	}
	if o.LastLoginTime != nil {
		toSerialize["LastLoginTime"] = o.LastLoginTime
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSession) UnmarshalJSON(bytes []byte) (err error) {
	type IamSessionWithoutEmbeddedStruct struct {
		AccountPermissions *[]IamAccountPermissions `json:"AccountPermissions,omitempty"`
		// The user agent IP address from which the session is launched.
		ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
		// Expiration time for the session.
		Expiration *time.Time `json:"Expiration,omitempty"`
		// Idle time expiration for the session.
		IdleTimeExpiration *time.Time `json:"IdleTimeExpiration,omitempty"`
		// The client address from which last login is initiated.
		LastLoginClient *string `json:"LastLoginClient,omitempty"`
		// The last login time for user.
		LastLoginTime *time.Time                 `json:"LastLoginTime,omitempty"`
		Permission    *IamPermissionRelationship `json:"Permission,omitempty"`
		User          *IamUserRelationship       `json:"User,omitempty"`
	}

	varIamSessionWithoutEmbeddedStruct := IamSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamSessionWithoutEmbeddedStruct)
	if err == nil {
		varIamSession := _IamSession{}
		varIamSession.AccountPermissions = varIamSessionWithoutEmbeddedStruct.AccountPermissions
		varIamSession.ClientIpAddress = varIamSessionWithoutEmbeddedStruct.ClientIpAddress
		varIamSession.Expiration = varIamSessionWithoutEmbeddedStruct.Expiration
		varIamSession.IdleTimeExpiration = varIamSessionWithoutEmbeddedStruct.IdleTimeExpiration
		varIamSession.LastLoginClient = varIamSessionWithoutEmbeddedStruct.LastLoginClient
		varIamSession.LastLoginTime = varIamSessionWithoutEmbeddedStruct.LastLoginTime
		varIamSession.Permission = varIamSessionWithoutEmbeddedStruct.Permission
		varIamSession.User = varIamSessionWithoutEmbeddedStruct.User
		*o = IamSession(varIamSession)
	} else {
		return err
	}

	varIamSession := _IamSession{}

	err = json.Unmarshal(bytes, &varIamSession)
	if err == nil {
		o.MoBaseMo = varIamSession.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AccountPermissions")
		delete(additionalProperties, "ClientIpAddress")
		delete(additionalProperties, "Expiration")
		delete(additionalProperties, "IdleTimeExpiration")
		delete(additionalProperties, "LastLoginClient")
		delete(additionalProperties, "LastLoginTime")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "User")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamSession struct {
	value *IamSession
	isSet bool
}

func (v NullableIamSession) Get() *IamSession {
	return v.value
}

func (v *NullableIamSession) Set(val *IamSession) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSession) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSession(val *IamSession) *NullableIamSession {
	return &NullableIamSession{value: val, isSet: true}
}

func (v NullableIamSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
