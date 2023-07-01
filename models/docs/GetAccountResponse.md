# GetAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMembers** | Pointer to [**[]GetAcccountRequestAccountMembersInner**](GetAcccountRequestAccountMembersInner.md) |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Banners** | Pointer to [**[]GetAcccountRequestBannersInner**](GetAcccountRequestBannersInner.md) |  | [optional] 
**BrokerAccountId** | Pointer to **string** |  | [optional] 
**BrokerId** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **int32** |  | [optional] 
**Deposits** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NetLiquidation** | Pointer to **string** |  | [optional] 
**OpenIpoOrders** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OpenOrderSize** | Pointer to **int32** |  | [optional] 
**OpenOrders** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Pdt** | Pointer to **bool** |  | [optional] 
**Positions** | Pointer to [**[]GetAccountResponsePositionsInner**](GetAccountResponsePositionsInner.md) |  | [optional] 
**Professional** | Pointer to **bool** |  | [optional] 
**SecAccountId** | Pointer to **int32** |  | [optional] 
**ShowUpgrade** | Pointer to **bool** |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossBase** | Pointer to **float32** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**Warning** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAccountResponse

`func NewGetAccountResponse() *GetAccountResponse`

NewGetAccountResponse instantiates a new GetAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponseWithDefaults

`func NewGetAccountResponseWithDefaults() *GetAccountResponse`

NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMembers

`func (o *GetAccountResponse) GetAccountMembers() []GetAcccountRequestAccountMembersInner`

GetAccountMembers returns the AccountMembers field if non-nil, zero value otherwise.

### GetAccountMembersOk

`func (o *GetAccountResponse) GetAccountMembersOk() (*[]GetAcccountRequestAccountMembersInner, bool)`

GetAccountMembersOk returns a tuple with the AccountMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMembers

`func (o *GetAccountResponse) SetAccountMembers(v []GetAcccountRequestAccountMembersInner)`

SetAccountMembers sets AccountMembers field to given value.

### HasAccountMembers

`func (o *GetAccountResponse) HasAccountMembers() bool`

HasAccountMembers returns a boolean if a field has been set.

### GetAccountType

`func (o *GetAccountResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetAccountResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetAccountResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetAccountResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBanners

`func (o *GetAccountResponse) GetBanners() []GetAcccountRequestBannersInner`

GetBanners returns the Banners field if non-nil, zero value otherwise.

### GetBannersOk

`func (o *GetAccountResponse) GetBannersOk() (*[]GetAcccountRequestBannersInner, bool)`

GetBannersOk returns a tuple with the Banners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanners

`func (o *GetAccountResponse) SetBanners(v []GetAcccountRequestBannersInner)`

SetBanners sets Banners field to given value.

### HasBanners

`func (o *GetAccountResponse) HasBanners() bool`

HasBanners returns a boolean if a field has been set.

### GetBrokerAccountId

`func (o *GetAccountResponse) GetBrokerAccountId() string`

GetBrokerAccountId returns the BrokerAccountId field if non-nil, zero value otherwise.

### GetBrokerAccountIdOk

`func (o *GetAccountResponse) GetBrokerAccountIdOk() (*string, bool)`

GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAccountId

`func (o *GetAccountResponse) SetBrokerAccountId(v string)`

SetBrokerAccountId sets BrokerAccountId field to given value.

### HasBrokerAccountId

`func (o *GetAccountResponse) HasBrokerAccountId() bool`

HasBrokerAccountId returns a boolean if a field has been set.

### GetBrokerId

`func (o *GetAccountResponse) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *GetAccountResponse) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *GetAccountResponse) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.

### HasBrokerId

`func (o *GetAccountResponse) HasBrokerId() bool`

HasBrokerId returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetAccountResponse) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetAccountResponse) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetAccountResponse) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetAccountResponse) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetDeposits

`func (o *GetAccountResponse) GetDeposits() []map[string]interface{}`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *GetAccountResponse) GetDepositsOk() (*[]map[string]interface{}, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *GetAccountResponse) SetDeposits(v []map[string]interface{})`

SetDeposits sets Deposits field to given value.

### HasDeposits

`func (o *GetAccountResponse) HasDeposits() bool`

HasDeposits returns a boolean if a field has been set.

### GetNetLiquidation

`func (o *GetAccountResponse) GetNetLiquidation() string`

GetNetLiquidation returns the NetLiquidation field if non-nil, zero value otherwise.

### GetNetLiquidationOk

`func (o *GetAccountResponse) GetNetLiquidationOk() (*string, bool)`

GetNetLiquidationOk returns a tuple with the NetLiquidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetLiquidation

`func (o *GetAccountResponse) SetNetLiquidation(v string)`

SetNetLiquidation sets NetLiquidation field to given value.

### HasNetLiquidation

`func (o *GetAccountResponse) HasNetLiquidation() bool`

HasNetLiquidation returns a boolean if a field has been set.

### GetOpenIpoOrders

`func (o *GetAccountResponse) GetOpenIpoOrders() []map[string]interface{}`

GetOpenIpoOrders returns the OpenIpoOrders field if non-nil, zero value otherwise.

### GetOpenIpoOrdersOk

`func (o *GetAccountResponse) GetOpenIpoOrdersOk() (*[]map[string]interface{}, bool)`

GetOpenIpoOrdersOk returns a tuple with the OpenIpoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIpoOrders

`func (o *GetAccountResponse) SetOpenIpoOrders(v []map[string]interface{})`

SetOpenIpoOrders sets OpenIpoOrders field to given value.

### HasOpenIpoOrders

`func (o *GetAccountResponse) HasOpenIpoOrders() bool`

HasOpenIpoOrders returns a boolean if a field has been set.

### GetOpenOrderSize

`func (o *GetAccountResponse) GetOpenOrderSize() int32`

GetOpenOrderSize returns the OpenOrderSize field if non-nil, zero value otherwise.

### GetOpenOrderSizeOk

`func (o *GetAccountResponse) GetOpenOrderSizeOk() (*int32, bool)`

GetOpenOrderSizeOk returns a tuple with the OpenOrderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderSize

`func (o *GetAccountResponse) SetOpenOrderSize(v int32)`

SetOpenOrderSize sets OpenOrderSize field to given value.

### HasOpenOrderSize

`func (o *GetAccountResponse) HasOpenOrderSize() bool`

HasOpenOrderSize returns a boolean if a field has been set.

### GetOpenOrders

`func (o *GetAccountResponse) GetOpenOrders() []map[string]interface{}`

GetOpenOrders returns the OpenOrders field if non-nil, zero value otherwise.

### GetOpenOrdersOk

`func (o *GetAccountResponse) GetOpenOrdersOk() (*[]map[string]interface{}, bool)`

GetOpenOrdersOk returns a tuple with the OpenOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrders

`func (o *GetAccountResponse) SetOpenOrders(v []map[string]interface{})`

SetOpenOrders sets OpenOrders field to given value.

### HasOpenOrders

`func (o *GetAccountResponse) HasOpenOrders() bool`

HasOpenOrders returns a boolean if a field has been set.

### GetPdt

`func (o *GetAccountResponse) GetPdt() bool`

GetPdt returns the Pdt field if non-nil, zero value otherwise.

### GetPdtOk

`func (o *GetAccountResponse) GetPdtOk() (*bool, bool)`

GetPdtOk returns a tuple with the Pdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdt

`func (o *GetAccountResponse) SetPdt(v bool)`

SetPdt sets Pdt field to given value.

### HasPdt

`func (o *GetAccountResponse) HasPdt() bool`

HasPdt returns a boolean if a field has been set.

### GetPositions

`func (o *GetAccountResponse) GetPositions() []GetAccountResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetAccountResponse) GetPositionsOk() (*[]GetAccountResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetAccountResponse) SetPositions(v []GetAccountResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetAccountResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetProfessional

`func (o *GetAccountResponse) GetProfessional() bool`

GetProfessional returns the Professional field if non-nil, zero value otherwise.

### GetProfessionalOk

`func (o *GetAccountResponse) GetProfessionalOk() (*bool, bool)`

GetProfessionalOk returns a tuple with the Professional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessional

`func (o *GetAccountResponse) SetProfessional(v bool)`

SetProfessional sets Professional field to given value.

### HasProfessional

`func (o *GetAccountResponse) HasProfessional() bool`

HasProfessional returns a boolean if a field has been set.

### GetSecAccountId

`func (o *GetAccountResponse) GetSecAccountId() int32`

GetSecAccountId returns the SecAccountId field if non-nil, zero value otherwise.

### GetSecAccountIdOk

`func (o *GetAccountResponse) GetSecAccountIdOk() (*int32, bool)`

GetSecAccountIdOk returns a tuple with the SecAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAccountId

`func (o *GetAccountResponse) SetSecAccountId(v int32)`

SetSecAccountId sets SecAccountId field to given value.

### HasSecAccountId

`func (o *GetAccountResponse) HasSecAccountId() bool`

HasSecAccountId returns a boolean if a field has been set.

### GetShowUpgrade

`func (o *GetAccountResponse) GetShowUpgrade() bool`

GetShowUpgrade returns the ShowUpgrade field if non-nil, zero value otherwise.

### GetShowUpgradeOk

`func (o *GetAccountResponse) GetShowUpgradeOk() (*bool, bool)`

GetShowUpgradeOk returns a tuple with the ShowUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUpgrade

`func (o *GetAccountResponse) SetShowUpgrade(v bool)`

SetShowUpgrade sets ShowUpgrade field to given value.

### HasShowUpgrade

`func (o *GetAccountResponse) HasShowUpgrade() bool`

HasShowUpgrade returns a boolean if a field has been set.

### GetTotalCost

`func (o *GetAccountResponse) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *GetAccountResponse) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *GetAccountResponse) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *GetAccountResponse) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountResponse) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountResponse) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountResponse) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountResponse) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossBase

`func (o *GetAccountResponse) GetUnrealizedProfitLossBase() float32`

GetUnrealizedProfitLossBase returns the UnrealizedProfitLossBase field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossBaseOk

`func (o *GetAccountResponse) GetUnrealizedProfitLossBaseOk() (*float32, bool)`

GetUnrealizedProfitLossBaseOk returns a tuple with the UnrealizedProfitLossBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossBase

`func (o *GetAccountResponse) SetUnrealizedProfitLossBase(v float32)`

SetUnrealizedProfitLossBase sets UnrealizedProfitLossBase field to given value.

### HasUnrealizedProfitLossBase

`func (o *GetAccountResponse) HasUnrealizedProfitLossBase() bool`

HasUnrealizedProfitLossBase returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountResponse) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountResponse) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountResponse) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountResponse) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetWarning

`func (o *GetAccountResponse) GetWarning() bool`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GetAccountResponse) GetWarningOk() (*bool, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GetAccountResponse) SetWarning(v bool)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GetAccountResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


