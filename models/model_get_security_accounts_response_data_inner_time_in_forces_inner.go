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

// checks if the GetSecurityAccountsResponseDataInnerTimeInForcesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSecurityAccountsResponseDataInnerTimeInForcesInner{}

// GetSecurityAccountsResponseDataInnerTimeInForcesInner struct for GetSecurityAccountsResponseDataInnerTimeInForcesInner
type GetSecurityAccountsResponseDataInnerTimeInForcesInner struct {
	Alias *string `json:"alias,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGetSecurityAccountsResponseDataInnerTimeInForcesInner instantiates a new GetSecurityAccountsResponseDataInnerTimeInForcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSecurityAccountsResponseDataInnerTimeInForcesInner() *GetSecurityAccountsResponseDataInnerTimeInForcesInner {
	this := GetSecurityAccountsResponseDataInnerTimeInForcesInner{}
	return &this
}

// NewGetSecurityAccountsResponseDataInnerTimeInForcesInnerWithDefaults instantiates a new GetSecurityAccountsResponseDataInnerTimeInForcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSecurityAccountsResponseDataInnerTimeInForcesInnerWithDefaults() *GetSecurityAccountsResponseDataInnerTimeInForcesInner {
	this := GetSecurityAccountsResponseDataInnerTimeInForcesInner{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) SetAlias(v string) {
	o.Alias = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetSecurityAccountsResponseDataInnerTimeInForcesInner) SetName(v string) {
	o.Name = &v
}

func (o GetSecurityAccountsResponseDataInnerTimeInForcesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSecurityAccountsResponseDataInnerTimeInForcesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner struct {
	value *GetSecurityAccountsResponseDataInnerTimeInForcesInner
	isSet bool
}

func (v NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner) Get() *GetSecurityAccountsResponseDataInnerTimeInForcesInner {
	return v.value
}

func (v *NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner) Set(val *GetSecurityAccountsResponseDataInnerTimeInForcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSecurityAccountsResponseDataInnerTimeInForcesInner(val *GetSecurityAccountsResponseDataInnerTimeInForcesInner) *NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner {
	return &NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner{value: val, isSet: true}
}

func (v NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSecurityAccountsResponseDataInnerTimeInForcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

