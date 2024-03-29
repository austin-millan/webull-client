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

// checks if the GetAccountsResponseV5AccountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountsResponseV5AccountsInner{}

// GetAccountsResponseV5AccountsInner struct for GetAccountsResponseV5AccountsInner
type GetAccountsResponseV5AccountsInner struct {
	SecAccountId *float32 `json:"secAccountId,omitempty"`
	BrokerId *float32 `json:"brokerId,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	BrokerAccountId *string `json:"brokerAccountId,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrencyId *float32 `json:"currencyId,omitempty"`
	Pdt *bool `json:"pdt,omitempty"`
	Professional *bool `json:"professional,omitempty"`
	ShowUpgrade *bool `json:"showUpgrade,omitempty"`
	TotalCost *string `json:"totalCost,omitempty"`
	NetLiquidation *string `json:"netLiquidation,omitempty"`
	UnrealizedProfitLoss *string `json:"unrealizedProfitLoss,omitempty"`
	UnrealizedProfitLossRate *string `json:"unrealizedProfitLossRate,omitempty"`
	UnrealizedProfitLossBase *string `json:"unrealizedProfitLossBase,omitempty"`
	Warning *bool `json:"warning,omitempty"`
	RemindModifyPwd *bool `json:"remindModifyPwd,omitempty"`
	AccountMembers []GetAccountsResponseV5AccountsInnerAccountMembersInner `json:"accountMembers,omitempty"`
	OpenOrderSize *float32 `json:"openOrderSize,omitempty"`
}

// NewGetAccountsResponseV5AccountsInner instantiates a new GetAccountsResponseV5AccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountsResponseV5AccountsInner() *GetAccountsResponseV5AccountsInner {
	this := GetAccountsResponseV5AccountsInner{}
	return &this
}

// NewGetAccountsResponseV5AccountsInnerWithDefaults instantiates a new GetAccountsResponseV5AccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountsResponseV5AccountsInnerWithDefaults() *GetAccountsResponseV5AccountsInner {
	this := GetAccountsResponseV5AccountsInner{}
	return &this
}

// GetSecAccountId returns the SecAccountId field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetSecAccountId() float32 {
	if o == nil || IsNil(o.SecAccountId) {
		var ret float32
		return ret
	}
	return *o.SecAccountId
}

// GetSecAccountIdOk returns a tuple with the SecAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetSecAccountIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SecAccountId) {
		return nil, false
	}
	return o.SecAccountId, true
}

// HasSecAccountId returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasSecAccountId() bool {
	if o != nil && !IsNil(o.SecAccountId) {
		return true
	}

	return false
}

// SetSecAccountId gets a reference to the given float32 and assigns it to the SecAccountId field.
func (o *GetAccountsResponseV5AccountsInner) SetSecAccountId(v float32) {
	o.SecAccountId = &v
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetBrokerId() float32 {
	if o == nil || IsNil(o.BrokerId) {
		var ret float32
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetBrokerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.BrokerId) {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasBrokerId() bool {
	if o != nil && !IsNil(o.BrokerId) {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given float32 and assigns it to the BrokerId field.
func (o *GetAccountsResponseV5AccountsInner) SetBrokerId(v float32) {
	o.BrokerId = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *GetAccountsResponseV5AccountsInner) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBrokerAccountId returns the BrokerAccountId field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetBrokerAccountId() string {
	if o == nil || IsNil(o.BrokerAccountId) {
		var ret string
		return ret
	}
	return *o.BrokerAccountId
}

// GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetBrokerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerAccountId) {
		return nil, false
	}
	return o.BrokerAccountId, true
}

// HasBrokerAccountId returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasBrokerAccountId() bool {
	if o != nil && !IsNil(o.BrokerAccountId) {
		return true
	}

	return false
}

// SetBrokerAccountId gets a reference to the given string and assigns it to the BrokerAccountId field.
func (o *GetAccountsResponseV5AccountsInner) SetBrokerAccountId(v string) {
	o.BrokerAccountId = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetAccountsResponseV5AccountsInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetCurrencyId() float32 {
	if o == nil || IsNil(o.CurrencyId) {
		var ret float32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetCurrencyIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given float32 and assigns it to the CurrencyId field.
func (o *GetAccountsResponseV5AccountsInner) SetCurrencyId(v float32) {
	o.CurrencyId = &v
}

// GetPdt returns the Pdt field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetPdt() bool {
	if o == nil || IsNil(o.Pdt) {
		var ret bool
		return ret
	}
	return *o.Pdt
}

// GetPdtOk returns a tuple with the Pdt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetPdtOk() (*bool, bool) {
	if o == nil || IsNil(o.Pdt) {
		return nil, false
	}
	return o.Pdt, true
}

// HasPdt returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasPdt() bool {
	if o != nil && !IsNil(o.Pdt) {
		return true
	}

	return false
}

// SetPdt gets a reference to the given bool and assigns it to the Pdt field.
func (o *GetAccountsResponseV5AccountsInner) SetPdt(v bool) {
	o.Pdt = &v
}

// GetProfessional returns the Professional field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetProfessional() bool {
	if o == nil || IsNil(o.Professional) {
		var ret bool
		return ret
	}
	return *o.Professional
}

// GetProfessionalOk returns a tuple with the Professional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetProfessionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Professional) {
		return nil, false
	}
	return o.Professional, true
}

// HasProfessional returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasProfessional() bool {
	if o != nil && !IsNil(o.Professional) {
		return true
	}

	return false
}

// SetProfessional gets a reference to the given bool and assigns it to the Professional field.
func (o *GetAccountsResponseV5AccountsInner) SetProfessional(v bool) {
	o.Professional = &v
}

// GetShowUpgrade returns the ShowUpgrade field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetShowUpgrade() bool {
	if o == nil || IsNil(o.ShowUpgrade) {
		var ret bool
		return ret
	}
	return *o.ShowUpgrade
}

// GetShowUpgradeOk returns a tuple with the ShowUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetShowUpgradeOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowUpgrade) {
		return nil, false
	}
	return o.ShowUpgrade, true
}

// HasShowUpgrade returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasShowUpgrade() bool {
	if o != nil && !IsNil(o.ShowUpgrade) {
		return true
	}

	return false
}

// SetShowUpgrade gets a reference to the given bool and assigns it to the ShowUpgrade field.
func (o *GetAccountsResponseV5AccountsInner) SetShowUpgrade(v bool) {
	o.ShowUpgrade = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetTotalCost() string {
	if o == nil || IsNil(o.TotalCost) {
		var ret string
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetTotalCostOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasTotalCost() bool {
	if o != nil && !IsNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given string and assigns it to the TotalCost field.
func (o *GetAccountsResponseV5AccountsInner) SetTotalCost(v string) {
	o.TotalCost = &v
}

// GetNetLiquidation returns the NetLiquidation field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetNetLiquidation() string {
	if o == nil || IsNil(o.NetLiquidation) {
		var ret string
		return ret
	}
	return *o.NetLiquidation
}

// GetNetLiquidationOk returns a tuple with the NetLiquidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetNetLiquidationOk() (*string, bool) {
	if o == nil || IsNil(o.NetLiquidation) {
		return nil, false
	}
	return o.NetLiquidation, true
}

// HasNetLiquidation returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasNetLiquidation() bool {
	if o != nil && !IsNil(o.NetLiquidation) {
		return true
	}

	return false
}

// SetNetLiquidation gets a reference to the given string and assigns it to the NetLiquidation field.
func (o *GetAccountsResponseV5AccountsInner) SetNetLiquidation(v string) {
	o.NetLiquidation = &v
}

// GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLoss() string {
	if o == nil || IsNil(o.UnrealizedProfitLoss) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLoss
}

// GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLoss) {
		return nil, false
	}
	return o.UnrealizedProfitLoss, true
}

// HasUnrealizedProfitLoss returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasUnrealizedProfitLoss() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLoss) {
		return true
	}

	return false
}

// SetUnrealizedProfitLoss gets a reference to the given string and assigns it to the UnrealizedProfitLoss field.
func (o *GetAccountsResponseV5AccountsInner) SetUnrealizedProfitLoss(v string) {
	o.UnrealizedProfitLoss = &v
}

// GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossRate() string {
	if o == nil || IsNil(o.UnrealizedProfitLossRate) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLossRate
}

// GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossRateOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLossRate) {
		return nil, false
	}
	return o.UnrealizedProfitLossRate, true
}

// HasUnrealizedProfitLossRate returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasUnrealizedProfitLossRate() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLossRate) {
		return true
	}

	return false
}

// SetUnrealizedProfitLossRate gets a reference to the given string and assigns it to the UnrealizedProfitLossRate field.
func (o *GetAccountsResponseV5AccountsInner) SetUnrealizedProfitLossRate(v string) {
	o.UnrealizedProfitLossRate = &v
}

// GetUnrealizedProfitLossBase returns the UnrealizedProfitLossBase field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossBase() string {
	if o == nil || IsNil(o.UnrealizedProfitLossBase) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfitLossBase
}

// GetUnrealizedProfitLossBaseOk returns a tuple with the UnrealizedProfitLossBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossBaseOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfitLossBase) {
		return nil, false
	}
	return o.UnrealizedProfitLossBase, true
}

// HasUnrealizedProfitLossBase returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasUnrealizedProfitLossBase() bool {
	if o != nil && !IsNil(o.UnrealizedProfitLossBase) {
		return true
	}

	return false
}

// SetUnrealizedProfitLossBase gets a reference to the given string and assigns it to the UnrealizedProfitLossBase field.
func (o *GetAccountsResponseV5AccountsInner) SetUnrealizedProfitLossBase(v string) {
	o.UnrealizedProfitLossBase = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetWarning() bool {
	if o == nil || IsNil(o.Warning) {
		var ret bool
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given bool and assigns it to the Warning field.
func (o *GetAccountsResponseV5AccountsInner) SetWarning(v bool) {
	o.Warning = &v
}

// GetRemindModifyPwd returns the RemindModifyPwd field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetRemindModifyPwd() bool {
	if o == nil || IsNil(o.RemindModifyPwd) {
		var ret bool
		return ret
	}
	return *o.RemindModifyPwd
}

// GetRemindModifyPwdOk returns a tuple with the RemindModifyPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetRemindModifyPwdOk() (*bool, bool) {
	if o == nil || IsNil(o.RemindModifyPwd) {
		return nil, false
	}
	return o.RemindModifyPwd, true
}

// HasRemindModifyPwd returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasRemindModifyPwd() bool {
	if o != nil && !IsNil(o.RemindModifyPwd) {
		return true
	}

	return false
}

// SetRemindModifyPwd gets a reference to the given bool and assigns it to the RemindModifyPwd field.
func (o *GetAccountsResponseV5AccountsInner) SetRemindModifyPwd(v bool) {
	o.RemindModifyPwd = &v
}

// GetAccountMembers returns the AccountMembers field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetAccountMembers() []GetAccountsResponseV5AccountsInnerAccountMembersInner {
	if o == nil || IsNil(o.AccountMembers) {
		var ret []GetAccountsResponseV5AccountsInnerAccountMembersInner
		return ret
	}
	return o.AccountMembers
}

// GetAccountMembersOk returns a tuple with the AccountMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetAccountMembersOk() ([]GetAccountsResponseV5AccountsInnerAccountMembersInner, bool) {
	if o == nil || IsNil(o.AccountMembers) {
		return nil, false
	}
	return o.AccountMembers, true
}

// HasAccountMembers returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasAccountMembers() bool {
	if o != nil && !IsNil(o.AccountMembers) {
		return true
	}

	return false
}

// SetAccountMembers gets a reference to the given []GetAccountsResponseV5AccountsInnerAccountMembersInner and assigns it to the AccountMembers field.
func (o *GetAccountsResponseV5AccountsInner) SetAccountMembers(v []GetAccountsResponseV5AccountsInnerAccountMembersInner) {
	o.AccountMembers = v
}

// GetOpenOrderSize returns the OpenOrderSize field value if set, zero value otherwise.
func (o *GetAccountsResponseV5AccountsInner) GetOpenOrderSize() float32 {
	if o == nil || IsNil(o.OpenOrderSize) {
		var ret float32
		return ret
	}
	return *o.OpenOrderSize
}

// GetOpenOrderSizeOk returns a tuple with the OpenOrderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsResponseV5AccountsInner) GetOpenOrderSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.OpenOrderSize) {
		return nil, false
	}
	return o.OpenOrderSize, true
}

// HasOpenOrderSize returns a boolean if a field has been set.
func (o *GetAccountsResponseV5AccountsInner) HasOpenOrderSize() bool {
	if o != nil && !IsNil(o.OpenOrderSize) {
		return true
	}

	return false
}

// SetOpenOrderSize gets a reference to the given float32 and assigns it to the OpenOrderSize field.
func (o *GetAccountsResponseV5AccountsInner) SetOpenOrderSize(v float32) {
	o.OpenOrderSize = &v
}

func (o GetAccountsResponseV5AccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountsResponseV5AccountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecAccountId) {
		toSerialize["secAccountId"] = o.SecAccountId
	}
	if !IsNil(o.BrokerId) {
		toSerialize["brokerId"] = o.BrokerId
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.BrokerAccountId) {
		toSerialize["brokerAccountId"] = o.BrokerAccountId
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currencyId"] = o.CurrencyId
	}
	if !IsNil(o.Pdt) {
		toSerialize["pdt"] = o.Pdt
	}
	if !IsNil(o.Professional) {
		toSerialize["professional"] = o.Professional
	}
	if !IsNil(o.ShowUpgrade) {
		toSerialize["showUpgrade"] = o.ShowUpgrade
	}
	if !IsNil(o.TotalCost) {
		toSerialize["totalCost"] = o.TotalCost
	}
	if !IsNil(o.NetLiquidation) {
		toSerialize["netLiquidation"] = o.NetLiquidation
	}
	if !IsNil(o.UnrealizedProfitLoss) {
		toSerialize["unrealizedProfitLoss"] = o.UnrealizedProfitLoss
	}
	if !IsNil(o.UnrealizedProfitLossRate) {
		toSerialize["unrealizedProfitLossRate"] = o.UnrealizedProfitLossRate
	}
	if !IsNil(o.UnrealizedProfitLossBase) {
		toSerialize["unrealizedProfitLossBase"] = o.UnrealizedProfitLossBase
	}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	if !IsNil(o.RemindModifyPwd) {
		toSerialize["remindModifyPwd"] = o.RemindModifyPwd
	}
	if !IsNil(o.AccountMembers) {
		toSerialize["accountMembers"] = o.AccountMembers
	}
	if !IsNil(o.OpenOrderSize) {
		toSerialize["openOrderSize"] = o.OpenOrderSize
	}
	return toSerialize, nil
}

type NullableGetAccountsResponseV5AccountsInner struct {
	value *GetAccountsResponseV5AccountsInner
	isSet bool
}

func (v NullableGetAccountsResponseV5AccountsInner) Get() *GetAccountsResponseV5AccountsInner {
	return v.value
}

func (v *NullableGetAccountsResponseV5AccountsInner) Set(val *GetAccountsResponseV5AccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountsResponseV5AccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountsResponseV5AccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountsResponseV5AccountsInner(val *GetAccountsResponseV5AccountsInner) *NullableGetAccountsResponseV5AccountsInner {
	return &NullableGetAccountsResponseV5AccountsInner{value: val, isSet: true}
}

func (v NullableGetAccountsResponseV5AccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountsResponseV5AccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


