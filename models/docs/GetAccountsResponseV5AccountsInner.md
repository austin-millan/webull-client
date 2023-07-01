# GetAccountsResponseV5AccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecAccountId** | Pointer to **float32** |  | [optional] 
**BrokerId** | Pointer to **float32** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**BrokerAccountId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **float32** |  | [optional] 
**Pdt** | Pointer to **bool** |  | [optional] 
**Professional** | Pointer to **bool** |  | [optional] 
**ShowUpgrade** | Pointer to **bool** |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**NetLiquidation** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossBase** | Pointer to **string** |  | [optional] 
**Warning** | Pointer to **bool** |  | [optional] 
**RemindModifyPwd** | Pointer to **bool** |  | [optional] 
**AccountMembers** | Pointer to [**[]GetAccountsResponseV5AccountsInnerAccountMembersInner**](GetAccountsResponseV5AccountsInnerAccountMembersInner.md) |  | [optional] 
**OpenOrderSize** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetAccountsResponseV5AccountsInner

`func NewGetAccountsResponseV5AccountsInner() *GetAccountsResponseV5AccountsInner`

NewGetAccountsResponseV5AccountsInner instantiates a new GetAccountsResponseV5AccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsResponseV5AccountsInnerWithDefaults

`func NewGetAccountsResponseV5AccountsInnerWithDefaults() *GetAccountsResponseV5AccountsInner`

NewGetAccountsResponseV5AccountsInnerWithDefaults instantiates a new GetAccountsResponseV5AccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecAccountId

`func (o *GetAccountsResponseV5AccountsInner) GetSecAccountId() float32`

GetSecAccountId returns the SecAccountId field if non-nil, zero value otherwise.

### GetSecAccountIdOk

`func (o *GetAccountsResponseV5AccountsInner) GetSecAccountIdOk() (*float32, bool)`

GetSecAccountIdOk returns a tuple with the SecAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAccountId

`func (o *GetAccountsResponseV5AccountsInner) SetSecAccountId(v float32)`

SetSecAccountId sets SecAccountId field to given value.

### HasSecAccountId

`func (o *GetAccountsResponseV5AccountsInner) HasSecAccountId() bool`

HasSecAccountId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountsResponseV5AccountsInner) GetBrokerId() float32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountsResponseV5AccountsInner) GetBrokerIdOk() (*float32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountsResponseV5AccountsInner) SetBrokerId(v float32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountsResponseV5AccountsInner) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetAccountType

`func (o *GetAccountsResponseV5AccountsInner) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetAccountsResponseV5AccountsInner) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetAccountsResponseV5AccountsInner) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetAccountsResponseV5AccountsInner) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBrokerAccountId

`func (o *GetAccountsResponseV5AccountsInner) GetBrokerAccountId() string`

GetBrokerAccountId returns the BrokerAccountId field if non-nil, zero value otherwise.

### GetBrokerAccountIdOk

`func (o *GetAccountsResponseV5AccountsInner) GetBrokerAccountIdOk() (*string, bool)`

GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAccountId

`func (o *GetAccountsResponseV5AccountsInner) SetBrokerAccountId(v string)`

SetBrokerAccountId sets BrokerAccountId field to given value.

### HasBrokerAccountId

`func (o *GetAccountsResponseV5AccountsInner) HasBrokerAccountId() bool`

HasBrokerAccountId returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountsResponseV5AccountsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountsResponseV5AccountsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountsResponseV5AccountsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountsResponseV5AccountsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetAccountsResponseV5AccountsInner) GetCurrencyId() float32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetAccountsResponseV5AccountsInner) GetCurrencyIdOk() (*float32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetAccountsResponseV5AccountsInner) SetCurrencyId(v float32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetAccountsResponseV5AccountsInner) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetPdt

`func (o *GetAccountsResponseV5AccountsInner) GetPdt() bool`

GetPdt returns the Pdt field if non-nil, zero value otherwise.

### GetPdtOk

`func (o *GetAccountsResponseV5AccountsInner) GetPdtOk() (*bool, bool)`

GetPdtOk returns a tuple with the Pdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdt

`func (o *GetAccountsResponseV5AccountsInner) SetPdt(v bool)`

SetPdt sets Pdt field to given value.

### HasPdt

`func (o *GetAccountsResponseV5AccountsInner) HasPdt() bool`

HasPdt returns a boolean if a field has been set.

### GetProfessional

`func (o *GetAccountsResponseV5AccountsInner) GetProfessional() bool`

GetProfessional returns the Professional field if non-nil, zero value otherwise.

### GetProfessionalOk

`func (o *GetAccountsResponseV5AccountsInner) GetProfessionalOk() (*bool, bool)`

GetProfessionalOk returns a tuple with the Professional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessional

`func (o *GetAccountsResponseV5AccountsInner) SetProfessional(v bool)`

SetProfessional sets Professional field to given value.

### HasProfessional

`func (o *GetAccountsResponseV5AccountsInner) HasProfessional() bool`

HasProfessional returns a boolean if a field has been set.

### GetShowUpgrade

`func (o *GetAccountsResponseV5AccountsInner) GetShowUpgrade() bool`

GetShowUpgrade returns the ShowUpgrade field if non-nil, zero value otherwise.

### GetShowUpgradeOk

`func (o *GetAccountsResponseV5AccountsInner) GetShowUpgradeOk() (*bool, bool)`

GetShowUpgradeOk returns a tuple with the ShowUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUpgrade

`func (o *GetAccountsResponseV5AccountsInner) SetShowUpgrade(v bool)`

SetShowUpgrade sets ShowUpgrade field to given value.

### HasShowUpgrade

`func (o *GetAccountsResponseV5AccountsInner) HasShowUpgrade() bool`

HasShowUpgrade returns a boolean if a field has been set.

### GetTotalCost

`func (o *GetAccountsResponseV5AccountsInner) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *GetAccountsResponseV5AccountsInner) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *GetAccountsResponseV5AccountsInner) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *GetAccountsResponseV5AccountsInner) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetNetLiquidation

`func (o *GetAccountsResponseV5AccountsInner) GetNetLiquidation() string`

GetNetLiquidation returns the NetLiquidation field if non-nil, zero value otherwise.

### GetNetLiquidationOk

`func (o *GetAccountsResponseV5AccountsInner) GetNetLiquidationOk() (*string, bool)`

GetNetLiquidationOk returns a tuple with the NetLiquidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetLiquidation

`func (o *GetAccountsResponseV5AccountsInner) SetNetLiquidation(v string)`

SetNetLiquidation sets NetLiquidation field to given value.

### HasNetLiquidation

`func (o *GetAccountsResponseV5AccountsInner) HasNetLiquidation() bool`

HasNetLiquidation returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5AccountsInner) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountsResponseV5AccountsInner) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5AccountsInner) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5AccountsInner) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetUnrealizedProfitLossBase

`func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossBase() string`

GetUnrealizedProfitLossBase returns the UnrealizedProfitLossBase field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossBaseOk

`func (o *GetAccountsResponseV5AccountsInner) GetUnrealizedProfitLossBaseOk() (*string, bool)`

GetUnrealizedProfitLossBaseOk returns a tuple with the UnrealizedProfitLossBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossBase

`func (o *GetAccountsResponseV5AccountsInner) SetUnrealizedProfitLossBase(v string)`

SetUnrealizedProfitLossBase sets UnrealizedProfitLossBase field to given value.

### HasUnrealizedProfitLossBase

`func (o *GetAccountsResponseV5AccountsInner) HasUnrealizedProfitLossBase() bool`

HasUnrealizedProfitLossBase returns a boolean if a field has been set.

### GetWarning

`func (o *GetAccountsResponseV5AccountsInner) GetWarning() bool`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GetAccountsResponseV5AccountsInner) GetWarningOk() (*bool, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GetAccountsResponseV5AccountsInner) SetWarning(v bool)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GetAccountsResponseV5AccountsInner) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetRemindModifyPwd

`func (o *GetAccountsResponseV5AccountsInner) GetRemindModifyPwd() bool`

GetRemindModifyPwd returns the RemindModifyPwd field if non-nil, zero value otherwise.

### GetRemindModifyPwdOk

`func (o *GetAccountsResponseV5AccountsInner) GetRemindModifyPwdOk() (*bool, bool)`

GetRemindModifyPwdOk returns a tuple with the RemindModifyPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindModifyPwd

`func (o *GetAccountsResponseV5AccountsInner) SetRemindModifyPwd(v bool)`

SetRemindModifyPwd sets RemindModifyPwd field to given value.

### HasRemindModifyPwd

`func (o *GetAccountsResponseV5AccountsInner) HasRemindModifyPwd() bool`

HasRemindModifyPwd returns a boolean if a field has been set.

### GetAccountMembers

`func (o *GetAccountsResponseV5AccountsInner) GetAccountMembers() []GetAccountsResponseV5AccountsInnerAccountMembersInner`

GetAccountMembers returns the AccountMembers field if non-nil, zero value otherwise.

### GetAccountMembersOk

`func (o *GetAccountsResponseV5AccountsInner) GetAccountMembersOk() (*[]GetAccountsResponseV5AccountsInnerAccountMembersInner, bool)`

GetAccountMembersOk returns a tuple with the AccountMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMembers

`func (o *GetAccountsResponseV5AccountsInner) SetAccountMembers(v []GetAccountsResponseV5AccountsInnerAccountMembersInner)`

SetAccountMembers sets AccountMembers field to given value.

### HasAccountMembers

`func (o *GetAccountsResponseV5AccountsInner) HasAccountMembers() bool`

HasAccountMembers returns a boolean if a field has been set.

### GetOpenOrderSize

`func (o *GetAccountsResponseV5AccountsInner) GetOpenOrderSize() float32`

GetOpenOrderSize returns the OpenOrderSize field if non-nil, zero value otherwise.

### GetOpenOrderSizeOk

`func (o *GetAccountsResponseV5AccountsInner) GetOpenOrderSizeOk() (*float32, bool)`

GetOpenOrderSizeOk returns a tuple with the OpenOrderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderSize

`func (o *GetAccountsResponseV5AccountsInner) SetOpenOrderSize(v float32)`

SetOpenOrderSize sets OpenOrderSize field to given value.

### HasOpenOrderSize

`func (o *GetAccountsResponseV5AccountsInner) HasOpenOrderSize() bool`

HasOpenOrderSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


