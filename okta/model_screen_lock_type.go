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

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// ScreenLockType the model 'ScreenLockType'
type ScreenLockType string

// List of ScreenLockType
const (
	SCREENLOCKTYPE_BIOMETRIC ScreenLockType = "BIOMETRIC"
	SCREENLOCKTYPE_PASSCODE  ScreenLockType = "PASSCODE"
)

// All allowed values of ScreenLockType enum
var AllowedScreenLockTypeEnumValues = []ScreenLockType{
	"BIOMETRIC",
	"PASSCODE",
}

func (v *ScreenLockType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScreenLockType(value)
	for _, existing := range AllowedScreenLockTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScreenLockType", value)
}

// NewScreenLockTypeFromValue returns a pointer to a valid ScreenLockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScreenLockTypeFromValue(v string) (*ScreenLockType, error) {
	ev := ScreenLockType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScreenLockType: valid values are %v", v, AllowedScreenLockTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScreenLockType) IsValid() bool {
	for _, existing := range AllowedScreenLockTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScreenLockType value
func (v ScreenLockType) Ptr() *ScreenLockType {
	return &v
}

type NullableScreenLockType struct {
	value *ScreenLockType
	isSet bool
}

func (v NullableScreenLockType) Get() *ScreenLockType {
	return v.value
}

func (v *NullableScreenLockType) Set(val *ScreenLockType) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenLockType) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenLockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenLockType(val *ScreenLockType) *NullableScreenLockType {
	return &NullableScreenLockType{value: val, isSet: true}
}

func (v NullableScreenLockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenLockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
