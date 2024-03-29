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

// checks if the GetFundamentalsResponseSimpleStatementInnerLabelsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFundamentalsResponseSimpleStatementInnerLabelsInner{}

// GetFundamentalsResponseSimpleStatementInnerLabelsInner struct for GetFundamentalsResponseSimpleStatementInnerLabelsInner
type GetFundamentalsResponseSimpleStatementInnerLabelsInner struct {
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGetFundamentalsResponseSimpleStatementInnerLabelsInner instantiates a new GetFundamentalsResponseSimpleStatementInnerLabelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundamentalsResponseSimpleStatementInnerLabelsInner() *GetFundamentalsResponseSimpleStatementInnerLabelsInner {
	this := GetFundamentalsResponseSimpleStatementInnerLabelsInner{}
	return &this
}

// NewGetFundamentalsResponseSimpleStatementInnerLabelsInnerWithDefaults instantiates a new GetFundamentalsResponseSimpleStatementInnerLabelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundamentalsResponseSimpleStatementInnerLabelsInnerWithDefaults() *GetFundamentalsResponseSimpleStatementInnerLabelsInner {
	this := GetFundamentalsResponseSimpleStatementInnerLabelsInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetFundamentalsResponseSimpleStatementInnerLabelsInner) SetName(v string) {
	o.Name = &v
}

func (o GetFundamentalsResponseSimpleStatementInnerLabelsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundamentalsResponseSimpleStatementInnerLabelsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner struct {
	value *GetFundamentalsResponseSimpleStatementInnerLabelsInner
	isSet bool
}

func (v NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner) Get() *GetFundamentalsResponseSimpleStatementInnerLabelsInner {
	return v.value
}

func (v *NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner) Set(val *GetFundamentalsResponseSimpleStatementInnerLabelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFundamentalsResponseSimpleStatementInnerLabelsInner(val *GetFundamentalsResponseSimpleStatementInnerLabelsInner) *NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner {
	return &NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner{value: val, isSet: true}
}

func (v NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundamentalsResponseSimpleStatementInnerLabelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


