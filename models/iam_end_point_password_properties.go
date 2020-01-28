// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IamEndPointPasswordProperties Iam:End Point Password Properties
//
// Password properties for endpoint users.
//
// swagger:model iamEndPointPasswordProperties
type IamEndPointPasswordProperties struct {
	MoBaseComplexType

	IamEndPointPasswordPropertiesAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IamEndPointPasswordProperties) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 IamEndPointPasswordPropertiesAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IamEndPointPasswordPropertiesAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IamEndPointPasswordProperties) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IamEndPointPasswordPropertiesAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this iam end point password properties
func (m *IamEndPointPasswordProperties) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IamEndPointPasswordPropertiesAO1P1
	if err := m.IamEndPointPasswordPropertiesAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IamEndPointPasswordProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamEndPointPasswordProperties) UnmarshalBinary(b []byte) error {
	var res IamEndPointPasswordProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IamEndPointPasswordPropertiesAO1P1 iam end point password properties a o1 p1
// swagger:model IamEndPointPasswordPropertiesAO1P1
type IamEndPointPasswordPropertiesAO1P1 struct {

	// Enables password expiry on the endpoint.
	//
	EnablePasswordExpiry *bool `json:"EnablePasswordExpiry,omitempty"`

	// Enables a strong password policy Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User's Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (!, @, #, $, %, ^, &, *, -, _, +, =).
	//
	//
	EnforceStrongPassword *bool `json:"EnforceStrongPassword,omitempty"`

	// Time period until when you can use the existing password, after it expires.
	//
	GracePeriod int64 `json:"GracePeriod,omitempty"`

	// The duration by when the password will expire.
	//
	NotificationPeriod int64 `json:"NotificationPeriod,omitempty"`

	// Set time period for password expiration. Value should be greater than notification period and grace period.
	//
	PasswordExpiryDuration int64 `json:"PasswordExpiryDuration,omitempty"`

	// Tracks password change history. Specifies in number of instances, that the new password was already used.
	//
	PasswordHistory int64 `json:"PasswordHistory,omitempty"`

	// iam end point password properties a o1 p1
	IamEndPointPasswordPropertiesAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IamEndPointPasswordPropertiesAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Enables password expiry on the endpoint.
		//
		EnablePasswordExpiry *bool `json:"EnablePasswordExpiry,omitempty"`

		// Enables a strong password policy Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User's Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (!, @, #, $, %, ^, &, *, -, _, +, =).
		//
		//
		EnforceStrongPassword *bool `json:"EnforceStrongPassword,omitempty"`

		// Time period until when you can use the existing password, after it expires.
		//
		GracePeriod int64 `json:"GracePeriod,omitempty"`

		// The duration by when the password will expire.
		//
		NotificationPeriod int64 `json:"NotificationPeriod,omitempty"`

		// Set time period for password expiration. Value should be greater than notification period and grace period.
		//
		PasswordExpiryDuration int64 `json:"PasswordExpiryDuration,omitempty"`

		// Tracks password change history. Specifies in number of instances, that the new password was already used.
		//
		PasswordHistory int64 `json:"PasswordHistory,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IamEndPointPasswordPropertiesAO1P1

	rcv.EnablePasswordExpiry = stage1.EnablePasswordExpiry

	rcv.EnforceStrongPassword = stage1.EnforceStrongPassword

	rcv.GracePeriod = stage1.GracePeriod

	rcv.NotificationPeriod = stage1.NotificationPeriod

	rcv.PasswordExpiryDuration = stage1.PasswordExpiryDuration

	rcv.PasswordHistory = stage1.PasswordHistory

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "EnablePasswordExpiry")

	delete(stage2, "EnforceStrongPassword")

	delete(stage2, "GracePeriod")

	delete(stage2, "NotificationPeriod")

	delete(stage2, "PasswordExpiryDuration")

	delete(stage2, "PasswordHistory")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.IamEndPointPasswordPropertiesAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IamEndPointPasswordPropertiesAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Enables password expiry on the endpoint.
		//
		EnablePasswordExpiry *bool `json:"EnablePasswordExpiry,omitempty"`

		// Enables a strong password policy Strong password requirements: A. The password must have a minimum of 8 and a maximum of 20 characters. B. The password must not contain the User's Name. C. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (!, @, #, $, %, ^, &, *, -, _, +, =).
		//
		//
		EnforceStrongPassword *bool `json:"EnforceStrongPassword,omitempty"`

		// Time period until when you can use the existing password, after it expires.
		//
		GracePeriod int64 `json:"GracePeriod,omitempty"`

		// The duration by when the password will expire.
		//
		NotificationPeriod int64 `json:"NotificationPeriod,omitempty"`

		// Set time period for password expiration. Value should be greater than notification period and grace period.
		//
		PasswordExpiryDuration int64 `json:"PasswordExpiryDuration,omitempty"`

		// Tracks password change history. Specifies in number of instances, that the new password was already used.
		//
		PasswordHistory int64 `json:"PasswordHistory,omitempty"`
	}

	stage1.EnablePasswordExpiry = m.EnablePasswordExpiry

	stage1.EnforceStrongPassword = m.EnforceStrongPassword

	stage1.GracePeriod = m.GracePeriod

	stage1.NotificationPeriod = m.NotificationPeriod

	stage1.PasswordExpiryDuration = m.PasswordExpiryDuration

	stage1.PasswordHistory = m.PasswordHistory

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IamEndPointPasswordPropertiesAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IamEndPointPasswordPropertiesAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this iam end point password properties a o1 p1
func (m *IamEndPointPasswordPropertiesAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamEndPointPasswordPropertiesAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamEndPointPasswordPropertiesAO1P1) UnmarshalBinary(b []byte) error {
	var res IamEndPointPasswordPropertiesAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
