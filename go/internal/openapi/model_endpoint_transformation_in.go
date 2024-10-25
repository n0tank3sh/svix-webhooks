/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EndpointTransformationIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointTransformationIn{}

// EndpointTransformationIn struct for EndpointTransformationIn
type EndpointTransformationIn struct {
	Code NullableString `json:"code,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewEndpointTransformationIn instantiates a new EndpointTransformationIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTransformationIn() *EndpointTransformationIn {
	this := EndpointTransformationIn{}
	return &this
}

// NewEndpointTransformationInWithDefaults instantiates a new EndpointTransformationIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTransformationInWithDefaults() *EndpointTransformationIn {
	this := EndpointTransformationIn{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointTransformationIn) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointTransformationIn) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *EndpointTransformationIn) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *EndpointTransformationIn) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *EndpointTransformationIn) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *EndpointTransformationIn) UnsetCode() {
	o.Code.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EndpointTransformationIn) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTransformationIn) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EndpointTransformationIn) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EndpointTransformationIn) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o EndpointTransformationIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTransformationIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableEndpointTransformationIn struct {
	value *EndpointTransformationIn
	isSet bool
}

func (v NullableEndpointTransformationIn) Get() *EndpointTransformationIn {
	return v.value
}

func (v *NullableEndpointTransformationIn) Set(val *EndpointTransformationIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTransformationIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTransformationIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTransformationIn(val *EndpointTransformationIn) *NullableEndpointTransformationIn {
	return &NullableEndpointTransformationIn{value: val, isSet: true}
}

func (v NullableEndpointTransformationIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTransformationIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


