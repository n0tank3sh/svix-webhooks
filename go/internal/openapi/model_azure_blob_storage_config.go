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

// checks if the AzureBlobStorageConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureBlobStorageConfig{}

// AzureBlobStorageConfig struct for AzureBlobStorageConfig
type AzureBlobStorageConfig struct {
	AccessKey string `json:"accessKey"`
	Account string `json:"account"`
	Container string `json:"container"`
}

type _AzureBlobStorageConfig AzureBlobStorageConfig

// NewAzureBlobStorageConfig instantiates a new AzureBlobStorageConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageConfig(accessKey string, account string, container string) *AzureBlobStorageConfig {
	this := AzureBlobStorageConfig{}
	this.AccessKey = accessKey
	this.Account = account
	this.Container = container
	return &this
}

// NewAzureBlobStorageConfigWithDefaults instantiates a new AzureBlobStorageConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageConfigWithDefaults() *AzureBlobStorageConfig {
	this := AzureBlobStorageConfig{}
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *AzureBlobStorageConfig) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageConfig) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *AzureBlobStorageConfig) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetAccount returns the Account field value
func (o *AzureBlobStorageConfig) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageConfig) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureBlobStorageConfig) SetAccount(v string) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureBlobStorageConfig) GetContainer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageConfig) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureBlobStorageConfig) SetContainer(v string) {
	o.Container = v
}

func (o AzureBlobStorageConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureBlobStorageConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessKey"] = o.AccessKey
	toSerialize["account"] = o.Account
	toSerialize["container"] = o.Container
	return toSerialize, nil
}

func (o *AzureBlobStorageConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessKey",
		"account",
		"container",
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

	varAzureBlobStorageConfig := _AzureBlobStorageConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureBlobStorageConfig)

	if err != nil {
		return err
	}

	*o = AzureBlobStorageConfig(varAzureBlobStorageConfig)

	return err
}

type NullableAzureBlobStorageConfig struct {
	value *AzureBlobStorageConfig
	isSet bool
}

func (v NullableAzureBlobStorageConfig) Get() *AzureBlobStorageConfig {
	return v.value
}

func (v *NullableAzureBlobStorageConfig) Set(val *AzureBlobStorageConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageConfig(val *AzureBlobStorageConfig) *NullableAzureBlobStorageConfig {
	return &NullableAzureBlobStorageConfig{value: val, isSet: true}
}

func (v NullableAzureBlobStorageConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

