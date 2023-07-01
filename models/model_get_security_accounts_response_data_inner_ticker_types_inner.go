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

// checks if the GetSecurityAccountsResponseDataInnerTickerTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSecurityAccountsResponseDataInnerTickerTypesInner{}

// GetSecurityAccountsResponseDataInnerTickerTypesInner struct for GetSecurityAccountsResponseDataInnerTickerTypesInner
type GetSecurityAccountsResponseDataInnerTickerTypesInner struct {
	OrderTypes []OrderType `json:"orderTypes,omitempty"`
	RegionId *int32 `json:"regionId,omitempty"`
	TickerType *int32 `json:"tickerType,omitempty"`
}

// NewGetSecurityAccountsResponseDataInnerTickerTypesInner instantiates a new GetSecurityAccountsResponseDataInnerTickerTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSecurityAccountsResponseDataInnerTickerTypesInner() *GetSecurityAccountsResponseDataInnerTickerTypesInner {
	this := GetSecurityAccountsResponseDataInnerTickerTypesInner{}
	return &this
}

// NewGetSecurityAccountsResponseDataInnerTickerTypesInnerWithDefaults instantiates a new GetSecurityAccountsResponseDataInnerTickerTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSecurityAccountsResponseDataInnerTickerTypesInnerWithDefaults() *GetSecurityAccountsResponseDataInnerTickerTypesInner {
	this := GetSecurityAccountsResponseDataInnerTickerTypesInner{}
	return &this
}

// GetOrderTypes returns the OrderTypes field value if set, zero value otherwise.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) GetOrderTypes() []OrderType {
	if o == nil || IsNil(o.OrderTypes) {
		var ret []OrderType
		return ret
	}
	return o.OrderTypes
}

// GetOrderTypesOk returns a tuple with the OrderTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) GetOrderTypesOk() ([]OrderType, bool) {
	if o == nil || IsNil(o.OrderTypes) {
		return nil, false
	}
	return o.OrderTypes, true
}

// HasOrderTypes returns a boolean if a field has been set.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) HasOrderTypes() bool {
	if o != nil && !IsNil(o.OrderTypes) {
		return true
	}

	return false
}

// SetOrderTypes gets a reference to the given []OrderType and assigns it to the OrderTypes field.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) SetOrderTypes(v []OrderType) {
	o.OrderTypes = v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) GetRegionId() int32 {
	if o == nil || IsNil(o.RegionId) {
		var ret int32
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) GetRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int32 and assigns it to the RegionId field.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) SetRegionId(v int32) {
	o.RegionId = &v
}

// GetTickerType returns the TickerType field value if set, zero value otherwise.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) GetTickerType() int32 {
	if o == nil || IsNil(o.TickerType) {
		var ret int32
		return ret
	}
	return *o.TickerType
}

// GetTickerTypeOk returns a tuple with the TickerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) GetTickerTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerType) {
		return nil, false
	}
	return o.TickerType, true
}

// HasTickerType returns a boolean if a field has been set.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) HasTickerType() bool {
	if o != nil && !IsNil(o.TickerType) {
		return true
	}

	return false
}

// SetTickerType gets a reference to the given int32 and assigns it to the TickerType field.
func (o *GetSecurityAccountsResponseDataInnerTickerTypesInner) SetTickerType(v int32) {
	o.TickerType = &v
}

func (o GetSecurityAccountsResponseDataInnerTickerTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSecurityAccountsResponseDataInnerTickerTypesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderTypes) {
		toSerialize["orderTypes"] = o.OrderTypes
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !IsNil(o.TickerType) {
		toSerialize["tickerType"] = o.TickerType
	}
	return toSerialize, nil
}

type NullableGetSecurityAccountsResponseDataInnerTickerTypesInner struct {
	value *GetSecurityAccountsResponseDataInnerTickerTypesInner
	isSet bool
}

func (v NullableGetSecurityAccountsResponseDataInnerTickerTypesInner) Get() *GetSecurityAccountsResponseDataInnerTickerTypesInner {
	return v.value
}

func (v *NullableGetSecurityAccountsResponseDataInnerTickerTypesInner) Set(val *GetSecurityAccountsResponseDataInnerTickerTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSecurityAccountsResponseDataInnerTickerTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSecurityAccountsResponseDataInnerTickerTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSecurityAccountsResponseDataInnerTickerTypesInner(val *GetSecurityAccountsResponseDataInnerTickerTypesInner) *NullableGetSecurityAccountsResponseDataInnerTickerTypesInner {
	return &NullableGetSecurityAccountsResponseDataInnerTickerTypesInner{value: val, isSet: true}
}

func (v NullableGetSecurityAccountsResponseDataInnerTickerTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSecurityAccountsResponseDataInnerTickerTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

