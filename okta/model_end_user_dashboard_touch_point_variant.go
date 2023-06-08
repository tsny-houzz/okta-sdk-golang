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

// EndUserDashboardTouchPointVariant the model 'EndUserDashboardTouchPointVariant'
type EndUserDashboardTouchPointVariant string

// List of EndUserDashboardTouchPointVariant
const (
	ENDUSERDASHBOARDTOUCHPOINTVARIANT_FULL_THEME                    EndUserDashboardTouchPointVariant = "FULL_THEME"
	ENDUSERDASHBOARDTOUCHPOINTVARIANT_LOGO_ON_FULL_WHITE_BACKGROUND EndUserDashboardTouchPointVariant = "LOGO_ON_FULL_WHITE_BACKGROUND"
	ENDUSERDASHBOARDTOUCHPOINTVARIANT_OKTA_DEFAULT                  EndUserDashboardTouchPointVariant = "OKTA_DEFAULT"
	ENDUSERDASHBOARDTOUCHPOINTVARIANT_WHITE_LOGO_BACKGROUND         EndUserDashboardTouchPointVariant = "WHITE_LOGO_BACKGROUND"
)

// All allowed values of EndUserDashboardTouchPointVariant enum
var AllowedEndUserDashboardTouchPointVariantEnumValues = []EndUserDashboardTouchPointVariant{
	"FULL_THEME",
	"LOGO_ON_FULL_WHITE_BACKGROUND",
	"OKTA_DEFAULT",
	"WHITE_LOGO_BACKGROUND",
}

func (v *EndUserDashboardTouchPointVariant) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EndUserDashboardTouchPointVariant(value)
	for _, existing := range AllowedEndUserDashboardTouchPointVariantEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EndUserDashboardTouchPointVariant", value)
}

// NewEndUserDashboardTouchPointVariantFromValue returns a pointer to a valid EndUserDashboardTouchPointVariant
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndUserDashboardTouchPointVariantFromValue(v string) (*EndUserDashboardTouchPointVariant, error) {
	ev := EndUserDashboardTouchPointVariant(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EndUserDashboardTouchPointVariant: valid values are %v", v, AllowedEndUserDashboardTouchPointVariantEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EndUserDashboardTouchPointVariant) IsValid() bool {
	for _, existing := range AllowedEndUserDashboardTouchPointVariantEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndUserDashboardTouchPointVariant value
func (v EndUserDashboardTouchPointVariant) Ptr() *EndUserDashboardTouchPointVariant {
	return &v
}

type NullableEndUserDashboardTouchPointVariant struct {
	value *EndUserDashboardTouchPointVariant
	isSet bool
}

func (v NullableEndUserDashboardTouchPointVariant) Get() *EndUserDashboardTouchPointVariant {
	return v.value
}

func (v *NullableEndUserDashboardTouchPointVariant) Set(val *EndUserDashboardTouchPointVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserDashboardTouchPointVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserDashboardTouchPointVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserDashboardTouchPointVariant(val *EndUserDashboardTouchPointVariant) *NullableEndUserDashboardTouchPointVariant {
	return &NullableEndUserDashboardTouchPointVariant{value: val, isSet: true}
}

func (v NullableEndUserDashboardTouchPointVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserDashboardTouchPointVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}