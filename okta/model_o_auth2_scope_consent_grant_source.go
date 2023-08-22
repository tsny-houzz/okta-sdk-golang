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

// OAuth2ScopeConsentGrantSource User type source that granted consent
type OAuth2ScopeConsentGrantSource string

// List of OAuth2ScopeConsentGrantSource
const (
	OAUTH2SCOPECONSENTGRANTSOURCE_ADMIN    OAuth2ScopeConsentGrantSource = "ADMIN"
	OAUTH2SCOPECONSENTGRANTSOURCE_END_USER OAuth2ScopeConsentGrantSource = "END_USER"
)

// All allowed values of OAuth2ScopeConsentGrantSource enum
var AllowedOAuth2ScopeConsentGrantSourceEnumValues = []OAuth2ScopeConsentGrantSource{
	"ADMIN",
	"END_USER",
}

func (v *OAuth2ScopeConsentGrantSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OAuth2ScopeConsentGrantSource(value)
	for _, existing := range AllowedOAuth2ScopeConsentGrantSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OAuth2ScopeConsentGrantSource", value)
}

// NewOAuth2ScopeConsentGrantSourceFromValue returns a pointer to a valid OAuth2ScopeConsentGrantSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOAuth2ScopeConsentGrantSourceFromValue(v string) (*OAuth2ScopeConsentGrantSource, error) {
	ev := OAuth2ScopeConsentGrantSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OAuth2ScopeConsentGrantSource: valid values are %v", v, AllowedOAuth2ScopeConsentGrantSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OAuth2ScopeConsentGrantSource) IsValid() bool {
	for _, existing := range AllowedOAuth2ScopeConsentGrantSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuth2ScopeConsentGrantSource value
func (v OAuth2ScopeConsentGrantSource) Ptr() *OAuth2ScopeConsentGrantSource {
	return &v
}

type NullableOAuth2ScopeConsentGrantSource struct {
	value *OAuth2ScopeConsentGrantSource
	isSet bool
}

func (v NullableOAuth2ScopeConsentGrantSource) Get() *OAuth2ScopeConsentGrantSource {
	return v.value
}

func (v *NullableOAuth2ScopeConsentGrantSource) Set(val *OAuth2ScopeConsentGrantSource) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopeConsentGrantSource) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopeConsentGrantSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopeConsentGrantSource(val *OAuth2ScopeConsentGrantSource) *NullableOAuth2ScopeConsentGrantSource {
	return &NullableOAuth2ScopeConsentGrantSource{value: val, isSet: true}
}

func (v NullableOAuth2ScopeConsentGrantSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopeConsentGrantSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}