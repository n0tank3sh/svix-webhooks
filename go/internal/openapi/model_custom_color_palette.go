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

// checks if the CustomColorPalette type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomColorPalette{}

// CustomColorPalette struct for CustomColorPalette
type CustomColorPalette struct {
	BackgroundHover NullableString `json:"backgroundHover,omitempty"`
	BackgroundPrimary NullableString `json:"backgroundPrimary,omitempty"`
	BackgroundSecondary NullableString `json:"backgroundSecondary,omitempty"`
	ButtonPrimary NullableString `json:"buttonPrimary,omitempty"`
	InteractiveAccent NullableString `json:"interactiveAccent,omitempty"`
	NavigationAccent NullableString `json:"navigationAccent,omitempty"`
	Primary NullableString `json:"primary,omitempty"`
	TextDanger NullableString `json:"textDanger,omitempty"`
	TextPrimary NullableString `json:"textPrimary,omitempty"`
}

// NewCustomColorPalette instantiates a new CustomColorPalette object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomColorPalette() *CustomColorPalette {
	this := CustomColorPalette{}
	return &this
}

// NewCustomColorPaletteWithDefaults instantiates a new CustomColorPalette object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomColorPaletteWithDefaults() *CustomColorPalette {
	this := CustomColorPalette{}
	return &this
}

// GetBackgroundHover returns the BackgroundHover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetBackgroundHover() string {
	if o == nil || IsNil(o.BackgroundHover.Get()) {
		var ret string
		return ret
	}
	return *o.BackgroundHover.Get()
}

// GetBackgroundHoverOk returns a tuple with the BackgroundHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetBackgroundHoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackgroundHover.Get(), o.BackgroundHover.IsSet()
}

// HasBackgroundHover returns a boolean if a field has been set.
func (o *CustomColorPalette) HasBackgroundHover() bool {
	if o != nil && o.BackgroundHover.IsSet() {
		return true
	}

	return false
}

// SetBackgroundHover gets a reference to the given NullableString and assigns it to the BackgroundHover field.
func (o *CustomColorPalette) SetBackgroundHover(v string) {
	o.BackgroundHover.Set(&v)
}
// SetBackgroundHoverNil sets the value for BackgroundHover to be an explicit nil
func (o *CustomColorPalette) SetBackgroundHoverNil() {
	o.BackgroundHover.Set(nil)
}

// UnsetBackgroundHover ensures that no value is present for BackgroundHover, not even an explicit nil
func (o *CustomColorPalette) UnsetBackgroundHover() {
	o.BackgroundHover.Unset()
}

// GetBackgroundPrimary returns the BackgroundPrimary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetBackgroundPrimary() string {
	if o == nil || IsNil(o.BackgroundPrimary.Get()) {
		var ret string
		return ret
	}
	return *o.BackgroundPrimary.Get()
}

// GetBackgroundPrimaryOk returns a tuple with the BackgroundPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetBackgroundPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackgroundPrimary.Get(), o.BackgroundPrimary.IsSet()
}

// HasBackgroundPrimary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasBackgroundPrimary() bool {
	if o != nil && o.BackgroundPrimary.IsSet() {
		return true
	}

	return false
}

// SetBackgroundPrimary gets a reference to the given NullableString and assigns it to the BackgroundPrimary field.
func (o *CustomColorPalette) SetBackgroundPrimary(v string) {
	o.BackgroundPrimary.Set(&v)
}
// SetBackgroundPrimaryNil sets the value for BackgroundPrimary to be an explicit nil
func (o *CustomColorPalette) SetBackgroundPrimaryNil() {
	o.BackgroundPrimary.Set(nil)
}

// UnsetBackgroundPrimary ensures that no value is present for BackgroundPrimary, not even an explicit nil
func (o *CustomColorPalette) UnsetBackgroundPrimary() {
	o.BackgroundPrimary.Unset()
}

// GetBackgroundSecondary returns the BackgroundSecondary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetBackgroundSecondary() string {
	if o == nil || IsNil(o.BackgroundSecondary.Get()) {
		var ret string
		return ret
	}
	return *o.BackgroundSecondary.Get()
}

// GetBackgroundSecondaryOk returns a tuple with the BackgroundSecondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetBackgroundSecondaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackgroundSecondary.Get(), o.BackgroundSecondary.IsSet()
}

// HasBackgroundSecondary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasBackgroundSecondary() bool {
	if o != nil && o.BackgroundSecondary.IsSet() {
		return true
	}

	return false
}

// SetBackgroundSecondary gets a reference to the given NullableString and assigns it to the BackgroundSecondary field.
func (o *CustomColorPalette) SetBackgroundSecondary(v string) {
	o.BackgroundSecondary.Set(&v)
}
// SetBackgroundSecondaryNil sets the value for BackgroundSecondary to be an explicit nil
func (o *CustomColorPalette) SetBackgroundSecondaryNil() {
	o.BackgroundSecondary.Set(nil)
}

// UnsetBackgroundSecondary ensures that no value is present for BackgroundSecondary, not even an explicit nil
func (o *CustomColorPalette) UnsetBackgroundSecondary() {
	o.BackgroundSecondary.Unset()
}

// GetButtonPrimary returns the ButtonPrimary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetButtonPrimary() string {
	if o == nil || IsNil(o.ButtonPrimary.Get()) {
		var ret string
		return ret
	}
	return *o.ButtonPrimary.Get()
}

// GetButtonPrimaryOk returns a tuple with the ButtonPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetButtonPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ButtonPrimary.Get(), o.ButtonPrimary.IsSet()
}

// HasButtonPrimary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasButtonPrimary() bool {
	if o != nil && o.ButtonPrimary.IsSet() {
		return true
	}

	return false
}

// SetButtonPrimary gets a reference to the given NullableString and assigns it to the ButtonPrimary field.
func (o *CustomColorPalette) SetButtonPrimary(v string) {
	o.ButtonPrimary.Set(&v)
}
// SetButtonPrimaryNil sets the value for ButtonPrimary to be an explicit nil
func (o *CustomColorPalette) SetButtonPrimaryNil() {
	o.ButtonPrimary.Set(nil)
}

// UnsetButtonPrimary ensures that no value is present for ButtonPrimary, not even an explicit nil
func (o *CustomColorPalette) UnsetButtonPrimary() {
	o.ButtonPrimary.Unset()
}

// GetInteractiveAccent returns the InteractiveAccent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetInteractiveAccent() string {
	if o == nil || IsNil(o.InteractiveAccent.Get()) {
		var ret string
		return ret
	}
	return *o.InteractiveAccent.Get()
}

// GetInteractiveAccentOk returns a tuple with the InteractiveAccent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetInteractiveAccentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InteractiveAccent.Get(), o.InteractiveAccent.IsSet()
}

// HasInteractiveAccent returns a boolean if a field has been set.
func (o *CustomColorPalette) HasInteractiveAccent() bool {
	if o != nil && o.InteractiveAccent.IsSet() {
		return true
	}

	return false
}

// SetInteractiveAccent gets a reference to the given NullableString and assigns it to the InteractiveAccent field.
func (o *CustomColorPalette) SetInteractiveAccent(v string) {
	o.InteractiveAccent.Set(&v)
}
// SetInteractiveAccentNil sets the value for InteractiveAccent to be an explicit nil
func (o *CustomColorPalette) SetInteractiveAccentNil() {
	o.InteractiveAccent.Set(nil)
}

// UnsetInteractiveAccent ensures that no value is present for InteractiveAccent, not even an explicit nil
func (o *CustomColorPalette) UnsetInteractiveAccent() {
	o.InteractiveAccent.Unset()
}

// GetNavigationAccent returns the NavigationAccent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetNavigationAccent() string {
	if o == nil || IsNil(o.NavigationAccent.Get()) {
		var ret string
		return ret
	}
	return *o.NavigationAccent.Get()
}

// GetNavigationAccentOk returns a tuple with the NavigationAccent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetNavigationAccentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NavigationAccent.Get(), o.NavigationAccent.IsSet()
}

// HasNavigationAccent returns a boolean if a field has been set.
func (o *CustomColorPalette) HasNavigationAccent() bool {
	if o != nil && o.NavigationAccent.IsSet() {
		return true
	}

	return false
}

// SetNavigationAccent gets a reference to the given NullableString and assigns it to the NavigationAccent field.
func (o *CustomColorPalette) SetNavigationAccent(v string) {
	o.NavigationAccent.Set(&v)
}
// SetNavigationAccentNil sets the value for NavigationAccent to be an explicit nil
func (o *CustomColorPalette) SetNavigationAccentNil() {
	o.NavigationAccent.Set(nil)
}

// UnsetNavigationAccent ensures that no value is present for NavigationAccent, not even an explicit nil
func (o *CustomColorPalette) UnsetNavigationAccent() {
	o.NavigationAccent.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetPrimary() string {
	if o == nil || IsNil(o.Primary.Get()) {
		var ret string
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasPrimary() bool {
	if o != nil && o.Primary.IsSet() {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given NullableString and assigns it to the Primary field.
func (o *CustomColorPalette) SetPrimary(v string) {
	o.Primary.Set(&v)
}
// SetPrimaryNil sets the value for Primary to be an explicit nil
func (o *CustomColorPalette) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
func (o *CustomColorPalette) UnsetPrimary() {
	o.Primary.Unset()
}

// GetTextDanger returns the TextDanger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetTextDanger() string {
	if o == nil || IsNil(o.TextDanger.Get()) {
		var ret string
		return ret
	}
	return *o.TextDanger.Get()
}

// GetTextDangerOk returns a tuple with the TextDanger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetTextDangerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextDanger.Get(), o.TextDanger.IsSet()
}

// HasTextDanger returns a boolean if a field has been set.
func (o *CustomColorPalette) HasTextDanger() bool {
	if o != nil && o.TextDanger.IsSet() {
		return true
	}

	return false
}

// SetTextDanger gets a reference to the given NullableString and assigns it to the TextDanger field.
func (o *CustomColorPalette) SetTextDanger(v string) {
	o.TextDanger.Set(&v)
}
// SetTextDangerNil sets the value for TextDanger to be an explicit nil
func (o *CustomColorPalette) SetTextDangerNil() {
	o.TextDanger.Set(nil)
}

// UnsetTextDanger ensures that no value is present for TextDanger, not even an explicit nil
func (o *CustomColorPalette) UnsetTextDanger() {
	o.TextDanger.Unset()
}

// GetTextPrimary returns the TextPrimary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetTextPrimary() string {
	if o == nil || IsNil(o.TextPrimary.Get()) {
		var ret string
		return ret
	}
	return *o.TextPrimary.Get()
}

// GetTextPrimaryOk returns a tuple with the TextPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetTextPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextPrimary.Get(), o.TextPrimary.IsSet()
}

// HasTextPrimary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasTextPrimary() bool {
	if o != nil && o.TextPrimary.IsSet() {
		return true
	}

	return false
}

// SetTextPrimary gets a reference to the given NullableString and assigns it to the TextPrimary field.
func (o *CustomColorPalette) SetTextPrimary(v string) {
	o.TextPrimary.Set(&v)
}
// SetTextPrimaryNil sets the value for TextPrimary to be an explicit nil
func (o *CustomColorPalette) SetTextPrimaryNil() {
	o.TextPrimary.Set(nil)
}

// UnsetTextPrimary ensures that no value is present for TextPrimary, not even an explicit nil
func (o *CustomColorPalette) UnsetTextPrimary() {
	o.TextPrimary.Unset()
}

func (o CustomColorPalette) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomColorPalette) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackgroundHover.IsSet() {
		toSerialize["backgroundHover"] = o.BackgroundHover.Get()
	}
	if o.BackgroundPrimary.IsSet() {
		toSerialize["backgroundPrimary"] = o.BackgroundPrimary.Get()
	}
	if o.BackgroundSecondary.IsSet() {
		toSerialize["backgroundSecondary"] = o.BackgroundSecondary.Get()
	}
	if o.ButtonPrimary.IsSet() {
		toSerialize["buttonPrimary"] = o.ButtonPrimary.Get()
	}
	if o.InteractiveAccent.IsSet() {
		toSerialize["interactiveAccent"] = o.InteractiveAccent.Get()
	}
	if o.NavigationAccent.IsSet() {
		toSerialize["navigationAccent"] = o.NavigationAccent.Get()
	}
	if o.Primary.IsSet() {
		toSerialize["primary"] = o.Primary.Get()
	}
	if o.TextDanger.IsSet() {
		toSerialize["textDanger"] = o.TextDanger.Get()
	}
	if o.TextPrimary.IsSet() {
		toSerialize["textPrimary"] = o.TextPrimary.Get()
	}
	return toSerialize, nil
}

type NullableCustomColorPalette struct {
	value *CustomColorPalette
	isSet bool
}

func (v NullableCustomColorPalette) Get() *CustomColorPalette {
	return v.value
}

func (v *NullableCustomColorPalette) Set(val *CustomColorPalette) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomColorPalette) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomColorPalette) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomColorPalette(val *CustomColorPalette) *NullableCustomColorPalette {
	return &NullableCustomColorPalette{value: val, isSet: true}
}

func (v NullableCustomColorPalette) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomColorPalette) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


