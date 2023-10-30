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

// MultifactorEnrollmentPolicyAuthenticatorStatus the model 'MultifactorEnrollmentPolicyAuthenticatorStatus'
type MultifactorEnrollmentPolicyAuthenticatorStatus string

// List of MultifactorEnrollmentPolicyAuthenticatorStatus
const (
	MULTIFACTORENROLLMENTPOLICYAUTHENTICATORSTATUS_NOT_ALLOWED MultifactorEnrollmentPolicyAuthenticatorStatus = "NOT_ALLOWED"
	MULTIFACTORENROLLMENTPOLICYAUTHENTICATORSTATUS_OPTIONAL    MultifactorEnrollmentPolicyAuthenticatorStatus = "OPTIONAL"
	MULTIFACTORENROLLMENTPOLICYAUTHENTICATORSTATUS_REQUIRED    MultifactorEnrollmentPolicyAuthenticatorStatus = "REQUIRED"
)

// All allowed values of MultifactorEnrollmentPolicyAuthenticatorStatus enum
var AllowedMultifactorEnrollmentPolicyAuthenticatorStatusEnumValues = []MultifactorEnrollmentPolicyAuthenticatorStatus{
	"NOT_ALLOWED",
	"OPTIONAL",
	"REQUIRED",
}

func (v *MultifactorEnrollmentPolicyAuthenticatorStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MultifactorEnrollmentPolicyAuthenticatorStatus(value)
	for _, existing := range AllowedMultifactorEnrollmentPolicyAuthenticatorStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MultifactorEnrollmentPolicyAuthenticatorStatus", value)
}

// NewMultifactorEnrollmentPolicyAuthenticatorStatusFromValue returns a pointer to a valid MultifactorEnrollmentPolicyAuthenticatorStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMultifactorEnrollmentPolicyAuthenticatorStatusFromValue(v string) (*MultifactorEnrollmentPolicyAuthenticatorStatus, error) {
	ev := MultifactorEnrollmentPolicyAuthenticatorStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MultifactorEnrollmentPolicyAuthenticatorStatus: valid values are %v", v, AllowedMultifactorEnrollmentPolicyAuthenticatorStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultifactorEnrollmentPolicyAuthenticatorStatus) IsValid() bool {
	for _, existing := range AllowedMultifactorEnrollmentPolicyAuthenticatorStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MultifactorEnrollmentPolicyAuthenticatorStatus value
func (v MultifactorEnrollmentPolicyAuthenticatorStatus) Ptr() *MultifactorEnrollmentPolicyAuthenticatorStatus {
	return &v
}

type NullableMultifactorEnrollmentPolicyAuthenticatorStatus struct {
	value *MultifactorEnrollmentPolicyAuthenticatorStatus
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorStatus) Get() *MultifactorEnrollmentPolicyAuthenticatorStatus {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorStatus) Set(val *MultifactorEnrollmentPolicyAuthenticatorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicyAuthenticatorStatus(val *MultifactorEnrollmentPolicyAuthenticatorStatus) *NullableMultifactorEnrollmentPolicyAuthenticatorStatus {
	return &NullableMultifactorEnrollmentPolicyAuthenticatorStatus{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
