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

// checks if the ReplaceOptionOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceOptionOrderRequest{}

// ReplaceOptionOrderRequest struct for ReplaceOptionOrderRequest
type ReplaceOptionOrderRequest struct {
	AuxPrice *float32 `json:"auxPrice,omitempty"`
	ComboId *string `json:"comboId,omitempty"`
	LmtPrice *float32 `json:"lmtPrice,omitempty"`
	OrderType *OrderType `json:"orderType,omitempty"`
	Orders []OptionOrder `json:"orders,omitempty"`
	SerialId *string `json:"serialId,omitempty"`
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`
}

// NewReplaceOptionOrderRequest instantiates a new ReplaceOptionOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceOptionOrderRequest() *ReplaceOptionOrderRequest {
	this := ReplaceOptionOrderRequest{}
	return &this
}

// NewReplaceOptionOrderRequestWithDefaults instantiates a new ReplaceOptionOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceOptionOrderRequestWithDefaults() *ReplaceOptionOrderRequest {
	this := ReplaceOptionOrderRequest{}
	return &this
}

// GetAuxPrice returns the AuxPrice field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetAuxPrice() float32 {
	if o == nil || IsNil(o.AuxPrice) {
		var ret float32
		return ret
	}
	return *o.AuxPrice
}

// GetAuxPriceOk returns a tuple with the AuxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetAuxPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.AuxPrice) {
		return nil, false
	}
	return o.AuxPrice, true
}

// HasAuxPrice returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasAuxPrice() bool {
	if o != nil && !IsNil(o.AuxPrice) {
		return true
	}

	return false
}

// SetAuxPrice gets a reference to the given float32 and assigns it to the AuxPrice field.
func (o *ReplaceOptionOrderRequest) SetAuxPrice(v float32) {
	o.AuxPrice = &v
}

// GetComboId returns the ComboId field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetComboId() string {
	if o == nil || IsNil(o.ComboId) {
		var ret string
		return ret
	}
	return *o.ComboId
}

// GetComboIdOk returns a tuple with the ComboId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetComboIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComboId) {
		return nil, false
	}
	return o.ComboId, true
}

// HasComboId returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasComboId() bool {
	if o != nil && !IsNil(o.ComboId) {
		return true
	}

	return false
}

// SetComboId gets a reference to the given string and assigns it to the ComboId field.
func (o *ReplaceOptionOrderRequest) SetComboId(v string) {
	o.ComboId = &v
}

// GetLmtPrice returns the LmtPrice field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetLmtPrice() float32 {
	if o == nil || IsNil(o.LmtPrice) {
		var ret float32
		return ret
	}
	return *o.LmtPrice
}

// GetLmtPriceOk returns a tuple with the LmtPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetLmtPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.LmtPrice) {
		return nil, false
	}
	return o.LmtPrice, true
}

// HasLmtPrice returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasLmtPrice() bool {
	if o != nil && !IsNil(o.LmtPrice) {
		return true
	}

	return false
}

// SetLmtPrice gets a reference to the given float32 and assigns it to the LmtPrice field.
func (o *ReplaceOptionOrderRequest) SetLmtPrice(v float32) {
	o.LmtPrice = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetOrderType() OrderType {
	if o == nil || IsNil(o.OrderType) {
		var ret OrderType
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetOrderTypeOk() (*OrderType, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given OrderType and assigns it to the OrderType field.
func (o *ReplaceOptionOrderRequest) SetOrderType(v OrderType) {
	o.OrderType = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetOrders() []OptionOrder {
	if o == nil || IsNil(o.Orders) {
		var ret []OptionOrder
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetOrdersOk() ([]OptionOrder, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OptionOrder and assigns it to the Orders field.
func (o *ReplaceOptionOrderRequest) SetOrders(v []OptionOrder) {
	o.Orders = v
}

// GetSerialId returns the SerialId field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetSerialId() string {
	if o == nil || IsNil(o.SerialId) {
		var ret string
		return ret
	}
	return *o.SerialId
}

// GetSerialIdOk returns a tuple with the SerialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetSerialIdOk() (*string, bool) {
	if o == nil || IsNil(o.SerialId) {
		return nil, false
	}
	return o.SerialId, true
}

// HasSerialId returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasSerialId() bool {
	if o != nil && !IsNil(o.SerialId) {
		return true
	}

	return false
}

// SetSerialId gets a reference to the given string and assigns it to the SerialId field.
func (o *ReplaceOptionOrderRequest) SetSerialId(v string) {
	o.SerialId = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *ReplaceOptionOrderRequest) GetTimeInForce() TimeInForce {
	if o == nil || IsNil(o.TimeInForce) {
		var ret TimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceOptionOrderRequest) GetTimeInForceOk() (*TimeInForce, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *ReplaceOptionOrderRequest) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given TimeInForce and assigns it to the TimeInForce field.
func (o *ReplaceOptionOrderRequest) SetTimeInForce(v TimeInForce) {
	o.TimeInForce = &v
}

func (o ReplaceOptionOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceOptionOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuxPrice) {
		toSerialize["auxPrice"] = o.AuxPrice
	}
	if !IsNil(o.ComboId) {
		toSerialize["comboId"] = o.ComboId
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

type NullableReplaceOptionOrderRequest struct {
	value *ReplaceOptionOrderRequest
	isSet bool
}

func (v NullableReplaceOptionOrderRequest) Get() *ReplaceOptionOrderRequest {
	return v.value
}

func (v *NullableReplaceOptionOrderRequest) Set(val *ReplaceOptionOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceOptionOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceOptionOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceOptionOrderRequest(val *ReplaceOptionOrderRequest) *NullableReplaceOptionOrderRequest {
	return &NullableReplaceOptionOrderRequest{value: val, isSet: true}
}

func (v NullableReplaceOptionOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceOptionOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


