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

// checks if the EventExampleIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventExampleIn{}

// EventExampleIn struct for EventExampleIn
type EventExampleIn struct {
	// The event type's name
	EventType string `json:"eventType" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
}

type _EventExampleIn EventExampleIn

// NewEventExampleIn instantiates a new EventExampleIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventExampleIn(eventType string) *EventExampleIn {
	this := EventExampleIn{}
	this.EventType = eventType
	return &this
}

// NewEventExampleInWithDefaults instantiates a new EventExampleIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventExampleInWithDefaults() *EventExampleIn {
	this := EventExampleIn{}
	return &this
}

// GetEventType returns the EventType field value
func (o *EventExampleIn) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EventExampleIn) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EventExampleIn) SetEventType(v string) {
	o.EventType = v
}

func (o EventExampleIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventExampleIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	return toSerialize, nil
}

func (o *EventExampleIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventType",
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

	varEventExampleIn := _EventExampleIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventExampleIn)

	if err != nil {
		return err
	}

	*o = EventExampleIn(varEventExampleIn)

	return err
}

type NullableEventExampleIn struct {
	value *EventExampleIn
	isSet bool
}

func (v NullableEventExampleIn) Get() *EventExampleIn {
	return v.value
}

func (v *NullableEventExampleIn) Set(val *EventExampleIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEventExampleIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEventExampleIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventExampleIn(val *EventExampleIn) *NullableEventExampleIn {
	return &NullableEventExampleIn{value: val, isSet: true}
}

func (v NullableEventExampleIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventExampleIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


