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

// checks if the GetOrdersItemOrdersInnerTickerPriceDefineListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrdersItemOrdersInnerTickerPriceDefineListInner{}

// GetOrdersItemOrdersInnerTickerPriceDefineListInner struct for GetOrdersItemOrdersInnerTickerPriceDefineListInner
type GetOrdersItemOrdersInnerTickerPriceDefineListInner struct {
	ContainBegin *bool `json:"containBegin,omitempty"`
	ContainEnd *bool `json:"containEnd,omitempty"`
	PriceUnit *string `json:"priceUnit,omitempty"`
	RangeBegin *string `json:"rangeBegin,omitempty"`
	RangeEnd *string `json:"rangeEnd,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
}

// NewGetOrdersItemOrdersInnerTickerPriceDefineListInner instantiates a new GetOrdersItemOrdersInnerTickerPriceDefineListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrdersItemOrdersInnerTickerPriceDefineListInner() *GetOrdersItemOrdersInnerTickerPriceDefineListInner {
	this := GetOrdersItemOrdersInnerTickerPriceDefineListInner{}
	return &this
}

// NewGetOrdersItemOrdersInnerTickerPriceDefineListInnerWithDefaults instantiates a new GetOrdersItemOrdersInnerTickerPriceDefineListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrdersItemOrdersInnerTickerPriceDefineListInnerWithDefaults() *GetOrdersItemOrdersInnerTickerPriceDefineListInner {
	this := GetOrdersItemOrdersInnerTickerPriceDefineListInner{}
	return &this
}

// GetContainBegin returns the ContainBegin field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetContainBegin() bool {
	if o == nil || IsNil(o.ContainBegin) {
		var ret bool
		return ret
	}
	return *o.ContainBegin
}

// GetContainBeginOk returns a tuple with the ContainBegin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetContainBeginOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainBegin) {
		return nil, false
	}
	return o.ContainBegin, true
}

// HasContainBegin returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) HasContainBegin() bool {
	if o != nil && !IsNil(o.ContainBegin) {
		return true
	}

	return false
}

// SetContainBegin gets a reference to the given bool and assigns it to the ContainBegin field.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) SetContainBegin(v bool) {
	o.ContainBegin = &v
}

// GetContainEnd returns the ContainEnd field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetContainEnd() bool {
	if o == nil || IsNil(o.ContainEnd) {
		var ret bool
		return ret
	}
	return *o.ContainEnd
}

// GetContainEndOk returns a tuple with the ContainEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetContainEndOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainEnd) {
		return nil, false
	}
	return o.ContainEnd, true
}

// HasContainEnd returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) HasContainEnd() bool {
	if o != nil && !IsNil(o.ContainEnd) {
		return true
	}

	return false
}

// SetContainEnd gets a reference to the given bool and assigns it to the ContainEnd field.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) SetContainEnd(v bool) {
	o.ContainEnd = &v
}

// GetPriceUnit returns the PriceUnit field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetPriceUnit() string {
	if o == nil || IsNil(o.PriceUnit) {
		var ret string
		return ret
	}
	return *o.PriceUnit
}

// GetPriceUnitOk returns a tuple with the PriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetPriceUnitOk() (*string, bool) {
	if o == nil || IsNil(o.PriceUnit) {
		return nil, false
	}
	return o.PriceUnit, true
}

// HasPriceUnit returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) HasPriceUnit() bool {
	if o != nil && !IsNil(o.PriceUnit) {
		return true
	}

	return false
}

// SetPriceUnit gets a reference to the given string and assigns it to the PriceUnit field.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) SetPriceUnit(v string) {
	o.PriceUnit = &v
}

// GetRangeBegin returns the RangeBegin field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetRangeBegin() string {
	if o == nil || IsNil(o.RangeBegin) {
		var ret string
		return ret
	}
	return *o.RangeBegin
}

// GetRangeBeginOk returns a tuple with the RangeBegin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetRangeBeginOk() (*string, bool) {
	if o == nil || IsNil(o.RangeBegin) {
		return nil, false
	}
	return o.RangeBegin, true
}

// HasRangeBegin returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) HasRangeBegin() bool {
	if o != nil && !IsNil(o.RangeBegin) {
		return true
	}

	return false
}

// SetRangeBegin gets a reference to the given string and assigns it to the RangeBegin field.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) SetRangeBegin(v string) {
	o.RangeBegin = &v
}

// GetRangeEnd returns the RangeEnd field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetRangeEnd() string {
	if o == nil || IsNil(o.RangeEnd) {
		var ret string
		return ret
	}
	return *o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetRangeEndOk() (*string, bool) {
	if o == nil || IsNil(o.RangeEnd) {
		return nil, false
	}
	return o.RangeEnd, true
}

// HasRangeEnd returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) HasRangeEnd() bool {
	if o != nil && !IsNil(o.RangeEnd) {
		return true
	}

	return false
}

// SetRangeEnd gets a reference to the given string and assigns it to the RangeEnd field.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) SetRangeEnd(v string) {
	o.RangeEnd = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *GetOrdersItemOrdersInnerTickerPriceDefineListInner) SetTickerId(v int32) {
	o.TickerId = &v
}

func (o GetOrdersItemOrdersInnerTickerPriceDefineListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrdersItemOrdersInnerTickerPriceDefineListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainBegin) {
		toSerialize["containBegin"] = o.ContainBegin
	}
	if !IsNil(o.ContainEnd) {
		toSerialize["containEnd"] = o.ContainEnd
	}
	if !IsNil(o.PriceUnit) {
		toSerialize["priceUnit"] = o.PriceUnit
	}
	if !IsNil(o.RangeBegin) {
		toSerialize["rangeBegin"] = o.RangeBegin
	}
	if !IsNil(o.RangeEnd) {
		toSerialize["rangeEnd"] = o.RangeEnd
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	return toSerialize, nil
}

type NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner struct {
	value *GetOrdersItemOrdersInnerTickerPriceDefineListInner
	isSet bool
}

func (v NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner) Get() *GetOrdersItemOrdersInnerTickerPriceDefineListInner {
	return v.value
}

func (v *NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner) Set(val *GetOrdersItemOrdersInnerTickerPriceDefineListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrdersItemOrdersInnerTickerPriceDefineListInner(val *GetOrdersItemOrdersInnerTickerPriceDefineListInner) *NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner {
	return &NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner{value: val, isSet: true}
}

func (v NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrdersItemOrdersInnerTickerPriceDefineListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


