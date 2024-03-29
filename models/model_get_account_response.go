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

// checks if the GetAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountResponse{}

// GetAccountResponse struct for GetAccountResponse
type GetAccountResponse struct {
	AccountMembers []GetAcccountRequestAccountMembersInner `json:"accountMembers,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	Banners []GetAcccountRequestBannersInner `json:"banners,omitempty"`
	BrokerAccountId *string `json:"brokerAccountId,omitempty"`
	BrokerId *int32 `json:"brokerId,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrencyId *int32 `json:"currencyId,omitempty"`
	Deposits []map[string]interface{} `json:"deposits,omitempty"`
	NetLiquidation *string `json:"netLiquidation,omitempty"`
	OpenIpoOrders []map[string]interface{} `json:"openIpoOrders,omitempty"`
	OpenOrderSize *int32 `json:"openOrderSize,omitempty"`
	OpenOrders []map[string]interface{} `json:"openOrders,omitempty"`
	Pdt *bool `json:"pdt,omitempty"`
	Positions []GetAccountResponsePositionsInner `json:"positions,omitempty"`
	Professional *bool `json:"professional,omitempty"`
	SecAccountId *int32 `json:"secAccountId,omitempty"`
	ShowUpgrade *bool `json:"showUpgrade,omitempty"`
	TotalCost *string `json:"totalCost,omitempty"`
	UnrealizedProfitLoss *string `json:"unrealizedProfitLoss,omitempty"`
	UnrealizedProfitLossBase *float32 `json:"unrealizedProfitLossBase,omitempty"`
	UnrealizedProfitLossRate *string `json:"unrealizedProfitLossRate,omitempty"`
	Warning *bool `json:"warning,omitempty"`
}

// NewGetAccountResponse instantiates a new GetAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountResponse() *GetAccountResponse {
	this := GetAccountResponse{}
	return &this
}

// NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountResponseWithDefaults() *GetAccountResponse {
	this := GetAccountResponse{}
	return &this
}

// GetAccountMembers returns the AccountMembers field value if set, zero value otherwise.
func (o *GetAccountResponse) GetAccountMembers() []GetAcccountRequestAccountMembersInner {
	if o == nil || IsNil(o.AccountMembers) {
		var ret []GetAcccountRequestAccountMembersInner
		return ret
	}
	return o.AccountMembers
}

// GetAccountMembersOk returns a tuple with the AccountMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetAccountMembersOk() ([]GetAcccountRequestAccountMembersInner, bool) {
	if o == nil || IsNil(o.AccountMembers) {
		return nil, false
	}
	return o.AccountMembers, true
}

// HasAccountMembers returns a boolean if a field has been set.
func (o *GetAccountResponse) HasAccountMembers() bool {
	if o != nil && !IsNil(o.AccountMembers) {
		return true
	}

	return false
}

// SetAccountMembers gets a reference to the given []GetAcccountRequestAccountMembersInner and assigns it to the AccountMembers field.
func (o *GetAccountResponse) SetAccountMembers(v []GetAcccountRequestAccountMembersInner) {
	o.AccountMembers = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *GetAccountResponse) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *GetAccountResponse) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *GetAccountResponse) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBanners returns the Banners field value if set, zero value otherwise.
func (o *GetAccountResponse) GetBanners() []GetAcccountRequestBannersInner {
	if o == nil || IsNil(o.Banners) {
		var ret []GetAcccountRequestBannersInner
		return ret
	}
	return o.Banners
}

// GetBannersOk returns a tuple with the Banners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBannersOk() ([]GetAcccountRequestBannersInner, bool) {
	if o == nil || IsNil(o.Banners) {
		return nil, false
	}
	return o.Banners, true
}

// HasBanners returns a boolean if a field has been set.
func (o *GetAccountResponse) HasBanners() bool {
	if o != nil && !IsNil(o.Banners) {
		return true
	}

	return false
}

// SetBanners gets a reference to the given []GetAcccountRequestBannersInner and assigns it to the Banners field.
func (o *GetAccountResponse) SetBanners(v []GetAcccountRequestBannersInner) {
	o.Banners = v
}

// GetBrokerAccountId returns the BrokerAccountId field value if set, zero value otherwise.
func (o *GetAccountResponse) GetBrokerAccountId() string {
	if o == nil || IsNil(o.BrokerAccountId) {
		var ret string
		return ret
	}
	return *o.BrokerAccountId
}

// GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBrokerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerAccountId) {
		return nil, false
	}
	return o.BrokerAccountId, true
}

// HasBrokerAccountId returns a boolean if a field has been set.
func (o *GetAccountResponse) HasBrokerAccountId() bool {
	if o != nil && !IsNil(o.BrokerAccountId) {
		return true
	}

	return false
}

// SetBrokerAccountId gets a reference to the given string and assigns it to the BrokerAccountId field.
func (o *GetAccountResponse) SetBrokerAccountId(v string) {
	o.BrokerAccountId = &v
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *GetAccountResponse) GetBrokerId() int32 {
	if o == nil || IsNil(o.BrokerId) {
		var ret int32
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetBrokerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BrokerId) {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *GetAccountResponse) HasBrokerId() bool {
	if o != nil && !IsNil(o.BrokerId) {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given int32 and assigns it to the BrokerId field.
func (o *GetAccountResponse) SetBrokerId(v int32) {
	o.BrokerId = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetAccountResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetAccountResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetAccountResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *GetAccountResponse) GetCurrencyId() int32 {
	if o == nil || IsNil(o.CurrencyId) {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *GetAccountResponse) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *GetAccountResponse) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetDeposits returns the Deposits field value if set, zero value otherwise.
func (o *GetAccountResponse) GetDeposits() []map[string]interface{} {
	if o == nil || IsNil(o.Deposits) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetDepositsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Deposits) {
		return nil, false
	}
	return o.Deposits, true
}

// HasDeposits returns a boolean if a field has been set.
func (o *GetAccountResponse) HasDeposits() bool {
	if o != nil && !IsNil(o.Deposits) {
		return true
	}

	return false
}

// SetDeposits gets a reference to the given []map[string]interface{} and assigns it to the Deposits field.
func (o *GetAccountResponse) SetDeposits(v []map[string]interface{}) {
	o.Deposits = v
}

// GetNetLiquidation returns the NetLiquidation field value if set, zero value otherwise.
func (o *GetAccountResponse) GetNetLiquidation() string {
	if o == nil || IsNil(o.NetLiquidation) {
		var ret string
		return ret
	}
	return *o.NetLiquidation
}

// GetNetLiquidationOk returns a tuple with the NetLiquidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetNetLiquidationOk() (*string, bool) {
	if o == nil || IsNil(o.NetLiquidation) {
		return nil, false
	}
	return o.NetLiquidation, true
}

// HasNetLiquidation returns a boolean if a field has been set.
func (o *GetAccountResponse) HasNetLiquidation() bool {
	if o != nil && !IsNil(o.NetLiquidation) {
		return true
	}

	return false
}

// SetNetLiquidation gets a reference to the given string and assigns it to the NetLiquidation field.
func (o *GetAccountResponse) SetNetLiquidation(v string) {
	o.NetLiquidation = &v
}

// GetOpenIpoOrders returns the OpenIpoOrders field value if set, zero value otherwise.
func (o *GetAccountResponse) GetOpenIpoOrders() []map[string]interface{} {
	if o == nil || IsNil(o.OpenIpoOrders) {
		var ret []map[string]interface{}
		return ret
	}
	return o.OpenIpoOrders
}

// GetOpenIpoOrdersOk returns a tuple with the OpenIpoOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetOpenIpoOrdersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.OpenIpoOrders) {
		return nil, false
	}
	return o.OpenIpoOrders, true
}

// HasOpenIpoOrders returns a boolean if a field has been set.
func (o *GetAccountResponse) HasOpenIpoOrders() bool {
	if o != nil && !IsNil(o.OpenIpoOrders) {
		return true
	}

	return false
}

// SetOpenIpoOrders gets a reference to the given []map[string]interface{} and assigns it to the OpenIpoOrders field.
func (o *GetAccountResponse) SetOpenIpoOrders(v []map[string]interface{}) {
	o.OpenIpoOrders = v
}

// GetOpenOrderSize returns the OpenOrderSize field value if set, zero value otherwise.
func (o *GetAccountResponse) GetOpenOrderSize() int32 {
	if o == nil || IsNil(o.OpenOrderSize) {
		var ret int32
		return ret
	}
	return *o.OpenOrderSize
}

// GetOpenOrderSizeOk returns a tuple with the OpenOrderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetOpenOrderSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenOrderSize) {
		return nil, false
	}
	return o.OpenOrderSize, true
}

// HasOpenOrderSize returns a boolean if a field has been set.
func (o *GetAccountResponse) HasOpenOrderSize() bool {
	if o != nil && !IsNil(o.OpenOrderSize) {
		return true
	}

	return false
}

// SetOpenOrderSize gets a reference to the given int32 and assigns it to the OpenOrderSize field.
func (o *GetAccountResponse) SetOpenOrderSize(v int32) {
	o.OpenOrderSize = &v
}

// GetOpenOrders returns the OpenOrders field value if set, zero value otherwise.
func (o *GetAccountResponse) GetOpenOrders() []map[string]interface{} {
	if o == nil || IsNil(o.OpenOrders) {
		var ret []map[string]interface{}
		return ret
	}
	return o.OpenOrders
}

// GetOpenOrdersOk returns a tuple with the OpenOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetOpenOrdersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.OpenOrders) {
		return nil, false
	}
	return o.OpenOrders, true
}

// HasOpenOrders returns a boolean if a field has been set.
func (o *GetAccountResponse) HasOpenOrders() bool {
	if o != nil && !IsNil(o.OpenOrders) {
		return true
	}

	return false
}

// SetOpenOrders gets a reference to the given []map[string]interface{} and assigns it to the OpenOrders field.
func (o *GetAccountResponse) SetOpenOrders(v []map[string]interface{}) {
	o.OpenOrders = v
}

// GetPdt returns the Pdt field value if set, zero value otherwise.
func (o *GetAccountResponse) GetPdt() bool {
	if o == nil || IsNil(o.Pdt) {
		var ret bool
		return ret
	}
	return *o.Pdt
}

// GetPdtOk returns a tuple with the Pdt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetPdtOk() (*bool, bool) {
	if o == nil || IsNil(o.Pdt) {
		return nil, false
	}
	return o.Pdt, true
}

// HasPdt returns a boolean if a field has been set.
func (o *GetAccountResponse) HasPdt() bool {
	if o != nil && !IsNil(o.Pdt) {
		return true
	}

	return false
}

// SetPdt gets a reference to the given bool and assigns it to the Pdt field.
func (o *GetAccountResponse) SetPdt(v bool) {
	o.Pdt = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *GetAccountResponse) GetPositions() []GetAccountResponsePositionsInner {
	if o == nil || IsNil(o.Positions) {
		var ret []GetAccountResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetPositionsOk() ([]GetAccountResponsePositionsInner, bool) {
	if o == nil || IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *GetAccountResponse) HasPositions() bool {
	if o != nil && !IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []GetAccountResponsePositionsInner and assigns it to the Positions field.
func (o *GetAccountResponse) SetPositions(v []GetAccountResponsePositionsInner) {
	o.Positions = v
}

// GetProfessional returns the Professional field value if set, zero value otherwise.
func (o *GetAccountResponse) GetProfessional() bool {
	if o == nil || IsNil(o.Professional) {
		var ret bool
		return ret
	}
	return *o.Professional
}

// GetProfessionalOk returns a tuple with the Professional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetProfessionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Professional) {
		return nil, false
	}
	return o.Professional, true
}

// HasProfessional returns a boolean if a field has been set.
func (o *GetAccountResponse) HasProfessional() bool {
	if o != nil && !IsNil(o.Professional) {
		return true
	}

	return false
}

// SetProfessional gets a reference to the given bool and assigns it to the Professional field.
func (o *GetAccountResponse) SetProfessional(v bool) {
	o.Professional = &v
}

// GetSecAccountId returns the SecAccountId field value if set, zero value otherwise.
func (o *GetAccountResponse) GetSecAccountId() int32 {
	if o == nil || IsNil(o.SecAccountId) {
		var ret int32
		return ret
	}
	return *o.SecAccountId
}

// GetSecAccountIdOk returns a tuple with the SecAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetSecAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SecAccountId) {
		return nil, false
	}
	return o.SecAccountId, true
}

// HasSecAccountId returns a boolean if a field has been set.
func (o *GetAccountResponse) HasSecAccountId() bool {
	if o != nil && !IsNil(o.SecAccountId) {
		return true
	}

	return false
}

// SetSecAccountId gets a reference to the given int32 and assigns it to the SecAccountId field.
func (o *GetAccountResponse) SetSecAccountId(v int32) {
	o.SecAccountId = &v
}

// GetShowUpgrade returns the ShowUpgrade field value if set, zero value otherwise.
func (o *GetAccountResponse) GetShowUpgrade() bool {
	if o == nil || IsNil(o.ShowUpgrade) {
		var ret bool
		return ret
	}
	return *o.ShowUpgrade
}

// GetShowUpgradeOk returns a tuple with the ShowUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetShowUpgradeOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowUpgrade) {
		return nil, false
	}
	return o.ShowUpgrade, true
}

// HasShowUpgrade returns a boolean if a field has been set.
func (o *GetAccountResponse) HasShowUpgrade() bool {
	if o != nil && !IsNil(o.ShowUpgrade) {
		return true
	}

	return false
}

// SetShowUpgrade gets a reference to the given bool and assigns it to the ShowUpgrade field.
func (o *GetAccountResponse) SetShowUpgrade(v bool) {
	o.ShowUpgrade = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *GetAccountResponse) GetTotalCost() string {
	if o == nil || IsNil(o.TotalCost) {
		var ret string
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetTotalCostOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *GetAccountResponse) HasTotalCost() bool {
	if o != nil && !IsNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given string and assigns it to the TotalCost field.
func (o *GetAccountResponse) SetTotalCost(v string) {
	o.TotalCost = &v
}

// GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field value if set, zero value otherwise.
func (o *GetAccountResponse) GetUnrealizedProfitLoss() string {
	if o == nil || IsNil(o.UnrealizedProfitLoss) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLoss
}

// GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetUnrealizedProfitLossOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLoss) {
		return nil, false
	}
	return o.UnrealizedProfitLoss, true
}

// HasUnrealizedProfitLoss returns a boolean if a field has been set.
func (o *GetAccountResponse) HasUnrealizedProfitLoss() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLoss) {
		return true
	}

	return false
}

// SetUnrealizedProfitLoss gets a reference to the given string and assigns it to the UnrealizedProfitLoss field.
func (o *GetAccountResponse) SetUnrealizedProfitLoss(v string) {
	o.UnrealizedProfitLoss = &v
}

// GetUnrealizedProfitLossBase returns the UnrealizedProfitLossBase field value if set, zero value otherwise.
func (o *GetAccountResponse) GetUnrealizedProfitLossBase() float32 {
	if o == nil || IsNil(o.UnrealizedProfitLossBase) {
		var ret float32
		return ret
	}
	return *o.UnrealizedProfitLossBase
}

// GetUnrealizedProfitLossBaseOk returns a tuple with the UnrealizedProfitLossBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetUnrealizedProfitLossBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLossBase) {
		return nil, false
	}
	return o.UnrealizedProfitLossBase, true
}

// HasUnrealizedProfitLossBase returns a boolean if a field has been set.
func (o *GetAccountResponse) HasUnrealizedProfitLossBase() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLossBase) {
		return true
	}

	return false
}

// SetUnrealizedProfitLossBase gets a reference to the given float32 and assigns it to the UnrealizedProfitLossBase field.
func (o *GetAccountResponse) SetUnrealizedProfitLossBase(v float32) {
	o.UnrealizedProfitLossBase = &v
}

// GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field value if set, zero value otherwise.
func (o *GetAccountResponse) GetUnrealizedProfitLossRate() string {
	if o == nil || IsNil(o.UnrealizedProfitLossRate) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLossRate
}

// GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetUnrealizedProfitLossRateOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLossRate) {
		return nil, false
	}
	return o.UnrealizedProfitLossRate, true
}

// HasUnrealizedProfitLossRate returns a boolean if a field has been set.
func (o *GetAccountResponse) HasUnrealizedProfitLossRate() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLossRate) {
		return true
	}

	return false
}

// SetUnrealizedProfitLossRate gets a reference to the given string and assigns it to the UnrealizedProfitLossRate field.
func (o *GetAccountResponse) SetUnrealizedProfitLossRate(v string) {
	o.UnrealizedProfitLossRate = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *GetAccountResponse) GetWarning() bool {
	if o == nil || IsNil(o.Warning) {
		var ret bool
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponse) GetWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *GetAccountResponse) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given bool and assigns it to the Warning field.
func (o *GetAccountResponse) SetWarning(v bool) {
	o.Warning = &v
}

func (o GetAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountMembers) {
		toSerialize["accountMembers"] = o.AccountMembers
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.Banners) {
		toSerialize["banners"] = o.Banners
	}
	if !IsNil(o.BrokerAccountId) {
		toSerialize["brokerAccountId"] = o.BrokerAccountId
	}
	if !IsNil(o.BrokerId) {
		toSerialize["brokerId"] = o.BrokerId
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currencyId"] = o.CurrencyId
	}
	if !IsNil(o.Deposits) {
		toSerialize["deposits"] = o.Deposits
	}
	if !IsNil(o.NetLiquidation) {
		toSerialize["netLiquidation"] = o.NetLiquidation
	}
	if !IsNil(o.OpenIpoOrders) {
		toSerialize["openIpoOrders"] = o.OpenIpoOrders
	}
	if !IsNil(o.OpenOrderSize) {
		toSerialize["openOrderSize"] = o.OpenOrderSize
	}
	if !IsNil(o.OpenOrders) {
		toSerialize["openOrders"] = o.OpenOrders
	}
	if !IsNil(o.Pdt) {
		toSerialize["pdt"] = o.Pdt
	}
	if !IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	if !IsNil(o.Professional) {
		toSerialize["professional"] = o.Professional
	}
	if !IsNil(o.SecAccountId) {
		toSerialize["secAccountId"] = o.SecAccountId
	}
	if !IsNil(o.ShowUpgrade) {
		toSerialize["showUpgrade"] = o.ShowUpgrade
	}
	if !IsNil(o.TotalCost) {
		toSerialize["totalCost"] = o.TotalCost
	}
	if !IsNil(o.UnrealizedProfitLoss) {
		toSerialize["unrealizedProfitLoss"] = o.UnrealizedProfitLoss
	}
	if !IsNil(o.UnrealizedProfitLossBase) {
		toSerialize["unrealizedProfitLossBase"] = o.UnrealizedProfitLossBase
	}
	if !IsNil(o.UnrealizedProfitLossRate) {
		toSerialize["unrealizedProfitLossRate"] = o.UnrealizedProfitLossRate
	}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return toSerialize, nil
}

type NullableGetAccountResponse struct {
	value *GetAccountResponse
	isSet bool
}

func (v NullableGetAccountResponse) Get() *GetAccountResponse {
	return v.value
}

func (v *NullableGetAccountResponse) Set(val *GetAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountResponse(val *GetAccountResponse) *NullableGetAccountResponse {
	return &NullableGetAccountResponse{value: val, isSet: true}
}

func (v NullableGetAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


