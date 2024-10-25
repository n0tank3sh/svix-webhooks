/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EndpointHeadersOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointHeadersOut{}

// EndpointHeadersOut The value of the headers is returned in the `headers` field.  Sensitive headers that have been redacted are returned in the sensitive field.
type EndpointHeadersOut struct {
	Headers map[string]string `json:"headers"`
	Sensitive []string `json:"sensitive"`
}

type _EndpointHeadersOut EndpointHeadersOut

// NewEndpointHeadersOut instantiates a new EndpointHeadersOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointHeadersOut(headers map[string]string, sensitive []string) *EndpointHeadersOut {
	this := EndpointHeadersOut{}
	this.Headers = headers
	this.Sensitive = sensitive
	return &this
}

// NewEndpointHeadersOutWithDefaults instantiates a new EndpointHeadersOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointHeadersOutWithDefaults() *EndpointHeadersOut {
	this := EndpointHeadersOut{}
	return &this
}

// GetHeaders returns the Headers field value
func (o *EndpointHeadersOut) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *EndpointHeadersOut) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *EndpointHeadersOut) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetSensitive returns the Sensitive field value
func (o *EndpointHeadersOut) GetSensitive() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value
// and a boolean to check if the value has been set.
func (o *EndpointHeadersOut) GetSensitiveOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// SetSensitive sets field value
func (o *EndpointHeadersOut) SetSensitive(v []string) {
	o.Sensitive = v
}

func (o EndpointHeadersOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointHeadersOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headers"] = o.Headers
	toSerialize["sensitive"] = o.Sensitive
	return toSerialize, nil
}

func (o *EndpointHeadersOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"headers",
		"sensitive",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEndpointHeadersOut := _EndpointHeadersOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointHeadersOut)

	if err != nil {
		return err
	}

	*o = EndpointHeadersOut(varEndpointHeadersOut)

	return err
}

type NullableEndpointHeadersOut struct {
	value *EndpointHeadersOut
	isSet bool
}

func (v NullableEndpointHeadersOut) Get() *EndpointHeadersOut {
	return v.value
}

func (v *NullableEndpointHeadersOut) Set(val *EndpointHeadersOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointHeadersOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointHeadersOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointHeadersOut(val *EndpointHeadersOut) *NullableEndpointHeadersOut {
	return &NullableEndpointHeadersOut{value: val, isSet: true}
}

func (v NullableEndpointHeadersOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointHeadersOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


