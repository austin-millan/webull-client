# GetAccountsResponseV5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetLiquidation** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLoss** | Pointer to **string** |  | [optional] 
**UnrealizedProfitLossRate** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **float32** |  | [optional] 
**Accounts** | Pointer to [**[]GetAccountsResponseV5AccountsInner**](GetAccountsResponseV5AccountsInner.md) |  | [optional] 
**Positions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OpenOrders** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OpenOrderSize** | Pointer to **float32** |  | [optional] 
**Positions2** | Pointer to [**[]GetAccountsResponseV5Positions2Inner**](GetAccountsResponseV5Positions2Inner.md) |  | [optional] 
**OpenOrders2** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OpenIpoOrders** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetAccountsResponseV5

`func NewGetAccountsResponseV5() *GetAccountsResponseV5`

NewGetAccountsResponseV5 instantiates a new GetAccountsResponseV5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsResponseV5WithDefaults

`func NewGetAccountsResponseV5WithDefaults() *GetAccountsResponseV5`

NewGetAccountsResponseV5WithDefaults instantiates a new GetAccountsResponseV5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetLiquidation

`func (o *GetAccountsResponseV5) GetNetLiquidation() string`

GetNetLiquidation returns the NetLiquidation field if non-nil, zero value otherwise.

### GetNetLiquidationOk

`func (o *GetAccountsResponseV5) GetNetLiquidationOk() (*string, bool)`

GetNetLiquidationOk returns a tuple with the NetLiquidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetLiquidation

`func (o *GetAccountsResponseV5) SetNetLiquidation(v string)`

SetNetLiquidation sets NetLiquidation field to given value.

### HasNetLiquidation

`func (o *GetAccountsResponseV5) HasNetLiquidation() bool`

HasNetLiquidation returns a boolean if a field has been set.

### GetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5) GetUnrealizedProfitLoss() string`

GetUnrealizedProfitLoss returns the UnrealizedProfitLoss field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossOk

`func (o *GetAccountsResponseV5) GetUnrealizedProfitLossOk() (*string, bool)`

GetUnrealizedProfitLossOk returns a tuple with the UnrealizedProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLoss

`func (o *GetAccountsResponseV5) SetUnrealizedProfitLoss(v string)`

SetUnrealizedProfitLoss sets UnrealizedProfitLoss field to given value.

### HasUnrealizedProfitLoss

`func (o *GetAccountsResponseV5) HasUnrealizedProfitLoss() bool`

HasUnrealizedProfitLoss returns a boolean if a field has been set.

### GetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5) GetUnrealizedProfitLossRate() string`

GetUnrealizedProfitLossRate returns the UnrealizedProfitLossRate field if non-nil, zero value otherwise.

### GetUnrealizedProfitLossRateOk

`func (o *GetAccountsResponseV5) GetUnrealizedProfitLossRateOk() (*string, bool)`

GetUnrealizedProfitLossRateOk returns a tuple with the UnrealizedProfitLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5) SetUnrealizedProfitLossRate(v string)`

SetUnrealizedProfitLossRate sets UnrealizedProfitLossRate field to given value.

### HasUnrealizedProfitLossRate

`func (o *GetAccountsResponseV5) HasUnrealizedProfitLossRate() bool`

HasUnrealizedProfitLossRate returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAccountsResponseV5) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccountsResponseV5) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccountsResponseV5) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAccountsResponseV5) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyId

`func (o *GetAccountsResponseV5) GetCurrencyId() float32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GetAccountsResponseV5) GetCurrencyIdOk() (*float32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GetAccountsResponseV5) SetCurrencyId(v float32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GetAccountsResponseV5) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetAccounts

`func (o *GetAccountsResponseV5) GetAccounts() []GetAccountsResponseV5AccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetAccountsResponseV5) GetAccountsOk() (*[]GetAccountsResponseV5AccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetAccountsResponseV5) SetAccounts(v []GetAccountsResponseV5AccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetAccountsResponseV5) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetPositions

`func (o *GetAccountsResponseV5) GetPositions() []map[string]interface{}`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetAccountsResponseV5) GetPositionsOk() (*[]map[string]interface{}, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetAccountsResponseV5) SetPositions(v []map[string]interface{})`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetAccountsResponseV5) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetOpenOrders

`func (o *GetAccountsResponseV5) GetOpenOrders() []map[string]interface{}`

GetOpenOrders returns the OpenOrders field if non-nil, zero value otherwise.

### GetOpenOrdersOk

`func (o *GetAccountsResponseV5) GetOpenOrdersOk() (*[]map[string]interface{}, bool)`

GetOpenOrdersOk returns a tuple with the OpenOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrders

`func (o *GetAccountsResponseV5) SetOpenOrders(v []map[string]interface{})`

SetOpenOrders sets OpenOrders field to given value.

### HasOpenOrders

`func (o *GetAccountsResponseV5) HasOpenOrders() bool`

HasOpenOrders returns a boolean if a field has been set.

### GetOpenOrderSize

`func (o *GetAccountsResponseV5) GetOpenOrderSize() float32`

GetOpenOrderSize returns the OpenOrderSize field if non-nil, zero value otherwise.

### GetOpenOrderSizeOk

`func (o *GetAccountsResponseV5) GetOpenOrderSizeOk() (*float32, bool)`

GetOpenOrderSizeOk returns a tuple with the OpenOrderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderSize

`func (o *GetAccountsResponseV5) SetOpenOrderSize(v float32)`

SetOpenOrderSize sets OpenOrderSize field to given value.

### HasOpenOrderSize

`func (o *GetAccountsResponseV5) HasOpenOrderSize() bool`

HasOpenOrderSize returns a boolean if a field has been set.

### GetPositions2

`func (o *GetAccountsResponseV5) GetPositions2() []GetAccountsResponseV5Positions2Inner`

GetPositions2 returns the Positions2 field if non-nil, zero value otherwise.

### GetPositions2Ok

`func (o *GetAccountsResponseV5) GetPositions2Ok() (*[]GetAccountsResponseV5Positions2Inner, bool)`

GetPositions2Ok returns a tuple with the Positions2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions2

`func (o *GetAccountsResponseV5) SetPositions2(v []GetAccountsResponseV5Positions2Inner)`

SetPositions2 sets Positions2 field to given value.

### HasPositions2

`func (o *GetAccountsResponseV5) HasPositions2() bool`

HasPositions2 returns a boolean if a field has been set.

### GetOpenOrders2

`func (o *GetAccountsResponseV5) GetOpenOrders2() []map[string]interface{}`

GetOpenOrders2 returns the OpenOrders2 field if non-nil, zero value otherwise.

### GetOpenOrders2Ok

`func (o *GetAccountsResponseV5) GetOpenOrders2Ok() (*[]map[string]interface{}, bool)`

GetOpenOrders2Ok returns a tuple with the OpenOrders2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrders2

`func (o *GetAccountsResponseV5) SetOpenOrders2(v []map[string]interface{})`

SetOpenOrders2 sets OpenOrders2 field to given value.

### HasOpenOrders2

`func (o *GetAccountsResponseV5) HasOpenOrders2() bool`

HasOpenOrders2 returns a boolean if a field has been set.

### GetOpenIpoOrders

`func (o *GetAccountsResponseV5) GetOpenIpoOrders() []map[string]interface{}`

GetOpenIpoOrders returns the OpenIpoOrders field if non-nil, zero value otherwise.

### GetOpenIpoOrdersOk

`func (o *GetAccountsResponseV5) GetOpenIpoOrdersOk() (*[]map[string]interface{}, bool)`

GetOpenIpoOrdersOk returns a tuple with the OpenIpoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIpoOrders

`func (o *GetAccountsResponseV5) SetOpenIpoOrders(v []map[string]interface{})`

SetOpenIpoOrders sets OpenIpoOrders field to given value.

### HasOpenIpoOrders

`func (o *GetAccountsResponseV5) HasOpenIpoOrders() bool`

HasOpenIpoOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


