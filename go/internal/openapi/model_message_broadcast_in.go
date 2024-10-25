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

// checks if the MessageBroadcastIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageBroadcastIn{}

// MessageBroadcastIn struct for MessageBroadcastIn
type MessageBroadcastIn struct {
	// List of free-form identifiers that endpoints can filter by
	Channels []string `json:"channels,omitempty"`
	// Optional unique identifier for the message
	EventId NullableString `json:"eventId,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The event type's name
	EventType string `json:"eventType" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	Payload map[string]interface{} `json:"payload"`
	// Optional number of hours to retain the message payload. Note that this is mutually exclusive with `payloadRetentionPeriod`.
	PayloadRetentionHours NullableInt64 `json:"payloadRetentionHours,omitempty"`
	// Optional number of days to retain the message payload. Defaults to 90. Note that this is mutually exclusive with `payloadRetentionHours`.
	PayloadRetentionPeriod NullableInt64 `json:"payloadRetentionPeriod,omitempty"`
}

type _MessageBroadcastIn MessageBroadcastIn

// NewMessageBroadcastIn instantiates a new MessageBroadcastIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageBroadcastIn(eventType string, payload map[string]interface{}) *MessageBroadcastIn {
	this := MessageBroadcastIn{}
	this.EventType = eventType
	this.Payload = payload
	var payloadRetentionPeriod int64 = 90
	this.PayloadRetentionPeriod = *NewNullableInt64(&payloadRetentionPeriod)
	return &this
}

// NewMessageBroadcastInWithDefaults instantiates a new MessageBroadcastIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageBroadcastInWithDefaults() *MessageBroadcastIn {
	this := MessageBroadcastIn{}
	var payloadRetentionPeriod int64 = 90
	this.PayloadRetentionPeriod = *NewNullableInt64(&payloadRetentionPeriod)
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageBroadcastIn) GetChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageBroadcastIn) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *MessageBroadcastIn) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *MessageBroadcastIn) SetChannels(v []string) {
	o.Channels = v
}

// GetEventId returns the EventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageBroadcastIn) GetEventId() string {
	if o == nil || IsNil(o.EventId.Get()) {
		var ret string
		return ret
	}
	return *o.EventId.Get()
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageBroadcastIn) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventId.Get(), o.EventId.IsSet()
}

// HasEventId returns a boolean if a field has been set.
func (o *MessageBroadcastIn) HasEventId() bool {
	if o != nil && o.EventId.IsSet() {
		return true
	}

	return false
}

// SetEventId gets a reference to the given NullableString and assigns it to the EventId field.
func (o *MessageBroadcastIn) SetEventId(v string) {
	o.EventId.Set(&v)
}
// SetEventIdNil sets the value for EventId to be an explicit nil
func (o *MessageBroadcastIn) SetEventIdNil() {
	o.EventId.Set(nil)
}

// UnsetEventId ensures that no value is present for EventId, not even an explicit nil
func (o *MessageBroadcastIn) UnsetEventId() {
	o.EventId.Unset()
}

// GetEventType returns the EventType field value
func (o *MessageBroadcastIn) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MessageBroadcastIn) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MessageBroadcastIn) SetEventType(v string) {
	o.EventType = v
}

// GetPayload returns the Payload field value
func (o *MessageBroadcastIn) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *MessageBroadcastIn) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *MessageBroadcastIn) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetPayloadRetentionHours returns the PayloadRetentionHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageBroadcastIn) GetPayloadRetentionHours() int64 {
	if o == nil || IsNil(o.PayloadRetentionHours.Get()) {
		var ret int64
		return ret
	}
	return *o.PayloadRetentionHours.Get()
}

// GetPayloadRetentionHoursOk returns a tuple with the PayloadRetentionHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageBroadcastIn) GetPayloadRetentionHoursOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayloadRetentionHours.Get(), o.PayloadRetentionHours.IsSet()
}

// HasPayloadRetentionHours returns a boolean if a field has been set.
func (o *MessageBroadcastIn) HasPayloadRetentionHours() bool {
	if o != nil && o.PayloadRetentionHours.IsSet() {
		return true
	}

	return false
}

// SetPayloadRetentionHours gets a reference to the given NullableInt64 and assigns it to the PayloadRetentionHours field.
func (o *MessageBroadcastIn) SetPayloadRetentionHours(v int64) {
	o.PayloadRetentionHours.Set(&v)
}
// SetPayloadRetentionHoursNil sets the value for PayloadRetentionHours to be an explicit nil
func (o *MessageBroadcastIn) SetPayloadRetentionHoursNil() {
	o.PayloadRetentionHours.Set(nil)
}

// UnsetPayloadRetentionHours ensures that no value is present for PayloadRetentionHours, not even an explicit nil
func (o *MessageBroadcastIn) UnsetPayloadRetentionHours() {
	o.PayloadRetentionHours.Unset()
}

// GetPayloadRetentionPeriod returns the PayloadRetentionPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageBroadcastIn) GetPayloadRetentionPeriod() int64 {
	if o == nil || IsNil(o.PayloadRetentionPeriod.Get()) {
		var ret int64
		return ret
	}
	return *o.PayloadRetentionPeriod.Get()
}

// GetPayloadRetentionPeriodOk returns a tuple with the PayloadRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageBroadcastIn) GetPayloadRetentionPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayloadRetentionPeriod.Get(), o.PayloadRetentionPeriod.IsSet()
}

// HasPayloadRetentionPeriod returns a boolean if a field has been set.
func (o *MessageBroadcastIn) HasPayloadRetentionPeriod() bool {
	if o != nil && o.PayloadRetentionPeriod.IsSet() {
		return true
	}

	return false
}

// SetPayloadRetentionPeriod gets a reference to the given NullableInt64 and assigns it to the PayloadRetentionPeriod field.
func (o *MessageBroadcastIn) SetPayloadRetentionPeriod(v int64) {
	o.PayloadRetentionPeriod.Set(&v)
}
// SetPayloadRetentionPeriodNil sets the value for PayloadRetentionPeriod to be an explicit nil
func (o *MessageBroadcastIn) SetPayloadRetentionPeriodNil() {
	o.PayloadRetentionPeriod.Set(nil)
}

// UnsetPayloadRetentionPeriod ensures that no value is present for PayloadRetentionPeriod, not even an explicit nil
func (o *MessageBroadcastIn) UnsetPayloadRetentionPeriod() {
	o.PayloadRetentionPeriod.Unset()
}

func (o MessageBroadcastIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageBroadcastIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if o.EventId.IsSet() {
		toSerialize["eventId"] = o.EventId.Get()
	}
	toSerialize["eventType"] = o.EventType
	toSerialize["payload"] = o.Payload
	if o.PayloadRetentionHours.IsSet() {
		toSerialize["payloadRetentionHours"] = o.PayloadRetentionHours.Get()
	}
	if o.PayloadRetentionPeriod.IsSet() {
		toSerialize["payloadRetentionPeriod"] = o.PayloadRetentionPeriod.Get()
	}
	return toSerialize, nil
}

func (o *MessageBroadcastIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventType",
		"payload",
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

	varMessageBroadcastIn := _MessageBroadcastIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageBroadcastIn)

	if err != nil {
		return err
	}

	*o = MessageBroadcastIn(varMessageBroadcastIn)

	return err
}

type NullableMessageBroadcastIn struct {
	value *MessageBroadcastIn
	isSet bool
}

func (v NullableMessageBroadcastIn) Get() *MessageBroadcastIn {
	return v.value
}

func (v *NullableMessageBroadcastIn) Set(val *MessageBroadcastIn) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageBroadcastIn) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageBroadcastIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageBroadcastIn(val *MessageBroadcastIn) *NullableMessageBroadcastIn {
	return &NullableMessageBroadcastIn{value: val, isSet: true}
}

func (v NullableMessageBroadcastIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageBroadcastIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


