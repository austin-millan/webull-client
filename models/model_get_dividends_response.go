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

// checks if the GetDividendsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDividendsResponse{}

// GetDividendsResponse struct for GetDividendsResponse
type GetDividendsResponse struct {
	DividendTotal *string `json:"dividendTotal,omitempty"`
	HasTax *bool `json:"hasTax,omitempty"`
}

// NewGetDividendsResponse instantiates a new GetDividendsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDividendsResponse() *GetDividendsResponse {
	this := GetDividendsResponse{}
	return &this
}

// NewGetDividendsResponseWithDefaults instantiates a new GetDividendsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDividendsResponseWithDefaults() *GetDividendsResponse {
	this := GetDividendsResponse{}
	return &this
}

// GetDividendTotal returns the DividendTotal field value if set, zero value otherwise.
func (o *GetDividendsResponse) GetDividendTotal() string {
	if o == nil || IsNil(o.DividendTotal) {
		var ret string
		return ret
	}
	return *o.DividendTotal
}

// GetDividendTotalOk returns a tuple with the DividendTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividendsResponse) GetDividendTotalOk() (*string, bool) {
	if o == nil || IsNil(o.DividendTotal) {
		return nil, false
	}
	return o.DividendTotal, true
}

// HasDividendTotal returns a boolean if a field has been set.
func (o *GetDividendsResponse) HasDividendTotal() bool {
	if o != nil && !IsNil(o.DividendTotal) {
		return true
	}

	return false
}

// SetDividendTotal gets a reference to the given string and assigns it to the DividendTotal field.
func (o *GetDividendsResponse) SetDividendTotal(v string) {
	o.DividendTotal = &v
}

// GetHasTax returns the HasTax field value if set, zero value otherwise.
func (o *GetDividendsResponse) GetHasTax() bool {
	if o == nil || IsNil(o.HasTax) {
		var ret bool
		return ret
	}
	return *o.HasTax
}

// GetHasTaxOk returns a tuple with the HasTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividendsResponse) GetHasTaxOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTax) {
		return nil, false
	}
	return o.HasTax, true
}

// HasHasTax returns a boolean if a field has been set.
func (o *GetDividendsResponse) HasHasTax() bool {
	if o != nil && !IsNil(o.HasTax) {
		return true
	}

	return false
}

// SetHasTax gets a reference to the given bool and assigns it to the HasTax field.
func (o *GetDividendsResponse) SetHasTax(v bool) {
	o.HasTax = &v
}

func (o GetDividendsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDividendsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DividendTotal) {
		toSerialize["dividendTotal"] = o.DividendTotal
	}
	if !IsNil(o.HasTax) {
		toSerialize["hasTax"] = o.HasTax
	}
	return toSerialize, nil
}

type NullableGetDividendsResponse struct {
	value *GetDividendsResponse
	isSet bool
}

func (v NullableGetDividendsResponse) Get() *GetDividendsResponse {
	return v.value
}

func (v *NullableGetDividendsResponse) Set(val *GetDividendsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDividendsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDividendsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDividendsResponse(val *GetDividendsResponse) *NullableGetDividendsResponse {
	return &NullableGetDividendsResponse{value: val, isSet: true}
}

func (v NullableGetDividendsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDividendsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


