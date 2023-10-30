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

// RiskProviderAction Action taken by Okta during authentication attempts based on the risk events sent by this provider
type RiskProviderAction string

// List of RiskProviderAction
const (
	RISKPROVIDERACTION_ENFORCE_AND_LOG RiskProviderAction = "enforce_and_log"
	RISKPROVIDERACTION_LOG_ONLY        RiskProviderAction = "log_only"
	RISKPROVIDERACTION_NONE            RiskProviderAction = "none"
)

// All allowed values of RiskProviderAction enum
var AllowedRiskProviderActionEnumValues = []RiskProviderAction{
	"enforce_and_log",
	"log_only",
	"none",
}

func (v *RiskProviderAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RiskProviderAction(value)
	for _, existing := range AllowedRiskProviderActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RiskProviderAction", value)
}

// NewRiskProviderActionFromValue returns a pointer to a valid RiskProviderAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskProviderActionFromValue(v string) (*RiskProviderAction, error) {
	ev := RiskProviderAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RiskProviderAction: valid values are %v", v, AllowedRiskProviderActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskProviderAction) IsValid() bool {
	for _, existing := range AllowedRiskProviderActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskProviderAction value
func (v RiskProviderAction) Ptr() *RiskProviderAction {
	return &v
}

type NullableRiskProviderAction struct {
	value *RiskProviderAction
	isSet bool
}

func (v NullableRiskProviderAction) Get() *RiskProviderAction {
	return v.value
}

func (v *NullableRiskProviderAction) Set(val *RiskProviderAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskProviderAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskProviderAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskProviderAction(val *RiskProviderAction) *NullableRiskProviderAction {
	return &NullableRiskProviderAction{value: val, isSet: true}
}

func (v NullableRiskProviderAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskProviderAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
