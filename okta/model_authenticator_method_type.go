/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// AuthenticatorMethodType the model 'AuthenticatorMethodType'
type AuthenticatorMethodType string

// List of AuthenticatorMethodType
const (
	AUTHENTICATORMETHODTYPE_CERT              AuthenticatorMethodType = "cert"
	AUTHENTICATORMETHODTYPE_DUO               AuthenticatorMethodType = "duo"
	AUTHENTICATORMETHODTYPE_EMAIL             AuthenticatorMethodType = "email"
	AUTHENTICATORMETHODTYPE_IDP               AuthenticatorMethodType = "idp"
	AUTHENTICATORMETHODTYPE_OTP               AuthenticatorMethodType = "otp"
	AUTHENTICATORMETHODTYPE_PASSWORD          AuthenticatorMethodType = "password"
	AUTHENTICATORMETHODTYPE_PUSH              AuthenticatorMethodType = "push"
	AUTHENTICATORMETHODTYPE_SECURITY_QUESTION AuthenticatorMethodType = "security_question"
	AUTHENTICATORMETHODTYPE_SIGNED_NONCE      AuthenticatorMethodType = "signed_nonce"
	AUTHENTICATORMETHODTYPE_SMS               AuthenticatorMethodType = "sms"
	AUTHENTICATORMETHODTYPE_TOTP              AuthenticatorMethodType = "totp"
	AUTHENTICATORMETHODTYPE_VOICE             AuthenticatorMethodType = "voice"
	AUTHENTICATORMETHODTYPE_WEBAUTHN          AuthenticatorMethodType = "webauthn"
)

// All allowed values of AuthenticatorMethodType enum
var AllowedAuthenticatorMethodTypeEnumValues = []AuthenticatorMethodType{
	"cert",
	"duo",
	"email",
	"idp",
	"otp",
	"password",
	"push",
	"security_question",
	"signed_nonce",
	"sms",
	"totp",
	"voice",
	"webauthn",
}

func (v *AuthenticatorMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticatorMethodType(value)
	for _, existing := range AllowedAuthenticatorMethodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticatorMethodType", value)
}

// NewAuthenticatorMethodTypeFromValue returns a pointer to a valid AuthenticatorMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticatorMethodTypeFromValue(v string) (*AuthenticatorMethodType, error) {
	ev := AuthenticatorMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticatorMethodType: valid values are %v", v, AllowedAuthenticatorMethodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticatorMethodType) IsValid() bool {
	for _, existing := range AllowedAuthenticatorMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthenticatorMethodType value
func (v AuthenticatorMethodType) Ptr() *AuthenticatorMethodType {
	return &v
}

type NullableAuthenticatorMethodType struct {
	value *AuthenticatorMethodType
	isSet bool
}

func (v NullableAuthenticatorMethodType) Get() *AuthenticatorMethodType {
	return v.value
}

func (v *NullableAuthenticatorMethodType) Set(val *AuthenticatorMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodType(val *AuthenticatorMethodType) *NullableAuthenticatorMethodType {
	return &NullableAuthenticatorMethodType{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}