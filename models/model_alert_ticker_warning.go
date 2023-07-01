/*
Webull API

Webull API Documentation

API version: 3.0.1
Contact: austin.millan@protonmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AlertTickerWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertTickerWarning{}

// AlertTickerWarning struct for AlertTickerWarning
type AlertTickerWarning struct {
	Remove *bool `json:"remove,omitempty"`
}

// NewAlertTickerWarning instantiates a new AlertTickerWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertTickerWarning() *AlertTickerWarning {
	this := AlertTickerWarning{}
	return &this
}

// NewAlertTickerWarningWithDefaults instantiates a new AlertTickerWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertTickerWarningWithDefaults() *AlertTickerWarning {
	this := AlertTickerWarning{}
	return &this
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *AlertTickerWarning) GetRemove() bool {
	if o == nil || IsNil(o.Remove) {
		var ret bool
		return ret
	}
	return *o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertTickerWarning) GetRemoveOk() (*bool, bool) {
	if o == nil || IsNil(o.Remove) {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *AlertTickerWarning) HasRemove() bool {
	if o != nil && !IsNil(o.Remove) {
		return true
	}

	return false
}

// SetRemove gets a reference to the given bool and assigns it to the Remove field.
func (o *AlertTickerWarning) SetRemove(v bool) {
	o.Remove = &v
}

func (o AlertTickerWarning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertTickerWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Remove) {
		toSerialize["remove"] = o.Remove
	}
	return toSerialize, nil
}

type NullableAlertTickerWarning struct {
	value *AlertTickerWarning
	isSet bool
}

func (v NullableAlertTickerWarning) Get() *AlertTickerWarning {
	return v.value
}

func (v *NullableAlertTickerWarning) Set(val *AlertTickerWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertTickerWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertTickerWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertTickerWarning(val *AlertTickerWarning) *NullableAlertTickerWarning {
	return &NullableAlertTickerWarning{value: val, isSet: true}
}

func (v NullableAlertTickerWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertTickerWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


