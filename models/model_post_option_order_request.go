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

// checks if the PostOptionOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostOptionOrderRequest{}

// PostOptionOrderRequest struct for PostOptionOrderRequest
type PostOptionOrderRequest struct {
	AuxPrice *float32 `json:"auxPrice,omitempty"`
	LmtPrice *float32 `json:"lmtPrice,omitempty"`
	OrderType *OrderType `json:"orderType,omitempty"`
	Orders []OptionOrder `json:"orders,omitempty"`
	SerialId *string `json:"serialId,omitempty"`
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`
}

// NewPostOptionOrderRequest instantiates a new PostOptionOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostOptionOrderRequest() *PostOptionOrderRequest {
	this := PostOptionOrderRequest{}
	return &this
}

// NewPostOptionOrderRequestWithDefaults instantiates a new PostOptionOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostOptionOrderRequestWithDefaults() *PostOptionOrderRequest {
	this := PostOptionOrderRequest{}
	return &this
}

// GetAuxPrice returns the AuxPrice field value if set, zero value otherwise.
func (o *PostOptionOrderRequest) GetAuxPrice() float32 {
	if o == nil || IsNil(o.AuxPrice) {
		var ret float32
		return ret
	}
	return *o.AuxPrice
}

// GetAuxPriceOk returns a tuple with the AuxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOptionOrderRequest) GetAuxPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.AuxPrice) {
		return nil, false
	}
	return o.AuxPrice, true
}

// HasAuxPrice returns a boolean if a field has been set.
func (o *PostOptionOrderRequest) HasAuxPrice() bool {
	if o != nil && !IsNil(o.AuxPrice) {
		return true
	}

	return false
}

// SetAuxPrice gets a reference to the given float32 and assigns it to the AuxPrice field.
func (o *PostOptionOrderRequest) SetAuxPrice(v float32) {
	o.AuxPrice = &v
}

// GetLmtPrice returns the LmtPrice field value if set, zero value otherwise.
func (o *PostOptionOrderRequest) GetLmtPrice() float32 {
	if o == nil || IsNil(o.LmtPrice) {
		var ret float32
		return ret
	}
	return *o.LmtPrice
}

// GetLmtPriceOk returns a tuple with the LmtPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOptionOrderRequest) GetLmtPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.LmtPrice) {
		return nil, false
	}
	return o.LmtPrice, true
}

// HasLmtPrice returns a boolean if a field has been set.
func (o *PostOptionOrderRequest) HasLmtPrice() bool {
	if o != nil && !IsNil(o.LmtPrice) {
		return true
	}

	return false
}

// SetLmtPrice gets a reference to the given float32 and assigns it to the LmtPrice field.
func (o *PostOptionOrderRequest) SetLmtPrice(v float32) {
	o.LmtPrice = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *PostOptionOrderRequest) GetOrderType() OrderType {
	if o == nil || IsNil(o.OrderType) {
		var ret OrderType
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOptionOrderRequest) GetOrderTypeOk() (*OrderType, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *PostOptionOrderRequest) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given OrderType and assigns it to the OrderType field.
func (o *PostOptionOrderRequest) SetOrderType(v OrderType) {
	o.OrderType = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *PostOptionOrderRequest) GetOrders() []OptionOrder {
	if o == nil || IsNil(o.Orders) {
		var ret []OptionOrder
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOptionOrderRequest) GetOrdersOk() ([]OptionOrder, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *PostOptionOrderRequest) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OptionOrder and assigns it to the Orders field.
func (o *PostOptionOrderRequest) SetOrders(v []OptionOrder) {
	o.Orders = v
}

// GetSerialId returns the SerialId field value if set, zero value otherwise.
func (o *PostOptionOrderRequest) GetSerialId() string {
	if o == nil || IsNil(o.SerialId) {
		var ret string
		return ret
	}
	return *o.SerialId
}

// GetSerialIdOk returns a tuple with the SerialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOptionOrderRequest) GetSerialIdOk() (*string, bool) {
	if o == nil || IsNil(o.SerialId) {
		return nil, false
	}
	return o.SerialId, true
}

// HasSerialId returns a boolean if a field has been set.
func (o *PostOptionOrderRequest) HasSerialId() bool {
	if o != nil && !IsNil(o.SerialId) {
		return true
	}

	return false
}

// SetSerialId gets a reference to the given string and assigns it to the SerialId field.
func (o *PostOptionOrderRequest) SetSerialId(v string) {
	o.SerialId = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *PostOptionOrderRequest) GetTimeInForce() TimeInForce {
	if o == nil || IsNil(o.TimeInForce) {
		var ret TimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOptionOrderRequest) GetTimeInForceOk() (*TimeInForce, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *PostOptionOrderRequest) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given TimeInForce and assigns it to the TimeInForce field.
func (o *PostOptionOrderRequest) SetTimeInForce(v TimeInForce) {
	o.TimeInForce = &v
}

func (o PostOptionOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostOptionOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuxPrice) {
		toSerialize["auxPrice"] = o.AuxPrice
	}
	if !IsNil(o.LmtPrice) {
		toSerialize["lmtPrice"] = o.LmtPrice
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.SerialId) {
		toSerialize["serialId"] = o.SerialId
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	return toSerialize, nil
}

type NullablePostOptionOrderRequest struct {
	value *PostOptionOrderRequest
	isSet bool
}

func (v NullablePostOptionOrderRequest) Get() *PostOptionOrderRequest {
	return v.value
}

func (v *NullablePostOptionOrderRequest) Set(val *PostOptionOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostOptionOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostOptionOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostOptionOrderRequest(val *PostOptionOrderRequest) *NullablePostOptionOrderRequest {
	return &NullablePostOptionOrderRequest{value: val, isSet: true}
}

func (v NullablePostOptionOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostOptionOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


