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

// checks if the GetAccountResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountResponsePositionsInner{}

// GetAccountResponsePositionsInner struct for GetAccountResponsePositionsInner
type GetAccountResponsePositionsInner struct {
	AssetType *string `json:"assetType,omitempty"`
	BrokerId *int32 `json:"brokerId,omitempty"`
	Cost *string `json:"cost,omitempty"`
	CostPrice *string `json:"costPrice,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ExchangeRate *string `json:"exchangeRate,omitempty"`
	Id *int32 `json:"id,omitempty"`
	LastOpenTime *string `json:"lastOpenTime,omitempty"`
	LastPrice *string `json:"lastPrice,omitempty"`
	Lock *bool `json:"lock,omitempty"`
	MarketValue *string `json:"marketValue,omitempty"`
	Position *string `json:"position,omitempty"`
	PositionProportion *string `json:"positionProportion,omitempty"`
	Ticker *GetAcccountRequestPositionsInnerTicker `json:"ticker,omitempty"`
	UnrealizedProfitLoss *string `json:"unrealizedProfitLoss,omitempty"`
	UnrealizedProfitLossRate *string `json:"unrealizedProfitLossRate,omitempty"`
	UpdatePositionTimeStamp *int64 `json:"updatePositionTimeStamp,omitempty"`
}

// NewGetAccountResponsePositionsInner instantiates a new GetAccountResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountResponsePositionsInner() *GetAccountResponsePositionsInner {
	this := GetAccountResponsePositionsInner{}
	return &this
}

// NewGetAccountResponsePositionsInnerWithDefaults instantiates a new GetAccountResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountResponsePositionsInnerWithDefaults() *GetAccountResponsePositionsInner {
	this := GetAccountResponsePositionsInner{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetAssetType() string {
	if o == nil || IsNil(o.AssetType) {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetAssetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *GetAccountResponsePositionsInner) SetAssetType(v string) {
	o.AssetType = &v
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetBrokerId() int32 {
	if o == nil || IsNil(o.BrokerId) {
		var ret int32
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetBrokerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BrokerId) {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasBrokerId() bool {
	if o != nil && !IsNil(o.BrokerId) {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given int32 and assigns it to the BrokerId field.
func (o *GetAccountResponsePositionsInner) SetBrokerId(v int32) {
	o.BrokerId = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetCost() string {
	if o == nil || IsNil(o.Cost) {
		var ret string
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetCostOk() (*string, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given string and assigns it to the Cost field.
func (o *GetAccountResponsePositionsInner) SetCost(v string) {
	o.Cost = &v
}

// GetCostPrice returns the CostPrice field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetCostPrice() string {
	if o == nil || IsNil(o.CostPrice) {
		var ret string
		return ret
	}
	return *o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetCostPriceOk() (*string, bool) {
	if o == nil || IsNil(o.CostPrice) {
		return nil, false
	}
	return o.CostPrice, true
}

// HasCostPrice returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasCostPrice() bool {
	if o != nil && !IsNil(o.CostPrice) {
		return true
	}

	return false
}

// SetCostPrice gets a reference to the given string and assigns it to the CostPrice field.
func (o *GetAccountResponsePositionsInner) SetCostPrice(v string) {
	o.CostPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetAccountResponsePositionsInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetExchangeRate() string {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret string
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetExchangeRateOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given string and assigns it to the ExchangeRate field.
func (o *GetAccountResponsePositionsInner) SetExchangeRate(v string) {
	o.ExchangeRate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetAccountResponsePositionsInner) SetId(v int32) {
	o.Id = &v
}

// GetLastOpenTime returns the LastOpenTime field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetLastOpenTime() string {
	if o == nil || IsNil(o.LastOpenTime) {
		var ret string
		return ret
	}
	return *o.LastOpenTime
}

// GetLastOpenTimeOk returns a tuple with the LastOpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetLastOpenTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastOpenTime) {
		return nil, false
	}
	return o.LastOpenTime, true
}

// HasLastOpenTime returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasLastOpenTime() bool {
	if o != nil && !IsNil(o.LastOpenTime) {
		return true
	}

	return false
}

// SetLastOpenTime gets a reference to the given string and assigns it to the LastOpenTime field.
func (o *GetAccountResponsePositionsInner) SetLastOpenTime(v string) {
	o.LastOpenTime = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetLastPrice() string {
	if o == nil || IsNil(o.LastPrice) {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetLastPriceOk() (*string, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *GetAccountResponsePositionsInner) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetLock() bool {
	if o == nil || IsNil(o.Lock) {
		var ret bool
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetLockOk() (*bool, bool) {
	if o == nil || IsNil(o.Lock) {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasLock() bool {
	if o != nil && !IsNil(o.Lock) {
		return true
	}

	return false
}

// SetLock gets a reference to the given bool and assigns it to the Lock field.
func (o *GetAccountResponsePositionsInner) SetLock(v bool) {
	o.Lock = &v
}

// GetMarketValue returns the MarketValue field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetMarketValue() string {
	if o == nil || IsNil(o.MarketValue) {
		var ret string
		return ret
	}
	return *o.MarketValue
}

// GetMarketValueOk returns a tuple with the MarketValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetMarketValueOk() (*string, bool) {
	if o == nil || IsNil(o.MarketValue) {
		return nil, false
	}
	return o.MarketValue, true
}

// HasMarketValue returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasMarketValue() bool {
	if o != nil && !IsNil(o.MarketValue) {
		return true
	}

	return false
}

// SetMarketValue gets a reference to the given string and assigns it to the MarketValue field.
func (o *GetAccountResponsePositionsInner) SetMarketValue(v string) {
	o.MarketValue = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *GetAccountResponsePositionsInner) SetPosition(v string) {
	o.Position = &v
}

// GetPositionProportion returns the PositionProportion field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetPositionProportion() string {
	if o == nil || IsNil(o.PositionProportion) {
		var ret string
		return ret
	}
	return *o.PositionProportion
}

// GetPositionProportionOk returns a tuple with the PositionProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetPositionProportionOk() (*string, bool) {
	if o == nil || IsNil(o.PositionProportion) {
		return nil, false
	}
	return o.PositionProportion, true
}

// HasPositionProportion returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasPositionProportion() bool {
	if o != nil && !IsNil(o.PositionProportion) {
		return true
	}

	return false
}

// SetPositionProportion gets a reference to the given string and assigns it to the PositionProportion field.
func (o *GetAccountResponsePositionsInner) SetPositionProportion(v string) {
	o.PositionProportion = &v
}

// GetTicker returns the Ticker field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetTicker() GetAcccountRequestPositionsInnerTicker {
	if o == nil || IsNil(o.Ticker) {
		var ret GetAcccountRequestPositionsInnerTicker
		return ret
	}
	return *o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetTickerOk() (*GetAcccountRequestPositionsInnerTicker, bool) {
	if o == nil || IsNil(o.Ticker) {
		return nil, false
	}
	return o.Ticker, true
}

// HasTicker returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasTicker() bool {
	if o != nil && !IsNil(o.Ticker) {
		return true
	}

	return false
}

// SetTicker gets a reference to the given GetAcccountRequestPositionsInnerTicker and assigns it to the Ticker field.
func (o *GetAccountResponsePositionsInner) SetTicker(v GetAcccountRequestPositionsInnerTicker) {
	o.Ticker = &v
}

// GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLoss() string {
	if o == nil || IsNil(o.UnrealizedProfitLoss) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLoss
}

// GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLossOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLoss) {
		return nil, false
	}
	return o.UnrealizedProfitLoss, true
}

// HasUnrealizedProfitLoss returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasUnrealizedProfitLoss() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLoss) {
		return true
	}

	return false
}

// SetUnrealizedProfitLoss gets a reference to the given string and assigns it to the UnrealizedProfitLoss field.
func (o *GetAccountResponsePositionsInner) SetUnrealizedProfitLoss(v string) {
	o.UnrealizedProfitLoss = &v
}

// GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLossRate() string {
	if o == nil || IsNil(o.UnrealizedProfitLossRate) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLossRate
}

// GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetUnrealizedProfitLossRateOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLossRate) {
		return nil, false
	}
	return o.UnrealizedProfitLossRate, true
}

// HasUnrealizedProfitLossRate returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasUnrealizedProfitLossRate() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLossRate) {
		return true
	}

	return false
}

// SetUnrealizedProfitLossRate gets a reference to the given string and assigns it to the UnrealizedProfitLossRate field.
func (o *GetAccountResponsePositionsInner) SetUnrealizedProfitLossRate(v string) {
	o.UnrealizedProfitLossRate = &v
}

// GetUpdatePositionTimeStamp returns the UpdatePositionTimeStamp field value if set, zero value otherwise.
func (o *GetAccountResponsePositionsInner) GetUpdatePositionTimeStamp() int64 {
	if o == nil || IsNil(o.UpdatePositionTimeStamp) {
		var ret int64
		return ret
	}
	return *o.UpdatePositionTimeStamp
}

// GetUpdatePositionTimeStampOk returns a tuple with the UpdatePositionTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponsePositionsInner) GetUpdatePositionTimeStampOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatePositionTimeStamp) {
		return nil, false
	}
	return o.UpdatePositionTimeStamp, true
}

// HasUpdatePositionTimeStamp returns a boolean if a field has been set.
func (o *GetAccountResponsePositionsInner) HasUpdatePositionTimeStamp() bool {
	if o != nil && !IsNil(o.UpdatePositionTimeStamp) {
		return true
	}

	return false
}

// SetUpdatePositionTimeStamp gets a reference to the given int64 and assigns it to the UpdatePositionTimeStamp field.
func (o *GetAccountResponsePositionsInner) SetUpdatePositionTimeStamp(v int64) {
	o.UpdatePositionTimeStamp = &v
}

func (o GetAccountResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountResponsePositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetType) {
		toSerialize["assetType"] = o.AssetType
	}
	if !IsNil(o.BrokerId) {
		toSerialize["brokerId"] = o.BrokerId
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.CostPrice) {
		toSerialize["costPrice"] = o.CostPrice
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastOpenTime) {
		toSerialize["lastOpenTime"] = o.LastOpenTime
	}
	if !IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !IsNil(o.Lock) {
		toSerialize["lock"] = o.Lock
	}
	if !IsNil(o.MarketValue) {
		toSerialize["marketValue"] = o.MarketValue
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.PositionProportion) {
		toSerialize["positionProportion"] = o.PositionProportion
	}
	if !IsNil(o.Ticker) {
		toSerialize["ticker"] = o.Ticker
	}
	if !IsNil(o.UnrealizedProfitLoss) {
		toSerialize["unrealizedProfitLoss"] = o.UnrealizedProfitLoss
	}
	if !IsNil(o.UnrealizedProfitLossRate) {
		toSerialize["unrealizedProfitLossRate"] = o.UnrealizedProfitLossRate
	}
	if !IsNil(o.UpdatePositionTimeStamp) {
		toSerialize["updatePositionTimeStamp"] = o.UpdatePositionTimeStamp
	}
	return toSerialize, nil
}

type NullableGetAccountResponsePositionsInner struct {
	value *GetAccountResponsePositionsInner
	isSet bool
}

func (v NullableGetAccountResponsePositionsInner) Get() *GetAccountResponsePositionsInner {
	return v.value
}

func (v *NullableGetAccountResponsePositionsInner) Set(val *GetAccountResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountResponsePositionsInner(val *GetAccountResponsePositionsInner) *NullableGetAccountResponsePositionsInner {
	return &NullableGetAccountResponsePositionsInner{value: val, isSet: true}
}

func (v NullableGetAccountResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


