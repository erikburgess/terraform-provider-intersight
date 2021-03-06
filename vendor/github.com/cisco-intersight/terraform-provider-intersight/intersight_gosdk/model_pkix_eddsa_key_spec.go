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

// PkixEddsaKeySpec The key pair is generated using Edwards-Curve Digital Signature Algorithm (EdDSA). The Edwards-curve Digital Signature Algorithm (EdDSA) is a variant of Schnorr's signature system with (possibly twisted) Edwards curves.
type PkixEddsaKeySpec struct {
	PkixKeyGenerationSpec
	// The EdDSA algorithm, as defined in RFC 8032.
	Algorithm            *string `json:"Algorithm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixEddsaKeySpec PkixEddsaKeySpec

// NewPkixEddsaKeySpec instantiates a new PkixEddsaKeySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixEddsaKeySpec() *PkixEddsaKeySpec {
	this := PkixEddsaKeySpec{}
	var algorithm string = "Ed25519"
	this.Algorithm = &algorithm
	return &this
}

// NewPkixEddsaKeySpecWithDefaults instantiates a new PkixEddsaKeySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixEddsaKeySpecWithDefaults() *PkixEddsaKeySpec {
	this := PkixEddsaKeySpec{}
	var algorithm string = "Ed25519"
	this.Algorithm = &algorithm
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *PkixEddsaKeySpec) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixEddsaKeySpec) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *PkixEddsaKeySpec) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *PkixEddsaKeySpec) SetAlgorithm(v string) {
	o.Algorithm = &v
}

func (o PkixEddsaKeySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPkixKeyGenerationSpec, errPkixKeyGenerationSpec := json.Marshal(o.PkixKeyGenerationSpec)
	if errPkixKeyGenerationSpec != nil {
		return []byte{}, errPkixKeyGenerationSpec
	}
	errPkixKeyGenerationSpec = json.Unmarshal([]byte(serializedPkixKeyGenerationSpec), &toSerialize)
	if errPkixKeyGenerationSpec != nil {
		return []byte{}, errPkixKeyGenerationSpec
	}
	if o.Algorithm != nil {
		toSerialize["Algorithm"] = o.Algorithm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixEddsaKeySpec) UnmarshalJSON(bytes []byte) (err error) {
	type PkixEddsaKeySpecWithoutEmbeddedStruct struct {
		// The EdDSA algorithm, as defined in RFC 8032.
		Algorithm *string `json:"Algorithm,omitempty"`
	}

	varPkixEddsaKeySpecWithoutEmbeddedStruct := PkixEddsaKeySpecWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPkixEddsaKeySpecWithoutEmbeddedStruct)
	if err == nil {
		varPkixEddsaKeySpec := _PkixEddsaKeySpec{}
		varPkixEddsaKeySpec.Algorithm = varPkixEddsaKeySpecWithoutEmbeddedStruct.Algorithm
		*o = PkixEddsaKeySpec(varPkixEddsaKeySpec)
	} else {
		return err
	}

	varPkixEddsaKeySpec := _PkixEddsaKeySpec{}

	err = json.Unmarshal(bytes, &varPkixEddsaKeySpec)
	if err == nil {
		o.PkixKeyGenerationSpec = varPkixEddsaKeySpec.PkixKeyGenerationSpec
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Algorithm")

		// remove fields from embedded structs
		reflectPkixKeyGenerationSpec := reflect.ValueOf(o.PkixKeyGenerationSpec)
		for i := 0; i < reflectPkixKeyGenerationSpec.Type().NumField(); i++ {
			t := reflectPkixKeyGenerationSpec.Type().Field(i)

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

type NullablePkixEddsaKeySpec struct {
	value *PkixEddsaKeySpec
	isSet bool
}

func (v NullablePkixEddsaKeySpec) Get() *PkixEddsaKeySpec {
	return v.value
}

func (v *NullablePkixEddsaKeySpec) Set(val *PkixEddsaKeySpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixEddsaKeySpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixEddsaKeySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixEddsaKeySpec(val *PkixEddsaKeySpec) *NullablePkixEddsaKeySpec {
	return &NullablePkixEddsaKeySpec{value: val, isSet: true}
}

func (v NullablePkixEddsaKeySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixEddsaKeySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
