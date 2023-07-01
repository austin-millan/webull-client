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
	"fmt"
)

// TimeInForce the model 'TimeInForce'
type TimeInForce string

// List of TimeInForce
const (
	GTC TimeInForce = "GTC"
	DAY TimeInForce = "DAY"
)

// All allowed values of TimeInForce enum
var AllowedTimeInForceEnumValues = []TimeInForce{
	"GTC",
	"DAY",
}

func (v *TimeInForce) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeInForce(value)
	for _, existing := range AllowedTimeInForceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeInForce", value)
}

// NewTimeInForceFromValue returns a pointer to a valid TimeInForce
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeInForceFromValue(v string) (*TimeInForce, error) {
	ev := TimeInForce(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeInForce: valid values are %v", v, AllowedTimeInForceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeInForce) IsValid() bool {
	for _, existing := range AllowedTimeInForceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeInForce value
func (v TimeInForce) Ptr() *TimeInForce {
	return &v
}

type NullableTimeInForce struct {
	value *TimeInForce
	isSet bool
}

func (v NullableTimeInForce) Get() *TimeInForce {
	return v.value
}

func (v *NullableTimeInForce) Set(val *TimeInForce) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeInForce) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeInForce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeInForce(val *TimeInForce) *NullableTimeInForce {
	return &NullableTimeInForce{value: val, isSet: true}
}

func (v NullableTimeInForce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeInForce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

