/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EventTypeExampleOut struct for EventTypeExampleOut
type EventTypeExampleOut struct {
	Example map[string]interface{} `json:"example"`
}

// NewEventTypeExampleOut instantiates a new EventTypeExampleOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeExampleOut(example map[string]interface{}) *EventTypeExampleOut {
	this := EventTypeExampleOut{}
	this.Example = example
	return &this
}

// NewEventTypeExampleOutWithDefaults instantiates a new EventTypeExampleOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeExampleOutWithDefaults() *EventTypeExampleOut {
	this := EventTypeExampleOut{}
	return &this
}

// GetExample returns the Example field value
func (o *EventTypeExampleOut) GetExample() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Example
}

// GetExampleOk returns a tuple with the Example field value
// and a boolean to check if the value has been set.
func (o *EventTypeExampleOut) GetExampleOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Example, true
}

// SetExample sets field value
func (o *EventTypeExampleOut) SetExample(v map[string]interface{}) {
	o.Example = v
}

func (o EventTypeExampleOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["example"] = o.Example
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypeExampleOut struct {
	value *EventTypeExampleOut
	isSet bool
}

func (v NullableEventTypeExampleOut) Get() *EventTypeExampleOut {
	return v.value
}

func (v *NullableEventTypeExampleOut) Set(val *EventTypeExampleOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeExampleOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeExampleOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeExampleOut(val *EventTypeExampleOut) *NullableEventTypeExampleOut {
	return &NullableEventTypeExampleOut{value: val, isSet: true}
}

func (v NullableEventTypeExampleOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeExampleOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


