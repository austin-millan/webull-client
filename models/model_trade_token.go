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

// checks if the TradeToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradeToken{}

// TradeToken struct for TradeToken
type TradeToken struct {
	Username *string `json:"username,omitempty"`
}

// NewTradeToken instantiates a new TradeToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeToken() *TradeToken {
	this := TradeToken{}
	return &this
}

// NewTradeTokenWithDefaults instantiates a new TradeToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeTokenWithDefaults() *TradeToken {
	this := TradeToken{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TradeToken) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeToken) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TradeToken) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TradeToken) SetUsername(v string) {
	o.Username = &v
}

func (o TradeToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableTradeToken struct {
	value *TradeToken
	isSet bool
}

func (v NullableTradeToken) Get() *TradeToken {
	return v.value
}

func (v *NullableTradeToken) Set(val *TradeToken) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeToken) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeToken(val *TradeToken) *NullableTradeToken {
	return &NullableTradeToken{value: val, isSet: true}
}

func (v NullableTradeToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


