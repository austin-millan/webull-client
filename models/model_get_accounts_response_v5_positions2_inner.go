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

// checks if the GetAccountsResponseV5Positions2Inner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountsResponseV5Positions2Inner{}

// GetAccountsResponseV5Positions2Inner struct for GetAccountsResponseV5Positions2Inner
type GetAccountsResponseV5Positions2Inner struct {
	ComboId *string `json:"comboId,omitempty"`
	ComboTickerType *string `json:"comboTickerType,omitempty"`
	Positions []GetAccountsResponseV5Positions2InnerPositionsInner `json:"positions,omitempty"`
}

// NewGetAccountsResponseV5Positions2Inner instantiates a new GetAccountsResponseV5Positions2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountsResponseV5Positions2Inner() *GetAccountsResponseV5Positions2Inner {
	this := GetAccountsResponseV5Positions2Inner{}
	return &this
}

// NewGetAccountsResponseV5Positions2InnerWithDefaults instantiates a new GetAccountsResponseV5Positions2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountsResponseV5Positions2InnerWithDefaults() *GetAccountsResponseV5Positions2Inner {
	this := GetAccountsResponseV5Positions2Inner{}
	return &this
}

// GetComboId returns the ComboId field value if set, zero value otherwise.
func (o *GetAccountsResponseV5Positions2Inner) GetComboId() string {
	if o == nil || IsNil(o.ComboId) {
		var ret string
		return ret
	}
	return *o.ComboId
}

// GetComboIdOk returns a tuple with the ComboId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5Positions2Inner) GetComboIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComboId) {
		return nil, false
	}
	return o.ComboId, true
}

// HasComboId returns a boolean if a field has been set.
func (o *GetAccountsResponseV5Positions2Inner) HasComboId() bool {
	if o != nil && !IsNil(o.ComboId) {
		return true
	}

	return false
}

// SetComboId gets a reference to the given string and assigns it to the ComboId field.
func (o *GetAccountsResponseV5Positions2Inner) SetComboId(v string) {
	o.ComboId = &v
}

// GetComboTickerType returns the ComboTickerType field value if set, zero value otherwise.
func (o *GetAccountsResponseV5Positions2Inner) GetComboTickerType() string {
	if o == nil || IsNil(o.ComboTickerType) {
		var ret string
		return ret
	}
	return *o.ComboTickerType
}

// GetComboTickerTypeOk returns a tuple with the ComboTickerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5Positions2Inner) GetComboTickerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComboTickerType) {
		return nil, false
	}
	return o.ComboTickerType, true
}

// HasComboTickerType returns a boolean if a field has been set.
func (o *GetAccountsResponseV5Positions2Inner) HasComboTickerType() bool {
	if o != nil && !IsNil(o.ComboTickerType) {
		return true
	}

	return false
}

// SetComboTickerType gets a reference to the given string and assigns it to the ComboTickerType field.
func (o *GetAccountsResponseV5Positions2Inner) SetComboTickerType(v string) {
	o.ComboTickerType = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *GetAccountsResponseV5Positions2Inner) GetPositions() []GetAccountsResponseV5Positions2InnerPositionsInner {
	if o == nil || IsNil(o.Positions) {
		var ret []GetAccountsResponseV5Positions2InnerPositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5Positions2Inner) GetPositionsOk() ([]GetAccountsResponseV5Positions2InnerPositionsInner, bool) {
	if o == nil || IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *GetAccountsResponseV5Positions2Inner) HasPositions() bool {
	if o != nil && !IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []GetAccountsResponseV5Positions2InnerPositionsInner and assigns it to the Positions field.
func (o *GetAccountsResponseV5Positions2Inner) SetPositions(v []GetAccountsResponseV5Positions2InnerPositionsInner) {
	o.Positions = v
}

func (o GetAccountsResponseV5Positions2Inner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountsResponseV5Positions2Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComboId) {
		toSerialize["comboId"] = o.ComboId
	}
	if !IsNil(o.ComboTickerType) {
		toSerialize["comboTickerType"] = o.ComboTickerType
	}
	if !IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	return toSerialize, nil
}

type NullableGetAccountsResponseV5Positions2Inner struct {
	value *GetAccountsResponseV5Positions2Inner
	isSet bool
}

func (v NullableGetAccountsResponseV5Positions2Inner) Get() *GetAccountsResponseV5Positions2Inner {
	return v.value
}

func (v *NullableGetAccountsResponseV5Positions2Inner) Set(val *GetAccountsResponseV5Positions2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountsResponseV5Positions2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountsResponseV5Positions2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountsResponseV5Positions2Inner(val *GetAccountsResponseV5Positions2Inner) *NullableGetAccountsResponseV5Positions2Inner {
	return &NullableGetAccountsResponseV5Positions2Inner{value: val, isSet: true}
}

func (v NullableGetAccountsResponseV5Positions2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountsResponseV5Positions2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


