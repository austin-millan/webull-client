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

// ExecutionType the model 'ExecutionType'
type ExecutionType string

// List of ExecutionType
const (
	MARKET ExecutionType = "market"
	LIMIT ExecutionType = "limit"
)

// All allowed values of ExecutionType enum
var AllowedExecutionTypeEnumValues = []ExecutionType{
	"market",
	"limit",
}

func (v *ExecutionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExecutionType(value)
	for _, existing := range AllowedExecutionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExecutionType", value)
}

// NewExecutionTypeFromValue returns a pointer to a valid ExecutionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExecutionTypeFromValue(v string) (*ExecutionType, error) {
	ev := ExecutionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExecutionType: valid values are %v", v, AllowedExecutionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExecutionType) IsValid() bool {
	for _, existing := range AllowedExecutionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionType value
func (v ExecutionType) Ptr() *ExecutionType {
	return &v
}

type NullableExecutionType struct {
	value *ExecutionType
	isSet bool
}

func (v NullableExecutionType) Get() *ExecutionType {
	return v.value
}

func (v *NullableExecutionType) Set(val *ExecutionType) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionType) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionType(val *ExecutionType) *NullableExecutionType {
	return &NullableExecutionType{value: val, isSet: true}
}

func (v NullableExecutionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

