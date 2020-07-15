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
)

// IamCertificateRequest The information required to generate a certificate signing request (CSR), which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.
type IamCertificateRequest struct {
	MoBaseMo
	// User input email address, an optional part of the subject of the certificate request.
	EmailAddress *string `json:"EmailAddress,omitempty"`
	// Name of the certificate request.
	Name *string `json:"Name,omitempty"`
	// Generated certificate signing request (CSR) in PEM format.
	Request *string `json:"Request,omitempty"`
	// Whether the user wants the generated CSR to be self-signed by the appliance.
	SelfSigned           *bool                          `json:"SelfSigned,omitempty"`
	Subject              *PkixDistinguishedName         `json:"Subject,omitempty"`
	SubjectAlternateName *PkixSubjectAlternateName      `json:"SubjectAlternateName,omitempty"`
	Account              *IamAccountRelationship        `json:"Account,omitempty"`
	Certificate          *IamCertificateRelationship    `json:"Certificate,omitempty"`
	PrivateKeySpec       *IamPrivateKeySpecRelationship `json:"PrivateKeySpec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamCertificateRequest IamCertificateRequest

// NewIamCertificateRequest instantiates a new IamCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamCertificateRequest() *IamCertificateRequest {
	this := IamCertificateRequest{}
	return &this
}

// NewIamCertificateRequestWithDefaults instantiates a new IamCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamCertificateRequestWithDefaults() *IamCertificateRequest {
	this := IamCertificateRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *IamCertificateRequest) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamCertificateRequest) SetName(v string) {
	o.Name = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *IamCertificateRequest) SetRequest(v string) {
	o.Request = &v
}

// GetSelfSigned returns the SelfSigned field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetSelfSigned() bool {
	if o == nil || o.SelfSigned == nil {
		var ret bool
		return ret
	}
	return *o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetSelfSignedOk() (*bool, bool) {
	if o == nil || o.SelfSigned == nil {
		return nil, false
	}
	return o.SelfSigned, true
}

// HasSelfSigned returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasSelfSigned() bool {
	if o != nil && o.SelfSigned != nil {
		return true
	}

	return false
}

// SetSelfSigned gets a reference to the given bool and assigns it to the SelfSigned field.
func (o *IamCertificateRequest) SetSelfSigned(v bool) {
	o.SelfSigned = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetSubject() PkixDistinguishedName {
	if o == nil || o.Subject == nil {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetSubjectOk() (*PkixDistinguishedName, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PkixDistinguishedName and assigns it to the Subject field.
func (o *IamCertificateRequest) SetSubject(v PkixDistinguishedName) {
	o.Subject = &v
}

// GetSubjectAlternateName returns the SubjectAlternateName field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetSubjectAlternateName() PkixSubjectAlternateName {
	if o == nil || o.SubjectAlternateName == nil {
		var ret PkixSubjectAlternateName
		return ret
	}
	return *o.SubjectAlternateName
}

// GetSubjectAlternateNameOk returns a tuple with the SubjectAlternateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetSubjectAlternateNameOk() (*PkixSubjectAlternateName, bool) {
	if o == nil || o.SubjectAlternateName == nil {
		return nil, false
	}
	return o.SubjectAlternateName, true
}

// HasSubjectAlternateName returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasSubjectAlternateName() bool {
	if o != nil && o.SubjectAlternateName != nil {
		return true
	}

	return false
}

// SetSubjectAlternateName gets a reference to the given PkixSubjectAlternateName and assigns it to the SubjectAlternateName field.
func (o *IamCertificateRequest) SetSubjectAlternateName(v PkixSubjectAlternateName) {
	o.SubjectAlternateName = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamCertificateRequest) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetCertificate() IamCertificateRelationship {
	if o == nil || o.Certificate == nil {
		var ret IamCertificateRelationship
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetCertificateOk() (*IamCertificateRelationship, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given IamCertificateRelationship and assigns it to the Certificate field.
func (o *IamCertificateRequest) SetCertificate(v IamCertificateRelationship) {
	o.Certificate = &v
}

// GetPrivateKeySpec returns the PrivateKeySpec field value if set, zero value otherwise.
func (o *IamCertificateRequest) GetPrivateKeySpec() IamPrivateKeySpecRelationship {
	if o == nil || o.PrivateKeySpec == nil {
		var ret IamPrivateKeySpecRelationship
		return ret
	}
	return *o.PrivateKeySpec
}

// GetPrivateKeySpecOk returns a tuple with the PrivateKeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequest) GetPrivateKeySpecOk() (*IamPrivateKeySpecRelationship, bool) {
	if o == nil || o.PrivateKeySpec == nil {
		return nil, false
	}
	return o.PrivateKeySpec, true
}

// HasPrivateKeySpec returns a boolean if a field has been set.
func (o *IamCertificateRequest) HasPrivateKeySpec() bool {
	if o != nil && o.PrivateKeySpec != nil {
		return true
	}

	return false
}

// SetPrivateKeySpec gets a reference to the given IamPrivateKeySpecRelationship and assigns it to the PrivateKeySpec field.
func (o *IamCertificateRequest) SetPrivateKeySpec(v IamPrivateKeySpecRelationship) {
	o.PrivateKeySpec = &v
}

func (o IamCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.EmailAddress != nil {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}
	if o.SelfSigned != nil {
		toSerialize["SelfSigned"] = o.SelfSigned
	}
	if o.Subject != nil {
		toSerialize["Subject"] = o.Subject
	}
	if o.SubjectAlternateName != nil {
		toSerialize["SubjectAlternateName"] = o.SubjectAlternateName
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}
	if o.PrivateKeySpec != nil {
		toSerialize["PrivateKeySpec"] = o.PrivateKeySpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamCertificateRequest) UnmarshalJSON(bytes []byte) (err error) {
	type IamCertificateRequestWithoutEmbeddedStruct struct {
		// User input email address, an optional part of the subject of the certificate request.
		EmailAddress *string `json:"EmailAddress,omitempty"`
		// Name of the certificate request.
		Name *string `json:"Name,omitempty"`
		// Generated certificate signing request (CSR) in PEM format.
		Request *string `json:"Request,omitempty"`
		// Whether the user wants the generated CSR to be self-signed by the appliance.
		SelfSigned           *bool                          `json:"SelfSigned,omitempty"`
		Subject              *PkixDistinguishedName         `json:"Subject,omitempty"`
		SubjectAlternateName *PkixSubjectAlternateName      `json:"SubjectAlternateName,omitempty"`
		Account              *IamAccountRelationship        `json:"Account,omitempty"`
		Certificate          *IamCertificateRelationship    `json:"Certificate,omitempty"`
		PrivateKeySpec       *IamPrivateKeySpecRelationship `json:"PrivateKeySpec,omitempty"`
	}

	varIamCertificateRequestWithoutEmbeddedStruct := IamCertificateRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamCertificateRequestWithoutEmbeddedStruct)
	if err == nil {
		varIamCertificateRequest := _IamCertificateRequest{}
		varIamCertificateRequest.EmailAddress = varIamCertificateRequestWithoutEmbeddedStruct.EmailAddress
		varIamCertificateRequest.Name = varIamCertificateRequestWithoutEmbeddedStruct.Name
		varIamCertificateRequest.Request = varIamCertificateRequestWithoutEmbeddedStruct.Request
		varIamCertificateRequest.SelfSigned = varIamCertificateRequestWithoutEmbeddedStruct.SelfSigned
		varIamCertificateRequest.Subject = varIamCertificateRequestWithoutEmbeddedStruct.Subject
		varIamCertificateRequest.SubjectAlternateName = varIamCertificateRequestWithoutEmbeddedStruct.SubjectAlternateName
		varIamCertificateRequest.Account = varIamCertificateRequestWithoutEmbeddedStruct.Account
		varIamCertificateRequest.Certificate = varIamCertificateRequestWithoutEmbeddedStruct.Certificate
		varIamCertificateRequest.PrivateKeySpec = varIamCertificateRequestWithoutEmbeddedStruct.PrivateKeySpec
		*o = IamCertificateRequest(varIamCertificateRequest)
	} else {
		return err
	}

	varIamCertificateRequest := _IamCertificateRequest{}

	err = json.Unmarshal(bytes, &varIamCertificateRequest)
	if err == nil {
		o.MoBaseMo = varIamCertificateRequest.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "EmailAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Request")
		delete(additionalProperties, "SelfSigned")
		delete(additionalProperties, "Subject")
		delete(additionalProperties, "SubjectAlternateName")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Certificate")
		delete(additionalProperties, "PrivateKeySpec")

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

type NullableIamCertificateRequest struct {
	value *IamCertificateRequest
	isSet bool
}

func (v NullableIamCertificateRequest) Get() *IamCertificateRequest {
	return v.value
}

func (v *NullableIamCertificateRequest) Set(val *IamCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIamCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIamCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamCertificateRequest(val *IamCertificateRequest) *NullableIamCertificateRequest {
	return &NullableIamCertificateRequest{value: val, isSet: true}
}

func (v NullableIamCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
