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
)

// BrandLinks struct for BrandLinks
type BrandLinks struct {
	Self                 *HrefObject `json:"self,omitempty"`
	Themes               *HrefObject `json:"themes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrandLinks BrandLinks

// NewBrandLinks instantiates a new BrandLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandLinks() *BrandLinks {
	this := BrandLinks{}
	return &this
}

// NewBrandLinksWithDefaults instantiates a new BrandLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandLinksWithDefaults() *BrandLinks {
	this := BrandLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *BrandLinks) GetSelf() HrefObject {
	if o == nil || o.Self == nil {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *BrandLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *BrandLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *BrandLinks) GetThemes() HrefObject {
	if o == nil || o.Themes == nil {
		var ret HrefObject
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLinks) GetThemesOk() (*HrefObject, bool) {
	if o == nil || o.Themes == nil {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *BrandLinks) HasThemes() bool {
	if o != nil && o.Themes != nil {
		return true
	}

	return false
}

// SetThemes gets a reference to the given HrefObject and assigns it to the Themes field.
func (o *BrandLinks) SetThemes(v HrefObject) {
	o.Themes = &v
}

func (o BrandLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Themes != nil {
		toSerialize["themes"] = o.Themes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BrandLinks) UnmarshalJSON(bytes []byte) (err error) {
	varBrandLinks := _BrandLinks{}

	err = json.Unmarshal(bytes, &varBrandLinks)
	if err == nil {
		*o = BrandLinks(varBrandLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "themes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBrandLinks struct {
	value *BrandLinks
	isSet bool
}

func (v NullableBrandLinks) Get() *BrandLinks {
	return v.value
}

func (v *NullableBrandLinks) Set(val *BrandLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandLinks(val *BrandLinks) *NullableBrandLinks {
	return &NullableBrandLinks{value: val, isSet: true}
}

func (v NullableBrandLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}