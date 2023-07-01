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

// checks if the GetAlertsResponseDataInnerEventWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlertsResponseDataInnerEventWarning{}

// GetAlertsResponseDataInnerEventWarning struct for GetAlertsResponseDataInnerEventWarning
type GetAlertsResponseDataInnerEventWarning struct {
	Del *bool `json:"del,omitempty"`
	ExchangeId *int32 `json:"exchangeId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Remove *bool `json:"remove,omitempty"`
	Rules []GetAlertsResponseDataInnerEventWarningRulesInner `json:"rules,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
	TickerType *int32 `json:"tickerType,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
}

// NewGetAlertsResponseDataInnerEventWarning instantiates a new GetAlertsResponseDataInnerEventWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlertsResponseDataInnerEventWarning() *GetAlertsResponseDataInnerEventWarning {
	this := GetAlertsResponseDataInnerEventWarning{}
	return &this
}

// NewGetAlertsResponseDataInnerEventWarningWithDefaults instantiates a new GetAlertsResponseDataInnerEventWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlertsResponseDataInnerEventWarningWithDefaults() *GetAlertsResponseDataInnerEventWarning {
	this := GetAlertsResponseDataInnerEventWarning{}
	return &this
}

// GetDel returns the Del field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetDel() bool {
	if o == nil || IsNil(o.Del) {
		var ret bool
		return ret
	}
	return *o.Del
}

// GetDelOk returns a tuple with the Del field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetDelOk() (*bool, bool) {
	if o == nil || IsNil(o.Del) {
		return nil, false
	}
	return o.Del, true
}

// HasDel returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasDel() bool {
	if o != nil && !IsNil(o.Del) {
		return true
	}

	return false
}

// SetDel gets a reference to the given bool and assigns it to the Del field.
func (o *GetAlertsResponseDataInnerEventWarning) SetDel(v bool) {
	o.Del = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetExchangeId() int32 {
	if o == nil || IsNil(o.ExchangeId) {
		var ret int32
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetExchangeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ExchangeId) {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasExchangeId() bool {
	if o != nil && !IsNil(o.ExchangeId) {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given int32 and assigns it to the ExchangeId field.
func (o *GetAlertsResponseDataInnerEventWarning) SetExchangeId(v int32) {
	o.ExchangeId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetAlertsResponseDataInnerEventWarning) SetId(v int32) {
	o.Id = &v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetRemove() bool {
	if o == nil || IsNil(o.Remove) {
		var ret bool
		return ret
	}
	return *o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetRemoveOk() (*bool, bool) {
	if o == nil || IsNil(o.Remove) {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasRemove() bool {
	if o != nil && !IsNil(o.Remove) {
		return true
	}

	return false
}

// SetRemove gets a reference to the given bool and assigns it to the Remove field.
func (o *GetAlertsResponseDataInnerEventWarning) SetRemove(v bool) {
	o.Remove = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetRules() []GetAlertsResponseDataInnerEventWarningRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetAlertsResponseDataInnerEventWarningRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetRulesOk() ([]GetAlertsResponseDataInnerEventWarningRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetAlertsResponseDataInnerEventWarningRulesInner and assigns it to the Rules field.
func (o *GetAlertsResponseDataInnerEventWarning) SetRules(v []GetAlertsResponseDataInnerEventWarningRulesInner) {
	o.Rules = v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *GetAlertsResponseDataInnerEventWarning) SetTickerId(v int32) {
	o.TickerId = &v
}

// GetTickerType returns the TickerType field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetTickerType() int32 {
	if o == nil || IsNil(o.TickerType) {
		var ret int32
		return ret
	}
	return *o.TickerType
}

// GetTickerTypeOk returns a tuple with the TickerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetTickerTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerType) {
		return nil, false
	}
	return o.TickerType, true
}

// HasTickerType returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasTickerType() bool {
	if o != nil && !IsNil(o.TickerType) {
		return true
	}

	return false
}

// SetTickerType gets a reference to the given int32 and assigns it to the TickerType field.
func (o *GetAlertsResponseDataInnerEventWarning) SetTickerType(v int32) {
	o.TickerType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInnerEventWarning) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInnerEventWarning) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInnerEventWarning) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *GetAlertsResponseDataInnerEventWarning) SetUserId(v int32) {
	o.UserId = &v
}

func (o GetAlertsResponseDataInnerEventWarning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlertsResponseDataInnerEventWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Del) {
		toSerialize["del"] = o.Del
	}
	if !IsNil(o.ExchangeId) {
		toSerialize["exchangeId"] = o.ExchangeId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Remove) {
		toSerialize["remove"] = o.Remove
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	if !IsNil(o.TickerType) {
		toSerialize["tickerType"] = o.TickerType
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableGetAlertsResponseDataInnerEventWarning struct {
	value *GetAlertsResponseDataInnerEventWarning
	isSet bool
}

func (v NullableGetAlertsResponseDataInnerEventWarning) Get() *GetAlertsResponseDataInnerEventWarning {
	return v.value
}

func (v *NullableGetAlertsResponseDataInnerEventWarning) Set(val *GetAlertsResponseDataInnerEventWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlertsResponseDataInnerEventWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlertsResponseDataInnerEventWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlertsResponseDataInnerEventWarning(val *GetAlertsResponseDataInnerEventWarning) *NullableGetAlertsResponseDataInnerEventWarning {
	return &NullableGetAlertsResponseDataInnerEventWarning{value: val, isSet: true}
}

func (v NullableGetAlertsResponseDataInnerEventWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlertsResponseDataInnerEventWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


