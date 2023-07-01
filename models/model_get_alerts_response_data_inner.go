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

// checks if the GetAlertsResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlertsResponseDataInner{}

// GetAlertsResponseDataInner struct for GetAlertsResponseDataInner
type GetAlertsResponseDataInner struct {
	DisExchangeCode *string `json:"disExchangeCode,omitempty"`
	DisSymbol *string `json:"disSymbol,omitempty"`
	EventWarning *GetAlertsResponseDataInnerEventWarning `json:"eventWarning,omitempty"`
	ExchangeCode *string `json:"exchangeCode,omitempty"`
	ExchangeTrade *bool `json:"exchangeTrade,omitempty"`
	RegionId *int32 `json:"regionId,omitempty"`
	ShowCode *string `json:"showCode,omitempty"`
	TickerId *int32 `json:"tickerId,omitempty"`
	TickerName *string `json:"tickerName,omitempty"`
	TickerSymbol *string `json:"tickerSymbol,omitempty"`
	TickerType *int32 `json:"tickerType,omitempty"`
	TickerWarning *GetAlertsResponseDataInnerTickerWarning `json:"tickerWarning,omitempty"`
	TinyName *string `json:"tinyName,omitempty"`
}

// NewGetAlertsResponseDataInner instantiates a new GetAlertsResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlertsResponseDataInner() *GetAlertsResponseDataInner {
	this := GetAlertsResponseDataInner{}
	return &this
}

// NewGetAlertsResponseDataInnerWithDefaults instantiates a new GetAlertsResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlertsResponseDataInnerWithDefaults() *GetAlertsResponseDataInner {
	this := GetAlertsResponseDataInner{}
	return &this
}

// GetDisExchangeCode returns the DisExchangeCode field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetDisExchangeCode() string {
	if o == nil || IsNil(o.DisExchangeCode) {
		var ret string
		return ret
	}
	return *o.DisExchangeCode
}

// GetDisExchangeCodeOk returns a tuple with the DisExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetDisExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisExchangeCode) {
		return nil, false
	}
	return o.DisExchangeCode, true
}

// HasDisExchangeCode returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasDisExchangeCode() bool {
	if o != nil && !IsNil(o.DisExchangeCode) {
		return true
	}

	return false
}

// SetDisExchangeCode gets a reference to the given string and assigns it to the DisExchangeCode field.
func (o *GetAlertsResponseDataInner) SetDisExchangeCode(v string) {
	o.DisExchangeCode = &v
}

// GetDisSymbol returns the DisSymbol field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetDisSymbol() string {
	if o == nil || IsNil(o.DisSymbol) {
		var ret string
		return ret
	}
	return *o.DisSymbol
}

// GetDisSymbolOk returns a tuple with the DisSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetDisSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.DisSymbol) {
		return nil, false
	}
	return o.DisSymbol, true
}

// HasDisSymbol returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasDisSymbol() bool {
	if o != nil && !IsNil(o.DisSymbol) {
		return true
	}

	return false
}

// SetDisSymbol gets a reference to the given string and assigns it to the DisSymbol field.
func (o *GetAlertsResponseDataInner) SetDisSymbol(v string) {
	o.DisSymbol = &v
}

// GetEventWarning returns the EventWarning field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetEventWarning() GetAlertsResponseDataInnerEventWarning {
	if o == nil || IsNil(o.EventWarning) {
		var ret GetAlertsResponseDataInnerEventWarning
		return ret
	}
	return *o.EventWarning
}

// GetEventWarningOk returns a tuple with the EventWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetEventWarningOk() (*GetAlertsResponseDataInnerEventWarning, bool) {
	if o == nil || IsNil(o.EventWarning) {
		return nil, false
	}
	return o.EventWarning, true
}

// HasEventWarning returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasEventWarning() bool {
	if o != nil && !IsNil(o.EventWarning) {
		return true
	}

	return false
}

// SetEventWarning gets a reference to the given GetAlertsResponseDataInnerEventWarning and assigns it to the EventWarning field.
func (o *GetAlertsResponseDataInner) SetEventWarning(v GetAlertsResponseDataInnerEventWarning) {
	o.EventWarning = &v
}

// GetExchangeCode returns the ExchangeCode field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetExchangeCode() string {
	if o == nil || IsNil(o.ExchangeCode) {
		var ret string
		return ret
	}
	return *o.ExchangeCode
}

// GetExchangeCodeOk returns a tuple with the ExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeCode) {
		return nil, false
	}
	return o.ExchangeCode, true
}

// HasExchangeCode returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasExchangeCode() bool {
	if o != nil && !IsNil(o.ExchangeCode) {
		return true
	}

	return false
}

// SetExchangeCode gets a reference to the given string and assigns it to the ExchangeCode field.
func (o *GetAlertsResponseDataInner) SetExchangeCode(v string) {
	o.ExchangeCode = &v
}

// GetExchangeTrade returns the ExchangeTrade field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetExchangeTrade() bool {
	if o == nil || IsNil(o.ExchangeTrade) {
		var ret bool
		return ret
	}
	return *o.ExchangeTrade
}

// GetExchangeTradeOk returns a tuple with the ExchangeTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetExchangeTradeOk() (*bool, bool) {
	if o == nil || IsNil(o.ExchangeTrade) {
		return nil, false
	}
	return o.ExchangeTrade, true
}

// HasExchangeTrade returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasExchangeTrade() bool {
	if o != nil && !IsNil(o.ExchangeTrade) {
		return true
	}

	return false
}

// SetExchangeTrade gets a reference to the given bool and assigns it to the ExchangeTrade field.
func (o *GetAlertsResponseDataInner) SetExchangeTrade(v bool) {
	o.ExchangeTrade = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetRegionId() int32 {
	if o == nil || IsNil(o.RegionId) {
		var ret int32
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int32 and assigns it to the RegionId field.
func (o *GetAlertsResponseDataInner) SetRegionId(v int32) {
	o.RegionId = &v
}

// GetShowCode returns the ShowCode field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetShowCode() string {
	if o == nil || IsNil(o.ShowCode) {
		var ret string
		return ret
	}
	return *o.ShowCode
}

// GetShowCodeOk returns a tuple with the ShowCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetShowCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShowCode) {
		return nil, false
	}
	return o.ShowCode, true
}

// HasShowCode returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasShowCode() bool {
	if o != nil && !IsNil(o.ShowCode) {
		return true
	}

	return false
}

// SetShowCode gets a reference to the given string and assigns it to the ShowCode field.
func (o *GetAlertsResponseDataInner) SetShowCode(v string) {
	o.ShowCode = &v
}

// GetTickerId returns the TickerId field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetTickerId() int32 {
	if o == nil || IsNil(o.TickerId) {
		var ret int32
		return ret
	}
	return *o.TickerId
}

// GetTickerIdOk returns a tuple with the TickerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetTickerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerId) {
		return nil, false
	}
	return o.TickerId, true
}

// HasTickerId returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasTickerId() bool {
	if o != nil && !IsNil(o.TickerId) {
		return true
	}

	return false
}

// SetTickerId gets a reference to the given int32 and assigns it to the TickerId field.
func (o *GetAlertsResponseDataInner) SetTickerId(v int32) {
	o.TickerId = &v
}

// GetTickerName returns the TickerName field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetTickerName() string {
	if o == nil || IsNil(o.TickerName) {
		var ret string
		return ret
	}
	return *o.TickerName
}

// GetTickerNameOk returns a tuple with the TickerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetTickerNameOk() (*string, bool) {
	if o == nil || IsNil(o.TickerName) {
		return nil, false
	}
	return o.TickerName, true
}

// HasTickerName returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasTickerName() bool {
	if o != nil && !IsNil(o.TickerName) {
		return true
	}

	return false
}

// SetTickerName gets a reference to the given string and assigns it to the TickerName field.
func (o *GetAlertsResponseDataInner) SetTickerName(v string) {
	o.TickerName = &v
}

// GetTickerSymbol returns the TickerSymbol field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetTickerSymbol() string {
	if o == nil || IsNil(o.TickerSymbol) {
		var ret string
		return ret
	}
	return *o.TickerSymbol
}

// GetTickerSymbolOk returns a tuple with the TickerSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetTickerSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.TickerSymbol) {
		return nil, false
	}
	return o.TickerSymbol, true
}

// HasTickerSymbol returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasTickerSymbol() bool {
	if o != nil && !IsNil(o.TickerSymbol) {
		return true
	}

	return false
}

// SetTickerSymbol gets a reference to the given string and assigns it to the TickerSymbol field.
func (o *GetAlertsResponseDataInner) SetTickerSymbol(v string) {
	o.TickerSymbol = &v
}

// GetTickerType returns the TickerType field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetTickerType() int32 {
	if o == nil || IsNil(o.TickerType) {
		var ret int32
		return ret
	}
	return *o.TickerType
}

// GetTickerTypeOk returns a tuple with the TickerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetTickerTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TickerType) {
		return nil, false
	}
	return o.TickerType, true
}

// HasTickerType returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasTickerType() bool {
	if o != nil && !IsNil(o.TickerType) {
		return true
	}

	return false
}

// SetTickerType gets a reference to the given int32 and assigns it to the TickerType field.
func (o *GetAlertsResponseDataInner) SetTickerType(v int32) {
	o.TickerType = &v
}

// GetTickerWarning returns the TickerWarning field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetTickerWarning() GetAlertsResponseDataInnerTickerWarning {
	if o == nil || IsNil(o.TickerWarning) {
		var ret GetAlertsResponseDataInnerTickerWarning
		return ret
	}
	return *o.TickerWarning
}

// GetTickerWarningOk returns a tuple with the TickerWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetTickerWarningOk() (*GetAlertsResponseDataInnerTickerWarning, bool) {
	if o == nil || IsNil(o.TickerWarning) {
		return nil, false
	}
	return o.TickerWarning, true
}

// HasTickerWarning returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasTickerWarning() bool {
	if o != nil && !IsNil(o.TickerWarning) {
		return true
	}

	return false
}

// SetTickerWarning gets a reference to the given GetAlertsResponseDataInnerTickerWarning and assigns it to the TickerWarning field.
func (o *GetAlertsResponseDataInner) SetTickerWarning(v GetAlertsResponseDataInnerTickerWarning) {
	o.TickerWarning = &v
}

// GetTinyName returns the TinyName field value if set, zero value otherwise.
func (o *GetAlertsResponseDataInner) GetTinyName() string {
	if o == nil || IsNil(o.TinyName) {
		var ret string
		return ret
	}
	return *o.TinyName
}

// GetTinyNameOk returns a tuple with the TinyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertsResponseDataInner) GetTinyNameOk() (*string, bool) {
	if o == nil || IsNil(o.TinyName) {
		return nil, false
	}
	return o.TinyName, true
}

// HasTinyName returns a boolean if a field has been set.
func (o *GetAlertsResponseDataInner) HasTinyName() bool {
	if o != nil && !IsNil(o.TinyName) {
		return true
	}

	return false
}

// SetTinyName gets a reference to the given string and assigns it to the TinyName field.
func (o *GetAlertsResponseDataInner) SetTinyName(v string) {
	o.TinyName = &v
}

func (o GetAlertsResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlertsResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisExchangeCode) {
		toSerialize["disExchangeCode"] = o.DisExchangeCode
	}
	if !IsNil(o.DisSymbol) {
		toSerialize["disSymbol"] = o.DisSymbol
	}
	if !IsNil(o.EventWarning) {
		toSerialize["eventWarning"] = o.EventWarning
	}
	if !IsNil(o.ExchangeCode) {
		toSerialize["exchangeCode"] = o.ExchangeCode
	}
	if !IsNil(o.ExchangeTrade) {
		toSerialize["exchangeTrade"] = o.ExchangeTrade
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !IsNil(o.ShowCode) {
		toSerialize["showCode"] = o.ShowCode
	}
	if !IsNil(o.TickerId) {
		toSerialize["tickerId"] = o.TickerId
	}
	if !IsNil(o.TickerName) {
		toSerialize["tickerName"] = o.TickerName
	}
	if !IsNil(o.TickerSymbol) {
		toSerialize["tickerSymbol"] = o.TickerSymbol
	}
	if !IsNil(o.TickerType) {
		toSerialize["tickerType"] = o.TickerType
	}
	if !IsNil(o.TickerWarning) {
		toSerialize["tickerWarning"] = o.TickerWarning
	}
	if !IsNil(o.TinyName) {
		toSerialize["tinyName"] = o.TinyName
	}
	return toSerialize, nil
}

type NullableGetAlertsResponseDataInner struct {
	value *GetAlertsResponseDataInner
	isSet bool
}

func (v NullableGetAlertsResponseDataInner) Get() *GetAlertsResponseDataInner {
	return v.value
}

func (v *NullableGetAlertsResponseDataInner) Set(val *GetAlertsResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlertsResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlertsResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlertsResponseDataInner(val *GetAlertsResponseDataInner) *NullableGetAlertsResponseDataInner {
	return &NullableGetAlertsResponseDataInner{value: val, isSet: true}
}

func (v NullableGetAlertsResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlertsResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


