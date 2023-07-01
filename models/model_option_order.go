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

// checks if the OptionOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionOrder{}

// OptionOrder struct for OptionOrder
type OptionOrder struct {
	Action *OrderSide `json:"action,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
	TickerType *string `json:"tickerType,omitempty"`
	TotalQuantity *int32 `json:"totalQuantity,omitempty"`
}

// NewOptionOrder instantiates a new OptionOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionOrder() *OptionOrder {
	this := OptionOrder{}
	var tickerType string = "OPTION"
	this.TickerType = &tickerType
	return &this
}

// NewOptionOrderWithDefaults instantiates a new OptionOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionOrderWithDefaults() *OptionOrder {
	this := OptionOrder{}
	var tickerType string = "OPTION"
	this.TickerType = &tickerType
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *OptionOrder) GetAction() OrderSide {
	if o == nil || IsNil(o.Action) {
		var ret OrderSide
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetActionOk() (*OrderSide, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *OptionOrder) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given OrderSide and assigns it to the Action field.
func (o *OptionOrder) SetAction(v OrderSide) {
	o.Action = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OptionOrder) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OptionOrder) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OptionOrder) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *OptionOrder) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *OptionOrder) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *OptionOrder) SetTickerId(v int32) {
	o.TickerId = &v
}

// GetTickerType returns the TickerType field value if set, zero value otherwise.
func (o *OptionOrder) GetTickerType() string {
	if o == nil || IsNil(o.TickerType) {
		var ret string
		return ret
	}
	return *o.TickerType
}

// GetTickerTypeOk returns a tuple with the TickerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetTickerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TickerType) {
		return nil, false
	}
	return o.TickerType, true
}

// HasTickerType returns a boolean if a field has been set.
func (o *OptionOrder) HasTickerType() bool {
	if o != nil && !IsNil(o.TickerType) {
		return true
	}

	return false
}

// SetTickerType gets a reference to the given string and assigns it to the TickerType field.
func (o *OptionOrder) SetTickerType(v string) {
	o.TickerType = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *OptionOrder) GetTotalQuantity() int32 {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetTotalQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *OptionOrder) HasTotalQuantity() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given int32 and assigns it to the TotalQuantity field.
func (o *OptionOrder) SetTotalQuantity(v int32) {
	o.TotalQuantity = &v
}

func (o OptionOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	if !IsNil(o.TickerType) {
		toSerialize["tickerType"] = o.TickerType
	}
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	return toSerialize, nil
}

type NullableOptionOrder struct {
	value *OptionOrder
	isSet bool
}

func (v NullableOptionOrder) Get() *OptionOrder {
	return v.value
}

func (v *NullableOptionOrder) Set(val *OptionOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionOrder(val *OptionOrder) *NullableOptionOrder {
	return &NullableOptionOrder{value: val, isSet: true}
}

func (v NullableOptionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


