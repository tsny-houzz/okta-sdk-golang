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
)

// U2fUserFactorProfile struct for U2fUserFactorProfile
type U2fUserFactorProfile struct {
	CredentialId         *string `json:"credentialId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _U2fUserFactorProfile U2fUserFactorProfile

// NewU2fUserFactorProfile instantiates a new U2fUserFactorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewU2fUserFactorProfile() *U2fUserFactorProfile {
	this := U2fUserFactorProfile{}
	return &this
}

// NewU2fUserFactorProfileWithDefaults instantiates a new U2fUserFactorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewU2fUserFactorProfileWithDefaults() *U2fUserFactorProfile {
	this := U2fUserFactorProfile{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *U2fUserFactorProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *U2fUserFactorProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *U2fUserFactorProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *U2fUserFactorProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o U2fUserFactorProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *U2fUserFactorProfile) UnmarshalJSON(bytes []byte) (err error) {
	varU2fUserFactorProfile := _U2fUserFactorProfile{}

	err = json.Unmarshal(bytes, &varU2fUserFactorProfile)
	if err == nil {
		*o = U2fUserFactorProfile(varU2fUserFactorProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableU2fUserFactorProfile struct {
	value *U2fUserFactorProfile
	isSet bool
}

func (v NullableU2fUserFactorProfile) Get() *U2fUserFactorProfile {
	return v.value
}

func (v *NullableU2fUserFactorProfile) Set(val *U2fUserFactorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableU2fUserFactorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableU2fUserFactorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableU2fUserFactorProfile(val *U2fUserFactorProfile) *NullableU2fUserFactorProfile {
	return &NullableU2fUserFactorProfile{value: val, isSet: true}
}

func (v NullableU2fUserFactorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableU2fUserFactorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
